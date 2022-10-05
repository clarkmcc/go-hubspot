# \SampleResponseApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CardsGetSample**](SampleResponseApi.md#CardsGetSample) | **Get** /crm/v3/extensions/cards/sample-response | Get sample card detail response



## CardsGetSample

> IntegratorCardPayloadResponse CardsGetSample(ctx).Execute()

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
    resp, r, err := apiClient.SampleResponseApi.CardsGetSample(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SampleResponseApi.CardsGetSample``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CardsGetSample`: IntegratorCardPayloadResponse
    fmt.Fprintf(os.Stdout, "Response from `SampleResponseApi.CardsGetSample`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCardsGetSampleRequest struct via the builder pattern


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

