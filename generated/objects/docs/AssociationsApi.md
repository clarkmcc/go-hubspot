# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletecrmv3objectsobjectTypeobjectIdassociationstoObjectTypetoObjectIdassociationTypeArchive**](AssociationsApi.md#Deletecrmv3objectsobjectTypeobjectIdassociationstoObjectTypetoObjectIdassociationTypeArchive) | **Delete** /crm/v3/objects/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId}/{associationType} | Remove an association between two objects
[**Getcrmv3objectsobjectTypeobjectIdassociationstoObjectTypeGetAll**](AssociationsApi.md#Getcrmv3objectsobjectTypeobjectIdassociationstoObjectTypeGetAll) | **Get** /crm/v3/objects/{objectType}/{objectId}/associations/{toObjectType} | List associations of an object by type
[**Putcrmv3objectsobjectTypeobjectIdassociationstoObjectTypetoObjectIdassociationTypeCreate**](AssociationsApi.md#Putcrmv3objectsobjectTypeobjectIdassociationstoObjectTypetoObjectIdassociationTypeCreate) | **Put** /crm/v3/objects/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId}/{associationType} | Associate an object with another object

# **Deletecrmv3objectsobjectTypeobjectIdassociationstoObjectTypetoObjectIdassociationTypeArchive**
> Deletecrmv3objectsobjectTypeobjectIdassociationstoObjectTypetoObjectIdassociationTypeArchive(ctx, objectType, objectId, toObjectType, toObjectId, associationType)
Remove an association between two objects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
  **toObjectType** | **string**|  | 
  **toObjectId** | **string**|  | 
  **associationType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcrmv3objectsobjectTypeobjectIdassociationstoObjectTypeGetAll**
> CollectionResponseAssociatedIdForwardPaging Getcrmv3objectsobjectTypeobjectIdassociationstoObjectTypeGetAll(ctx, objectType, objectId, toObjectType, optional)
List associations of an object by type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
  **toObjectType** | **string**|  | 
 **optional** | ***AssociationsApiGetcrmv3objectsobjectTypeobjectIdassociationstoObjectTypeGetAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssociationsApiGetcrmv3objectsobjectTypeobjectIdassociationstoObjectTypeGetAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **after** | **optional.String**| The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **optional.Int32**| The maximum number of results to display per page. | [default to 500]

### Return type

[**CollectionResponseAssociatedIdForwardPaging**](CollectionResponseAssociatedIdForwardPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Putcrmv3objectsobjectTypeobjectIdassociationstoObjectTypetoObjectIdassociationTypeCreate**
> SimplePublicObjectWithAssociations Putcrmv3objectsobjectTypeobjectIdassociationstoObjectTypetoObjectIdassociationTypeCreate(ctx, objectType, objectId, toObjectType, toObjectId, associationType)
Associate an object with another object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
  **toObjectType** | **string**|  | 
  **toObjectId** | **string**|  | 
  **associationType** | **string**|  | 

### Return type

[**SimplePublicObjectWithAssociations**](SimplePublicObjectWithAssociations.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

