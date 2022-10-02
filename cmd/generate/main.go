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
	"strings"
)

var (
	directoryUrl = "https://api.hubspot.com/api-catalog-public/v1/apis"
	replacer     = strings.NewReplacer(" ", "_", "-", "_")
)

//preProcessMap contains a grouped list of operations to rename before generating typings.
var preProcessMap = map[string][]preProcessEntry{
	"Companies": {
		{old: "get-/crm/v3/objects/companies_getPage", new: "GetPage"},
		{old: "post-/crm/v3/objects/companies_create", new: "Create"},
		{old: "post-/crm/v3/objects/companies/merge_merge", new: "Merge"},
		{old: "post-/crm/v3/objects/companies/search_doSearch", new: "Search"},
		{old: "get-/crm/v3/objects/companies/{companyId}_getById", new: "Get"},
		{old: "delete-/crm/v3/objects/companies/{companyId}_archive", new: "Archive"},
		{old: "patch-/crm/v3/objects/companies/{companyId}_update", new: "Update"},
		{old: "post-/crm/v3/objects/companies/batch/archive_archive", new: "BatchArchive"},
		{old: "post-/crm/v3/objects/companies/batch/create_create", new: "BatchCreate"},
		{old: "post-/crm/v3/objects/companies/batch/read_read", new: "BatchRead"},
		{old: "post-/crm/v3/objects/companies/batch/update_update", new: "BatchUpdate"},
		{old: "get-/crm/v3/objects/companies/{companyId}/associations/{toObjectType}_getAll", new: "AssociationsGet"},
		{old: "put-/crm/v3/objects/companies/{companyId}/associations/{toObjectType}/{toObjectId}/{associationType}_create", new: "AssociationsCreate"},
		{old: "delete-/crm/v3/objects/companies/{companyId}/associations/{toObjectType}/{toObjectId}/{associationType}_archive", new: "AssociationsArchive"},
	},
	"Contacts": {
		{old: "get-/crm/v3/objects/contacts_getPage", new: "GetPage"},
		{old: "post-/crm/v3/objects/contacts_create", new: "Create"},
		{old: "post-/crm/v3/objects/contacts/gdpr-delete_purge", new: "Delete"},
		{old: "post-/crm/v3/objects/contacts/merge_merge", new: "Merge"},
		{old: "post-/crm/v3/objects/contacts/search_doSearch", new: "Search"},
		{old: "get-/crm/v3/objects/contacts/{contactId}_getById", new: "Get"},
		{old: "delete-/crm/v3/objects/contacts/{contactId}_archive", new: "Archive"},
		{old: "patch-/crm/v3/objects/contacts/{contactId}_update", new: "Update"},
		{old: "post-/crm/v3/objects/contacts/batch/archive_archive", new: "BatchArchive"},
		{old: "post-/crm/v3/objects/contacts/batch/create_create", new: "BatchCreate"},
		{old: "post-/crm/v3/objects/contacts/batch/read_read", new: "BatchRead"},
		{old: "post-/crm/v3/objects/contacts/batch/update_update", new: "BatchUpdate"},
		{old: "get-/crm/v3/objects/contacts/{contactId}/associations/{toObjectType}_getAll", new: "AssociationsGet"},
		{old: "put-/crm/v3/objects/contacts/{contactId}/associations/{toObjectType}/{toObjectId}/{associationType}_create", new: "AssociationsCreate"},
		{old: "delete-/crm/v3/objects/contacts/{contactId}/associations/{toObjectType}/{toObjectId}/{associationType}_archive", new: "AssociationsArchive"},
	},
	"Events": {
		{old: "get-/events/v3/events_getPage", new: "GetPage"},
	},
	"Objects": {
		{old: "get-/crm/v3/objects/{objectType}_getPage", new: "GetPage"},
		{old: "post-/crm/v3/objects/{objectType}_create", new: "Create"},
		{old: "post-/crm/v3/objects/{objectType}/batch/archive_archive", new: "BatchArchive"},
		{old: "post-/crm/v3/objects/{objectType}/batch/create_create", new: "BatchCreate"},
		{old: "post-/crm/v3/objects/{objectType}/batch/read_read", new: "BatchRead"},
		{old: "post-/crm/v3/objects/{objectType}/batch/update_update", new: "BatchUpdate"},
		{old: "post-/crm/v3/objects/{objectType}/gdpr-delete_purge", new: "Delete"},
		{old: "post-/crm/v3/objects/{objectType}/merge_merge", new: "Merge"},
		{old: "post-/crm/v3/objects/{objectType}/search_doSearch", new: "Search"},
		{old: "get-/crm/v3/objects/{objectType}/{objectId}_getById", new: "Get"},
		{old: "delete-/crm/v3/objects/{objectType}/{objectId}_archive", new: "Archive"},
		{old: "patch-/crm/v3/objects/{objectType}/{objectId}_update", new: "Update"},
		{old: "get-/crm/v3/objects/{objectType}/{objectId}/associations/{toObjectType}_getAll", new: "AssociationsGet"},
		{old: "put-/crm/v3/objects/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId}/{associationType}_create", new: "AssociationsCreate"},
		{old: "delete-/crm/v3/objects/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId}/{associationType}_archive", new: "AssociationsArchive"},
	},
}

// preProcessEntry contains a mapping from an old operation to a new operation.
type preProcessEntry struct {
	old string
	new string
}

// Directory is the schema for the https://api.hubspot.com/api-catalog-public/v1/apis page which lists all
// the publicly available APIs.
type Directory struct {
	Results []struct {
		Name     string `json:"name"`
		Features map[string]struct {
			OpenAPI string `json:"openAPI"`
			Stage   string `json:"stage"`
		} `json:"features"`
	} `json:"results"`
}

func preProcessResponse(name string, url string) (string, error) {
	res, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	result := string(b)

	entries := preProcessMap[name]

	for _, entry := range entries {
		result = strings.ReplaceAll(result, entry.old, entry.new)
	}

	filename := "./schema/" + name + ".json"
	err = ioutil.WriteFile(filename, []byte(result), 0644)

	if err != nil {
		return "", err
	}

	return filename, nil
}

func main() {
	res, err := http.Get(directoryUrl)
	if err != nil {
		panic(err)
	}
	var r Directory
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
			filename, err := preProcessResponse(name, feature.OpenAPI)

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
				"-o", "./generated/"+name,
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
