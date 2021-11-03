# TaxType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | The code/ID of the tax in the external accounting system. | 
**Name** | Pointer to **string** | The display name of the tax. | [optional] 

## Methods

### NewTaxType

`func NewTaxType(code string, ) *TaxType`

NewTaxType instantiates a new TaxType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxTypeWithDefaults

`func NewTaxTypeWithDefaults() *TaxType`

NewTaxTypeWithDefaults instantiates a new TaxType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TaxType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TaxType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TaxType) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *TaxType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaxType) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


