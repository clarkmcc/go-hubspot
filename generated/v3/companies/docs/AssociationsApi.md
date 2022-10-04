# \AssociationsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociationsArchive**](AssociationsApi.md#AssociationsArchive) | **Delete** /crm/v3/objects/companies/{companyId}/associations/{toObjectType}/{toObjectId}/{associationType} | Remove an association between two companies
[**AssociationsCreate**](AssociationsApi.md#AssociationsCreate) | **Put** /crm/v3/objects/companies/{companyId}/associations/{toObjectType}/{toObjectId}/{associationType} | Associate a company with another object
[**AssociationsGet**](AssociationsApi.md#AssociationsGet) | **Get** /crm/v3/objects/companies/{companyId}/associations/{toObjectType} | List associations of a company by type



## AssociationsArchive

> AssociationsArchive(ctx, companyId, toObjectType, toObjectId, associationType).Execute()

Remove an association between two companies

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
    companyId := "companyId_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    toObjectId := "toObjectId_example" // string | 
    associationType := "associationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsApi.AssociationsArchive(context.Background(), companyId, toObjectType, toObjectId, associationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.AssociationsArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**toObjectType** | **string** |  | 
**toObjectId** | **string** |  | 
**associationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssociationsArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssociationsCreate

> SimplePublicObjectWithAssociations AssociationsCreate(ctx, companyId, toObjectType, toObjectId, associationType).Execute()

Associate a company with another object

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
    companyId := "companyId_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    toObjectId := "toObjectId_example" // string | 
    associationType := "associationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsApi.AssociationsCreate(context.Background(), companyId, toObjectType, toObjectId, associationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.AssociationsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssociationsCreate`: SimplePublicObjectWithAssociations
    fmt.Fprintf(os.Stdout, "Response from `AssociationsApi.AssociationsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**toObjectType** | **string** |  | 
**toObjectId** | **string** |  | 
**associationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssociationsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**SimplePublicObjectWithAssociations**](SimplePublicObjectWithAssociations.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssociationsGet

> CollectionResponseAssociatedIdForwardPaging AssociationsGet(ctx, companyId, toObjectType).After(after).Limit(limit).Execute()

List associations of a company by type

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
    companyId := "companyId_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
    limit := int32(56) // int32 | The maximum number of results to display per page. (optional) (default to 500)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsApi.AssociationsGet(context.Background(), companyId, toObjectType).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.AssociationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssociationsGet`: CollectionResponseAssociatedIdForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `AssociationsApi.AssociationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssociationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to display per page. | [default to 500]

### Return type

[**CollectionResponseAssociatedIdForwardPaging**](CollectionResponseAssociatedIdForwardPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

