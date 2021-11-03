# InvoiceReadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalInvoiceNumber** | Pointer to **string** | The invoice number. Note that this is _not_ the ID of the invoice, but the number that the billed customer will see. | [optional] 
**TotalAmountBilled** | **float32** | The total amount that this invoice is for. | 
**BalanceDue** | **float32** | The remaining balance due for the invoice. | 
**CurrencyCode** | **string** | The ISO 4217 currency code that represents the currency of the invoice. | 
**DueDate** | **string** | The due date of the invoice | 
**ExternalRecipientId** | **string** | The id of the customer in the external accounting system that the invoice was sent to. | 
**ReceivedByRecipientDate** | Pointer to **int64** | The datetime that that invoice was sent to the customer. | [optional] 
**ExternalCreateDateTime** | Pointer to **int64** | The datetime that the invoice was created in the external accounting system. | [optional] 
**IsVoided** | **bool** | Indicated is the invoice has been voided or not. | 
**CreatedAt** | **time.Time** | The datetime that the invoice was created in HubSpot. | 
**UpdatedAt** | **time.Time** | The datetime that the invoice was last updated in HubSpot. | 
**ArchivedAt** | Pointer to **time.Time** |  | [optional] 
**Archived** | **bool** |  | 
**ExternalAccountId** | **string** | The id of the account in the external accounting system that this invoice is related to. | 
**InvoiceStatus** | **string** | The status of the invoice | 
**Id** | **string** | The id of this invoice in the external accounting system. | 

## Methods

### NewInvoiceReadResponse

`func NewInvoiceReadResponse(totalAmountBilled float32, balanceDue float32, currencyCode string, dueDate string, externalRecipientId string, isVoided bool, createdAt time.Time, updatedAt time.Time, archived bool, externalAccountId string, invoiceStatus string, id string, ) *InvoiceReadResponse`

NewInvoiceReadResponse instantiates a new InvoiceReadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceReadResponseWithDefaults

`func NewInvoiceReadResponseWithDefaults() *InvoiceReadResponse`

NewInvoiceReadResponseWithDefaults instantiates a new InvoiceReadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalInvoiceNumber

`func (o *InvoiceReadResponse) GetExternalInvoiceNumber() string`

GetExternalInvoiceNumber returns the ExternalInvoiceNumber field if non-nil, zero value otherwise.

### GetExternalInvoiceNumberOk

`func (o *InvoiceReadResponse) GetExternalInvoiceNumberOk() (*string, bool)`

GetExternalInvoiceNumberOk returns a tuple with the ExternalInvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalInvoiceNumber

`func (o *InvoiceReadResponse) SetExternalInvoiceNumber(v string)`

SetExternalInvoiceNumber sets ExternalInvoiceNumber field to given value.

### HasExternalInvoiceNumber

`func (o *InvoiceReadResponse) HasExternalInvoiceNumber() bool`

HasExternalInvoiceNumber returns a boolean if a field has been set.

### GetTotalAmountBilled

`func (o *InvoiceReadResponse) GetTotalAmountBilled() float32`

GetTotalAmountBilled returns the TotalAmountBilled field if non-nil, zero value otherwise.

### GetTotalAmountBilledOk

`func (o *InvoiceReadResponse) GetTotalAmountBilledOk() (*float32, bool)`

GetTotalAmountBilledOk returns a tuple with the TotalAmountBilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountBilled

`func (o *InvoiceReadResponse) SetTotalAmountBilled(v float32)`

SetTotalAmountBilled sets TotalAmountBilled field to given value.


### GetBalanceDue

`func (o *InvoiceReadResponse) GetBalanceDue() float32`

GetBalanceDue returns the BalanceDue field if non-nil, zero value otherwise.

### GetBalanceDueOk

`func (o *InvoiceReadResponse) GetBalanceDueOk() (*float32, bool)`

GetBalanceDueOk returns a tuple with the BalanceDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceDue

`func (o *InvoiceReadResponse) SetBalanceDue(v float32)`

SetBalanceDue sets BalanceDue field to given value.


### GetCurrencyCode

`func (o *InvoiceReadResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *InvoiceReadResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *InvoiceReadResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetDueDate

`func (o *InvoiceReadResponse) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoiceReadResponse) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoiceReadResponse) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.


### GetExternalRecipientId

`func (o *InvoiceReadResponse) GetExternalRecipientId() string`

GetExternalRecipientId returns the ExternalRecipientId field if non-nil, zero value otherwise.

### GetExternalRecipientIdOk

`func (o *InvoiceReadResponse) GetExternalRecipientIdOk() (*string, bool)`

GetExternalRecipientIdOk returns a tuple with the ExternalRecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRecipientId

`func (o *InvoiceReadResponse) SetExternalRecipientId(v string)`

SetExternalRecipientId sets ExternalRecipientId field to given value.


### GetReceivedByRecipientDate

`func (o *InvoiceReadResponse) GetReceivedByRecipientDate() int64`

GetReceivedByRecipientDate returns the ReceivedByRecipientDate field if non-nil, zero value otherwise.

### GetReceivedByRecipientDateOk

`func (o *InvoiceReadResponse) GetReceivedByRecipientDateOk() (*int64, bool)`

GetReceivedByRecipientDateOk returns a tuple with the ReceivedByRecipientDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedByRecipientDate

`func (o *InvoiceReadResponse) SetReceivedByRecipientDate(v int64)`

SetReceivedByRecipientDate sets ReceivedByRecipientDate field to given value.

### HasReceivedByRecipientDate

`func (o *InvoiceReadResponse) HasReceivedByRecipientDate() bool`

HasReceivedByRecipientDate returns a boolean if a field has been set.

### GetExternalCreateDateTime

`func (o *InvoiceReadResponse) GetExternalCreateDateTime() int64`

GetExternalCreateDateTime returns the ExternalCreateDateTime field if non-nil, zero value otherwise.

### GetExternalCreateDateTimeOk

`func (o *InvoiceReadResponse) GetExternalCreateDateTimeOk() (*int64, bool)`

GetExternalCreateDateTimeOk returns a tuple with the ExternalCreateDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCreateDateTime

`func (o *InvoiceReadResponse) SetExternalCreateDateTime(v int64)`

SetExternalCreateDateTime sets ExternalCreateDateTime field to given value.

### HasExternalCreateDateTime

`func (o *InvoiceReadResponse) HasExternalCreateDateTime() bool`

HasExternalCreateDateTime returns a boolean if a field has been set.

### GetIsVoided

`func (o *InvoiceReadResponse) GetIsVoided() bool`

GetIsVoided returns the IsVoided field if non-nil, zero value otherwise.

### GetIsVoidedOk

`func (o *InvoiceReadResponse) GetIsVoidedOk() (*bool, bool)`

GetIsVoidedOk returns a tuple with the IsVoided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVoided

`func (o *InvoiceReadResponse) SetIsVoided(v bool)`

SetIsVoided sets IsVoided field to given value.


### GetCreatedAt

`func (o *InvoiceReadResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceReadResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceReadResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *InvoiceReadResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvoiceReadResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvoiceReadResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetArchivedAt

`func (o *InvoiceReadResponse) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *InvoiceReadResponse) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *InvoiceReadResponse) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *InvoiceReadResponse) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetArchived

`func (o *InvoiceReadResponse) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *InvoiceReadResponse) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *InvoiceReadResponse) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetExternalAccountId

`func (o *InvoiceReadResponse) GetExternalAccountId() string`

GetExternalAccountId returns the ExternalAccountId field if non-nil, zero value otherwise.

### GetExternalAccountIdOk

`func (o *InvoiceReadResponse) GetExternalAccountIdOk() (*string, bool)`

GetExternalAccountIdOk returns a tuple with the ExternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAccountId

`func (o *InvoiceReadResponse) SetExternalAccountId(v string)`

SetExternalAccountId sets ExternalAccountId field to given value.


### GetInvoiceStatus

`func (o *InvoiceReadResponse) GetInvoiceStatus() string`

GetInvoiceStatus returns the InvoiceStatus field if non-nil, zero value otherwise.

### GetInvoiceStatusOk

`func (o *InvoiceReadResponse) GetInvoiceStatusOk() (*string, bool)`

GetInvoiceStatusOk returns a tuple with the InvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatus

`func (o *InvoiceReadResponse) SetInvoiceStatus(v string)`

SetInvoiceStatus sets InvoiceStatus field to given value.


### GetId

`func (o *InvoiceReadResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceReadResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceReadResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


