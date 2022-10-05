# \TemplatesApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TemplateCreate**](TemplatesApi.md#TemplateCreate) | **Post** /crm/v3/timeline/{appId}/event-templates | Create an event template for your app
[**TemplateGetAll**](TemplatesApi.md#TemplateGetAll) | **Get** /crm/v3/timeline/{appId}/event-templates | List all event templates for your app
[**TemplatesArchive**](TemplatesApi.md#TemplatesArchive) | **Delete** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId} | Deletes an event template for the app
[**TemplatesGetByID**](TemplatesApi.md#TemplatesGetByID) | **Get** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId} | Gets a specific event template for your app
[**TemplatesUpdate**](TemplatesApi.md#TemplatesUpdate) | **Put** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId} | Update an existing event template



## TemplateCreate

> TimelineEventTemplate TemplateCreate(ctx, appId).TimelineEventTemplateCreateRequest(timelineEventTemplateCreateRequest).Execute()

Create an event template for your app



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
    timelineEventTemplateCreateRequest := *openapiclient.NewTimelineEventTemplateCreateRequest("Name_example", []openapiclient.TimelineEventTemplateToken{*openapiclient.NewTimelineEventTemplateToken("Label_example", []openapiclient.TimelineEventTemplateTokenOption{*openapiclient.NewTimelineEventTemplateTokenOption("Value_example", "Label_example")}, "Name_example", "Type_example")}, "ObjectType_example") // TimelineEventTemplateCreateRequest | The new event template definition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateCreate(context.Background(), appId).TimelineEventTemplateCreateRequest(timelineEventTemplateCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateCreate`: TimelineEventTemplate
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timelineEventTemplateCreateRequest** | [**TimelineEventTemplateCreateRequest**](TimelineEventTemplateCreateRequest.md) | The new event template definition. | 

### Return type

[**TimelineEventTemplate**](TimelineEventTemplate.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateGetAll

> CollectionResponseTimelineEventTemplateNoPaging TemplateGetAll(ctx, appId).Execute()

List all event templates for your app



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateGetAll(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateGetAll`: CollectionResponseTimelineEventTemplateNoPaging
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionResponseTimelineEventTemplateNoPaging**](CollectionResponseTimelineEventTemplateNoPaging.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesArchive

> TemplatesArchive(ctx, eventTemplateId, appId).Execute()

Deletes an event template for the app



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
    appId := int32(56) // int32 | The ID of the target app.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplatesArchive(context.Background(), eventTemplateId, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesArchiveRequest struct via the builder pattern


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


## TemplatesGetByID

> TimelineEventTemplate TemplatesGetByID(ctx, eventTemplateId, appId).Execute()

Gets a specific event template for your app



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
    appId := int32(56) // int32 | The ID of the target app.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplatesGetByID(context.Background(), eventTemplateId, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesGetByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesGetByID`: TimelineEventTemplate
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesGetByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesGetByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TimelineEventTemplate**](TimelineEventTemplate.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesUpdate

> TimelineEventTemplate TemplatesUpdate(ctx, eventTemplateId, appId).TimelineEventTemplateUpdateRequest(timelineEventTemplateUpdateRequest).Execute()

Update an existing event template



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
    appId := int32(56) // int32 | The ID of the target app.
    timelineEventTemplateUpdateRequest := *openapiclient.NewTimelineEventTemplateUpdateRequest("Name_example", []openapiclient.TimelineEventTemplateToken{*openapiclient.NewTimelineEventTemplateToken("Label_example", []openapiclient.TimelineEventTemplateTokenOption{*openapiclient.NewTimelineEventTemplateTokenOption("Value_example", "Label_example")}, "Name_example", "Type_example")}, "Id_example") // TimelineEventTemplateUpdateRequest | The updated event template definition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplatesUpdate(context.Background(), eventTemplateId, appId).TimelineEventTemplateUpdateRequest(timelineEventTemplateUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesUpdate`: TimelineEventTemplate
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timelineEventTemplateUpdateRequest** | [**TimelineEventTemplateUpdateRequest**](TimelineEventTemplateUpdateRequest.md) | The updated event template definition. | 

### Return type

[**TimelineEventTemplate**](TimelineEventTemplate.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

