# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Postcrmv3objectsticketsbatcharchiveArchive**](BatchApi.md#Postcrmv3objectsticketsbatcharchiveArchive) | **Post** /crm/v3/objects/tickets/batch/archive | Archive a batch of tickets by ID
[**Postcrmv3objectsticketsbatchcreateCreate**](BatchApi.md#Postcrmv3objectsticketsbatchcreateCreate) | **Post** /crm/v3/objects/tickets/batch/create | Create a batch of tickets
[**Postcrmv3objectsticketsbatchreadRead**](BatchApi.md#Postcrmv3objectsticketsbatchreadRead) | **Post** /crm/v3/objects/tickets/batch/read | Read a batch of tickets by internal ID, or unique property values
[**Postcrmv3objectsticketsbatchupdateUpdate**](BatchApi.md#Postcrmv3objectsticketsbatchupdateUpdate) | **Post** /crm/v3/objects/tickets/batch/update | Update a batch of tickets

# **Postcrmv3objectsticketsbatcharchiveArchive**
> Postcrmv3objectsticketsbatcharchiveArchive(ctx, body)
Archive a batch of tickets by ID

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

# **Postcrmv3objectsticketsbatchcreateCreate**
> BatchResponseSimplePublicObject Postcrmv3objectsticketsbatchcreateCreate(ctx, body)
Create a batch of tickets

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

# **Postcrmv3objectsticketsbatchreadRead**
> BatchResponseSimplePublicObject Postcrmv3objectsticketsbatchreadRead(ctx, body, optional)
Read a batch of tickets by internal ID, or unique property values

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchReadInputSimplePublicObjectId**](BatchReadInputSimplePublicObjectId.md)|  | 
 **optional** | ***BatchApiPostcrmv3objectsticketsbatchreadReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BatchApiPostcrmv3objectsticketsbatchreadReadOpts struct
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

# **Postcrmv3objectsticketsbatchupdateUpdate**
> BatchResponseSimplePublicObject Postcrmv3objectsticketsbatchupdateUpdate(ctx, body)
Update a batch of tickets

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

