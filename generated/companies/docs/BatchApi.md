# \BatchApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmV3ObjectsCompaniesBatchArchiveArchive**](BatchApi.md#PostCrmV3ObjectsCompaniesBatchArchiveArchive) | **Post** /crm/v3/objects/companies/batch/archive | Archive a batch of companies by ID
[**PostCrmV3ObjectsCompaniesBatchCreateCreate**](BatchApi.md#PostCrmV3ObjectsCompaniesBatchCreateCreate) | **Post** /crm/v3/objects/companies/batch/create | Create a batch of companies
[**PostCrmV3ObjectsCompaniesBatchReadRead**](BatchApi.md#PostCrmV3ObjectsCompaniesBatchReadRead) | **Post** /crm/v3/objects/companies/batch/read | Read a batch of companies by internal ID, or unique property values
[**PostCrmV3ObjectsCompaniesBatchUpdateUpdate**](BatchApi.md#PostCrmV3ObjectsCompaniesBatchUpdateUpdate) | **Post** /crm/v3/objects/companies/batch/update | Update a batch of companies



## PostCrmV3ObjectsCompaniesBatchArchiveArchive

> PostCrmV3ObjectsCompaniesBatchArchiveArchive(ctx).BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId).Execute()

Archive a batch of companies by ID

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
    resp, r, err := apiClient.BatchApi.PostCrmV3ObjectsCompaniesBatchArchiveArchive(context.Background()).BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV3ObjectsCompaniesBatchArchiveArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsCompaniesBatchArchiveArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputSimplePublicObjectId** | [**BatchInputSimplePublicObjectId**](BatchInputSimplePublicObjectId.md) |  | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ObjectsCompaniesBatchCreateCreate

> BatchResponseSimplePublicObject PostCrmV3ObjectsCompaniesBatchCreateCreate(ctx).BatchInputSimplePublicObjectInput(batchInputSimplePublicObjectInput).Execute()

Create a batch of companies

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
    resp, r, err := apiClient.BatchApi.PostCrmV3ObjectsCompaniesBatchCreateCreate(context.Background()).BatchInputSimplePublicObjectInput(batchInputSimplePublicObjectInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV3ObjectsCompaniesBatchCreateCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ObjectsCompaniesBatchCreateCreate`: BatchResponseSimplePublicObject
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.PostCrmV3ObjectsCompaniesBatchCreateCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsCompaniesBatchCreateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputSimplePublicObjectInput** | [**BatchInputSimplePublicObjectInput**](BatchInputSimplePublicObjectInput.md) |  | 

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ObjectsCompaniesBatchReadRead

> BatchResponseSimplePublicObject PostCrmV3ObjectsCompaniesBatchReadRead(ctx).BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId).Archived(archived).Execute()

Read a batch of companies by internal ID, or unique property values

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
    resp, r, err := apiClient.BatchApi.PostCrmV3ObjectsCompaniesBatchReadRead(context.Background()).BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV3ObjectsCompaniesBatchReadRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ObjectsCompaniesBatchReadRead`: BatchResponseSimplePublicObject
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.PostCrmV3ObjectsCompaniesBatchReadRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsCompaniesBatchReadReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchReadInputSimplePublicObjectId** | [**BatchReadInputSimplePublicObjectId**](BatchReadInputSimplePublicObjectId.md) |  | 
 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ObjectsCompaniesBatchUpdateUpdate

> BatchResponseSimplePublicObject PostCrmV3ObjectsCompaniesBatchUpdateUpdate(ctx).BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput).Execute()

Update a batch of companies

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
    resp, r, err := apiClient.BatchApi.PostCrmV3ObjectsCompaniesBatchUpdateUpdate(context.Background()).BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV3ObjectsCompaniesBatchUpdateUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ObjectsCompaniesBatchUpdateUpdate`: BatchResponseSimplePublicObject
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.PostCrmV3ObjectsCompaniesBatchUpdateUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsCompaniesBatchUpdateUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputSimplePublicObjectBatchInput** | [**BatchInputSimplePublicObjectBatchInput**](BatchInputSimplePublicObjectBatchInput.md) |  | 

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

