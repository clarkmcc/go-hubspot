# \SubscriptionsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionsArchive**](SubscriptionsApi.md#SubscriptionsArchive) | **Delete** /webhooks/v3/{appId}/subscriptions/{subscriptionId} | 
[**SubscriptionsBatchUpdate**](SubscriptionsApi.md#SubscriptionsBatchUpdate) | **Post** /webhooks/v3/{appId}/subscriptions/batch/update | 
[**SubscriptionsCreate**](SubscriptionsApi.md#SubscriptionsCreate) | **Post** /webhooks/v3/{appId}/subscriptions | 
[**SubscriptionsGetAll**](SubscriptionsApi.md#SubscriptionsGetAll) | **Get** /webhooks/v3/{appId}/subscriptions | 
[**SubscriptionsGetByID**](SubscriptionsApi.md#SubscriptionsGetByID) | **Get** /webhooks/v3/{appId}/subscriptions/{subscriptionId} | 
[**SubscriptionsUpdate**](SubscriptionsApi.md#SubscriptionsUpdate) | **Patch** /webhooks/v3/{appId}/subscriptions/{subscriptionId} | 



## SubscriptionsArchive

> SubscriptionsArchive(ctx, subscriptionId, appId).Execute()



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
    subscriptionId := int32(56) // int32 | 
    appId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsApi.SubscriptionsArchive(context.Background(), subscriptionId, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.SubscriptionsArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **int32** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsBatchUpdate

> BatchResponseSubscriptionResponse SubscriptionsBatchUpdate(ctx, appId).BatchInputSubscriptionBatchUpdateRequest(batchInputSubscriptionBatchUpdateRequest).Execute()



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
    appId := int32(56) // int32 | 
    batchInputSubscriptionBatchUpdateRequest := *openapiclient.NewBatchInputSubscriptionBatchUpdateRequest([]openapiclient.SubscriptionBatchUpdateRequest{*openapiclient.NewSubscriptionBatchUpdateRequest(int32(123), false)}) // BatchInputSubscriptionBatchUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsApi.SubscriptionsBatchUpdate(context.Background(), appId).BatchInputSubscriptionBatchUpdateRequest(batchInputSubscriptionBatchUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.SubscriptionsBatchUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsBatchUpdate`: BatchResponseSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.SubscriptionsBatchUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsBatchUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputSubscriptionBatchUpdateRequest** | [**BatchInputSubscriptionBatchUpdateRequest**](BatchInputSubscriptionBatchUpdateRequest.md) |  | 

### Return type

[**BatchResponseSubscriptionResponse**](BatchResponseSubscriptionResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsCreate

> SubscriptionResponse SubscriptionsCreate(ctx, appId).SubscriptionCreateRequest(subscriptionCreateRequest).Execute()



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
    appId := int32(56) // int32 | 
    subscriptionCreateRequest := *openapiclient.NewSubscriptionCreateRequest("EventType_example") // SubscriptionCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsApi.SubscriptionsCreate(context.Background(), appId).SubscriptionCreateRequest(subscriptionCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.SubscriptionsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsCreate`: SubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.SubscriptionsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionCreateRequest** | [**SubscriptionCreateRequest**](SubscriptionCreateRequest.md) |  | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsGetAll

> SubscriptionListResponse SubscriptionsGetAll(ctx, appId).Execute()



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
    appId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsApi.SubscriptionsGetAll(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.SubscriptionsGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsGetAll`: SubscriptionListResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.SubscriptionsGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriptionListResponse**](SubscriptionListResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsGetByID

> SubscriptionResponse SubscriptionsGetByID(ctx, subscriptionId, appId).Execute()



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
    subscriptionId := int32(56) // int32 | 
    appId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsApi.SubscriptionsGetByID(context.Background(), subscriptionId, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.SubscriptionsGetByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsGetByID`: SubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.SubscriptionsGetByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **int32** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsGetByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsUpdate

> SubscriptionResponse SubscriptionsUpdate(ctx, subscriptionId, appId).SubscriptionPatchRequest(subscriptionPatchRequest).Execute()



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
    subscriptionId := int32(56) // int32 | 
    appId := int32(56) // int32 | 
    subscriptionPatchRequest := *openapiclient.NewSubscriptionPatchRequest() // SubscriptionPatchRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsApi.SubscriptionsUpdate(context.Background(), subscriptionId, appId).SubscriptionPatchRequest(subscriptionPatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.SubscriptionsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsUpdate`: SubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.SubscriptionsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **int32** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **subscriptionPatchRequest** | [**SubscriptionPatchRequest**](SubscriptionPatchRequest.md) |  | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

