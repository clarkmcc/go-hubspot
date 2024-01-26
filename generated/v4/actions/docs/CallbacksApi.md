# \CallbacksApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallbackComplete**](CallbacksApi.md#CallbackComplete) | **Post** /automation/v4/actions/callbacks/{callbackId}/complete | Completes a single callback
[**CallbackCompleteBatch**](CallbacksApi.md#CallbackCompleteBatch) | **Post** /automation/v4/actions/callbacks/complete | Completes a batch of callbacks



## CallbackComplete

> CallbackComplete(ctx, callbackId).CallbackCompletionRequest(callbackCompletionRequest).Execute()

Completes a single callback

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
    callbackId := "callbackId_example" // string | 
    callbackCompletionRequest := *openapiclient.NewCallbackCompletionRequest(map[string]string{"key": "Inner_example"}) // CallbackCompletionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CallbacksApi.CallbackComplete(context.Background(), callbackId).CallbackCompletionRequest(callbackCompletionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.CallbackComplete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callbackId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCallbackCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **callbackCompletionRequest** | [**CallbackCompletionRequest**](CallbackCompletionRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallbackCompleteBatch

> CallbackCompleteBatch(ctx).BatchInputCallbackCompletionBatchRequest(batchInputCallbackCompletionBatchRequest).Execute()

Completes a batch of callbacks

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
    batchInputCallbackCompletionBatchRequest := *openapiclient.NewBatchInputCallbackCompletionBatchRequest([]openapiclient.CallbackCompletionBatchRequest{*openapiclient.NewCallbackCompletionBatchRequest(map[string]string{"key": "Inner_example"}, "CallbackId_example")}) // BatchInputCallbackCompletionBatchRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CallbacksApi.CallbackCompleteBatch(context.Background()).BatchInputCallbackCompletionBatchRequest(batchInputCallbackCompletionBatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.CallbackCompleteBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallbackCompleteBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputCallbackCompletionBatchRequest** | [**BatchInputCallbackCompletionBatchRequest**](BatchInputCallbackCompletionBatchRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

