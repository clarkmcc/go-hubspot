# \AttendanceSubscriberStateChangesApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttendanceCreate**](AttendanceSubscriberStateChangesApi.md#AttendanceCreate) | **Post** /marketing/v3/marketing-events/attendance/{externalEventId}/{subscriberState}/create | Record
[**AttendanceEmailCreate**](AttendanceSubscriberStateChangesApi.md#AttendanceEmailCreate) | **Post** /marketing/v3/marketing-events/attendance/{externalEventId}/{subscriberState}/email-create | Record



## AttendanceCreate

> BatchResponseSubscriberVidResponse AttendanceCreate(ctx, externalEventId, subscriberState).BatchInputMarketingEventSubscriber(batchInputMarketingEventSubscriber).ExternalAccountId(externalAccountId).Execute()

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
    subscriberState := "subscriberState_example" // string | The new subscriber state for the HubSpot contacts and the specified marketing event
    batchInputMarketingEventSubscriber := *openapiclient.NewBatchInputMarketingEventSubscriber([]openapiclient.MarketingEventSubscriber{*openapiclient.NewMarketingEventSubscriber(int64(123))}) // BatchInputMarketingEventSubscriber | The details of the contacts to subscribe to the event. Parameters of join and left time if state is Attended.
    externalAccountId := "externalAccountId_example" // string | The account id associated with the marketing event (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttendanceSubscriberStateChangesApi.AttendanceCreate(context.Background(), externalEventId, subscriberState).BatchInputMarketingEventSubscriber(batchInputMarketingEventSubscriber).ExternalAccountId(externalAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttendanceSubscriberStateChangesApi.AttendanceCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttendanceCreate`: BatchResponseSubscriberVidResponse
    fmt.Fprintf(os.Stdout, "Response from `AttendanceSubscriberStateChangesApi.AttendanceCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event | 
**subscriberState** | **string** | The new subscriber state for the HubSpot contacts and the specified marketing event | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttendanceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputMarketingEventSubscriber** | [**BatchInputMarketingEventSubscriber**](BatchInputMarketingEventSubscriber.md) | The details of the contacts to subscribe to the event. Parameters of join and left time if state is Attended. | 
 **externalAccountId** | **string** | The account id associated with the marketing event | 

### Return type

[**BatchResponseSubscriberVidResponse**](BatchResponseSubscriberVidResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttendanceEmailCreate

> BatchResponseSubscriberEmailResponse AttendanceEmailCreate(ctx, externalEventId, subscriberState).BatchInputMarketingEventEmailSubscriber(batchInputMarketingEventEmailSubscriber).ExternalAccountId(externalAccountId).Execute()

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
    subscriberState := "subscriberState_example" // string | The new subscriber state for the HubSpot contacts and the specified marketing event
    batchInputMarketingEventEmailSubscriber := *openapiclient.NewBatchInputMarketingEventEmailSubscriber([]openapiclient.MarketingEventEmailSubscriber{*openapiclient.NewMarketingEventEmailSubscriber(int64(123), "Email_example")}) // BatchInputMarketingEventEmailSubscriber | The details of the contacts to subscribe to the event. Parameters of join and left time if state is Attended.
    externalAccountId := "externalAccountId_example" // string | The account id associated with the marketing event (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttendanceSubscriberStateChangesApi.AttendanceEmailCreate(context.Background(), externalEventId, subscriberState).BatchInputMarketingEventEmailSubscriber(batchInputMarketingEventEmailSubscriber).ExternalAccountId(externalAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttendanceSubscriberStateChangesApi.AttendanceEmailCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttendanceEmailCreate`: BatchResponseSubscriberEmailResponse
    fmt.Fprintf(os.Stdout, "Response from `AttendanceSubscriberStateChangesApi.AttendanceEmailCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event | 
**subscriberState** | **string** | The new subscriber state for the HubSpot contacts and the specified marketing event | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttendanceEmailCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputMarketingEventEmailSubscriber** | [**BatchInputMarketingEventEmailSubscriber**](BatchInputMarketingEventEmailSubscriber.md) | The details of the contacts to subscribe to the event. Parameters of join and left time if state is Attended. | 
 **externalAccountId** | **string** | The account id associated with the marketing event | 

### Return type

[**BatchResponseSubscriberEmailResponse**](BatchResponseSubscriberEmailResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

