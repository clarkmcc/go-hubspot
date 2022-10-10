# \DefinitionApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPage**](DefinitionApi.md#GetPage) | **Get** /communication-preferences/v3/definitions | Get subscription definitions



## GetPage

> SubscriptionDefinitionsResponse GetPage(ctx).Execute()

Get subscription definitions



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
    resp, r, err := apiClient.DefinitionApi.GetPage(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefinitionApi.GetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPage`: SubscriptionDefinitionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefinitionApi.GetPage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageRequest struct via the builder pattern


### Return type

[**SubscriptionDefinitionsResponse**](SubscriptionDefinitionsResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

