# \PipelineStageAuditsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StagesGetAudit**](PipelineStageAuditsApi.md#StagesGetAudit) | **Get** /crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId}/audit | Return an audit of all changes to the pipeline stage



## StagesGetAudit

> CollectionResponsePublicAuditInfoNoPaging StagesGetAudit(ctx, objectType, stageId).Execute()

Return an audit of all changes to the pipeline stage



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
    stageId := "stageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineStageAuditsApi.StagesGetAudit(context.Background(), objectType, stageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineStageAuditsApi.StagesGetAudit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesGetAudit`: CollectionResponsePublicAuditInfoNoPaging
    fmt.Fprintf(os.Stdout, "Response from `PipelineStageAuditsApi.StagesGetAudit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesGetAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CollectionResponsePublicAuditInfoNoPaging**](CollectionResponsePublicAuditInfoNoPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

