# \EventsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIntegratorsTimelineV3EventsEventTemplateIdEventIdDetailGetDetailById**](EventsApi.md#GetIntegratorsTimelineV3EventsEventTemplateIdEventIdDetailGetDetailById) | **Get** /crm/v3/timeline/events/{eventTemplateId}/{eventId}/detail | Gets the detailTemplate as rendered
[**GetIntegratorsTimelineV3EventsEventTemplateIdEventIdGetById**](EventsApi.md#GetIntegratorsTimelineV3EventsEventTemplateIdEventIdGetById) | **Get** /crm/v3/timeline/events/{eventTemplateId}/{eventId} | Gets the event
[**GetIntegratorsTimelineV3EventsEventTemplateIdEventIdRenderGetRenderById**](EventsApi.md#GetIntegratorsTimelineV3EventsEventTemplateIdEventIdRenderGetRenderById) | **Get** /crm/v3/timeline/events/{eventTemplateId}/{eventId}/render | Renders the header or detail as HTML
[**PostIntegratorsTimelineV3EventsBatchCreateCreateBatch**](EventsApi.md#PostIntegratorsTimelineV3EventsBatchCreateCreateBatch) | **Post** /crm/v3/timeline/events/batch/create | Creates multiple events
[**PostIntegratorsTimelineV3EventsCreate**](EventsApi.md#PostIntegratorsTimelineV3EventsCreate) | **Post** /crm/v3/timeline/events | Create a single event



## GetIntegratorsTimelineV3EventsEventTemplateIdEventIdDetailGetDetailById

> EventDetail GetIntegratorsTimelineV3EventsEventTemplateIdEventIdDetailGetDetailById(ctx, eventTemplateId, eventId).Execute()

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
    resp, r, err := apiClient.EventsApi.GetIntegratorsTimelineV3EventsEventTemplateIdEventIdDetailGetDetailById(context.Background(), eventTemplateId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetIntegratorsTimelineV3EventsEventTemplateIdEventIdDetailGetDetailById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegratorsTimelineV3EventsEventTemplateIdEventIdDetailGetDetailById`: EventDetail
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetIntegratorsTimelineV3EventsEventTemplateIdEventIdDetailGetDetailById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**eventId** | **string** | The event ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegratorsTimelineV3EventsEventTemplateIdEventIdDetailGetDetailByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EventDetail**](EventDetail.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegratorsTimelineV3EventsEventTemplateIdEventIdGetById

> TimelineEventResponse GetIntegratorsTimelineV3EventsEventTemplateIdEventIdGetById(ctx, eventTemplateId, eventId).Execute()

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
    resp, r, err := apiClient.EventsApi.GetIntegratorsTimelineV3EventsEventTemplateIdEventIdGetById(context.Background(), eventTemplateId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetIntegratorsTimelineV3EventsEventTemplateIdEventIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegratorsTimelineV3EventsEventTemplateIdEventIdGetById`: TimelineEventResponse
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetIntegratorsTimelineV3EventsEventTemplateIdEventIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**eventId** | **string** | The event ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegratorsTimelineV3EventsEventTemplateIdEventIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TimelineEventResponse**](TimelineEventResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegratorsTimelineV3EventsEventTemplateIdEventIdRenderGetRenderById

> string GetIntegratorsTimelineV3EventsEventTemplateIdEventIdRenderGetRenderById(ctx, eventTemplateId, eventId).Detail(detail).Execute()

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
    resp, r, err := apiClient.EventsApi.GetIntegratorsTimelineV3EventsEventTemplateIdEventIdRenderGetRenderById(context.Background(), eventTemplateId, eventId).Detail(detail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetIntegratorsTimelineV3EventsEventTemplateIdEventIdRenderGetRenderById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegratorsTimelineV3EventsEventTemplateIdEventIdRenderGetRenderById`: string
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetIntegratorsTimelineV3EventsEventTemplateIdEventIdRenderGetRenderById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**eventId** | **string** | The event ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegratorsTimelineV3EventsEventTemplateIdEventIdRenderGetRenderByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **detail** | **bool** | Set to &#39;true&#39;, we want to render the &#x60;detailTemplate&#x60; instead of the &#x60;headerTemplate&#x60;. | 

### Return type

**string**

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegratorsTimelineV3EventsBatchCreateCreateBatch

> BatchResponseTimelineEventResponse PostIntegratorsTimelineV3EventsBatchCreateCreateBatch(ctx).BatchInputTimelineEvent(batchInputTimelineEvent).Execute()

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
    resp, r, err := apiClient.EventsApi.PostIntegratorsTimelineV3EventsBatchCreateCreateBatch(context.Background()).BatchInputTimelineEvent(batchInputTimelineEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.PostIntegratorsTimelineV3EventsBatchCreateCreateBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostIntegratorsTimelineV3EventsBatchCreateCreateBatch`: BatchResponseTimelineEventResponse
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.PostIntegratorsTimelineV3EventsBatchCreateCreateBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegratorsTimelineV3EventsBatchCreateCreateBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputTimelineEvent** | [**BatchInputTimelineEvent**](BatchInputTimelineEvent.md) | The timeline event definition. | 

### Return type

[**BatchResponseTimelineEventResponse**](BatchResponseTimelineEventResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegratorsTimelineV3EventsCreate

> TimelineEventResponse PostIntegratorsTimelineV3EventsCreate(ctx).TimelineEvent(timelineEvent).Execute()

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
    resp, r, err := apiClient.EventsApi.PostIntegratorsTimelineV3EventsCreate(context.Background()).TimelineEvent(timelineEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.PostIntegratorsTimelineV3EventsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostIntegratorsTimelineV3EventsCreate`: TimelineEventResponse
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.PostIntegratorsTimelineV3EventsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegratorsTimelineV3EventsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timelineEvent** | [**TimelineEvent**](TimelineEvent.md) | The timeline event definition. | 

### Return type

[**TimelineEventResponse**](TimelineEventResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

