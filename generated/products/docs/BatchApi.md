# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Postcrmv3objectsproductsbatcharchiveArchive**](BatchApi.md#Postcrmv3objectsproductsbatcharchiveArchive) | **Post** /crm/v3/objects/products/batch/archive | Archive a batch of products by ID
[**Postcrmv3objectsproductsbatchcreateCreate**](BatchApi.md#Postcrmv3objectsproductsbatchcreateCreate) | **Post** /crm/v3/objects/products/batch/create | Create a batch of products
[**Postcrmv3objectsproductsbatchreadRead**](BatchApi.md#Postcrmv3objectsproductsbatchreadRead) | **Post** /crm/v3/objects/products/batch/read | Read a batch of products by internal ID, or unique property values
[**Postcrmv3objectsproductsbatchupdateUpdate**](BatchApi.md#Postcrmv3objectsproductsbatchupdateUpdate) | **Post** /crm/v3/objects/products/batch/update | Update a batch of products

# **Postcrmv3objectsproductsbatcharchiveArchive**
> Postcrmv3objectsproductsbatcharchiveArchive(ctx, body)
Archive a batch of products by ID

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

# **Postcrmv3objectsproductsbatchcreateCreate**
> BatchResponseSimplePublicObject Postcrmv3objectsproductsbatchcreateCreate(ctx, body)
Create a batch of products

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

# **Postcrmv3objectsproductsbatchreadRead**
> BatchResponseSimplePublicObject Postcrmv3objectsproductsbatchreadRead(ctx, body, optional)
Read a batch of products by internal ID, or unique property values

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchReadInputSimplePublicObjectId**](BatchReadInputSimplePublicObjectId.md)|  | 
 **optional** | ***BatchApiPostcrmv3objectsproductsbatchreadReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BatchApiPostcrmv3objectsproductsbatchreadReadOpts struct
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

# **Postcrmv3objectsproductsbatchupdateUpdate**
> BatchResponseSimplePublicObject Postcrmv3objectsproductsbatchupdateUpdate(ctx, body)
Update a batch of products

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

