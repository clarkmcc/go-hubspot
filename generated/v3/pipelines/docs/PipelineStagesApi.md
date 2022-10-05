# \PipelineStagesApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StagesArchive**](PipelineStagesApi.md#StagesArchive) | **Delete** /crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId} | Delete a pipeline stage
[**StagesCreate**](PipelineStagesApi.md#StagesCreate) | **Post** /crm/v3/pipelines/{objectType}/{pipelineId}/stages | Create a pipeline stage
[**StagesGetAll**](PipelineStagesApi.md#StagesGetAll) | **Get** /crm/v3/pipelines/{objectType}/{pipelineId}/stages | Return all stages of a pipeline
[**StagesGetByID**](PipelineStagesApi.md#StagesGetByID) | **Get** /crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId} | Return a pipeline stage by ID
[**StagesReplace**](PipelineStagesApi.md#StagesReplace) | **Put** /crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId} | Replace a pipeline stage
[**StagesUpdate**](PipelineStagesApi.md#StagesUpdate) | **Patch** /crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId} | Update a pipeline stage



## StagesArchive

> StagesArchive(ctx, objectType, pipelineId, stageId).Execute()

Delete a pipeline stage



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
    stageId := "stageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineStagesApi.StagesArchive(context.Background(), objectType, pipelineId, stageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineStagesApi.StagesArchive``: %v\n", err)
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
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesCreate

> PipelineStage StagesCreate(ctx, objectType, pipelineId).PipelineStageInput(pipelineStageInput).Execute()

Create a pipeline stage



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
    pipelineStageInput := *openapiclient.NewPipelineStageInput("Label_example", int32(123), map[string]string{"key": "Inner_example"}) // PipelineStageInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineStagesApi.StagesCreate(context.Background(), objectType, pipelineId).PipelineStageInput(pipelineStageInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineStagesApi.StagesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesCreate`: PipelineStage
    fmt.Fprintf(os.Stdout, "Response from `PipelineStagesApi.StagesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pipelineStageInput** | [**PipelineStageInput**](PipelineStageInput.md) |  | 

### Return type

[**PipelineStage**](PipelineStage.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesGetAll

> CollectionResponsePipelineStageNoPaging StagesGetAll(ctx, objectType, pipelineId).Execute()

Return all stages of a pipeline



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
    resp, r, err := apiClient.PipelineStagesApi.StagesGetAll(context.Background(), objectType, pipelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineStagesApi.StagesGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesGetAll`: CollectionResponsePipelineStageNoPaging
    fmt.Fprintf(os.Stdout, "Response from `PipelineStagesApi.StagesGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CollectionResponsePipelineStageNoPaging**](CollectionResponsePipelineStageNoPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesGetByID

> PipelineStage StagesGetByID(ctx, objectType, pipelineId, stageId).Execute()

Return a pipeline stage by ID



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
    stageId := "stageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineStagesApi.StagesGetByID(context.Background(), objectType, pipelineId, stageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineStagesApi.StagesGetByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesGetByID`: PipelineStage
    fmt.Fprintf(os.Stdout, "Response from `PipelineStagesApi.StagesGetByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesGetByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PipelineStage**](PipelineStage.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesReplace

> PipelineStage StagesReplace(ctx, objectType, pipelineId, stageId).PipelineStageInput(pipelineStageInput).Execute()

Replace a pipeline stage



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
    stageId := "stageId_example" // string | 
    pipelineStageInput := *openapiclient.NewPipelineStageInput("Label_example", int32(123), map[string]string{"key": "Inner_example"}) // PipelineStageInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineStagesApi.StagesReplace(context.Background(), objectType, pipelineId, stageId).PipelineStageInput(pipelineStageInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineStagesApi.StagesReplace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesReplace`: PipelineStage
    fmt.Fprintf(os.Stdout, "Response from `PipelineStagesApi.StagesReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pipelineStageInput** | [**PipelineStageInput**](PipelineStageInput.md) |  | 

### Return type

[**PipelineStage**](PipelineStage.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesUpdate

> PipelineStage StagesUpdate(ctx, objectType, pipelineId, stageId).PipelineStagePatchInput(pipelineStagePatchInput).Execute()

Update a pipeline stage



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
    stageId := "stageId_example" // string | 
    pipelineStagePatchInput := *openapiclient.NewPipelineStagePatchInput(map[string]string{"key": "Inner_example"}) // PipelineStagePatchInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineStagesApi.StagesUpdate(context.Background(), objectType, pipelineId, stageId).PipelineStagePatchInput(pipelineStagePatchInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineStagesApi.StagesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUpdate`: PipelineStage
    fmt.Fprintf(os.Stdout, "Response from `PipelineStagesApi.StagesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pipelineStagePatchInput** | [**PipelineStagePatchInput**](PipelineStagePatchInput.md) |  | 

### Return type

[**PipelineStage**](PipelineStage.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

