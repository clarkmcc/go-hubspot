# \RedirectsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCmsV3UrlRedirectsUrlRedirectIdArchive**](RedirectsApi.md#DeleteCmsV3UrlRedirectsUrlRedirectIdArchive) | **Delete** /cms/v3/url-redirects/{urlRedirectId} | Delete a redirect
[**GetCmsV3UrlRedirectsGetPage**](RedirectsApi.md#GetCmsV3UrlRedirectsGetPage) | **Get** /cms/v3/url-redirects/ | Get current redirects
[**GetCmsV3UrlRedirectsUrlRedirectIdGetById**](RedirectsApi.md#GetCmsV3UrlRedirectsUrlRedirectIdGetById) | **Get** /cms/v3/url-redirects/{urlRedirectId} | Get details for a redirect
[**PatchCmsV3UrlRedirectsUrlRedirectIdUpdate**](RedirectsApi.md#PatchCmsV3UrlRedirectsUrlRedirectIdUpdate) | **Patch** /cms/v3/url-redirects/{urlRedirectId} | Update a redirect
[**PostCmsV3UrlRedirectsCreate**](RedirectsApi.md#PostCmsV3UrlRedirectsCreate) | **Post** /cms/v3/url-redirects/ | Create a redirect



## DeleteCmsV3UrlRedirectsUrlRedirectIdArchive

> DeleteCmsV3UrlRedirectsUrlRedirectIdArchive(ctx, urlRedirectId).Execute()

Delete a redirect



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
    urlRedirectId := "urlRedirectId_example" // string | The ID of the target redirect.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RedirectsApi.DeleteCmsV3UrlRedirectsUrlRedirectIdArchive(context.Background(), urlRedirectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RedirectsApi.DeleteCmsV3UrlRedirectsUrlRedirectIdArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlRedirectId** | **string** | The ID of the target redirect. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCmsV3UrlRedirectsUrlRedirectIdArchiveRequest struct via the builder pattern


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


## GetCmsV3UrlRedirectsGetPage

> CollectionResponseWithTotalUrlMapping GetCmsV3UrlRedirectsGetPage(ctx).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).Properties(properties).After(after).Before(before).Limit(limit).Archived(archived).Execute()

Get current redirects



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    createdAt := time.Now() // time.Time | Only return redirects created on exactly this date. (optional)
    createdAfter := time.Now() // time.Time | Only return redirects created after this date. (optional)
    createdBefore := time.Now() // time.Time | Only return redirects created before this date. (optional)
    updatedAt := time.Now() // time.Time | Only return redirects last updated on exactly this date. (optional)
    updatedAfter := time.Now() // time.Time | Only return redirects last updated after this date. (optional)
    updatedBefore := time.Now() // time.Time | Only return redirects last updated before this date. (optional)
    sort := []string{"Inner_example"} // []string |  (optional)
    properties := []string{"Inner_example"} // []string |  (optional)
    after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
    before := "before_example" // string |  (optional)
    limit := int32(56) // int32 | Maximum number of result per page (optional)
    archived := true // bool | Whether to return only results that have been archived. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RedirectsApi.GetCmsV3UrlRedirectsGetPage(context.Background()).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).Properties(properties).After(after).Before(before).Limit(limit).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RedirectsApi.GetCmsV3UrlRedirectsGetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3UrlRedirectsGetPage`: CollectionResponseWithTotalUrlMapping
    fmt.Fprintf(os.Stdout, "Response from `RedirectsApi.GetCmsV3UrlRedirectsGetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3UrlRedirectsGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAt** | **time.Time** | Only return redirects created on exactly this date. | 
 **createdAfter** | **time.Time** | Only return redirects created after this date. | 
 **createdBefore** | **time.Time** | Only return redirects created before this date. | 
 **updatedAt** | **time.Time** | Only return redirects last updated on exactly this date. | 
 **updatedAfter** | **time.Time** | Only return redirects last updated after this date. | 
 **updatedBefore** | **time.Time** | Only return redirects last updated before this date. | 
 **sort** | **[]string** |  | 
 **properties** | **[]string** |  | 
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **before** | **string** |  | 
 **limit** | **int32** | Maximum number of result per page | 
 **archived** | **bool** | Whether to return only results that have been archived. | 

### Return type

[**CollectionResponseWithTotalUrlMapping**](CollectionResponseWithTotalUrlMapping.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3UrlRedirectsUrlRedirectIdGetById

> UrlMapping GetCmsV3UrlRedirectsUrlRedirectIdGetById(ctx, urlRedirectId).Execute()

Get details for a redirect



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
    urlRedirectId := "urlRedirectId_example" // string | The ID of the target redirect.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RedirectsApi.GetCmsV3UrlRedirectsUrlRedirectIdGetById(context.Background(), urlRedirectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RedirectsApi.GetCmsV3UrlRedirectsUrlRedirectIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3UrlRedirectsUrlRedirectIdGetById`: UrlMapping
    fmt.Fprintf(os.Stdout, "Response from `RedirectsApi.GetCmsV3UrlRedirectsUrlRedirectIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlRedirectId** | **string** | The ID of the target redirect. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3UrlRedirectsUrlRedirectIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UrlMapping**](UrlMapping.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCmsV3UrlRedirectsUrlRedirectIdUpdate

> UrlMapping PatchCmsV3UrlRedirectsUrlRedirectIdUpdate(ctx, urlRedirectId).UrlMapping(urlMapping).Execute()

Update a redirect



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
    urlRedirectId := "urlRedirectId_example" // string | 
    urlMapping := *openapiclient.NewUrlMapping(int64(123), int32(123), int64(123), int64(123), int32(123), int32(123), "RoutePrefix_example", "Destination_example", int32(123), int64(123), false, false, false, false, false, false, false, "Name_example", int32(123), int64(123), "Note_example", "Label_example", false, "CosObjectType_example", int64(123)) // UrlMapping |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RedirectsApi.PatchCmsV3UrlRedirectsUrlRedirectIdUpdate(context.Background(), urlRedirectId).UrlMapping(urlMapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RedirectsApi.PatchCmsV3UrlRedirectsUrlRedirectIdUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCmsV3UrlRedirectsUrlRedirectIdUpdate`: UrlMapping
    fmt.Fprintf(os.Stdout, "Response from `RedirectsApi.PatchCmsV3UrlRedirectsUrlRedirectIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlRedirectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCmsV3UrlRedirectsUrlRedirectIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **urlMapping** | [**UrlMapping**](UrlMapping.md) |  | 

### Return type

[**UrlMapping**](UrlMapping.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3UrlRedirectsCreate

> UrlMapping PostCmsV3UrlRedirectsCreate(ctx).UrlMappingCreateRequestBody(urlMappingCreateRequestBody).Execute()

Create a redirect



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
    urlMappingCreateRequestBody := *openapiclient.NewUrlMappingCreateRequestBody("RoutePrefix_example", "Destination_example", int32(123)) // UrlMappingCreateRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RedirectsApi.PostCmsV3UrlRedirectsCreate(context.Background()).UrlMappingCreateRequestBody(urlMappingCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RedirectsApi.PostCmsV3UrlRedirectsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3UrlRedirectsCreate`: UrlMapping
    fmt.Fprintf(os.Stdout, "Response from `RedirectsApi.PostCmsV3UrlRedirectsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3UrlRedirectsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **urlMappingCreateRequestBody** | [**UrlMappingCreateRequestBody**](UrlMappingCreateRequestBody.md) |  | 

### Return type

[**UrlMapping**](UrlMapping.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

