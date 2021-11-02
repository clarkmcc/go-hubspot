# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Postcrmv3objectsdealsbatcharchiveArchive**](BatchApi.md#Postcrmv3objectsdealsbatcharchiveArchive) | **Post** /crm/v3/objects/deals/batch/archive | Archive a batch of deals by ID
[**Postcrmv3objectsdealsbatchcreateCreate**](BatchApi.md#Postcrmv3objectsdealsbatchcreateCreate) | **Post** /crm/v3/objects/deals/batch/create | Create a batch of deals
[**Postcrmv3objectsdealsbatchreadRead**](BatchApi.md#Postcrmv3objectsdealsbatchreadRead) | **Post** /crm/v3/objects/deals/batch/read | Read a batch of deals by internal ID, or unique property values
[**Postcrmv3objectsdealsbatchupdateUpdate**](BatchApi.md#Postcrmv3objectsdealsbatchupdateUpdate) | **Post** /crm/v3/objects/deals/batch/update | Update a batch of deals

# **Postcrmv3objectsdealsbatcharchiveArchive**
> Postcrmv3objectsdealsbatcharchiveArchive(ctx, body)
Archive a batch of deals by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputSimplePublicObjectId**](BatchInputSimplePublicObjectId.md)|  | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3objectsdealsbatchcreateCreate**
> BatchResponseSimplePublicObject Postcrmv3objectsdealsbatchcreateCreate(ctx, body)
Create a batch of deals

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputSimplePublicObjectInput**](BatchInputSimplePublicObjectInput.md)|  | 

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3objectsdealsbatchreadRead**
> BatchResponseSimplePublicObject Postcrmv3objectsdealsbatchreadRead(ctx, body, optional)
Read a batch of deals by internal ID, or unique property values

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchReadInputSimplePublicObjectId**](BatchReadInputSimplePublicObjectId.md)|  | 
 **optional** | ***BatchApiPostcrmv3objectsdealsbatchreadReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BatchApiPostcrmv3objectsdealsbatchreadReadOpts struct
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

# **Postcrmv3objectsdealsbatchupdateUpdate**
> BatchResponseSimplePublicObject Postcrmv3objectsdealsbatchupdateUpdate(ctx, body)
Update a batch of deals

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputSimplePublicObjectBatchInput**](BatchInputSimplePublicObjectBatchInput.md)|  | 

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

