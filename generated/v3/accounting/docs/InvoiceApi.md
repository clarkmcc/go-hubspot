# \InvoiceApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoiceCreatePayment**](InvoiceApi.md#InvoiceCreatePayment) | **Post** /crm/v3/extensions/accounting/invoice/{invoiceId}/payment | Records an invoice payment
[**InvoiceGetByID**](InvoiceApi.md#InvoiceGetByID) | **Get** /crm/v3/extensions/accounting/invoice/{invoiceId} | Get invoice data
[**InvoiceUpdate**](InvoiceApi.md#InvoiceUpdate) | **Patch** /crm/v3/extensions/accounting/invoice/{invoiceId} | Update an invoice



## InvoiceCreatePayment

> InvoiceUpdateResponse InvoiceCreatePayment(ctx, invoiceId).InvoiceCreatePaymentRequest(invoiceCreatePaymentRequest).AccountId(accountId).Execute()

Records an invoice payment



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
    invoiceId := "invoiceId_example" // string | The ID of the invoice. This is the invoice ID from the external accounting system.
    invoiceCreatePaymentRequest := *openapiclient.NewInvoiceCreatePaymentRequest(float32(123), "CurrencyCode_example", time.Now(), "ExternalPaymentId_example") // InvoiceCreatePaymentRequest | The payment information
    accountId := "accountId_example" // string | The ID of the account that the invoice belongs to. This is the account ID from the external accounting system. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceApi.InvoiceCreatePayment(context.Background(), invoiceId).InvoiceCreatePaymentRequest(invoiceCreatePaymentRequest).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.InvoiceCreatePayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvoiceCreatePayment`: InvoiceUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.InvoiceCreatePayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | The ID of the invoice. This is the invoice ID from the external accounting system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceCreatePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoiceCreatePaymentRequest** | [**InvoiceCreatePaymentRequest**](InvoiceCreatePaymentRequest.md) | The payment information | 
 **accountId** | **string** | The ID of the account that the invoice belongs to. This is the account ID from the external accounting system. | 

### Return type

[**InvoiceUpdateResponse**](InvoiceUpdateResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceGetByID

> InvoiceReadResponse InvoiceGetByID(ctx, invoiceId).AccountId(accountId).Execute()

Get invoice data



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
    invoiceId := "invoiceId_example" // string | The ID of the invoice. This is the invoice ID from the external accounting system.
    accountId := "accountId_example" // string | The ID of the account that the invoice belongs to. This is the account ID from the external accounting system.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceApi.InvoiceGetByID(context.Background(), invoiceId).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.InvoiceGetByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvoiceGetByID`: InvoiceReadResponse
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.InvoiceGetByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | The ID of the invoice. This is the invoice ID from the external accounting system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceGetByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | The ID of the account that the invoice belongs to. This is the account ID from the external accounting system. | 

### Return type

[**InvoiceReadResponse**](InvoiceReadResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceUpdate

> InvoiceUpdateResponse InvoiceUpdate(ctx, invoiceId).AccountId(accountId).InvoiceUpdateRequest(invoiceUpdateRequest).Execute()

Update an invoice



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
    invoiceId := "invoiceId_example" // string | The ID of the invoice. This is the invoice ID from the external accounting system.
    accountId := "accountId_example" // string | The ID of the account that the invoice belongs to. This is the account ID from the external accounting system.
    invoiceUpdateRequest := *openapiclient.NewInvoiceUpdateRequest() // InvoiceUpdateRequest | The invoice data to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceApi.InvoiceUpdate(context.Background(), invoiceId).AccountId(accountId).InvoiceUpdateRequest(invoiceUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.InvoiceUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvoiceUpdate`: InvoiceUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.InvoiceUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | The ID of the invoice. This is the invoice ID from the external accounting system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | The ID of the account that the invoice belongs to. This is the account ID from the external accounting system. | 
 **invoiceUpdateRequest** | [**InvoiceUpdateRequest**](InvoiceUpdateRequest.md) | The invoice data to update | 

### Return type

[**InvoiceUpdateResponse**](InvoiceUpdateResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

