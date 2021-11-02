# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Postcrmv3extensionsaccountingcallbackcustomerCreaterequestIdCreateCustomer**](CallbacksApi.md#Postcrmv3extensionsaccountingcallbackcustomerCreaterequestIdCreateCustomer) | **Post** /crm/v3/extensions/accounting/callback/customer-create/{requestId} | Endpoint for customer creation response
[**Postcrmv3extensionsaccountingcallbackcustomerSearchrequestIdDoCustomerSearch**](CallbacksApi.md#Postcrmv3extensionsaccountingcallbackcustomerSearchrequestIdDoCustomerSearch) | **Post** /crm/v3/extensions/accounting/callback/customer-search/{requestId} | Endpoint for customer search response
[**Postcrmv3extensionsaccountingcallbackexchangeRaterequestIdCreateExchangeRate**](CallbacksApi.md#Postcrmv3extensionsaccountingcallbackexchangeRaterequestIdCreateExchangeRate) | **Post** /crm/v3/extensions/accounting/callback/exchange-rate/{requestId} | Endpoint for exchange rate response
[**Postcrmv3extensionsaccountingcallbackinvoiceCreaterequestIdCreateInvoice**](CallbacksApi.md#Postcrmv3extensionsaccountingcallbackinvoiceCreaterequestIdCreateInvoice) | **Post** /crm/v3/extensions/accounting/callback/invoice-create/{requestId} | Endpoint for invoice creation response
[**Postcrmv3extensionsaccountingcallbackinvoicePdfrequestIdInvoicePdf**](CallbacksApi.md#Postcrmv3extensionsaccountingcallbackinvoicePdfrequestIdInvoicePdf) | **Post** /crm/v3/extensions/accounting/callback/invoice-pdf/{requestId} | Endpoint for PDF content of invoice
[**Postcrmv3extensionsaccountingcallbackinvoiceSearchrequestIdDoInvoiceSearch**](CallbacksApi.md#Postcrmv3extensionsaccountingcallbackinvoiceSearchrequestIdDoInvoiceSearch) | **Post** /crm/v3/extensions/accounting/callback/invoice-search/{requestId} | Endpoint for invoice search response
[**Postcrmv3extensionsaccountingcallbackinvoicesrequestIdGetById**](CallbacksApi.md#Postcrmv3extensionsaccountingcallbackinvoicesrequestIdGetById) | **Post** /crm/v3/extensions/accounting/callback/invoices/{requestId} | Endpoint for invoice get-by-id response
[**Postcrmv3extensionsaccountingcallbackproductSearchrequestIdDoProductSearch**](CallbacksApi.md#Postcrmv3extensionsaccountingcallbackproductSearchrequestIdDoProductSearch) | **Post** /crm/v3/extensions/accounting/callback/product-search/{requestId} | Endpoint for product search response
[**Postcrmv3extensionsaccountingcallbacktaxSearchrequestIdDoTaxSearch**](CallbacksApi.md#Postcrmv3extensionsaccountingcallbacktaxSearchrequestIdDoTaxSearch) | **Post** /crm/v3/extensions/accounting/callback/tax-search/{requestId} | Endpoint for taxes search response
[**Postcrmv3extensionsaccountingcallbacktermsrequestIdCreateTerm**](CallbacksApi.md#Postcrmv3extensionsaccountingcallbacktermsrequestIdCreateTerm) | **Post** /crm/v3/extensions/accounting/callback/terms/{requestId} | Endpoint for terms search response

# **Postcrmv3extensionsaccountingcallbackcustomerCreaterequestIdCreateCustomer**
> Postcrmv3extensionsaccountingcallbackcustomerCreaterequestIdCreateCustomer(ctx, body, requestId)
Endpoint for customer creation response

Call this endpoint with the response to a customer creation request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResultIdAccountingResponse**](ResultIdAccountingResponse.md)| The ID of the created customer. | 
  **requestId** | **string**| The ID of the request that this response is for | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3extensionsaccountingcallbackcustomerSearchrequestIdDoCustomerSearch**
> Postcrmv3extensionsaccountingcallbackcustomerSearchrequestIdDoCustomerSearch(ctx, body, requestId)
Endpoint for customer search response

Call this endpoint with the response to a customer search request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CustomerSearchResponseExternal**](CustomerSearchResponseExternal.md)| The result of the customer search request. | 
  **requestId** | **string**| The ID of the request that this response is for | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3extensionsaccountingcallbackexchangeRaterequestIdCreateExchangeRate**
> Postcrmv3extensionsaccountingcallbackexchangeRaterequestIdCreateExchangeRate(ctx, body, requestId)
Endpoint for exchange rate response

Call this endpoint with the response to an exchange rate request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExchangeRateResponse**](ExchangeRateResponse.md)| The result of the exchange rate request. | 
  **requestId** | **string**| The ID of the request that this response is for | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3extensionsaccountingcallbackinvoiceCreaterequestIdCreateInvoice**
> Postcrmv3extensionsaccountingcallbackinvoiceCreaterequestIdCreateInvoice(ctx, body, requestId)
Endpoint for invoice creation response

Call this endpoint with the response to a invoice creation request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResultIdAccountingResponse**](ResultIdAccountingResponse.md)| The ID of the created invoice. | 
  **requestId** | **string**| The ID of the request that this response is for | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3extensionsaccountingcallbackinvoicePdfrequestIdInvoicePdf**
> Postcrmv3extensionsaccountingcallbackinvoicePdfrequestIdInvoicePdf(ctx, body, requestId)
Endpoint for PDF content of invoice

Call this endpoint with the PDF content of a requested invoice.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InvoicePdfResponse**](InvoicePdfResponse.md)| The bytes of the invoice PDF. | 
  **requestId** | **string**| The ID of the request that this response is for | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3extensionsaccountingcallbackinvoiceSearchrequestIdDoInvoiceSearch**
> Postcrmv3extensionsaccountingcallbackinvoiceSearchrequestIdDoInvoiceSearch(ctx, body, requestId)
Endpoint for invoice search response

Call this endpoint with the response to a invoice search request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InvoiceSearchResponse**](InvoiceSearchResponse.md)| The result of the invoice search request. | 
  **requestId** | **string**| The ID of the request that this response is for | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3extensionsaccountingcallbackinvoicesrequestIdGetById**
> Postcrmv3extensionsaccountingcallbackinvoicesrequestIdGetById(ctx, body, requestId)
Endpoint for invoice get-by-id response

Call this endpoint with the response to a invoice get-by-id request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InvoicesResponseExternal**](InvoicesResponseExternal.md)| The result of the invoice request. | 
  **requestId** | **string**| The ID of the request that this response is for | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3extensionsaccountingcallbackproductSearchrequestIdDoProductSearch**
> Postcrmv3extensionsaccountingcallbackproductSearchrequestIdDoProductSearch(ctx, body, requestId)
Endpoint for product search response

Call this endpoint with the response to a product search request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProductSearchResponse**](ProductSearchResponse.md)| The result of the product search request. | 
  **requestId** | **string**| The ID of the request that this response is for | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3extensionsaccountingcallbacktaxSearchrequestIdDoTaxSearch**
> Postcrmv3extensionsaccountingcallbacktaxSearchrequestIdDoTaxSearch(ctx, body, requestId)
Endpoint for taxes search response

Call this endpoint with the response to a taxes search request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TaxSearchResponse**](TaxSearchResponse.md)| The result of the taxes search request. | 
  **requestId** | **string**| The ID of the request that this response is for | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3extensionsaccountingcallbacktermsrequestIdCreateTerm**
> Postcrmv3extensionsaccountingcallbacktermsrequestIdCreateTerm(ctx, body, requestId)
Endpoint for terms search response

Call this endpoint with the response to a terms search request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TermsResponse**](TermsResponse.md)| The result of the terms search | 
  **requestId** | **string**| The ID of the request that this response is for | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

