# \TemplatesApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdArchive**](TemplatesApi.md#DeleteIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdArchive) | **Delete** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId} | Deletes an event template for the app
[**GetIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdGetById**](TemplatesApi.md#GetIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdGetById) | **Get** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId} | Gets a specific event template for your app
[**GetIntegratorsTimelineV3AppIdEventTemplatesGetAll**](TemplatesApi.md#GetIntegratorsTimelineV3AppIdEventTemplatesGetAll) | **Get** /crm/v3/timeline/{appId}/event-templates | List all event templates for your app
[**PostIntegratorsTimelineV3AppIdEventTemplatesCreate**](TemplatesApi.md#PostIntegratorsTimelineV3AppIdEventTemplatesCreate) | **Post** /crm/v3/timeline/{appId}/event-templates | Create an event template for your app
[**PutIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdUpdate**](TemplatesApi.md#PutIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdUpdate) | **Put** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId} | Update an existing event template



## DeleteIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdArchive

> DeleteIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdArchive(ctx, eventTemplateId, appId).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplatesApi.DeleteIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdArchive(context.Background(), eventTemplateId, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.DeleteIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdArchiveRequest struct via the builder pattern


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


## GetIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdGetById

> TimelineEventTemplate GetIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdGetById(ctx, eventTemplateId, appId).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplatesApi.GetIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdGetById(context.Background(), eventTemplateId, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.GetIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdGetById`: TimelineEventTemplate
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.GetIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdGetByIdRequest struct via the builder pattern


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


## GetIntegratorsTimelineV3AppIdEventTemplatesGetAll

> CollectionResponseTimelineEventTemplateNoPaging GetIntegratorsTimelineV3AppIdEventTemplatesGetAll(ctx, appId).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplatesApi.GetIntegratorsTimelineV3AppIdEventTemplatesGetAll(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.GetIntegratorsTimelineV3AppIdEventTemplatesGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegratorsTimelineV3AppIdEventTemplatesGetAll`: CollectionResponseTimelineEventTemplateNoPaging
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.GetIntegratorsTimelineV3AppIdEventTemplatesGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegratorsTimelineV3AppIdEventTemplatesGetAllRequest struct via the builder pattern


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


## PostIntegratorsTimelineV3AppIdEventTemplatesCreate

> TimelineEventTemplate PostIntegratorsTimelineV3AppIdEventTemplatesCreate(ctx, appId).TimelineEventTemplateCreateRequest(timelineEventTemplateCreateRequest).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplatesApi.PostIntegratorsTimelineV3AppIdEventTemplatesCreate(context.Background(), appId).TimelineEventTemplateCreateRequest(timelineEventTemplateCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.PostIntegratorsTimelineV3AppIdEventTemplatesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostIntegratorsTimelineV3AppIdEventTemplatesCreate`: TimelineEventTemplate
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.PostIntegratorsTimelineV3AppIdEventTemplatesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegratorsTimelineV3AppIdEventTemplatesCreateRequest struct via the builder pattern


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


## PutIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdUpdate

> TimelineEventTemplate PutIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdUpdate(ctx, eventTemplateId, appId).TimelineEventTemplateUpdateRequest(timelineEventTemplateUpdateRequest).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplatesApi.PutIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdUpdate(context.Background(), eventTemplateId, appId).TimelineEventTemplateUpdateRequest(timelineEventTemplateUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.PutIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdUpdate`: TimelineEventTemplate
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.PutIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdUpdateRequest struct via the builder pattern


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

