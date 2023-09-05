# \RefreshTokensApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOauthV1RefreshTokensTokenArchive**](RefreshTokensApi.md#DeleteOauthV1RefreshTokensTokenArchive) | **Delete** /oauth/v1/refresh-tokens/{token} | 
[**GetOauthV1RefreshTokensTokenGet**](RefreshTokensApi.md#GetOauthV1RefreshTokensTokenGet) | **Get** /oauth/v1/refresh-tokens/{token} | 



## DeleteOauthV1RefreshTokensTokenArchive

> DeleteOauthV1RefreshTokensTokenArchive(ctx, token).Execute()



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
    resp, r, err := apiClient.RefreshTokensApi.DeleteOauthV1RefreshTokensTokenArchive(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefreshTokensApi.DeleteOauthV1RefreshTokensTokenArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOauthV1RefreshTokensTokenArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOauthV1RefreshTokensTokenGet

> RefreshTokenInfoResponse GetOauthV1RefreshTokensTokenGet(ctx, token).Execute()



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
    resp, r, err := apiClient.RefreshTokensApi.GetOauthV1RefreshTokensTokenGet(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefreshTokensApi.GetOauthV1RefreshTokensTokenGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOauthV1RefreshTokensTokenGet`: RefreshTokenInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `RefreshTokensApi.GetOauthV1RefreshTokensTokenGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthV1RefreshTokensTokenGetRequest struct via the builder pattern


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

