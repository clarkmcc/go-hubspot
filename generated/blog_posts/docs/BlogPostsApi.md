# \BlogPostsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCmsV3BlogsPostsObjectIdArchive**](BlogPostsApi.md#DeleteCmsV3BlogsPostsObjectIdArchive) | **Delete** /cms/v3/blogs/posts/{objectId} | Delete a Blog Post
[**GetCmsV3BlogsPostsGetPage**](BlogPostsApi.md#GetCmsV3BlogsPostsGetPage) | **Get** /cms/v3/blogs/posts | Get all Blog Posts
[**GetCmsV3BlogsPostsObjectIdDraftGetDraftById**](BlogPostsApi.md#GetCmsV3BlogsPostsObjectIdDraftGetDraftById) | **Get** /cms/v3/blogs/posts/{objectId}/draft | Retrieve the full draft version of the Blog Post
[**GetCmsV3BlogsPostsObjectIdGetById**](BlogPostsApi.md#GetCmsV3BlogsPostsObjectIdGetById) | **Get** /cms/v3/blogs/posts/{objectId} | Retrieve a Blog Post
[**GetCmsV3BlogsPostsObjectIdRevisionsGetPreviousVersions**](BlogPostsApi.md#GetCmsV3BlogsPostsObjectIdRevisionsGetPreviousVersions) | **Get** /cms/v3/blogs/posts/{objectId}/revisions | Retrieves all the previous versions of a blog post
[**GetCmsV3BlogsPostsObjectIdRevisionsRevisionIdGetPreviousVersion**](BlogPostsApi.md#GetCmsV3BlogsPostsObjectIdRevisionsRevisionIdGetPreviousVersion) | **Get** /cms/v3/blogs/posts/{objectId}/revisions/{revisionId} | Retrieves a previous version of a blog post
[**PatchCmsV3BlogsPostsObjectIdDraftUpdateDraft**](BlogPostsApi.md#PatchCmsV3BlogsPostsObjectIdDraftUpdateDraft) | **Patch** /cms/v3/blogs/posts/{objectId}/draft | Update a Blog Post draft
[**PatchCmsV3BlogsPostsObjectIdUpdate**](BlogPostsApi.md#PatchCmsV3BlogsPostsObjectIdUpdate) | **Patch** /cms/v3/blogs/posts/{objectId} | Update a Blog Post
[**PostCmsV3BlogsPostsBatchArchiveArchiveBatch**](BlogPostsApi.md#PostCmsV3BlogsPostsBatchArchiveArchiveBatch) | **Post** /cms/v3/blogs/posts/batch/archive | Delete a batch of Blog Posts
[**PostCmsV3BlogsPostsBatchCreateCreateBatch**](BlogPostsApi.md#PostCmsV3BlogsPostsBatchCreateCreateBatch) | **Post** /cms/v3/blogs/posts/batch/create | Create a batch of Blog Posts
[**PostCmsV3BlogsPostsBatchReadReadBatch**](BlogPostsApi.md#PostCmsV3BlogsPostsBatchReadReadBatch) | **Post** /cms/v3/blogs/posts/batch/read | Retrieve a batch of Blog Posts
[**PostCmsV3BlogsPostsBatchUpdateUpdateBatch**](BlogPostsApi.md#PostCmsV3BlogsPostsBatchUpdateUpdateBatch) | **Post** /cms/v3/blogs/posts/batch/update | Update a batch of Blog Posts
[**PostCmsV3BlogsPostsCloneClone**](BlogPostsApi.md#PostCmsV3BlogsPostsCloneClone) | **Post** /cms/v3/blogs/posts/clone | Clone a Blog Post
[**PostCmsV3BlogsPostsCreate**](BlogPostsApi.md#PostCmsV3BlogsPostsCreate) | **Post** /cms/v3/blogs/posts | Create a new Blog Post
[**PostCmsV3BlogsPostsMultiLanguageAttachToLangGroupAttachToLangGroup**](BlogPostsApi.md#PostCmsV3BlogsPostsMultiLanguageAttachToLangGroupAttachToLangGroup) | **Post** /cms/v3/blogs/posts/multi-language/attach-to-lang-group | Attach a Blog Post to a multi-language group
[**PostCmsV3BlogsPostsMultiLanguageCreateLanguageVariationCreateLangVariation**](BlogPostsApi.md#PostCmsV3BlogsPostsMultiLanguageCreateLanguageVariationCreateLangVariation) | **Post** /cms/v3/blogs/posts/multi-language/create-language-variation | Create a new language variation
[**PostCmsV3BlogsPostsMultiLanguageDetachFromLangGroupDetachFromLangGroup**](BlogPostsApi.md#PostCmsV3BlogsPostsMultiLanguageDetachFromLangGroupDetachFromLangGroup) | **Post** /cms/v3/blogs/posts/multi-language/detach-from-lang-group | Detach a Blog Post from a multi-language group
[**PostCmsV3BlogsPostsMultiLanguageUpdateLanguagesUpdateLangs**](BlogPostsApi.md#PostCmsV3BlogsPostsMultiLanguageUpdateLanguagesUpdateLangs) | **Post** /cms/v3/blogs/posts/multi-language/update-languages | Update languages of multi-language group
[**PostCmsV3BlogsPostsObjectIdDraftPushLivePushLive**](BlogPostsApi.md#PostCmsV3BlogsPostsObjectIdDraftPushLivePushLive) | **Post** /cms/v3/blogs/posts/{objectId}/draft/push-live | Push Blog Post draft edits live
[**PostCmsV3BlogsPostsObjectIdDraftResetResetDraft**](BlogPostsApi.md#PostCmsV3BlogsPostsObjectIdDraftResetResetDraft) | **Post** /cms/v3/blogs/posts/{objectId}/draft/reset | Reset the Blog Post draft to the live version
[**PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion**](BlogPostsApi.md#PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion) | **Post** /cms/v3/blogs/posts/{objectId}/revisions/{revisionId}/restore | Restore a previous version of a blog post
[**PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft**](BlogPostsApi.md#PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft) | **Post** /cms/v3/blogs/posts/{objectId}/revisions/{revisionId}/restore-to-draft | Restore a previous version of a blog post, to the draft version of the blog post
[**PostCmsV3BlogsPostsScheduleSchedule**](BlogPostsApi.md#PostCmsV3BlogsPostsScheduleSchedule) | **Post** /cms/v3/blogs/posts/schedule | Schedule a Blog Post to be Published
[**PutCmsV3BlogsPostsMultiLanguageSetNewLangPrimarySetLangPrimary**](BlogPostsApi.md#PutCmsV3BlogsPostsMultiLanguageSetNewLangPrimarySetLangPrimary) | **Put** /cms/v3/blogs/posts/multi-language/set-new-lang-primary | Set a new primary language



## DeleteCmsV3BlogsPostsObjectIdArchive

> DeleteCmsV3BlogsPostsObjectIdArchive(ctx, objectId).Archived(archived).Execute()

Delete a Blog Post



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    objectId := "objectId_example" // string | The Blog Post id.
    archived := true // bool | Whether to return only results that have been archived. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.DeleteCmsV3BlogsPostsObjectIdArchive(context.Background(), objectId).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.DeleteCmsV3BlogsPostsObjectIdArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCmsV3BlogsPostsObjectIdArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Whether to return only results that have been archived. | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3BlogsPostsGetPage

> CollectionResponseWithTotalBlogPostForwardPaging GetCmsV3BlogsPostsGetPage(ctx).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Execute()

Get all Blog Posts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    createdAt := time.Now() // time.Time | Only return Blog Posts created at exactly the specified time. (optional)
    createdAfter := time.Now() // time.Time | Only return Blog Posts created after the specified time. (optional)
    createdBefore := time.Now() // time.Time | Only return Blog Posts created before the specified time. (optional)
    updatedAt := time.Now() // time.Time | Only return Blog Posts last updated at exactly the specified time. (optional)
    updatedAfter := time.Now() // time.Time | Only return Blog Posts last updated after the specified time. (optional)
    updatedBefore := time.Now() // time.Time | Only return Blog Posts last updated before the specified time. (optional)
    sort := []string{"Inner_example"} // []string | Specifies which fields to use for sorting results. Valid fields are `name`, `createdAt`, `updatedAt`, `createdBy`, `updatedBy`. `createdAt` will be used by default. (optional)
    after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
    limit := int32(56) // int32 | The maximum number of results to return. Default is 20. (optional)
    archived := true // bool | Specifies whether to return deleted Blog Posts. Defaults to `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.GetCmsV3BlogsPostsGetPage(context.Background()).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.GetCmsV3BlogsPostsGetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3BlogsPostsGetPage`: CollectionResponseWithTotalBlogPostForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.GetCmsV3BlogsPostsGetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsPostsGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAt** | **time.Time** | Only return Blog Posts created at exactly the specified time. | 
 **createdAfter** | **time.Time** | Only return Blog Posts created after the specified time. | 
 **createdBefore** | **time.Time** | Only return Blog Posts created before the specified time. | 
 **updatedAt** | **time.Time** | Only return Blog Posts last updated at exactly the specified time. | 
 **updatedAfter** | **time.Time** | Only return Blog Posts last updated after the specified time. | 
 **updatedBefore** | **time.Time** | Only return Blog Posts last updated before the specified time. | 
 **sort** | **[]string** | Specifies which fields to use for sorting results. Valid fields are &#x60;name&#x60;, &#x60;createdAt&#x60;, &#x60;updatedAt&#x60;, &#x60;createdBy&#x60;, &#x60;updatedBy&#x60;. &#x60;createdAt&#x60; will be used by default. | 
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is 20. | 
 **archived** | **bool** | Specifies whether to return deleted Blog Posts. Defaults to &#x60;false&#x60;. | 

### Return type

[**CollectionResponseWithTotalBlogPostForwardPaging**](CollectionResponseWithTotalBlogPostForwardPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3BlogsPostsObjectIdDraftGetDraftById

> BlogPost GetCmsV3BlogsPostsObjectIdDraftGetDraftById(ctx, objectId).Execute()

Retrieve the full draft version of the Blog Post



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    objectId := "objectId_example" // string | The Blog Post id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.GetCmsV3BlogsPostsObjectIdDraftGetDraftById(context.Background(), objectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.GetCmsV3BlogsPostsObjectIdDraftGetDraftById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3BlogsPostsObjectIdDraftGetDraftById`: BlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.GetCmsV3BlogsPostsObjectIdDraftGetDraftById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsPostsObjectIdDraftGetDraftByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3BlogsPostsObjectIdGetById

> BlogPost GetCmsV3BlogsPostsObjectIdGetById(ctx, objectId).Archived(archived).Execute()

Retrieve a Blog Post



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    objectId := "objectId_example" // string | The Blog Post id.
    archived := true // bool | Specifies whether to return deleted Blog Posts. Defaults to `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.GetCmsV3BlogsPostsObjectIdGetById(context.Background(), objectId).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.GetCmsV3BlogsPostsObjectIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3BlogsPostsObjectIdGetById`: BlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.GetCmsV3BlogsPostsObjectIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsPostsObjectIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Specifies whether to return deleted Blog Posts. Defaults to &#x60;false&#x60;. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3BlogsPostsObjectIdRevisionsGetPreviousVersions

> CollectionResponseWithTotalVersionBlogPost GetCmsV3BlogsPostsObjectIdRevisionsGetPreviousVersions(ctx, objectId).After(after).Before(before).Limit(limit).Execute()

Retrieves all the previous versions of a blog post



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    objectId := "objectId_example" // string | The Blog Post id.
    after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
    before := "before_example" // string |  (optional)
    limit := int32(56) // int32 | The maximum number of results to return. Default is 100. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.GetCmsV3BlogsPostsObjectIdRevisionsGetPreviousVersions(context.Background(), objectId).After(after).Before(before).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.GetCmsV3BlogsPostsObjectIdRevisionsGetPreviousVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3BlogsPostsObjectIdRevisionsGetPreviousVersions`: CollectionResponseWithTotalVersionBlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.GetCmsV3BlogsPostsObjectIdRevisionsGetPreviousVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsPostsObjectIdRevisionsGetPreviousVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **before** | **string** |  | 
 **limit** | **int32** | The maximum number of results to return. Default is 100. | 

### Return type

[**CollectionResponseWithTotalVersionBlogPost**](CollectionResponseWithTotalVersionBlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3BlogsPostsObjectIdRevisionsRevisionIdGetPreviousVersion

> VersionBlogPost GetCmsV3BlogsPostsObjectIdRevisionsRevisionIdGetPreviousVersion(ctx, objectId, revisionId).Execute()

Retrieves a previous version of a blog post



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    objectId := "objectId_example" // string | The Blog Post id.
    revisionId := "revisionId_example" // string | The Blog Post version id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.GetCmsV3BlogsPostsObjectIdRevisionsRevisionIdGetPreviousVersion(context.Background(), objectId, revisionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.GetCmsV3BlogsPostsObjectIdRevisionsRevisionIdGetPreviousVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3BlogsPostsObjectIdRevisionsRevisionIdGetPreviousVersion`: VersionBlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.GetCmsV3BlogsPostsObjectIdRevisionsRevisionIdGetPreviousVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 
**revisionId** | **string** | The Blog Post version id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsPostsObjectIdRevisionsRevisionIdGetPreviousVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VersionBlogPost**](VersionBlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCmsV3BlogsPostsObjectIdDraftUpdateDraft

> BlogPost PatchCmsV3BlogsPostsObjectIdDraftUpdateDraft(ctx, objectId).BlogPost(blogPost).Execute()

Update a Blog Post draft



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    objectId := "objectId_example" // string | The Blog Post id.
    blogPost := *openapiclient.NewBlogPost("Id_example", "Slug_example", "ContentGroupId_example", "Campaign_example", int32(123), "State_example", "Name_example", "MabExperimentId_example", false, "AuthorName_example", "AbTestId_example", "CreatedById_example", "UpdatedById_example", "Domain_example", "AbStatus_example", "FolderId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "Language_example", "TranslatedFromId_example", map[string]ContentLanguageVariation{"key": *openapiclient.NewContentLanguageVariation(int64(123), "Name_example", "Slug_example", "State_example", "AuthorName_example", "Password_example", false, []map[string]interface{}{map[string]interface{}(123)}, "Campaign_example", false, time.Now(), time.Now(), time.Now())}, int32(123), "DynamicPageDataSourceId_example", "BlogAuthorId_example", []int64{int64(123)}, "HtmlTitle_example", false, false, "PostBody_example", "PostSummary_example", "RssBody_example", "RssSummary_example", false, false, int64(123), "PageExpiryRedirectUrl_example", int64(123), false, false, false, false, "FeaturedImage_example", "FeaturedImageAltText_example", "LinkRelCanonicalUrl_example", "ContentTypeCategory_example", []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}}, "MetaDescription_example", "HeadHtml_example", "FooterHtml_example", false, false, []map[string]interface{}{map[string]interface{}(123)}, map[string]LayoutSection{"key": *openapiclient.NewLayoutSection(int32(123), int32(123), "Name_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": *openapiclient.NewLayoutSection(int32(123), int32(123), "Name_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": }}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData(*openapiclient.NewStyles("VerticalAlignment_example", *openapiclient.NewRGBAColor(int32(123), int32(123), int32(123), float32(123)), *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), *openapiclient.NewGradient(*openapiclient.NewSideOrCorner("VerticalSide_example", "HorizontalSide_example"), *openapiclient.NewAngle(float32(123), "Units_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop(*openapiclient.NewRGBAColor(int32(123), int32(123), int32(123), float32(123)))}), int32(123), false, "FlexboxPositioning_example"), "CssClass_example")}, []openapiclient.LayoutSection{}, "CssClass_example", "CssStyle_example", "CssId_example", *openapiclient.NewStyles("VerticalAlignment_example", , *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), *openapiclient.NewGradient(*openapiclient.NewSideOrCorner("VerticalSide_example", "HorizontalSide_example"), *openapiclient.NewAngle(float32(123), "Units_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop()}), int32(123), false, "FlexboxPositioning_example"))}}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData(, "CssClass_example")}, []openapiclient.LayoutSection{}, "CssClass_example", "CssStyle_example", "CssId_example", )}, map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "Url_example", "Password_example", "CurrentState_example", time.Now(), time.Now(), time.Now(), time.Now()) // BlogPost | The JSON representation of the updated Blog Post to be applied to the draft.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.PatchCmsV3BlogsPostsObjectIdDraftUpdateDraft(context.Background(), objectId).BlogPost(blogPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.PatchCmsV3BlogsPostsObjectIdDraftUpdateDraft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCmsV3BlogsPostsObjectIdDraftUpdateDraft`: BlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.PatchCmsV3BlogsPostsObjectIdDraftUpdateDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCmsV3BlogsPostsObjectIdDraftUpdateDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blogPost** | [**BlogPost**](BlogPost.md) | The JSON representation of the updated Blog Post to be applied to the draft. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCmsV3BlogsPostsObjectIdUpdate

> BlogPost PatchCmsV3BlogsPostsObjectIdUpdate(ctx, objectId).BlogPost(blogPost).Archived(archived).Execute()

Update a Blog Post



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    objectId := "objectId_example" // string | The Blog Post id.
    blogPost := *openapiclient.NewBlogPost("Id_example", "Slug_example", "ContentGroupId_example", "Campaign_example", int32(123), "State_example", "Name_example", "MabExperimentId_example", false, "AuthorName_example", "AbTestId_example", "CreatedById_example", "UpdatedById_example", "Domain_example", "AbStatus_example", "FolderId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "Language_example", "TranslatedFromId_example", map[string]ContentLanguageVariation{"key": *openapiclient.NewContentLanguageVariation(int64(123), "Name_example", "Slug_example", "State_example", "AuthorName_example", "Password_example", false, []map[string]interface{}{map[string]interface{}(123)}, "Campaign_example", false, time.Now(), time.Now(), time.Now())}, int32(123), "DynamicPageDataSourceId_example", "BlogAuthorId_example", []int64{int64(123)}, "HtmlTitle_example", false, false, "PostBody_example", "PostSummary_example", "RssBody_example", "RssSummary_example", false, false, int64(123), "PageExpiryRedirectUrl_example", int64(123), false, false, false, false, "FeaturedImage_example", "FeaturedImageAltText_example", "LinkRelCanonicalUrl_example", "ContentTypeCategory_example", []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}}, "MetaDescription_example", "HeadHtml_example", "FooterHtml_example", false, false, []map[string]interface{}{map[string]interface{}(123)}, map[string]LayoutSection{"key": *openapiclient.NewLayoutSection(int32(123), int32(123), "Name_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": *openapiclient.NewLayoutSection(int32(123), int32(123), "Name_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": }}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData(*openapiclient.NewStyles("VerticalAlignment_example", *openapiclient.NewRGBAColor(int32(123), int32(123), int32(123), float32(123)), *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), *openapiclient.NewGradient(*openapiclient.NewSideOrCorner("VerticalSide_example", "HorizontalSide_example"), *openapiclient.NewAngle(float32(123), "Units_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop(*openapiclient.NewRGBAColor(int32(123), int32(123), int32(123), float32(123)))}), int32(123), false, "FlexboxPositioning_example"), "CssClass_example")}, []openapiclient.LayoutSection{}, "CssClass_example", "CssStyle_example", "CssId_example", *openapiclient.NewStyles("VerticalAlignment_example", , *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), *openapiclient.NewGradient(*openapiclient.NewSideOrCorner("VerticalSide_example", "HorizontalSide_example"), *openapiclient.NewAngle(float32(123), "Units_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop()}), int32(123), false, "FlexboxPositioning_example"))}}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData(, "CssClass_example")}, []openapiclient.LayoutSection{}, "CssClass_example", "CssStyle_example", "CssId_example", )}, map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "Url_example", "Password_example", "CurrentState_example", time.Now(), time.Now(), time.Now(), time.Now()) // BlogPost | The JSON representation of the updated Blog Post.
    archived := true // bool | Specifies whether to update deleted Blog Posts. Defaults to `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.PatchCmsV3BlogsPostsObjectIdUpdate(context.Background(), objectId).BlogPost(blogPost).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.PatchCmsV3BlogsPostsObjectIdUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCmsV3BlogsPostsObjectIdUpdate`: BlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.PatchCmsV3BlogsPostsObjectIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCmsV3BlogsPostsObjectIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blogPost** | [**BlogPost**](BlogPost.md) | The JSON representation of the updated Blog Post. | 
 **archived** | **bool** | Specifies whether to update deleted Blog Posts. Defaults to &#x60;false&#x60;. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsBatchArchiveArchiveBatch

> PostCmsV3BlogsPostsBatchArchiveArchiveBatch(ctx).BatchInputString(batchInputString).Execute()

Delete a batch of Blog Posts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of Blog Post ids.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.PostCmsV3BlogsPostsBatchArchiveArchiveBatch(context.Background()).BatchInputString(batchInputString).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.PostCmsV3BlogsPostsBatchArchiveArchiveBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsBatchArchiveArchiveBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Blog Post ids. | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsBatchCreateCreateBatch

> BatchResponseBlogPost PostCmsV3BlogsPostsBatchCreateCreateBatch(ctx).BatchInputBlogPost(batchInputBlogPost).Execute()

Create a batch of Blog Posts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    batchInputBlogPost := *openapiclient.NewBatchInputBlogPost([]openapiclient.BlogPost{*openapiclient.NewBlogPost("Id_example", "Slug_example", "ContentGroupId_example", "Campaign_example", int32(123), "State_example", "Name_example", "MabExperimentId_example", false, "AuthorName_example", "AbTestId_example", "CreatedById_example", "UpdatedById_example", "Domain_example", "AbStatus_example", "FolderId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "Language_example", "TranslatedFromId_example", map[string]ContentLanguageVariation{"key": *openapiclient.NewContentLanguageVariation(int64(123), "Name_example", "Slug_example", "State_example", "AuthorName_example", "Password_example", false, []map[string]interface{}{map[string]interface{}(123)}, "Campaign_example", false, time.Now(), time.Now(), time.Now())}, int32(123), "DynamicPageDataSourceId_example", "BlogAuthorId_example", []int64{int64(123)}, "HtmlTitle_example", false, false, "PostBody_example", "PostSummary_example", "RssBody_example", "RssSummary_example", false, false, int64(123), "PageExpiryRedirectUrl_example", int64(123), false, false, false, false, "FeaturedImage_example", "FeaturedImageAltText_example", "LinkRelCanonicalUrl_example", "ContentTypeCategory_example", []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}}, "MetaDescription_example", "HeadHtml_example", "FooterHtml_example", false, false, []map[string]interface{}{map[string]interface{}(123)}, map[string]LayoutSection{"key": *openapiclient.NewLayoutSection(int32(123), int32(123), "Name_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": *openapiclient.NewLayoutSection(int32(123), int32(123), "Name_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": }}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData(*openapiclient.NewStyles("VerticalAlignment_example", *openapiclient.NewRGBAColor(int32(123), int32(123), int32(123), float32(123)), *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), *openapiclient.NewGradient(*openapiclient.NewSideOrCorner("VerticalSide_example", "HorizontalSide_example"), *openapiclient.NewAngle(float32(123), "Units_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop(*openapiclient.NewRGBAColor(int32(123), int32(123), int32(123), float32(123)))}), int32(123), false, "FlexboxPositioning_example"), "CssClass_example")}, []openapiclient.LayoutSection{}, "CssClass_example", "CssStyle_example", "CssId_example", *openapiclient.NewStyles("VerticalAlignment_example", , *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), *openapiclient.NewGradient(*openapiclient.NewSideOrCorner("VerticalSide_example", "HorizontalSide_example"), *openapiclient.NewAngle(float32(123), "Units_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop()}), int32(123), false, "FlexboxPositioning_example"))}}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData(, "CssClass_example")}, []openapiclient.LayoutSection{}, "CssClass_example", "CssStyle_example", "CssId_example", )}, map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "Url_example", "Password_example", "CurrentState_example", time.Now(), time.Now(), time.Now(), time.Now())}) // BatchInputBlogPost | The JSON array of new Blog Posts to create.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.PostCmsV3BlogsPostsBatchCreateCreateBatch(context.Background()).BatchInputBlogPost(batchInputBlogPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.PostCmsV3BlogsPostsBatchCreateCreateBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3BlogsPostsBatchCreateCreateBatch`: BatchResponseBlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.PostCmsV3BlogsPostsBatchCreateCreateBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsBatchCreateCreateBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputBlogPost** | [**BatchInputBlogPost**](BatchInputBlogPost.md) | The JSON array of new Blog Posts to create. | 

### Return type

[**BatchResponseBlogPost**](BatchResponseBlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsBatchReadReadBatch

> BatchResponseBlogPost PostCmsV3BlogsPostsBatchReadReadBatch(ctx).BatchInputString(batchInputString).Archived(archived).Execute()

Retrieve a batch of Blog Posts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of Blog Post ids.
    archived := true // bool | Specifies whether to return deleted Blog Posts. Defaults to `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.PostCmsV3BlogsPostsBatchReadReadBatch(context.Background()).BatchInputString(batchInputString).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.PostCmsV3BlogsPostsBatchReadReadBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3BlogsPostsBatchReadReadBatch`: BatchResponseBlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.PostCmsV3BlogsPostsBatchReadReadBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsBatchReadReadBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Blog Post ids. | 
 **archived** | **bool** | Specifies whether to return deleted Blog Posts. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseBlogPost**](BatchResponseBlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsBatchUpdateUpdateBatch

> BatchResponseBlogPost PostCmsV3BlogsPostsBatchUpdateUpdateBatch(ctx).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()

Update a batch of Blog Posts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    batchInputJsonNode := *openapiclient.NewBatchInputJsonNode([]map[string]interface{}{map[string]interface{}(123)}) // BatchInputJsonNode | A JSON array of the JSON representations of the updated Blog Posts.
    archived := true // bool | Specifies whether to update deleted Blog Posts. Defaults to `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.PostCmsV3BlogsPostsBatchUpdateUpdateBatch(context.Background()).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.PostCmsV3BlogsPostsBatchUpdateUpdateBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3BlogsPostsBatchUpdateUpdateBatch`: BatchResponseBlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.PostCmsV3BlogsPostsBatchUpdateUpdateBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsBatchUpdateUpdateBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputJsonNode** | [**BatchInputJsonNode**](BatchInputJsonNode.md) | A JSON array of the JSON representations of the updated Blog Posts. | 
 **archived** | **bool** | Specifies whether to update deleted Blog Posts. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseBlogPost**](BatchResponseBlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsCloneClone

> BlogPost PostCmsV3BlogsPostsCloneClone(ctx).ContentCloneRequestVNext(contentCloneRequestVNext).Execute()

Clone a Blog Post



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    contentCloneRequestVNext := *openapiclient.NewContentCloneRequestVNext("Id_example") // ContentCloneRequestVNext | The JSON representation of the ContentCloneRequest object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.PostCmsV3BlogsPostsCloneClone(context.Background()).ContentCloneRequestVNext(contentCloneRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.PostCmsV3BlogsPostsCloneClone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3BlogsPostsCloneClone`: BlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.PostCmsV3BlogsPostsCloneClone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsCloneCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentCloneRequestVNext** | [**ContentCloneRequestVNext**](ContentCloneRequestVNext.md) | The JSON representation of the ContentCloneRequest object. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsCreate

> BlogPost PostCmsV3BlogsPostsCreate(ctx).BlogPost(blogPost).Execute()

Create a new Blog Post



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    blogPost := *openapiclient.NewBlogPost("Id_example", "Slug_example", "ContentGroupId_example", "Campaign_example", int32(123), "State_example", "Name_example", "MabExperimentId_example", false, "AuthorName_example", "AbTestId_example", "CreatedById_example", "UpdatedById_example", "Domain_example", "AbStatus_example", "FolderId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "Language_example", "TranslatedFromId_example", map[string]ContentLanguageVariation{"key": *openapiclient.NewContentLanguageVariation(int64(123), "Name_example", "Slug_example", "State_example", "AuthorName_example", "Password_example", false, []map[string]interface{}{map[string]interface{}(123)}, "Campaign_example", false, time.Now(), time.Now(), time.Now())}, int32(123), "DynamicPageDataSourceId_example", "BlogAuthorId_example", []int64{int64(123)}, "HtmlTitle_example", false, false, "PostBody_example", "PostSummary_example", "RssBody_example", "RssSummary_example", false, false, int64(123), "PageExpiryRedirectUrl_example", int64(123), false, false, false, false, "FeaturedImage_example", "FeaturedImageAltText_example", "LinkRelCanonicalUrl_example", "ContentTypeCategory_example", []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}}, "MetaDescription_example", "HeadHtml_example", "FooterHtml_example", false, false, []map[string]interface{}{map[string]interface{}(123)}, map[string]LayoutSection{"key": *openapiclient.NewLayoutSection(int32(123), int32(123), "Name_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": *openapiclient.NewLayoutSection(int32(123), int32(123), "Name_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": }}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData(*openapiclient.NewStyles("VerticalAlignment_example", *openapiclient.NewRGBAColor(int32(123), int32(123), int32(123), float32(123)), *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), *openapiclient.NewGradient(*openapiclient.NewSideOrCorner("VerticalSide_example", "HorizontalSide_example"), *openapiclient.NewAngle(float32(123), "Units_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop(*openapiclient.NewRGBAColor(int32(123), int32(123), int32(123), float32(123)))}), int32(123), false, "FlexboxPositioning_example"), "CssClass_example")}, []openapiclient.LayoutSection{}, "CssClass_example", "CssStyle_example", "CssId_example", *openapiclient.NewStyles("VerticalAlignment_example", , *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), *openapiclient.NewGradient(*openapiclient.NewSideOrCorner("VerticalSide_example", "HorizontalSide_example"), *openapiclient.NewAngle(float32(123), "Units_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop()}), int32(123), false, "FlexboxPositioning_example"))}}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData(, "CssClass_example")}, []openapiclient.LayoutSection{}, "CssClass_example", "CssStyle_example", "CssId_example", )}, map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "Url_example", "Password_example", "CurrentState_example", time.Now(), time.Now(), time.Now(), time.Now()) // BlogPost | The JSON representation of a new Blog Post.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.PostCmsV3BlogsPostsCreate(context.Background()).BlogPost(blogPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.PostCmsV3BlogsPostsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3BlogsPostsCreate`: BlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.PostCmsV3BlogsPostsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blogPost** | [**BlogPost**](BlogPost.md) | The JSON representation of a new Blog Post. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsMultiLanguageAttachToLangGroupAttachToLangGroup

> Error PostCmsV3BlogsPostsMultiLanguageAttachToLangGroupAttachToLangGroup(ctx).AttachToLangPrimaryRequestVNext(attachToLangPrimaryRequestVNext).Execute()

Attach a Blog Post to a multi-language group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    attachToLangPrimaryRequestVNext := *openapiclient.NewAttachToLangPrimaryRequestVNext("Id_example", "Language_example", "PrimaryId_example") // AttachToLangPrimaryRequestVNext | The JSON representation of the AttachToLangPrimaryRequest object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.PostCmsV3BlogsPostsMultiLanguageAttachToLangGroupAttachToLangGroup(context.Background()).AttachToLangPrimaryRequestVNext(attachToLangPrimaryRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.PostCmsV3BlogsPostsMultiLanguageAttachToLangGroupAttachToLangGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3BlogsPostsMultiLanguageAttachToLangGroupAttachToLangGroup`: Error
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.PostCmsV3BlogsPostsMultiLanguageAttachToLangGroupAttachToLangGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsMultiLanguageAttachToLangGroupAttachToLangGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachToLangPrimaryRequestVNext** | [**AttachToLangPrimaryRequestVNext**](AttachToLangPrimaryRequestVNext.md) | The JSON representation of the AttachToLangPrimaryRequest object. | 

### Return type

[**Error**](Error.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsMultiLanguageCreateLanguageVariationCreateLangVariation

> BlogPost PostCmsV3BlogsPostsMultiLanguageCreateLanguageVariationCreateLangVariation(ctx).BlogPostLanguageCloneRequestVNext(blogPostLanguageCloneRequestVNext).Execute()

Create a new language variation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    blogPostLanguageCloneRequestVNext := *openapiclient.NewBlogPostLanguageCloneRequestVNext("Id_example") // BlogPostLanguageCloneRequestVNext | The JSON representation of the BlogPostLanguageCloneRequestVNext object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.PostCmsV3BlogsPostsMultiLanguageCreateLanguageVariationCreateLangVariation(context.Background()).BlogPostLanguageCloneRequestVNext(blogPostLanguageCloneRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.PostCmsV3BlogsPostsMultiLanguageCreateLanguageVariationCreateLangVariation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3BlogsPostsMultiLanguageCreateLanguageVariationCreateLangVariation`: BlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.PostCmsV3BlogsPostsMultiLanguageCreateLanguageVariationCreateLangVariation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsMultiLanguageCreateLanguageVariationCreateLangVariationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blogPostLanguageCloneRequestVNext** | [**BlogPostLanguageCloneRequestVNext**](BlogPostLanguageCloneRequestVNext.md) | The JSON representation of the BlogPostLanguageCloneRequestVNext object. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsMultiLanguageDetachFromLangGroupDetachFromLangGroup

> Error PostCmsV3BlogsPostsMultiLanguageDetachFromLangGroupDetachFromLangGroup(ctx).DetachFromLangGroupRequestVNext(detachFromLangGroupRequestVNext).Execute()

Detach a Blog Post from a multi-language group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    detachFromLangGroupRequestVNext := *openapiclient.NewDetachFromLangGroupRequestVNext("Id_example") // DetachFromLangGroupRequestVNext | The JSON representation of the DetachFromLangGroupRequest object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.PostCmsV3BlogsPostsMultiLanguageDetachFromLangGroupDetachFromLangGroup(context.Background()).DetachFromLangGroupRequestVNext(detachFromLangGroupRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.PostCmsV3BlogsPostsMultiLanguageDetachFromLangGroupDetachFromLangGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3BlogsPostsMultiLanguageDetachFromLangGroupDetachFromLangGroup`: Error
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.PostCmsV3BlogsPostsMultiLanguageDetachFromLangGroupDetachFromLangGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsMultiLanguageDetachFromLangGroupDetachFromLangGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **detachFromLangGroupRequestVNext** | [**DetachFromLangGroupRequestVNext**](DetachFromLangGroupRequestVNext.md) | The JSON representation of the DetachFromLangGroupRequest object. | 

### Return type

[**Error**](Error.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsMultiLanguageUpdateLanguagesUpdateLangs

> Error PostCmsV3BlogsPostsMultiLanguageUpdateLanguagesUpdateLangs(ctx).UpdateLanguagesRequestVNext(updateLanguagesRequestVNext).Execute()

Update languages of multi-language group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    updateLanguagesRequestVNext := *openapiclient.NewUpdateLanguagesRequestVNext("PrimaryId_example", map[string]string{"key": "Inner_example"}) // UpdateLanguagesRequestVNext | The JSON representation of the SetNewLanguagePrimaryRequest object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.PostCmsV3BlogsPostsMultiLanguageUpdateLanguagesUpdateLangs(context.Background()).UpdateLanguagesRequestVNext(updateLanguagesRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.PostCmsV3BlogsPostsMultiLanguageUpdateLanguagesUpdateLangs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3BlogsPostsMultiLanguageUpdateLanguagesUpdateLangs`: Error
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.PostCmsV3BlogsPostsMultiLanguageUpdateLanguagesUpdateLangs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsMultiLanguageUpdateLanguagesUpdateLangsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateLanguagesRequestVNext** | [**UpdateLanguagesRequestVNext**](UpdateLanguagesRequestVNext.md) | The JSON representation of the SetNewLanguagePrimaryRequest object. | 

### Return type

[**Error**](Error.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsObjectIdDraftPushLivePushLive

> PostCmsV3BlogsPostsObjectIdDraftPushLivePushLive(ctx, objectId).Execute()

Push Blog Post draft edits live



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    objectId := "objectId_example" // string | The id of the Blog Post for which it's draft will be pushed live.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.PostCmsV3BlogsPostsObjectIdDraftPushLivePushLive(context.Background(), objectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.PostCmsV3BlogsPostsObjectIdDraftPushLivePushLive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The id of the Blog Post for which it&#39;s draft will be pushed live. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsObjectIdDraftPushLivePushLiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsObjectIdDraftResetResetDraft

> PostCmsV3BlogsPostsObjectIdDraftResetResetDraft(ctx, objectId).Execute()

Reset the Blog Post draft to the live version



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    objectId := "objectId_example" // string | The id of the Blog Post for which it's draft will be reset.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.PostCmsV3BlogsPostsObjectIdDraftResetResetDraft(context.Background(), objectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.PostCmsV3BlogsPostsObjectIdDraftResetResetDraft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The id of the Blog Post for which it&#39;s draft will be reset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsObjectIdDraftResetResetDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion

> BlogPost PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion(ctx, objectId, revisionId).Execute()

Restore a previous version of a blog post



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    objectId := "objectId_example" // string | The Blog Post id.
    revisionId := "revisionId_example" // string | The Blog Post version id to restore.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion(context.Background(), objectId, revisionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion`: BlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreRestorePreviousVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 
**revisionId** | **string** | The Blog Post version id to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreRestorePreviousVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft

> BlogPost PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft(ctx, objectId, revisionId).Execute()

Restore a previous version of a blog post, to the draft version of the blog post



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    objectId := "objectId_example" // string | The Blog Post id.
    revisionId := int64(789) // int64 | The Blog Post version id to restore.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft(context.Background(), objectId, revisionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft`: BlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 
**revisionId** | **int64** | The Blog Post version id to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraftRestorePreviousVersionToDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsScheduleSchedule

> PostCmsV3BlogsPostsScheduleSchedule(ctx).ContentScheduleRequestVNext(contentScheduleRequestVNext).Execute()

Schedule a Blog Post to be Published



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    contentScheduleRequestVNext := *openapiclient.NewContentScheduleRequestVNext("Id_example", time.Now()) // ContentScheduleRequestVNext | The JSON representation of the ContentScheduleRequestVNext object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.PostCmsV3BlogsPostsScheduleSchedule(context.Background()).ContentScheduleRequestVNext(contentScheduleRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.PostCmsV3BlogsPostsScheduleSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsScheduleScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentScheduleRequestVNext** | [**ContentScheduleRequestVNext**](ContentScheduleRequestVNext.md) | The JSON representation of the ContentScheduleRequestVNext object. | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCmsV3BlogsPostsMultiLanguageSetNewLangPrimarySetLangPrimary

> PutCmsV3BlogsPostsMultiLanguageSetNewLangPrimarySetLangPrimary(ctx).SetNewLanguagePrimaryRequestVNext(setNewLanguagePrimaryRequestVNext).Execute()

Set a new primary language



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    setNewLanguagePrimaryRequestVNext := *openapiclient.NewSetNewLanguagePrimaryRequestVNext("Id_example") // SetNewLanguagePrimaryRequestVNext | The JSON representation of the SetNewLanguagePrimaryRequest object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostsApi.PutCmsV3BlogsPostsMultiLanguageSetNewLangPrimarySetLangPrimary(context.Background()).SetNewLanguagePrimaryRequestVNext(setNewLanguagePrimaryRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.PutCmsV3BlogsPostsMultiLanguageSetNewLangPrimarySetLangPrimary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutCmsV3BlogsPostsMultiLanguageSetNewLangPrimarySetLangPrimaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setNewLanguagePrimaryRequestVNext** | [**SetNewLanguagePrimaryRequestVNext**](SetNewLanguagePrimaryRequestVNext.md) | The JSON representation of the SetNewLanguagePrimaryRequest object. | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

