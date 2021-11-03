# \ExtractApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCmsV3SourceCodeExtractPath**](ExtractApi.md#PostCmsV3SourceCodeExtractPath) | **Post** /cms/v3/source-code/extract/{path} | Extracts a zip file



## PostCmsV3SourceCodeExtractPath

> PostCmsV3SourceCodeExtractPath(ctx, path).Execute()

Extracts a zip file



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | The file system location of the zip file.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExtractApi.PostCmsV3SourceCodeExtractPath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtractApi.PostCmsV3SourceCodeExtractPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The file system location of the zip file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3SourceCodeExtractPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

