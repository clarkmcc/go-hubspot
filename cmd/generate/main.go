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
	directoryUrl = "https://api.hubspot.com/api-catalog-public/v1/apis"
	replacer     = strings.NewReplacer(" ", "_", "-", "_")
)

//preProcessMap contains a grouped list of operations to rename before generating typings.
var preProcessMap = map[string][]preProcessEntry{
	"Companies": {
		{old: "get-/crm/.*/objects/companies_getPage", new: "GetPage"},
		{old: "post-/crm/.*/objects/companies_create", new: "Create"},
		{old: "post-/crm/.*/objects/companies/merge_merge", new: "Merge"},
		{old: "post-/crm/.*/objects/companies/search_doSearch", new: "Search"},
		{old: "get-/crm/.*/objects/companies/{companyId}_getById", new: "Get"},
		{old: "delete-/crm/.*/objects/companies/{companyId}_archive", new: "Archive"},
		{old: "patch-/crm/.*/objects/companies/{companyId}_update", new: "Update"},
		{old: "post-/crm/.*/objects/companies/batch/archive_archive", new: "BatchArchive"},
		{old: "post-/crm/.*/objects/companies/batch/create_create", new: "BatchCreate"},
		{old: "post-/crm/.*/objects/companies/batch/read_read", new: "BatchRead"},
		{old: "post-/crm/.*/objects/companies/batch/update_update", new: "BatchUpdate"},
		{old: "get-/crm/.*/objects/companies/{companyId}/associations/{toObjectType}_getAll", new: "AssociationsGet"},
		{old: "put-/crm/.*/objects/companies/{companyId}/associations/{toObjectType}/{toObjectId}/{associationType}_create", new: "AssociationsCreate"},
		{old: "delete-/crm/.*/objects/companies/{companyId}/associations/{toObjectType}/{toObjectId}/{associationType}_archive", new: "AssociationsArchive"},
	},
	"Contacts": {
		{old: "get-/crm/.*/objects/contacts_getPage", new: "GetPage"},
		{old: "post-/crm/.*/objects/contacts_create", new: "Create"},
		{old: "post-/crm/.*/objects/contacts/gdpr-delete_purge", new: "Delete"},
		{old: "post-/crm/.*/objects/contacts/merge_merge", new: "Merge"},
		{old: "post-/crm/.*/objects/contacts/search_doSearch", new: "Search"},
		{old: "get-/crm/.*/objects/contacts/{contactId}_getById", new: "Get"},
		{old: "delete-/crm/.*/objects/contacts/{contactId}_archive", new: "Archive"},
		{old: "patch-/crm/.*/objects/contacts/{contactId}_update", new: "Update"},
		{old: "post-/crm/.*/objects/contacts/batch/archive_archive", new: "BatchArchive"},
		{old: "post-/crm/.*/objects/contacts/batch/create_create", new: "BatchCreate"},
		{old: "post-/crm/.*/objects/contacts/batch/read_read", new: "BatchRead"},
		{old: "post-/crm/.*/objects/contacts/batch/update_update", new: "BatchUpdate"},
		{old: "get-/crm/.*/objects/contacts/{contactId}/associations/{toObjectType}_getAll", new: "AssociationsGet"},
		{old: "put-/crm/.*/objects/contacts/{contactId}/associations/{toObjectType}/{toObjectId}/{associationType}_create", new: "AssociationsCreate"},
		{old: "delete-/crm/.*/objects/contacts/{contactId}/associations/{toObjectType}/{toObjectId}/{associationType}_archive", new: "AssociationsArchive"},
	},
	"Events": {
		{old: "get-/events/.*/events_getPage", new: "GetPage"},
	},
	"Objects": {
		{old: "get-/crm/.*/objects/{objectType}_getPage", new: "GetPage"},
		{old: "post-/crm/.*/objects/{objectType}_create", new: "Create"},
		{old: "post-/crm/.*/objects/{objectType}/batch/archive_archive", new: "BatchArchive"},
		{old: "post-/crm/.*/objects/{objectType}/batch/create_create", new: "BatchCreate"},
		{old: "post-/crm/.*/objects/{objectType}/batch/read_read", new: "BatchRead"},
		{old: "post-/crm/.*/objects/{objectType}/batch/update_update", new: "BatchUpdate"},
		{old: "post-/crm/.*/objects/{objectType}/gdpr-delete_purge", new: "Delete"},
		{old: "post-/crm/.*/objects/{objectType}/merge_merge", new: "Merge"},
		{old: "post-/crm/.*/objects/{objectType}/search_doSearch", new: "Search"},
		{old: "get-/crm/.*/objects/{objectType}/{objectId}_getById", new: "Get"},
		{old: "delete-/crm/.*/objects/{objectType}/{objectId}_archive", new: "Archive"},
		{old: "patch-/crm/.*/objects/{objectType}/{objectId}_update", new: "Update"},
		{old: "get-/crm/.*/objects/{objectType}/{objectId}/associations/{toObjectType}_getAll", new: "AssociationsGet"},
		{old: "put-/crm/.*/objects/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId}/{associationType}_create", new: "AssociationsCreate"},
		{old: "delete-/crm/.*/objects/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId}/{associationType}_archive", new: "AssociationsArchive"},
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

// directory is the schema for the https://api.hubspot.com/api-catalog-public/v1/apis page which lists all
// the publicly available APIs.
type directory struct {
	Results []struct {
		Name     string `json:"name"`
		Features map[string]struct {
			OpenAPI string `json:"openAPI"`
			Stage   string `json:"stage"`
		} `json:"features"`
	} `json:"results"`
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
	for _, group := range r.Results {
		for name, feature := range group.Features {
			schema, err := retrieveSchema(feature.OpenAPI)

			if err != nil {
				panic(err)
			}

			filename, err := preProcessSchema(name, schema)

			if err != nil {
				panic(err)
			}

			version, err := versionFromSchema(schema)

			if err != nil {
				panic(err)
			}

			//if i > 0 {
			//	break outer
			//}
			name = replacer.Replace(strings.ToLower(name))
			fmt.Printf("generating group/feature %s/%s\n", strings.ToLower(group.Name), name)
			_, err = exec.Command(generator, "generate",
				"-i", filename,
				"-g", "go",
				"--package-name", name,
				"--additional-properties=isGoSubmodule=false",
				"--skip-validate-spec",
				"-o", "./generated/"+version+"/"+name,
			).CombinedOutput()
			if err != nil {
				panic(err)
			}
			//i++
		}
	}

	fmt.Println("updating generated files")
	err = filepath.Walk("generated", func(path string, info fs.FileInfo, err error) error {
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
			find := []byte("if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {\n\t\t\tif apiKey, ok := auth[\"hapikey\"]; ok {\n\t\t\t\tvar key string\n\t\t\t\tif apiKey.Prefix != \"\" {\n\t\t\t\t\tkey = apiKey.Prefix + \" \" + apiKey.Key\n\t\t\t\t} else {\n\t\t\t\t\tkey = apiKey.Key\n\t\t\t\t}\n\t\t\t\tlocalVarQueryParams.Add(\"hapikey\", key)\n\t\t\t}\n\t\t}")
			replace := []byte("if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {\n\t\t\tauth.Apply(hubspot.AuthorizationRequest{\n\t\t\t\tQueryParams: localVarQueryParams,\n\t\t\t\tFormParams:  localVarFormParams,\n\t\t\t\tHeaders:     localVarHeaderParams,\n\t\t\t})\n\t\t}")
			if bytes.Contains(b, find) {
				// Replace generated API key auth with custom API key auth
				b = bytes.ReplaceAll(b, find, replace)

				// Add imports for authorization package
				idx := bytes.Index(b, []byte("\"net/url\""))
				if idx >= 0 {
					b = bytes.Join([][]byte{b[:idx], []byte("\t\"github.com/clarkmcc/go-hubspot\""), b[idx:]}, []byte("\n"))
				}

				err = os.Remove(path)
				file, err := os.Create(path)
				if err != nil {
					return err
				}
				defer file.Close()
				_, err = file.Write(b)
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
}
