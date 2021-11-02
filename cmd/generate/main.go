package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	swagger, err := exec.LookPath("swagger-codegen")
	if err != nil {
		panic(err)
	}
	for _, group := range r.Results {
		for name, feature := range group.Features {
			name = replacer.Replace(strings.ToLower(name))
			fmt.Printf("generating group/feature %s/%s\n", strings.ToLower(group.Name), name)
			_, err := exec.Command(swagger, "generate", "-i", feature.OpenAPI, "--additional-properties", "packageName="+name, "-l", "go", "-o", "./generated/"+name).CombinedOutput()
			if err != nil {
				panic(err)
			}
		}
	}

	// The client configuration structs have a base path with a trailing slash like
	// https://api.hubapi.com/ this. When concatenated with paths by the swagger-codegen,
	// it turns into https://api.hubapi.com//my/endpoint which always returns a 404. We'll
	// replace the base path with a base path that does not have a trailing slash.
	fmt.Println("fixing base path in generated clients")
	err = filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if info.Name() != "configuration.go" {
			return nil
		}
		var b []byte
		{
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			b, err = ioutil.ReadAll(file)
			if err != nil {
				return err
			}
		}
		b = bytes.Replace(b, []byte("https://api.hubapi.com/"), []byte("https://api.hubapi.com"), 1)
		err = os.Remove(path)
		if err != nil {
			return err
		}
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		_, err = file.Write(b)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
