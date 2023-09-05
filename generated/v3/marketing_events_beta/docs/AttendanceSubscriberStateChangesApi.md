# \AttendanceSubscriberStateChangesApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttendanceCreateCreate**](AttendanceSubscriberStateChangesApi.md#AttendanceCreateCreate) | **Post** /marketing/v3/marketing-events/attendance/{externalEventId}/{subscriberState}/create | Record
[**AttendanceEmailCreateCreateByEmail**](AttendanceSubscriberStateChangesApi.md#AttendanceEmailCreateCreateByEmail) | **Post** /marketing/v3/marketing-events/attendance/{externalEventId}/{subscriberState}/email-create | Record



## AttendanceCreateCreate

> BatchResponseSubscriberVidResponse AttendanceCreateCreate(ctx, externalEventId, subscriberState).BatchInputMarketingEventSubscriber(batchInputMarketingEventSubscriber).ExternalAccountId(externalAccountId).Execute()

Record



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
    externalEventId := "externalEventId_example" // string | The id of the marketing event
    subscriberState := "subscriberState_example" // string | The new subscriber state for the HubSpot contacts and the specified marketing event. For example: 'register', 'attend' or 'cancel'.
    batchInputMarketingEventSubscriber := *openapiclient.NewBatchInputMarketingEventSubscriber([]openapiclient.MarketingEventSubscriber{*openapiclient.NewMarketingEventSubscriber(int64(123))}) // BatchInputMarketingEventSubscriber | The details of the contacts to subscribe to the event. Parameters of join and left time if state is Attended.
    externalAccountId := "externalAccountId_example" // string | The account id associated with the marketing event (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttendanceSubscriberStateChangesApi.AttendanceCreateCreate(context.Background(), externalEventId, subscriberState).BatchInputMarketingEventSubscriber(batchInputMarketingEventSubscriber).ExternalAccountId(externalAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttendanceSubscriberStateChangesApi.AttendanceCreateCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttendanceCreateCreate`: BatchResponseSubscriberVidResponse
    fmt.Fprintf(os.Stdout, "Response from `AttendanceSubscriberStateChangesApi.AttendanceCreateCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event | 
**subscriberState** | **string** | The new subscriber state for the HubSpot contacts and the specified marketing event. For example: &#39;register&#39;, &#39;attend&#39; or &#39;cancel&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttendanceCreateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputMarketingEventSubscriber** | [**BatchInputMarketingEventSubscriber**](BatchInputMarketingEventSubscriber.md) | The details of the contacts to subscribe to the event. Parameters of join and left time if state is Attended. | 
 **externalAccountId** | **string** | The account id associated with the marketing event | 

### Return type

[**BatchResponseSubscriberVidResponse**](BatchResponseSubscriberVidResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttendanceEmailCreateCreateByEmail

> BatchResponseSubscriberEmailResponse AttendanceEmailCreateCreateByEmail(ctx, externalEventId, subscriberState).BatchInputMarketingEventEmailSubscriber(batchInputMarketingEventEmailSubscriber).ExternalAccountId(externalAccountId).Execute()

Record



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
    externalEventId := "externalEventId_example" // string | The id of the marketing event
    subscriberState := "subscriberState_example" // string | The new subscriber state for the HubSpot contacts and the specified marketing event. For example: 'register', 'attend' or 'cancel'.
    batchInputMarketingEventEmailSubscriber := *openapiclient.NewBatchInputMarketingEventEmailSubscriber([]openapiclient.MarketingEventEmailSubscriber{*openapiclient.NewMarketingEventEmailSubscriber("Email_example", int64(123))}) // BatchInputMarketingEventEmailSubscriber | The details of the contacts to subscribe to the event. Parameters of join and left time if state is Attended.
    externalAccountId := "externalAccountId_example" // string | The account id associated with the marketing event (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttendanceSubscriberStateChangesApi.AttendanceEmailCreateCreateByEmail(context.Background(), externalEventId, subscriberState).BatchInputMarketingEventEmailSubscriber(batchInputMarketingEventEmailSubscriber).ExternalAccountId(externalAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttendanceSubscriberStateChangesApi.AttendanceEmailCreateCreateByEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttendanceEmailCreateCreateByEmail`: BatchResponseSubscriberEmailResponse
    fmt.Fprintf(os.Stdout, "Response from `AttendanceSubscriberStateChangesApi.AttendanceEmailCreateCreateByEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event | 
**subscriberState** | **string** | The new subscriber state for the HubSpot contacts and the specified marketing event. For example: &#39;register&#39;, &#39;attend&#39; or &#39;cancel&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttendanceEmailCreateCreateByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputMarketingEventEmailSubscriber** | [**BatchInputMarketingEventEmailSubscriber**](BatchInputMarketingEventEmailSubscriber.md) | The details of the contacts to subscribe to the event. Parameters of join and left time if state is Attended. | 
 **externalAccountId** | **string** | The account id associated with the marketing event | 

### Return type

[**BatchResponseSubscriberEmailResponse**](BatchResponseSubscriberEmailResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

