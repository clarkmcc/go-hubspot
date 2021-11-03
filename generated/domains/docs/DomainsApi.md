# \DomainsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCmsV3DomainsDomainIdGetById**](DomainsApi.md#GetCmsV3DomainsDomainIdGetById) | **Get** /cms/v3/domains/{domainId} | Get a single domain
[**GetCmsV3DomainsGetPage**](DomainsApi.md#GetCmsV3DomainsGetPage) | **Get** /cms/v3/domains/ | Get current domains



## GetCmsV3DomainsDomainIdGetById

> Domain GetCmsV3DomainsDomainIdGetById(ctx, domainId).Archived(archived).Execute()

Get a single domain



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
    domainId := "domainId_example" // string | The unique ID of the domain.
    archived := true // bool | Whether to return only results that have been archived. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsApi.GetCmsV3DomainsDomainIdGetById(context.Background(), domainId).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.GetCmsV3DomainsDomainIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3DomainsDomainIdGetById`: Domain
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.GetCmsV3DomainsDomainIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | The unique ID of the domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3DomainsDomainIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Whether to return only results that have been archived. | 

### Return type

[**Domain**](Domain.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3DomainsGetPage

> CollectionResponseWithTotalDomain GetCmsV3DomainsGetPage(ctx).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).Properties(properties).After(after).Before(before).Limit(limit).Archived(archived).Execute()

Get current domains



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
    createdAt := int64(789) // int64 | Only return domains created at this date. (optional)
    createdAfter := int64(789) // int64 | Only return domains created after this date. (optional)
    createdBefore := int64(789) // int64 | Only return domains created before this date. (optional)
    updatedAt := int64(789) // int64 | Only return domains updated at this date. (optional)
    updatedAfter := int64(789) // int64 | Only return domains updated after this date. (optional)
    updatedBefore := int64(789) // int64 | Only return domains updated before this date. (optional)
    sort := []string{"Inner_example"} // []string |  (optional)
    properties := []string{"Inner_example"} // []string |  (optional)
    after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
    before := "before_example" // string |  (optional)
    limit := int32(56) // int32 | Maximum number of results per page. (optional)
    archived := true // bool | Whether to return only results that have been archived. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsApi.GetCmsV3DomainsGetPage(context.Background()).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).Properties(properties).After(after).Before(before).Limit(limit).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.GetCmsV3DomainsGetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3DomainsGetPage`: CollectionResponseWithTotalDomain
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.GetCmsV3DomainsGetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3DomainsGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAt** | **int64** | Only return domains created at this date. | 
 **createdAfter** | **int64** | Only return domains created after this date. | 
 **createdBefore** | **int64** | Only return domains created before this date. | 
 **updatedAt** | **int64** | Only return domains updated at this date. | 
 **updatedAfter** | **int64** | Only return domains updated after this date. | 
 **updatedBefore** | **int64** | Only return domains updated before this date. | 
 **sort** | **[]string** |  | 
 **properties** | **[]string** |  | 
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **before** | **string** |  | 
 **limit** | **int32** | Maximum number of results per page. | 
 **archived** | **bool** | Whether to return only results that have been archived. | 

### Return type

[**CollectionResponseWithTotalDomain**](CollectionResponseWithTotalDomain.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

