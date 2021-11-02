# InvoiceUpdateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalInvoiceNumber** | **string** |  | [optional] [default to null]
**CurrencyCode** | **string** | The ISO 4217 currency code that represents the currency used in the invoice to bill the recipient | [optional] [default to null]
**DueDate** | **string** | The ISO-8601 due date of the invoice. | [optional] [default to null]
**ExternalRecipientId** | **string** | The ID of the invoice recipient. This is the recipient ID from the external accounting system. | [optional] [default to null]
**ReceivedByRecipientDate** | **int64** |  | [optional] [default to null]
**IsVoided** | **bool** | States if the invoice is voided or not. | [optional] [default to null]
**ReceivedByCustomerDate** | **string** | The ISO-8601 datetime of when the customer received the invoice. | [optional] [default to null]
**InvoiceNumber** | **string** | The number / name of the invoice. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

