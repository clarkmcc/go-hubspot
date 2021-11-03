# \AssociationsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive**](AssociationsApi.md#DeleteCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive) | **Delete** /crm/v3/objects/line_items/{lineItemId}/associations/{toObjectType}/{toObjectId}/{associationType} | Remove an association between two line items
[**GetCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeGetAll**](AssociationsApi.md#GetCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeGetAll) | **Get** /crm/v3/objects/line_items/{lineItemId}/associations/{toObjectType} | List associations of a line item by type
[**PutCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate**](AssociationsApi.md#PutCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate) | **Put** /crm/v3/objects/line_items/{lineItemId}/associations/{toObjectType}/{toObjectId}/{associationType} | Associate a line item with another object



## DeleteCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive

> DeleteCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive(ctx, lineItemId, toObjectType, toObjectId, associationType).Execute()

Remove an association between two line items

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
    lineItemId := "lineItemId_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    toObjectId := "toObjectId_example" // string | 
    associationType := "associationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssociationsApi.DeleteCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive(context.Background(), lineItemId, toObjectType, toObjectId, associationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.DeleteCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineItemId** | **string** |  | 
**toObjectType** | **string** |  | 
**toObjectId** | **string** |  | 
**associationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveRequest struct via the builder pattern


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


## GetCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeGetAll

> CollectionResponseAssociatedIdForwardPaging GetCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeGetAll(ctx, lineItemId, toObjectType).After(after).Limit(limit).Execute()

List associations of a line item by type

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
    lineItemId := "lineItemId_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
    limit := int32(56) // int32 | The maximum number of results to display per page. (optional) (default to 500)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssociationsApi.GetCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeGetAll(context.Background(), lineItemId, toObjectType).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.GetCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeGetAll`: CollectionResponseAssociatedIdForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `AssociationsApi.GetCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineItemId** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to display per page. | [default to 500]

### Return type

[**CollectionResponseAssociatedIdForwardPaging**](CollectionResponseAssociatedIdForwardPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate

> SimplePublicObjectWithAssociations PutCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate(ctx, lineItemId, toObjectType, toObjectId, associationType).Execute()

Associate a line item with another object

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
    lineItemId := "lineItemId_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    toObjectId := "toObjectId_example" // string | 
    associationType := "associationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssociationsApi.PutCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate(context.Background(), lineItemId, toObjectType, toObjectId, associationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.PutCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate`: SimplePublicObjectWithAssociations
    fmt.Fprintf(os.Stdout, "Response from `AssociationsApi.PutCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineItemId** | **string** |  | 
**toObjectType** | **string** |  | 
**toObjectId** | **string** |  | 
**associationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV3ObjectsLineItemsLineItemIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**SimplePublicObjectWithAssociations**](SimplePublicObjectWithAssociations.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

