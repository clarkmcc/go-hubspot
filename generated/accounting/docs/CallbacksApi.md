# \CallbacksApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmV3ExtensionsAccountingCallbackCustomerCreateRequestIdCreateCustomer**](CallbacksApi.md#PostCrmV3ExtensionsAccountingCallbackCustomerCreateRequestIdCreateCustomer) | **Post** /crm/v3/extensions/accounting/callback/customer-create/{requestId} | Endpoint for customer creation response
[**PostCrmV3ExtensionsAccountingCallbackCustomerSearchRequestIdDoCustomerSearch**](CallbacksApi.md#PostCrmV3ExtensionsAccountingCallbackCustomerSearchRequestIdDoCustomerSearch) | **Post** /crm/v3/extensions/accounting/callback/customer-search/{requestId} | Endpoint for customer search response
[**PostCrmV3ExtensionsAccountingCallbackExchangeRateRequestIdCreateExchangeRate**](CallbacksApi.md#PostCrmV3ExtensionsAccountingCallbackExchangeRateRequestIdCreateExchangeRate) | **Post** /crm/v3/extensions/accounting/callback/exchange-rate/{requestId} | Endpoint for exchange rate response
[**PostCrmV3ExtensionsAccountingCallbackInvoiceCreateRequestIdCreateInvoice**](CallbacksApi.md#PostCrmV3ExtensionsAccountingCallbackInvoiceCreateRequestIdCreateInvoice) | **Post** /crm/v3/extensions/accounting/callback/invoice-create/{requestId} | Endpoint for invoice creation response
[**PostCrmV3ExtensionsAccountingCallbackInvoicePdfRequestIdInvoicePdf**](CallbacksApi.md#PostCrmV3ExtensionsAccountingCallbackInvoicePdfRequestIdInvoicePdf) | **Post** /crm/v3/extensions/accounting/callback/invoice-pdf/{requestId} | Endpoint for PDF content of invoice
[**PostCrmV3ExtensionsAccountingCallbackInvoiceSearchRequestIdDoInvoiceSearch**](CallbacksApi.md#PostCrmV3ExtensionsAccountingCallbackInvoiceSearchRequestIdDoInvoiceSearch) | **Post** /crm/v3/extensions/accounting/callback/invoice-search/{requestId} | Endpoint for invoice search response
[**PostCrmV3ExtensionsAccountingCallbackInvoicesRequestIdGetById**](CallbacksApi.md#PostCrmV3ExtensionsAccountingCallbackInvoicesRequestIdGetById) | **Post** /crm/v3/extensions/accounting/callback/invoices/{requestId} | Endpoint for invoice get-by-id response
[**PostCrmV3ExtensionsAccountingCallbackProductSearchRequestIdDoProductSearch**](CallbacksApi.md#PostCrmV3ExtensionsAccountingCallbackProductSearchRequestIdDoProductSearch) | **Post** /crm/v3/extensions/accounting/callback/product-search/{requestId} | Endpoint for product search response
[**PostCrmV3ExtensionsAccountingCallbackTaxSearchRequestIdDoTaxSearch**](CallbacksApi.md#PostCrmV3ExtensionsAccountingCallbackTaxSearchRequestIdDoTaxSearch) | **Post** /crm/v3/extensions/accounting/callback/tax-search/{requestId} | Endpoint for taxes search response
[**PostCrmV3ExtensionsAccountingCallbackTermsRequestIdCreateTerm**](CallbacksApi.md#PostCrmV3ExtensionsAccountingCallbackTermsRequestIdCreateTerm) | **Post** /crm/v3/extensions/accounting/callback/terms/{requestId} | Endpoint for terms search response



## PostCrmV3ExtensionsAccountingCallbackCustomerCreateRequestIdCreateCustomer

> PostCrmV3ExtensionsAccountingCallbackCustomerCreateRequestIdCreateCustomer(ctx, requestId).ResultIdAccountingResponse(resultIdAccountingResponse).Execute()

Endpoint for customer creation response



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
    requestId := "requestId_example" // string | The ID of the request that this response is for
    resultIdAccountingResponse := *openapiclient.NewResultIdAccountingResponse("OK", "js-1") // ResultIdAccountingResponse | The ID of the created customer.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallbacksApi.PostCrmV3ExtensionsAccountingCallbackCustomerCreateRequestIdCreateCustomer(context.Background(), requestId).ResultIdAccountingResponse(resultIdAccountingResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.PostCrmV3ExtensionsAccountingCallbackCustomerCreateRequestIdCreateCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | The ID of the request that this response is for | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExtensionsAccountingCallbackCustomerCreateRequestIdCreateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resultIdAccountingResponse** | [**ResultIdAccountingResponse**](ResultIdAccountingResponse.md) | The ID of the created customer. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ExtensionsAccountingCallbackCustomerSearchRequestIdDoCustomerSearch

> PostCrmV3ExtensionsAccountingCallbackCustomerSearchRequestIdDoCustomerSearch(ctx, requestId).CustomerSearchResponseExternal(customerSearchResponseExternal).Execute()

Endpoint for customer search response



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
    requestId := "requestId_example" // string | The ID of the request that this response is for
    customerSearchResponseExternal := *openapiclient.NewCustomerSearchResponseExternal("OK", []openapiclient.AccountingExtensionCustomer{*openapiclient.NewAccountingExtensionCustomer("John Smith", "js-1")}) // CustomerSearchResponseExternal | The result of the customer search request.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallbacksApi.PostCrmV3ExtensionsAccountingCallbackCustomerSearchRequestIdDoCustomerSearch(context.Background(), requestId).CustomerSearchResponseExternal(customerSearchResponseExternal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.PostCrmV3ExtensionsAccountingCallbackCustomerSearchRequestIdDoCustomerSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | The ID of the request that this response is for | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExtensionsAccountingCallbackCustomerSearchRequestIdDoCustomerSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerSearchResponseExternal** | [**CustomerSearchResponseExternal**](CustomerSearchResponseExternal.md) | The result of the customer search request. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ExtensionsAccountingCallbackExchangeRateRequestIdCreateExchangeRate

> PostCrmV3ExtensionsAccountingCallbackExchangeRateRequestIdCreateExchangeRate(ctx, requestId).ExchangeRateResponse(exchangeRateResponse).Execute()

Endpoint for exchange rate response



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
    requestId := "requestId_example" // string | The ID of the request that this response is for
    exchangeRateResponse := *openapiclient.NewExchangeRateResponse("OK", float32(1.003), "JPY", "USD") // ExchangeRateResponse | The result of the exchange rate request.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallbacksApi.PostCrmV3ExtensionsAccountingCallbackExchangeRateRequestIdCreateExchangeRate(context.Background(), requestId).ExchangeRateResponse(exchangeRateResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.PostCrmV3ExtensionsAccountingCallbackExchangeRateRequestIdCreateExchangeRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | The ID of the request that this response is for | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExtensionsAccountingCallbackExchangeRateRequestIdCreateExchangeRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exchangeRateResponse** | [**ExchangeRateResponse**](ExchangeRateResponse.md) | The result of the exchange rate request. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ExtensionsAccountingCallbackInvoiceCreateRequestIdCreateInvoice

> PostCrmV3ExtensionsAccountingCallbackInvoiceCreateRequestIdCreateInvoice(ctx, requestId).ResultIdAccountingResponse(resultIdAccountingResponse).Execute()

Endpoint for invoice creation response



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
    requestId := "requestId_example" // string | The ID of the request that this response is for
    resultIdAccountingResponse := *openapiclient.NewResultIdAccountingResponse("OK", "js-1") // ResultIdAccountingResponse | The ID of the created invoice.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallbacksApi.PostCrmV3ExtensionsAccountingCallbackInvoiceCreateRequestIdCreateInvoice(context.Background(), requestId).ResultIdAccountingResponse(resultIdAccountingResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.PostCrmV3ExtensionsAccountingCallbackInvoiceCreateRequestIdCreateInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | The ID of the request that this response is for | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExtensionsAccountingCallbackInvoiceCreateRequestIdCreateInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resultIdAccountingResponse** | [**ResultIdAccountingResponse**](ResultIdAccountingResponse.md) | The ID of the created invoice. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ExtensionsAccountingCallbackInvoicePdfRequestIdInvoicePdf

> PostCrmV3ExtensionsAccountingCallbackInvoicePdfRequestIdInvoicePdf(ctx, requestId).InvoicePdfResponse(invoicePdfResponse).Execute()

Endpoint for PDF content of invoice



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
    requestId := "requestId_example" // string | The ID of the request that this response is for
    invoicePdfResponse := *openapiclient.NewInvoicePdfResponse("Invoice_example") // InvoicePdfResponse | The bytes of the invoice PDF.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallbacksApi.PostCrmV3ExtensionsAccountingCallbackInvoicePdfRequestIdInvoicePdf(context.Background(), requestId).InvoicePdfResponse(invoicePdfResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.PostCrmV3ExtensionsAccountingCallbackInvoicePdfRequestIdInvoicePdf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | The ID of the request that this response is for | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExtensionsAccountingCallbackInvoicePdfRequestIdInvoicePdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoicePdfResponse** | [**InvoicePdfResponse**](InvoicePdfResponse.md) | The bytes of the invoice PDF. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ExtensionsAccountingCallbackInvoiceSearchRequestIdDoInvoiceSearch

> PostCrmV3ExtensionsAccountingCallbackInvoiceSearchRequestIdDoInvoiceSearch(ctx, requestId).InvoiceSearchResponse(invoiceSearchResponse).Execute()

Endpoint for invoice search response



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    requestId := "requestId_example" // string | The ID of the request that this response is for
    invoiceSearchResponse := *openapiclient.NewInvoiceSearchResponse([]openapiclient.AccountingExtensionInvoice{*openapiclient.NewAccountingExtensionInvoice(float32(100.5), time.Now(), "USD", "https://myapp.com/invoices/1243a2", "John Smith", "OVERDUE")}) // InvoiceSearchResponse | The result of the invoice search request.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallbacksApi.PostCrmV3ExtensionsAccountingCallbackInvoiceSearchRequestIdDoInvoiceSearch(context.Background(), requestId).InvoiceSearchResponse(invoiceSearchResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.PostCrmV3ExtensionsAccountingCallbackInvoiceSearchRequestIdDoInvoiceSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | The ID of the request that this response is for | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExtensionsAccountingCallbackInvoiceSearchRequestIdDoInvoiceSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoiceSearchResponse** | [**InvoiceSearchResponse**](InvoiceSearchResponse.md) | The result of the invoice search request. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ExtensionsAccountingCallbackInvoicesRequestIdGetById

> PostCrmV3ExtensionsAccountingCallbackInvoicesRequestIdGetById(ctx, requestId).InvoicesResponseExternal(invoicesResponseExternal).Execute()

Endpoint for invoice get-by-id response



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    requestId := "requestId_example" // string | The ID of the request that this response is for
    invoicesResponseExternal := *openapiclient.NewInvoicesResponseExternal([]openapiclient.AccountingExtensionInvoice{*openapiclient.NewAccountingExtensionInvoice(float32(100.5), time.Now(), "USD", "https://myapp.com/invoices/1243a2", "John Smith", "OVERDUE")}) // InvoicesResponseExternal | The result of the invoice request.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallbacksApi.PostCrmV3ExtensionsAccountingCallbackInvoicesRequestIdGetById(context.Background(), requestId).InvoicesResponseExternal(invoicesResponseExternal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.PostCrmV3ExtensionsAccountingCallbackInvoicesRequestIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | The ID of the request that this response is for | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExtensionsAccountingCallbackInvoicesRequestIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoicesResponseExternal** | [**InvoicesResponseExternal**](InvoicesResponseExternal.md) | The result of the invoice request. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ExtensionsAccountingCallbackProductSearchRequestIdDoProductSearch

> PostCrmV3ExtensionsAccountingCallbackProductSearchRequestIdDoProductSearch(ctx, requestId).ProductSearchResponse(productSearchResponse).Execute()

Endpoint for product search response



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
    requestId := "requestId_example" // string | The ID of the request that this response is for
    productSearchResponse := *openapiclient.NewProductSearchResponse([]openapiclient.Product{*openapiclient.NewProduct(*openapiclient.NewUnitPrice(false), false, "Marketing Services", "prod-123")}) // ProductSearchResponse | The result of the product search request.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallbacksApi.PostCrmV3ExtensionsAccountingCallbackProductSearchRequestIdDoProductSearch(context.Background(), requestId).ProductSearchResponse(productSearchResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.PostCrmV3ExtensionsAccountingCallbackProductSearchRequestIdDoProductSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | The ID of the request that this response is for | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExtensionsAccountingCallbackProductSearchRequestIdDoProductSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productSearchResponse** | [**ProductSearchResponse**](ProductSearchResponse.md) | The result of the product search request. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ExtensionsAccountingCallbackTaxSearchRequestIdDoTaxSearch

> PostCrmV3ExtensionsAccountingCallbackTaxSearchRequestIdDoTaxSearch(ctx, requestId).TaxSearchResponse(taxSearchResponse).Execute()

Endpoint for taxes search response



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
    requestId := "requestId_example" // string | The ID of the request that this response is for
    taxSearchResponse := *openapiclient.NewTaxSearchResponse([]openapiclient.Tax{*openapiclient.NewTax("tax-1", float32(8.05), "Local Sales Tax")}) // TaxSearchResponse | The result of the taxes search request.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallbacksApi.PostCrmV3ExtensionsAccountingCallbackTaxSearchRequestIdDoTaxSearch(context.Background(), requestId).TaxSearchResponse(taxSearchResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.PostCrmV3ExtensionsAccountingCallbackTaxSearchRequestIdDoTaxSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | The ID of the request that this response is for | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExtensionsAccountingCallbackTaxSearchRequestIdDoTaxSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taxSearchResponse** | [**TaxSearchResponse**](TaxSearchResponse.md) | The result of the taxes search request. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ExtensionsAccountingCallbackTermsRequestIdCreateTerm

> PostCrmV3ExtensionsAccountingCallbackTermsRequestIdCreateTerm(ctx, requestId).TermsResponse(termsResponse).Execute()

Endpoint for terms search response



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
    requestId := "requestId_example" // string | The ID of the request that this response is for
    termsResponse := *openapiclient.NewTermsResponse([]openapiclient.AccountingExtensionTerm{*openapiclient.NewAccountingExtensionTerm("Net 30", "net-30")}) // TermsResponse | The result of the terms search

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallbacksApi.PostCrmV3ExtensionsAccountingCallbackTermsRequestIdCreateTerm(context.Background(), requestId).TermsResponse(termsResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.PostCrmV3ExtensionsAccountingCallbackTermsRequestIdCreateTerm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | The ID of the request that this response is for | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExtensionsAccountingCallbackTermsRequestIdCreateTermRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **termsResponse** | [**TermsResponse**](TermsResponse.md) | The result of the terms search | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

