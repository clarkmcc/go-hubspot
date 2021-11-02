# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetcommunicationPreferencesv3statusemailemailAddressGetEmailStatus**](StatusApi.md#GetcommunicationPreferencesv3statusemailemailAddressGetEmailStatus) | **Get** /communication-preferences/v3/status/email/{emailAddress} | Get subscription statuses for a contact
[**PostcommunicationPreferencesv3subscribeSubscribe**](StatusApi.md#PostcommunicationPreferencesv3subscribeSubscribe) | **Post** /communication-preferences/v3/subscribe | Subscribe a contact
[**PostcommunicationPreferencesv3unsubscribeUnsubscribe**](StatusApi.md#PostcommunicationPreferencesv3unsubscribeUnsubscribe) | **Post** /communication-preferences/v3/unsubscribe | Unsubscribe a contact

# **GetcommunicationPreferencesv3statusemailemailAddressGetEmailStatus**
> PublicSubscriptionStatusesResponse GetcommunicationPreferencesv3statusemailemailAddressGetEmailStatus(ctx, emailAddress)
Get subscription statuses for a contact

Returns a list of subscriptions and their status for a given contact.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailAddress** | **string**|  | 

### Return type

[**PublicSubscriptionStatusesResponse**](PublicSubscriptionStatusesResponse.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostcommunicationPreferencesv3subscribeSubscribe**
> PublicSubscriptionStatus PostcommunicationPreferencesv3subscribeSubscribe(ctx, body)
Subscribe a contact

Subscribes a contact to the given subscription type. This API is not valid to use for subscribing a contact at a brand or portal level and will return an error.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PublicUpdateSubscriptionStatusRequest**](PublicUpdateSubscriptionStatusRequest.md)|  | 

### Return type

[**PublicSubscriptionStatus**](PublicSubscriptionStatus.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostcommunicationPreferencesv3unsubscribeUnsubscribe**
> PublicSubscriptionStatus PostcommunicationPreferencesv3unsubscribeUnsubscribe(ctx, body)
Unsubscribe a contact

Unsubscribes a contact from the given subscription type. This API is not valid to use for unsubscribing a contact at a brand or portal level and will return an error.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PublicUpdateSubscriptionStatusRequest**](PublicUpdateSubscriptionStatusRequest.md)|  | 

### Return type

[**PublicSubscriptionStatus**](PublicSubscriptionStatus.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

