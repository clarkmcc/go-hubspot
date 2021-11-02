# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getcrmv3associationsfromObjectTypetoObjectTypetypesGetAll**](TypesApi.md#Getcrmv3associationsfromObjectTypetoObjectTypetypesGetAll) | **Get** /crm/v3/associations/{fromObjectType}/{toObjectType}/types | List association types

# **Getcrmv3associationsfromObjectTypetoObjectTypetypesGetAll**
> CollectionResponsePublicAssociationDefiniton Getcrmv3associationsfromObjectTypetoObjectTypetypesGetAll(ctx, fromObjectType, toObjectType)
List association types

List all the valid association types available between two object types

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fromObjectType** | **string**|  | 
  **toObjectType** | **string**|  | 

### Return type

[**CollectionResponsePublicAssociationDefiniton**](CollectionResponsePublicAssociationDefiniton.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

