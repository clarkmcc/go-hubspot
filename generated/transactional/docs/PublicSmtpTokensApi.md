# \PublicSmtpTokensApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMarketingV3TransactionalSmtpTokensTokenIdArchiveToken**](PublicSmtpTokensApi.md#DeleteMarketingV3TransactionalSmtpTokensTokenIdArchiveToken) | **Delete** /marketing/v3/transactional/smtp-tokens/{tokenId} | Delete a single token by ID.
[**GetMarketingV3TransactionalSmtpTokensGetTokensPage**](PublicSmtpTokensApi.md#GetMarketingV3TransactionalSmtpTokensGetTokensPage) | **Get** /marketing/v3/transactional/smtp-tokens | Query SMTP API tokens by campaign name or an emailCampaignId.
[**GetMarketingV3TransactionalSmtpTokensTokenIdGetTokenById**](PublicSmtpTokensApi.md#GetMarketingV3TransactionalSmtpTokensTokenIdGetTokenById) | **Get** /marketing/v3/transactional/smtp-tokens/{tokenId} | Query a single token by ID.
[**PostMarketingV3TransactionalSmtpTokensCreateToken**](PublicSmtpTokensApi.md#PostMarketingV3TransactionalSmtpTokensCreateToken) | **Post** /marketing/v3/transactional/smtp-tokens | Create a SMTP API token.
[**PostMarketingV3TransactionalSmtpTokensTokenIdPasswordResetResetPassword**](PublicSmtpTokensApi.md#PostMarketingV3TransactionalSmtpTokensTokenIdPasswordResetResetPassword) | **Post** /marketing/v3/transactional/smtp-tokens/{tokenId}/password-reset | Reset the password of an existing token.



## DeleteMarketingV3TransactionalSmtpTokensTokenIdArchiveToken

> DeleteMarketingV3TransactionalSmtpTokensTokenIdArchiveToken(ctx, tokenId).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicSmtpTokensApi.DeleteMarketingV3TransactionalSmtpTokensTokenIdArchiveToken(context.Background(), tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSmtpTokensApi.DeleteMarketingV3TransactionalSmtpTokensTokenIdArchiveToken``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMarketingV3TransactionalSmtpTokensTokenIdArchiveTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3TransactionalSmtpTokensGetTokensPage

> CollectionResponseSmtpApiTokenView GetMarketingV3TransactionalSmtpTokensGetTokensPage(ctx).CampaignName(campaignName).EmailCampaignId(emailCampaignId).After(after).Limit(limit).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicSmtpTokensApi.GetMarketingV3TransactionalSmtpTokensGetTokensPage(context.Background()).CampaignName(campaignName).EmailCampaignId(emailCampaignId).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSmtpTokensApi.GetMarketingV3TransactionalSmtpTokensGetTokensPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketingV3TransactionalSmtpTokensGetTokensPage`: CollectionResponseSmtpApiTokenView
    fmt.Fprintf(os.Stdout, "Response from `PublicSmtpTokensApi.GetMarketingV3TransactionalSmtpTokensGetTokensPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3TransactionalSmtpTokensGetTokensPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignName** | **string** | A name for the campaign tied to the SMTP API token. | 
 **emailCampaignId** | **string** | Identifier assigned to the campaign provided during the token creation. | 
 **after** | **string** | Starting point to get the next set of results. | 
 **limit** | **int32** | Maximum number of tokens to return. | 

### Return type

[**CollectionResponseSmtpApiTokenView**](CollectionResponseSmtpApiTokenView.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3TransactionalSmtpTokensTokenIdGetTokenById

> SmtpApiTokenView GetMarketingV3TransactionalSmtpTokensTokenIdGetTokenById(ctx, tokenId).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicSmtpTokensApi.GetMarketingV3TransactionalSmtpTokensTokenIdGetTokenById(context.Background(), tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSmtpTokensApi.GetMarketingV3TransactionalSmtpTokensTokenIdGetTokenById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketingV3TransactionalSmtpTokensTokenIdGetTokenById`: SmtpApiTokenView
    fmt.Fprintf(os.Stdout, "Response from `PublicSmtpTokensApi.GetMarketingV3TransactionalSmtpTokensTokenIdGetTokenById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | Identifier generated when a token is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3TransactionalSmtpTokensTokenIdGetTokenByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SmtpApiTokenView**](SmtpApiTokenView.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3TransactionalSmtpTokensCreateToken

> SmtpApiTokenView PostMarketingV3TransactionalSmtpTokensCreateToken(ctx).SmtpApiTokenRequestEgg(smtpApiTokenRequestEgg).Execute()

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
    smtpApiTokenRequestEgg := *openapiclient.NewSmtpApiTokenRequestEgg(false, "CampaignName_example") // SmtpApiTokenRequestEgg | A request object that includes the campaign name tied to the token and whether contacts should be created for recipients of emails. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicSmtpTokensApi.PostMarketingV3TransactionalSmtpTokensCreateToken(context.Background()).SmtpApiTokenRequestEgg(smtpApiTokenRequestEgg).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSmtpTokensApi.PostMarketingV3TransactionalSmtpTokensCreateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMarketingV3TransactionalSmtpTokensCreateToken`: SmtpApiTokenView
    fmt.Fprintf(os.Stdout, "Response from `PublicSmtpTokensApi.PostMarketingV3TransactionalSmtpTokensCreateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3TransactionalSmtpTokensCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smtpApiTokenRequestEgg** | [**SmtpApiTokenRequestEgg**](SmtpApiTokenRequestEgg.md) | A request object that includes the campaign name tied to the token and whether contacts should be created for recipients of emails. | 

### Return type

[**SmtpApiTokenView**](SmtpApiTokenView.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3TransactionalSmtpTokensTokenIdPasswordResetResetPassword

> SmtpApiTokenView PostMarketingV3TransactionalSmtpTokensTokenIdPasswordResetResetPassword(ctx, tokenId).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicSmtpTokensApi.PostMarketingV3TransactionalSmtpTokensTokenIdPasswordResetResetPassword(context.Background(), tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSmtpTokensApi.PostMarketingV3TransactionalSmtpTokensTokenIdPasswordResetResetPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMarketingV3TransactionalSmtpTokensTokenIdPasswordResetResetPassword`: SmtpApiTokenView
    fmt.Fprintf(os.Stdout, "Response from `PublicSmtpTokensApi.PostMarketingV3TransactionalSmtpTokensTokenIdPasswordResetResetPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | Identifier generated when a token is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3TransactionalSmtpTokensTokenIdPasswordResetResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SmtpApiTokenView**](SmtpApiTokenView.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

