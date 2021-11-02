# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getcmsv3sourceCodeenvironmentmetadatapath**](MetadataApi.md#Getcmsv3sourceCodeenvironmentmetadatapath) | **Get** /cms/v3/source-code/{environment}/metadata/{path} | Get the metadata for a file

# **Getcmsv3sourceCodeenvironmentmetadatapath**
> AssetFileMetadata Getcmsv3sourceCodeenvironmentmetadatapath(ctx, environment, path)
Get the metadata for a file

Gets the metadata object for the file at the specified path in the specified environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environment** | **string**| The environment of the file (\&quot;draft\&quot; or \&quot;published\&quot;). | 
  **path** | **string**| The file system location of the file. | 

### Return type

[**AssetFileMetadata**](AssetFileMetadata.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

