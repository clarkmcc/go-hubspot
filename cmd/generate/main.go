package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	// directoryUrl is HubSpot's public API spec catalog. The previous
	// "api-catalog-public/v1/apis" endpoint was retired and now returns 404; the
	// current catalog lives here and uses the date-based versioning scheme.
	directoryUrl = "https://api.hubspot.com/public/api/spec/v2/specs"

	// targetVersion selects which date-based API release to generate. HubSpot
	// moved from semantic (v1..v4) to date-based (YYYY-MM) versions; legacy
	// versions are unsupported from Sep 2027. 2026-09 is the current target.
	targetVersion = "2026-09"

	// targetAPIs is the allowlist of catalog API names to generate, mapped to the
	// Go package name they are emitted under. We generate the CRM-core surface
	// rather than all ~170 catalog entries. The generic namer (deriveName) keeps
	// method names collision-free for any API added here, so extending coverage
	// is just a matter of adding entries.
	targetAPIs = map[string]string{
		"Contacts":       "contacts",
		"Companies":      "companies",
		"Deals":          "deals",
		"Tickets":        "tickets",
		"Products":       "products",
		"Line Items":     "lineitems",
		"Quotes":         "quotes",
		"Pipelines":      "pipelines",
		"Properties":     "properties",
		"Crm Owners":     "owners",
		"Custom Objects": "objects",
		"Associations":   "associations",
		"Schemas":        "schemas",
	}
)

// nameOverrides lets a specific operation (matched by HTTP method + a regexp on
// its path) be given an explicit Go method name, overriding the generic namer
// below. Use this to pin stable names that downstream code depends on, or to
// improve a generically-derived name. Keyed by catalog API name.
//
// The generic namer (see deriveName) already produces compiling, collision-free
// names for every operation, so overrides are optional — they exist to keep
// ergonomic, stable method names for the operations this org consumes heavily.
var nameOverrides = map[string][]nameOverride{
	"Contacts": {
		{method: "post", path: `/contacts/search$`, name: "Search"},
		{method: "post", path: `/contacts/batch/create$`, name: "BatchCreate"},
		{method: "post", path: `/contacts/batch/update$`, name: "BatchUpdate"},
		{method: "post", path: `/contacts/batch/archive$`, name: "BatchArchive"},
		{method: "post", path: `/contacts/batch/read$`, name: "BatchRead"},
	},
	"Custom Objects": {
		{method: "post", path: `/\{objectType\}/search$`, name: "Search"},
		{method: "post", path: `/\{objectType\}/batch/create$`, name: "BatchCreate"},
		{method: "post", path: `/\{objectType\}/batch/update$`, name: "BatchUpdate"},
		{method: "post", path: `/\{objectType\}/batch/archive$`, name: "BatchArchive"},
		{method: "post", path: `/\{objectType\}/batch/read$`, name: "BatchRead"},
	},
	"Associations": {
		// The legacy "/types" read is gone in date-based versions; association
		// type ids are now read from "/labels" (GET, tag Basic -> BasicAPIService).
		{method: "get", path: `/\{fromObjectType\}/\{toObjectType\}/labels$`, name: "GetLabels"},
		{method: "post", path: `/\{fromObjectType\}/\{toObjectType\}/batch/create$`, name: "BatchCreate"},
		{method: "post", path: `/\{fromObjectType\}/\{toObjectType\}/batch/archive$`, name: "BatchArchive"},
	},
}

// nameOverride pins an explicit Go method name for an operation identified by
// HTTP method and a regexp matched against the operation's path.
type nameOverride struct {
	method string
	path   string
	name   string
}

// schemaOpenAPI contains the OpenAPI schema definition.
type schemaOpenAPI struct {
	Info struct {
		Title   string `json:"title"`
		Version string `json:"version"`
	} `json:"info"`
}

// directory is the schema for the https://api.hubspot.com/public/api/spec/v2/specs
// catalog page which lists all publicly available APIs and their releases.
type directory struct {
	Results []directoryAPI `json:"results"`
}

// directoryAPI is a single API in the catalog, with one entry per released version.
type directoryAPI struct {
	Name     string `json:"name"`
	Group    string `json:"group"`
	Releases []struct {
		ReleaseVersion string `json:"releaseVersion"`
		Stage          string `json:"stage"`
		OpenAPI        string `json:"openApi"`
	} `json:"releases"`
}

// openAPIFor returns the OpenAPI spec URL for the requested release version of
// this API, or "" if that version is not published.
func (d directoryAPI) openAPIFor(version string) string {
	for _, rel := range d.Releases {
		if rel.ReleaseVersion == version {
			return rel.OpenAPI
		}
	}
	return ""
}

// retrieveSchema downloads a schema from the given url.
// It returns the schema text as a string.
func retrieveSchema(url string) (string, error) {
	res, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	schema, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(schema), nil
}

// preProcessSchema rewrites operation IDs in the OpenAPI spec so that
// openapi-generator produces clean, collision-free Go method names, then saves
// the result to ./schema/<group>.json and returns the path.
//
// HubSpot's date-based specs expose very long auto-generated operationIds (e.g.
// "post-/crm/objects/2026-09/contacts/search_/crm/objects/2026-09-beta/{objectType}/search")
// which would otherwise become unusable Go method names. For each operation we
// pick a name as follows:
//  1. an explicit entry in nameOverrides (matched by method + path regexp), else
//  2. a generic name derived from the HTTP verb + meaningful path segments.
//
// Because openapi-generator groups operations into one <Tag>APIService per tag,
// method-name collisions only matter within a tag. We resolve collisions per tag
// in three escalating passes (base name -> name including path params -> numeric
// suffix) so the output always compiles.
func preProcessSchema(group string, schema string) (string, error) {
	var doc map[string]any
	if err := json.Unmarshal([]byte(schema), &doc); err != nil {
		return "", fmt.Errorf("parsing %s spec: %w", group, err)
	}

	paths, _ := doc["paths"].(map[string]any)

	// Compile overrides for this API once.
	type compiledOverride struct {
		method string
		re     *regexp.Regexp
		name   string
	}
	var overrides []compiledOverride
	for _, o := range nameOverrides[group] {
		overrides = append(overrides, compiledOverride{
			method: strings.ToLower(o.method),
			re:     regexp.MustCompile(o.path),
			name:   o.name,
		})
	}

	// Collect operations grouped by tag, preserving a deterministic order.
	type op struct {
		node     map[string]any
		method   string
		path     string
		tag      string
		override string // explicit name, or "" to use the generic namer
	}
	var ops []op
	var pathKeys []string
	for p := range paths {
		pathKeys = append(pathKeys, p)
	}
	sort.Strings(pathKeys)
	for _, p := range pathKeys {
		methods, _ := paths[p].(map[string]any)
		var mKeys []string
		for m := range methods {
			mKeys = append(mKeys, m)
		}
		sort.Strings(mKeys)
		for _, m := range mKeys {
			if !isHTTPMethod(m) {
				continue
			}
			node, ok := methods[m].(map[string]any)
			if !ok {
				continue
			}
			tag := "Default"
			if tags, ok := node["tags"].([]any); ok && len(tags) > 0 {
				if t, ok := tags[0].(string); ok {
					tag = t
				}
			}
			ov := ""
			for _, o := range overrides {
				if o.method == strings.ToLower(m) && o.re.MatchString(p) {
					ov = o.name
					break
				}
			}
			ops = append(ops, op{node: node, method: strings.ToLower(m), path: p, tag: tag, override: ov})
		}
	}

	// Assign names with collision resolution. openapi-generator derives both the
	// method name AND a package-level request type (Api<Name>Request) from the
	// operationId, so names must be unique across the WHOLE package, not just
	// within a tag/APIService. We therefore dedupe against a single set.
	used := map[string]bool{}
	// Pass 0: explicit overrides claim their names first.
	for i := range ops {
		o := &ops[i]
		if o.override != "" {
			o.node["operationId"] = o.override
			used[o.override] = true
		}
	}
	// Passes 1-3 for the rest.
	for i := range ops {
		o := &ops[i]
		if o.override != "" {
			continue
		}
		candidates := []string{
			deriveName(o.method, o.path, false),
			deriveName(o.method, o.path, true),
		}
		name := ""
		for _, c := range candidates {
			if c != "" && !used[c] {
				name = c
				break
			}
		}
		if name == "" {
			base := deriveName(o.method, o.path, true)
			for n := 2; ; n++ {
				c := base + strconv.Itoa(n)
				if !used[c] {
					name = c
					break
				}
			}
		}
		o.node["operationId"] = name
		used[name] = true
	}

	out, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return "", fmt.Errorf("re-encoding %s spec: %w", group, err)
	}

	filename := "./schema/" + group + ".json"
	if err := ioutil.WriteFile(filename, out, 0644); err != nil {
		return "", err
	}
	return filename, nil
}

// isHTTPMethod reports whether key is an OpenAPI operation verb (not e.g.
// "parameters" or "summary" which also appear under a path item).
func isHTTPMethod(key string) bool {
	switch strings.ToLower(key) {
	case "get", "post", "put", "patch", "delete", "head", "options", "trace":
		return true
	}
	return false
}

var (
	nameVerb        = map[string]string{"get": "Get", "post": "Post", "put": "Put", "patch": "Patch", "delete": "Delete", "head": "Head", "options": "Options", "trace": "Trace"}
	versionSegmentR = regexp.MustCompile(`^(v\d+|\d{4}-\d{2}(-beta)?)$`)
	// nameDropSegments are leading "area" path segments that add no meaning to a
	// per-API method name (the package already scopes them).
	nameDropSegments = map[string]bool{"crm": true, "objects": true, "associations": true, "crm-object-schemas": true}
)

// deriveName builds a Go method name from an HTTP verb and a path. When
// includeParams is false, path parameters ({id}) are skipped; when true they are
// included as "By<Param>" segments to disambiguate otherwise-identical names.
func deriveName(method, path string, includeParams bool) string {
	var parts []string
	for _, seg := range strings.Split(path, "/") {
		if seg == "" || versionSegmentR.MatchString(seg) {
			continue
		}
		if strings.HasPrefix(seg, "{") {
			if includeParams {
				parts = append(parts, "By"+camelSegment(strings.Trim(seg, "{}")))
			}
			continue
		}
		if nameDropSegments[seg] {
			continue
		}
		parts = append(parts, camelSegment(seg))
	}
	verb := nameVerb[strings.ToLower(method)]
	if len(parts) == 0 {
		return verb + "Root"
	}
	return verb + strings.Join(parts, "")
}

// camelSegment converts a single path segment (which may contain '-' or '_')
// into an exported CamelCase identifier fragment.
func camelSegment(seg string) string {
	var b strings.Builder
	for _, part := range regexp.MustCompile(`[-_]`).Split(seg, -1) {
		if part == "" {
			continue
		}
		b.WriteString(strings.ToUpper(part[:1]))
		b.WriteString(part[1:])
	}
	return b.String()
}

// versionFromSchema retrieves the OpenAPI schema version from the input JSON string.
// It returns the API version.
func versionFromSchema(input string) (string, error) {
	var schema schemaOpenAPI
	err := json.Unmarshal([]byte(input), &schema)

	if err != nil {
		return "", err
	}

	return schema.Info.Version, nil
}

func main() {
	res, err := http.Get(directoryUrl)
	if err != nil {
		panic(err)
	}
	var r directory
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		panic(err)
	}
	err = os.Setenv("GIT_USER_ID", "clarkmcc")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("GIT_REPO_ID", "go-hubspot")
	if err != nil {
		panic(err)
	}
	generator, err := exec.LookPath("openapi-generator")
	if err != nil {
		panic(err)
	}

	//i := 0
	//outer:
	generated := map[string]bool{}
	for _, api := range r.Results {
		pkg, wanted := targetAPIs[api.Name]
		if !wanted || generated[api.Name] {
			continue
		}
		openAPIURL := api.openAPIFor(targetVersion)
		if openAPIURL == "" {
			fmt.Printf("WARNING: API %q has no %s release in catalog; skipping\n", api.Name, targetVersion)
			continue
		}

		schema, err := retrieveSchema(openAPIURL)
		if err != nil {
			panic(err)
		}

		// preProcessSchema keys renames by the catalog API name.
		filename, err := preProcessSchema(api.Name, schema)
		if err != nil {
			panic(err)
		}

		version, err := versionFromSchema(schema)
		if err != nil {
			panic(err)
		}

		fmt.Printf("generating %s/%s -> generated/%s/%s\n", strings.ToLower(api.Group), api.Name, version, pkg)
		out, err := exec.Command(generator, "generate",
			"-i", filename,
			"-g", "go",
			"--package-name", pkg,
			"--additional-properties=isGoSubmodule=false",
			"--skip-validate-spec",
			"-o", "./generated/"+version+"/"+pkg,
		).CombinedOutput()
		if err != nil {
			fmt.Println(string(out))
			panic(err)
		}
		generated[api.Name] = true
	}

	// Warn about any targeted API we never generated (e.g. absent from the
	// catalog response for this version), so coverage gaps are visible.
	for name := range targetAPIs {
		if !generated[name] {
			fmt.Printf("WARNING: targeted API %q was not generated (not found with a %s release)\n", name, targetVersion)
		}
	}

	fmt.Println("updating generated files")
	// Only post-process the version we just generated, so we never touch other
	// (possibly older) client trees under generated/.
	err = filepath.Walk(filepath.Join("generated", targetVersion), func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		// API files check for the presence of an API key in the request context using type
		// assertion. Because we generate each module separately, each module has its own
		// type for storing the API key in the context (contacts.APIKey vs deals.APIKey). In
		// order to allow us to share the same authorizer across packages, we'll do a find
		// and replace on all the package-specific API key structs and replace them with our
		// own global structs.
		if strings.HasSuffix(info.Name(), ".go") {
			// Open the file, read it, then close it
			var b []byte
			err = func() error {
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()
				b, err = ioutil.ReadAll(file)
				if err != nil {
					return err
				}
				return nil
			}()
			if err != nil {
				return err
			}

			// openapi-generator 7.x centralizes authentication in client.go's
			// prepareRequest ("Walk through any authentication.") rather than
			// emitting a per-operation API-key block. We inject our shared
			// hubspot.Authorizer there so a private-app token (or any Authorizer)
			// set via hubspot.WithAuthorizer is applied to every request.
			if info.Name() == "client.go" && bytes.Contains(b, authWalkAnchor) &&
				!bytes.Contains(b, authInjectMarker) {
				b = bytes.Replace(b, authWalkAnchor, append(append([]byte{}, authWalkAnchor...), authInject...), 1)

				// Add the import for the authorization package (once).
				if !bytes.Contains(b, []byte("\"github.com/clarkmcc/go-hubspot\"")) {
					idx := bytes.Index(b, []byte("\"net/url\""))
					if idx >= 0 {
						b = bytes.Join([][]byte{b[:idx], []byte("\t\"github.com/clarkmcc/go-hubspot\""), b[idx:]}, []byte("\n"))
					}
				}

				err = os.WriteFile(path, b, 0666)
				if err != nil {
					return err
				}
			}
		}
		// The client configuration gets generated with a go module for each generated
		// client. We'll go through and delete each submodule so that we can take advantage
		// of the root parent module.
		if info.Name() != "go.mod" && info.Name() != "go.sum" && info.Name() != "git_push.sh" {
			return nil
		}
		return os.Remove(path)
	})
	if err != nil {
		panic(err)
	}

	// openapi-generator 7.x emits a test/ directory of stub tests that import a
	// placeholder module (github.com/GIT_USER_ID/GIT_REPO_ID) and break `go vet`.
	// Remove them so the generated packages are self-contained under our module.
	for _, pkg := range targetAPIs {
		testDir := filepath.Join("generated", targetVersion, pkg, "test")
		if err := os.RemoveAll(testDir); err != nil {
			panic(err)
		}
	}

	// Format everything we generated.
	fmt.Println("formatting generated files")
	if out, err := exec.Command("gofmt", "-w", filepath.Join("generated", targetVersion)).CombinedOutput(); err != nil {
		fmt.Println(string(out))
		panic(err)
	}
}

// authWalkAnchor marks the point in generated client.go (prepareRequest) where
// openapi-generator 7.x applies authentication. We append our injection right
// after it.
var authWalkAnchor = []byte("// Walk through any authentication.")

// authInjectMarker is a unique substring of authInject used to detect whether a
// client.go has already been patched, so the generator is idempotent.
var authInjectMarker = []byte("hubspot.ApplyAuthorizerFromContext(ctx, localVarRequest)")

// authInject applies any hubspot.Authorizer stored on the context (e.g. a
// private-app TokenAuthorizer) to the outgoing *http.Request. The Authorizer
// writes into a header/query/form bag which we then copy onto the request. This
// replaces the per-operation API-key block emitted by older generator versions.
var authInject = []byte(`

		// HubSpot shared authorizer (injected by cmd/generate). Applies a
		// private-app token or any other Authorizer set via hubspot.WithAuthorizer.
		// Implemented as a helper in the hubspot package so this generated code does
		// not reference net/url, which prepareRequest shadows with a local variable.
		hubspot.ApplyAuthorizerFromContext(ctx, localVarRequest)`)
