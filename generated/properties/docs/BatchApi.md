# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Postcrmv3propertiesobjectTypebatcharchiveArchive**](BatchApi.md#Postcrmv3propertiesobjectTypebatcharchiveArchive) | **Post** /crm/v3/properties/{objectType}/batch/archive | Archive a batch of properties
[**Postcrmv3propertiesobjectTypebatchcreateCreate**](BatchApi.md#Postcrmv3propertiesobjectTypebatchcreateCreate) | **Post** /crm/v3/properties/{objectType}/batch/create | Create a batch of properties
[**Postcrmv3propertiesobjectTypebatchreadRead**](BatchApi.md#Postcrmv3propertiesobjectTypebatchreadRead) | **Post** /crm/v3/properties/{objectType}/batch/read | Read a batch of properties

# **Postcrmv3propertiesobjectTypebatcharchiveArchive**
> Postcrmv3propertiesobjectTypebatcharchiveArchive(ctx, body, objectType)
Archive a batch of properties

Archive a provided list of properties. This method will return a 204 No Content response on success regardless of the initial state of the property (e.g. active, already archived, non-existent).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputPropertyName**](BatchInputPropertyName.md)|  | 
  **objectType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3propertiesobjectTypebatchcreateCreate**
> BatchResponseProperty Postcrmv3propertiesobjectTypebatchcreateCreate(ctx, body, objectType)
Create a batch of properties

Create a batch of properties using the same rules as when creating an individual property.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputPropertyCreate**](BatchInputPropertyCreate.md)|  | 
  **objectType** | **string**|  | 

### Return type

[**BatchResponseProperty**](BatchResponseProperty.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3propertiesobjectTypebatchreadRead**
> BatchResponseProperty Postcrmv3propertiesobjectTypebatchreadRead(ctx, body, objectType)
Read a batch of properties

Read a provided list of properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchReadInputPropertyName**](BatchReadInputPropertyName.md)|  | 
  **objectType** | **string**|  | 

### Return type

[**BatchResponseProperty**](BatchResponseProperty.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

