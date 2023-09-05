# \ValidationApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Validate**](ValidationApi.md#Validate) | **Post** /cms/v3/source-code/{environment}/validate/{path} | Validate the contents of a file



## Validate

> Error Validate(ctx, path).File(file).Execute()

Validate the contents of a file



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
    path := "path_example" // string | The file system location of the file.
    file := os.NewFile(1234, "some_file") // *os.File | The file to validate. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ValidationApi.Validate(context.Background(), path).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ValidationApi.Validate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Validate`: Error
    fmt.Fprintf(os.Stdout, "Response from `ValidationApi.Validate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The file system location of the file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The file to validate. | 

### Return type

[**Error**](Error.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

