# CreateUserAccountRequestExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The id of the account in your system. | 
**AccountName** | **string** | The name of the account in your system. This is normally the name visible to your users. | 
**CurrencyCode** | **string** | The default currency that this account uses. | 

## Methods

### NewCreateUserAccountRequestExternal

`func NewCreateUserAccountRequestExternal(accountId string, accountName string, currencyCode string, ) *CreateUserAccountRequestExternal`

NewCreateUserAccountRequestExternal instantiates a new CreateUserAccountRequestExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserAccountRequestExternalWithDefaults

`func NewCreateUserAccountRequestExternalWithDefaults() *CreateUserAccountRequestExternal`

NewCreateUserAccountRequestExternalWithDefaults instantiates a new CreateUserAccountRequestExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CreateUserAccountRequestExternal) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateUserAccountRequestExternal) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateUserAccountRequestExternal) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAccountName

`func (o *CreateUserAccountRequestExternal) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *CreateUserAccountRequestExternal) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *CreateUserAccountRequestExternal) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetCurrencyCode

`func (o *CreateUserAccountRequestExternal) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CreateUserAccountRequestExternal) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CreateUserAccountRequestExternal) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


