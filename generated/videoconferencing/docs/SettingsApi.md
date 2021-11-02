# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletecrmv3extensionsvideoconferencingsettingsappIdArchive**](SettingsApi.md#Deletecrmv3extensionsvideoconferencingsettingsappIdArchive) | **Delete** /crm/v3/extensions/videoconferencing/settings/{appId} | Delete settings
[**Getcrmv3extensionsvideoconferencingsettingsappIdGetById**](SettingsApi.md#Getcrmv3extensionsvideoconferencingsettingsappIdGetById) | **Get** /crm/v3/extensions/videoconferencing/settings/{appId} | Get settings
[**Putcrmv3extensionsvideoconferencingsettingsappIdReplace**](SettingsApi.md#Putcrmv3extensionsvideoconferencingsettingsappIdReplace) | **Put** /crm/v3/extensions/videoconferencing/settings/{appId} | Update settings

# **Deletecrmv3extensionsvideoconferencingsettingsappIdArchive**
> Deletecrmv3extensionsvideoconferencingsettingsappIdArchive(ctx, appId)
Delete settings

Deletes the settings for a video conference application with the specified ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| The ID of the video conference application. This is the identifier of the application created in your HubSpot developer portal. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcrmv3extensionsvideoconferencingsettingsappIdGetById**
> ExternalSettings Getcrmv3extensionsvideoconferencingsettingsappIdGetById(ctx, appId)
Get settings

Return the settings for a video conference application with the specified ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| The ID of the video conference application. This is the identifier of the application created in your HubSpot developer portal. | 

### Return type

[**ExternalSettings**](ExternalSettings.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Putcrmv3extensionsvideoconferencingsettingsappIdReplace**
> ExternalSettings Putcrmv3extensionsvideoconferencingsettingsappIdReplace(ctx, body, appId)
Update settings

Updates the settings for a video conference application with the specified ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExternalSettings**](ExternalSettings.md)|  | 
  **appId** | **int32**| The ID of the video conference application. This is the identifier of the application created in your HubSpot developer portal. | 

### Return type

[**ExternalSettings**](ExternalSettings.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

