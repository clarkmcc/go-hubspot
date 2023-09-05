# \SourceCodeExtractApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExtractAsync**](SourceCodeExtractApi.md#ExtractAsync) | **Post** /cms/v3/source-code/extract/async | 
[**ExtractGetAsyncStatus**](SourceCodeExtractApi.md#ExtractGetAsyncStatus) | **Get** /cms/v3/source-code/extract/async/tasks/{taskId}/status | 



## ExtractAsync

> TaskLocator ExtractAsync(ctx).FileExtractRequest(fileExtractRequest).Execute()



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
    resp, r, err := apiClient.SourceCodeExtractApi.ExtractAsync(context.Background()).FileExtractRequest(fileExtractRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceCodeExtractApi.ExtractAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtractAsync`: TaskLocator
    fmt.Fprintf(os.Stdout, "Response from `SourceCodeExtractApi.ExtractAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExtractAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileExtractRequest** | [**FileExtractRequest**](FileExtractRequest.md) |  | 

### Return type

[**TaskLocator**](TaskLocator.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtractGetAsyncStatus

> ActionResponse ExtractGetAsyncStatus(ctx, taskId).Execute()



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
    resp, r, err := apiClient.SourceCodeExtractApi.ExtractGetAsyncStatus(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceCodeExtractApi.ExtractGetAsyncStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtractGetAsyncStatus`: ActionResponse
    fmt.Fprintf(os.Stdout, "Response from `SourceCodeExtractApi.ExtractGetAsyncStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtractGetAsyncStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionResponse**](ActionResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

