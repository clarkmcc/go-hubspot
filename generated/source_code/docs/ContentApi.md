# \ContentApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCmsV3SourceCodeEnvironmentContentPathArchive**](ContentApi.md#DeleteCmsV3SourceCodeEnvironmentContentPathArchive) | **Delete** /cms/v3/source-code/{environment}/content/{path} | Delete a file
[**GetCmsV3SourceCodeEnvironmentContentPathGet**](ContentApi.md#GetCmsV3SourceCodeEnvironmentContentPathGet) | **Get** /cms/v3/source-code/{environment}/content/{path} | Download a file
[**PostCmsV3SourceCodeEnvironmentContentPathCreate**](ContentApi.md#PostCmsV3SourceCodeEnvironmentContentPathCreate) | **Post** /cms/v3/source-code/{environment}/content/{path} | Create a file
[**PutCmsV3SourceCodeEnvironmentContentPathReplace**](ContentApi.md#PutCmsV3SourceCodeEnvironmentContentPathReplace) | **Put** /cms/v3/source-code/{environment}/content/{path} | Create or update a file



## DeleteCmsV3SourceCodeEnvironmentContentPathArchive

> DeleteCmsV3SourceCodeEnvironmentContentPathArchive(ctx, environment, path).Execute()

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
    resp, r, err := apiClient.ContentApi.DeleteCmsV3SourceCodeEnvironmentContentPathArchive(context.Background(), environment, path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.DeleteCmsV3SourceCodeEnvironmentContentPathArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCmsV3SourceCodeEnvironmentContentPathArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3SourceCodeEnvironmentContentPathGet

> Error GetCmsV3SourceCodeEnvironmentContentPathGet(ctx, environment, path).Execute()

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
    resp, r, err := apiClient.ContentApi.GetCmsV3SourceCodeEnvironmentContentPathGet(context.Background(), environment, path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.GetCmsV3SourceCodeEnvironmentContentPathGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3SourceCodeEnvironmentContentPathGet`: Error
    fmt.Fprintf(os.Stdout, "Response from `ContentApi.GetCmsV3SourceCodeEnvironmentContentPathGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environment** | **string** | The environment of the file (\&quot;draft\&quot; or \&quot;published\&quot;). | 
**path** | **string** | The file system location of the file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3SourceCodeEnvironmentContentPathGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Error**](Error.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3SourceCodeEnvironmentContentPathCreate

> AssetFileMetadata PostCmsV3SourceCodeEnvironmentContentPathCreate(ctx, environment, path).File(file).Execute()

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
    resp, r, err := apiClient.ContentApi.PostCmsV3SourceCodeEnvironmentContentPathCreate(context.Background(), environment, path).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.PostCmsV3SourceCodeEnvironmentContentPathCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3SourceCodeEnvironmentContentPathCreate`: AssetFileMetadata
    fmt.Fprintf(os.Stdout, "Response from `ContentApi.PostCmsV3SourceCodeEnvironmentContentPathCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environment** | **string** | The environment of the file (\&quot;draft\&quot; or \&quot;published\&quot;). | 
**path** | **string** | The file system location of the file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3SourceCodeEnvironmentContentPathCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | The file to upload. | 

### Return type

[**AssetFileMetadata**](AssetFileMetadata.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCmsV3SourceCodeEnvironmentContentPathReplace

> AssetFileMetadata PutCmsV3SourceCodeEnvironmentContentPathReplace(ctx, environment, path).File(file).Execute()

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
    resp, r, err := apiClient.ContentApi.PutCmsV3SourceCodeEnvironmentContentPathReplace(context.Background(), environment, path).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.PutCmsV3SourceCodeEnvironmentContentPathReplace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutCmsV3SourceCodeEnvironmentContentPathReplace`: AssetFileMetadata
    fmt.Fprintf(os.Stdout, "Response from `ContentApi.PutCmsV3SourceCodeEnvironmentContentPathReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environment** | **string** | The environment of the file (\&quot;draft\&quot; or \&quot;published\&quot;). | 
**path** | **string** | The file system location of the file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCmsV3SourceCodeEnvironmentContentPathReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | The file to upload. | 

### Return type

[**AssetFileMetadata**](AssetFileMetadata.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

