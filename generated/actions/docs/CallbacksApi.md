# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Postautomationv4actionscallbackscallbackIdcompleteComplete**](CallbacksApi.md#Postautomationv4actionscallbackscallbackIdcompleteComplete) | **Post** /automation/v4/actions/callbacks/{callbackId}/complete | Complete a callback
[**Postautomationv4actionscallbackscompleteCompleteBatch**](CallbacksApi.md#Postautomationv4actionscallbackscompleteCompleteBatch) | **Post** /automation/v4/actions/callbacks/complete | Complete a batch of callbacks

# **Postautomationv4actionscallbackscallbackIdcompleteComplete**
> Postautomationv4actionscallbackscallbackIdcompleteComplete(ctx, body, callbackId)
Complete a callback

Completes the given action callback.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CallbackCompletionRequest**](CallbackCompletionRequest.md)| The result of the completed action. | 
  **callbackId** | **string**| The ID of the target app. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postautomationv4actionscallbackscompleteCompleteBatch**
> Postautomationv4actionscallbackscompleteCompleteBatch(ctx, body)
Complete a batch of callbacks

Completes the given action callbacks.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputCallbackCompletionBatchRequest**](BatchInputCallbackCompletionBatchRequest.md)| The result of the completed action. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

