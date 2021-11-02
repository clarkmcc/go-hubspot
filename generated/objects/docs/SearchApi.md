# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Postcrmv3objectsobjectTypesearchDoSearch**](SearchApi.md#Postcrmv3objectsobjectTypesearchDoSearch) | **Post** /crm/v3/objects/{objectType}/search | 

# **Postcrmv3objectsobjectTypesearchDoSearch**
> CollectionResponseWithTotalSimplePublicObjectForwardPaging Postcrmv3objectsobjectTypesearchDoSearch(ctx, body, objectType)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PublicObjectSearchRequest**](PublicObjectSearchRequest.md)|  | 
  **objectType** | **string**|  | 

### Return type

[**CollectionResponseWithTotalSimplePublicObjectForwardPaging**](CollectionResponseWithTotalSimplePublicObjectForwardPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

