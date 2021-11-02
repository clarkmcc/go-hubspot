# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletecmsv3blogspostsobjectIdArchive**](BlogPostApi.md#Deletecmsv3blogspostsobjectIdArchive) | **Delete** /cms/v3/blogs/posts/{objectId} | Delete a Blog Post
[**Getcmsv3blogspostsGetPage**](BlogPostApi.md#Getcmsv3blogspostsGetPage) | **Get** /cms/v3/blogs/posts | Get all Blog Posts
[**Getcmsv3blogspostsobjectIdGetById**](BlogPostApi.md#Getcmsv3blogspostsobjectIdGetById) | **Get** /cms/v3/blogs/posts/{objectId} | Retrieve a Blog Post
[**Getcmsv3blogspostsobjectIddraftGetDraftById**](BlogPostApi.md#Getcmsv3blogspostsobjectIddraftGetDraftById) | **Get** /cms/v3/blogs/posts/{objectId}/draft | Retrieve the full draft version of the Blog Post
[**Getcmsv3blogspostsobjectIdrevisionsGetPreviousVersions**](BlogPostApi.md#Getcmsv3blogspostsobjectIdrevisionsGetPreviousVersions) | **Get** /cms/v3/blogs/posts/{objectId}/revisions | Retrieves all the previous versions of a blog post
[**Getcmsv3blogspostsobjectIdrevisionsrevisionIdGetPreviousVersion**](BlogPostApi.md#Getcmsv3blogspostsobjectIdrevisionsrevisionIdGetPreviousVersion) | **Get** /cms/v3/blogs/posts/{objectId}/revisions/{revisionId} | Retrieves a previous version of a blog post
[**Patchcmsv3blogspostsobjectIdUpdate**](BlogPostApi.md#Patchcmsv3blogspostsobjectIdUpdate) | **Patch** /cms/v3/blogs/posts/{objectId} | Update a Blog Post
[**Patchcmsv3blogspostsobjectIddraftUpdateDraft**](BlogPostApi.md#Patchcmsv3blogspostsobjectIddraftUpdateDraft) | **Patch** /cms/v3/blogs/posts/{objectId}/draft | Update a Blog Post draft
[**Postcmsv3blogspostsCreate**](BlogPostApi.md#Postcmsv3blogspostsCreate) | **Post** /cms/v3/blogs/posts | Create a new Blog Post
[**Postcmsv3blogspostsbatcharchiveArchiveBatch**](BlogPostApi.md#Postcmsv3blogspostsbatcharchiveArchiveBatch) | **Post** /cms/v3/blogs/posts/batch/archive | Archive a batch of Blog Posts
[**Postcmsv3blogspostsbatchcreateCreateBatch**](BlogPostApi.md#Postcmsv3blogspostsbatchcreateCreateBatch) | **Post** /cms/v3/blogs/posts/batch/create | Create a batch of Blog Posts
[**Postcmsv3blogspostsbatchreadReadBatch**](BlogPostApi.md#Postcmsv3blogspostsbatchreadReadBatch) | **Post** /cms/v3/blogs/posts/batch/read | Retrieve a batch of Blog Posts
[**Postcmsv3blogspostsbatchupdateUpdateBatch**](BlogPostApi.md#Postcmsv3blogspostsbatchupdateUpdateBatch) | **Post** /cms/v3/blogs/posts/batch/update | Update a batch of Blog Posts
[**Postcmsv3blogspostscloneClone**](BlogPostApi.md#Postcmsv3blogspostscloneClone) | **Post** /cms/v3/blogs/posts/clone | Clone a Blog Post
[**Postcmsv3blogspostsobjectIddraftpushLivePushLive**](BlogPostApi.md#Postcmsv3blogspostsobjectIddraftpushLivePushLive) | **Post** /cms/v3/blogs/posts/{objectId}/draft/push-live | Push Blog Post draft edits live
[**Postcmsv3blogspostsobjectIddraftresetResetDraft**](BlogPostApi.md#Postcmsv3blogspostsobjectIddraftresetResetDraft) | **Post** /cms/v3/blogs/posts/{objectId}/draft/reset | Reset the Blog Post draft to the live version
[**Postcmsv3blogspostsobjectIdrevisionsrevisionIdrestoreRestorePreviousVersion**](BlogPostApi.md#Postcmsv3blogspostsobjectIdrevisionsrevisionIdrestoreRestorePreviousVersion) | **Post** /cms/v3/blogs/posts/{objectId}/revisions/{revisionId}/restore | Restore a previous version of a blog post
[**Postcmsv3blogspostsobjectIdrevisionsrevisionIdrestoreToDraftRestorePreviousVersionToDraft**](BlogPostApi.md#Postcmsv3blogspostsobjectIdrevisionsrevisionIdrestoreToDraftRestorePreviousVersionToDraft) | **Post** /cms/v3/blogs/posts/{objectId}/revisions/{revisionId}/restore-to-draft | Restore a previous version of a blog post, to the draft version of the blog post
[**Postcmsv3blogspostsscheduleSchedule**](BlogPostApi.md#Postcmsv3blogspostsscheduleSchedule) | **Post** /cms/v3/blogs/posts/schedule | Schedule a Blog Post to be Published

# **Deletecmsv3blogspostsobjectIdArchive**
> Deletecmsv3blogspostsobjectIdArchive(ctx, objectId, optional)
Delete a Blog Post

Delete the Blog Post object identified by the id in the path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**| The Blog Post id. | 
 **optional** | ***BlogPostApiDeletecmsv3blogspostsobjectIdArchiveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlogPostApiDeletecmsv3blogspostsobjectIdArchiveOpts struct
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

# **Getcmsv3blogspostsGetPage**
> CollectionResponseWithTotalBlogPostForwardPaging Getcmsv3blogspostsGetPage(ctx, optional)
Get all Blog Posts

Get the list of blog posts. Supports paging and filtering. This method would be useful for an integration that examined these models and used an external service to suggest edits. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BlogPostApiGetcmsv3blogspostsGetPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlogPostApiGetcmsv3blogspostsGetPageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAt** | **optional.Time**| Only return Blog Posts created at exactly the specified time. | 
 **createdAfter** | **optional.Time**| Only return Blog Posts created after the specified time. | 
 **createdBefore** | **optional.Time**| Only return Blog Posts created before the specified time. | 
 **updatedAt** | **optional.Time**| Only return Blog Posts last updated at exactly the specified time. | 
 **updatedAfter** | **optional.Time**| Only return Blog Posts last updated after the specified time. | 
 **updatedBefore** | **optional.Time**| Only return Blog Posts last updated before the specified time. | 
 **sort** | [**optional.Interface of []string**](string.md)| Specifies which fields to use for sorting results. Valid fields are &#x60;name&#x60;, &#x60;createdAt&#x60;, &#x60;updatedAt&#x60;, &#x60;createdBy&#x60;, &#x60;updatedBy&#x60;. &#x60;createdAt&#x60; will be used by default. | 
 **after** | **optional.String**| The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **optional.Int32**| The maximum number of results to return. Default is 100. | 
 **archived** | **optional.Bool**| Specifies whether to return archived Blog Posts. Defaults to &#x60;false&#x60;. | 

### Return type

[**CollectionResponseWithTotalBlogPostForwardPaging**](CollectionResponseWithTotalBlogPostForwardPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcmsv3blogspostsobjectIdGetById**
> BlogPost Getcmsv3blogspostsobjectIdGetById(ctx, objectId, optional)
Retrieve a Blog Post

Retrieve the Blog Post object identified by the id in the path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**| The Blog Post id. | 
 **optional** | ***BlogPostApiGetcmsv3blogspostsobjectIdGetByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlogPostApiGetcmsv3blogspostsobjectIdGetByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **optional.Bool**| Specifies whether to return archived Blog Posts. Defaults to &#x60;false&#x60;. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcmsv3blogspostsobjectIddraftGetDraftById**
> BlogPost Getcmsv3blogspostsobjectIddraftGetDraftById(ctx, objectId)
Retrieve the full draft version of the Blog Post

Retrieve the full draft version of the Blog Post.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**| The Blog Post id. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcmsv3blogspostsobjectIdrevisionsGetPreviousVersions**
> CollectionResponseWithTotalVersionBlogPost Getcmsv3blogspostsobjectIdrevisionsGetPreviousVersions(ctx, objectId, optional)
Retrieves all the previous versions of a blog post

Retrieves all the previous versions of a blog post.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**| The Blog Post id. | 
 **optional** | ***BlogPostApiGetcmsv3blogspostsobjectIdrevisionsGetPreviousVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlogPostApiGetcmsv3blogspostsobjectIdrevisionsGetPreviousVersionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **optional.String**| The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **before** | **optional.String**|  | 
 **limit** | **optional.Int32**| The maximum number of results to return. Default is 100. | 

### Return type

[**CollectionResponseWithTotalVersionBlogPost**](CollectionResponseWithTotalVersionBlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcmsv3blogspostsobjectIdrevisionsrevisionIdGetPreviousVersion**
> VersionBlogPost Getcmsv3blogspostsobjectIdrevisionsrevisionIdGetPreviousVersion(ctx, objectId, revisionId)
Retrieves a previous version of a blog post

Retrieves a previous version of a blog post.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**| The Blog Post id. | 
  **revisionId** | **string**| The Blog Post version id. | 

### Return type

[**VersionBlogPost**](VersionBlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patchcmsv3blogspostsobjectIdUpdate**
> BlogPost Patchcmsv3blogspostsobjectIdUpdate(ctx, body, objectId, optional)
Update a Blog Post

Sparse updates a single Blog Post object identified by the id in the path. All the column values need not be specified. Only the that need to be modified can be specified. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BlogPost**](BlogPost.md)| The JSON representation of the updated Blog Post. | 
  **objectId** | **string**| The Blog Post id. | 
 **optional** | ***BlogPostApiPatchcmsv3blogspostsobjectIdUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlogPostApiPatchcmsv3blogspostsobjectIdUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **archived** | **optional.**| Specifies whether to update archived Blog Posts. Defaults to &#x60;false&#x60;. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patchcmsv3blogspostsobjectIddraftUpdateDraft**
> BlogPost Patchcmsv3blogspostsobjectIddraftUpdateDraft(ctx, body, objectId)
Update a Blog Post draft

Sparse updates the draft version of a single Blog Post object identified by the id in the path. All the column values need not be specified. Only the that need to be modified can be specified. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BlogPost**](BlogPost.md)| The JSON representation of the updated Blog Post to be applied to the draft. | 
  **objectId** | **string**| The Blog Post id. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3blogspostsCreate**
> BlogPost Postcmsv3blogspostsCreate(ctx, body)
Create a new Blog Post

Create a new Blog Post.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BlogPost**](BlogPost.md)| The JSON representation of a new Blog Post. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3blogspostsbatcharchiveArchiveBatch**
> Postcmsv3blogspostsbatcharchiveArchiveBatch(ctx, body)
Archive a batch of Blog Posts

Delete the Blog Post objects identified in the request body. Note: This is not the same as the in-app `archive` function. To perform an in-app `archive` send an normal update with the `archived` field set to true.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputString**](BatchInputString.md)| The JSON array of Blog Post ids. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3blogspostsbatchcreateCreateBatch**
> BatchResponseBlogPost Postcmsv3blogspostsbatchcreateCreateBatch(ctx, body)
Create a batch of Blog Posts

Create the Blog Post objects detailed in the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputBlogPost**](BatchInputBlogPost.md)| The JSON array of new Blog Posts to create. | 

### Return type

[**BatchResponseBlogPost**](BatchResponseBlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3blogspostsbatchreadReadBatch**
> BatchResponseBlogPost Postcmsv3blogspostsbatchreadReadBatch(ctx, body, optional)
Retrieve a batch of Blog Posts

Retrieve the Blog Post objects identified in the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputString**](BatchInputString.md)| The JSON array of Blog Post ids. | 
 **optional** | ***BlogPostApiPostcmsv3blogspostsbatchreadReadBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlogPostApiPostcmsv3blogspostsbatchreadReadBatchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **optional.**| Specifies whether to return archived Blog Posts. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseBlogPost**](BatchResponseBlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3blogspostsbatchupdateUpdateBatch**
> BatchResponseBlogPost Postcmsv3blogspostsbatchupdateUpdateBatch(ctx, body, optional)
Update a batch of Blog Posts

Update the Blog Post objects identified in the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputJsonNode**](BatchInputJsonNode.md)|  | 
 **optional** | ***BlogPostApiPostcmsv3blogspostsbatchupdateUpdateBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlogPostApiPostcmsv3blogspostsbatchupdateUpdateBatchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **optional.**| Whether to return only results that have been archived. | 

### Return type

[**BatchResponseBlogPost**](BatchResponseBlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3blogspostscloneClone**
> BlogPost Postcmsv3blogspostscloneClone(ctx, body)
Clone a Blog Post

Clone a Blog.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ContentCloneRequestVNext**](ContentCloneRequestVNext.md)| The JSON representation of the ContentCloneRequest object. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3blogspostsobjectIddraftpushLivePushLive**
> Postcmsv3blogspostsobjectIddraftpushLivePushLive(ctx, objectId)
Push Blog Post draft edits live

Take any changes from the draft version of the Blog Post and apply them to the live version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**| The id of the Blog Post for which it&#x27;s draft will be pushed live. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3blogspostsobjectIddraftresetResetDraft**
> Postcmsv3blogspostsobjectIddraftresetResetDraft(ctx, objectId)
Reset the Blog Post draft to the live version

Discards any edits and resets the draft to the live version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**| The id of the Blog Post for which it&#x27;s draft will be reset. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3blogspostsobjectIdrevisionsrevisionIdrestoreRestorePreviousVersion**
> BlogPost Postcmsv3blogspostsobjectIdrevisionsrevisionIdrestoreRestorePreviousVersion(ctx, objectId, revisionId)
Restore a previous version of a blog post

Takes a specified version of a blog post and restores it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**| The Blog Post id. | 
  **revisionId** | **string**| The Blog Post version id to restore. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3blogspostsobjectIdrevisionsrevisionIdrestoreToDraftRestorePreviousVersionToDraft**
> BlogPost Postcmsv3blogspostsobjectIdrevisionsrevisionIdrestoreToDraftRestorePreviousVersionToDraft(ctx, objectId, revisionId)
Restore a previous version of a blog post, to the draft version of the blog post

Takes a specified version of a blog post, sets it as the new draft version of the blog post.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**| The Blog Post id. | 
  **revisionId** | **int64**| The Blog Post version id to restore. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3blogspostsscheduleSchedule**
> Postcmsv3blogspostsscheduleSchedule(ctx, body)
Schedule a Blog Post to be Published

Schedule a Blog Post to be Published.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ContentScheduleRequestVNext**](ContentScheduleRequestVNext.md)| The JSON representation of the ContentCloneRequestVNext object. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

