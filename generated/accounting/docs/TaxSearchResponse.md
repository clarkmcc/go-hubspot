# TaxSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **string** | Designates if the response is a success (&#39;OK&#39;) or failure (&#39;ERR&#39;). | [optional] 
**Taxes** | [**[]Tax**](Tax.md) | The list of taxes that matched the search criteria | 

## Methods

### NewTaxSearchResponse

`func NewTaxSearchResponse(taxes []Tax, ) *TaxSearchResponse`

NewTaxSearchResponse instantiates a new TaxSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxSearchResponseWithDefaults

`func NewTaxSearchResponseWithDefaults() *TaxSearchResponse`

NewTaxSearchResponseWithDefaults instantiates a new TaxSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *TaxSearchResponse) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TaxSearchResponse) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TaxSearchResponse) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *TaxSearchResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetTaxes

`func (o *TaxSearchResponse) GetTaxes() []Tax`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *TaxSearchResponse) GetTaxesOk() (*[]Tax, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *TaxSearchResponse) SetTaxes(v []Tax)`

SetTaxes sets Taxes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


