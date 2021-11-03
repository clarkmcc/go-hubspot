package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
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
	for _, group := range r.Results {
		for name, feature := range group.Features {
			name = replacer.Replace(strings.ToLower(name))
			fmt.Printf("generating group/feature %s/%s\n", strings.ToLower(group.Name), name)
			_, err := exec.Command(generator, "generate",
				"-i", feature.OpenAPI,
				"-g", "go",
				"--package-name", name,
				"--additional-properties=isGoSubmodule=false",
				"--skip-validate-spec",
				"-o", "./generated/"+name,
			).CombinedOutput()
			if err != nil {
				panic(err)
			}
		}
	}

	// The client configuration gets generated with a go module for each generated
	// client. We'll go through and delete each submodule so that we can take advantage
	// of the root parent module.
	fmt.Println("fixing base path in generated clients")
	err = filepath.Walk("generated", func(path string, info fs.FileInfo, err error) error {
		if info.Name() != "go.mod" && info.Name() != "go.sum" {
			return nil
		}
		return os.Remove(path)
	})
	if err != nil {
		panic(err)
	}
}
