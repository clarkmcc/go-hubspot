# \BlogAuthorsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCmsV3BlogsAuthorsObjectIdArchive**](BlogAuthorsApi.md#DeleteCmsV3BlogsAuthorsObjectIdArchive) | **Delete** /cms/v3/blogs/authors/{objectId} | Delete a Blog Author
[**GetCmsV3BlogsAuthorsGetPage**](BlogAuthorsApi.md#GetCmsV3BlogsAuthorsGetPage) | **Get** /cms/v3/blogs/authors | Get all Blog Authors
[**GetCmsV3BlogsAuthorsObjectIdGetById**](BlogAuthorsApi.md#GetCmsV3BlogsAuthorsObjectIdGetById) | **Get** /cms/v3/blogs/authors/{objectId} | Retrieve a Blog Author
[**PatchCmsV3BlogsAuthorsObjectIdUpdate**](BlogAuthorsApi.md#PatchCmsV3BlogsAuthorsObjectIdUpdate) | **Patch** /cms/v3/blogs/authors/{objectId} | Update a Blog Author
[**PostCmsV3BlogsAuthorsBatchArchiveArchiveBatch**](BlogAuthorsApi.md#PostCmsV3BlogsAuthorsBatchArchiveArchiveBatch) | **Post** /cms/v3/blogs/authors/batch/archive | Delete a batch of Blog Authors
[**PostCmsV3BlogsAuthorsBatchCreateCreateBatch**](BlogAuthorsApi.md#PostCmsV3BlogsAuthorsBatchCreateCreateBatch) | **Post** /cms/v3/blogs/authors/batch/create | Create a batch of Blog Authors
[**PostCmsV3BlogsAuthorsBatchReadReadBatch**](BlogAuthorsApi.md#PostCmsV3BlogsAuthorsBatchReadReadBatch) | **Post** /cms/v3/blogs/authors/batch/read | Retrieve a batch of Blog Authors
[**PostCmsV3BlogsAuthorsBatchUpdateUpdateBatch**](BlogAuthorsApi.md#PostCmsV3BlogsAuthorsBatchUpdateUpdateBatch) | **Post** /cms/v3/blogs/authors/batch/update | Update a batch of Blog Authors
[**PostCmsV3BlogsAuthorsCreate**](BlogAuthorsApi.md#PostCmsV3BlogsAuthorsCreate) | **Post** /cms/v3/blogs/authors | Create a new Blog Author
[**PostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroupAttachToLangGroup**](BlogAuthorsApi.md#PostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroupAttachToLangGroup) | **Post** /cms/v3/blogs/authors/multi-language/attach-to-lang-group | Attach a Blog Author to a multi-language group
[**PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariationCreateLangVariation**](BlogAuthorsApi.md#PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariationCreateLangVariation) | **Post** /cms/v3/blogs/authors/multi-language/create-language-variation | Create a new language variation
[**PostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroupDetachFromLangGroup**](BlogAuthorsApi.md#PostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroupDetachFromLangGroup) | **Post** /cms/v3/blogs/authors/multi-language/detach-from-lang-group | Detach a Blog Author from a multi-language group
[**PostCmsV3BlogsAuthorsMultiLanguageUpdateLanguagesUpdateLangs**](BlogAuthorsApi.md#PostCmsV3BlogsAuthorsMultiLanguageUpdateLanguagesUpdateLangs) | **Post** /cms/v3/blogs/authors/multi-language/update-languages | Update languages of multi-language group
[**PutCmsV3BlogsAuthorsMultiLanguageSetNewLangPrimarySetLangPrimary**](BlogAuthorsApi.md#PutCmsV3BlogsAuthorsMultiLanguageSetNewLangPrimarySetLangPrimary) | **Put** /cms/v3/blogs/authors/multi-language/set-new-lang-primary | Set a new primary language



## DeleteCmsV3BlogsAuthorsObjectIdArchive

> DeleteCmsV3BlogsAuthorsObjectIdArchive(ctx, objectId).Archived(archived).Execute()

Delete a Blog Author



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
    objectId := "objectId_example" // string | The Blog Author id.
    archived := true // bool | Whether to return only results that have been archived. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogAuthorsApi.DeleteCmsV3BlogsAuthorsObjectIdArchive(context.Background(), objectId).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.DeleteCmsV3BlogsAuthorsObjectIdArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Author id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCmsV3BlogsAuthorsObjectIdArchiveRequest struct via the builder pattern


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


## GetCmsV3BlogsAuthorsGetPage

> CollectionResponseWithTotalBlogAuthorForwardPaging GetCmsV3BlogsAuthorsGetPage(ctx).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Execute()

Get all Blog Authors



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
    createdAt := time.Now() // time.Time | Only return Blog Authors created at exactly the specified time. (optional)
    createdAfter := time.Now() // time.Time | Only return Blog Authors created after the specified time. (optional)
    createdBefore := time.Now() // time.Time | Only return Blog Authors created before the specified time. (optional)
    updatedAt := time.Now() // time.Time | Only return Blog Authors last updated at exactly the specified time. (optional)
    updatedAfter := time.Now() // time.Time | Only return Blog Authors last updated after the specified time. (optional)
    updatedBefore := time.Now() // time.Time | Only return Blog Authors last updated before the specified time. (optional)
    sort := []string{"Inner_example"} // []string | Specifies which fields to use for sorting results. Valid fields are `name`, `createdAt`, `updatedAt`, `createdBy`, `updatedBy`. `createdAt` will be used by default. (optional)
    after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
    limit := int32(56) // int32 | The maximum number of results to return. Default is 100. (optional)
    archived := true // bool | Specifies whether to return deleted Blog Authors. Defaults to `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogAuthorsApi.GetCmsV3BlogsAuthorsGetPage(context.Background()).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.GetCmsV3BlogsAuthorsGetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3BlogsAuthorsGetPage`: CollectionResponseWithTotalBlogAuthorForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsApi.GetCmsV3BlogsAuthorsGetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsAuthorsGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAt** | **time.Time** | Only return Blog Authors created at exactly the specified time. | 
 **createdAfter** | **time.Time** | Only return Blog Authors created after the specified time. | 
 **createdBefore** | **time.Time** | Only return Blog Authors created before the specified time. | 
 **updatedAt** | **time.Time** | Only return Blog Authors last updated at exactly the specified time. | 
 **updatedAfter** | **time.Time** | Only return Blog Authors last updated after the specified time. | 
 **updatedBefore** | **time.Time** | Only return Blog Authors last updated before the specified time. | 
 **sort** | **[]string** | Specifies which fields to use for sorting results. Valid fields are &#x60;name&#x60;, &#x60;createdAt&#x60;, &#x60;updatedAt&#x60;, &#x60;createdBy&#x60;, &#x60;updatedBy&#x60;. &#x60;createdAt&#x60; will be used by default. | 
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is 100. | 
 **archived** | **bool** | Specifies whether to return deleted Blog Authors. Defaults to &#x60;false&#x60;. | 

### Return type

[**CollectionResponseWithTotalBlogAuthorForwardPaging**](CollectionResponseWithTotalBlogAuthorForwardPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3BlogsAuthorsObjectIdGetById

> BlogAuthor GetCmsV3BlogsAuthorsObjectIdGetById(ctx, objectId).Archived(archived).Execute()

Retrieve a Blog Author



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
    objectId := "objectId_example" // string | The Blog Author id.
    archived := true // bool | Specifies whether to return deleted Blog Authors. Defaults to `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogAuthorsApi.GetCmsV3BlogsAuthorsObjectIdGetById(context.Background(), objectId).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.GetCmsV3BlogsAuthorsObjectIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3BlogsAuthorsObjectIdGetById`: BlogAuthor
    fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsApi.GetCmsV3BlogsAuthorsObjectIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Author id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsAuthorsObjectIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Specifies whether to return deleted Blog Authors. Defaults to &#x60;false&#x60;. | 

### Return type

[**BlogAuthor**](BlogAuthor.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCmsV3BlogsAuthorsObjectIdUpdate

> BlogAuthor PatchCmsV3BlogsAuthorsObjectIdUpdate(ctx, objectId).BlogAuthor(blogAuthor).Archived(archived).Execute()

Update a Blog Author



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
    objectId := "objectId_example" // string | The Blog Author id.
    blogAuthor := *openapiclient.NewBlogAuthor("Id_example", "FullName_example", "Email_example", "Slug_example", "Language_example", int64(123), "Name_example", "DisplayName_example", "Bio_example", "Website_example", "Twitter_example", "Facebook_example", "Linkedin_example", "Avatar_example", time.Now(), time.Now(), time.Now()) // BlogAuthor | The JSON representation of the updated Blog Author.
    archived := true // bool | Specifies whether to update deleted Blog Authors. Defaults to `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogAuthorsApi.PatchCmsV3BlogsAuthorsObjectIdUpdate(context.Background(), objectId).BlogAuthor(blogAuthor).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.PatchCmsV3BlogsAuthorsObjectIdUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCmsV3BlogsAuthorsObjectIdUpdate`: BlogAuthor
    fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsApi.PatchCmsV3BlogsAuthorsObjectIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Author id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCmsV3BlogsAuthorsObjectIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blogAuthor** | [**BlogAuthor**](BlogAuthor.md) | The JSON representation of the updated Blog Author. | 
 **archived** | **bool** | Specifies whether to update deleted Blog Authors. Defaults to &#x60;false&#x60;. | 

### Return type

[**BlogAuthor**](BlogAuthor.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsAuthorsBatchArchiveArchiveBatch

> PostCmsV3BlogsAuthorsBatchArchiveArchiveBatch(ctx).BatchInputString(batchInputString).Execute()

Delete a batch of Blog Authors



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
    batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of Blog Author ids.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogAuthorsApi.PostCmsV3BlogsAuthorsBatchArchiveArchiveBatch(context.Background()).BatchInputString(batchInputString).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.PostCmsV3BlogsAuthorsBatchArchiveArchiveBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsBatchArchiveArchiveBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Blog Author ids. | 

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


## PostCmsV3BlogsAuthorsBatchCreateCreateBatch

> BatchResponseBlogAuthor PostCmsV3BlogsAuthorsBatchCreateCreateBatch(ctx).BatchInputBlogAuthor(batchInputBlogAuthor).Execute()

Create a batch of Blog Authors



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
    batchInputBlogAuthor := *openapiclient.NewBatchInputBlogAuthor([]openapiclient.BlogAuthor{*openapiclient.NewBlogAuthor("Id_example", "FullName_example", "Email_example", "Slug_example", "Language_example", int64(123), "Name_example", "DisplayName_example", "Bio_example", "Website_example", "Twitter_example", "Facebook_example", "Linkedin_example", "Avatar_example", time.Now(), time.Now(), time.Now())}) // BatchInputBlogAuthor | The JSON array of new Blog Authors to create.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogAuthorsApi.PostCmsV3BlogsAuthorsBatchCreateCreateBatch(context.Background()).BatchInputBlogAuthor(batchInputBlogAuthor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.PostCmsV3BlogsAuthorsBatchCreateCreateBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3BlogsAuthorsBatchCreateCreateBatch`: BatchResponseBlogAuthor
    fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsApi.PostCmsV3BlogsAuthorsBatchCreateCreateBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsBatchCreateCreateBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputBlogAuthor** | [**BatchInputBlogAuthor**](BatchInputBlogAuthor.md) | The JSON array of new Blog Authors to create. | 

### Return type

[**BatchResponseBlogAuthor**](BatchResponseBlogAuthor.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsAuthorsBatchReadReadBatch

> BatchResponseBlogAuthor PostCmsV3BlogsAuthorsBatchReadReadBatch(ctx).BatchInputString(batchInputString).Archived(archived).Execute()

Retrieve a batch of Blog Authors



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
    batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of Blog Author ids.
    archived := true // bool | Specifies whether to return deleted Blog Authors. Defaults to `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogAuthorsApi.PostCmsV3BlogsAuthorsBatchReadReadBatch(context.Background()).BatchInputString(batchInputString).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.PostCmsV3BlogsAuthorsBatchReadReadBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3BlogsAuthorsBatchReadReadBatch`: BatchResponseBlogAuthor
    fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsApi.PostCmsV3BlogsAuthorsBatchReadReadBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsBatchReadReadBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Blog Author ids. | 
 **archived** | **bool** | Specifies whether to return deleted Blog Authors. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseBlogAuthor**](BatchResponseBlogAuthor.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsAuthorsBatchUpdateUpdateBatch

> BatchResponseBlogAuthor PostCmsV3BlogsAuthorsBatchUpdateUpdateBatch(ctx).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()

Update a batch of Blog Authors



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
    batchInputJsonNode := *openapiclient.NewBatchInputJsonNode([]map[string]interface{}{map[string]interface{}(123)}) // BatchInputJsonNode | A JSON array of the JSON representations of the updated Blog Authors.
    archived := true // bool | Specifies whether to update deleted Blog Authors. Defaults to `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogAuthorsApi.PostCmsV3BlogsAuthorsBatchUpdateUpdateBatch(context.Background()).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.PostCmsV3BlogsAuthorsBatchUpdateUpdateBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3BlogsAuthorsBatchUpdateUpdateBatch`: BatchResponseBlogAuthor
    fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsApi.PostCmsV3BlogsAuthorsBatchUpdateUpdateBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsBatchUpdateUpdateBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputJsonNode** | [**BatchInputJsonNode**](BatchInputJsonNode.md) | A JSON array of the JSON representations of the updated Blog Authors. | 
 **archived** | **bool** | Specifies whether to update deleted Blog Authors. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseBlogAuthor**](BatchResponseBlogAuthor.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsAuthorsCreate

> BlogAuthor PostCmsV3BlogsAuthorsCreate(ctx).BlogAuthor(blogAuthor).Execute()

Create a new Blog Author



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
    blogAuthor := *openapiclient.NewBlogAuthor("Id_example", "FullName_example", "Email_example", "Slug_example", "Language_example", int64(123), "Name_example", "DisplayName_example", "Bio_example", "Website_example", "Twitter_example", "Facebook_example", "Linkedin_example", "Avatar_example", time.Now(), time.Now(), time.Now()) // BlogAuthor | The JSON representation of a new Blog Author.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogAuthorsApi.PostCmsV3BlogsAuthorsCreate(context.Background()).BlogAuthor(blogAuthor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.PostCmsV3BlogsAuthorsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3BlogsAuthorsCreate`: BlogAuthor
    fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsApi.PostCmsV3BlogsAuthorsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blogAuthor** | [**BlogAuthor**](BlogAuthor.md) | The JSON representation of a new Blog Author. | 

### Return type

[**BlogAuthor**](BlogAuthor.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroupAttachToLangGroup

> Error PostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroupAttachToLangGroup(ctx).AttachToLangPrimaryRequestVNext(attachToLangPrimaryRequestVNext).Execute()

Attach a Blog Author to a multi-language group



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
    resp, r, err := apiClient.BlogAuthorsApi.PostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroupAttachToLangGroup(context.Background()).AttachToLangPrimaryRequestVNext(attachToLangPrimaryRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.PostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroupAttachToLangGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroupAttachToLangGroup`: Error
    fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsApi.PostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroupAttachToLangGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroupAttachToLangGroupRequest struct via the builder pattern


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


## PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariationCreateLangVariation

> BlogAuthor PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariationCreateLangVariation(ctx).BlogAuthorCloneRequestVNext(blogAuthorCloneRequestVNext).Execute()

Create a new language variation



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
    blogAuthorCloneRequestVNext := *openapiclient.NewBlogAuthorCloneRequestVNext("Id_example", *openapiclient.NewBlogAuthor("Id_example", "FullName_example", "Email_example", "Slug_example", "Language_example", int64(123), "Name_example", "DisplayName_example", "Bio_example", "Website_example", "Twitter_example", "Facebook_example", "Linkedin_example", "Avatar_example", time.Now(), time.Now(), time.Now())) // BlogAuthorCloneRequestVNext | The JSON representation of the ContentLanguageCloneRequest object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogAuthorsApi.PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariationCreateLangVariation(context.Background()).BlogAuthorCloneRequestVNext(blogAuthorCloneRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariationCreateLangVariation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariationCreateLangVariation`: BlogAuthor
    fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsApi.PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariationCreateLangVariation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariationCreateLangVariationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blogAuthorCloneRequestVNext** | [**BlogAuthorCloneRequestVNext**](BlogAuthorCloneRequestVNext.md) | The JSON representation of the ContentLanguageCloneRequest object. | 

### Return type

[**BlogAuthor**](BlogAuthor.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroupDetachFromLangGroup

> Error PostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroupDetachFromLangGroup(ctx).DetachFromLangGroupRequestVNext(detachFromLangGroupRequestVNext).Execute()

Detach a Blog Author from a multi-language group



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
    resp, r, err := apiClient.BlogAuthorsApi.PostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroupDetachFromLangGroup(context.Background()).DetachFromLangGroupRequestVNext(detachFromLangGroupRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.PostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroupDetachFromLangGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroupDetachFromLangGroup`: Error
    fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsApi.PostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroupDetachFromLangGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroupDetachFromLangGroupRequest struct via the builder pattern


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


## PostCmsV3BlogsAuthorsMultiLanguageUpdateLanguagesUpdateLangs

> Error PostCmsV3BlogsAuthorsMultiLanguageUpdateLanguagesUpdateLangs(ctx).UpdateLanguagesRequestVNext(updateLanguagesRequestVNext).Execute()

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
    updateLanguagesRequestVNext := *openapiclient.NewUpdateLanguagesRequestVNext("PrimaryId_example", map[string]string{"key": "Inner_example"}) // UpdateLanguagesRequestVNext | The JSON representation of the UpdateLanguagesRequest object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogAuthorsApi.PostCmsV3BlogsAuthorsMultiLanguageUpdateLanguagesUpdateLangs(context.Background()).UpdateLanguagesRequestVNext(updateLanguagesRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.PostCmsV3BlogsAuthorsMultiLanguageUpdateLanguagesUpdateLangs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3BlogsAuthorsMultiLanguageUpdateLanguagesUpdateLangs`: Error
    fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsApi.PostCmsV3BlogsAuthorsMultiLanguageUpdateLanguagesUpdateLangs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsMultiLanguageUpdateLanguagesUpdateLangsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateLanguagesRequestVNext** | [**UpdateLanguagesRequestVNext**](UpdateLanguagesRequestVNext.md) | The JSON representation of the UpdateLanguagesRequest object. | 

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


## PutCmsV3BlogsAuthorsMultiLanguageSetNewLangPrimarySetLangPrimary

> PutCmsV3BlogsAuthorsMultiLanguageSetNewLangPrimarySetLangPrimary(ctx).SetNewLanguagePrimaryRequestVNext(setNewLanguagePrimaryRequestVNext).Execute()

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
    resp, r, err := apiClient.BlogAuthorsApi.PutCmsV3BlogsAuthorsMultiLanguageSetNewLangPrimarySetLangPrimary(context.Background()).SetNewLanguagePrimaryRequestVNext(setNewLanguagePrimaryRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.PutCmsV3BlogsAuthorsMultiLanguageSetNewLangPrimarySetLangPrimary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutCmsV3BlogsAuthorsMultiLanguageSetNewLangPrimarySetLangPrimaryRequest struct via the builder pattern


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

