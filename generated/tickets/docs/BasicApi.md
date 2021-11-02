# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletecrmv3objectsticketsticketIdArchive**](BasicApi.md#Deletecrmv3objectsticketsticketIdArchive) | **Delete** /crm/v3/objects/tickets/{ticketId} | Archive
[**Getcrmv3objectsticketsGetPage**](BasicApi.md#Getcrmv3objectsticketsGetPage) | **Get** /crm/v3/objects/tickets | List
[**Getcrmv3objectsticketsticketIdGetById**](BasicApi.md#Getcrmv3objectsticketsticketIdGetById) | **Get** /crm/v3/objects/tickets/{ticketId} | Read
[**Patchcrmv3objectsticketsticketIdUpdate**](BasicApi.md#Patchcrmv3objectsticketsticketIdUpdate) | **Patch** /crm/v3/objects/tickets/{ticketId} | Update
[**Postcrmv3objectsticketsCreate**](BasicApi.md#Postcrmv3objectsticketsCreate) | **Post** /crm/v3/objects/tickets | Create

# **Deletecrmv3objectsticketsticketIdArchive**
> Deletecrmv3objectsticketsticketIdArchive(ctx, ticketId)
Archive

Move an Object identified by `{ticketId}` to the recycling bin.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ticketId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcrmv3objectsticketsGetPage**
> CollectionResponseSimplePublicObjectWithAssociationsForwardPaging Getcrmv3objectsticketsGetPage(ctx, optional)
List

Read a page of tickets. Control what is returned via the `properties` query param.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BasicApiGetcrmv3objectsticketsGetPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BasicApiGetcrmv3objectsticketsGetPageOpts struct
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

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcrmv3objectsticketsticketIdGetById**
> SimplePublicObjectWithAssociations Getcrmv3objectsticketsticketIdGetById(ctx, ticketId, optional)
Read

Read an Object identified by `{ticketId}`. `{ticketId}` refers to the internal object ID by default, or optionally any unique property value as specified by the `idProperty` query param.  Control what is returned via the `properties` query param.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ticketId** | **string**|  | 
 **optional** | ***BasicApiGetcrmv3objectsticketsticketIdGetByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BasicApiGetcrmv3objectsticketsticketIdGetByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **properties** | [**optional.Interface of []string**](string.md)| A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored. | 
 **associations** | [**optional.Interface of []string**](string.md)| A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored. | 
 **archived** | **optional.Bool**| Whether to return only results that have been archived. | [default to false]
 **idProperty** | **optional.String**| The name of a property whose values are unique for this object type | 

### Return type

[**SimplePublicObjectWithAssociations**](SimplePublicObjectWithAssociations.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patchcrmv3objectsticketsticketIdUpdate**
> SimplePublicObject Patchcrmv3objectsticketsticketIdUpdate(ctx, body, ticketId, optional)
Update

Perform a partial update of an Object identified by `{ticketId}`. `{ticketId}` refers to the internal object ID by default, or optionally any unique property value as specified by the `idProperty` query param. Provided property values will be overwritten. Read-only and non-existent properties will be ignored. Properties values can be cleared by passing an empty string.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SimplePublicObjectInput**](SimplePublicObjectInput.md)|  | 
  **ticketId** | **string**|  | 
 **optional** | ***BasicApiPatchcrmv3objectsticketsticketIdUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BasicApiPatchcrmv3objectsticketsticketIdUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idProperty** | **optional.**| The name of a property whose values are unique for this object type | 

### Return type

[**SimplePublicObject**](SimplePublicObject.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3objectsticketsCreate**
> SimplePublicObject Postcrmv3objectsticketsCreate(ctx, body)
Create

Create a ticket with the given properties and return a copy of the object, including the ID. Documentation and examples for creating standard tickets is provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SimplePublicObjectInput**](SimplePublicObjectInput.md)|  | 

### Return type

[**SimplePublicObject**](SimplePublicObject.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

