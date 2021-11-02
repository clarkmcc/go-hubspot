# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deleteoauthv1refreshTokenstokenArchiveRefreshToken**](RefreshTokensApi.md#Deleteoauthv1refreshTokenstokenArchiveRefreshToken) | **Delete** /oauth/v1/refresh-tokens/{token} | 
[**Getoauthv1refreshTokenstokenGetRefreshToken**](RefreshTokensApi.md#Getoauthv1refreshTokenstokenGetRefreshToken) | **Get** /oauth/v1/refresh-tokens/{token} | 

# **Deleteoauthv1refreshTokenstokenArchiveRefreshToken**
> ModelError Deleteoauthv1refreshTokenstokenArchiveRefreshToken(ctx, token)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getoauthv1refreshTokenstokenGetRefreshToken**
> RefreshTokenInfoResponse Getoauthv1refreshTokenstokenGetRefreshToken(ctx, token)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

[**RefreshTokenInfoResponse**](RefreshTokenInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

