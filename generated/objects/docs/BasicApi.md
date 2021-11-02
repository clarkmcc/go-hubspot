# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletecrmv3objectsobjectTypeobjectIdArchive**](BasicApi.md#Deletecrmv3objectsobjectTypeobjectIdArchive) | **Delete** /crm/v3/objects/{objectType}/{objectId} | Archive
[**Getcrmv3objectsobjectTypeGetPage**](BasicApi.md#Getcrmv3objectsobjectTypeGetPage) | **Get** /crm/v3/objects/{objectType} | List
[**Getcrmv3objectsobjectTypeobjectIdGetById**](BasicApi.md#Getcrmv3objectsobjectTypeobjectIdGetById) | **Get** /crm/v3/objects/{objectType}/{objectId} | Read
[**Patchcrmv3objectsobjectTypeobjectIdUpdate**](BasicApi.md#Patchcrmv3objectsobjectTypeobjectIdUpdate) | **Patch** /crm/v3/objects/{objectType}/{objectId} | Update
[**Postcrmv3objectsobjectTypeCreate**](BasicApi.md#Postcrmv3objectsobjectTypeCreate) | **Post** /crm/v3/objects/{objectType} | Create

# **Deletecrmv3objectsobjectTypeobjectIdArchive**
> Deletecrmv3objectsobjectTypeobjectIdArchive(ctx, objectType, objectId)
Archive

Move an Object identified by `{objectId}` to the recycling bin.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcrmv3objectsobjectTypeGetPage**
> CollectionResponseSimplePublicObjectWithAssociationsForwardPaging Getcrmv3objectsobjectTypeGetPage(ctx, objectType, optional)
List

Read a page of objects. Control what is returned via the `properties` query param.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
 **optional** | ***BasicApiGetcrmv3objectsobjectTypeGetPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BasicApiGetcrmv3objectsobjectTypeGetPageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The maximum number of results to display per page. | [default to 10]
 **after** | **optional.String**| The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **properties** | [**optional.Interface of []string**](string.md)| A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored. | 
 **associations** | [**optional.Interface of []string**](string.md)| A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored. | 
 **archived** | **optional.Bool**| Whether to return only results that have been archived. | [default to false]

### Return type

[**CollectionResponseSimplePublicObjectWithAssociationsForwardPaging**](CollectionResponseSimplePublicObjectWithAssociationsForwardPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcrmv3objectsobjectTypeobjectIdGetById**
> SimplePublicObjectWithAssociations Getcrmv3objectsobjectTypeobjectIdGetById(ctx, objectType, objectId, optional)
Read

Read an Object identified by `{objectId}`. `{objectId}` refers to the internal object ID by default, or optionally any unique property value as specified by the `idProperty` query param.  Control what is returned via the `properties` query param.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
 **optional** | ***BasicApiGetcrmv3objectsobjectTypeobjectIdGetByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BasicApiGetcrmv3objectsobjectTypeobjectIdGetByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **properties** | [**optional.Interface of []string**](string.md)| A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored. | 
 **associations** | [**optional.Interface of []string**](string.md)| A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored. | 
 **archived** | **optional.Bool**| Whether to return only results that have been archived. | [default to false]
 **idProperty** | **optional.String**| The name of a property whose values are unique for this object type | 

### Return type

[**SimplePublicObjectWithAssociations**](SimplePublicObjectWithAssociations.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patchcrmv3objectsobjectTypeobjectIdUpdate**
> SimplePublicObject Patchcrmv3objectsobjectTypeobjectIdUpdate(ctx, body, objectType, objectId, optional)
Update

Perform a partial update of an Object identified by `{objectId}`. `{objectId}` refers to the internal object ID by default, or optionally any unique property value as specified by the `idProperty` query param. Provided property values will be overwritten. Read-only and non-existent properties will be ignored. Properties values can be cleared by passing an empty string.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SimplePublicObjectInput**](SimplePublicObjectInput.md)|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
 **optional** | ***BasicApiPatchcrmv3objectsobjectTypeobjectIdUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BasicApiPatchcrmv3objectsobjectTypeobjectIdUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **idProperty** | **optional.**| The name of a property whose values are unique for this object type | 

### Return type

[**SimplePublicObject**](SimplePublicObject.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3objectsobjectTypeCreate**
> SimplePublicObject Postcrmv3objectsobjectTypeCreate(ctx, body, objectType)
Create

Create a CRM object with the given properties and return a copy of the object, including the ID. Documentation and examples for creating standard objects is provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SimplePublicObjectInput**](SimplePublicObjectInput.md)|  | 
  **objectType** | **string**|  | 

### Return type

[**SimplePublicObject**](SimplePublicObject.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

