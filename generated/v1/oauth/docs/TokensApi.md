# \TokensApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToken**](TokensApi.md#CreateToken) | **Post** /oauth/v1/token | 



## CreateToken

> TokenResponseIF CreateToken(ctx).GrantType(grantType).Code(code).RedirectUri(redirectUri).ClientId(clientId).ClientSecret(clientSecret).RefreshToken(refreshToken).Execute()



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
    grantType := "grantType_example" // string |  (optional)
    code := "code_example" // string |  (optional)
    redirectUri := "redirectUri_example" // string |  (optional)
    clientId := "clientId_example" // string |  (optional)
    clientSecret := "clientSecret_example" // string |  (optional)
    refreshToken := "refreshToken_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensApi.CreateToken(context.Background()).GrantType(grantType).Code(code).RedirectUri(redirectUri).ClientId(clientId).ClientSecret(clientSecret).RefreshToken(refreshToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.CreateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateToken`: TokenResponseIF
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.CreateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** |  | 
 **code** | **string** |  | 
 **redirectUri** | **string** |  | 
 **clientId** | **string** |  | 
 **clientSecret** | **string** |  | 
 **refreshToken** | **string** |  | 

### Return type

[**TokenResponseIF**](TokenResponseIF.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

