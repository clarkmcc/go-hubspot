# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletecrmv3propertiesobjectTypegroupsgroupNameArchive**](GroupsApi.md#Deletecrmv3propertiesobjectTypegroupsgroupNameArchive) | **Delete** /crm/v3/properties/{objectType}/groups/{groupName} | Archive a property group
[**Getcrmv3propertiesobjectTypegroupsGetAll**](GroupsApi.md#Getcrmv3propertiesobjectTypegroupsGetAll) | **Get** /crm/v3/properties/{objectType}/groups | Read all property groups
[**Getcrmv3propertiesobjectTypegroupsgroupNameGetByName**](GroupsApi.md#Getcrmv3propertiesobjectTypegroupsgroupNameGetByName) | **Get** /crm/v3/properties/{objectType}/groups/{groupName} | Read a property group
[**Patchcrmv3propertiesobjectTypegroupsgroupNameUpdate**](GroupsApi.md#Patchcrmv3propertiesobjectTypegroupsgroupNameUpdate) | **Patch** /crm/v3/properties/{objectType}/groups/{groupName} | Update a property group
[**Postcrmv3propertiesobjectTypegroupsCreate**](GroupsApi.md#Postcrmv3propertiesobjectTypegroupsCreate) | **Post** /crm/v3/properties/{objectType}/groups | Create a property group

# **Deletecrmv3propertiesobjectTypegroupsgroupNameArchive**
> Deletecrmv3propertiesobjectTypegroupsgroupNameArchive(ctx, objectType, groupName)
Archive a property group

Move a property group identified by {groupName} to the recycling bin.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **groupName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcrmv3propertiesobjectTypegroupsGetAll**
> CollectionResponsePropertyGroup Getcrmv3propertiesobjectTypegroupsGetAll(ctx, objectType)
Read all property groups

Read all existing property groups for the specified object type and HubSpot account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 

### Return type

[**CollectionResponsePropertyGroup**](CollectionResponsePropertyGroup.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcrmv3propertiesobjectTypegroupsgroupNameGetByName**
> PropertyGroup Getcrmv3propertiesobjectTypegroupsgroupNameGetByName(ctx, objectType, groupName)
Read a property group

Read a property group identified by {groupName}.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **groupName** | **string**|  | 

### Return type

[**PropertyGroup**](PropertyGroup.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patchcrmv3propertiesobjectTypegroupsgroupNameUpdate**
> PropertyGroup Patchcrmv3propertiesobjectTypegroupsgroupNameUpdate(ctx, body, objectType, groupName)
Update a property group

Perform a partial update of a property group identified by {groupName}. Provided fields will be overwritten.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PropertyGroupUpdate**](PropertyGroupUpdate.md)|  | 
  **objectType** | **string**|  | 
  **groupName** | **string**|  | 

### Return type

[**PropertyGroup**](PropertyGroup.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3propertiesobjectTypegroupsCreate**
> PropertyGroup Postcrmv3propertiesobjectTypegroupsCreate(ctx, body, objectType)
Create a property group

Create and return a copy of a new property group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PropertyGroupCreate**](PropertyGroupCreate.md)|  | 
  **objectType** | **string**|  | 

### Return type

[**PropertyGroup**](PropertyGroup.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

