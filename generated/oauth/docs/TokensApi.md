# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Postoauthv1tokenCreateToken**](TokensApi.md#Postoauthv1tokenCreateToken) | **Post** /oauth/v1/token | 

# **Postoauthv1tokenCreateToken**
> TokenResponseIf Postoauthv1tokenCreateToken(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TokensApiPostoauthv1tokenCreateTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokensApiPostoauthv1tokenCreateTokenOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **optional.**|  | 
 **code** | **optional.**|  | 
 **redirectUri** | **optional.**|  | 
 **clientId** | **optional.**|  | 
 **clientSecret** | **optional.**|  | 
 **refreshToken** | **optional.**|  | 

### Return type

[**TokenResponseIf**](TokenResponseIF.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

