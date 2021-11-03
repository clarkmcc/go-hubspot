# InvoiceSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **string** | Designates if the response is a success (&#39;OK&#39;) or failure (&#39;ERR&#39;). | [optional] 
**Invoices** | [**[]AccountingExtensionInvoice**](AccountingExtensionInvoice.md) | The list of invoices that matched the search criteria. | 

## Methods

### NewInvoiceSearchResponse

`func NewInvoiceSearchResponse(invoices []AccountingExtensionInvoice, ) *InvoiceSearchResponse`

NewInvoiceSearchResponse instantiates a new InvoiceSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceSearchResponseWithDefaults

`func NewInvoiceSearchResponseWithDefaults() *InvoiceSearchResponse`

NewInvoiceSearchResponseWithDefaults instantiates a new InvoiceSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *InvoiceSearchResponse) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *InvoiceSearchResponse) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *InvoiceSearchResponse) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *InvoiceSearchResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetInvoices

`func (o *InvoiceSearchResponse) GetInvoices() []AccountingExtensionInvoice`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *InvoiceSearchResponse) GetInvoicesOk() (*[]AccountingExtensionInvoice, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *InvoiceSearchResponse) SetInvoices(v []AccountingExtensionInvoice)`

SetInvoices sets Invoices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


