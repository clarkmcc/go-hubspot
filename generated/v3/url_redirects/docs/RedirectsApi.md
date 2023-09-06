# \RedirectsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Archive**](RedirectsApi.md#Archive) | **Delete** /cms/v3/url-redirects/{urlRedirectId} | Delete a redirect
[**Create**](RedirectsApi.md#Create) | **Post** /cms/v3/url-redirects/ | Create a redirect
[**GetByID**](RedirectsApi.md#GetByID) | **Get** /cms/v3/url-redirects/{urlRedirectId} | Get details for a redirect
[**GetPage**](RedirectsApi.md#GetPage) | **Get** /cms/v3/url-redirects/ | Get current redirects
[**Update**](RedirectsApi.md#Update) | **Patch** /cms/v3/url-redirects/{urlRedirectId} | Update a redirect



## Archive

> Archive(ctx, urlRedirectId).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RedirectsApi.Archive(context.Background(), urlRedirectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RedirectsApi.Archive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiArchiveRequest struct via the builder pattern


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


## Create

> UrlMapping Create(ctx).UrlMappingCreateRequestBody(urlMappingCreateRequestBody).Execute()

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
    urlMappingCreateRequestBody := *openapiclient.NewUrlMappingCreateRequestBody("RoutePrefix_example", "Destination_example", int32(123)) // UrlMappingCreateRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RedirectsApi.Create(context.Background()).UrlMappingCreateRequestBody(urlMappingCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RedirectsApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: UrlMapping
    fmt.Fprintf(os.Stdout, "Response from `RedirectsApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **urlMappingCreateRequestBody** | [**UrlMappingCreateRequestBody**](UrlMappingCreateRequestBody.md) |  | 

### Return type

[**UrlMapping**](UrlMapping.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByID

> UrlMapping GetByID(ctx, urlRedirectId).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RedirectsApi.GetByID(context.Background(), urlRedirectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RedirectsApi.GetByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetByID`: UrlMapping
    fmt.Fprintf(os.Stdout, "Response from `RedirectsApi.GetByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlRedirectId** | **string** | The ID of the target redirect. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UrlMapping**](UrlMapping.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPage

> CollectionResponseWithTotalUrlMappingForwardPaging GetPage(ctx).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Execute()

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
    after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
    limit := int32(56) // int32 | Maximum number of result per page (optional)
    archived := true // bool | Whether to return only results that have been archived. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RedirectsApi.GetPage(context.Background()).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RedirectsApi.GetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPage`: CollectionResponseWithTotalUrlMappingForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `RedirectsApi.GetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAt** | **time.Time** | Only return redirects created on exactly this date. | 
 **createdAfter** | **time.Time** | Only return redirects created after this date. | 
 **createdBefore** | **time.Time** | Only return redirects created before this date. | 
 **updatedAt** | **time.Time** | Only return redirects last updated on exactly this date. | 
 **updatedAfter** | **time.Time** | Only return redirects last updated after this date. | 
 **updatedBefore** | **time.Time** | Only return redirects last updated before this date. | 
 **sort** | **[]string** |  | 
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | Maximum number of result per page | 
 **archived** | **bool** | Whether to return only results that have been archived. | 

### Return type

[**CollectionResponseWithTotalUrlMappingForwardPaging**](CollectionResponseWithTotalUrlMappingForwardPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UrlMapping Update(ctx, urlRedirectId).UrlMapping(urlMapping).Execute()

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
    urlMapping := *openapiclient.NewUrlMapping("Id_example", "RoutePrefix_example", "Destination_example", int32(123), false, false, false, false, false, false, int32(123)) // UrlMapping | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RedirectsApi.Update(context.Background(), urlRedirectId).UrlMapping(urlMapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RedirectsApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: UrlMapping
    fmt.Fprintf(os.Stdout, "Response from `RedirectsApi.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlRedirectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **urlMapping** | [**UrlMapping**](UrlMapping.md) |  | 

### Return type

[**UrlMapping**](UrlMapping.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

