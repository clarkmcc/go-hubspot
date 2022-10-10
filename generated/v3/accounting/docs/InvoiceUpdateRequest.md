# InvoiceUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalInvoiceNumber** | Pointer to **string** |  | [optional] 
**CurrencyCode** | Pointer to **string** | The ISO 4217 currency code that represents the currency used in the invoice to bill the recipient | [optional] 
**DueDate** | Pointer to **string** | The ISO-8601 due date of the invoice. | [optional] 
**ExternalRecipientId** | Pointer to **string** | The ID of the invoice recipient. This is the recipient ID from the external accounting system. | [optional] 
**ReceivedByRecipientDate** | Pointer to **int64** |  | [optional] 
**IsVoided** | Pointer to **bool** | States if the invoice is voided or not. | [optional] 
**ReceivedByCustomerDate** | Pointer to **string** | The ISO-8601 datetime of when the customer received the invoice. | [optional] 
**InvoiceNumber** | Pointer to **string** | The number / name of the invoice. | [optional] 

## Methods

### NewInvoiceUpdateRequest

`func NewInvoiceUpdateRequest() *InvoiceUpdateRequest`

NewInvoiceUpdateRequest instantiates a new InvoiceUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceUpdateRequestWithDefaults

`func NewInvoiceUpdateRequestWithDefaults() *InvoiceUpdateRequest`

NewInvoiceUpdateRequestWithDefaults instantiates a new InvoiceUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalInvoiceNumber

`func (o *InvoiceUpdateRequest) GetExternalInvoiceNumber() string`

GetExternalInvoiceNumber returns the ExternalInvoiceNumber field if non-nil, zero value otherwise.

### GetExternalInvoiceNumberOk

`func (o *InvoiceUpdateRequest) GetExternalInvoiceNumberOk() (*string, bool)`

GetExternalInvoiceNumberOk returns a tuple with the ExternalInvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalInvoiceNumber

`func (o *InvoiceUpdateRequest) SetExternalInvoiceNumber(v string)`

SetExternalInvoiceNumber sets ExternalInvoiceNumber field to given value.

### HasExternalInvoiceNumber

`func (o *InvoiceUpdateRequest) HasExternalInvoiceNumber() bool`

HasExternalInvoiceNumber returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *InvoiceUpdateRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *InvoiceUpdateRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *InvoiceUpdateRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *InvoiceUpdateRequest) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDueDate

`func (o *InvoiceUpdateRequest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoiceUpdateRequest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoiceUpdateRequest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InvoiceUpdateRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetExternalRecipientId

`func (o *InvoiceUpdateRequest) GetExternalRecipientId() string`

GetExternalRecipientId returns the ExternalRecipientId field if non-nil, zero value otherwise.

### GetExternalRecipientIdOk

`func (o *InvoiceUpdateRequest) GetExternalRecipientIdOk() (*string, bool)`

GetExternalRecipientIdOk returns a tuple with the ExternalRecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRecipientId

`func (o *InvoiceUpdateRequest) SetExternalRecipientId(v string)`

SetExternalRecipientId sets ExternalRecipientId field to given value.

### HasExternalRecipientId

`func (o *InvoiceUpdateRequest) HasExternalRecipientId() bool`

HasExternalRecipientId returns a boolean if a field has been set.

### GetReceivedByRecipientDate

`func (o *InvoiceUpdateRequest) GetReceivedByRecipientDate() int64`

GetReceivedByRecipientDate returns the ReceivedByRecipientDate field if non-nil, zero value otherwise.

### GetReceivedByRecipientDateOk

`func (o *InvoiceUpdateRequest) GetReceivedByRecipientDateOk() (*int64, bool)`

GetReceivedByRecipientDateOk returns a tuple with the ReceivedByRecipientDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedByRecipientDate

`func (o *InvoiceUpdateRequest) SetReceivedByRecipientDate(v int64)`

SetReceivedByRecipientDate sets ReceivedByRecipientDate field to given value.

### HasReceivedByRecipientDate

`func (o *InvoiceUpdateRequest) HasReceivedByRecipientDate() bool`

HasReceivedByRecipientDate returns a boolean if a field has been set.

### GetIsVoided

`func (o *InvoiceUpdateRequest) GetIsVoided() bool`

GetIsVoided returns the IsVoided field if non-nil, zero value otherwise.

### GetIsVoidedOk

`func (o *InvoiceUpdateRequest) GetIsVoidedOk() (*bool, bool)`

GetIsVoidedOk returns a tuple with the IsVoided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVoided

`func (o *InvoiceUpdateRequest) SetIsVoided(v bool)`

SetIsVoided sets IsVoided field to given value.

### HasIsVoided

`func (o *InvoiceUpdateRequest) HasIsVoided() bool`

HasIsVoided returns a boolean if a field has been set.

### GetReceivedByCustomerDate

`func (o *InvoiceUpdateRequest) GetReceivedByCustomerDate() string`

GetReceivedByCustomerDate returns the ReceivedByCustomerDate field if non-nil, zero value otherwise.

### GetReceivedByCustomerDateOk

`func (o *InvoiceUpdateRequest) GetReceivedByCustomerDateOk() (*string, bool)`

GetReceivedByCustomerDateOk returns a tuple with the ReceivedByCustomerDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedByCustomerDate

`func (o *InvoiceUpdateRequest) SetReceivedByCustomerDate(v string)`

SetReceivedByCustomerDate sets ReceivedByCustomerDate field to given value.

### HasReceivedByCustomerDate

`func (o *InvoiceUpdateRequest) HasReceivedByCustomerDate() bool`

HasReceivedByCustomerDate returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *InvoiceUpdateRequest) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *InvoiceUpdateRequest) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *InvoiceUpdateRequest) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *InvoiceUpdateRequest) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


