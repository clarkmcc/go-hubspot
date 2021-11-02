# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletecrmv3objectslineItemslineItemIdassociationstoObjectTypetoObjectIdassociationTypeArchive**](AssociationsApi.md#Deletecrmv3objectslineItemslineItemIdassociationstoObjectTypetoObjectIdassociationTypeArchive) | **Delete** /crm/v3/objects/line_items/{lineItemId}/associations/{toObjectType}/{toObjectId}/{associationType} | Remove an association between two line items
[**Getcrmv3objectslineItemslineItemIdassociationstoObjectTypeGetAll**](AssociationsApi.md#Getcrmv3objectslineItemslineItemIdassociationstoObjectTypeGetAll) | **Get** /crm/v3/objects/line_items/{lineItemId}/associations/{toObjectType} | List associations of a line item by type
[**Putcrmv3objectslineItemslineItemIdassociationstoObjectTypetoObjectIdassociationTypeCreate**](AssociationsApi.md#Putcrmv3objectslineItemslineItemIdassociationstoObjectTypetoObjectIdassociationTypeCreate) | **Put** /crm/v3/objects/line_items/{lineItemId}/associations/{toObjectType}/{toObjectId}/{associationType} | Associate a line item with another object

# **Deletecrmv3objectslineItemslineItemIdassociationstoObjectTypetoObjectIdassociationTypeArchive**
> Deletecrmv3objectslineItemslineItemIdassociationstoObjectTypetoObjectIdassociationTypeArchive(ctx, lineItemId, toObjectType, toObjectId, associationType)
Remove an association between two line items

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lineItemId** | **string**|  | 
  **toObjectType** | **string**|  | 
  **toObjectId** | **string**|  | 
  **associationType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcrmv3objectslineItemslineItemIdassociationstoObjectTypeGetAll**
> CollectionResponseAssociatedIdForwardPaging Getcrmv3objectslineItemslineItemIdassociationstoObjectTypeGetAll(ctx, lineItemId, toObjectType, optional)
List associations of a line item by type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lineItemId** | **string**|  | 
  **toObjectType** | **string**|  | 
 **optional** | ***AssociationsApiGetcrmv3objectslineItemslineItemIdassociationstoObjectTypeGetAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssociationsApiGetcrmv3objectslineItemslineItemIdassociationstoObjectTypeGetAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **optional.String**| The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **optional.Int32**| The maximum number of results to display per page. | [default to 500]

### Return type

[**CollectionResponseAssociatedIdForwardPaging**](CollectionResponseAssociatedIdForwardPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Putcrmv3objectslineItemslineItemIdassociationstoObjectTypetoObjectIdassociationTypeCreate**
> SimplePublicObjectWithAssociations Putcrmv3objectslineItemslineItemIdassociationstoObjectTypetoObjectIdassociationTypeCreate(ctx, lineItemId, toObjectType, toObjectId, associationType)
Associate a line item with another object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lineItemId** | **string**|  | 
  **toObjectType** | **string**|  | 
  **toObjectId** | **string**|  | 
  **associationType** | **string**|  | 

### Return type

[**SimplePublicObjectWithAssociations**](SimplePublicObjectWithAssociations.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

