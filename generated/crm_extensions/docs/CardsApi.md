# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletecrmv3extensionscardsappIdcardIdArchive**](CardsApi.md#Deletecrmv3extensionscardsappIdcardIdArchive) | **Delete** /crm/v3/extensions/cards/{appId}/{cardId} | Delete a card
[**Getcrmv3extensionscardsappIdGetAll**](CardsApi.md#Getcrmv3extensionscardsappIdGetAll) | **Get** /crm/v3/extensions/cards/{appId} | Get all cards
[**Getcrmv3extensionscardsappIdcardIdGetById**](CardsApi.md#Getcrmv3extensionscardsappIdcardIdGetById) | **Get** /crm/v3/extensions/cards/{appId}/{cardId} | Get a card.
[**Patchcrmv3extensionscardsappIdcardIdUpdate**](CardsApi.md#Patchcrmv3extensionscardsappIdcardIdUpdate) | **Patch** /crm/v3/extensions/cards/{appId}/{cardId} | Update a card
[**Postcrmv3extensionscardsappIdCreate**](CardsApi.md#Postcrmv3extensionscardsappIdCreate) | **Post** /crm/v3/extensions/cards/{appId} | Create a new card

# **Deletecrmv3extensionscardsappIdcardIdArchive**
> Deletecrmv3extensionscardsappIdcardIdArchive(ctx, appId, cardId)
Delete a card

Permanently deletes a card definition with the given ID. Once deleted, data fetch requests for this card will no longer be sent to your service. This can't be undone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| The ID of the target app. | 
  **cardId** | **string**| The ID of the card to delete. | 

### Return type

 (empty response body)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcrmv3extensionscardsappIdGetAll**
> CardListResponse Getcrmv3extensionscardsappIdGetAll(ctx, appId)
Get all cards

Returns a list of cards for a given app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| The ID of the target app. | 

### Return type

[**CardListResponse**](CardListResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcrmv3extensionscardsappIdcardIdGetById**
> CardResponse Getcrmv3extensionscardsappIdcardIdGetById(ctx, appId, cardId)
Get a card.

Returns the definition for a card with the given ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| The ID of the target app. | 
  **cardId** | **string**| The ID of the target card. | 

### Return type

[**CardResponse**](CardResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patchcrmv3extensionscardsappIdcardIdUpdate**
> CardResponse Patchcrmv3extensionscardsappIdcardIdUpdate(ctx, body, appId, cardId)
Update a card

Update a card definition with new details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CardPatchRequest**](CardPatchRequest.md)| Card definition fields to be updated. | 
  **appId** | **int32**| The ID of the target app. | 
  **cardId** | **string**| The ID of the card to update. | 

### Return type

[**CardResponse**](CardResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3extensionscardsappIdCreate**
> CardResponse Postcrmv3extensionscardsappIdCreate(ctx, body, appId)
Create a new card

Defines a new card that will become active on an account when this app is installed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CardCreateRequest**](CardCreateRequest.md)| The new card definition. | 
  **appId** | **int32**| The ID of the target app. | 

### Return type

[**CardResponse**](CardResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

