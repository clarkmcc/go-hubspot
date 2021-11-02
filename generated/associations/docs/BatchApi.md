# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Postcrmv3associationsfromObjectTypetoObjectTypebatcharchiveArchive**](BatchApi.md#Postcrmv3associationsfromObjectTypetoObjectTypebatcharchiveArchive) | **Post** /crm/v3/associations/{fromObjectType}/{toObjectType}/batch/archive | Archive a batch of associations
[**Postcrmv3associationsfromObjectTypetoObjectTypebatchcreateCreate**](BatchApi.md#Postcrmv3associationsfromObjectTypetoObjectTypebatchcreateCreate) | **Post** /crm/v3/associations/{fromObjectType}/{toObjectType}/batch/create | Create a batch of associations
[**Postcrmv3associationsfromObjectTypetoObjectTypebatchreadRead**](BatchApi.md#Postcrmv3associationsfromObjectTypetoObjectTypebatchreadRead) | **Post** /crm/v3/associations/{fromObjectType}/{toObjectType}/batch/read | Read a batch of associations

# **Postcrmv3associationsfromObjectTypetoObjectTypebatcharchiveArchive**
> Postcrmv3associationsfromObjectTypetoObjectTypebatcharchiveArchive(ctx, fromObjectType, toObjectType, optional)
Archive a batch of associations

Remove the associations between all pairs of objects identified in the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fromObjectType** | **string**|  | 
  **toObjectType** | **string**|  | 
 **optional** | ***BatchApiPostcrmv3associationsfromObjectTypetoObjectTypebatcharchiveArchiveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BatchApiPostcrmv3associationsfromObjectTypetoObjectTypebatcharchiveArchiveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of BatchInputPublicAssociation**](BatchInputPublicAssociation.md)|  | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3associationsfromObjectTypetoObjectTypebatchcreateCreate**
> BatchResponsePublicAssociation Postcrmv3associationsfromObjectTypetoObjectTypebatchcreateCreate(ctx, fromObjectType, toObjectType, optional)
Create a batch of associations

Associate all pairs of objects identified in the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fromObjectType** | **string**|  | 
  **toObjectType** | **string**|  | 
 **optional** | ***BatchApiPostcrmv3associationsfromObjectTypetoObjectTypebatchcreateCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BatchApiPostcrmv3associationsfromObjectTypetoObjectTypebatchcreateCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of BatchInputPublicAssociation**](BatchInputPublicAssociation.md)|  | 

### Return type

[**BatchResponsePublicAssociation**](BatchResponsePublicAssociation.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3associationsfromObjectTypetoObjectTypebatchreadRead**
> BatchResponsePublicAssociationMulti Postcrmv3associationsfromObjectTypetoObjectTypebatchreadRead(ctx, fromObjectType, toObjectType, optional)
Read a batch of associations

Get the IDs of all `{toObjectType}` objects associated with those specified in the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fromObjectType** | **string**|  | 
  **toObjectType** | **string**|  | 
 **optional** | ***BatchApiPostcrmv3associationsfromObjectTypetoObjectTypebatchreadReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BatchApiPostcrmv3associationsfromObjectTypetoObjectTypebatchreadReadOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of BatchInputPublicObjectId**](BatchInputPublicObjectId.md)|  | 

### Return type

[**BatchResponsePublicAssociationMulti**](BatchResponsePublicAssociationMulti.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

