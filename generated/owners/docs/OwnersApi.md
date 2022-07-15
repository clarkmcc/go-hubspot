# \OwnersApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCrmV3OwnersGetPage**](OwnersApi.md#GetCrmV3OwnersGetPage) | **Get** /crm/v3/owners/ | Get a page of owners
[**GetCrmV3OwnersOwnerIdGetById**](OwnersApi.md#GetCrmV3OwnersOwnerIdGetById) | **Get** /crm/v3/owners/{ownerId} | Read an owner by given &#x60;id&#x60; or &#x60;userId&#x60;



## GetCrmV3OwnersGetPage

> CollectionResponsePublicOwnerForwardPaging GetCrmV3OwnersGetPage(ctx).Email(email).After(after).Limit(limit).Archived(archived).Execute()

Get a page of owners

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
    email := "email_example" // string | Filter by email address (optional) (optional)
    after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
    limit := int32(56) // int32 | The maximum number of results to display per page. (optional) (default to 100)
    archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OwnersApi.GetCrmV3OwnersGetPage(context.Background()).Email(email).After(after).Limit(limit).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OwnersApi.GetCrmV3OwnersGetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3OwnersGetPage`: CollectionResponsePublicOwnerForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `OwnersApi.GetCrmV3OwnersGetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3OwnersGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Filter by email address (optional) | 
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to display per page. | [default to 100]
 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

[**CollectionResponsePublicOwnerForwardPaging**](CollectionResponsePublicOwnerForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3OwnersOwnerIdGetById

> PublicOwner GetCrmV3OwnersOwnerIdGetById(ctx, ownerId).IdProperty(idProperty).Archived(archived).Execute()

Read an owner by given `id` or `userId`

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
    ownerId := int32(56) // int32 | 
    idProperty := "idProperty_example" // string |  (optional) (default to "id")
    archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OwnersApi.GetCrmV3OwnersOwnerIdGetById(context.Background(), ownerId).IdProperty(idProperty).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OwnersApi.GetCrmV3OwnersOwnerIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3OwnersOwnerIdGetById`: PublicOwner
    fmt.Fprintf(os.Stdout, "Response from `OwnersApi.GetCrmV3OwnersOwnerIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3OwnersOwnerIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idProperty** | **string** |  | [default to &quot;id&quot;]
 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

[**PublicOwner**](PublicOwner.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

