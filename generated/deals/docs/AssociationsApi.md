# \AssociationsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3ObjectsDealsDealIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive**](AssociationsApi.md#DeleteCrmV3ObjectsDealsDealIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive) | **Delete** /crm/v3/objects/deals/{dealId}/associations/{toObjectType}/{toObjectId}/{associationType} | Remove an association between two deals
[**GetCrmV3ObjectsDealsDealIdAssociationsToObjectTypeGetAll**](AssociationsApi.md#GetCrmV3ObjectsDealsDealIdAssociationsToObjectTypeGetAll) | **Get** /crm/v3/objects/deals/{dealId}/associations/{toObjectType} | List associations of a deal by type
[**PutCrmV3ObjectsDealsDealIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate**](AssociationsApi.md#PutCrmV3ObjectsDealsDealIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate) | **Put** /crm/v3/objects/deals/{dealId}/associations/{toObjectType}/{toObjectId}/{associationType} | Associate a deal with another object



## DeleteCrmV3ObjectsDealsDealIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive

> DeleteCrmV3ObjectsDealsDealIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive(ctx, dealId, toObjectType, toObjectId, associationType).Execute()

Remove an association between two deals

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
    dealId := "dealId_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    toObjectId := "toObjectId_example" // string | 
    associationType := "associationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsApi.DeleteCrmV3ObjectsDealsDealIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive(context.Background(), dealId, toObjectType, toObjectId, associationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.DeleteCrmV3ObjectsDealsDealIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealId** | **string** |  | 
**toObjectType** | **string** |  | 
**toObjectId** | **string** |  | 
**associationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3ObjectsDealsDealIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3ObjectsDealsDealIdAssociationsToObjectTypeGetAll

> CollectionResponseAssociatedIdForwardPaging GetCrmV3ObjectsDealsDealIdAssociationsToObjectTypeGetAll(ctx, dealId, toObjectType).After(after).Limit(limit).Execute()

List associations of a deal by type

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
    dealId := "dealId_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
    limit := int32(56) // int32 | The maximum number of results to display per page. (optional) (default to 500)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsApi.GetCrmV3ObjectsDealsDealIdAssociationsToObjectTypeGetAll(context.Background(), dealId, toObjectType).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.GetCrmV3ObjectsDealsDealIdAssociationsToObjectTypeGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3ObjectsDealsDealIdAssociationsToObjectTypeGetAll`: CollectionResponseAssociatedIdForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `AssociationsApi.GetCrmV3ObjectsDealsDealIdAssociationsToObjectTypeGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealId** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ObjectsDealsDealIdAssociationsToObjectTypeGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to display per page. | [default to 500]

### Return type

[**CollectionResponseAssociatedIdForwardPaging**](CollectionResponseAssociatedIdForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmV3ObjectsDealsDealIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate

> SimplePublicObjectWithAssociations PutCrmV3ObjectsDealsDealIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate(ctx, dealId, toObjectType, toObjectId, associationType).Execute()

Associate a deal with another object

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
    dealId := "dealId_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    toObjectId := "toObjectId_example" // string | 
    associationType := "associationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsApi.PutCrmV3ObjectsDealsDealIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate(context.Background(), dealId, toObjectType, toObjectId, associationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.PutCrmV3ObjectsDealsDealIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutCrmV3ObjectsDealsDealIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate`: SimplePublicObjectWithAssociations
    fmt.Fprintf(os.Stdout, "Response from `AssociationsApi.PutCrmV3ObjectsDealsDealIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealId** | **string** |  | 
**toObjectType** | **string** |  | 
**toObjectId** | **string** |  | 
**associationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV3ObjectsDealsDealIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**SimplePublicObjectWithAssociations**](SimplePublicObjectWithAssociations.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

