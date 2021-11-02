# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Postcrmv3objectslineItemsbatcharchiveArchive**](BatchApi.md#Postcrmv3objectslineItemsbatcharchiveArchive) | **Post** /crm/v3/objects/line_items/batch/archive | Archive a batch of line items by ID
[**Postcrmv3objectslineItemsbatchcreateCreate**](BatchApi.md#Postcrmv3objectslineItemsbatchcreateCreate) | **Post** /crm/v3/objects/line_items/batch/create | Create a batch of line items
[**Postcrmv3objectslineItemsbatchreadRead**](BatchApi.md#Postcrmv3objectslineItemsbatchreadRead) | **Post** /crm/v3/objects/line_items/batch/read | Read a batch of line items by internal ID, or unique property values
[**Postcrmv3objectslineItemsbatchupdateUpdate**](BatchApi.md#Postcrmv3objectslineItemsbatchupdateUpdate) | **Post** /crm/v3/objects/line_items/batch/update | Update a batch of line items

# **Postcrmv3objectslineItemsbatcharchiveArchive**
> Postcrmv3objectslineItemsbatcharchiveArchive(ctx, body)
Archive a batch of line items by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputSimplePublicObjectId**](BatchInputSimplePublicObjectId.md)|  | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3objectslineItemsbatchcreateCreate**
> BatchResponseSimplePublicObject Postcrmv3objectslineItemsbatchcreateCreate(ctx, body)
Create a batch of line items

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputSimplePublicObjectInput**](BatchInputSimplePublicObjectInput.md)|  | 

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3objectslineItemsbatchreadRead**
> BatchResponseSimplePublicObject Postcrmv3objectslineItemsbatchreadRead(ctx, body, optional)
Read a batch of line items by internal ID, or unique property values

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchReadInputSimplePublicObjectId**](BatchReadInputSimplePublicObjectId.md)|  | 
 **optional** | ***BatchApiPostcrmv3objectslineItemsbatchreadReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BatchApiPostcrmv3objectslineItemsbatchreadReadOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **optional.**| Whether to return only results that have been archived. | [default to false]

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3objectslineItemsbatchupdateUpdate**
> BatchResponseSimplePublicObject Postcrmv3objectslineItemsbatchupdateUpdate(ctx, body)
Update a batch of line items

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputSimplePublicObjectBatchInput**](BatchInputSimplePublicObjectBatchInput.md)|  | 

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

