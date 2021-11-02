# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletecmsv3blogstagsobjectIdArchive**](TagApi.md#Deletecmsv3blogstagsobjectIdArchive) | **Delete** /cms/v3/blogs/tags/{objectId} | Delete a Blog Tag
[**Getcmsv3blogstagsGetPage**](TagApi.md#Getcmsv3blogstagsGetPage) | **Get** /cms/v3/blogs/tags | Get all Blog Tags
[**Getcmsv3blogstagsobjectIdGetById**](TagApi.md#Getcmsv3blogstagsobjectIdGetById) | **Get** /cms/v3/blogs/tags/{objectId} | Retrieve a Blog Tag
[**Patchcmsv3blogstagsobjectIdUpdate**](TagApi.md#Patchcmsv3blogstagsobjectIdUpdate) | **Patch** /cms/v3/blogs/tags/{objectId} | Update a Blog Tag
[**Postcmsv3blogstagsCreate**](TagApi.md#Postcmsv3blogstagsCreate) | **Post** /cms/v3/blogs/tags | Create a new Blog Tag
[**Postcmsv3blogstagsbatcharchiveArchiveBatch**](TagApi.md#Postcmsv3blogstagsbatcharchiveArchiveBatch) | **Post** /cms/v3/blogs/tags/batch/archive | Archive a batch of Blog Tags
[**Postcmsv3blogstagsbatchcreateCreateBatch**](TagApi.md#Postcmsv3blogstagsbatchcreateCreateBatch) | **Post** /cms/v3/blogs/tags/batch/create | Create a batch of Blog Tags
[**Postcmsv3blogstagsbatchreadReadBatch**](TagApi.md#Postcmsv3blogstagsbatchreadReadBatch) | **Post** /cms/v3/blogs/tags/batch/read | Retrieve a batch of Blog Tags
[**Postcmsv3blogstagsbatchupdateUpdateBatch**](TagApi.md#Postcmsv3blogstagsbatchupdateUpdateBatch) | **Post** /cms/v3/blogs/tags/batch/update | Update a batch of Blog Tags

# **Deletecmsv3blogstagsobjectIdArchive**
> Deletecmsv3blogstagsobjectIdArchive(ctx, objectId, optional)
Delete a Blog Tag

Delete the Blog Tag object identified by the id in the path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**| The Blog Tag id. | 
 **optional** | ***TagApiDeletecmsv3blogstagsobjectIdArchiveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TagApiDeletecmsv3blogstagsobjectIdArchiveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **optional.Bool**| Whether to return only results that have been archived. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcmsv3blogstagsGetPage**
> CollectionResponseWithTotalTagForwardPaging Getcmsv3blogstagsGetPage(ctx, optional)
Get all Blog Tags

Get the list of blog tags. Supports paging and filtering. This method would be useful for an integration that examined these models and used an external service to suggest edits. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TagApiGetcmsv3blogstagsGetPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TagApiGetcmsv3blogstagsGetPageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAt** | **optional.Time**| Only return Blog Tags created at exactly the specified time. | 
 **createdAfter** | **optional.Time**| Only return Blog Tags created after the specified time. | 
 **createdBefore** | **optional.Time**| Only return Blog Tags created before the specified time. | 
 **updatedAt** | **optional.Time**| Only return Blog Tags last updated at exactly the specified time. | 
 **updatedAfter** | **optional.Time**| Only return Blog Tags last updated after the specified time. | 
 **updatedBefore** | **optional.Time**| Only return Blog Tags last updated before the specified time. | 
 **sort** | [**optional.Interface of []string**](string.md)| Specifies which fields to use for sorting results. Valid fields are &#x60;name&#x60;, &#x60;createdAt&#x60;, &#x60;updatedAt&#x60;, &#x60;createdBy&#x60;, &#x60;updatedBy&#x60;. &#x60;createdAt&#x60; will be used by default. | 
 **after** | **optional.String**| The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **optional.Int32**| The maximum number of results to return. Default is 100. | 
 **archived** | **optional.Bool**| Specifies whether to return archived Blog Tags. Defaults to &#x60;false&#x60;. | 

### Return type

[**CollectionResponseWithTotalTagForwardPaging**](CollectionResponseWithTotalTagForwardPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcmsv3blogstagsobjectIdGetById**
> Tag Getcmsv3blogstagsobjectIdGetById(ctx, objectId, optional)
Retrieve a Blog Tag

Retrieve the Blog Tag object identified by the id in the path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**| The Blog Tag id. | 
 **optional** | ***TagApiGetcmsv3blogstagsobjectIdGetByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TagApiGetcmsv3blogstagsobjectIdGetByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **optional.Bool**| Specifies whether to return archived Blog Tags. Defaults to &#x60;false&#x60;. | 

### Return type

[**Tag**](Tag.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patchcmsv3blogstagsobjectIdUpdate**
> Tag Patchcmsv3blogstagsobjectIdUpdate(ctx, body, objectId, optional)
Update a Blog Tag

Sparse updates a single Blog Tag object identified by the id in the path. All the column values need not be specified. Only the that need to be modified can be specified. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Tag**](Tag.md)| The JSON representation of the updated Blog Tag. | 
  **objectId** | **string**| The Blog Tag id. | 
 **optional** | ***TagApiPatchcmsv3blogstagsobjectIdUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TagApiPatchcmsv3blogstagsobjectIdUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **archived** | **optional.**| Specifies whether to update archived Blog Tags. Defaults to &#x60;false&#x60;. | 

### Return type

[**Tag**](Tag.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3blogstagsCreate**
> Tag Postcmsv3blogstagsCreate(ctx, body)
Create a new Blog Tag

Create a new Blog Tag.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Tag**](Tag.md)| The JSON representation of a new Blog Tag. | 

### Return type

[**Tag**](Tag.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3blogstagsbatcharchiveArchiveBatch**
> Postcmsv3blogstagsbatcharchiveArchiveBatch(ctx, body)
Archive a batch of Blog Tags

Delete the Blog Tag objects identified in the request body. Note: This is not the same as the in-app `archive` function.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputString**](BatchInputString.md)| The JSON array of Blog Tag ids. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3blogstagsbatchcreateCreateBatch**
> BatchResponseTag Postcmsv3blogstagsbatchcreateCreateBatch(ctx, body)
Create a batch of Blog Tags

Create the Blog Tag objects detailed in the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputTag**](BatchInputTag.md)| The JSON array of new Blog Tags to create. | 

### Return type

[**BatchResponseTag**](BatchResponseTag.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3blogstagsbatchreadReadBatch**
> BatchResponseTag Postcmsv3blogstagsbatchreadReadBatch(ctx, body, optional)
Retrieve a batch of Blog Tags

Retrieve the Blog Tag objects identified in the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputString**](BatchInputString.md)| The JSON array of Blog Tag ids. | 
 **optional** | ***TagApiPostcmsv3blogstagsbatchreadReadBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TagApiPostcmsv3blogstagsbatchreadReadBatchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **optional.**| Specifies whether to return archived Blog Tags. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseTag**](BatchResponseTag.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3blogstagsbatchupdateUpdateBatch**
> BatchResponseTag Postcmsv3blogstagsbatchupdateUpdateBatch(ctx, body, optional)
Update a batch of Blog Tags

Update the Blog Tag objects identified in the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputJsonNode**](BatchInputJsonNode.md)| A JSON array of the JSON representations of the updated Blog Tags. | 
 **optional** | ***TagApiPostcmsv3blogstagsbatchupdateUpdateBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TagApiPostcmsv3blogstagsbatchupdateUpdateBatchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **optional.**| Specifies whether to update archived Blog Tags. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseTag**](BatchResponseTag.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

