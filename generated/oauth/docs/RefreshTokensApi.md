# \RefreshTokensApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOauthV1RefreshTokensTokenArchiveRefreshToken**](RefreshTokensApi.md#DeleteOauthV1RefreshTokensTokenArchiveRefreshToken) | **Delete** /oauth/v1/refresh-tokens/{token} | 
[**GetOauthV1RefreshTokensTokenGetRefreshToken**](RefreshTokensApi.md#GetOauthV1RefreshTokensTokenGetRefreshToken) | **Get** /oauth/v1/refresh-tokens/{token} | 



## DeleteOauthV1RefreshTokensTokenArchiveRefreshToken

> Error DeleteOauthV1RefreshTokensTokenArchiveRefreshToken(ctx, token).Execute()



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RefreshTokensApi.DeleteOauthV1RefreshTokensTokenArchiveRefreshToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefreshTokensApi.DeleteOauthV1RefreshTokensTokenArchiveRefreshToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOauthV1RefreshTokensTokenArchiveRefreshToken`: Error
    fmt.Fprintf(os.Stdout, "Response from `RefreshTokensApi.DeleteOauthV1RefreshTokensTokenArchiveRefreshToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOauthV1RefreshTokensTokenArchiveRefreshTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Error**](Error.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOauthV1RefreshTokensTokenGetRefreshToken

> RefreshTokenInfoResponse GetOauthV1RefreshTokensTokenGetRefreshToken(ctx, token).Execute()



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RefreshTokensApi.GetOauthV1RefreshTokensTokenGetRefreshToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefreshTokensApi.GetOauthV1RefreshTokensTokenGetRefreshToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOauthV1RefreshTokensTokenGetRefreshToken`: RefreshTokenInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `RefreshTokensApi.GetOauthV1RefreshTokensTokenGetRefreshToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthV1RefreshTokensTokenGetRefreshTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RefreshTokenInfoResponse**](RefreshTokenInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

