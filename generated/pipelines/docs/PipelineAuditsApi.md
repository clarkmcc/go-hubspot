# \PipelineAuditsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCrmV3PipelinesObjectTypePipelineIdAuditGetAudit**](PipelineAuditsApi.md#GetCrmV3PipelinesObjectTypePipelineIdAuditGetAudit) | **Get** /crm/v3/pipelines/{objectType}/{pipelineId}/audit | Return an audit of all changes to the pipeline



## GetCrmV3PipelinesObjectTypePipelineIdAuditGetAudit

> CollectionResponsePublicAuditInfoNoPaging GetCrmV3PipelinesObjectTypePipelineIdAuditGetAudit(ctx, objectType, pipelineId).Execute()

Return an audit of all changes to the pipeline



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
    resp, r, err := apiClient.PipelineAuditsApi.GetCrmV3PipelinesObjectTypePipelineIdAuditGetAudit(context.Background(), objectType, pipelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineAuditsApi.GetCrmV3PipelinesObjectTypePipelineIdAuditGetAudit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3PipelinesObjectTypePipelineIdAuditGetAudit`: CollectionResponsePublicAuditInfoNoPaging
    fmt.Fprintf(os.Stdout, "Response from `PipelineAuditsApi.GetCrmV3PipelinesObjectTypePipelineIdAuditGetAudit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3PipelinesObjectTypePipelineIdAuditGetAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CollectionResponsePublicAuditInfoNoPaging**](CollectionResponsePublicAuditInfoNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

