# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getcrmv3ownersGetPage**](OwnersApi.md#Getcrmv3ownersGetPage) | **Get** /crm/v3/owners/ | Get a page of owners
[**Getcrmv3ownersownerIdGetById**](OwnersApi.md#Getcrmv3ownersownerIdGetById) | **Get** /crm/v3/owners/{ownerId} | Read an owner by given &#x60;id&#x60; or &#x60;userId&#x60;

# **Getcrmv3ownersGetPage**
> CollectionResponsePublicOwnerForwardPaging Getcrmv3ownersGetPage(ctx, optional)
Get a page of owners

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OwnersApiGetcrmv3ownersGetPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OwnersApiGetcrmv3ownersGetPageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **optional.String**| Filter by email address (optional) | 
 **after** | **optional.String**| The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **optional.Int32**| The maximum number of results to display per page. | [default to 100]
 **archived** | **optional.Bool**| Whether to return only results that have been archived. | [default to false]

### Return type

[**CollectionResponsePublicOwnerForwardPaging**](CollectionResponsePublicOwnerForwardPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcrmv3ownersownerIdGetById**
> PublicOwner Getcrmv3ownersownerIdGetById(ctx, ownerId, optional)
Read an owner by given `id` or `userId`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerId** | **int32**|  | 
 **optional** | ***OwnersApiGetcrmv3ownersownerIdGetByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OwnersApiGetcrmv3ownersownerIdGetByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idProperty** | **optional.String**|  | [default to id]
 **archived** | **optional.Bool**| Whether to return only results that have been archived. | [default to false]

### Return type

[**PublicOwner**](PublicOwner.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

