# \DefinitionsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeIdArchive**](DefinitionsApi.md#DeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeIdArchive) | **Delete** /crm/v4/associations/{fromObjectType}/{toObjectType}/labels/{associationTypeId} | Delete
[**GetCrmV4AssociationsFromObjectTypeToObjectTypeLabelsGetAll**](DefinitionsApi.md#GetCrmV4AssociationsFromObjectTypeToObjectTypeLabelsGetAll) | **Get** /crm/v4/associations/{fromObjectType}/{toObjectType}/labels | Read
[**PostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreate**](DefinitionsApi.md#PostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreate) | **Post** /crm/v4/associations/{fromObjectType}/{toObjectType}/labels | Create
[**PutCrmV4AssociationsFromObjectTypeToObjectTypeLabelsUpdate**](DefinitionsApi.md#PutCrmV4AssociationsFromObjectTypeToObjectTypeLabelsUpdate) | **Put** /crm/v4/associations/{fromObjectType}/{toObjectType}/labels | Update



## DeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeIdArchive

> DeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeIdArchive(ctx, fromObjectType, toObjectType, associationTypeId).Execute()

Delete



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
    associationTypeId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefinitionsApi.DeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeIdArchive(context.Background(), fromObjectType, toObjectType, associationTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefinitionsApi.DeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeIdArchive``: %v\n", err)
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
**associationTypeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeIdArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV4AssociationsFromObjectTypeToObjectTypeLabelsGetAll

> CollectionResponseAssociationSpecWithLabelNoPaging GetCrmV4AssociationsFromObjectTypeToObjectTypeLabelsGetAll(ctx, fromObjectType, toObjectType).Execute()

Read



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefinitionsApi.GetCrmV4AssociationsFromObjectTypeToObjectTypeLabelsGetAll(context.Background(), fromObjectType, toObjectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefinitionsApi.GetCrmV4AssociationsFromObjectTypeToObjectTypeLabelsGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV4AssociationsFromObjectTypeToObjectTypeLabelsGetAll`: CollectionResponseAssociationSpecWithLabelNoPaging
    fmt.Fprintf(os.Stdout, "Response from `DefinitionsApi.GetCrmV4AssociationsFromObjectTypeToObjectTypeLabelsGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV4AssociationsFromObjectTypeToObjectTypeLabelsGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CollectionResponseAssociationSpecWithLabelNoPaging**](CollectionResponseAssociationSpecWithLabelNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreate

> CollectionResponseAssociationSpecWithLabelNoPaging PostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreate(ctx, fromObjectType, toObjectType).PublicAssociationDefinitionCreateRequest(publicAssociationDefinitionCreateRequest).Execute()

Create



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
    publicAssociationDefinitionCreateRequest := *openapiclient.NewPublicAssociationDefinitionCreateRequest("Label_example", "Name_example") // PublicAssociationDefinitionCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefinitionsApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreate(context.Background(), fromObjectType, toObjectType).PublicAssociationDefinitionCreateRequest(publicAssociationDefinitionCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefinitionsApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreate`: CollectionResponseAssociationSpecWithLabelNoPaging
    fmt.Fprintf(os.Stdout, "Response from `DefinitionsApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **publicAssociationDefinitionCreateRequest** | [**PublicAssociationDefinitionCreateRequest**](PublicAssociationDefinitionCreateRequest.md) |  | 

### Return type

[**CollectionResponseAssociationSpecWithLabelNoPaging**](CollectionResponseAssociationSpecWithLabelNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmV4AssociationsFromObjectTypeToObjectTypeLabelsUpdate

> PutCrmV4AssociationsFromObjectTypeToObjectTypeLabelsUpdate(ctx, fromObjectType, toObjectType).PublicAssociationDefinitionUpdateRequest(publicAssociationDefinitionUpdateRequest).Execute()

Update



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
    publicAssociationDefinitionUpdateRequest := *openapiclient.NewPublicAssociationDefinitionUpdateRequest("Label_example", int32(123)) // PublicAssociationDefinitionUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefinitionsApi.PutCrmV4AssociationsFromObjectTypeToObjectTypeLabelsUpdate(context.Background(), fromObjectType, toObjectType).PublicAssociationDefinitionUpdateRequest(publicAssociationDefinitionUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefinitionsApi.PutCrmV4AssociationsFromObjectTypeToObjectTypeLabelsUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutCrmV4AssociationsFromObjectTypeToObjectTypeLabelsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **publicAssociationDefinitionUpdateRequest** | [**PublicAssociationDefinitionUpdateRequest**](PublicAssociationDefinitionUpdateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

