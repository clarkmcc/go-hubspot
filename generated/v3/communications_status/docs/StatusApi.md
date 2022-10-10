# \StatusApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEmailStatus**](StatusApi.md#GetEmailStatus) | **Get** /communication-preferences/v3/status/email/{emailAddress} | Get subscription statuses for a contact
[**Subscribe**](StatusApi.md#Subscribe) | **Post** /communication-preferences/v3/subscribe | Subscribe a contact
[**Unsubscribe**](StatusApi.md#Unsubscribe) | **Post** /communication-preferences/v3/unsubscribe | Unsubscribe a contact



## GetEmailStatus

> PublicSubscriptionStatusesResponse GetEmailStatus(ctx, emailAddress).Execute()

Get subscription statuses for a contact



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
    emailAddress := "emailAddress_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatusApi.GetEmailStatus(context.Background(), emailAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatusApi.GetEmailStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailStatus`: PublicSubscriptionStatusesResponse
    fmt.Fprintf(os.Stdout, "Response from `StatusApi.GetEmailStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailAddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicSubscriptionStatusesResponse**](PublicSubscriptionStatusesResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Subscribe

> PublicSubscriptionStatus Subscribe(ctx).PublicUpdateSubscriptionStatusRequest(publicUpdateSubscriptionStatusRequest).Execute()

Subscribe a contact



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
    publicUpdateSubscriptionStatusRequest := *openapiclient.NewPublicUpdateSubscriptionStatusRequest("EmailAddress_example", "SubscriptionId_example") // PublicUpdateSubscriptionStatusRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatusApi.Subscribe(context.Background()).PublicUpdateSubscriptionStatusRequest(publicUpdateSubscriptionStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatusApi.Subscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Subscribe`: PublicSubscriptionStatus
    fmt.Fprintf(os.Stdout, "Response from `StatusApi.Subscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicUpdateSubscriptionStatusRequest** | [**PublicUpdateSubscriptionStatusRequest**](PublicUpdateSubscriptionStatusRequest.md) |  | 

### Return type

[**PublicSubscriptionStatus**](PublicSubscriptionStatus.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Unsubscribe

> PublicSubscriptionStatus Unsubscribe(ctx).PublicUpdateSubscriptionStatusRequest(publicUpdateSubscriptionStatusRequest).Execute()

Unsubscribe a contact



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
    publicUpdateSubscriptionStatusRequest := *openapiclient.NewPublicUpdateSubscriptionStatusRequest("EmailAddress_example", "SubscriptionId_example") // PublicUpdateSubscriptionStatusRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatusApi.Unsubscribe(context.Background()).PublicUpdateSubscriptionStatusRequest(publicUpdateSubscriptionStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatusApi.Unsubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Unsubscribe`: PublicSubscriptionStatus
    fmt.Fprintf(os.Stdout, "Response from `StatusApi.Unsubscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicUpdateSubscriptionStatusRequest** | [**PublicUpdateSubscriptionStatusRequest**](PublicUpdateSubscriptionStatusRequest.md) |  | 

### Return type

[**PublicSubscriptionStatus**](PublicSubscriptionStatus.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

