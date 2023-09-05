# \EventsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchCreate**](EventsApi.md#BatchCreate) | **Post** /crm/v3/timeline/events/batch/create | Creates multiple events
[**Create**](EventsApi.md#Create) | **Post** /crm/v3/timeline/events | Create a single event
[**GetByID**](EventsApi.md#GetByID) | **Get** /crm/v3/timeline/events/{eventTemplateId}/{eventId} | Gets the event
[**GetDetailByID**](EventsApi.md#GetDetailByID) | **Get** /crm/v3/timeline/events/{eventTemplateId}/{eventId}/detail | Gets the detailTemplate as rendered
[**GetRenderByID**](EventsApi.md#GetRenderByID) | **Get** /crm/v3/timeline/events/{eventTemplateId}/{eventId}/render | Renders the header or detail as HTML



## BatchCreate

> BatchCreate(ctx).BatchInputTimelineEvent(batchInputTimelineEvent).Execute()

Creates multiple events



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
    batchInputTimelineEvent := *openapiclient.NewBatchInputTimelineEvent([]openapiclient.TimelineEvent{*openapiclient.NewTimelineEvent("EventTemplateId_example", map[string]string{"key": "Inner_example"})}) // BatchInputTimelineEvent | The timeline event definition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.BatchCreate(context.Background()).BatchInputTimelineEvent(batchInputTimelineEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.BatchCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputTimelineEvent** | [**BatchInputTimelineEvent**](BatchInputTimelineEvent.md) | The timeline event definition. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> TimelineEventResponse Create(ctx).TimelineEvent(timelineEvent).Execute()

Create a single event



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
    timelineEvent := *openapiclient.NewTimelineEvent("EventTemplateId_example", map[string]string{"key": "Inner_example"}) // TimelineEvent | The timeline event definition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.Create(context.Background()).TimelineEvent(timelineEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: TimelineEventResponse
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timelineEvent** | [**TimelineEvent**](TimelineEvent.md) | The timeline event definition. | 

### Return type

[**TimelineEventResponse**](TimelineEventResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByID

> TimelineEventResponse GetByID(ctx, eventTemplateId, eventId).Execute()

Gets the event



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
    eventTemplateId := "eventTemplateId_example" // string | The event template ID.
    eventId := "eventId_example" // string | The event ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetByID(context.Background(), eventTemplateId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetByID`: TimelineEventResponse
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**eventId** | **string** | The event ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TimelineEventResponse**](TimelineEventResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDetailByID

> EventDetail GetDetailByID(ctx, eventTemplateId, eventId).Execute()

Gets the detailTemplate as rendered



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
    eventTemplateId := "eventTemplateId_example" // string | The event template ID.
    eventId := "eventId_example" // string | The event ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetDetailByID(context.Background(), eventTemplateId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetDetailByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDetailByID`: EventDetail
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetDetailByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**eventId** | **string** | The event ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDetailByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EventDetail**](EventDetail.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRenderByID

> string GetRenderByID(ctx, eventTemplateId, eventId).Detail(detail).Execute()

Renders the header or detail as HTML



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
    eventTemplateId := "eventTemplateId_example" // string | The event template ID.
    eventId := "eventId_example" // string | The event ID.
    detail := true // bool | Set to 'true', we want to render the `detailTemplate` instead of the `headerTemplate`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetRenderByID(context.Background(), eventTemplateId, eventId).Detail(detail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetRenderByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRenderByID`: string
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetRenderByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**eventId** | **string** | The event ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRenderByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **detail** | **bool** | Set to &#39;true&#39;, we want to render the &#x60;detailTemplate&#x60; instead of the &#x60;headerTemplate&#x60;. | 

### Return type

**string**

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

