# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getcmsv3domainsGetPage**](DomainsApi.md#Getcmsv3domainsGetPage) | **Get** /cms/v3/domains/ | Get current domains
[**Getcmsv3domainsdomainIdGetById**](DomainsApi.md#Getcmsv3domainsdomainIdGetById) | **Get** /cms/v3/domains/{domainId} | Get a single domain

# **Getcmsv3domainsGetPage**
> CollectionResponseWithTotalDomain Getcmsv3domainsGetPage(ctx, optional)
Get current domains

Returns all existing domains that have been created. Results can be limited and filtered by creation or updated date.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DomainsApiGetcmsv3domainsGetPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomainsApiGetcmsv3domainsGetPageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAt** | **optional.Int64**| Only return domains created at this date. | 
 **createdAfter** | **optional.Int64**| Only return domains created after this date. | 
 **createdBefore** | **optional.Int64**| Only return domains created before this date. | 
 **updatedAt** | **optional.Int64**| Only return domains updated at this date. | 
 **updatedAfter** | **optional.Int64**| Only return domains updated after this date. | 
 **updatedBefore** | **optional.Int64**| Only return domains updated before this date. | 
 **sort** | [**optional.Interface of []string**](string.md)|  | 
 **properties** | [**optional.Interface of []string**](string.md)|  | 
 **after** | **optional.String**| The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **before** | **optional.String**|  | 
 **limit** | **optional.Int32**| Maximum number of results per page. | 
 **archived** | **optional.Bool**| Whether to return only results that have been archived. | 

### Return type

[**CollectionResponseWithTotalDomain**](CollectionResponseWithTotalDomain.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcmsv3domainsdomainIdGetById**
> Domain Getcmsv3domainsdomainIdGetById(ctx, domainId, optional)
Get a single domain

Returns a single domains with the id specified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| The unique ID of the domain. | 
 **optional** | ***DomainsApiGetcmsv3domainsdomainIdGetByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomainsApiGetcmsv3domainsdomainIdGetByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **optional.Bool**| Whether to return only results that have been archived. | 

### Return type

[**Domain**](Domain.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

