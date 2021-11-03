# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | The country of the address. | [optional] 
**CountrySubDivisionCode** | Pointer to **string** | A region of the county of the address.  May represent county, state etc. | [optional] 
**City** | Pointer to **string** | The city of the address. | [optional] 
**PostalCode** | Pointer to **string** | The postcode/zipcode of the address. | [optional] 
**LineOne** | Pointer to **string** | The first line of the address. | [optional] 

## Methods

### NewAddress

`func NewAddress() *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *Address) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Address) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Address) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Address) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountrySubDivisionCode

`func (o *Address) GetCountrySubDivisionCode() string`

GetCountrySubDivisionCode returns the CountrySubDivisionCode field if non-nil, zero value otherwise.

### GetCountrySubDivisionCodeOk

`func (o *Address) GetCountrySubDivisionCodeOk() (*string, bool)`

GetCountrySubDivisionCodeOk returns a tuple with the CountrySubDivisionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountrySubDivisionCode

`func (o *Address) SetCountrySubDivisionCode(v string)`

SetCountrySubDivisionCode sets CountrySubDivisionCode field to given value.

### HasCountrySubDivisionCode

`func (o *Address) HasCountrySubDivisionCode() bool`

HasCountrySubDivisionCode returns a boolean if a field has been set.

### GetCity

`func (o *Address) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Address) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Address) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Address) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetPostalCode

`func (o *Address) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Address) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Address) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *Address) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetLineOne

`func (o *Address) GetLineOne() string`

GetLineOne returns the LineOne field if non-nil, zero value otherwise.

### GetLineOneOk

`func (o *Address) GetLineOneOk() (*string, bool)`

GetLineOneOk returns a tuple with the LineOne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineOne

`func (o *Address) SetLineOne(v string)`

SetLineOne sets LineOne field to given value.

### HasLineOne

`func (o *Address) HasLineOne() bool`

HasLineOne returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


