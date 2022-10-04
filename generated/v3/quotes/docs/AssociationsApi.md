# \AssociationsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive**](AssociationsApi.md#DeleteCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive) | **Delete** /crm/v3/objects/quotes/{quoteId}/associations/{toObjectType}/{toObjectId}/{associationType} | Remove an association between two quotes
[**GetCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeGetAll**](AssociationsApi.md#GetCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeGetAll) | **Get** /crm/v3/objects/quotes/{quoteId}/associations/{toObjectType} | List associations of a quote by type
[**PutCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate**](AssociationsApi.md#PutCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate) | **Put** /crm/v3/objects/quotes/{quoteId}/associations/{toObjectType}/{toObjectId}/{associationType} | Associate a quote with another object



## DeleteCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive

> DeleteCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive(ctx, quoteId, toObjectType, toObjectId, associationType).Execute()

Remove an association between two quotes

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
    quoteId := "quoteId_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    toObjectId := "toObjectId_example" // string | 
    associationType := "associationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsApi.DeleteCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive(context.Background(), quoteId, toObjectType, toObjectId, associationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.DeleteCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 
**toObjectType** | **string** |  | 
**toObjectId** | **string** |  | 
**associationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveRequest struct via the builder pattern


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


## GetCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeGetAll

> CollectionResponseAssociatedIdForwardPaging GetCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeGetAll(ctx, quoteId, toObjectType).After(after).Limit(limit).Execute()

List associations of a quote by type

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
    quoteId := "quoteId_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
    limit := int32(56) // int32 | The maximum number of results to display per page. (optional) (default to 500)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsApi.GetCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeGetAll(context.Background(), quoteId, toObjectType).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.GetCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeGetAll`: CollectionResponseAssociatedIdForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `AssociationsApi.GetCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeGetAllRequest struct via the builder pattern


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


## PutCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate

> SimplePublicObjectWithAssociations PutCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate(ctx, quoteId, toObjectType, toObjectId, associationType).Execute()

Associate a quote with another object

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
    quoteId := "quoteId_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    toObjectId := "toObjectId_example" // string | 
    associationType := "associationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsApi.PutCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate(context.Background(), quoteId, toObjectType, toObjectId, associationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.PutCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate`: SimplePublicObjectWithAssociations
    fmt.Fprintf(os.Stdout, "Response from `AssociationsApi.PutCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 
**toObjectType** | **string** |  | 
**toObjectId** | **string** |  | 
**associationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV3ObjectsQuotesQuoteIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateRequest struct via the builder pattern


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

