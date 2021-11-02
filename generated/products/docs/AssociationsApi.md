# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletecrmv3objectsproductsproductIdassociationstoObjectTypetoObjectIdassociationTypeArchive**](AssociationsApi.md#Deletecrmv3objectsproductsproductIdassociationstoObjectTypetoObjectIdassociationTypeArchive) | **Delete** /crm/v3/objects/products/{productId}/associations/{toObjectType}/{toObjectId}/{associationType} | Remove an association between two products
[**Getcrmv3objectsproductsproductIdassociationstoObjectTypeGetAll**](AssociationsApi.md#Getcrmv3objectsproductsproductIdassociationstoObjectTypeGetAll) | **Get** /crm/v3/objects/products/{productId}/associations/{toObjectType} | List associations of a product by type
[**Putcrmv3objectsproductsproductIdassociationstoObjectTypetoObjectIdassociationTypeCreate**](AssociationsApi.md#Putcrmv3objectsproductsproductIdassociationstoObjectTypetoObjectIdassociationTypeCreate) | **Put** /crm/v3/objects/products/{productId}/associations/{toObjectType}/{toObjectId}/{associationType} | Associate a product with another object

# **Deletecrmv3objectsproductsproductIdassociationstoObjectTypetoObjectIdassociationTypeArchive**
> Deletecrmv3objectsproductsproductIdassociationstoObjectTypetoObjectIdassociationTypeArchive(ctx, productId, toObjectType, toObjectId, associationType)
Remove an association between two products

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | **string**|  | 
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

# **Getcrmv3objectsproductsproductIdassociationstoObjectTypeGetAll**
> CollectionResponseAssociatedIdForwardPaging Getcrmv3objectsproductsproductIdassociationstoObjectTypeGetAll(ctx, productId, toObjectType, optional)
List associations of a product by type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | **string**|  | 
  **toObjectType** | **string**|  | 
 **optional** | ***AssociationsApiGetcrmv3objectsproductsproductIdassociationstoObjectTypeGetAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssociationsApiGetcrmv3objectsproductsproductIdassociationstoObjectTypeGetAllOpts struct
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

# **Putcrmv3objectsproductsproductIdassociationstoObjectTypetoObjectIdassociationTypeCreate**
> SimplePublicObjectWithAssociations Putcrmv3objectsproductsproductIdassociationstoObjectTypetoObjectIdassociationTypeCreate(ctx, productId, toObjectType, toObjectId, associationType)
Associate a product with another object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | **string**|  | 
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

