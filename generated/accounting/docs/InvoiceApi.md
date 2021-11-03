# \InvoiceApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCrmV3ExtensionsAccountingInvoiceInvoiceIdGetById**](InvoiceApi.md#GetCrmV3ExtensionsAccountingInvoiceInvoiceIdGetById) | **Get** /crm/v3/extensions/accounting/invoice/{invoiceId} | Get invoice data
[**PatchCrmV3ExtensionsAccountingInvoiceInvoiceIdUpdate**](InvoiceApi.md#PatchCrmV3ExtensionsAccountingInvoiceInvoiceIdUpdate) | **Patch** /crm/v3/extensions/accounting/invoice/{invoiceId} | Update an invoice
[**PostCrmV3ExtensionsAccountingInvoiceInvoiceIdPaymentCreatePayment**](InvoiceApi.md#PostCrmV3ExtensionsAccountingInvoiceInvoiceIdPaymentCreatePayment) | **Post** /crm/v3/extensions/accounting/invoice/{invoiceId}/payment | Records an invoice payment



## GetCrmV3ExtensionsAccountingInvoiceInvoiceIdGetById

> InvoiceReadResponse GetCrmV3ExtensionsAccountingInvoiceInvoiceIdGetById(ctx, invoiceId).AccountId(accountId).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoiceApi.GetCrmV3ExtensionsAccountingInvoiceInvoiceIdGetById(context.Background(), invoiceId).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.GetCrmV3ExtensionsAccountingInvoiceInvoiceIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3ExtensionsAccountingInvoiceInvoiceIdGetById`: InvoiceReadResponse
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.GetCrmV3ExtensionsAccountingInvoiceInvoiceIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | The ID of the invoice. This is the invoice ID from the external accounting system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ExtensionsAccountingInvoiceInvoiceIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | The ID of the account that the invoice belongs to. This is the account ID from the external accounting system. | 

### Return type

[**InvoiceReadResponse**](InvoiceReadResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCrmV3ExtensionsAccountingInvoiceInvoiceIdUpdate

> InvoiceUpdateResponse PatchCrmV3ExtensionsAccountingInvoiceInvoiceIdUpdate(ctx, invoiceId).AccountId(accountId).InvoiceUpdateRequest(invoiceUpdateRequest).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoiceApi.PatchCrmV3ExtensionsAccountingInvoiceInvoiceIdUpdate(context.Background(), invoiceId).AccountId(accountId).InvoiceUpdateRequest(invoiceUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.PatchCrmV3ExtensionsAccountingInvoiceInvoiceIdUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCrmV3ExtensionsAccountingInvoiceInvoiceIdUpdate`: InvoiceUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.PatchCrmV3ExtensionsAccountingInvoiceInvoiceIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | The ID of the invoice. This is the invoice ID from the external accounting system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCrmV3ExtensionsAccountingInvoiceInvoiceIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | The ID of the account that the invoice belongs to. This is the account ID from the external accounting system. | 
 **invoiceUpdateRequest** | [**InvoiceUpdateRequest**](InvoiceUpdateRequest.md) | The invoice data to update | 

### Return type

[**InvoiceUpdateResponse**](InvoiceUpdateResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ExtensionsAccountingInvoiceInvoiceIdPaymentCreatePayment

> InvoiceUpdateResponse PostCrmV3ExtensionsAccountingInvoiceInvoiceIdPaymentCreatePayment(ctx, invoiceId).InvoiceCreatePaymentRequest(invoiceCreatePaymentRequest).AccountId(accountId).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoiceApi.PostCrmV3ExtensionsAccountingInvoiceInvoiceIdPaymentCreatePayment(context.Background(), invoiceId).InvoiceCreatePaymentRequest(invoiceCreatePaymentRequest).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.PostCrmV3ExtensionsAccountingInvoiceInvoiceIdPaymentCreatePayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ExtensionsAccountingInvoiceInvoiceIdPaymentCreatePayment`: InvoiceUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.PostCrmV3ExtensionsAccountingInvoiceInvoiceIdPaymentCreatePayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | The ID of the invoice. This is the invoice ID from the external accounting system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExtensionsAccountingInvoiceInvoiceIdPaymentCreatePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoiceCreatePaymentRequest** | [**InvoiceCreatePaymentRequest**](InvoiceCreatePaymentRequest.md) | The payment information | 
 **accountId** | **string** | The ID of the account that the invoice belongs to. This is the account ID from the external accounting system. | 

### Return type

[**InvoiceUpdateResponse**](InvoiceUpdateResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

