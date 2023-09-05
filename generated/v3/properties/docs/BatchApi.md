# \BatchApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchArchive**](BatchApi.md#BatchArchive) | **Post** /crm/v3/properties/{objectType}/batch/archive | Archive a batch of properties
[**BatchCreate**](BatchApi.md#BatchCreate) | **Post** /crm/v3/properties/{objectType}/batch/create | Create a batch of properties
[**BatchRead**](BatchApi.md#BatchRead) | **Post** /crm/v3/properties/{objectType}/batch/read | Read a batch of properties



## BatchArchive

> BatchArchive(ctx, objectType).BatchInputPropertyName(batchInputPropertyName).Execute()

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
    resp, r, err := apiClient.BatchApi.BatchArchive(context.Background(), objectType).BatchInputPropertyName(batchInputPropertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBatchArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputPropertyName** | [**BatchInputPropertyName**](BatchInputPropertyName.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchCreate

> BatchResponseProperty BatchCreate(ctx, objectType).BatchInputPropertyCreate(batchInputPropertyCreate).Execute()

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
    resp, r, err := apiClient.BatchApi.BatchCreate(context.Background(), objectType).BatchInputPropertyCreate(batchInputPropertyCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchCreate`: BatchResponseProperty
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.BatchCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputPropertyCreate** | [**BatchInputPropertyCreate**](BatchInputPropertyCreate.md) |  | 

### Return type

[**BatchResponseProperty**](BatchResponseProperty.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchRead

> BatchResponseProperty BatchRead(ctx, objectType).BatchReadInputPropertyName(batchReadInputPropertyName).Execute()

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
    batchReadInputPropertyName := *openapiclient.NewBatchReadInputPropertyName(false, []openapiclient.PropertyName{*openapiclient.NewPropertyName("Name_example")}) // BatchReadInputPropertyName | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.BatchRead(context.Background(), objectType).BatchReadInputPropertyName(batchReadInputPropertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchRead`: BatchResponseProperty
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.BatchRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchReadInputPropertyName** | [**BatchReadInputPropertyName**](BatchReadInputPropertyName.md) |  | 

### Return type

[**BatchResponseProperty**](BatchResponseProperty.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

