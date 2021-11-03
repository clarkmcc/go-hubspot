# InvoicesResponseExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **string** | Designates if the response is a success (&#39;OK&#39;) or failure (&#39;ERR&#39;). | [optional] 
**Invoices** | [**[]AccountingExtensionInvoice**](AccountingExtensionInvoice.md) | The list of invoices that were found for the request. | 

## Methods

### NewInvoicesResponseExternal

`func NewInvoicesResponseExternal(invoices []AccountingExtensionInvoice, ) *InvoicesResponseExternal`

NewInvoicesResponseExternal instantiates a new InvoicesResponseExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicesResponseExternalWithDefaults

`func NewInvoicesResponseExternalWithDefaults() *InvoicesResponseExternal`

NewInvoicesResponseExternalWithDefaults instantiates a new InvoicesResponseExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *InvoicesResponseExternal) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *InvoicesResponseExternal) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *InvoicesResponseExternal) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *InvoicesResponseExternal) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetInvoices

`func (o *InvoicesResponseExternal) GetInvoices() []AccountingExtensionInvoice`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *InvoicesResponseExternal) GetInvoicesOk() (*[]AccountingExtensionInvoice, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *InvoicesResponseExternal) SetInvoices(v []AccountingExtensionInvoice)`

SetInvoices sets Invoices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


