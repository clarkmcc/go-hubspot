# \StatusApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCommunicationPreferencesV3StatusEmailEmailAddressGetEmailStatus**](StatusApi.md#GetCommunicationPreferencesV3StatusEmailEmailAddressGetEmailStatus) | **Get** /communication-preferences/v3/status/email/{emailAddress} | Get subscription statuses for a contact
[**PostCommunicationPreferencesV3SubscribeSubscribe**](StatusApi.md#PostCommunicationPreferencesV3SubscribeSubscribe) | **Post** /communication-preferences/v3/subscribe | Subscribe a contact
[**PostCommunicationPreferencesV3UnsubscribeUnsubscribe**](StatusApi.md#PostCommunicationPreferencesV3UnsubscribeUnsubscribe) | **Post** /communication-preferences/v3/unsubscribe | Unsubscribe a contact



## GetCommunicationPreferencesV3StatusEmailEmailAddressGetEmailStatus

> PublicSubscriptionStatusesResponse GetCommunicationPreferencesV3StatusEmailEmailAddressGetEmailStatus(ctx, emailAddress).Execute()

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
    resp, r, err := apiClient.StatusApi.GetCommunicationPreferencesV3StatusEmailEmailAddressGetEmailStatus(context.Background(), emailAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatusApi.GetCommunicationPreferencesV3StatusEmailEmailAddressGetEmailStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommunicationPreferencesV3StatusEmailEmailAddressGetEmailStatus`: PublicSubscriptionStatusesResponse
    fmt.Fprintf(os.Stdout, "Response from `StatusApi.GetCommunicationPreferencesV3StatusEmailEmailAddressGetEmailStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailAddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommunicationPreferencesV3StatusEmailEmailAddressGetEmailStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicSubscriptionStatusesResponse**](PublicSubscriptionStatusesResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommunicationPreferencesV3SubscribeSubscribe

> PublicSubscriptionStatus PostCommunicationPreferencesV3SubscribeSubscribe(ctx).PublicUpdateSubscriptionStatusRequest(publicUpdateSubscriptionStatusRequest).Execute()

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
    resp, r, err := apiClient.StatusApi.PostCommunicationPreferencesV3SubscribeSubscribe(context.Background()).PublicUpdateSubscriptionStatusRequest(publicUpdateSubscriptionStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatusApi.PostCommunicationPreferencesV3SubscribeSubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCommunicationPreferencesV3SubscribeSubscribe`: PublicSubscriptionStatus
    fmt.Fprintf(os.Stdout, "Response from `StatusApi.PostCommunicationPreferencesV3SubscribeSubscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCommunicationPreferencesV3SubscribeSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicUpdateSubscriptionStatusRequest** | [**PublicUpdateSubscriptionStatusRequest**](PublicUpdateSubscriptionStatusRequest.md) |  | 

### Return type

[**PublicSubscriptionStatus**](PublicSubscriptionStatus.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommunicationPreferencesV3UnsubscribeUnsubscribe

> PublicSubscriptionStatus PostCommunicationPreferencesV3UnsubscribeUnsubscribe(ctx).PublicUpdateSubscriptionStatusRequest(publicUpdateSubscriptionStatusRequest).Execute()

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
    resp, r, err := apiClient.StatusApi.PostCommunicationPreferencesV3UnsubscribeUnsubscribe(context.Background()).PublicUpdateSubscriptionStatusRequest(publicUpdateSubscriptionStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatusApi.PostCommunicationPreferencesV3UnsubscribeUnsubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCommunicationPreferencesV3UnsubscribeUnsubscribe`: PublicSubscriptionStatus
    fmt.Fprintf(os.Stdout, "Response from `StatusApi.PostCommunicationPreferencesV3UnsubscribeUnsubscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCommunicationPreferencesV3UnsubscribeUnsubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicUpdateSubscriptionStatusRequest** | [**PublicUpdateSubscriptionStatusRequest**](PublicUpdateSubscriptionStatusRequest.md) |  | 

### Return type

[**PublicSubscriptionStatus**](PublicSubscriptionStatus.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

