# Tax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | The code/ID of the tax in the external accounting system. | 
**Percentage** | **float32** | The tax percentage.  For example, 8.05 represents a 8.05% tax rate. | 
**Name** | **string** | The display name of the tax. | 

## Methods

### NewTax

`func NewTax(code string, percentage float32, name string, ) *Tax`

NewTax instantiates a new Tax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxWithDefaults

`func NewTaxWithDefaults() *Tax`

NewTaxWithDefaults instantiates a new Tax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Tax) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Tax) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Tax) SetCode(v string)`

SetCode sets Code field to given value.


### GetPercentage

`func (o *Tax) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *Tax) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *Tax) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.


### GetName

`func (o *Tax) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tax) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tax) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


