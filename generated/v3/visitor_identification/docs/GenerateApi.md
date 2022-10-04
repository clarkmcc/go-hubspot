# \GenerateApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostVisitorIdentificationV3TokensCreateGenerateToken**](GenerateApi.md#PostVisitorIdentificationV3TokensCreateGenerateToken) | **Post** /conversations/v3/visitor-identification/tokens/create | Generate a token



## PostVisitorIdentificationV3TokensCreateGenerateToken

> IdentificationTokenResponse PostVisitorIdentificationV3TokensCreateGenerateToken(ctx).IdentificationTokenGenerationRequest(identificationTokenGenerationRequest).Execute()

Generate a token



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
    identificationTokenGenerationRequest := *openapiclient.NewIdentificationTokenGenerationRequest("Email_example") // IdentificationTokenGenerationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenerateApi.PostVisitorIdentificationV3TokensCreateGenerateToken(context.Background()).IdentificationTokenGenerationRequest(identificationTokenGenerationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenerateApi.PostVisitorIdentificationV3TokensCreateGenerateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostVisitorIdentificationV3TokensCreateGenerateToken`: IdentificationTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `GenerateApi.PostVisitorIdentificationV3TokensCreateGenerateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostVisitorIdentificationV3TokensCreateGenerateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identificationTokenGenerationRequest** | [**IdentificationTokenGenerationRequest**](IdentificationTokenGenerationRequest.md) |  | 

### Return type

[**IdentificationTokenResponse**](IdentificationTokenResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

