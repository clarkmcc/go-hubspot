# \TypesApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCrmV3AssociationsFromObjectTypeToObjectTypeTypesGetAll**](TypesApi.md#GetCrmV3AssociationsFromObjectTypeToObjectTypeTypesGetAll) | **Get** /crm/v3/associations/{fromObjectType}/{toObjectType}/types | List association types



## GetCrmV3AssociationsFromObjectTypeToObjectTypeTypesGetAll

> CollectionResponsePublicAssociationDefiniton GetCrmV3AssociationsFromObjectTypeToObjectTypeTypesGetAll(ctx, fromObjectType, toObjectType).Execute()

List association types



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
    resp, r, err := apiClient.TypesApi.GetCrmV3AssociationsFromObjectTypeToObjectTypeTypesGetAll(context.Background(), fromObjectType, toObjectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesApi.GetCrmV3AssociationsFromObjectTypeToObjectTypeTypesGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3AssociationsFromObjectTypeToObjectTypeTypesGetAll`: CollectionResponsePublicAssociationDefiniton
    fmt.Fprintf(os.Stdout, "Response from `TypesApi.GetCrmV3AssociationsFromObjectTypeToObjectTypeTypesGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3AssociationsFromObjectTypeToObjectTypeTypesGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CollectionResponsePublicAssociationDefiniton**](CollectionResponsePublicAssociationDefiniton.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

