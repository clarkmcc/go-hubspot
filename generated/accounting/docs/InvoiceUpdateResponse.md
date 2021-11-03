# InvoiceUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalInvoiceNumber** | Pointer to **string** |  | [optional] 
**TotalAmountBilled** | **float32** |  | 
**BalanceDue** | **float32** |  | 
**CurrencyCode** | **string** |  | 
**DueDate** | **string** |  | 
**ExternalRecipientId** | **string** |  | 
**ReceivedByRecipientDate** | Pointer to **int64** |  | [optional] 
**ExternalCreateDateTime** | Pointer to **int64** |  | [optional] 
**IsVoided** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**ArchivedAt** | Pointer to **time.Time** |  | [optional] 
**Archived** | **bool** |  | 
**ExternalAccountId** | **string** |  | 
**InvoiceStatus** | **string** |  | 
**Id** | **string** |  | 

## Methods

### NewInvoiceUpdateResponse

`func NewInvoiceUpdateResponse(totalAmountBilled float32, balanceDue float32, currencyCode string, dueDate string, externalRecipientId string, isVoided bool, createdAt time.Time, updatedAt time.Time, archived bool, externalAccountId string, invoiceStatus string, id string, ) *InvoiceUpdateResponse`

NewInvoiceUpdateResponse instantiates a new InvoiceUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceUpdateResponseWithDefaults

`func NewInvoiceUpdateResponseWithDefaults() *InvoiceUpdateResponse`

NewInvoiceUpdateResponseWithDefaults instantiates a new InvoiceUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalInvoiceNumber

`func (o *InvoiceUpdateResponse) GetExternalInvoiceNumber() string`

GetExternalInvoiceNumber returns the ExternalInvoiceNumber field if non-nil, zero value otherwise.

### GetExternalInvoiceNumberOk

`func (o *InvoiceUpdateResponse) GetExternalInvoiceNumberOk() (*string, bool)`

GetExternalInvoiceNumberOk returns a tuple with the ExternalInvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalInvoiceNumber

`func (o *InvoiceUpdateResponse) SetExternalInvoiceNumber(v string)`

SetExternalInvoiceNumber sets ExternalInvoiceNumber field to given value.

### HasExternalInvoiceNumber

`func (o *InvoiceUpdateResponse) HasExternalInvoiceNumber() bool`

HasExternalInvoiceNumber returns a boolean if a field has been set.

### GetTotalAmountBilled

`func (o *InvoiceUpdateResponse) GetTotalAmountBilled() float32`

GetTotalAmountBilled returns the TotalAmountBilled field if non-nil, zero value otherwise.

### GetTotalAmountBilledOk

`func (o *InvoiceUpdateResponse) GetTotalAmountBilledOk() (*float32, bool)`

GetTotalAmountBilledOk returns a tuple with the TotalAmountBilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountBilled

`func (o *InvoiceUpdateResponse) SetTotalAmountBilled(v float32)`

SetTotalAmountBilled sets TotalAmountBilled field to given value.


### GetBalanceDue

`func (o *InvoiceUpdateResponse) GetBalanceDue() float32`

GetBalanceDue returns the BalanceDue field if non-nil, zero value otherwise.

### GetBalanceDueOk

`func (o *InvoiceUpdateResponse) GetBalanceDueOk() (*float32, bool)`

GetBalanceDueOk returns a tuple with the BalanceDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceDue

`func (o *InvoiceUpdateResponse) SetBalanceDue(v float32)`

SetBalanceDue sets BalanceDue field to given value.


### GetCurrencyCode

`func (o *InvoiceUpdateResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *InvoiceUpdateResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *InvoiceUpdateResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetDueDate

`func (o *InvoiceUpdateResponse) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoiceUpdateResponse) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoiceUpdateResponse) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.


### GetExternalRecipientId

`func (o *InvoiceUpdateResponse) GetExternalRecipientId() string`

GetExternalRecipientId returns the ExternalRecipientId field if non-nil, zero value otherwise.

### GetExternalRecipientIdOk

`func (o *InvoiceUpdateResponse) GetExternalRecipientIdOk() (*string, bool)`

GetExternalRecipientIdOk returns a tuple with the ExternalRecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRecipientId

`func (o *InvoiceUpdateResponse) SetExternalRecipientId(v string)`

SetExternalRecipientId sets ExternalRecipientId field to given value.


### GetReceivedByRecipientDate

`func (o *InvoiceUpdateResponse) GetReceivedByRecipientDate() int64`

GetReceivedByRecipientDate returns the ReceivedByRecipientDate field if non-nil, zero value otherwise.

### GetReceivedByRecipientDateOk

`func (o *InvoiceUpdateResponse) GetReceivedByRecipientDateOk() (*int64, bool)`

GetReceivedByRecipientDateOk returns a tuple with the ReceivedByRecipientDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedByRecipientDate

`func (o *InvoiceUpdateResponse) SetReceivedByRecipientDate(v int64)`

SetReceivedByRecipientDate sets ReceivedByRecipientDate field to given value.

### HasReceivedByRecipientDate

`func (o *InvoiceUpdateResponse) HasReceivedByRecipientDate() bool`

HasReceivedByRecipientDate returns a boolean if a field has been set.

### GetExternalCreateDateTime

`func (o *InvoiceUpdateResponse) GetExternalCreateDateTime() int64`

GetExternalCreateDateTime returns the ExternalCreateDateTime field if non-nil, zero value otherwise.

### GetExternalCreateDateTimeOk

`func (o *InvoiceUpdateResponse) GetExternalCreateDateTimeOk() (*int64, bool)`

GetExternalCreateDateTimeOk returns a tuple with the ExternalCreateDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCreateDateTime

`func (o *InvoiceUpdateResponse) SetExternalCreateDateTime(v int64)`

SetExternalCreateDateTime sets ExternalCreateDateTime field to given value.

### HasExternalCreateDateTime

`func (o *InvoiceUpdateResponse) HasExternalCreateDateTime() bool`

HasExternalCreateDateTime returns a boolean if a field has been set.

### GetIsVoided

`func (o *InvoiceUpdateResponse) GetIsVoided() bool`

GetIsVoided returns the IsVoided field if non-nil, zero value otherwise.

### GetIsVoidedOk

`func (o *InvoiceUpdateResponse) GetIsVoidedOk() (*bool, bool)`

GetIsVoidedOk returns a tuple with the IsVoided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVoided

`func (o *InvoiceUpdateResponse) SetIsVoided(v bool)`

SetIsVoided sets IsVoided field to given value.


### GetCreatedAt

`func (o *InvoiceUpdateResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceUpdateResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceUpdateResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *InvoiceUpdateResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvoiceUpdateResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvoiceUpdateResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetArchivedAt

`func (o *InvoiceUpdateResponse) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *InvoiceUpdateResponse) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *InvoiceUpdateResponse) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *InvoiceUpdateResponse) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetArchived

`func (o *InvoiceUpdateResponse) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *InvoiceUpdateResponse) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *InvoiceUpdateResponse) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetExternalAccountId

`func (o *InvoiceUpdateResponse) GetExternalAccountId() string`

GetExternalAccountId returns the ExternalAccountId field if non-nil, zero value otherwise.

### GetExternalAccountIdOk

`func (o *InvoiceUpdateResponse) GetExternalAccountIdOk() (*string, bool)`

GetExternalAccountIdOk returns a tuple with the ExternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAccountId

`func (o *InvoiceUpdateResponse) SetExternalAccountId(v string)`

SetExternalAccountId sets ExternalAccountId field to given value.


### GetInvoiceStatus

`func (o *InvoiceUpdateResponse) GetInvoiceStatus() string`

GetInvoiceStatus returns the InvoiceStatus field if non-nil, zero value otherwise.

### GetInvoiceStatusOk

`func (o *InvoiceUpdateResponse) GetInvoiceStatusOk() (*string, bool)`

GetInvoiceStatusOk returns a tuple with the InvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatus

`func (o *InvoiceUpdateResponse) SetInvoiceStatus(v string)`

SetInvoiceStatus sets InvoiceStatus field to given value.


### GetId

`func (o *InvoiceUpdateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceUpdateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceUpdateResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


