# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletewebhooksv3appIdsettingsClear**](SettingsApi.md#Deletewebhooksv3appIdsettingsClear) | **Delete** /webhooks/v3/{appId}/settings | Clear webhook settings
[**Getwebhooksv3appIdsettingsGetAll**](SettingsApi.md#Getwebhooksv3appIdsettingsGetAll) | **Get** /webhooks/v3/{appId}/settings | Get webhook settings
[**Putwebhooksv3appIdsettingsConfigure**](SettingsApi.md#Putwebhooksv3appIdsettingsConfigure) | **Put** /webhooks/v3/{appId}/settings | Configure webhook settings

# **Deletewebhooksv3appIdsettingsClear**
> Deletewebhooksv3appIdsettingsClear(ctx, appId)
Clear webhook settings

Resets webhook target URL to empty, and max concurrency limit to `0` for the given app. This will effectively pause all webhook subscriptions until new settings are provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| The ID of the target app. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getwebhooksv3appIdsettingsGetAll**
> SettingsResponse Getwebhooksv3appIdsettingsGetAll(ctx, appId)
Get webhook settings

Returns the current state of webhook settings for the given app. These settings include the app's configured target URL and max concurrency limit.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| The ID of the target app. | 

### Return type

[**SettingsResponse**](SettingsResponse.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Putwebhooksv3appIdsettingsConfigure**
> SettingsResponse Putwebhooksv3appIdsettingsConfigure(ctx, body, appId)
Configure webhook settings

Used to set the webhook target URL and max concurrency limit for the given app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SettingsChangeRequest**](SettingsChangeRequest.md)| Settings state to create new with or replace existing settings with. | 
  **appId** | **int32**| The ID of the target app. | 

### Return type

[**SettingsResponse**](SettingsResponse.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

