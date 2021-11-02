# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletecmsv3blogsauthorsobjectIdArchive**](AuthorApi.md#Deletecmsv3blogsauthorsobjectIdArchive) | **Delete** /cms/v3/blogs/authors/{objectId} | Delete a Blog Author
[**Getcmsv3blogsauthorsGetPage**](AuthorApi.md#Getcmsv3blogsauthorsGetPage) | **Get** /cms/v3/blogs/authors | Get all Blog Authors
[**Getcmsv3blogsauthorsobjectIdGetById**](AuthorApi.md#Getcmsv3blogsauthorsobjectIdGetById) | **Get** /cms/v3/blogs/authors/{objectId} | Retrieve a Blog Author
[**Patchcmsv3blogsauthorsobjectIdUpdate**](AuthorApi.md#Patchcmsv3blogsauthorsobjectIdUpdate) | **Patch** /cms/v3/blogs/authors/{objectId} | Update a Blog Author
[**Postcmsv3blogsauthorsCreate**](AuthorApi.md#Postcmsv3blogsauthorsCreate) | **Post** /cms/v3/blogs/authors | Create a new Blog Author
[**Postcmsv3blogsauthorsbatcharchiveArchiveBatch**](AuthorApi.md#Postcmsv3blogsauthorsbatcharchiveArchiveBatch) | **Post** /cms/v3/blogs/authors/batch/archive | Archive a batch of Blog Authors
[**Postcmsv3blogsauthorsbatchcreateCreateBatch**](AuthorApi.md#Postcmsv3blogsauthorsbatchcreateCreateBatch) | **Post** /cms/v3/blogs/authors/batch/create | Create a batch of Blog Authors
[**Postcmsv3blogsauthorsbatchreadReadBatch**](AuthorApi.md#Postcmsv3blogsauthorsbatchreadReadBatch) | **Post** /cms/v3/blogs/authors/batch/read | Retrieve a batch of Blog Authors
[**Postcmsv3blogsauthorsbatchupdateUpdateBatch**](AuthorApi.md#Postcmsv3blogsauthorsbatchupdateUpdateBatch) | **Post** /cms/v3/blogs/authors/batch/update | Update a batch of Blog Authors

# **Deletecmsv3blogsauthorsobjectIdArchive**
> Deletecmsv3blogsauthorsobjectIdArchive(ctx, objectId, optional)
Delete a Blog Author

Delete the Blog Author object identified by the id in the path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**| The Blog Author id. | 
 **optional** | ***AuthorApiDeletecmsv3blogsauthorsobjectIdArchiveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorApiDeletecmsv3blogsauthorsobjectIdArchiveOpts struct
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

# **Getcmsv3blogsauthorsGetPage**
> CollectionResponseWithTotalBlogAuthorForwardPaging Getcmsv3blogsauthorsGetPage(ctx, optional)
Get all Blog Authors

Get the list of blog authors. Supports paging and filtering. This method would be useful for an integration that examined these models and used an external service to suggest edits. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthorApiGetcmsv3blogsauthorsGetPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorApiGetcmsv3blogsauthorsGetPageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAt** | **optional.Time**| Only return Blog Authors created at exactly the specified time. | 
 **createdAfter** | **optional.Time**| Only return Blog Authors created after the specified time. | 
 **createdBefore** | **optional.Time**| Only return Blog Authors created before the specified time. | 
 **updatedAt** | **optional.Time**| Only return Blog Authors last updated at exactly the specified time. | 
 **updatedAfter** | **optional.Time**| Only return Blog Authors last updated after the specified time. | 
 **updatedBefore** | **optional.Time**| Only return Blog Authors last updated before the specified time. | 
 **sort** | [**optional.Interface of []string**](string.md)| Specifies which fields to use for sorting results. Valid fields are &#x60;name&#x60;, &#x60;createdAt&#x60;, &#x60;updatedAt&#x60;, &#x60;createdBy&#x60;, &#x60;updatedBy&#x60;. &#x60;createdAt&#x60; will be used by default. | 
 **after** | **optional.String**| The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **optional.Int32**| The maximum number of results to return. Default is 100. | 
 **archived** | **optional.Bool**| Specifies whether to return archived Blog Authors. Defaults to &#x60;false&#x60;. | 

### Return type

[**CollectionResponseWithTotalBlogAuthorForwardPaging**](CollectionResponseWithTotalBlogAuthorForwardPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcmsv3blogsauthorsobjectIdGetById**
> BlogAuthor Getcmsv3blogsauthorsobjectIdGetById(ctx, objectId, optional)
Retrieve a Blog Author

Retrieve the Blog Author object identified by the id in the path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**| The Blog Author id. | 
 **optional** | ***AuthorApiGetcmsv3blogsauthorsobjectIdGetByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorApiGetcmsv3blogsauthorsobjectIdGetByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **optional.Bool**| Specifies whether to return archived Blog Authors. Defaults to &#x60;false&#x60;. | 

### Return type

[**BlogAuthor**](BlogAuthor.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patchcmsv3blogsauthorsobjectIdUpdate**
> BlogAuthor Patchcmsv3blogsauthorsobjectIdUpdate(ctx, body, objectId, optional)
Update a Blog Author

Sparse updates a single Blog Author object identified by the id in the path. All the column values need not be specified. Only the that need to be modified can be specified. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BlogAuthor**](BlogAuthor.md)| The JSON representation of the updated Blog Author. | 
  **objectId** | **string**| The Blog Author id. | 
 **optional** | ***AuthorApiPatchcmsv3blogsauthorsobjectIdUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorApiPatchcmsv3blogsauthorsobjectIdUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **archived** | **optional.**| Specifies whether to update archived Blog Authors. Defaults to &#x60;false&#x60;. | 

### Return type

[**BlogAuthor**](BlogAuthor.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3blogsauthorsCreate**
> BlogAuthor Postcmsv3blogsauthorsCreate(ctx, body)
Create a new Blog Author

Create a new Blog Author.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BlogAuthor**](BlogAuthor.md)| The JSON representation of a new Blog Author. | 

### Return type

[**BlogAuthor**](BlogAuthor.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3blogsauthorsbatcharchiveArchiveBatch**
> Postcmsv3blogsauthorsbatcharchiveArchiveBatch(ctx, body)
Archive a batch of Blog Authors

Delete the Blog Author objects identified in the request body. Note: This is not the same as the in-app `archive` function.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputString**](BatchInputString.md)| The JSON array of Blog Author ids. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3blogsauthorsbatchcreateCreateBatch**
> BatchResponseBlogAuthor Postcmsv3blogsauthorsbatchcreateCreateBatch(ctx, body)
Create a batch of Blog Authors

Create the Blog Author objects detailed in the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputBlogAuthor**](BatchInputBlogAuthor.md)| The JSON array of new Blog Authors to create. | 

### Return type

[**BatchResponseBlogAuthor**](BatchResponseBlogAuthor.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3blogsauthorsbatchreadReadBatch**
> BatchResponseBlogAuthor Postcmsv3blogsauthorsbatchreadReadBatch(ctx, body, optional)
Retrieve a batch of Blog Authors

Retrieve the Blog Author objects identified in the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputString**](BatchInputString.md)| The JSON array of Blog Author ids. | 
 **optional** | ***AuthorApiPostcmsv3blogsauthorsbatchreadReadBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorApiPostcmsv3blogsauthorsbatchreadReadBatchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **optional.**| Specifies whether to return archived Blog Authors. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseBlogAuthor**](BatchResponseBlogAuthor.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3blogsauthorsbatchupdateUpdateBatch**
> BatchResponseBlogAuthor Postcmsv3blogsauthorsbatchupdateUpdateBatch(ctx, body, optional)
Update a batch of Blog Authors

Update the Blog Author objects identified in the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputJsonNode**](BatchInputJsonNode.md)| A JSON array of the JSON representations of the updated Blog Authors. | 
 **optional** | ***AuthorApiPostcmsv3blogsauthorsbatchupdateUpdateBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorApiPostcmsv3blogsauthorsbatchupdateUpdateBatchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **optional.**| Specifies whether to update archived Blog Authors. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseBlogAuthor**](BatchResponseBlogAuthor.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

