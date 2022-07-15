# \AssociationsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3ObjectsContactsContactIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive**](AssociationsApi.md#DeleteCrmV3ObjectsContactsContactIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive) | **Delete** /crm/v3/objects/contacts/{contactId}/associations/{toObjectType}/{toObjectId}/{associationType} | Remove an association between two contacts
[**GetCrmV3ObjectsContactsContactIdAssociationsToObjectTypeGetAll**](AssociationsApi.md#GetCrmV3ObjectsContactsContactIdAssociationsToObjectTypeGetAll) | **Get** /crm/v3/objects/contacts/{contactId}/associations/{toObjectType} | List associations of a contact by type
[**PutCrmV3ObjectsContactsContactIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate**](AssociationsApi.md#PutCrmV3ObjectsContactsContactIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate) | **Put** /crm/v3/objects/contacts/{contactId}/associations/{toObjectType}/{toObjectId}/{associationType} | Associate a contact with another object



## DeleteCrmV3ObjectsContactsContactIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive

> DeleteCrmV3ObjectsContactsContactIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive(ctx, contactId, toObjectType, toObjectId, associationType).Execute()

Remove an association between two contacts

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
    contactId := "contactId_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    toObjectId := "toObjectId_example" // string | 
    associationType := "associationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsApi.DeleteCrmV3ObjectsContactsContactIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive(context.Background(), contactId, toObjectType, toObjectId, associationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.DeleteCrmV3ObjectsContactsContactIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 
**toObjectType** | **string** |  | 
**toObjectId** | **string** |  | 
**associationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3ObjectsContactsContactIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveRequest struct via the builder pattern


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


## GetCrmV3ObjectsContactsContactIdAssociationsToObjectTypeGetAll

> CollectionResponseAssociatedIdForwardPaging GetCrmV3ObjectsContactsContactIdAssociationsToObjectTypeGetAll(ctx, contactId, toObjectType).After(after).Limit(limit).Execute()

List associations of a contact by type

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
    contactId := "contactId_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
    limit := int32(56) // int32 | The maximum number of results to display per page. (optional) (default to 500)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsApi.GetCrmV3ObjectsContactsContactIdAssociationsToObjectTypeGetAll(context.Background(), contactId, toObjectType).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.GetCrmV3ObjectsContactsContactIdAssociationsToObjectTypeGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3ObjectsContactsContactIdAssociationsToObjectTypeGetAll`: CollectionResponseAssociatedIdForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `AssociationsApi.GetCrmV3ObjectsContactsContactIdAssociationsToObjectTypeGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ObjectsContactsContactIdAssociationsToObjectTypeGetAllRequest struct via the builder pattern


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


## PutCrmV3ObjectsContactsContactIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate

> SimplePublicObjectWithAssociations PutCrmV3ObjectsContactsContactIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate(ctx, contactId, toObjectType, toObjectId, associationType).Execute()

Associate a contact with another object

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
    contactId := "contactId_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    toObjectId := "toObjectId_example" // string | 
    associationType := "associationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsApi.PutCrmV3ObjectsContactsContactIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate(context.Background(), contactId, toObjectType, toObjectId, associationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.PutCrmV3ObjectsContactsContactIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutCrmV3ObjectsContactsContactIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate`: SimplePublicObjectWithAssociations
    fmt.Fprintf(os.Stdout, "Response from `AssociationsApi.PutCrmV3ObjectsContactsContactIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 
**toObjectType** | **string** |  | 
**toObjectId** | **string** |  | 
**associationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV3ObjectsContactsContactIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateRequest struct via the builder pattern


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

