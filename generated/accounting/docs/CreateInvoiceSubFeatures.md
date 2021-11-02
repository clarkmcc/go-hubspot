# CreateInvoiceSubFeatures

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateCustomer** | **bool** | Indicates if a new customer can be created in the external accounting system. | [default to null]
**Taxes** | **bool** | Indicates if taxes can be specified by the HubSpot user for line items. | [default to null]
**ExchangeRates** | **bool** | Indicates if the external accounting system supports fetching exchange rates when create an invoice in HubSpot for a customer billed in a different currency. | [default to null]
**Terms** | **bool** | Indicates if the external accounting system supports fetching payment terms. | [default to null]
**InvoiceComments** | **bool** | Indicates if a visible comment can be added to invoices. | [default to null]
**InvoiceDiscounts** | **bool** | Indicates if invoice-level discounts can be added to invoices. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

