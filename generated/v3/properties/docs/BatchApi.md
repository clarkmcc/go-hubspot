# \BatchApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmV3PropertiesObjectTypeBatchArchiveArchive**](BatchApi.md#PostCrmV3PropertiesObjectTypeBatchArchiveArchive) | **Post** /crm/v3/properties/{objectType}/batch/archive | Archive a batch of properties
[**PostCrmV3PropertiesObjectTypeBatchCreateCreate**](BatchApi.md#PostCrmV3PropertiesObjectTypeBatchCreateCreate) | **Post** /crm/v3/properties/{objectType}/batch/create | Create a batch of properties
[**PostCrmV3PropertiesObjectTypeBatchReadRead**](BatchApi.md#PostCrmV3PropertiesObjectTypeBatchReadRead) | **Post** /crm/v3/properties/{objectType}/batch/read | Read a batch of properties



## PostCrmV3PropertiesObjectTypeBatchArchiveArchive

> PostCrmV3PropertiesObjectTypeBatchArchiveArchive(ctx, objectType).BatchInputPropertyName(batchInputPropertyName).Execute()

Archive a batch of properties



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
    objectType := "objectType_example" // string | 
    batchInputPropertyName := *openapiclient.NewBatchInputPropertyName([]openapiclient.PropertyName{*openapiclient.NewPropertyName("Name_example")}) // BatchInputPropertyName | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostCrmV3PropertiesObjectTypeBatchArchiveArchive(context.Background(), objectType).BatchInputPropertyName(batchInputPropertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV3PropertiesObjectTypeBatchArchiveArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3PropertiesObjectTypeBatchArchiveArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputPropertyName** | [**BatchInputPropertyName**](BatchInputPropertyName.md) |  | 

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


## PostCrmV3PropertiesObjectTypeBatchCreateCreate

> BatchResponseProperty PostCrmV3PropertiesObjectTypeBatchCreateCreate(ctx, objectType).BatchInputPropertyCreate(batchInputPropertyCreate).Execute()

Create a batch of properties



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
    objectType := "objectType_example" // string | 
    batchInputPropertyCreate := *openapiclient.NewBatchInputPropertyCreate([]openapiclient.PropertyCreate{*openapiclient.NewPropertyCreate("Name_example", "Label_example", "Type_example", "FieldType_example", "GroupName_example")}) // BatchInputPropertyCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostCrmV3PropertiesObjectTypeBatchCreateCreate(context.Background(), objectType).BatchInputPropertyCreate(batchInputPropertyCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV3PropertiesObjectTypeBatchCreateCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3PropertiesObjectTypeBatchCreateCreate`: BatchResponseProperty
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.PostCrmV3PropertiesObjectTypeBatchCreateCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3PropertiesObjectTypeBatchCreateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputPropertyCreate** | [**BatchInputPropertyCreate**](BatchInputPropertyCreate.md) |  | 

### Return type

[**BatchResponseProperty**](BatchResponseProperty.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3PropertiesObjectTypeBatchReadRead

> BatchResponseProperty PostCrmV3PropertiesObjectTypeBatchReadRead(ctx, objectType).BatchReadInputPropertyName(batchReadInputPropertyName).Execute()

Read a batch of properties



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
    objectType := "objectType_example" // string | 
    batchReadInputPropertyName := *openapiclient.NewBatchReadInputPropertyName([]openapiclient.PropertyName{*openapiclient.NewPropertyName("Name_example")}, false) // BatchReadInputPropertyName | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostCrmV3PropertiesObjectTypeBatchReadRead(context.Background(), objectType).BatchReadInputPropertyName(batchReadInputPropertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV3PropertiesObjectTypeBatchReadRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3PropertiesObjectTypeBatchReadRead`: BatchResponseProperty
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.PostCrmV3PropertiesObjectTypeBatchReadRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3PropertiesObjectTypeBatchReadReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchReadInputPropertyName** | [**BatchReadInputPropertyName**](BatchReadInputPropertyName.md) |  | 

### Return type

[**BatchResponseProperty**](BatchResponseProperty.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

