# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deleteintegratorstimelinev3appIdeventTemplateseventTemplateIdArchive**](TemplatesApi.md#Deleteintegratorstimelinev3appIdeventTemplateseventTemplateIdArchive) | **Delete** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId} | Deletes an event template for the app
[**Getintegratorstimelinev3appIdeventTemplatesGetAll**](TemplatesApi.md#Getintegratorstimelinev3appIdeventTemplatesGetAll) | **Get** /crm/v3/timeline/{appId}/event-templates | List all event templates for your app
[**Getintegratorstimelinev3appIdeventTemplateseventTemplateIdGetById**](TemplatesApi.md#Getintegratorstimelinev3appIdeventTemplateseventTemplateIdGetById) | **Get** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId} | Gets a specific event template for your app
[**Postintegratorstimelinev3appIdeventTemplatesCreate**](TemplatesApi.md#Postintegratorstimelinev3appIdeventTemplatesCreate) | **Post** /crm/v3/timeline/{appId}/event-templates | Create an event template for your app
[**Putintegratorstimelinev3appIdeventTemplateseventTemplateIdUpdate**](TemplatesApi.md#Putintegratorstimelinev3appIdeventTemplateseventTemplateIdUpdate) | **Put** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId} | Update an existing event template

# **Deleteintegratorstimelinev3appIdeventTemplateseventTemplateIdArchive**
> Deleteintegratorstimelinev3appIdeventTemplateseventTemplateIdArchive(ctx, eventTemplateId, appId)
Deletes an event template for the app

This will delete the event template. All associated events will be removed from search results and the timeline UI.  This action can't be undone, so it's highly recommended that you stop using any associated events before deleting a template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventTemplateId** | **string**| The event template ID. | 
  **appId** | **int32**| The ID of the target app. | 

### Return type

 (empty response body)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getintegratorstimelinev3appIdeventTemplatesGetAll**
> CollectionResponseTimelineEventTemplateNoPaging Getintegratorstimelinev3appIdeventTemplatesGetAll(ctx, appId)
List all event templates for your app

Use this to list all event templates owned by your app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| The ID of the target app. | 

### Return type

[**CollectionResponseTimelineEventTemplateNoPaging**](CollectionResponseTimelineEventTemplateNoPaging.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getintegratorstimelinev3appIdeventTemplateseventTemplateIdGetById**
> TimelineEventTemplate Getintegratorstimelinev3appIdeventTemplateseventTemplateIdGetById(ctx, eventTemplateId, appId)
Gets a specific event template for your app

View the current state of a specific template and its tokens.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventTemplateId** | **string**| The event template ID. | 
  **appId** | **int32**| The ID of the target app. | 

### Return type

[**TimelineEventTemplate**](TimelineEventTemplate.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postintegratorstimelinev3appIdeventTemplatesCreate**
> TimelineEventTemplate Postintegratorstimelinev3appIdeventTemplatesCreate(ctx, body, appId)
Create an event template for your app

Event templates define the general structure for a custom timeline event. This includes formatted copy for its heading and details, as well as any custom property definitions. The event could be something like viewing a video, registering for a webinar, or filling out a survey. A single app can define multiple event templates.  Event templates will be created for contacts by default, but they can be created for companies, tickets, and deals as well.  Each event template contains its own set of tokens and `Markdown` templates. These tokens can be associated with any CRM object properties via the `objectPropertyName` field to fully build out CRM objects.  You must create an event template before you can create events.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TimelineEventTemplateCreateRequest**](TimelineEventTemplateCreateRequest.md)| The new event template definition. | 
  **appId** | **int32**| The ID of the target app. | 

### Return type

[**TimelineEventTemplate**](TimelineEventTemplate.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Putintegratorstimelinev3appIdeventTemplateseventTemplateIdUpdate**
> TimelineEventTemplate Putintegratorstimelinev3appIdeventTemplateseventTemplateIdUpdate(ctx, body, eventTemplateId, appId)
Update an existing event template

Updates an existing template and its tokens. This is primarily used to update the headerTemplate/detailTemplate, and those changes will take effect for existing events.  You can also update or replace all the tokens in the template here instead of doing individual API calls on the `/tokens` endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TimelineEventTemplateUpdateRequest**](TimelineEventTemplateUpdateRequest.md)| The updated event template definition. | 
  **eventTemplateId** | **string**| The event template ID. | 
  **appId** | **int32**| The ID of the target app. | 

### Return type

[**TimelineEventTemplate**](TimelineEventTemplate.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

