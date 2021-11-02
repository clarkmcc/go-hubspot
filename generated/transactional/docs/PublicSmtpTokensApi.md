# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletemarketingv3transactionalsmtpTokenstokenIdArchiveToken**](PublicSmtpTokensApi.md#Deletemarketingv3transactionalsmtpTokenstokenIdArchiveToken) | **Delete** /marketing/v3/transactional/smtp-tokens/{tokenId} | Delete a single token by ID.
[**Getmarketingv3transactionalsmtpTokensGetTokensPage**](PublicSmtpTokensApi.md#Getmarketingv3transactionalsmtpTokensGetTokensPage) | **Get** /marketing/v3/transactional/smtp-tokens | Query SMTP API tokens by campaign name or an emailCampaignId.
[**Getmarketingv3transactionalsmtpTokenstokenIdGetTokenById**](PublicSmtpTokensApi.md#Getmarketingv3transactionalsmtpTokenstokenIdGetTokenById) | **Get** /marketing/v3/transactional/smtp-tokens/{tokenId} | Query a single token by ID.
[**Postmarketingv3transactionalsmtpTokensCreateToken**](PublicSmtpTokensApi.md#Postmarketingv3transactionalsmtpTokensCreateToken) | **Post** /marketing/v3/transactional/smtp-tokens | Create a SMTP API token.
[**Postmarketingv3transactionalsmtpTokenstokenIdpasswordResetResetPassword**](PublicSmtpTokensApi.md#Postmarketingv3transactionalsmtpTokenstokenIdpasswordResetResetPassword) | **Post** /marketing/v3/transactional/smtp-tokens/{tokenId}/password-reset | Reset the password of an existing token.

# **Deletemarketingv3transactionalsmtpTokenstokenIdArchiveToken**
> Deletemarketingv3transactionalsmtpTokenstokenIdArchiveToken(ctx, tokenId)
Delete a single token by ID.

Delete a single token by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenId** | **string**| Identifier generated when a token is created. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getmarketingv3transactionalsmtpTokensGetTokensPage**
> CollectionResponseSmtpApiTokenView Getmarketingv3transactionalsmtpTokensGetTokensPage(ctx, optional)
Query SMTP API tokens by campaign name or an emailCampaignId.

Query multiple SMTP API tokens by campaign name or a single token by emailCampaignId.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PublicSmtpTokensApiGetmarketingv3transactionalsmtpTokensGetTokensPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicSmtpTokensApiGetmarketingv3transactionalsmtpTokensGetTokensPageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignName** | **optional.String**| A name for the campaign tied to the SMTP API token. | 
 **emailCampaignId** | **optional.String**| Identifier assigned to the campaign provided during the token creation. | 
 **after** | **optional.String**| Starting point to get the next set of results. | 
 **limit** | **optional.Int32**| Maximum number of tokens to return. | 

### Return type

[**CollectionResponseSmtpApiTokenView**](CollectionResponseSmtpApiTokenView.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getmarketingv3transactionalsmtpTokenstokenIdGetTokenById**
> SmtpApiTokenView Getmarketingv3transactionalsmtpTokenstokenIdGetTokenById(ctx, tokenId)
Query a single token by ID.

Query a single token by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenId** | **string**| Identifier generated when a token is created. | 

### Return type

[**SmtpApiTokenView**](SmtpApiTokenView.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postmarketingv3transactionalsmtpTokensCreateToken**
> SmtpApiTokenView Postmarketingv3transactionalsmtpTokensCreateToken(ctx, optional)
Create a SMTP API token.

Create a SMTP API token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PublicSmtpTokensApiPostmarketingv3transactionalsmtpTokensCreateTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicSmtpTokensApiPostmarketingv3transactionalsmtpTokensCreateTokenOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SmtpApiTokenRequestEgg**](SmtpApiTokenRequestEgg.md)| A request object that includes the campaign name tied to the token and whether contacts should be created for recipients of emails. | 

### Return type

[**SmtpApiTokenView**](SmtpApiTokenView.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postmarketingv3transactionalsmtpTokenstokenIdpasswordResetResetPassword**
> SmtpApiTokenView Postmarketingv3transactionalsmtpTokenstokenIdpasswordResetResetPassword(ctx, tokenId)
Reset the password of an existing token.

Allows the creation of a replacement password for a given token. Once the password is successfully reset, the old password for the token will be invalid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenId** | **string**| Identifier generated when a token is created. | 

### Return type

[**SmtpApiTokenView**](SmtpApiTokenView.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

