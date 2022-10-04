# \BatchApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchArchiveArchive**](BatchApi.md#PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchArchiveArchive) | **Post** /crm/v3/associations/{fromObjectType}/{toObjectType}/batch/archive | Archive a batch of associations
[**PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchCreateCreate**](BatchApi.md#PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchCreateCreate) | **Post** /crm/v3/associations/{fromObjectType}/{toObjectType}/batch/create | Create a batch of associations
[**PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchReadRead**](BatchApi.md#PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchReadRead) | **Post** /crm/v3/associations/{fromObjectType}/{toObjectType}/batch/read | Read a batch of associations



## PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchArchiveArchive

> PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchArchiveArchive(ctx, fromObjectType, toObjectType).BatchInputPublicAssociation(batchInputPublicAssociation).Execute()

Archive a batch of associations



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
    fromObjectType := "fromObjectType_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    batchInputPublicAssociation := *openapiclient.NewBatchInputPublicAssociation([]openapiclient.PublicAssociation{*openapiclient.NewPublicAssociation(*openapiclient.NewPublicObjectId("Id_example"), *openapiclient.NewPublicObjectId("Id_example"), "Type_example")}) // BatchInputPublicAssociation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchArchiveArchive(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociation(batchInputPublicAssociation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchArchiveArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3AssociationsFromObjectTypeToObjectTypeBatchArchiveArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicAssociation** | [**BatchInputPublicAssociation**](BatchInputPublicAssociation.md) |  | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchCreateCreate

> BatchResponsePublicAssociation PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchCreateCreate(ctx, fromObjectType, toObjectType).BatchInputPublicAssociation(batchInputPublicAssociation).Execute()

Create a batch of associations



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
    fromObjectType := "fromObjectType_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    batchInputPublicAssociation := *openapiclient.NewBatchInputPublicAssociation([]openapiclient.PublicAssociation{*openapiclient.NewPublicAssociation(*openapiclient.NewPublicObjectId("Id_example"), *openapiclient.NewPublicObjectId("Id_example"), "Type_example")}) // BatchInputPublicAssociation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchCreateCreate(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociation(batchInputPublicAssociation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchCreateCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchCreateCreate`: BatchResponsePublicAssociation
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchCreateCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3AssociationsFromObjectTypeToObjectTypeBatchCreateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicAssociation** | [**BatchInputPublicAssociation**](BatchInputPublicAssociation.md) |  | 

### Return type

[**BatchResponsePublicAssociation**](BatchResponsePublicAssociation.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchReadRead

> BatchResponsePublicAssociationMulti PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchReadRead(ctx, fromObjectType, toObjectType).BatchInputPublicObjectId(batchInputPublicObjectId).Execute()

Read a batch of associations



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
    fromObjectType := "fromObjectType_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    batchInputPublicObjectId := *openapiclient.NewBatchInputPublicObjectId([]openapiclient.PublicObjectId{*openapiclient.NewPublicObjectId("Id_example")}) // BatchInputPublicObjectId |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchReadRead(context.Background(), fromObjectType, toObjectType).BatchInputPublicObjectId(batchInputPublicObjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchReadRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchReadRead`: BatchResponsePublicAssociationMulti
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchReadRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3AssociationsFromObjectTypeToObjectTypeBatchReadReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicObjectId** | [**BatchInputPublicObjectId**](BatchInputPublicObjectId.md) |  | 

### Return type

[**BatchResponsePublicAssociationMulti**](BatchResponsePublicAssociationMulti.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

