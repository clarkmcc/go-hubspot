# \BlogAuthorsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Archive**](BlogAuthorsApi.md#Archive) | **Delete** /cms/v3/blogs/authors/{objectId} | Delete a Blog Author
[**AttachToLanguageGroup**](BlogAuthorsApi.md#AttachToLanguageGroup) | **Post** /cms/v3/blogs/authors/multi-language/attach-to-lang-group | Attach a Blog Author to a multi-language group
[**BatchArchive**](BlogAuthorsApi.md#BatchArchive) | **Post** /cms/v3/blogs/authors/batch/archive | Delete a batch of Blog Authors
[**BatchCreate**](BlogAuthorsApi.md#BatchCreate) | **Post** /cms/v3/blogs/authors/batch/create | Create a batch of Blog Authors
[**BatchRead**](BlogAuthorsApi.md#BatchRead) | **Post** /cms/v3/blogs/authors/batch/read | Retrieve a batch of Blog Authors
[**BatchUpdate**](BlogAuthorsApi.md#BatchUpdate) | **Post** /cms/v3/blogs/authors/batch/update | Update a batch of Blog Authors
[**Create**](BlogAuthorsApi.md#Create) | **Post** /cms/v3/blogs/authors | Create a new Blog Author
[**CreateLanguageVariation**](BlogAuthorsApi.md#CreateLanguageVariation) | **Post** /cms/v3/blogs/authors/multi-language/create-language-variation | Create a new language variation
[**DetachFromLanguageGroup**](BlogAuthorsApi.md#DetachFromLanguageGroup) | **Post** /cms/v3/blogs/authors/multi-language/detach-from-lang-group | Detach a Blog Author from a multi-language group
[**GetByID**](BlogAuthorsApi.md#GetByID) | **Get** /cms/v3/blogs/authors/{objectId} | Retrieve a Blog Author
[**GetPage**](BlogAuthorsApi.md#GetPage) | **Get** /cms/v3/blogs/authors | Get all Blog Authors
[**SetLanguagePrimary**](BlogAuthorsApi.md#SetLanguagePrimary) | **Put** /cms/v3/blogs/authors/multi-language/set-new-lang-primary | Set a new primary language
[**Update**](BlogAuthorsApi.md#Update) | **Patch** /cms/v3/blogs/authors/{objectId} | Update a Blog Author
[**UpdateLanguages**](BlogAuthorsApi.md#UpdateLanguages) | **Post** /cms/v3/blogs/authors/multi-language/update-languages | Update languages of multi-language group



## Archive

> Archive(ctx, objectId).Archived(archived).Execute()

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
    resp, r, err := apiClient.BlogAuthorsApi.Archive(context.Background(), objectId).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.Archive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Whether to return only results that have been archived. | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachToLanguageGroup

> AttachToLanguageGroup(ctx).AttachToLangPrimaryRequestVNext(attachToLangPrimaryRequestVNext).Execute()

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
    attachToLangPrimaryRequestVNext := *openapiclient.NewAttachToLangPrimaryRequestVNext("Language_example", "Id_example", "PrimaryId_example") // AttachToLangPrimaryRequestVNext | The JSON representation of the AttachToLangPrimaryRequest object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogAuthorsApi.AttachToLanguageGroup(context.Background()).AttachToLangPrimaryRequestVNext(attachToLangPrimaryRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.AttachToLanguageGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachToLanguageGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachToLangPrimaryRequestVNext** | [**AttachToLangPrimaryRequestVNext**](AttachToLangPrimaryRequestVNext.md) | The JSON representation of the AttachToLangPrimaryRequest object. | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchArchive

> BatchArchive(ctx).BatchInputString(batchInputString).Execute()

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
    resp, r, err := apiClient.BlogAuthorsApi.BatchArchive(context.Background()).BatchInputString(batchInputString).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.BatchArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Blog Author ids. | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchCreate

> BatchResponseBlogAuthor BatchCreate(ctx).BatchInputBlogAuthor(batchInputBlogAuthor).Execute()

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
    batchInputBlogAuthor := *openapiclient.NewBatchInputBlogAuthor([]openapiclient.BlogAuthor{*openapiclient.NewBlogAuthor("Website_example", "DisplayName_example", time.Now(), "Facebook_example", "FullName_example", "Bio_example", "Language_example", "Linkedin_example", "Avatar_example", int64(123), "Twitter_example", time.Now(), "Name_example", "Id_example", time.Now(), "Email_example", "Slug_example")}) // BatchInputBlogAuthor | The JSON array of new Blog Authors to create.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogAuthorsApi.BatchCreate(context.Background()).BatchInputBlogAuthor(batchInputBlogAuthor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.BatchCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchCreate`: BatchResponseBlogAuthor
    fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsApi.BatchCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputBlogAuthor** | [**BatchInputBlogAuthor**](BatchInputBlogAuthor.md) | The JSON array of new Blog Authors to create. | 

### Return type

[**BatchResponseBlogAuthor**](BatchResponseBlogAuthor.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchRead

> BatchResponseBlogAuthor BatchRead(ctx).BatchInputString(batchInputString).Archived(archived).Execute()

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
    resp, r, err := apiClient.BlogAuthorsApi.BatchRead(context.Background()).BatchInputString(batchInputString).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.BatchRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchRead`: BatchResponseBlogAuthor
    fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsApi.BatchRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Blog Author ids. | 
 **archived** | **bool** | Specifies whether to return deleted Blog Authors. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseBlogAuthor**](BatchResponseBlogAuthor.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchUpdate

> BatchResponseBlogAuthor BatchUpdate(ctx).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()

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
    resp, r, err := apiClient.BlogAuthorsApi.BatchUpdate(context.Background()).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.BatchUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchUpdate`: BatchResponseBlogAuthor
    fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsApi.BatchUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputJsonNode** | [**BatchInputJsonNode**](BatchInputJsonNode.md) | A JSON array of the JSON representations of the updated Blog Authors. | 
 **archived** | **bool** | Specifies whether to update deleted Blog Authors. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseBlogAuthor**](BatchResponseBlogAuthor.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> BlogAuthor Create(ctx).BlogAuthor(blogAuthor).Execute()

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
    blogAuthor := *openapiclient.NewBlogAuthor("Website_example", "DisplayName_example", time.Now(), "Facebook_example", "FullName_example", "Bio_example", "Language_example", "Linkedin_example", "Avatar_example", int64(123), "Twitter_example", time.Now(), "Name_example", "Id_example", time.Now(), "Email_example", "Slug_example") // BlogAuthor | The JSON representation of a new Blog Author.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogAuthorsApi.Create(context.Background()).BlogAuthor(blogAuthor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: BlogAuthor
    fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blogAuthor** | [**BlogAuthor**](BlogAuthor.md) | The JSON representation of a new Blog Author. | 

### Return type

[**BlogAuthor**](BlogAuthor.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLanguageVariation

> BlogAuthor CreateLanguageVariation(ctx).BlogAuthorCloneRequestVNext(blogAuthorCloneRequestVNext).Execute()

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
    blogAuthorCloneRequestVNext := *openapiclient.NewBlogAuthorCloneRequestVNext("Id_example", *openapiclient.NewBlogAuthor("Website_example", "DisplayName_example", time.Now(), "Facebook_example", "FullName_example", "Bio_example", "Language_example", "Linkedin_example", "Avatar_example", int64(123), "Twitter_example", time.Now(), "Name_example", "Id_example", time.Now(), "Email_example", "Slug_example")) // BlogAuthorCloneRequestVNext | The JSON representation of the ContentLanguageCloneRequest object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogAuthorsApi.CreateLanguageVariation(context.Background()).BlogAuthorCloneRequestVNext(blogAuthorCloneRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.CreateLanguageVariation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLanguageVariation`: BlogAuthor
    fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsApi.CreateLanguageVariation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLanguageVariationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blogAuthorCloneRequestVNext** | [**BlogAuthorCloneRequestVNext**](BlogAuthorCloneRequestVNext.md) | The JSON representation of the ContentLanguageCloneRequest object. | 

### Return type

[**BlogAuthor**](BlogAuthor.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachFromLanguageGroup

> DetachFromLanguageGroup(ctx).DetachFromLangGroupRequestVNext(detachFromLangGroupRequestVNext).Execute()

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
    resp, r, err := apiClient.BlogAuthorsApi.DetachFromLanguageGroup(context.Background()).DetachFromLangGroupRequestVNext(detachFromLangGroupRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.DetachFromLanguageGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDetachFromLanguageGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **detachFromLangGroupRequestVNext** | [**DetachFromLangGroupRequestVNext**](DetachFromLangGroupRequestVNext.md) | The JSON representation of the DetachFromLangGroupRequest object. | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByID

> BlogAuthor GetByID(ctx, objectId).Archived(archived).Property(property).Execute()

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
    property := "property_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogAuthorsApi.GetByID(context.Background(), objectId).Archived(archived).Property(property).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.GetByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetByID`: BlogAuthor
    fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsApi.GetByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Author id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Specifies whether to return deleted Blog Authors. Defaults to &#x60;false&#x60;. | 
 **property** | **string** |  | 

### Return type

[**BlogAuthor**](BlogAuthor.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPage

> CollectionResponseWithTotalBlogAuthorForwardPaging GetPage(ctx).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Property(property).Execute()

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
    property := "property_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogAuthorsApi.GetPage(context.Background()).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Property(property).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.GetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPage`: CollectionResponseWithTotalBlogAuthorForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsApi.GetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPageRequest struct via the builder pattern


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
 **property** | **string** |  | 

### Return type

[**CollectionResponseWithTotalBlogAuthorForwardPaging**](CollectionResponseWithTotalBlogAuthorForwardPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetLanguagePrimary

> SetLanguagePrimary(ctx).SetNewLanguagePrimaryRequestVNext(setNewLanguagePrimaryRequestVNext).Execute()

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
    resp, r, err := apiClient.BlogAuthorsApi.SetLanguagePrimary(context.Background()).SetNewLanguagePrimaryRequestVNext(setNewLanguagePrimaryRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.SetLanguagePrimary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetLanguagePrimaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setNewLanguagePrimaryRequestVNext** | [**SetNewLanguagePrimaryRequestVNext**](SetNewLanguagePrimaryRequestVNext.md) | The JSON representation of the SetNewLanguagePrimaryRequest object. | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> BlogAuthor Update(ctx, objectId).BlogAuthor(blogAuthor).Archived(archived).Execute()

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
    blogAuthor := *openapiclient.NewBlogAuthor("Website_example", "DisplayName_example", time.Now(), "Facebook_example", "FullName_example", "Bio_example", "Language_example", "Linkedin_example", "Avatar_example", int64(123), "Twitter_example", time.Now(), "Name_example", "Id_example", time.Now(), "Email_example", "Slug_example") // BlogAuthor | The JSON representation of the updated Blog Author.
    archived := true // bool | Specifies whether to update deleted Blog Authors. Defaults to `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogAuthorsApi.Update(context.Background(), objectId).BlogAuthor(blogAuthor).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: BlogAuthor
    fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsApi.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Author id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blogAuthor** | [**BlogAuthor**](BlogAuthor.md) | The JSON representation of the updated Blog Author. | 
 **archived** | **bool** | Specifies whether to update deleted Blog Authors. Defaults to &#x60;false&#x60;. | 

### Return type

[**BlogAuthor**](BlogAuthor.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLanguages

> UpdateLanguages(ctx).UpdateLanguagesRequestVNext(updateLanguagesRequestVNext).Execute()

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
    updateLanguagesRequestVNext := *openapiclient.NewUpdateLanguagesRequestVNext(map[string]string{"key": "Inner_example"}, "PrimaryId_example") // UpdateLanguagesRequestVNext | The JSON representation of the UpdateLanguagesRequest object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogAuthorsApi.UpdateLanguages(context.Background()).UpdateLanguagesRequestVNext(updateLanguagesRequestVNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsApi.UpdateLanguages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLanguagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateLanguagesRequestVNext** | [**UpdateLanguagesRequestVNext**](UpdateLanguagesRequestVNext.md) | The JSON representation of the UpdateLanguagesRequest object. | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

