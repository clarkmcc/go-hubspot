# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getcrmv3importsGetPage**](CoreApi.md#Getcrmv3importsGetPage) | **Get** /crm/v3/imports/ | Get active imports
[**Getcrmv3importsimportIdGetById**](CoreApi.md#Getcrmv3importsimportIdGetById) | **Get** /crm/v3/imports/{importId} | Get the information on any import
[**Postcrmv3importsCreate**](CoreApi.md#Postcrmv3importsCreate) | **Post** /crm/v3/imports/ | Start a new import
[**Postcrmv3importsimportIdcancelCancel**](CoreApi.md#Postcrmv3importsimportIdcancelCancel) | **Post** /crm/v3/imports/{importId}/cancel | Cancel an active import

# **Getcrmv3importsGetPage**
> CollectionResponsePublicImportResponse Getcrmv3importsGetPage(ctx, optional)
Get active imports

Returns a paged list of active imports for this account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CoreApiGetcrmv3importsGetPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CoreApiGetcrmv3importsGetPageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **optional.String**| The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **before** | **optional.String**|  | 
 **limit** | **optional.Int32**| The maximum number of results to display per page. | 

### Return type

[**CollectionResponsePublicImportResponse**](CollectionResponsePublicImportResponse.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcrmv3importsimportIdGetById**
> PublicImportResponse Getcrmv3importsimportIdGetById(ctx, importId)
Get the information on any import

A complete summary of an import record, including any updates.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **importId** | **int64**|  | 

### Return type

[**PublicImportResponse**](PublicImportResponse.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3importsCreate**
> PublicImportResponse Postcrmv3importsCreate(ctx, optional)
Start a new import

Begins importing data from the specified file resources. This uploads the corresponding file and uses the import request object to convert rows in the files to objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CoreApiPostcrmv3importsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CoreApiPostcrmv3importsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **files** | **optional.Interface of *os.File****optional.**|  | 
 **importRequest** | **optional.**|  | 

### Return type

[**PublicImportResponse**](PublicImportResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3importsimportIdcancelCancel**
> ActionResponse Postcrmv3importsimportIdcancelCancel(ctx, importId)
Cancel an active import

This allows a developer to cancel an active import.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **importId** | **int64**|  | 

### Return type

[**ActionResponse**](ActionResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

