# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getmarketingv3marketingEventsappIdsettings**](SettingsApi.md#Getmarketingv3marketingEventsappIdsettings) | **Get** /marketing/v3/marketing-events/{appId}/settings | Retrieve the application settings
[**Postmarketingv3marketingEventsappIdsettings**](SettingsApi.md#Postmarketingv3marketingEventsappIdsettings) | **Post** /marketing/v3/marketing-events/{appId}/settings | Update the application settings

# **Getmarketingv3marketingEventsappIdsettings**
> EventDetailSettings Getmarketingv3marketingEventsappIdsettings(ctx, appId)
Retrieve the application settings

Retrieve the current settings for the application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| The id of the application to retrieve the settings for. | 

### Return type

[**EventDetailSettings**](EventDetailSettings.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey), [hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postmarketingv3marketingEventsappIdsettings**
> EventDetailSettings Postmarketingv3marketingEventsappIdsettings(ctx, body, appId)
Update the application settings

Create or update the current settings for the application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventDetailSettingsUrl**](EventDetailSettingsUrl.md)| The new application settings | 
  **appId** | **int32**| The id of the application to update the settings for. | 

### Return type

[**EventDetailSettings**](EventDetailSettings.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey), [hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

