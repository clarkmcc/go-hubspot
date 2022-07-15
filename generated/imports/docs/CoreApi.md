# \CoreApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCrmV3ImportsGetPage**](CoreApi.md#GetCrmV3ImportsGetPage) | **Get** /crm/v3/imports/ | Get active imports
[**GetCrmV3ImportsImportIdGetById**](CoreApi.md#GetCrmV3ImportsImportIdGetById) | **Get** /crm/v3/imports/{importId} | Get the information on any import
[**PostCrmV3ImportsCreate**](CoreApi.md#PostCrmV3ImportsCreate) | **Post** /crm/v3/imports/ | Start a new import
[**PostCrmV3ImportsImportIdCancelCancel**](CoreApi.md#PostCrmV3ImportsImportIdCancelCancel) | **Post** /crm/v3/imports/{importId}/cancel | Cancel an active import



## GetCrmV3ImportsGetPage

> CollectionResponsePublicImportResponse GetCrmV3ImportsGetPage(ctx).After(after).Before(before).Limit(limit).Execute()

Get active imports



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
    after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
    before := "before_example" // string |  (optional)
    limit := int32(56) // int32 | The maximum number of results to display per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreApi.GetCrmV3ImportsGetPage(context.Background()).After(after).Before(before).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.GetCrmV3ImportsGetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3ImportsGetPage`: CollectionResponsePublicImportResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreApi.GetCrmV3ImportsGetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ImportsGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **before** | **string** |  | 
 **limit** | **int32** | The maximum number of results to display per page. | 

### Return type

[**CollectionResponsePublicImportResponse**](CollectionResponsePublicImportResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3ImportsImportIdGetById

> PublicImportResponse GetCrmV3ImportsImportIdGetById(ctx, importId).Execute()

Get the information on any import



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
    importId := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreApi.GetCrmV3ImportsImportIdGetById(context.Background(), importId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.GetCrmV3ImportsImportIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3ImportsImportIdGetById`: PublicImportResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreApi.GetCrmV3ImportsImportIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**importId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ImportsImportIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicImportResponse**](PublicImportResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ImportsCreate

> PublicImportResponse PostCrmV3ImportsCreate(ctx).Files(files).ImportRequest(importRequest).Execute()

Start a new import



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
    files := os.NewFile(1234, "some_file") // *os.File | A list of files containing the data to import (optional)
    importRequest := "importRequest_example" // string | JSON formatted metadata about the import. This includes a name for the import and the column mappings for each file. See the overview tab for more on the required format. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreApi.PostCrmV3ImportsCreate(context.Background()).Files(files).ImportRequest(importRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.PostCrmV3ImportsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ImportsCreate`: PublicImportResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreApi.PostCrmV3ImportsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ImportsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **files** | ***os.File** | A list of files containing the data to import | 
 **importRequest** | **string** | JSON formatted metadata about the import. This includes a name for the import and the column mappings for each file. See the overview tab for more on the required format. | 

### Return type

[**PublicImportResponse**](PublicImportResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ImportsImportIdCancelCancel

> ActionResponse PostCrmV3ImportsImportIdCancelCancel(ctx, importId).Execute()

Cancel an active import



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
    importId := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreApi.PostCrmV3ImportsImportIdCancelCancel(context.Background(), importId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.PostCrmV3ImportsImportIdCancelCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ImportsImportIdCancelCancel`: ActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreApi.PostCrmV3ImportsImportIdCancelCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**importId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ImportsImportIdCancelCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionResponse**](ActionResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

