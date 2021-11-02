# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getintegratorstimelinev3eventseventTemplateIdeventIdGetById**](EventsApi.md#Getintegratorstimelinev3eventseventTemplateIdeventIdGetById) | **Get** /crm/v3/timeline/events/{eventTemplateId}/{eventId} | Gets the event
[**Getintegratorstimelinev3eventseventTemplateIdeventIddetailGetDetailById**](EventsApi.md#Getintegratorstimelinev3eventseventTemplateIdeventIddetailGetDetailById) | **Get** /crm/v3/timeline/events/{eventTemplateId}/{eventId}/detail | Gets the detailTemplate as rendered
[**Getintegratorstimelinev3eventseventTemplateIdeventIdrenderGetRenderById**](EventsApi.md#Getintegratorstimelinev3eventseventTemplateIdeventIdrenderGetRenderById) | **Get** /crm/v3/timeline/events/{eventTemplateId}/{eventId}/render | Renders the header or detail as HTML
[**Postintegratorstimelinev3eventsCreate**](EventsApi.md#Postintegratorstimelinev3eventsCreate) | **Post** /crm/v3/timeline/events | Create a single event
[**Postintegratorstimelinev3eventsbatchcreateCreateBatch**](EventsApi.md#Postintegratorstimelinev3eventsbatchcreateCreateBatch) | **Post** /crm/v3/timeline/events/batch/create | Creates multiple events

# **Getintegratorstimelinev3eventseventTemplateIdeventIdGetById**
> TimelineEventResponse Getintegratorstimelinev3eventseventTemplateIdeventIdGetById(ctx, eventTemplateId, eventId)
Gets the event

This returns the previously created event. It contains all existing info for the event, but not necessarily the CRM object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventTemplateId** | **string**| The event template ID. | 
  **eventId** | **string**| The event ID. | 

### Return type

[**TimelineEventResponse**](TimelineEventResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getintegratorstimelinev3eventseventTemplateIdeventIddetailGetDetailById**
> EventDetail Getintegratorstimelinev3eventseventTemplateIdeventIddetailGetDetailById(ctx, eventTemplateId, eventId)
Gets the detailTemplate as rendered

This will take the `detailTemplate` from the event template and return an object rendering the specified event. If the template references `extraData` that isn't found in the event, it will be ignored and we'll render without it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventTemplateId** | **string**| The event template ID. | 
  **eventId** | **string**| The event ID. | 

### Return type

[**EventDetail**](EventDetail.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getintegratorstimelinev3eventseventTemplateIdeventIdrenderGetRenderById**
> string Getintegratorstimelinev3eventseventTemplateIdeventIdrenderGetRenderById(ctx, eventTemplateId, eventId, optional)
Renders the header or detail as HTML

This will take either the `headerTemplate` or `detailTemplate` from the event template and render for the specified event as HTML. If the template references `extraData` that isn't found in the event, it will be ignored and we'll render without it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventTemplateId** | **string**| The event template ID. | 
  **eventId** | **string**| The event ID. | 
 **optional** | ***EventsApiGetintegratorstimelinev3eventseventTemplateIdeventIdrenderGetRenderByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventsApiGetintegratorstimelinev3eventseventTemplateIdeventIdrenderGetRenderByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **detail** | **optional.Bool**| Set to &#x27;true&#x27;, we want to render the &#x60;detailTemplate&#x60; instead of the &#x60;headerTemplate&#x60;. | 

### Return type

**string**

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/html, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postintegratorstimelinev3eventsCreate**
> TimelineEventResponse Postintegratorstimelinev3eventsCreate(ctx, body)
Create a single event

Creates an instance of a timeline event based on an event template. Once created, this event is immutable on the object timeline and cannot be modified. If the event template was configured to update object properties via `objectPropertyName`, this call will also attempt to updates those properties, or add them if they don't exist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TimelineEvent**](TimelineEvent.md)| The timeline event definition. | 

### Return type

[**TimelineEventResponse**](TimelineEventResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postintegratorstimelinev3eventsbatchcreateCreateBatch**
> BatchResponseTimelineEventResponse Postintegratorstimelinev3eventsbatchcreateCreateBatch(ctx, body)
Creates multiple events

Creates multiple instances of timeline events based on an event template. Once created, these event are immutable on the object timeline and cannot be modified. If the event template was configured to update object properties via `objectPropertyName`, this call will also attempt to updates those properties, or add them if they don't exist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputTimelineEvent**](BatchInputTimelineEvent.md)| The timeline event definition. | 

### Return type

[**BatchResponseTimelineEventResponse**](BatchResponseTimelineEventResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

