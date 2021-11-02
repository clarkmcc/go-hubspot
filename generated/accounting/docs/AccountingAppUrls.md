# AccountingAppUrls

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GetInvoiceUrl** | **string** | A URL that specifies the endpoint where invoices can be retrieved. | [default to null]
**SearchCustomerUrl** | **string** | A URL that specifies the endpoint where a customer search can be performed. | [default to null]
**GetInvoicePdfUrl** | **string** | A URL that specifies the endpoint where an invoice PDF can be retrieved. | [default to null]
**CustomerUrlTemplate** | **string** | A template URL that indicates the endpoint where a customer can be fetched by ID. Note that ${CUSTOMER_ID} in this URL will be replaced by the actual customer ID. For example: https://myapp.com/api/customers/${CUSTOMER_ID} | [default to null]
**ProductUrlTemplate** | **string** | A template URL that indicates the endpoint where a product can be fetched by ID. Note that ${PRODUCT_ID} in this URL will be replaced by the actual product ID. For example: https://myapp.com/api/products/${PRODUCT_ID} | [default to null]
**InvoiceUrlTemplate** | **string** | A template URL that indicates the endpoint where an invoice can be fetched by ID. Note that ${INVOICE_ID} in this URL will be replaced by the actual invoice ID. For example: https://myapp.com/api/invoices/${INVOICE_ID} | [default to null]
**CreateInvoiceUrl** | **string** | A URL that specifies the endpoint where an invoices can be created. | [optional] [default to null]
**SearchInvoiceUrl** | **string** | A URL that specifies the endpoint where an invoice search can be performed. | [optional] [default to null]
**SearchProductUrl** | **string** | A URL that specifies the endpoint where a product search can be performed. | [optional] [default to null]
**GetTermsUrl** | **string** | A URL that specifies the endpoint where payment terms can be retrieved. | [optional] [default to null]
**CreateCustomerUrl** | **string** | A URL that specifies the endpoint where a new customer can be created. | [optional] [default to null]
**SearchTaxUrl** | **string** | A URL that specifies the endpoint where a tax search can be performed. | [optional] [default to null]
**ExchangeRateUrl** | **string** | A URL that specifies the endpoint where exchange rates can be queried. | [optional] [default to null]
**SearchUrl** | **string** |  | [optional] [default to null]
**SearchCountUrl** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

