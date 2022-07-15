# \BatchApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmV3ObjectsProductsBatchArchiveArchive**](BatchApi.md#PostCrmV3ObjectsProductsBatchArchiveArchive) | **Post** /crm/v3/objects/products/batch/archive | Archive a batch of products by ID
[**PostCrmV3ObjectsProductsBatchCreateCreate**](BatchApi.md#PostCrmV3ObjectsProductsBatchCreateCreate) | **Post** /crm/v3/objects/products/batch/create | Create a batch of products
[**PostCrmV3ObjectsProductsBatchReadRead**](BatchApi.md#PostCrmV3ObjectsProductsBatchReadRead) | **Post** /crm/v3/objects/products/batch/read | Read a batch of products by internal ID, or unique property values
[**PostCrmV3ObjectsProductsBatchUpdateUpdate**](BatchApi.md#PostCrmV3ObjectsProductsBatchUpdateUpdate) | **Post** /crm/v3/objects/products/batch/update | Update a batch of products



## PostCrmV3ObjectsProductsBatchArchiveArchive

> PostCrmV3ObjectsProductsBatchArchiveArchive(ctx).BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId).Execute()

Archive a batch of products by ID

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
    batchInputSimplePublicObjectId := *openapiclient.NewBatchInputSimplePublicObjectId([]openapiclient.SimplePublicObjectId{*openapiclient.NewSimplePublicObjectId("Id_example")}) // BatchInputSimplePublicObjectId | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostCrmV3ObjectsProductsBatchArchiveArchive(context.Background()).BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV3ObjectsProductsBatchArchiveArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsProductsBatchArchiveArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputSimplePublicObjectId** | [**BatchInputSimplePublicObjectId**](BatchInputSimplePublicObjectId.md) |  | 

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


## PostCrmV3ObjectsProductsBatchCreateCreate

> BatchResponseSimplePublicObject PostCrmV3ObjectsProductsBatchCreateCreate(ctx).BatchInputSimplePublicObjectInput(batchInputSimplePublicObjectInput).Execute()

Create a batch of products

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
    batchInputSimplePublicObjectInput := *openapiclient.NewBatchInputSimplePublicObjectInput([]openapiclient.SimplePublicObjectInput{*openapiclient.NewSimplePublicObjectInput(map[string]string{"key": "Inner_example"})}) // BatchInputSimplePublicObjectInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostCrmV3ObjectsProductsBatchCreateCreate(context.Background()).BatchInputSimplePublicObjectInput(batchInputSimplePublicObjectInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV3ObjectsProductsBatchCreateCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ObjectsProductsBatchCreateCreate`: BatchResponseSimplePublicObject
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.PostCrmV3ObjectsProductsBatchCreateCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsProductsBatchCreateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputSimplePublicObjectInput** | [**BatchInputSimplePublicObjectInput**](BatchInputSimplePublicObjectInput.md) |  | 

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ObjectsProductsBatchReadRead

> BatchResponseSimplePublicObject PostCrmV3ObjectsProductsBatchReadRead(ctx).BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId).Archived(archived).Execute()

Read a batch of products by internal ID, or unique property values

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
    batchReadInputSimplePublicObjectId := *openapiclient.NewBatchReadInputSimplePublicObjectId([]string{"Properties_example"}, []string{"PropertiesWithHistory_example"}, []openapiclient.SimplePublicObjectId{*openapiclient.NewSimplePublicObjectId("Id_example")}) // BatchReadInputSimplePublicObjectId | 
    archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostCrmV3ObjectsProductsBatchReadRead(context.Background()).BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV3ObjectsProductsBatchReadRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ObjectsProductsBatchReadRead`: BatchResponseSimplePublicObject
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.PostCrmV3ObjectsProductsBatchReadRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsProductsBatchReadReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchReadInputSimplePublicObjectId** | [**BatchReadInputSimplePublicObjectId**](BatchReadInputSimplePublicObjectId.md) |  | 
 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ObjectsProductsBatchUpdateUpdate

> BatchResponseSimplePublicObject PostCrmV3ObjectsProductsBatchUpdateUpdate(ctx).BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput).Execute()

Update a batch of products

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
    batchInputSimplePublicObjectBatchInput := *openapiclient.NewBatchInputSimplePublicObjectBatchInput([]openapiclient.SimplePublicObjectBatchInput{*openapiclient.NewSimplePublicObjectBatchInput(map[string]string{"key": "Inner_example"}, "Id_example")}) // BatchInputSimplePublicObjectBatchInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostCrmV3ObjectsProductsBatchUpdateUpdate(context.Background()).BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV3ObjectsProductsBatchUpdateUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ObjectsProductsBatchUpdateUpdate`: BatchResponseSimplePublicObject
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.PostCrmV3ObjectsProductsBatchUpdateUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsProductsBatchUpdateUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputSimplePublicObjectBatchInput** | [**BatchInputSimplePublicObjectBatchInput**](BatchInputSimplePublicObjectBatchInput.md) |  | 

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

