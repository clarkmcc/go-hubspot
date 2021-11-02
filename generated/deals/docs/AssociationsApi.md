# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletecrmv3objectsdealsdealIdassociationstoObjectTypetoObjectIdassociationTypeArchive**](AssociationsApi.md#Deletecrmv3objectsdealsdealIdassociationstoObjectTypetoObjectIdassociationTypeArchive) | **Delete** /crm/v3/objects/deals/{dealId}/associations/{toObjectType}/{toObjectId}/{associationType} | Remove an association between two deals
[**Getcrmv3objectsdealsdealIdassociationstoObjectTypeGetAll**](AssociationsApi.md#Getcrmv3objectsdealsdealIdassociationstoObjectTypeGetAll) | **Get** /crm/v3/objects/deals/{dealId}/associations/{toObjectType} | List associations of a deal by type
[**Putcrmv3objectsdealsdealIdassociationstoObjectTypetoObjectIdassociationTypeCreate**](AssociationsApi.md#Putcrmv3objectsdealsdealIdassociationstoObjectTypetoObjectIdassociationTypeCreate) | **Put** /crm/v3/objects/deals/{dealId}/associations/{toObjectType}/{toObjectId}/{associationType} | Associate a deal with another object

# **Deletecrmv3objectsdealsdealIdassociationstoObjectTypetoObjectIdassociationTypeArchive**
> Deletecrmv3objectsdealsdealIdassociationstoObjectTypetoObjectIdassociationTypeArchive(ctx, dealId, toObjectType, toObjectId, associationType)
Remove an association between two deals

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dealId** | **string**|  | 
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

# **Getcrmv3objectsdealsdealIdassociationstoObjectTypeGetAll**
> CollectionResponseAssociatedIdForwardPaging Getcrmv3objectsdealsdealIdassociationstoObjectTypeGetAll(ctx, dealId, toObjectType, optional)
List associations of a deal by type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dealId** | **string**|  | 
  **toObjectType** | **string**|  | 
 **optional** | ***AssociationsApiGetcrmv3objectsdealsdealIdassociationstoObjectTypeGetAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssociationsApiGetcrmv3objectsdealsdealIdassociationstoObjectTypeGetAllOpts struct
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

# **Putcrmv3objectsdealsdealIdassociationstoObjectTypetoObjectIdassociationTypeCreate**
> SimplePublicObjectWithAssociations Putcrmv3objectsdealsdealIdassociationstoObjectTypetoObjectIdassociationTypeCreate(ctx, dealId, toObjectType, toObjectId, associationType)
Associate a deal with another object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dealId** | **string**|  | 
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

