# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Postmarketingv3transactionalsingleEmailsendSendEmail**](SingleSendApi.md#Postmarketingv3transactionalsingleEmailsendSendEmail) | **Post** /marketing/v3/transactional/single-email/send | Send a single transactional email asynchronously.

# **Postmarketingv3transactionalsingleEmailsendSendEmail**
> EmailSendStatusView Postmarketingv3transactionalsingleEmailsendSendEmail(ctx, optional)
Send a single transactional email asynchronously.

Asynchronously send a transactional email. Returns the status of the email send with a statusId that can be used to continuously query for the status using the Email Send Status API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SingleSendApiPostmarketingv3transactionalsingleEmailsendSendEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SingleSendApiPostmarketingv3transactionalsingleEmailsendSendEmailOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PublicSingleSendRequestEgg**](PublicSingleSendRequestEgg.md)| A request object describing the email to send. | 

### Return type

[**EmailSendStatusView**](EmailSendStatusView.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

