# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletecmsv3sourceCodeenvironmentcontentpath**](ContentApi.md#Deletecmsv3sourceCodeenvironmentcontentpath) | **Delete** /cms/v3/source-code/{environment}/content/{path} | Delete a file
[**Getcmsv3sourceCodeenvironmentcontentpath**](ContentApi.md#Getcmsv3sourceCodeenvironmentcontentpath) | **Get** /cms/v3/source-code/{environment}/content/{path} | Download a file
[**Postcmsv3sourceCodeenvironmentcontentpath**](ContentApi.md#Postcmsv3sourceCodeenvironmentcontentpath) | **Post** /cms/v3/source-code/{environment}/content/{path} | Create a file
[**Putcmsv3sourceCodeenvironmentcontentpath**](ContentApi.md#Putcmsv3sourceCodeenvironmentcontentpath) | **Put** /cms/v3/source-code/{environment}/content/{path} | Create or update a file

# **Deletecmsv3sourceCodeenvironmentcontentpath**
> Deletecmsv3sourceCodeenvironmentcontentpath(ctx, environment, path)
Delete a file

Deletes the file at the specified path in the specified environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environment** | **string**| The environment of the file (\&quot;draft\&quot; or \&quot;published\&quot;). | 
  **path** | **string**| The file system location of the file. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcmsv3sourceCodeenvironmentcontentpath**
> ModelError Getcmsv3sourceCodeenvironmentcontentpath(ctx, environment, path)
Download a file

Downloads the byte contents of the file at the specified path in the specified environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environment** | **string**| The environment of the file (\&quot;draft\&quot; or \&quot;published\&quot;). | 
  **path** | **string**| The file system location of the file. | 

### Return type

[**ModelError**](Error.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3sourceCodeenvironmentcontentpath**
> AssetFileMetadata Postcmsv3sourceCodeenvironmentcontentpath(ctx, environment, path, optional)
Create a file

Creates a file at the specified path in the specified environment. Accepts multipart/form-data content type. Throws an error if a file already exists at the specified path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environment** | **string**| The environment of the file (\&quot;draft\&quot; or \&quot;published\&quot;). | 
  **path** | **string**| The file system location of the file. | 
 **optional** | ***ContentApiPostcmsv3sourceCodeenvironmentcontentpathOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentApiPostcmsv3sourceCodeenvironmentcontentpathOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | **optional.Interface of *os.File****optional.**|  | 

### Return type

[**AssetFileMetadata**](AssetFileMetadata.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Putcmsv3sourceCodeenvironmentcontentpath**
> AssetFileMetadata Putcmsv3sourceCodeenvironmentcontentpath(ctx, environment, path, optional)
Create or update a file

Upserts a file at the specified path in the specified environment. Accepts multipart/form-data content type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environment** | **string**| The environment of the file (\&quot;draft\&quot; or \&quot;published\&quot;). | 
  **path** | **string**| The file system location of the file. | 
 **optional** | ***ContentApiPutcmsv3sourceCodeenvironmentcontentpathOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentApiPutcmsv3sourceCodeenvironmentcontentpathOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | **optional.Interface of *os.File****optional.**|  | 

### Return type

[**AssetFileMetadata**](AssetFileMetadata.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

