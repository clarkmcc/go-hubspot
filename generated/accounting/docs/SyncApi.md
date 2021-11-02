# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Postcrmv3extensionsaccountingsyncappIdcontactsCreateContact**](SyncApi.md#Postcrmv3extensionsaccountingsyncappIdcontactsCreateContact) | **Post** /crm/v3/extensions/accounting/sync/{appId}/contacts | Import contacts
[**Postcrmv3extensionsaccountingsyncappIdproductsCreateProduct**](SyncApi.md#Postcrmv3extensionsaccountingsyncappIdproductsCreateProduct) | **Post** /crm/v3/extensions/accounting/sync/{appId}/products | Import products

# **Postcrmv3extensionsaccountingsyncappIdcontactsCreateContact**
> ActionResponse Postcrmv3extensionsaccountingsyncappIdcontactsCreateContact(ctx, body, appId)
Import contacts

Imports contacts' properties from an external accounting system to HubSpot. Import details, including property mappings, must be configured previously in HubSpot infrastructure.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SyncContactsRequest**](SyncContactsRequest.md)|  | 
  **appId** | **int32**| The ID of the accounting app. This is the identifier of the application created in your HubSpot developer portal. | 

### Return type

[**ActionResponse**](ActionResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3extensionsaccountingsyncappIdproductsCreateProduct**
> ActionResponse Postcrmv3extensionsaccountingsyncappIdproductsCreateProduct(ctx, body, appId)
Import products

Imports products' properties from an external accounting system to HubSpot. Import details, including property mappings, must be configured previously in HubSpot infrastructure.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SyncProductsRequest**](SyncProductsRequest.md)|  | 
  **appId** | **int32**| The ID of the accounting app. This is the identifier of the application created in your HubSpot developer portal. | 

### Return type

[**ActionResponse**](ActionResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

