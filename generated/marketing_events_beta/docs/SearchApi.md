# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getmarketingv3marketingEventseventssearch**](SearchApi.md#Getmarketingv3marketingEventseventssearch) | **Get** /marketing/v3/marketing-events/events/search | Search for marketing events

# **Getmarketingv3marketingEventseventssearch**
> CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging Getmarketingv3marketingEventseventssearch(ctx, q)
Search for marketing events

Search for marketing events that have an event id that starts with the query string

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **q** | **string**| The partial event id to search for | 

### Return type

[**CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging**](CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

