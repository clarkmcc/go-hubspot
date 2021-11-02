# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletecmsv3urlRedirectsurlRedirectIdArchive**](RedirectsApi.md#Deletecmsv3urlRedirectsurlRedirectIdArchive) | **Delete** /cms/v3/url-redirects/{urlRedirectId} | Delete a redirect
[**Getcmsv3urlRedirectsGetPage**](RedirectsApi.md#Getcmsv3urlRedirectsGetPage) | **Get** /cms/v3/url-redirects/ | Get current redirects
[**Getcmsv3urlRedirectsurlRedirectIdGetById**](RedirectsApi.md#Getcmsv3urlRedirectsurlRedirectIdGetById) | **Get** /cms/v3/url-redirects/{urlRedirectId} | Get details for a redirect
[**Patchcmsv3urlRedirectsurlRedirectIdUpdate**](RedirectsApi.md#Patchcmsv3urlRedirectsurlRedirectIdUpdate) | **Patch** /cms/v3/url-redirects/{urlRedirectId} | Update a redirect
[**Postcmsv3urlRedirectsCreate**](RedirectsApi.md#Postcmsv3urlRedirectsCreate) | **Post** /cms/v3/url-redirects/ | Create a redirect

# **Deletecmsv3urlRedirectsurlRedirectIdArchive**
> Deletecmsv3urlRedirectsurlRedirectIdArchive(ctx, urlRedirectId)
Delete a redirect

Delete one existing redirect, so it is no longer mapped.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **urlRedirectId** | **string**| The ID of the target redirect. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcmsv3urlRedirectsGetPage**
> CollectionResponseWithTotalUrlMapping Getcmsv3urlRedirectsGetPage(ctx, optional)
Get current redirects

Returns all existing URL redirects. Results can be limited and filtered by creation or updated date.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RedirectsApiGetcmsv3urlRedirectsGetPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RedirectsApiGetcmsv3urlRedirectsGetPageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAt** | **optional.Time**| Only return redirects created on exactly this date. | 
 **createdAfter** | **optional.Time**| Only return redirects created after this date. | 
 **createdBefore** | **optional.Time**| Only return redirects created before this date. | 
 **updatedAt** | **optional.Time**| Only return redirects last updated on exactly this date. | 
 **updatedAfter** | **optional.Time**| Only return redirects last updated after this date. | 
 **updatedBefore** | **optional.Time**| Only return redirects last updated before this date. | 
 **sort** | [**optional.Interface of []string**](string.md)|  | 
 **properties** | [**optional.Interface of []string**](string.md)|  | 
 **after** | **optional.String**| The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **before** | **optional.String**|  | 
 **limit** | **optional.Int32**| Maximum number of result per page | 
 **archived** | **optional.Bool**| Whether to return only results that have been archived. | 

### Return type

[**CollectionResponseWithTotalUrlMapping**](CollectionResponseWithTotalUrlMapping.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcmsv3urlRedirectsurlRedirectIdGetById**
> UrlMapping Getcmsv3urlRedirectsurlRedirectIdGetById(ctx, urlRedirectId)
Get details for a redirect

Returns the details for a single existing URL redirect by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **urlRedirectId** | **string**| The ID of the target redirect. | 

### Return type

[**UrlMapping**](UrlMapping.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patchcmsv3urlRedirectsurlRedirectIdUpdate**
> UrlMapping Patchcmsv3urlRedirectsurlRedirectIdUpdate(ctx, urlRedirectId, optional)
Update a redirect

Updates the settings for an existing URL redirect.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **urlRedirectId** | **string**|  | 
 **optional** | ***RedirectsApiPatchcmsv3urlRedirectsurlRedirectIdUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RedirectsApiPatchcmsv3urlRedirectsurlRedirectIdUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UrlMapping**](UrlMapping.md)|  | 

### Return type

[**UrlMapping**](UrlMapping.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3urlRedirectsCreate**
> UrlMapping Postcmsv3urlRedirectsCreate(ctx, optional)
Create a redirect

Creates and configures a new URL redirect.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RedirectsApiPostcmsv3urlRedirectsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RedirectsApiPostcmsv3urlRedirectsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UrlMappingCreateRequestBody**](UrlMappingCreateRequestBody.md)|  | 

### Return type

[**UrlMapping**](UrlMapping.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

