# \PublicImportsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCrmV3ImportsImportIdErrorsGetErrors**](PublicImportsApi.md#GetCrmV3ImportsImportIdErrorsGetErrors) | **Get** /crm/v3/imports/{importId}/errors | 



## GetCrmV3ImportsImportIdErrorsGetErrors

> CollectionResponsePublicImportErrorForwardPaging GetCrmV3ImportsImportIdErrorsGetErrors(ctx, importId).After(after).Limit(limit).Execute()



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
    importId := int64(789) // int64 | 
    after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
    limit := int32(56) // int32 | The maximum number of results to display per page. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicImportsApi.GetCrmV3ImportsImportIdErrorsGetErrors(context.Background(), importId).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicImportsApi.GetCrmV3ImportsImportIdErrorsGetErrors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3ImportsImportIdErrorsGetErrors`: CollectionResponsePublicImportErrorForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `PublicImportsApi.GetCrmV3ImportsImportIdErrorsGetErrors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**importId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ImportsImportIdErrorsGetErrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to display per page. | 

### Return type

[**CollectionResponsePublicImportErrorForwardPaging**](CollectionResponsePublicImportErrorForwardPaging.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

