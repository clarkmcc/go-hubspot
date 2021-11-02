# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getcrmv3extensionsaccountinginvoiceinvoiceIdGetById**](InvoiceApi.md#Getcrmv3extensionsaccountinginvoiceinvoiceIdGetById) | **Get** /crm/v3/extensions/accounting/invoice/{invoiceId} | Get invoice data
[**Patchcrmv3extensionsaccountinginvoiceinvoiceIdUpdate**](InvoiceApi.md#Patchcrmv3extensionsaccountinginvoiceinvoiceIdUpdate) | **Patch** /crm/v3/extensions/accounting/invoice/{invoiceId} | Update an invoice
[**Postcrmv3extensionsaccountinginvoiceinvoiceIdpaymentCreatePayment**](InvoiceApi.md#Postcrmv3extensionsaccountinginvoiceinvoiceIdpaymentCreatePayment) | **Post** /crm/v3/extensions/accounting/invoice/{invoiceId}/payment | Records an invoice payment

# **Getcrmv3extensionsaccountinginvoiceinvoiceIdGetById**
> InvoiceReadResponse Getcrmv3extensionsaccountinginvoiceinvoiceIdGetById(ctx, invoiceId, accountId)
Get invoice data

Returns invoice data for an Accounting account from the specified ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **invoiceId** | **string**| The ID of the invoice. This is the invoice ID from the external accounting system. | 
  **accountId** | **string**| The ID of the account that the invoice belongs to. This is the account ID from the external accounting system. | 

### Return type

[**InvoiceReadResponse**](InvoiceReadResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patchcrmv3extensionsaccountinginvoiceinvoiceIdUpdate**
> InvoiceUpdateResponse Patchcrmv3extensionsaccountinginvoiceinvoiceIdUpdate(ctx, body, invoiceId, accountId)
Update an invoice

Updates an Invoice by the given ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InvoiceUpdateRequest**](InvoiceUpdateRequest.md)| The invoice data to update | 
  **invoiceId** | **string**| The ID of the invoice. This is the invoice ID from the external accounting system. | 
  **accountId** | **string**| The ID of the account that the invoice belongs to. This is the account ID from the external accounting system. | 

### Return type

[**InvoiceUpdateResponse**](InvoiceUpdateResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3extensionsaccountinginvoiceinvoiceIdpaymentCreatePayment**
> InvoiceUpdateResponse Postcrmv3extensionsaccountinginvoiceinvoiceIdpaymentCreatePayment(ctx, body, invoiceId, optional)
Records an invoice payment

Records an payment against an invoice.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InvoiceCreatePaymentRequest**](InvoiceCreatePaymentRequest.md)| The payment information | 
  **invoiceId** | **string**| The ID of the invoice. This is the invoice ID from the external accounting system. | 
 **optional** | ***InvoiceApiPostcrmv3extensionsaccountinginvoiceinvoiceIdpaymentCreatePaymentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvoiceApiPostcrmv3extensionsaccountinginvoiceinvoiceIdpaymentCreatePaymentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountId** | **optional.**| The ID of the account that the invoice belongs to. This is the account ID from the external accounting system. | 

### Return type

[**InvoiceUpdateResponse**](InvoiceUpdateResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

