# \PipelineStagesApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3PipelinesObjectTypePipelineIdStagesStageIdArchive**](PipelineStagesApi.md#DeleteCrmV3PipelinesObjectTypePipelineIdStagesStageIdArchive) | **Delete** /crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId} | Archive a pipeline stage
[**GetCrmV3PipelinesObjectTypePipelineIdStagesGetAll**](PipelineStagesApi.md#GetCrmV3PipelinesObjectTypePipelineIdStagesGetAll) | **Get** /crm/v3/pipelines/{objectType}/{pipelineId}/stages | Return all stages of a pipeline
[**GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdGetById**](PipelineStagesApi.md#GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdGetById) | **Get** /crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId} | Return a pipeline stage by ID
[**PatchCrmV3PipelinesObjectTypePipelineIdStagesStageIdUpdate**](PipelineStagesApi.md#PatchCrmV3PipelinesObjectTypePipelineIdStagesStageIdUpdate) | **Patch** /crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId} | Update a pipeline stage
[**PostCrmV3PipelinesObjectTypePipelineIdStagesCreate**](PipelineStagesApi.md#PostCrmV3PipelinesObjectTypePipelineIdStagesCreate) | **Post** /crm/v3/pipelines/{objectType}/{pipelineId}/stages | Create a pipeline stage
[**PutCrmV3PipelinesObjectTypePipelineIdStagesStageIdReplace**](PipelineStagesApi.md#PutCrmV3PipelinesObjectTypePipelineIdStagesStageIdReplace) | **Put** /crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId} | Replace a pipeline stage



## DeleteCrmV3PipelinesObjectTypePipelineIdStagesStageIdArchive

> DeleteCrmV3PipelinesObjectTypePipelineIdStagesStageIdArchive(ctx, objectType, pipelineId, stageId).Execute()

Archive a pipeline stage



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelineStagesApi.DeleteCrmV3PipelinesObjectTypePipelineIdStagesStageIdArchive(context.Background(), objectType, pipelineId, stageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineStagesApi.DeleteCrmV3PipelinesObjectTypePipelineIdStagesStageIdArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCrmV3PipelinesObjectTypePipelineIdStagesStageIdArchiveRequest struct via the builder pattern


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


## GetCrmV3PipelinesObjectTypePipelineIdStagesGetAll

> CollectionResponsePipelineStage GetCrmV3PipelinesObjectTypePipelineIdStagesGetAll(ctx, objectType, pipelineId).Archived(archived).Execute()

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
    archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelineStagesApi.GetCrmV3PipelinesObjectTypePipelineIdStagesGetAll(context.Background(), objectType, pipelineId).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineStagesApi.GetCrmV3PipelinesObjectTypePipelineIdStagesGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3PipelinesObjectTypePipelineIdStagesGetAll`: CollectionResponsePipelineStage
    fmt.Fprintf(os.Stdout, "Response from `PipelineStagesApi.GetCrmV3PipelinesObjectTypePipelineIdStagesGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3PipelinesObjectTypePipelineIdStagesGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

[**CollectionResponsePipelineStage**](CollectionResponsePipelineStage.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdGetById

> PipelineStage GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdGetById(ctx, objectType, pipelineId, stageId).Archived(archived).Execute()

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
    archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelineStagesApi.GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdGetById(context.Background(), objectType, pipelineId, stageId).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineStagesApi.GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdGetById`: PipelineStage
    fmt.Fprintf(os.Stdout, "Response from `PipelineStagesApi.GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdGetById`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetCrmV3PipelinesObjectTypePipelineIdStagesStageIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

[**PipelineStage**](PipelineStage.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCrmV3PipelinesObjectTypePipelineIdStagesStageIdUpdate

> PipelineStage PatchCrmV3PipelinesObjectTypePipelineIdStagesStageIdUpdate(ctx, objectType, pipelineId, stageId).Archived(archived).PipelineStagePatchInput(pipelineStagePatchInput).Execute()

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
    archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)
    pipelineStagePatchInput := *openapiclient.NewPipelineStagePatchInput(map[string]string{"key": "Inner_example"}) // PipelineStagePatchInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelineStagesApi.PatchCrmV3PipelinesObjectTypePipelineIdStagesStageIdUpdate(context.Background(), objectType, pipelineId, stageId).Archived(archived).PipelineStagePatchInput(pipelineStagePatchInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineStagesApi.PatchCrmV3PipelinesObjectTypePipelineIdStagesStageIdUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCrmV3PipelinesObjectTypePipelineIdStagesStageIdUpdate`: PipelineStage
    fmt.Fprintf(os.Stdout, "Response from `PipelineStagesApi.PatchCrmV3PipelinesObjectTypePipelineIdStagesStageIdUpdate`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPatchCrmV3PipelinesObjectTypePipelineIdStagesStageIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]
 **pipelineStagePatchInput** | [**PipelineStagePatchInput**](PipelineStagePatchInput.md) |  | 

### Return type

[**PipelineStage**](PipelineStage.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3PipelinesObjectTypePipelineIdStagesCreate

> PipelineStage PostCrmV3PipelinesObjectTypePipelineIdStagesCreate(ctx, objectType, pipelineId).PipelineStageInput(pipelineStageInput).Execute()

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
    pipelineStageInput := *openapiclient.NewPipelineStageInput("Label_example", int32(123), map[string]string{"key": "Inner_example"}) // PipelineStageInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelineStagesApi.PostCrmV3PipelinesObjectTypePipelineIdStagesCreate(context.Background(), objectType, pipelineId).PipelineStageInput(pipelineStageInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineStagesApi.PostCrmV3PipelinesObjectTypePipelineIdStagesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3PipelinesObjectTypePipelineIdStagesCreate`: PipelineStage
    fmt.Fprintf(os.Stdout, "Response from `PipelineStagesApi.PostCrmV3PipelinesObjectTypePipelineIdStagesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3PipelinesObjectTypePipelineIdStagesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pipelineStageInput** | [**PipelineStageInput**](PipelineStageInput.md) |  | 

### Return type

[**PipelineStage**](PipelineStage.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmV3PipelinesObjectTypePipelineIdStagesStageIdReplace

> PipelineStage PutCrmV3PipelinesObjectTypePipelineIdStagesStageIdReplace(ctx, objectType, pipelineId, stageId).PipelineStageInput(pipelineStageInput).Execute()

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
    pipelineStageInput := *openapiclient.NewPipelineStageInput("Label_example", int32(123), map[string]string{"key": "Inner_example"}) // PipelineStageInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelineStagesApi.PutCrmV3PipelinesObjectTypePipelineIdStagesStageIdReplace(context.Background(), objectType, pipelineId, stageId).PipelineStageInput(pipelineStageInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineStagesApi.PutCrmV3PipelinesObjectTypePipelineIdStagesStageIdReplace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutCrmV3PipelinesObjectTypePipelineIdStagesStageIdReplace`: PipelineStage
    fmt.Fprintf(os.Stdout, "Response from `PipelineStagesApi.PutCrmV3PipelinesObjectTypePipelineIdStagesStageIdReplace`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPutCrmV3PipelinesObjectTypePipelineIdStagesStageIdReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pipelineStageInput** | [**PipelineStageInput**](PipelineStageInput.md) |  | 

### Return type

[**PipelineStage**](PipelineStage.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

