# \SampleResponseApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCrmV3ExtensionsCardsSampleResponseGetCardsSampleResponse**](SampleResponseApi.md#GetCrmV3ExtensionsCardsSampleResponseGetCardsSampleResponse) | **Get** /crm/v3/extensions/cards/sample-response | Get sample card detail response



## GetCrmV3ExtensionsCardsSampleResponseGetCardsSampleResponse

> IntegratorCardPayloadResponse GetCrmV3ExtensionsCardsSampleResponseGetCardsSampleResponse(ctx).Execute()

Get sample card detail response



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SampleResponseApi.GetCrmV3ExtensionsCardsSampleResponseGetCardsSampleResponse(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SampleResponseApi.GetCrmV3ExtensionsCardsSampleResponseGetCardsSampleResponse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3ExtensionsCardsSampleResponseGetCardsSampleResponse`: IntegratorCardPayloadResponse
    fmt.Fprintf(os.Stdout, "Response from `SampleResponseApi.GetCrmV3ExtensionsCardsSampleResponseGetCardsSampleResponse`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ExtensionsCardsSampleResponseGetCardsSampleResponseRequest struct via the builder pattern


### Return type

[**IntegratorCardPayloadResponse**](IntegratorCardPayloadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

