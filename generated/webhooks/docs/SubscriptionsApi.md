# \SubscriptionsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchive**](SubscriptionsApi.md#DeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchive) | **Delete** /webhooks/v3/{appId}/subscriptions/{subscriptionId} | Delete a subscription
[**GetWebhooksV3AppIdSubscriptionsGetAll**](SubscriptionsApi.md#GetWebhooksV3AppIdSubscriptionsGetAll) | **Get** /webhooks/v3/{appId}/subscriptions | Get subscription details
[**GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetById**](SubscriptionsApi.md#GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetById) | **Get** /webhooks/v3/{appId}/subscriptions/{subscriptionId} | Get subscription
[**PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdate**](SubscriptionsApi.md#PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdate) | **Patch** /webhooks/v3/{appId}/subscriptions/{subscriptionId} | Update a subscription
[**PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatch**](SubscriptionsApi.md#PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatch) | **Post** /webhooks/v3/{appId}/subscriptions/batch/update | Batch update subscriptions
[**PostWebhooksV3AppIdSubscriptionsCreate**](SubscriptionsApi.md#PostWebhooksV3AppIdSubscriptionsCreate) | **Post** /webhooks/v3/{appId}/subscriptions | Subscribe to an event



## DeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchive

> DeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchive(ctx, subscriptionId, appId).Execute()

Delete a subscription



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
    subscriptionId := int32(56) // int32 | The ID of subscription to delete.
    appId := int32(56) // int32 | The ID of the target app.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsApi.DeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchive(context.Background(), subscriptionId, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.DeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **int32** | The ID of subscription to delete. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooksV3AppIdSubscriptionsGetAll

> SubscriptionListResponse GetWebhooksV3AppIdSubscriptionsGetAll(ctx, appId).Execute()

Get subscription details



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
    appId := int32(56) // int32 | The ID of the target app.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsApi.GetWebhooksV3AppIdSubscriptionsGetAll(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.GetWebhooksV3AppIdSubscriptionsGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhooksV3AppIdSubscriptionsGetAll`: SubscriptionListResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.GetWebhooksV3AppIdSubscriptionsGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksV3AppIdSubscriptionsGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriptionListResponse**](SubscriptionListResponse.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetById

> SubscriptionResponse GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetById(ctx, appId, subscriptionId).Execute()

Get subscription



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
    appId := int32(56) // int32 | The ID of the target app.
    subscriptionId := int32(56) // int32 | The ID of the target subscription.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsApi.GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetById(context.Background(), appId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetById`: SubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 
**subscriptionId** | **int32** | The ID of the target subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksV3AppIdSubscriptionsSubscriptionIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdate

> SubscriptionResponse PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdate(ctx, subscriptionId, appId).SubscriptionPatchRequest(subscriptionPatchRequest).Execute()

Update a subscription



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
    subscriptionId := int32(56) // int32 | The ID of the subscription to update.
    appId := int32(56) // int32 | The ID of the target app.
    subscriptionPatchRequest := *openapiclient.NewSubscriptionPatchRequest() // SubscriptionPatchRequest | Updated details for the subscription.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsApi.PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdate(context.Background(), subscriptionId, appId).SubscriptionPatchRequest(subscriptionPatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdate`: SubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **int32** | The ID of the subscription to update. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **subscriptionPatchRequest** | [**SubscriptionPatchRequest**](SubscriptionPatchRequest.md) | Updated details for the subscription. | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatch

> BatchResponseSubscriptionResponse PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatch(ctx, appId).BatchInputSubscriptionBatchUpdateRequest(batchInputSubscriptionBatchUpdateRequest).Execute()

Batch update subscriptions



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
    appId := int32(56) // int32 | The app ID of the target app.
    batchInputSubscriptionBatchUpdateRequest := *openapiclient.NewBatchInputSubscriptionBatchUpdateRequest([]openapiclient.SubscriptionBatchUpdateRequest{*openapiclient.NewSubscriptionBatchUpdateRequest(int32(123), false)}) // BatchInputSubscriptionBatchUpdateRequest | Updated details for the specified subscriptions.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsApi.PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatch(context.Background(), appId).BatchInputSubscriptionBatchUpdateRequest(batchInputSubscriptionBatchUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatch`: BatchResponseSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The app ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputSubscriptionBatchUpdateRequest** | [**BatchInputSubscriptionBatchUpdateRequest**](BatchInputSubscriptionBatchUpdateRequest.md) | Updated details for the specified subscriptions. | 

### Return type

[**BatchResponseSubscriptionResponse**](BatchResponseSubscriptionResponse.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWebhooksV3AppIdSubscriptionsCreate

> SubscriptionResponse PostWebhooksV3AppIdSubscriptionsCreate(ctx, appId).SubscriptionCreateRequest(subscriptionCreateRequest).Execute()

Subscribe to an event



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
    appId := int32(56) // int32 | The ID of the target app.
    subscriptionCreateRequest := *openapiclient.NewSubscriptionCreateRequest("EventType_example") // SubscriptionCreateRequest | Details about the new subscription.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsApi.PostWebhooksV3AppIdSubscriptionsCreate(context.Background(), appId).SubscriptionCreateRequest(subscriptionCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.PostWebhooksV3AppIdSubscriptionsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostWebhooksV3AppIdSubscriptionsCreate`: SubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.PostWebhooksV3AppIdSubscriptionsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWebhooksV3AppIdSubscriptionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionCreateRequest** | [**SubscriptionCreateRequest**](SubscriptionCreateRequest.md) | Details about the new subscription. | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

