# \AssociationsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociationsArchive**](AssociationsApi.md#AssociationsArchive) | **Delete** /crm/v3/objects/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId}/{associationType} | Remove an association between two objects
[**AssociationsCreate**](AssociationsApi.md#AssociationsCreate) | **Put** /crm/v3/objects/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId}/{associationType} | Associate an object with another object
[**AssociationsGetAll**](AssociationsApi.md#AssociationsGetAll) | **Get** /crm/v3/objects/{objectType}/{objectId}/associations/{toObjectType} | List associations of an object by type



## AssociationsArchive

> AssociationsArchive(ctx, objectType, objectId, toObjectType, toObjectId, associationType).Execute()

Remove an association between two objects

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
    objectType := "objectType_example" // string | 
    objectId := "objectId_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    toObjectId := "toObjectId_example" // string | 
    associationType := "associationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsApi.AssociationsArchive(context.Background(), objectType, objectId, toObjectType, toObjectId, associationType).Execute()
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
**objectType** | **string** |  | 
**objectId** | **string** |  | 
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

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssociationsCreate

> SimplePublicObjectWithAssociations AssociationsCreate(ctx, objectType, objectId, toObjectType, toObjectId, associationType).Execute()

Associate an object with another object

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
    objectType := "objectType_example" // string | 
    objectId := "objectId_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    toObjectId := "toObjectId_example" // string | 
    associationType := "associationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsApi.AssociationsCreate(context.Background(), objectType, objectId, toObjectType, toObjectId, associationType).Execute()
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
**objectType** | **string** |  | 
**objectId** | **string** |  | 
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

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssociationsGetAll

> CollectionResponseAssociatedIdForwardPaging AssociationsGetAll(ctx, objectType, objectId, toObjectType).After(after).Limit(limit).Execute()

List associations of an object by type

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
    objectType := "objectType_example" // string | 
    objectId := "objectId_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
    limit := int32(56) // int32 | The maximum number of results to display per page. (optional) (default to 500)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsApi.AssociationsGetAll(context.Background(), objectType, objectId, toObjectType).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.AssociationsGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssociationsGetAll`: CollectionResponseAssociatedIdForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `AssociationsApi.AssociationsGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**objectId** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssociationsGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to display per page. | [default to 500]

### Return type

[**CollectionResponseAssociatedIdForwardPaging**](CollectionResponseAssociatedIdForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

