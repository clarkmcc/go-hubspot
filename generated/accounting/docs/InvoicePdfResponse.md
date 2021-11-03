# InvoicePdfResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **string** | Designates if the response is a success (&#39;OK&#39;) or failure (&#39;ERR&#39;). | [optional] 
**Invoice** | **string** | The bytes of the invoice PDF. | 

## Methods

### NewInvoicePdfResponse

`func NewInvoicePdfResponse(invoice string, ) *InvoicePdfResponse`

NewInvoicePdfResponse instantiates a new InvoicePdfResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicePdfResponseWithDefaults

`func NewInvoicePdfResponseWithDefaults() *InvoicePdfResponse`

NewInvoicePdfResponseWithDefaults instantiates a new InvoicePdfResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *InvoicePdfResponse) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *InvoicePdfResponse) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *InvoicePdfResponse) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *InvoicePdfResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetInvoice

`func (o *InvoicePdfResponse) GetInvoice() string`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *InvoicePdfResponse) GetInvoiceOk() (*string, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *InvoicePdfResponse) SetInvoice(v string)`

SetInvoice sets Invoice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


