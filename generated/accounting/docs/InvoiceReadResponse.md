# InvoiceReadResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalInvoiceNumber** | **string** | The invoice number. Note that this is _not_ the ID of the invoice, but the number that the billed customer will see. | [optional] [default to null]
**TotalAmountBilled** | **float64** | The total amount that this invoice is for. | [default to null]
**BalanceDue** | **float64** | The remaining balance due for the invoice. | [default to null]
**CurrencyCode** | **string** | The ISO 4217 currency code that represents the currency of the invoice. | [default to null]
**DueDate** | **string** | The due date of the invoice | [default to null]
**ExternalRecipientId** | **string** | The id of the customer in the external accounting system that the invoice was sent to. | [default to null]
**ReceivedByRecipientDate** | **int64** | The datetime that that invoice was sent to the customer. | [optional] [default to null]
**ExternalCreateDateTime** | **int64** | The datetime that the invoice was created in the external accounting system. | [optional] [default to null]
**IsVoided** | **bool** | Indicated is the invoice has been voided or not. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The datetime that the invoice was created in HubSpot. | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The datetime that the invoice was last updated in HubSpot. | [default to null]
**ArchivedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Archived** | **bool** |  | [default to null]
**ExternalAccountId** | **string** | The id of the account in the external accounting system that this invoice is related to. | [default to null]
**InvoiceStatus** | **string** | The status of the invoice | [default to null]
**Id** | **string** | The id of this invoice in the external accounting system. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

