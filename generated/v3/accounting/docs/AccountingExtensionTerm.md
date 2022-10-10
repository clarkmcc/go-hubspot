# AccountingExtensionTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DueDate** | Pointer to **string** | The due date for payment of the invoice, in ISO-8601 date format (yyyy-MM-dd) | [optional] 
**Name** | **string** | The display name of the payment terms. | 
**Id** | **string** | The ID of the payment terms in the external accounting system. | 
**DueDays** | Pointer to **int32** | The number of days that these payment terms represent. | [optional] 

## Methods

### NewAccountingExtensionTerm

`func NewAccountingExtensionTerm(name string, id string, ) *AccountingExtensionTerm`

NewAccountingExtensionTerm instantiates a new AccountingExtensionTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountingExtensionTermWithDefaults

`func NewAccountingExtensionTermWithDefaults() *AccountingExtensionTerm`

NewAccountingExtensionTermWithDefaults instantiates a new AccountingExtensionTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDueDate

`func (o *AccountingExtensionTerm) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *AccountingExtensionTerm) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *AccountingExtensionTerm) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *AccountingExtensionTerm) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetName

`func (o *AccountingExtensionTerm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountingExtensionTerm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountingExtensionTerm) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *AccountingExtensionTerm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountingExtensionTerm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountingExtensionTerm) SetId(v string)`

SetId sets Id field to given value.


### GetDueDays

`func (o *AccountingExtensionTerm) GetDueDays() int32`

GetDueDays returns the DueDays field if non-nil, zero value otherwise.

### GetDueDaysOk

`func (o *AccountingExtensionTerm) GetDueDaysOk() (*int32, bool)`

GetDueDaysOk returns a tuple with the DueDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDays

`func (o *AccountingExtensionTerm) SetDueDays(v int32)`

SetDueDays sets DueDays field to given value.

### HasDueDays

`func (o *AccountingExtensionTerm) HasDueDays() bool`

HasDueDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


