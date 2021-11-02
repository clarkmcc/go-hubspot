# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Postmarketingv3marketingEventseventsexternalEventIdsubscriberStateemailUpsert**](SubscriberStateChangesApi.md#Postmarketingv3marketingEventseventsexternalEventIdsubscriberStateemailUpsert) | **Post** /marketing/v3/marketing-events/events/{externalEventId}/{subscriberState}/email-upsert | Record
[**Postmarketingv3marketingEventseventsexternalEventIdsubscriberStateupsert**](SubscriberStateChangesApi.md#Postmarketingv3marketingEventseventsexternalEventIdsubscriberStateupsert) | **Post** /marketing/v3/marketing-events/events/{externalEventId}/{subscriberState}/upsert | Record

# **Postmarketingv3marketingEventseventsexternalEventIdsubscriberStateemailUpsert**
> ModelError Postmarketingv3marketingEventseventsexternalEventIdsubscriberStateemailUpsert(ctx, body, externalEventId, subscriberState, externalAccountId)
Record

Record a subscription state between multiple HubSpot contacts and a marketing event, using contact email addresses.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputMarketingEventEmailSubscriber**](BatchInputMarketingEventEmailSubscriber.md)| The details of the contacts to subscribe to the event | 
  **externalEventId** | **string**| The id of the marketing event | 
  **subscriberState** | **string**| The new subscriber state for the HubSpot contacts and the specified marketing event | 
  **externalAccountId** | **string**| The account id associated with the marketing event | 

### Return type

[**ModelError**](Error.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postmarketingv3marketingEventseventsexternalEventIdsubscriberStateupsert**
> ModelError Postmarketingv3marketingEventseventsexternalEventIdsubscriberStateupsert(ctx, body, externalEventId, subscriberState, externalAccountId)
Record

Record a subscription state between multiple HubSpot contacts and a marketing event, using HubSpot contact ids.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputMarketingEventSubscriber**](BatchInputMarketingEventSubscriber.md)| The details of the contacts to subscribe to the event | 
  **externalEventId** | **string**| The id of the marketing event | 
  **subscriberState** | **string**| The new subscriber state for the HubSpot contacts and the specified marketing event | 
  **externalAccountId** | **string**| The account id associated with the marketing event | 

### Return type

[**ModelError**](Error.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

