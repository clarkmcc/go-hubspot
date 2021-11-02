# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletecrmv3extensionscallingappIdsettingsArchive**](SettingsApi.md#Deletecrmv3extensionscallingappIdsettingsArchive) | **Delete** /crm/v3/extensions/calling/{appId}/settings | Delete calling settings
[**Getcrmv3extensionscallingappIdsettingsGetById**](SettingsApi.md#Getcrmv3extensionscallingappIdsettingsGetById) | **Get** /crm/v3/extensions/calling/{appId}/settings | Get calling settings
[**Patchcrmv3extensionscallingappIdsettingsUpdate**](SettingsApi.md#Patchcrmv3extensionscallingappIdsettingsUpdate) | **Patch** /crm/v3/extensions/calling/{appId}/settings | Update settings
[**Postcrmv3extensionscallingappIdsettingsCreate**](SettingsApi.md#Postcrmv3extensionscallingappIdsettingsCreate) | **Post** /crm/v3/extensions/calling/{appId}/settings | Configure a calling extension

# **Deletecrmv3extensionscallingappIdsettingsArchive**
> Deletecrmv3extensionscallingappIdsettingsArchive(ctx, appId)
Delete calling settings

Deletes this calling extension. This will remove your service as an option for all connected accounts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| The ID of the target app. | 

### Return type

 (empty response body)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcrmv3extensionscallingappIdsettingsGetById**
> SettingsResponse Getcrmv3extensionscallingappIdsettingsGetById(ctx, appId)
Get calling settings

Returns the calling extension settings configured for your app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| The ID of the target app. | 

### Return type

[**SettingsResponse**](SettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patchcrmv3extensionscallingappIdsettingsUpdate**
> SettingsResponse Patchcrmv3extensionscallingappIdsettingsUpdate(ctx, body, appId)
Update settings

Updates existing calling extension settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SettingsPatchRequest**](SettingsPatchRequest.md)| Updated details for the settings. | 
  **appId** | **int32**| The ID of the target app. | 

### Return type

[**SettingsResponse**](SettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3extensionscallingappIdsettingsCreate**
> SettingsResponse Postcrmv3extensionscallingappIdsettingsCreate(ctx, body, appId)
Configure a calling extension

Used to set the menu label, target iframe URL, and dimensions for your calling extension.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SettingsRequest**](SettingsRequest.md)| Settings state to create with. | 
  **appId** | **int32**| The ID of the target app. | 

### Return type

[**SettingsResponse**](SettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

