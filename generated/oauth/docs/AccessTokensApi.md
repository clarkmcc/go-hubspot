# \AccessTokensApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOauthV1AccessTokensTokenGetAccessToken**](AccessTokensApi.md#GetOauthV1AccessTokensTokenGetAccessToken) | **Get** /oauth/v1/access-tokens/{token} | 



## GetOauthV1AccessTokensTokenGetAccessToken

> AccessTokenInfoResponse GetOauthV1AccessTokensTokenGetAccessToken(ctx, token).Execute()



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
    token := "token_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessTokensApi.GetOauthV1AccessTokensTokenGetAccessToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensApi.GetOauthV1AccessTokensTokenGetAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOauthV1AccessTokensTokenGetAccessToken`: AccessTokenInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessTokensApi.GetOauthV1AccessTokensTokenGetAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthV1AccessTokensTokenGetAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessTokenInfoResponse**](AccessTokenInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

