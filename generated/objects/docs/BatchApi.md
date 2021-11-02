# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Postcrmv3objectsobjectTypebatcharchiveArchive**](BatchApi.md#Postcrmv3objectsobjectTypebatcharchiveArchive) | **Post** /crm/v3/objects/{objectType}/batch/archive | Archive a batch of objects by ID
[**Postcrmv3objectsobjectTypebatchcreateCreate**](BatchApi.md#Postcrmv3objectsobjectTypebatchcreateCreate) | **Post** /crm/v3/objects/{objectType}/batch/create | Create a batch of objects
[**Postcrmv3objectsobjectTypebatchreadRead**](BatchApi.md#Postcrmv3objectsobjectTypebatchreadRead) | **Post** /crm/v3/objects/{objectType}/batch/read | Read a batch of objects by internal ID, or unique property values
[**Postcrmv3objectsobjectTypebatchupdateUpdate**](BatchApi.md#Postcrmv3objectsobjectTypebatchupdateUpdate) | **Post** /crm/v3/objects/{objectType}/batch/update | Update a batch of objects

# **Postcrmv3objectsobjectTypebatcharchiveArchive**
> Postcrmv3objectsobjectTypebatcharchiveArchive(ctx, body, objectType)
Archive a batch of objects by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputSimplePublicObjectId**](BatchInputSimplePublicObjectId.md)|  | 
  **objectType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3objectsobjectTypebatchcreateCreate**
> BatchResponseSimplePublicObject Postcrmv3objectsobjectTypebatchcreateCreate(ctx, body, objectType)
Create a batch of objects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputSimplePublicObjectInput**](BatchInputSimplePublicObjectInput.md)|  | 
  **objectType** | **string**|  | 

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3objectsobjectTypebatchreadRead**
> BatchResponseSimplePublicObject Postcrmv3objectsobjectTypebatchreadRead(ctx, body, objectType, optional)
Read a batch of objects by internal ID, or unique property values

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchReadInputSimplePublicObjectId**](BatchReadInputSimplePublicObjectId.md)|  | 
  **objectType** | **string**|  | 
 **optional** | ***BatchApiPostcrmv3objectsobjectTypebatchreadReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BatchApiPostcrmv3objectsobjectTypebatchreadReadOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **archived** | **optional.**| Whether to return only results that have been archived. | [default to false]

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3objectsobjectTypebatchupdateUpdate**
> BatchResponseSimplePublicObject Postcrmv3objectsobjectTypebatchupdateUpdate(ctx, body, objectType)
Update a batch of objects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputSimplePublicObjectBatchInput**](BatchInputSimplePublicObjectBatchInput.md)|  | 
  **objectType** | **string**|  | 

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

