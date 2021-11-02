# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Postmarketingv3marketingEventseventsdelete**](BatchApi.md#Postmarketingv3marketingEventseventsdelete) | **Post** /marketing/v3/marketing-events/events/delete | Delete multiple marketing events
[**Postmarketingv3marketingEventseventsupsert**](BatchApi.md#Postmarketingv3marketingEventseventsupsert) | **Post** /marketing/v3/marketing-events/events/upsert | Create or update multiple marketing events

# **Postmarketingv3marketingEventseventsdelete**
> ModelError Postmarketingv3marketingEventseventsdelete(ctx, body)
Delete multiple marketing events

Bulk delete a number of marketing events in HubSpot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputMarketingEventExternalUniqueIdentifier**](BatchInputMarketingEventExternalUniqueIdentifier.md)| The details of the marketing events to delete | 

### Return type

[**ModelError**](Error.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postmarketingv3marketingEventseventsupsert**
> BatchResponseMarketingEventPublicDefaultResponse Postmarketingv3marketingEventseventsupsert(ctx, body)
Create or update multiple marketing events

Upset multiple Marketing Event. If there is an existing Marketing event with the specified id, it will be updated; otherwise a new event will be created.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputMarketingEventCreateRequestParams**](BatchInputMarketingEventCreateRequestParams.md)| The details of the marketing events to upsert | 

### Return type

[**BatchResponseMarketingEventPublicDefaultResponse**](BatchResponseMarketingEventPublicDefaultResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

