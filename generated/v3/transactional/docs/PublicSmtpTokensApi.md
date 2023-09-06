# \PublicSmtpTokensApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveToken**](PublicSmtpTokensApi.md#ArchiveToken) | **Delete** /marketing/v3/transactional/smtp-tokens/{tokenId} | Delete a single token by ID.
[**CreateToken**](PublicSmtpTokensApi.md#CreateToken) | **Post** /marketing/v3/transactional/smtp-tokens | Create a SMTP API token.
[**GetTokenByID**](PublicSmtpTokensApi.md#GetTokenByID) | **Get** /marketing/v3/transactional/smtp-tokens/{tokenId} | Query a single token by ID.
[**GetTokensPage**](PublicSmtpTokensApi.md#GetTokensPage) | **Get** /marketing/v3/transactional/smtp-tokens | Query SMTP API tokens by campaign name or an emailCampaignId.
[**ResetPassword**](PublicSmtpTokensApi.md#ResetPassword) | **Post** /marketing/v3/transactional/smtp-tokens/{tokenId}/password-reset | Reset the password of an existing token.



## ArchiveToken

> ArchiveToken(ctx, tokenId).Execute()

Delete a single token by ID.



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
    tokenId := "tokenId_example" // string | Identifier generated when a token is created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicSmtpTokensApi.ArchiveToken(context.Background(), tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSmtpTokensApi.ArchiveToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | Identifier generated when a token is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateToken

> SmtpApiTokenView CreateToken(ctx).SmtpApiTokenRequestEgg(smtpApiTokenRequestEgg).Execute()

Create a SMTP API token.



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
    smtpApiTokenRequestEgg := *openapiclient.NewSmtpApiTokenRequestEgg(false, "CampaignName_example") // SmtpApiTokenRequestEgg | A request object that includes the campaign name tied to the token and whether contacts should be created for email recipients.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicSmtpTokensApi.CreateToken(context.Background()).SmtpApiTokenRequestEgg(smtpApiTokenRequestEgg).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSmtpTokensApi.CreateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateToken`: SmtpApiTokenView
    fmt.Fprintf(os.Stdout, "Response from `PublicSmtpTokensApi.CreateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smtpApiTokenRequestEgg** | [**SmtpApiTokenRequestEgg**](SmtpApiTokenRequestEgg.md) | A request object that includes the campaign name tied to the token and whether contacts should be created for email recipients. | 

### Return type

[**SmtpApiTokenView**](SmtpApiTokenView.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenByID

> SmtpApiTokenView GetTokenByID(ctx, tokenId).Execute()

Query a single token by ID.



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
    tokenId := "tokenId_example" // string | Identifier generated when a token is created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicSmtpTokensApi.GetTokenByID(context.Background(), tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSmtpTokensApi.GetTokenByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenByID`: SmtpApiTokenView
    fmt.Fprintf(os.Stdout, "Response from `PublicSmtpTokensApi.GetTokenByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | Identifier generated when a token is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SmtpApiTokenView**](SmtpApiTokenView.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokensPage

> CollectionResponseSmtpApiTokenViewForwardPaging GetTokensPage(ctx).CampaignName(campaignName).EmailCampaignId(emailCampaignId).After(after).Limit(limit).Execute()

Query SMTP API tokens by campaign name or an emailCampaignId.



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
    campaignName := "campaignName_example" // string | A name for the campaign tied to the SMTP API token. (optional)
    emailCampaignId := "emailCampaignId_example" // string | Identifier assigned to the campaign provided during the token creation. (optional)
    after := "after_example" // string | Starting point to get the next set of results. (optional)
    limit := int32(56) // int32 | Maximum number of tokens to return. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicSmtpTokensApi.GetTokensPage(context.Background()).CampaignName(campaignName).EmailCampaignId(emailCampaignId).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSmtpTokensApi.GetTokensPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokensPage`: CollectionResponseSmtpApiTokenViewForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `PublicSmtpTokensApi.GetTokensPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokensPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignName** | **string** | A name for the campaign tied to the SMTP API token. | 
 **emailCampaignId** | **string** | Identifier assigned to the campaign provided during the token creation. | 
 **after** | **string** | Starting point to get the next set of results. | 
 **limit** | **int32** | Maximum number of tokens to return. | 

### Return type

[**CollectionResponseSmtpApiTokenViewForwardPaging**](CollectionResponseSmtpApiTokenViewForwardPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPassword

> SmtpApiTokenView ResetPassword(ctx, tokenId).Execute()

Reset the password of an existing token.



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
    tokenId := "tokenId_example" // string | Identifier generated when a token is created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicSmtpTokensApi.ResetPassword(context.Background(), tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSmtpTokensApi.ResetPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetPassword`: SmtpApiTokenView
    fmt.Fprintf(os.Stdout, "Response from `PublicSmtpTokensApi.ResetPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | Identifier generated when a token is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SmtpApiTokenView**](SmtpApiTokenView.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

