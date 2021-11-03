# UnitPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | The actual unit price amount. | [optional] 
**TaxIncluded** | **bool** | Indicates if the unit price amount already includes taxes. | 

## Methods

### NewUnitPrice

`func NewUnitPrice(taxIncluded bool, ) *UnitPrice`

NewUnitPrice instantiates a new UnitPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnitPriceWithDefaults

`func NewUnitPriceWithDefaults() *UnitPrice`

NewUnitPriceWithDefaults instantiates a new UnitPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UnitPrice) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnitPrice) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnitPrice) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnitPrice) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTaxIncluded

`func (o *UnitPrice) GetTaxIncluded() bool`

GetTaxIncluded returns the TaxIncluded field if non-nil, zero value otherwise.

### GetTaxIncludedOk

`func (o *UnitPrice) GetTaxIncludedOk() (*bool, bool)`

GetTaxIncludedOk returns a tuple with the TaxIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIncluded

`func (o *UnitPrice) SetTaxIncluded(v bool)`

SetTaxIncluded sets TaxIncluded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


