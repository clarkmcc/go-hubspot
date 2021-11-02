# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletecrmv3objectscompaniescompanyIdArchive**](BasicApi.md#Deletecrmv3objectscompaniescompanyIdArchive) | **Delete** /crm/v3/objects/companies/{companyId} | Archive
[**Getcrmv3objectscompaniesGetPage**](BasicApi.md#Getcrmv3objectscompaniesGetPage) | **Get** /crm/v3/objects/companies | List
[**Getcrmv3objectscompaniescompanyIdGetById**](BasicApi.md#Getcrmv3objectscompaniescompanyIdGetById) | **Get** /crm/v3/objects/companies/{companyId} | Read
[**Patchcrmv3objectscompaniescompanyIdUpdate**](BasicApi.md#Patchcrmv3objectscompaniescompanyIdUpdate) | **Patch** /crm/v3/objects/companies/{companyId} | Update
[**Postcrmv3objectscompaniesCreate**](BasicApi.md#Postcrmv3objectscompaniesCreate) | **Post** /crm/v3/objects/companies | Create

# **Deletecrmv3objectscompaniescompanyIdArchive**
> Deletecrmv3objectscompaniescompanyIdArchive(ctx, companyId)
Archive

Move an Object identified by `{companyId}` to the recycling bin.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcrmv3objectscompaniesGetPage**
> CollectionResponseSimplePublicObjectWithAssociationsForwardPaging Getcrmv3objectscompaniesGetPage(ctx, optional)
List

Read a page of companies. Control what is returned via the `properties` query param.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BasicApiGetcrmv3objectscompaniesGetPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BasicApiGetcrmv3objectscompaniesGetPageOpts struct
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

# **Getcrmv3objectscompaniescompanyIdGetById**
> SimplePublicObjectWithAssociations Getcrmv3objectscompaniescompanyIdGetById(ctx, companyId, optional)
Read

Read an Object identified by `{companyId}`. `{companyId}` refers to the internal object ID by default, or optionally any unique property value as specified by the `idProperty` query param.  Control what is returned via the `properties` query param.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | **string**|  | 
 **optional** | ***BasicApiGetcrmv3objectscompaniescompanyIdGetByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BasicApiGetcrmv3objectscompaniescompanyIdGetByIdOpts struct
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

# **Patchcrmv3objectscompaniescompanyIdUpdate**
> SimplePublicObject Patchcrmv3objectscompaniescompanyIdUpdate(ctx, body, companyId, optional)
Update

Perform a partial update of an Object identified by `{companyId}`. `{companyId}` refers to the internal object ID by default, or optionally any unique property value as specified by the `idProperty` query param. Provided property values will be overwritten. Read-only and non-existent properties will be ignored. Properties values can be cleared by passing an empty string.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SimplePublicObjectInput**](SimplePublicObjectInput.md)|  | 
  **companyId** | **string**|  | 
 **optional** | ***BasicApiPatchcrmv3objectscompaniescompanyIdUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BasicApiPatchcrmv3objectscompaniescompanyIdUpdateOpts struct
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

# **Postcrmv3objectscompaniesCreate**
> SimplePublicObject Postcrmv3objectscompaniesCreate(ctx, body)
Create

Create a company with the given properties and return a copy of the object, including the ID. Documentation and examples for creating standard companies is provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SimplePublicObjectInput**](SimplePublicObjectInput.md)|  | 

### Return type

[**SimplePublicObject**](SimplePublicObject.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

