# \ContentApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentArchive**](ContentApi.md#ContentArchive) | **Delete** /cms/v3/source-code/{environment}/content/{path} | Delete a file
[**ContentCreate**](ContentApi.md#ContentCreate) | **Post** /cms/v3/source-code/{environment}/content/{path} | Create a file
[**ContentGet**](ContentApi.md#ContentGet) | **Get** /cms/v3/source-code/{environment}/content/{path} | Download a file
[**ContentReplace**](ContentApi.md#ContentReplace) | **Put** /cms/v3/source-code/{environment}/content/{path} | Create or update a file



## ContentArchive

> ContentArchive(ctx, environment, path).Execute()

Delete a file



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
    environment := "environment_example" // string | The environment of the file (\"draft\" or \"published\").
    path := "path_example" // string | The file system location of the file.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentApi.ContentArchive(context.Background(), environment, path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.ContentArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environment** | **string** | The environment of the file (\&quot;draft\&quot; or \&quot;published\&quot;). | 
**path** | **string** | The file system location of the file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentCreate

> AssetFileMetadata ContentCreate(ctx, environment, path).File(file).Execute()

Create a file



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
    environment := "environment_example" // string | The environment of the file (\"draft\" or \"published\").
    path := "path_example" // string | The file system location of the file.
    file := os.NewFile(1234, "some_file") // *os.File | The file to upload. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentApi.ContentCreate(context.Background(), environment, path).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.ContentCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentCreate`: AssetFileMetadata
    fmt.Fprintf(os.Stdout, "Response from `ContentApi.ContentCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environment** | **string** | The environment of the file (\&quot;draft\&quot; or \&quot;published\&quot;). | 
**path** | **string** | The file system location of the file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | The file to upload. | 

### Return type

[**AssetFileMetadata**](AssetFileMetadata.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentGet

> Error ContentGet(ctx, environment, path).Execute()

Download a file



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
    environment := "environment_example" // string | The environment of the file (\"draft\" or \"published\").
    path := "path_example" // string | The file system location of the file.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentApi.ContentGet(context.Background(), environment, path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.ContentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentGet`: Error
    fmt.Fprintf(os.Stdout, "Response from `ContentApi.ContentGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environment** | **string** | The environment of the file (\&quot;draft\&quot; or \&quot;published\&quot;). | 
**path** | **string** | The file system location of the file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Error**](Error.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentReplace

> AssetFileMetadata ContentReplace(ctx, environment, path).File(file).Execute()

Create or update a file



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
    environment := "environment_example" // string | The environment of the file (\"draft\" or \"published\").
    path := "path_example" // string | The file system location of the file.
    file := os.NewFile(1234, "some_file") // *os.File | The file to upload. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentApi.ContentReplace(context.Background(), environment, path).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.ContentReplace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentReplace`: AssetFileMetadata
    fmt.Fprintf(os.Stdout, "Response from `ContentApi.ContentReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environment** | **string** | The environment of the file (\&quot;draft\&quot; or \&quot;published\&quot;). | 
**path** | **string** | The file system location of the file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | The file to upload. | 

### Return type

[**AssetFileMetadata**](AssetFileMetadata.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

