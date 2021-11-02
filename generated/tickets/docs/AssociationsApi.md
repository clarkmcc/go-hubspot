# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletecrmv3objectsticketsticketIdassociationstoObjectTypetoObjectIdassociationTypeArchive**](AssociationsApi.md#Deletecrmv3objectsticketsticketIdassociationstoObjectTypetoObjectIdassociationTypeArchive) | **Delete** /crm/v3/objects/tickets/{ticketId}/associations/{toObjectType}/{toObjectId}/{associationType} | Remove an association between two tickets
[**Getcrmv3objectsticketsticketIdassociationstoObjectTypeGetAll**](AssociationsApi.md#Getcrmv3objectsticketsticketIdassociationstoObjectTypeGetAll) | **Get** /crm/v3/objects/tickets/{ticketId}/associations/{toObjectType} | List associations of a ticket by type
[**Putcrmv3objectsticketsticketIdassociationstoObjectTypetoObjectIdassociationTypeCreate**](AssociationsApi.md#Putcrmv3objectsticketsticketIdassociationstoObjectTypetoObjectIdassociationTypeCreate) | **Put** /crm/v3/objects/tickets/{ticketId}/associations/{toObjectType}/{toObjectId}/{associationType} | Associate a ticket with another object

# **Deletecrmv3objectsticketsticketIdassociationstoObjectTypetoObjectIdassociationTypeArchive**
> Deletecrmv3objectsticketsticketIdassociationstoObjectTypetoObjectIdassociationTypeArchive(ctx, ticketId, toObjectType, toObjectId, associationType)
Remove an association between two tickets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ticketId** | **string**|  | 
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

# **Getcrmv3objectsticketsticketIdassociationstoObjectTypeGetAll**
> CollectionResponseAssociatedIdForwardPaging Getcrmv3objectsticketsticketIdassociationstoObjectTypeGetAll(ctx, ticketId, toObjectType, optional)
List associations of a ticket by type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ticketId** | **string**|  | 
  **toObjectType** | **string**|  | 
 **optional** | ***AssociationsApiGetcrmv3objectsticketsticketIdassociationstoObjectTypeGetAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssociationsApiGetcrmv3objectsticketsticketIdassociationstoObjectTypeGetAllOpts struct
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

# **Putcrmv3objectsticketsticketIdassociationstoObjectTypetoObjectIdassociationTypeCreate**
> SimplePublicObjectWithAssociations Putcrmv3objectsticketsticketIdassociationstoObjectTypetoObjectIdassociationTypeCreate(ctx, ticketId, toObjectType, toObjectId, associationType)
Associate a ticket with another object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ticketId** | **string**|  | 
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

