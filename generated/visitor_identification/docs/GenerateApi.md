# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostvisitorIdentificationv3tokenscreateGenerateToken**](GenerateApi.md#PostvisitorIdentificationv3tokenscreateGenerateToken) | **Post** /conversations/v3/visitor-identification/tokens/create | Generate a token

# **PostvisitorIdentificationv3tokenscreateGenerateToken**
> IdentificationTokenResponse PostvisitorIdentificationv3tokenscreateGenerateToken(ctx, body)
Generate a token

Generates a new visitor identification token. This token will be unique every time this endpoint is called, even if called with the same email address. This token is temporary and will expire after 12 hours

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdentificationTokenGenerationRequest**](IdentificationTokenGenerationRequest.md)|  | 

### Return type

[**IdentificationTokenResponse**](IdentificationTokenResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

