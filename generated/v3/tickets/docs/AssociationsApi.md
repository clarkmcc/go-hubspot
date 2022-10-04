# \AssociationsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive**](AssociationsApi.md#DeleteCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive) | **Delete** /crm/v3/objects/tickets/{ticketId}/associations/{toObjectType}/{toObjectId}/{associationType} | Remove an association between two tickets
[**GetCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeGetAll**](AssociationsApi.md#GetCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeGetAll) | **Get** /crm/v3/objects/tickets/{ticketId}/associations/{toObjectType} | List associations of a ticket by type
[**PutCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate**](AssociationsApi.md#PutCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate) | **Put** /crm/v3/objects/tickets/{ticketId}/associations/{toObjectType}/{toObjectId}/{associationType} | Associate a ticket with another object



## DeleteCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive

> DeleteCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive(ctx, ticketId, toObjectType, toObjectId, associationType).Execute()

Remove an association between two tickets

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
    ticketId := "ticketId_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    toObjectId := "toObjectId_example" // string | 
    associationType := "associationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsApi.DeleteCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive(context.Background(), ticketId, toObjectType, toObjectId, associationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.DeleteCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **string** |  | 
**toObjectType** | **string** |  | 
**toObjectId** | **string** |  | 
**associationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveRequest struct via the builder pattern


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


## GetCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeGetAll

> CollectionResponseAssociatedIdForwardPaging GetCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeGetAll(ctx, ticketId, toObjectType).After(after).Limit(limit).Execute()

List associations of a ticket by type

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
    ticketId := "ticketId_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
    limit := int32(56) // int32 | The maximum number of results to display per page. (optional) (default to 500)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsApi.GetCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeGetAll(context.Background(), ticketId, toObjectType).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.GetCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeGetAll`: CollectionResponseAssociatedIdForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `AssociationsApi.GetCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeGetAllRequest struct via the builder pattern


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


## PutCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate

> SimplePublicObjectWithAssociations PutCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate(ctx, ticketId, toObjectType, toObjectId, associationType).Execute()

Associate a ticket with another object

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
    ticketId := "ticketId_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    toObjectId := "toObjectId_example" // string | 
    associationType := "associationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsApi.PutCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate(context.Background(), ticketId, toObjectType, toObjectId, associationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.PutCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate`: SimplePublicObjectWithAssociations
    fmt.Fprintf(os.Stdout, "Response from `AssociationsApi.PutCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **string** |  | 
**toObjectType** | **string** |  | 
**toObjectId** | **string** |  | 
**associationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV3ObjectsTicketsTicketIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateRequest struct via the builder pattern


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

