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
	"strings"
)

var (
	// directoryUrl is HubSpot's public API spec catalog. The previous
	// "api-catalog-public/v1/apis" endpoint was retired and now returns 404; the
	// current catalog lives here and uses the date-based versioning scheme.
	directoryUrl = "https://api.hubspot.com/public/api/spec/v2/specs"
	replacer     = strings.NewReplacer(" ", "_", "-", "_")

	// targetVersion selects which date-based API release to generate. HubSpot
	// moved from semantic (v1..v4) to date-based (YYYY-MM) versions; legacy
	// versions are unsupported from Sep 2027. 2026-09 is the current target.
	targetVersion = "2026-09"

	// targetAPIs is the allowlist of catalog API names to generate, mapped to the
	// Go package name they are emitted under. We intentionally generate only the
	// APIs this project consumes rather than all ~170 catalog entries: operation
	// renames (preProcessMap) are only maintained for these, and generating the
	// full catalog is slow and noisy.
	targetAPIs = map[string]string{
		"Contacts":       "contacts",
		"Custom Objects": "objects",
		"Associations":   "associations",
	}
)

// preProcessMap contains a grouped list of operations to rename before generating typings.
// preProcessMap contains, per catalog API name, a list of operationId rewrites
// applied to the raw OpenAPI spec JSON before generation. HubSpot's date-based
// specs expose long, auto-generated operationIds (e.g.
// "post-/crm/objects/2026-09/contacts/search_/crm/objects/2026-09-beta/{objectType}/search")
// which openapi-generator would turn into unusable Go method names. Each entry
// matches the whole operationId string and replaces it with a clean method name.
// The version segment is matched loosely ([0-9-]+(?:-beta)?) so the same rules
// keep working across date-based releases.
var preProcessMap = map[string][]preProcessEntry{
	"Contacts": {
		{old: `post-/crm/objects/[0-9-]+(?:-beta)?/contacts/search_[^"\\]*`, new: "Search"},
		{old: `post-/crm/objects/[0-9-]+(?:-beta)?/contacts/batch/create_[^"\\]*`, new: "BatchCreate"},
		{old: `post-/crm/objects/[0-9-]+(?:-beta)?/contacts/batch/update_[^"\\]*`, new: "BatchUpdate"},
		{old: `post-/crm/objects/[0-9-]+(?:-beta)?/contacts/batch/archive_[^"\\]*`, new: "BatchArchive"},
		{old: `post-/crm/objects/[0-9-]+(?:-beta)?/contacts/batch/read_[^"\\]*`, new: "BatchRead"},
	},
	"Custom Objects": {
		{old: `post-/crm/objects/[0-9-]+(?:-beta)?/\{objectType\}/search_[^"\\]*`, new: "Search"},
		{old: `post-/crm/objects/[0-9-]+(?:-beta)?/\{objectType\}/batch/create_[^"\\]*`, new: "BatchCreate"},
		{old: `post-/crm/objects/[0-9-]+(?:-beta)?/\{objectType\}/batch/update_[^"\\]*`, new: "BatchUpdate"},
		{old: `post-/crm/objects/[0-9-]+(?:-beta)?/\{objectType\}/batch/archive_[^"\\]*`, new: "BatchArchive"},
		{old: `post-/crm/objects/[0-9-]+(?:-beta)?/\{objectType\}/batch/read_[^"\\]*`, new: "BatchRead"},
	},
	"Associations": {
		// The legacy "/types" read is gone in date-based versions; association
		// type ids are now read from "/labels" (GET, tag Basic -> BasicAPIService).
		{old: `get-/crm/associations/[0-9-]+(?:-beta)?/\{fromObjectType\}/\{toObjectType\}/labels_[^"\\]*`, new: "GetLabels"},
		{old: `post-/crm/associations/[0-9-]+(?:-beta)?/\{fromObjectType\}/\{toObjectType\}/batch/create_[^"\\]*`, new: "BatchCreate"},
		{old: `post-/crm/associations/[0-9-]+(?:-beta)?/\{fromObjectType\}/\{toObjectType\}/batch/archive_[^"\\]*`, new: "BatchArchive"},
	},
}

// preProcessEntry contains a mapping from an old operation to a new operation.
type preProcessEntry struct {
	old string
	new string
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

// preProcessSchema replaces operation names in the given schema and saves the results to file.
// It returns a path to the saved file.
func preProcessSchema(group string, schema string) (string, error) {
	entries := preProcessMap[group]

	for _, entry := range entries {
		pattern := regexp.MustCompile(entry.old)
		schema = pattern.ReplaceAllString(schema, entry.new)
	}

	filename := "./schema/" + group + ".json"
	err := ioutil.WriteFile(filename, []byte(schema), 0644)

	if err != nil {
		return "", err
	}

	return filename, nil
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
	for _, api := range r.Results {
		pkg, wanted := targetAPIs[api.Name]
		if !wanted {
			continue
		}
		openAPIURL := api.openAPIFor(targetVersion)
		if openAPIURL == "" {
			panic(fmt.Sprintf("API %q has no %s release in catalog", api.Name, targetVersion))
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
