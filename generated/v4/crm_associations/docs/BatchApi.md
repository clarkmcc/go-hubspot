# \BatchApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveArchive**](BatchApi.md#PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveArchive) | **Post** /crm/v4/associations/{fromObjectType}/{toObjectType}/batch/archive | Delete
[**PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefault**](BatchApi.md#PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefault) | **Post** /crm/v4/associations/{fromObjectType}/{toObjectType}/batch/associate/default |  Create Default Associations
[**PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreate**](BatchApi.md#PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreate) | **Post** /crm/v4/associations/{fromObjectType}/{toObjectType}/batch/create | Create
[**PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveArchiveLabels**](BatchApi.md#PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveArchiveLabels) | **Post** /crm/v4/associations/{fromObjectType}/{toObjectType}/batch/labels/archive | Delete Specific Labels
[**PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPage**](BatchApi.md#PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPage) | **Post** /crm/v4/associations/{fromObjectType}/{toObjectType}/batch/read | Read



## PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveArchive

> PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveArchive(ctx, fromObjectType, toObjectType).BatchInputPublicAssociationMultiArchive(batchInputPublicAssociationMultiArchive).Execute()

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
    batchInputPublicAssociationMultiArchive := *openapiclient.NewBatchInputPublicAssociationMultiArchive([]openapiclient.PublicAssociationMultiArchive{*openapiclient.NewPublicAssociationMultiArchive(*openapiclient.NewPublicObjectId("Id_example"), []openapiclient.PublicObjectId{*openapiclient.NewPublicObjectId("Id_example")})}) // BatchInputPublicAssociationMultiArchive | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveArchive(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociationMultiArchive(batchInputPublicAssociationMultiArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicAssociationMultiArchive** | [**BatchInputPublicAssociationMultiArchive**](BatchInputPublicAssociationMultiArchive.md) |  | 

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


## PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefault

> BatchResponsePublicDefaultAssociation PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefault(ctx, fromObjectType, toObjectType).BatchInputPublicDefaultAssociationMultiPost(batchInputPublicDefaultAssociationMultiPost).Execute()

 Create Default Associations



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
    batchInputPublicDefaultAssociationMultiPost := *openapiclient.NewBatchInputPublicDefaultAssociationMultiPost([]openapiclient.PublicDefaultAssociationMultiPost{*openapiclient.NewPublicDefaultAssociationMultiPost(*openapiclient.NewPublicObjectId("Id_example"), *openapiclient.NewPublicObjectId("Id_example"))}) // BatchInputPublicDefaultAssociationMultiPost | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefault(context.Background(), fromObjectType, toObjectType).BatchInputPublicDefaultAssociationMultiPost(batchInputPublicDefaultAssociationMultiPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefault`: BatchResponsePublicDefaultAssociation
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicDefaultAssociationMultiPost** | [**BatchInputPublicDefaultAssociationMultiPost**](BatchInputPublicDefaultAssociationMultiPost.md) |  | 

### Return type

[**BatchResponsePublicDefaultAssociation**](BatchResponsePublicDefaultAssociation.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreate

> BatchResponseLabelsBetweenObjectPair PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreate(ctx, fromObjectType, toObjectType).BatchInputPublicAssociationMultiPost(batchInputPublicAssociationMultiPost).Execute()

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
    batchInputPublicAssociationMultiPost := *openapiclient.NewBatchInputPublicAssociationMultiPost([]openapiclient.PublicAssociationMultiPost{*openapiclient.NewPublicAssociationMultiPost(*openapiclient.NewPublicObjectId("Id_example"), *openapiclient.NewPublicObjectId("Id_example"), []openapiclient.AssociationSpec{*openapiclient.NewAssociationSpec("AssociationCategory_example", int32(123))})}) // BatchInputPublicAssociationMultiPost | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreate(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociationMultiPost(batchInputPublicAssociationMultiPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreate`: BatchResponseLabelsBetweenObjectPair
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicAssociationMultiPost** | [**BatchInputPublicAssociationMultiPost**](BatchInputPublicAssociationMultiPost.md) |  | 

### Return type

[**BatchResponseLabelsBetweenObjectPair**](BatchResponseLabelsBetweenObjectPair.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveArchiveLabels

> PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveArchiveLabels(ctx, fromObjectType, toObjectType).BatchInputPublicAssociationMultiPost(batchInputPublicAssociationMultiPost).Execute()

Delete Specific Labels



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
    batchInputPublicAssociationMultiPost := *openapiclient.NewBatchInputPublicAssociationMultiPost([]openapiclient.PublicAssociationMultiPost{*openapiclient.NewPublicAssociationMultiPost(*openapiclient.NewPublicObjectId("Id_example"), *openapiclient.NewPublicObjectId("Id_example"), []openapiclient.AssociationSpec{*openapiclient.NewAssociationSpec("AssociationCategory_example", int32(123))})}) // BatchInputPublicAssociationMultiPost | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveArchiveLabels(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociationMultiPost(batchInputPublicAssociationMultiPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveArchiveLabels``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveArchiveLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicAssociationMultiPost** | [**BatchInputPublicAssociationMultiPost**](BatchInputPublicAssociationMultiPost.md) |  | 

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


## PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPage

> BatchResponsePublicAssociationMultiWithLabel PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPage(ctx, fromObjectType, toObjectType).BatchInputPublicFetchAssociationsBatchRequest(batchInputPublicFetchAssociationsBatchRequest).Execute()

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
    batchInputPublicFetchAssociationsBatchRequest := *openapiclient.NewBatchInputPublicFetchAssociationsBatchRequest([]openapiclient.PublicFetchAssociationsBatchRequest{*openapiclient.NewPublicFetchAssociationsBatchRequest("Id_example")}) // BatchInputPublicFetchAssociationsBatchRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPage(context.Background(), fromObjectType, toObjectType).BatchInputPublicFetchAssociationsBatchRequest(batchInputPublicFetchAssociationsBatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPage`: BatchResponsePublicAssociationMultiWithLabel
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicFetchAssociationsBatchRequest** | [**BatchInputPublicFetchAssociationsBatchRequest**](BatchInputPublicFetchAssociationsBatchRequest.md) |  | 

### Return type

[**BatchResponsePublicAssociationMultiWithLabel**](BatchResponsePublicAssociationMultiWithLabel.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

