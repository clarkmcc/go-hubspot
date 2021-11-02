# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Posteventsv3send**](BehavioralEventsTrackingApi.md#Posteventsv3send) | **Post** /events/v3/send | Sends Custom Behavioral Event

# **Posteventsv3send**
> Posteventsv3send(ctx, body)
Sends Custom Behavioral Event

Endpoint to send an instance of a behavioral event

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BehavioralEventHttpCompletionRequest**](BehavioralEventHttpCompletionRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

