# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Postcmsv3sourceCodeextractpath**](ExtractApi.md#Postcmsv3sourceCodeextractpath) | **Post** /cms/v3/source-code/extract/{path} | Extracts a zip file

# **Postcmsv3sourceCodeextractpath**
> Postcmsv3sourceCodeextractpath(ctx, path)
Extracts a zip file

Extracts a zip file in the file system. The zip file will be extracted in-place and not be deleted automatically.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The file system location of the zip file. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

