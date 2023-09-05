# \TypesApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAll**](TypesApi.md#GetAll) | **Get** /crm/v3/associations/{fromObjectType}/{toObjectType}/types | List association types



## GetAll

> CollectionResponsePublicAssociationDefinitionNoPaging GetAll(ctx, fromObjectType, toObjectType).Execute()

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
    resp, r, err := apiClient.TypesApi.GetAll(context.Background(), fromObjectType, toObjectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesApi.GetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAll`: CollectionResponsePublicAssociationDefinitionNoPaging
    fmt.Fprintf(os.Stdout, "Response from `TypesApi.GetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CollectionResponsePublicAssociationDefinitionNoPaging**](CollectionResponsePublicAssociationDefinitionNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

