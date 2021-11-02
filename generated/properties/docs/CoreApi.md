# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletecrmv3propertiesobjectTypepropertyNameArchive**](CoreApi.md#Deletecrmv3propertiesobjectTypepropertyNameArchive) | **Delete** /crm/v3/properties/{objectType}/{propertyName} | Archive a property
[**Getcrmv3propertiesobjectTypeGetAll**](CoreApi.md#Getcrmv3propertiesobjectTypeGetAll) | **Get** /crm/v3/properties/{objectType} | Read all properties
[**Getcrmv3propertiesobjectTypepropertyNameGetByName**](CoreApi.md#Getcrmv3propertiesobjectTypepropertyNameGetByName) | **Get** /crm/v3/properties/{objectType}/{propertyName} | Read a property
[**Patchcrmv3propertiesobjectTypepropertyNameUpdate**](CoreApi.md#Patchcrmv3propertiesobjectTypepropertyNameUpdate) | **Patch** /crm/v3/properties/{objectType}/{propertyName} | Update a property
[**Postcrmv3propertiesobjectTypeCreate**](CoreApi.md#Postcrmv3propertiesobjectTypeCreate) | **Post** /crm/v3/properties/{objectType} | Create a property

# **Deletecrmv3propertiesobjectTypepropertyNameArchive**
> Deletecrmv3propertiesobjectTypepropertyNameArchive(ctx, objectType, propertyName)
Archive a property

Move a property identified by {propertyName} to the recycling bin.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **propertyName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcrmv3propertiesobjectTypeGetAll**
> CollectionResponseProperty Getcrmv3propertiesobjectTypeGetAll(ctx, objectType, optional)
Read all properties

Read all existing properties for the specified object type and HubSpot account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
 **optional** | ***CoreApiGetcrmv3propertiesobjectTypeGetAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CoreApiGetcrmv3propertiesobjectTypeGetAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **optional.Bool**| Whether to return only results that have been archived. | [default to false]

### Return type

[**CollectionResponseProperty**](CollectionResponseProperty.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcrmv3propertiesobjectTypepropertyNameGetByName**
> Property Getcrmv3propertiesobjectTypepropertyNameGetByName(ctx, objectType, propertyName, optional)
Read a property

Read a property identified by {propertyName}.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **propertyName** | **string**|  | 
 **optional** | ***CoreApiGetcrmv3propertiesobjectTypepropertyNameGetByNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CoreApiGetcrmv3propertiesobjectTypepropertyNameGetByNameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **archived** | **optional.Bool**| Whether to return only results that have been archived. | [default to false]

### Return type

[**Property**](Property.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patchcrmv3propertiesobjectTypepropertyNameUpdate**
> Property Patchcrmv3propertiesobjectTypepropertyNameUpdate(ctx, body, objectType, propertyName)
Update a property

Perform a partial update of a property identified by {propertyName}. Provided fields will be overwritten.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PropertyUpdate**](PropertyUpdate.md)|  | 
  **objectType** | **string**|  | 
  **propertyName** | **string**|  | 

### Return type

[**Property**](Property.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3propertiesobjectTypeCreate**
> Property Postcrmv3propertiesobjectTypeCreate(ctx, body, objectType)
Create a property

Create and return a copy of a new property for the specified object type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PropertyCreate**](PropertyCreate.md)|  | 
  **objectType** | **string**|  | 

### Return type

[**Property**](Property.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

