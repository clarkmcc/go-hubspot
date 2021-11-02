# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Postcmsv3sourceCodeenvironmentvalidatepath**](ValidationApi.md#Postcmsv3sourceCodeenvironmentvalidatepath) | **Post** /cms/v3/source-code/{environment}/validate/{path} | Validate the contents of a file

# **Postcmsv3sourceCodeenvironmentvalidatepath**
> ModelError Postcmsv3sourceCodeenvironmentvalidatepath(ctx, path, optional)
Validate the contents of a file

Validates the file contents passed to the endpoint given a specified path and environment. Accepts multipart/form-data content type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The file system location of the file. | 
 **optional** | ***ValidationApiPostcmsv3sourceCodeenvironmentvalidatepathOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidationApiPostcmsv3sourceCodeenvironmentvalidatepathOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **optional.Interface of *os.File****optional.**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

