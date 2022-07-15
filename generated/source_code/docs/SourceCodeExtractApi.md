# \SourceCodeExtractApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCmsV3SourceCodeExtractAsyncTasksTaskIdStatusGetAsyncStatus**](SourceCodeExtractApi.md#GetCmsV3SourceCodeExtractAsyncTasksTaskIdStatusGetAsyncStatus) | **Get** /cms/v3/source-code/extract/async/tasks/{taskId}/status | 
[**PostCmsV3SourceCodeExtractAsyncDoAsync**](SourceCodeExtractApi.md#PostCmsV3SourceCodeExtractAsyncDoAsync) | **Post** /cms/v3/source-code/extract/async | 



## GetCmsV3SourceCodeExtractAsyncTasksTaskIdStatusGetAsyncStatus

> ActionResponse GetCmsV3SourceCodeExtractAsyncTasksTaskIdStatusGetAsyncStatus(ctx, taskId).Execute()



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
    taskId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceCodeExtractApi.GetCmsV3SourceCodeExtractAsyncTasksTaskIdStatusGetAsyncStatus(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceCodeExtractApi.GetCmsV3SourceCodeExtractAsyncTasksTaskIdStatusGetAsyncStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3SourceCodeExtractAsyncTasksTaskIdStatusGetAsyncStatus`: ActionResponse
    fmt.Fprintf(os.Stdout, "Response from `SourceCodeExtractApi.GetCmsV3SourceCodeExtractAsyncTasksTaskIdStatusGetAsyncStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3SourceCodeExtractAsyncTasksTaskIdStatusGetAsyncStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionResponse**](ActionResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3SourceCodeExtractAsyncDoAsync

> TaskLocator PostCmsV3SourceCodeExtractAsyncDoAsync(ctx).FileExtractRequest(fileExtractRequest).Execute()



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
    fileExtractRequest := *openapiclient.NewFileExtractRequest("Path_example") // FileExtractRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceCodeExtractApi.PostCmsV3SourceCodeExtractAsyncDoAsync(context.Background()).FileExtractRequest(fileExtractRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceCodeExtractApi.PostCmsV3SourceCodeExtractAsyncDoAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3SourceCodeExtractAsyncDoAsync`: TaskLocator
    fmt.Fprintf(os.Stdout, "Response from `SourceCodeExtractApi.PostCmsV3SourceCodeExtractAsyncDoAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3SourceCodeExtractAsyncDoAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileExtractRequest** | [**FileExtractRequest**](FileExtractRequest.md) |  | 

### Return type

[**TaskLocator**](TaskLocator.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

