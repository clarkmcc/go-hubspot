# AccountingExtensionCustomer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to **string** | The customer&#39;s email address | [optional] 
**Name** | **string** | The customer&#39;s full name | 
**Id** | **string** | The ID of the customer in the external accounting system. | 
**BillingAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**CurrencyCode** | Pointer to **string** | The ISO 4217 currency code that represents the currency the customer should be billed in. | [optional] 

## Methods

### NewAccountingExtensionCustomer

`func NewAccountingExtensionCustomer(name string, id string, ) *AccountingExtensionCustomer`

NewAccountingExtensionCustomer instantiates a new AccountingExtensionCustomer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountingExtensionCustomerWithDefaults

`func NewAccountingExtensionCustomerWithDefaults() *AccountingExtensionCustomer`

NewAccountingExtensionCustomerWithDefaults instantiates a new AccountingExtensionCustomer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *AccountingExtensionCustomer) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *AccountingExtensionCustomer) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *AccountingExtensionCustomer) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *AccountingExtensionCustomer) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetName

`func (o *AccountingExtensionCustomer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountingExtensionCustomer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountingExtensionCustomer) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *AccountingExtensionCustomer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountingExtensionCustomer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountingExtensionCustomer) SetId(v string)`

SetId sets Id field to given value.


### GetBillingAddress

`func (o *AccountingExtensionCustomer) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *AccountingExtensionCustomer) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *AccountingExtensionCustomer) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *AccountingExtensionCustomer) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *AccountingExtensionCustomer) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AccountingExtensionCustomer) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AccountingExtensionCustomer) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *AccountingExtensionCustomer) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


