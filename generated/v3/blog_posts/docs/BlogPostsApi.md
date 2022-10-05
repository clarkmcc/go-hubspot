# \BlogPostsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Archive**](BlogPostsApi.md#Archive) | **Delete** /cms/v3/blogs/posts/{objectId} | Delete a Blog Post
[**AttachToLangGroup**](BlogPostsApi.md#AttachToLangGroup) | **Post** /cms/v3/blogs/posts/multi-language/attach-to-lang-group | Attach a Blog Post to a multi-language group
[**BatchArchive**](BlogPostsApi.md#BatchArchive) | **Post** /cms/v3/blogs/posts/batch/archive | Delete a batch of Blog Posts
[**BatchCreate**](BlogPostsApi.md#BatchCreate) | **Post** /cms/v3/blogs/posts/batch/create | Create a batch of Blog Posts
[**BatchRead**](BlogPostsApi.md#BatchRead) | **Post** /cms/v3/blogs/posts/batch/read | Retrieve a batch of Blog Posts
[**BatchUpdate**](BlogPostsApi.md#BatchUpdate) | **Post** /cms/v3/blogs/posts/batch/update | Update a batch of Blog Posts
[**Clone**](BlogPostsApi.md#Clone) | **Post** /cms/v3/blogs/posts/clone | Clone a Blog Post
[**Create**](BlogPostsApi.md#Create) | **Post** /cms/v3/blogs/posts | Create a new Blog Post
[**CreateLangVariation**](BlogPostsApi.md#CreateLangVariation) | **Post** /cms/v3/blogs/posts/multi-language/create-language-variation | Create a new language variation
[**DetachFromLangGroup**](BlogPostsApi.md#DetachFromLangGroup) | **Post** /cms/v3/blogs/posts/multi-language/detach-from-lang-group | Detach a Blog Post from a multi-language group
[**GetByID**](BlogPostsApi.md#GetByID) | **Get** /cms/v3/blogs/posts/{objectId} | Retrieve a Blog Post
[**GetDraftByID**](BlogPostsApi.md#GetDraftByID) | **Get** /cms/v3/blogs/posts/{objectId}/draft | Retrieve the full draft version of the Blog Post
[**GetPage**](BlogPostsApi.md#GetPage) | **Get** /cms/v3/blogs/posts | Get all Blog Posts
[**GetPreviousVersion**](BlogPostsApi.md#GetPreviousVersion) | **Get** /cms/v3/blogs/posts/{objectId}/revisions/{revisionId} | Retrieves a previous version of a blog post
[**GetPreviousVersions**](BlogPostsApi.md#GetPreviousVersions) | **Get** /cms/v3/blogs/posts/{objectId}/revisions | Retrieves all the previous versions of a blog post
[**PushLive**](BlogPostsApi.md#PushLive) | **Post** /cms/v3/blogs/posts/{objectId}/draft/push-live | Push Blog Post draft edits live
[**ResetDraft**](BlogPostsApi.md#ResetDraft) | **Post** /cms/v3/blogs/posts/{objectId}/draft/reset | Reset the Blog Post draft to the live version
[**RestorePreviousVersion**](BlogPostsApi.md#RestorePreviousVersion) | **Post** /cms/v3/blogs/posts/{objectId}/revisions/{revisionId}/restore | Restore a previous version of a blog post
[**RestorePreviousVersionToDraft**](BlogPostsApi.md#RestorePreviousVersionToDraft) | **Post** /cms/v3/blogs/posts/{objectId}/revisions/{revisionId}/restore-to-draft | Restore a previous version of a blog post, to the draft version of the blog post
[**Schedule**](BlogPostsApi.md#Schedule) | **Post** /cms/v3/blogs/posts/schedule | Schedule a Blog Post to be Published
[**SetLangPrimary**](BlogPostsApi.md#SetLangPrimary) | **Put** /cms/v3/blogs/posts/multi-language/set-new-lang-primary | Set a new primary language
[**Update**](BlogPostsApi.md#Update) | **Patch** /cms/v3/blogs/posts/{objectId} | Update a Blog Post
[**UpdateDraft**](BlogPostsApi.md#UpdateDraft) | **Patch** /cms/v3/blogs/posts/{objectId}/draft | Update a Blog Post draft
[**UpdateLangs**](BlogPostsApi.md#UpdateLangs) | **Post** /cms/v3/blogs/posts/multi-language/update-languages | Update languages of multi-language group



## Archive

> Archive(ctx, objectId).Archived(archived).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.Archive(context.Background(), objectId).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.Archive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Whether to return only results that have been archived. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachToLangGroup

> Error AttachToLangGroup(ctx).AttachToLangPrimaryRequestVNext(attachToLangPrimaryRequestVNext).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.AttachToLangGroup(context.Background()).AttachToLangPrimaryRequestVNext(attachToLangPrimaryRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.AttachToLangGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachToLangGroup`: Error
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.AttachToLangGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachToLangGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachToLangPrimaryRequestVNext** | [**AttachToLangPrimaryRequestVNext**](AttachToLangPrimaryRequestVNext.md) | The JSON representation of the AttachToLangPrimaryRequest object. | 

### Return type

[**Error**](Error.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchArchive

> BatchArchive(ctx).BatchInputString(batchInputString).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.BatchArchive(context.Background()).BatchInputString(batchInputString).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.BatchArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Blog Post ids. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchCreate

> BatchResponseBlogPost BatchCreate(ctx).BatchInputBlogPost(batchInputBlogPost).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.BatchCreate(context.Background()).BatchInputBlogPost(batchInputBlogPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.BatchCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchCreate`: BatchResponseBlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.BatchCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputBlogPost** | [**BatchInputBlogPost**](BatchInputBlogPost.md) | The JSON array of new Blog Posts to create. | 

### Return type

[**BatchResponseBlogPost**](BatchResponseBlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchRead

> BatchResponseBlogPost BatchRead(ctx).BatchInputString(batchInputString).Archived(archived).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.BatchRead(context.Background()).BatchInputString(batchInputString).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.BatchRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchRead`: BatchResponseBlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.BatchRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Blog Post ids. | 
 **archived** | **bool** | Specifies whether to return deleted Blog Posts. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseBlogPost**](BatchResponseBlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchUpdate

> BatchResponseBlogPost BatchUpdate(ctx).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.BatchUpdate(context.Background()).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.BatchUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchUpdate`: BatchResponseBlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.BatchUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputJsonNode** | [**BatchInputJsonNode**](BatchInputJsonNode.md) | A JSON array of the JSON representations of the updated Blog Posts. | 
 **archived** | **bool** | Specifies whether to update deleted Blog Posts. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseBlogPost**](BatchResponseBlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Clone

> BlogPost Clone(ctx).ContentCloneRequestVNext(contentCloneRequestVNext).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.Clone(context.Background()).ContentCloneRequestVNext(contentCloneRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.Clone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Clone`: BlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.Clone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentCloneRequestVNext** | [**ContentCloneRequestVNext**](ContentCloneRequestVNext.md) | The JSON representation of the ContentCloneRequest object. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> BlogPost Create(ctx).BlogPost(blogPost).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.Create(context.Background()).BlogPost(blogPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: BlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blogPost** | [**BlogPost**](BlogPost.md) | The JSON representation of a new Blog Post. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLangVariation

> BlogPost CreateLangVariation(ctx).BlogPostLanguageCloneRequestVNext(blogPostLanguageCloneRequestVNext).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.CreateLangVariation(context.Background()).BlogPostLanguageCloneRequestVNext(blogPostLanguageCloneRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.CreateLangVariation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLangVariation`: BlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.CreateLangVariation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLangVariationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blogPostLanguageCloneRequestVNext** | [**BlogPostLanguageCloneRequestVNext**](BlogPostLanguageCloneRequestVNext.md) | The JSON representation of the BlogPostLanguageCloneRequestVNext object. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachFromLangGroup

> Error DetachFromLangGroup(ctx).DetachFromLangGroupRequestVNext(detachFromLangGroupRequestVNext).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.DetachFromLangGroup(context.Background()).DetachFromLangGroupRequestVNext(detachFromLangGroupRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.DetachFromLangGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachFromLangGroup`: Error
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.DetachFromLangGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDetachFromLangGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **detachFromLangGroupRequestVNext** | [**DetachFromLangGroupRequestVNext**](DetachFromLangGroupRequestVNext.md) | The JSON representation of the DetachFromLangGroupRequest object. | 

### Return type

[**Error**](Error.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByID

> BlogPost GetByID(ctx, objectId).Archived(archived).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.GetByID(context.Background(), objectId).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.GetByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetByID`: BlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.GetByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Specifies whether to return deleted Blog Posts. Defaults to &#x60;false&#x60;. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDraftByID

> BlogPost GetDraftByID(ctx, objectId).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.GetDraftByID(context.Background(), objectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.GetDraftByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDraftByID`: BlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.GetDraftByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDraftByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPage

> CollectionResponseWithTotalBlogPostForwardPaging GetPage(ctx).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.GetPage(context.Background()).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.GetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPage`: CollectionResponseWithTotalBlogPostForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.GetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPageRequest struct via the builder pattern


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

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPreviousVersion

> VersionBlogPost GetPreviousVersion(ctx, objectId, revisionId).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.GetPreviousVersion(context.Background(), objectId, revisionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.GetPreviousVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPreviousVersion`: VersionBlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.GetPreviousVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 
**revisionId** | **string** | The Blog Post version id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPreviousVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VersionBlogPost**](VersionBlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPreviousVersions

> CollectionResponseWithTotalVersionBlogPost GetPreviousVersions(ctx, objectId).After(after).Before(before).Limit(limit).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.GetPreviousVersions(context.Background(), objectId).After(after).Before(before).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.GetPreviousVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPreviousVersions`: CollectionResponseWithTotalVersionBlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.GetPreviousVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPreviousVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **before** | **string** |  | 
 **limit** | **int32** | The maximum number of results to return. Default is 100. | 

### Return type

[**CollectionResponseWithTotalVersionBlogPost**](CollectionResponseWithTotalVersionBlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PushLive

> PushLive(ctx, objectId).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.PushLive(context.Background(), objectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.PushLive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPushLiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetDraft

> ResetDraft(ctx, objectId).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.ResetDraft(context.Background(), objectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.ResetDraft``: %v\n", err)
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

Other parameters are passed through a pointer to a apiResetDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestorePreviousVersion

> BlogPost RestorePreviousVersion(ctx, objectId, revisionId).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.RestorePreviousVersion(context.Background(), objectId, revisionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.RestorePreviousVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestorePreviousVersion`: BlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.RestorePreviousVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 
**revisionId** | **string** | The Blog Post version id to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestorePreviousVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestorePreviousVersionToDraft

> BlogPost RestorePreviousVersionToDraft(ctx, objectId, revisionId).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.RestorePreviousVersionToDraft(context.Background(), objectId, revisionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.RestorePreviousVersionToDraft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestorePreviousVersionToDraft`: BlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.RestorePreviousVersionToDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 
**revisionId** | **int64** | The Blog Post version id to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestorePreviousVersionToDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Schedule

> Schedule(ctx).ContentScheduleRequestVNext(contentScheduleRequestVNext).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.Schedule(context.Background()).ContentScheduleRequestVNext(contentScheduleRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.Schedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentScheduleRequestVNext** | [**ContentScheduleRequestVNext**](ContentScheduleRequestVNext.md) | The JSON representation of the ContentScheduleRequestVNext object. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetLangPrimary

> SetLangPrimary(ctx).SetNewLanguagePrimaryRequestVNext(setNewLanguagePrimaryRequestVNext).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.SetLangPrimary(context.Background()).SetNewLanguagePrimaryRequestVNext(setNewLanguagePrimaryRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.SetLangPrimary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetLangPrimaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setNewLanguagePrimaryRequestVNext** | [**SetNewLanguagePrimaryRequestVNext**](SetNewLanguagePrimaryRequestVNext.md) | The JSON representation of the SetNewLanguagePrimaryRequest object. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> BlogPost Update(ctx, objectId).BlogPost(blogPost).Archived(archived).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.Update(context.Background(), objectId).BlogPost(blogPost).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: BlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blogPost** | [**BlogPost**](BlogPost.md) | The JSON representation of the updated Blog Post. | 
 **archived** | **bool** | Specifies whether to update deleted Blog Posts. Defaults to &#x60;false&#x60;. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDraft

> BlogPost UpdateDraft(ctx, objectId).BlogPost(blogPost).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.UpdateDraft(context.Background(), objectId).BlogPost(blogPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.UpdateDraft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDraft`: BlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.UpdateDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blogPost** | [**BlogPost**](BlogPost.md) | The JSON representation of the updated Blog Post to be applied to the draft. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLangs

> Error UpdateLangs(ctx).UpdateLanguagesRequestVNext(updateLanguagesRequestVNext).Execute()

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
    resp, r, err := apiClient.BlogPostsApi.UpdateLangs(context.Background()).UpdateLanguagesRequestVNext(updateLanguagesRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsApi.UpdateLangs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLangs`: Error
    fmt.Fprintf(os.Stdout, "Response from `BlogPostsApi.UpdateLangs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLangsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateLanguagesRequestVNext** | [**UpdateLanguagesRequestVNext**](UpdateLanguagesRequestVNext.md) | The JSON representation of the SetNewLanguagePrimaryRequest object. | 

### Return type

[**Error**](Error.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

