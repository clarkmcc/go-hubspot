# \PipelinesApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Archive**](PipelinesApi.md#Archive) | **Delete** /crm/v3/pipelines/{objectType}/{pipelineId} | Delete a pipeline
[**Create**](PipelinesApi.md#Create) | **Post** /crm/v3/pipelines/{objectType} | Create a pipeline
[**GetAll**](PipelinesApi.md#GetAll) | **Get** /crm/v3/pipelines/{objectType} | Retrieve all pipelines
[**GetByID**](PipelinesApi.md#GetByID) | **Get** /crm/v3/pipelines/{objectType}/{pipelineId} | Return a pipeline by ID
[**Replace**](PipelinesApi.md#Replace) | **Put** /crm/v3/pipelines/{objectType}/{pipelineId} | Replace a pipeline
[**Update**](PipelinesApi.md#Update) | **Patch** /crm/v3/pipelines/{objectType}/{pipelineId} | Update a pipeline



## Archive

> Archive(ctx, objectType, pipelineId).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).Execute()

Delete a pipeline



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
    pipelineId := "pipelineId_example" // string | 
    validateReferencesBeforeDelete := true // bool |  (optional) (default to false)
    validateDealStageUsagesBeforeDelete := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelinesApi.Archive(context.Background(), objectType, pipelineId).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.Archive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateReferencesBeforeDelete** | **bool** |  | [default to false]
 **validateDealStageUsagesBeforeDelete** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> Pipeline Create(ctx, objectType).PipelineInput(pipelineInput).Execute()

Create a pipeline



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
    pipelineInput := *openapiclient.NewPipelineInput("Label_example", int32(123), []openapiclient.PipelineStageInput{*openapiclient.NewPipelineStageInput("Label_example", int32(123), map[string]string{"key": "Inner_example"})}) // PipelineInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelinesApi.Create(context.Background(), objectType).PipelineInput(pipelineInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pipelineInput** | [**PipelineInput**](PipelineInput.md) |  | 

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAll

> CollectionResponsePipelineNoPaging GetAll(ctx, objectType).Execute()

Retrieve all pipelines



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelinesApi.GetAll(context.Background(), objectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAll`: CollectionResponsePipelineNoPaging
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionResponsePipelineNoPaging**](CollectionResponsePipelineNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByID

> Pipeline GetByID(ctx, objectType, pipelineId).Execute()

Return a pipeline by ID



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
    pipelineId := "pipelineId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelinesApi.GetByID(context.Background(), objectType, pipelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetByID`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Replace

> Pipeline Replace(ctx, objectType, pipelineId).PipelineInput(pipelineInput).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).Execute()

Replace a pipeline



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
    pipelineId := "pipelineId_example" // string | 
    pipelineInput := *openapiclient.NewPipelineInput("Label_example", int32(123), []openapiclient.PipelineStageInput{*openapiclient.NewPipelineStageInput("Label_example", int32(123), map[string]string{"key": "Inner_example"})}) // PipelineInput | 
    validateReferencesBeforeDelete := true // bool |  (optional) (default to false)
    validateDealStageUsagesBeforeDelete := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelinesApi.Replace(context.Background(), objectType, pipelineId).PipelineInput(pipelineInput).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.Replace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Replace`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.Replace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pipelineInput** | [**PipelineInput**](PipelineInput.md) |  | 
 **validateReferencesBeforeDelete** | **bool** |  | [default to false]
 **validateDealStageUsagesBeforeDelete** | **bool** |  | [default to false]

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Pipeline Update(ctx, objectType, pipelineId).PipelinePatchInput(pipelinePatchInput).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).Execute()

Update a pipeline



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
    pipelineId := "pipelineId_example" // string | 
    pipelinePatchInput := *openapiclient.NewPipelinePatchInput() // PipelinePatchInput | 
    validateReferencesBeforeDelete := true // bool |  (optional) (default to false)
    validateDealStageUsagesBeforeDelete := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelinesApi.Update(context.Background(), objectType, pipelineId).PipelinePatchInput(pipelinePatchInput).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pipelinePatchInput** | [**PipelinePatchInput**](PipelinePatchInput.md) |  | 
 **validateReferencesBeforeDelete** | **bool** |  | [default to false]
 **validateDealStageUsagesBeforeDelete** | **bool** |  | [default to false]

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

