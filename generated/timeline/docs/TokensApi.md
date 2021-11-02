# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deleteintegratorstimelinev3appIdeventTemplateseventTemplateIdtokenstokenNameArchive**](TokensApi.md#Deleteintegratorstimelinev3appIdeventTemplateseventTemplateIdtokenstokenNameArchive) | **Delete** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId}/tokens/{tokenName} | Removes a token from the event template
[**Postintegratorstimelinev3appIdeventTemplateseventTemplateIdtokensCreate**](TokensApi.md#Postintegratorstimelinev3appIdeventTemplateseventTemplateIdtokensCreate) | **Post** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId}/tokens | Adds a token to an existing event template
[**Putintegratorstimelinev3appIdeventTemplateseventTemplateIdtokenstokenNameUpdate**](TokensApi.md#Putintegratorstimelinev3appIdeventTemplateseventTemplateIdtokenstokenNameUpdate) | **Put** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId}/tokens/{tokenName} | Updates an existing token on an event template

# **Deleteintegratorstimelinev3appIdeventTemplateseventTemplateIdtokenstokenNameArchive**
> Deleteintegratorstimelinev3appIdeventTemplateseventTemplateIdtokenstokenNameArchive(ctx, eventTemplateId, tokenName, appId)
Removes a token from the event template

This will remove the token from an existing template. Existing events and CRM objects will still retain the token and its mapped object properties, but new ones will not.  The timeline will still display this property for older CRM objects if it's still referenced in the template `Markdown`. New events will not.  Any lists or reports referencing deleted tokens will no longer return new contacts, but old ones will still exist in the lists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventTemplateId** | **string**| The event template ID. | 
  **tokenName** | **string**| The token name. | 
  **appId** | **int32**| The ID of the target app. | 

### Return type

 (empty response body)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postintegratorstimelinev3appIdeventTemplateseventTemplateIdtokensCreate**
> TimelineEventTemplateToken Postintegratorstimelinev3appIdeventTemplateseventTemplateIdtokensCreate(ctx, body, eventTemplateId, appId)
Adds a token to an existing event template

Once you've defined an event template, it's likely that you'll want to define tokens for it as well. You can do this on the event template itself or update individual tokens here.  Event type tokens allow you to attach custom data to events displayed in a timeline or used for list segmentation.  You can also use `objectPropertyName` to associate any CRM object properties. This will allow you to fully build out CRM objects.  Token names should be unique across the template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TimelineEventTemplateToken**](TimelineEventTemplateToken.md)| The new token definition. | 
  **eventTemplateId** | **string**| The event template ID. | 
  **appId** | **int32**| The ID of the target app. | 

### Return type

[**TimelineEventTemplateToken**](TimelineEventTemplateToken.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Putintegratorstimelinev3appIdeventTemplateseventTemplateIdtokenstokenNameUpdate**
> TimelineEventTemplateToken Putintegratorstimelinev3appIdeventTemplateseventTemplateIdtokenstokenNameUpdate(ctx, body, eventTemplateId, tokenName, appId)
Updates an existing token on an event template

This will update the existing token on an event template. Name and type can't be changed on existing tokens.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TimelineEventTemplateTokenUpdateRequest**](TimelineEventTemplateTokenUpdateRequest.md)| The updated token definition. | 
  **eventTemplateId** | **string**| The event template ID. | 
  **tokenName** | **string**| The token name. | 
  **appId** | **int32**| The ID of the target app. | 

### Return type

[**TimelineEventTemplateToken**](TimelineEventTemplateToken.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

