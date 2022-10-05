# \CallbacksApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallbackCreateCustomer**](CallbacksApi.md#CallbackCreateCustomer) | **Post** /crm/v3/extensions/accounting/callback/customer-create/{requestId} | Endpoint for customer creation response
[**CallbackCreateExchangeRate**](CallbacksApi.md#CallbackCreateExchangeRate) | **Post** /crm/v3/extensions/accounting/callback/exchange-rate/{requestId} | Endpoint for exchange rate response
[**CallbackCreateInvoice**](CallbacksApi.md#CallbackCreateInvoice) | **Post** /crm/v3/extensions/accounting/callback/invoice-create/{requestId} | Endpoint for invoice creation response
[**CallbackCreateTerm**](CallbacksApi.md#CallbackCreateTerm) | **Post** /crm/v3/extensions/accounting/callback/terms/{requestId} | Endpoint for terms search response
[**CallbackDoInvoiceSearch**](CallbacksApi.md#CallbackDoInvoiceSearch) | **Post** /crm/v3/extensions/accounting/callback/invoice-search/{requestId} | Endpoint for invoice search response
[**CallbackDoProductSearch**](CallbacksApi.md#CallbackDoProductSearch) | **Post** /crm/v3/extensions/accounting/callback/product-search/{requestId} | Endpoint for product search response
[**CallbackDoSearchCustomer**](CallbacksApi.md#CallbackDoSearchCustomer) | **Post** /crm/v3/extensions/accounting/callback/customer-search/{requestId} | Endpoint for customer search response
[**CallbackDoTaxSearch**](CallbacksApi.md#CallbackDoTaxSearch) | **Post** /crm/v3/extensions/accounting/callback/tax-search/{requestId} | Endpoint for taxes search response
[**CallbackGetByID**](CallbacksApi.md#CallbackGetByID) | **Post** /crm/v3/extensions/accounting/callback/invoices/{requestId} | Endpoint for invoice get-by-id response
[**CallbackInvoicePDF**](CallbacksApi.md#CallbackInvoicePDF) | **Post** /crm/v3/extensions/accounting/callback/invoice-pdf/{requestId} | Endpoint for PDF content of invoice



## CallbackCreateCustomer

> CallbackCreateCustomer(ctx, requestId).ResultIdAccountingResponse(resultIdAccountingResponse).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CallbacksApi.CallbackCreateCustomer(context.Background(), requestId).ResultIdAccountingResponse(resultIdAccountingResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.CallbackCreateCustomer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCallbackCreateCustomerRequest struct via the builder pattern


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


## CallbackCreateExchangeRate

> CallbackCreateExchangeRate(ctx, requestId).ExchangeRateResponse(exchangeRateResponse).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CallbacksApi.CallbackCreateExchangeRate(context.Background(), requestId).ExchangeRateResponse(exchangeRateResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.CallbackCreateExchangeRate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCallbackCreateExchangeRateRequest struct via the builder pattern


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


## CallbackCreateInvoice

> CallbackCreateInvoice(ctx, requestId).ResultIdAccountingResponse(resultIdAccountingResponse).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CallbacksApi.CallbackCreateInvoice(context.Background(), requestId).ResultIdAccountingResponse(resultIdAccountingResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.CallbackCreateInvoice``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCallbackCreateInvoiceRequest struct via the builder pattern


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


## CallbackCreateTerm

> CallbackCreateTerm(ctx, requestId).TermsResponse(termsResponse).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CallbacksApi.CallbackCreateTerm(context.Background(), requestId).TermsResponse(termsResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.CallbackCreateTerm``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCallbackCreateTermRequest struct via the builder pattern


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


## CallbackDoInvoiceSearch

> CallbackDoInvoiceSearch(ctx, requestId).InvoiceSearchResponse(invoiceSearchResponse).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CallbacksApi.CallbackDoInvoiceSearch(context.Background(), requestId).InvoiceSearchResponse(invoiceSearchResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.CallbackDoInvoiceSearch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCallbackDoInvoiceSearchRequest struct via the builder pattern


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


## CallbackDoProductSearch

> CallbackDoProductSearch(ctx, requestId).ProductSearchResponse(productSearchResponse).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CallbacksApi.CallbackDoProductSearch(context.Background(), requestId).ProductSearchResponse(productSearchResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.CallbackDoProductSearch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCallbackDoProductSearchRequest struct via the builder pattern


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


## CallbackDoSearchCustomer

> CallbackDoSearchCustomer(ctx, requestId).CustomerSearchResponseExternal(customerSearchResponseExternal).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CallbacksApi.CallbackDoSearchCustomer(context.Background(), requestId).CustomerSearchResponseExternal(customerSearchResponseExternal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.CallbackDoSearchCustomer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCallbackDoSearchCustomerRequest struct via the builder pattern


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


## CallbackDoTaxSearch

> CallbackDoTaxSearch(ctx, requestId).TaxSearchResponse(taxSearchResponse).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CallbacksApi.CallbackDoTaxSearch(context.Background(), requestId).TaxSearchResponse(taxSearchResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.CallbackDoTaxSearch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCallbackDoTaxSearchRequest struct via the builder pattern


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


## CallbackGetByID

> CallbackGetByID(ctx, requestId).InvoicesResponseExternal(invoicesResponseExternal).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CallbacksApi.CallbackGetByID(context.Background(), requestId).InvoicesResponseExternal(invoicesResponseExternal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.CallbackGetByID``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCallbackGetByIDRequest struct via the builder pattern


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


## CallbackInvoicePDF

> CallbackInvoicePDF(ctx, requestId).InvoicePdfResponse(invoicePdfResponse).Execute()

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
    invoicePdfResponse := *openapiclient.NewInvoicePdfResponse(string(123)) // InvoicePdfResponse | The bytes of the invoice PDF.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CallbacksApi.CallbackInvoicePDF(context.Background(), requestId).InvoicePdfResponse(invoicePdfResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbacksApi.CallbackInvoicePDF``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCallbackInvoicePDFRequest struct via the builder pattern


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

