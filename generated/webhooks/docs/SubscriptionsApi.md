# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletewebhooksv3appIdsubscriptionssubscriptionIdArchive**](SubscriptionsApi.md#Deletewebhooksv3appIdsubscriptionssubscriptionIdArchive) | **Delete** /webhooks/v3/{appId}/subscriptions/{subscriptionId} | Delete a subscription
[**Getwebhooksv3appIdsubscriptionsGetAll**](SubscriptionsApi.md#Getwebhooksv3appIdsubscriptionsGetAll) | **Get** /webhooks/v3/{appId}/subscriptions | Get subscription details
[**Getwebhooksv3appIdsubscriptionssubscriptionIdGetById**](SubscriptionsApi.md#Getwebhooksv3appIdsubscriptionssubscriptionIdGetById) | **Get** /webhooks/v3/{appId}/subscriptions/{subscriptionId} | Get subscription
[**Patchwebhooksv3appIdsubscriptionssubscriptionIdUpdate**](SubscriptionsApi.md#Patchwebhooksv3appIdsubscriptionssubscriptionIdUpdate) | **Patch** /webhooks/v3/{appId}/subscriptions/{subscriptionId} | Update a subscription
[**Postwebhooksv3appIdsubscriptionsCreate**](SubscriptionsApi.md#Postwebhooksv3appIdsubscriptionsCreate) | **Post** /webhooks/v3/{appId}/subscriptions | Subscribe to an event
[**Postwebhooksv3appIdsubscriptionsbatchupdateUpdateBatch**](SubscriptionsApi.md#Postwebhooksv3appIdsubscriptionsbatchupdateUpdateBatch) | **Post** /webhooks/v3/{appId}/subscriptions/batch/update | Batch update subscriptions

# **Deletewebhooksv3appIdsubscriptionssubscriptionIdArchive**
> Deletewebhooksv3appIdsubscriptionssubscriptionIdArchive(ctx, subscriptionId, appId)
Delete a subscription

Permanently deletes a subscription. This cannot be undone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **int32**| The ID of subscription to delete. | 
  **appId** | **int32**| The ID of the target app. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getwebhooksv3appIdsubscriptionsGetAll**
> SubscriptionListResponse Getwebhooksv3appIdsubscriptionsGetAll(ctx, appId)
Get subscription details

Returns full details for all existing subscriptions for the given app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| The ID of the target app. | 

### Return type

[**SubscriptionListResponse**](SubscriptionListResponse.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getwebhooksv3appIdsubscriptionssubscriptionIdGetById**
> SubscriptionResponse Getwebhooksv3appIdsubscriptionssubscriptionIdGetById(ctx, appId, subscriptionId)
Get subscription

Returns details about a subscription.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| The ID of the target app. | 
  **subscriptionId** | **int32**| The ID of the target subscription. | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patchwebhooksv3appIdsubscriptionssubscriptionIdUpdate**
> SubscriptionResponse Patchwebhooksv3appIdsubscriptionssubscriptionIdUpdate(ctx, body, subscriptionId, appId)
Update a subscription

Updates the details for an existing subscription.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubscriptionPatchRequest**](SubscriptionPatchRequest.md)| Updated details for the subscription. | 
  **subscriptionId** | **int32**| The ID of the subscription to update. | 
  **appId** | **int32**| The ID of the target app. | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postwebhooksv3appIdsubscriptionsCreate**
> SubscriptionResponse Postwebhooksv3appIdsubscriptionsCreate(ctx, body, appId)
Subscribe to an event

Creates a new webhook subscription for the given app. Each subscription in an app must be unique.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubscriptionCreateRequest**](SubscriptionCreateRequest.md)| Details about the new subscription. | 
  **appId** | **int32**| The ID of the target app. | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postwebhooksv3appIdsubscriptionsbatchupdateUpdateBatch**
> BatchResponseSubscriptionResponse Postwebhooksv3appIdsubscriptionsbatchupdateUpdateBatch(ctx, body, appId)
Batch update subscriptions

Activates or deactivates target app subscriptions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputSubscriptionBatchUpdateRequest**](BatchInputSubscriptionBatchUpdateRequest.md)| Updated details for the specified subscriptions. | 
  **appId** | **int32**| The app ID of the target app. | 

### Return type

[**BatchResponseSubscriptionResponse**](BatchResponseSubscriptionResponse.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

