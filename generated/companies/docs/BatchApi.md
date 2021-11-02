# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Postcrmv3objectscompaniesbatcharchiveArchive**](BatchApi.md#Postcrmv3objectscompaniesbatcharchiveArchive) | **Post** /crm/v3/objects/companies/batch/archive | Archive a batch of companies by ID
[**Postcrmv3objectscompaniesbatchcreateCreate**](BatchApi.md#Postcrmv3objectscompaniesbatchcreateCreate) | **Post** /crm/v3/objects/companies/batch/create | Create a batch of companies
[**Postcrmv3objectscompaniesbatchreadRead**](BatchApi.md#Postcrmv3objectscompaniesbatchreadRead) | **Post** /crm/v3/objects/companies/batch/read | Read a batch of companies by internal ID, or unique property values
[**Postcrmv3objectscompaniesbatchupdateUpdate**](BatchApi.md#Postcrmv3objectscompaniesbatchupdateUpdate) | **Post** /crm/v3/objects/companies/batch/update | Update a batch of companies

# **Postcrmv3objectscompaniesbatcharchiveArchive**
> Postcrmv3objectscompaniesbatcharchiveArchive(ctx, body)
Archive a batch of companies by ID

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

# **Postcrmv3objectscompaniesbatchcreateCreate**
> BatchResponseSimplePublicObject Postcrmv3objectscompaniesbatchcreateCreate(ctx, body)
Create a batch of companies

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

# **Postcrmv3objectscompaniesbatchreadRead**
> BatchResponseSimplePublicObject Postcrmv3objectscompaniesbatchreadRead(ctx, body, optional)
Read a batch of companies by internal ID, or unique property values

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchReadInputSimplePublicObjectId**](BatchReadInputSimplePublicObjectId.md)|  | 
 **optional** | ***BatchApiPostcrmv3objectscompaniesbatchreadReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BatchApiPostcrmv3objectscompaniesbatchreadReadOpts struct
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

# **Postcrmv3objectscompaniesbatchupdateUpdate**
> BatchResponseSimplePublicObject Postcrmv3objectscompaniesbatchupdateUpdate(ctx, body)
Update a batch of companies

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

