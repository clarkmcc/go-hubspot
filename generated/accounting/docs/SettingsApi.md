# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getcrmv3extensionsaccountingsettingsappIdGetById**](SettingsApi.md#Getcrmv3extensionsaccountingsettingsappIdGetById) | **Get** /crm/v3/extensions/accounting/settings/{appId} | Get URL settings
[**Putcrmv3extensionsaccountingsettingsappIdReplace**](SettingsApi.md#Putcrmv3extensionsaccountingsettingsappIdReplace) | **Put** /crm/v3/extensions/accounting/settings/{appId} | Add/Update URL Settings

# **Getcrmv3extensionsaccountingsettingsappIdGetById**
> AccountingAppSettings Getcrmv3extensionsaccountingsettingsappIdGetById(ctx, appId)
Get URL settings

Returns the URL settings for an accounting app with the specified ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| The ID of the accounting app. This is the identifier of the application created in your HubSpot developer portal. | 

### Return type

[**AccountingAppSettings**](AccountingAppSettings.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Putcrmv3extensionsaccountingsettingsappIdReplace**
> Putcrmv3extensionsaccountingsettingsappIdReplace(ctx, body, appId)
Add/Update URL Settings

Add/Update the URL settings for an accounting app with the specified ID.  All URLs must use the `https` protocol.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountingAppSettings**](AccountingAppSettings.md)|  | 
  **appId** | **int32**| The ID of the accounting app. This is the identifier of the application created in your HubSpot developer portal. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

