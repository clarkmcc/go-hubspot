# \TagApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCmsV3BlogsTagsObjectIdArchive**](TagApi.md#DeleteCmsV3BlogsTagsObjectIdArchive) | **Delete** /cms/v3/blogs/tags/{objectId} | Delete a Blog Tag
[**GetCmsV3BlogsTagsGetPage**](TagApi.md#GetCmsV3BlogsTagsGetPage) | **Get** /cms/v3/blogs/tags | Get all Blog Tags
[**GetCmsV3BlogsTagsObjectIdGetById**](TagApi.md#GetCmsV3BlogsTagsObjectIdGetById) | **Get** /cms/v3/blogs/tags/{objectId} | Retrieve a Blog Tag
[**PatchCmsV3BlogsTagsObjectIdUpdate**](TagApi.md#PatchCmsV3BlogsTagsObjectIdUpdate) | **Patch** /cms/v3/blogs/tags/{objectId} | Update a Blog Tag
[**PostCmsV3BlogsTagsBatchArchiveArchiveBatch**](TagApi.md#PostCmsV3BlogsTagsBatchArchiveArchiveBatch) | **Post** /cms/v3/blogs/tags/batch/archive | Archive a batch of Blog Tags
[**PostCmsV3BlogsTagsBatchCreateCreateBatch**](TagApi.md#PostCmsV3BlogsTagsBatchCreateCreateBatch) | **Post** /cms/v3/blogs/tags/batch/create | Create a batch of Blog Tags
[**PostCmsV3BlogsTagsBatchReadReadBatch**](TagApi.md#PostCmsV3BlogsTagsBatchReadReadBatch) | **Post** /cms/v3/blogs/tags/batch/read | Retrieve a batch of Blog Tags
[**PostCmsV3BlogsTagsBatchUpdateUpdateBatch**](TagApi.md#PostCmsV3BlogsTagsBatchUpdateUpdateBatch) | **Post** /cms/v3/blogs/tags/batch/update | Update a batch of Blog Tags
[**PostCmsV3BlogsTagsCreate**](TagApi.md#PostCmsV3BlogsTagsCreate) | **Post** /cms/v3/blogs/tags | Create a new Blog Tag



## DeleteCmsV3BlogsTagsObjectIdArchive

> DeleteCmsV3BlogsTagsObjectIdArchive(ctx, objectId).Archived(archived).Execute()

Delete a Blog Tag



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
    objectId := "objectId_example" // string | The Blog Tag id.
    archived := true // bool | Whether to return only results that have been archived. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagApi.DeleteCmsV3BlogsTagsObjectIdArchive(context.Background(), objectId).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagApi.DeleteCmsV3BlogsTagsObjectIdArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Tag id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCmsV3BlogsTagsObjectIdArchiveRequest struct via the builder pattern


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


## GetCmsV3BlogsTagsGetPage

> CollectionResponseWithTotalTagForwardPaging GetCmsV3BlogsTagsGetPage(ctx).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Execute()

Get all Blog Tags



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
    createdAt := time.Now() // time.Time | Only return Blog Tags created at exactly the specified time. (optional)
    createdAfter := time.Now() // time.Time | Only return Blog Tags created after the specified time. (optional)
    createdBefore := time.Now() // time.Time | Only return Blog Tags created before the specified time. (optional)
    updatedAt := time.Now() // time.Time | Only return Blog Tags last updated at exactly the specified time. (optional)
    updatedAfter := time.Now() // time.Time | Only return Blog Tags last updated after the specified time. (optional)
    updatedBefore := time.Now() // time.Time | Only return Blog Tags last updated before the specified time. (optional)
    sort := []string{"Inner_example"} // []string | Specifies which fields to use for sorting results. Valid fields are `name`, `createdAt`, `updatedAt`, `createdBy`, `updatedBy`. `createdAt` will be used by default. (optional)
    after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
    limit := int32(56) // int32 | The maximum number of results to return. Default is 100. (optional)
    archived := true // bool | Specifies whether to return archived Blog Tags. Defaults to `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagApi.GetCmsV3BlogsTagsGetPage(context.Background()).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagApi.GetCmsV3BlogsTagsGetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3BlogsTagsGetPage`: CollectionResponseWithTotalTagForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `TagApi.GetCmsV3BlogsTagsGetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsTagsGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAt** | **time.Time** | Only return Blog Tags created at exactly the specified time. | 
 **createdAfter** | **time.Time** | Only return Blog Tags created after the specified time. | 
 **createdBefore** | **time.Time** | Only return Blog Tags created before the specified time. | 
 **updatedAt** | **time.Time** | Only return Blog Tags last updated at exactly the specified time. | 
 **updatedAfter** | **time.Time** | Only return Blog Tags last updated after the specified time. | 
 **updatedBefore** | **time.Time** | Only return Blog Tags last updated before the specified time. | 
 **sort** | **[]string** | Specifies which fields to use for sorting results. Valid fields are &#x60;name&#x60;, &#x60;createdAt&#x60;, &#x60;updatedAt&#x60;, &#x60;createdBy&#x60;, &#x60;updatedBy&#x60;. &#x60;createdAt&#x60; will be used by default. | 
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is 100. | 
 **archived** | **bool** | Specifies whether to return archived Blog Tags. Defaults to &#x60;false&#x60;. | 

### Return type

[**CollectionResponseWithTotalTagForwardPaging**](CollectionResponseWithTotalTagForwardPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3BlogsTagsObjectIdGetById

> Tag GetCmsV3BlogsTagsObjectIdGetById(ctx, objectId).Archived(archived).Execute()

Retrieve a Blog Tag



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
    objectId := "objectId_example" // string | The Blog Tag id.
    archived := true // bool | Specifies whether to return archived Blog Tags. Defaults to `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagApi.GetCmsV3BlogsTagsObjectIdGetById(context.Background(), objectId).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagApi.GetCmsV3BlogsTagsObjectIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3BlogsTagsObjectIdGetById`: Tag
    fmt.Fprintf(os.Stdout, "Response from `TagApi.GetCmsV3BlogsTagsObjectIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Tag id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsTagsObjectIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Specifies whether to return archived Blog Tags. Defaults to &#x60;false&#x60;. | 

### Return type

[**Tag**](Tag.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCmsV3BlogsTagsObjectIdUpdate

> Tag PatchCmsV3BlogsTagsObjectIdUpdate(ctx, objectId).Tag(tag).Archived(archived).Execute()

Update a Blog Tag



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
    objectId := "objectId_example" // string | The Blog Tag id.
    tag := *openapiclient.NewTag("Id_example", "Name_example", time.Now(), time.Now(), time.Now()) // Tag | The JSON representation of the updated Blog Tag.
    archived := true // bool | Specifies whether to update archived Blog Tags. Defaults to `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagApi.PatchCmsV3BlogsTagsObjectIdUpdate(context.Background(), objectId).Tag(tag).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagApi.PatchCmsV3BlogsTagsObjectIdUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCmsV3BlogsTagsObjectIdUpdate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `TagApi.PatchCmsV3BlogsTagsObjectIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Tag id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCmsV3BlogsTagsObjectIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | [**Tag**](Tag.md) | The JSON representation of the updated Blog Tag. | 
 **archived** | **bool** | Specifies whether to update archived Blog Tags. Defaults to &#x60;false&#x60;. | 

### Return type

[**Tag**](Tag.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsTagsBatchArchiveArchiveBatch

> PostCmsV3BlogsTagsBatchArchiveArchiveBatch(ctx).BatchInputString(batchInputString).Execute()

Archive a batch of Blog Tags



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
    batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of Blog Tag ids.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagApi.PostCmsV3BlogsTagsBatchArchiveArchiveBatch(context.Background()).BatchInputString(batchInputString).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagApi.PostCmsV3BlogsTagsBatchArchiveArchiveBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsTagsBatchArchiveArchiveBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Blog Tag ids. | 

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


## PostCmsV3BlogsTagsBatchCreateCreateBatch

> map[string]interface{} PostCmsV3BlogsTagsBatchCreateCreateBatch(ctx).BatchInputTag(batchInputTag).Execute()

Create a batch of Blog Tags



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
    batchInputTag := *openapiclient.NewBatchInputTag([]openapiclient.Tag{*openapiclient.NewTag("Id_example", "Name_example", time.Now(), time.Now(), time.Now())}) // BatchInputTag | The JSON array of new Blog Tags to create.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagApi.PostCmsV3BlogsTagsBatchCreateCreateBatch(context.Background()).BatchInputTag(batchInputTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagApi.PostCmsV3BlogsTagsBatchCreateCreateBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3BlogsTagsBatchCreateCreateBatch`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TagApi.PostCmsV3BlogsTagsBatchCreateCreateBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsTagsBatchCreateCreateBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputTag** | [**BatchInputTag**](BatchInputTag.md) | The JSON array of new Blog Tags to create. | 

### Return type

**map[string]interface{}**

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsTagsBatchReadReadBatch

> map[string]interface{} PostCmsV3BlogsTagsBatchReadReadBatch(ctx).BatchInputString(batchInputString).Archived(archived).Execute()

Retrieve a batch of Blog Tags



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
    batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of Blog Tag ids.
    archived := true // bool | Specifies whether to return archived Blog Tags. Defaults to `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagApi.PostCmsV3BlogsTagsBatchReadReadBatch(context.Background()).BatchInputString(batchInputString).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagApi.PostCmsV3BlogsTagsBatchReadReadBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3BlogsTagsBatchReadReadBatch`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TagApi.PostCmsV3BlogsTagsBatchReadReadBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsTagsBatchReadReadBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Blog Tag ids. | 
 **archived** | **bool** | Specifies whether to return archived Blog Tags. Defaults to &#x60;false&#x60;. | 

### Return type

**map[string]interface{}**

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsTagsBatchUpdateUpdateBatch

> map[string]interface{} PostCmsV3BlogsTagsBatchUpdateUpdateBatch(ctx).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()

Update a batch of Blog Tags



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
    batchInputJsonNode := *openapiclient.NewBatchInputJsonNode([]map[string]interface{}{map[string]interface{}(123)}) // BatchInputJsonNode | A JSON array of the JSON representations of the updated Blog Tags.
    archived := true // bool | Specifies whether to update archived Blog Tags. Defaults to `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagApi.PostCmsV3BlogsTagsBatchUpdateUpdateBatch(context.Background()).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagApi.PostCmsV3BlogsTagsBatchUpdateUpdateBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3BlogsTagsBatchUpdateUpdateBatch`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TagApi.PostCmsV3BlogsTagsBatchUpdateUpdateBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsTagsBatchUpdateUpdateBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputJsonNode** | [**BatchInputJsonNode**](BatchInputJsonNode.md) | A JSON array of the JSON representations of the updated Blog Tags. | 
 **archived** | **bool** | Specifies whether to update archived Blog Tags. Defaults to &#x60;false&#x60;. | 

### Return type

**map[string]interface{}**

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsTagsCreate

> Tag PostCmsV3BlogsTagsCreate(ctx).Tag(tag).Execute()

Create a new Blog Tag



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
    tag := *openapiclient.NewTag("Id_example", "Name_example", time.Now(), time.Now(), time.Now()) // Tag | The JSON representation of a new Blog Tag.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagApi.PostCmsV3BlogsTagsCreate(context.Background()).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagApi.PostCmsV3BlogsTagsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3BlogsTagsCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `TagApi.PostCmsV3BlogsTagsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsTagsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | [**Tag**](Tag.md) | The JSON representation of a new Blog Tag. | 

### Return type

[**Tag**](Tag.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

