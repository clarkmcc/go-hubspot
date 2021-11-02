# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getcrmv3importsimportIderrorsGetErrors**](PublicImportsApi.md#Getcrmv3importsimportIderrorsGetErrors) | **Get** /crm/v3/imports/{importId}/errors | 

# **Getcrmv3importsimportIderrorsGetErrors**
> CollectionResponsePublicImportErrorForwardPaging Getcrmv3importsimportIderrorsGetErrors(ctx, importId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **importId** | **int64**|  | 
 **optional** | ***PublicImportsApiGetcrmv3importsimportIderrorsGetErrorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicImportsApiGetcrmv3importsimportIderrorsGetErrorsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **optional.String**| The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **optional.Int32**| The maximum number of results to display per page. | 

### Return type

[**CollectionResponsePublicImportErrorForwardPaging**](CollectionResponsePublicImportErrorForwardPaging.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

