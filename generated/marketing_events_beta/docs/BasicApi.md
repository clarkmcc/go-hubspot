# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletemarketingv3marketingEventseventsexternalEventId**](BasicApi.md#Deletemarketingv3marketingEventseventsexternalEventId) | **Delete** /marketing/v3/marketing-events/events/{externalEventId} | Delete a marketing event
[**Getmarketingv3marketingEventseventsexternalEventId**](BasicApi.md#Getmarketingv3marketingEventseventsexternalEventId) | **Get** /marketing/v3/marketing-events/events/{externalEventId} | Get a marketing event
[**Patchmarketingv3marketingEventseventsexternalEventId**](BasicApi.md#Patchmarketingv3marketingEventseventsexternalEventId) | **Patch** /marketing/v3/marketing-events/events/{externalEventId} | Update a marketing event
[**Postmarketingv3marketingEventsevents**](BasicApi.md#Postmarketingv3marketingEventsevents) | **Post** /marketing/v3/marketing-events/events | Create a marketing event
[**Postmarketingv3marketingEventseventsexternalEventIdcancel**](BasicApi.md#Postmarketingv3marketingEventseventsexternalEventIdcancel) | **Post** /marketing/v3/marketing-events/events/{externalEventId}/cancel | Mark a marketing event as cancelled
[**Putmarketingv3marketingEventseventsexternalEventId**](BasicApi.md#Putmarketingv3marketingEventseventsexternalEventId) | **Put** /marketing/v3/marketing-events/events/{externalEventId} | Create or update a marketing event

# **Deletemarketingv3marketingEventseventsexternalEventId**
> Deletemarketingv3marketingEventseventsexternalEventId(ctx, externalEventId, externalAccountId)
Delete a marketing event

Deletes an existing Marketing Event with the specified id, if one exists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **externalEventId** | **string**| The id of the marketing event to delete | 
  **externalAccountId** | **string**| The account id associated with the marketing event | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getmarketingv3marketingEventseventsexternalEventId**
> MarketingEventPublicReadResponse Getmarketingv3marketingEventseventsexternalEventId(ctx, externalEventId, externalAccountId)
Get a marketing event

Returns the details of the Marketing Event with the specified id, if one exists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **externalEventId** | **string**| The id of the marketing event to return | 
  **externalAccountId** | **string**| The account id associated with the marketing event | 

### Return type

[**MarketingEventPublicReadResponse**](MarketingEventPublicReadResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patchmarketingv3marketingEventseventsexternalEventId**
> MarketingEventPublicDefaultResponse Patchmarketingv3marketingEventseventsexternalEventId(ctx, body, externalEventId, externalAccountId)
Update a marketing event

Updates an existing Marketing Event with the specified id, if one exists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MarketingEventUpdateRequestParams**](MarketingEventUpdateRequestParams.md)| The details of the marketing event to update | 
  **externalEventId** | **string**| The id of the marketing event to update | 
  **externalAccountId** | **string**| The account id associated with the marketing event | 

### Return type

[**MarketingEventPublicDefaultResponse**](MarketingEventPublicDefaultResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postmarketingv3marketingEventsevents**
> MarketingEventDefaultResponse Postmarketingv3marketingEventsevents(ctx, body)
Create a marketing event

Creates a new marketing event in HubSpot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MarketingEventCreateRequestParams**](MarketingEventCreateRequestParams.md)| The details of the marketing event to create | 

### Return type

[**MarketingEventDefaultResponse**](MarketingEventDefaultResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postmarketingv3marketingEventseventsexternalEventIdcancel**
> MarketingEventDefaultResponse Postmarketingv3marketingEventseventsexternalEventIdcancel(ctx, externalEventId, externalAccountId)
Mark a marketing event as cancelled

Mark a marketing event as cancelled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **externalEventId** | **string**| The id of the marketing event to mark as cancelled | 
  **externalAccountId** | **string**| The account id associated with the marketing event | 

### Return type

[**MarketingEventDefaultResponse**](MarketingEventDefaultResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Putmarketingv3marketingEventseventsexternalEventId**
> MarketingEventPublicDefaultResponse Putmarketingv3marketingEventseventsexternalEventId(ctx, body, externalEventId)
Create or update a marketing event

Upsets a Marketing Event. If there is an existing Marketing event with the specified id, it will be updated; otherwise a new event will be created.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MarketingEventCreateRequestParams**](MarketingEventCreateRequestParams.md)| The details of the marketing event to upsert | 
  **externalEventId** | **string**| The id of the marketing event to upsert | 

### Return type

[**MarketingEventPublicDefaultResponse**](MarketingEventPublicDefaultResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

