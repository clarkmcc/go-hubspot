# \TemplatesApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3TimelineAppIdEventTemplatesEventTemplateIdArchive**](TemplatesApi.md#DeleteCrmV3TimelineAppIdEventTemplatesEventTemplateIdArchive) | **Delete** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId} | Deletes an event template for the app
[**GetCrmV3TimelineAppIdEventTemplatesEventTemplateIdGetById**](TemplatesApi.md#GetCrmV3TimelineAppIdEventTemplatesEventTemplateIdGetById) | **Get** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId} | Gets a specific event template for your app
[**GetCrmV3TimelineAppIdEventTemplatesGetAll**](TemplatesApi.md#GetCrmV3TimelineAppIdEventTemplatesGetAll) | **Get** /crm/v3/timeline/{appId}/event-templates | List all event templates for your app
[**PostCrmV3TimelineAppIdEventTemplatesCreate**](TemplatesApi.md#PostCrmV3TimelineAppIdEventTemplatesCreate) | **Post** /crm/v3/timeline/{appId}/event-templates | Create an event template for your app
[**PutCrmV3TimelineAppIdEventTemplatesEventTemplateIdUpdate**](TemplatesApi.md#PutCrmV3TimelineAppIdEventTemplatesEventTemplateIdUpdate) | **Put** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId} | Update an existing event template



## DeleteCrmV3TimelineAppIdEventTemplatesEventTemplateIdArchive

> DeleteCrmV3TimelineAppIdEventTemplatesEventTemplateIdArchive(ctx, eventTemplateId, appId).Execute()

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
    resp, r, err := apiClient.TemplatesApi.DeleteCrmV3TimelineAppIdEventTemplatesEventTemplateIdArchive(context.Background(), eventTemplateId, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.DeleteCrmV3TimelineAppIdEventTemplatesEventTemplateIdArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCrmV3TimelineAppIdEventTemplatesEventTemplateIdArchiveRequest struct via the builder pattern


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


## GetCrmV3TimelineAppIdEventTemplatesEventTemplateIdGetById

> TimelineEventTemplate GetCrmV3TimelineAppIdEventTemplatesEventTemplateIdGetById(ctx, eventTemplateId, appId).Execute()

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
    resp, r, err := apiClient.TemplatesApi.GetCrmV3TimelineAppIdEventTemplatesEventTemplateIdGetById(context.Background(), eventTemplateId, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.GetCrmV3TimelineAppIdEventTemplatesEventTemplateIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3TimelineAppIdEventTemplatesEventTemplateIdGetById`: TimelineEventTemplate
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.GetCrmV3TimelineAppIdEventTemplatesEventTemplateIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3TimelineAppIdEventTemplatesEventTemplateIdGetByIdRequest struct via the builder pattern


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


## GetCrmV3TimelineAppIdEventTemplatesGetAll

> CollectionResponseTimelineEventTemplateNoPaging GetCrmV3TimelineAppIdEventTemplatesGetAll(ctx, appId).Execute()

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
    resp, r, err := apiClient.TemplatesApi.GetCrmV3TimelineAppIdEventTemplatesGetAll(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.GetCrmV3TimelineAppIdEventTemplatesGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3TimelineAppIdEventTemplatesGetAll`: CollectionResponseTimelineEventTemplateNoPaging
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.GetCrmV3TimelineAppIdEventTemplatesGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3TimelineAppIdEventTemplatesGetAllRequest struct via the builder pattern


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


## PostCrmV3TimelineAppIdEventTemplatesCreate

> TimelineEventTemplate PostCrmV3TimelineAppIdEventTemplatesCreate(ctx, appId).TimelineEventTemplateCreateRequest(timelineEventTemplateCreateRequest).Execute()

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
    timelineEventTemplateCreateRequest := *openapiclient.NewTimelineEventTemplateCreateRequest("PetSpot Registration", []openapiclient.TimelineEventTemplateToken{*openapiclient.NewTimelineEventTemplateToken("petType", "Pet Type", "enumeration")}, "contacts") // TimelineEventTemplateCreateRequest | The new event template definition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.PostCrmV3TimelineAppIdEventTemplatesCreate(context.Background(), appId).TimelineEventTemplateCreateRequest(timelineEventTemplateCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.PostCrmV3TimelineAppIdEventTemplatesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3TimelineAppIdEventTemplatesCreate`: TimelineEventTemplate
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.PostCrmV3TimelineAppIdEventTemplatesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3TimelineAppIdEventTemplatesCreateRequest struct via the builder pattern


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


## PutCrmV3TimelineAppIdEventTemplatesEventTemplateIdUpdate

> TimelineEventTemplate PutCrmV3TimelineAppIdEventTemplatesEventTemplateIdUpdate(ctx, eventTemplateId, appId).TimelineEventTemplateUpdateRequest(timelineEventTemplateUpdateRequest).Execute()

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
    timelineEventTemplateUpdateRequest := *openapiclient.NewTimelineEventTemplateUpdateRequest("PetSpot Registration", []openapiclient.TimelineEventTemplateToken{*openapiclient.NewTimelineEventTemplateToken("petType", "Pet Type", "enumeration")}, "1001298") // TimelineEventTemplateUpdateRequest | The updated event template definition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.PutCrmV3TimelineAppIdEventTemplatesEventTemplateIdUpdate(context.Background(), eventTemplateId, appId).TimelineEventTemplateUpdateRequest(timelineEventTemplateUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.PutCrmV3TimelineAppIdEventTemplatesEventTemplateIdUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutCrmV3TimelineAppIdEventTemplatesEventTemplateIdUpdate`: TimelineEventTemplate
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.PutCrmV3TimelineAppIdEventTemplatesEventTemplateIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV3TimelineAppIdEventTemplatesEventTemplateIdUpdateRequest struct via the builder pattern


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

