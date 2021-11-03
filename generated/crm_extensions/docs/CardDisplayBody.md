# CardDisplayBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | [**[]CardDisplayProperty**](CardDisplayProperty.md) | Card display properties. These will will be rendered as \&quot;label : value\&quot; pairs in the card UI. See the [example card](#) in the overview docs for more details. | 

## Methods

### NewCardDisplayBody

`func NewCardDisplayBody(properties []CardDisplayProperty, ) *CardDisplayBody`

NewCardDisplayBody instantiates a new CardDisplayBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardDisplayBodyWithDefaults

`func NewCardDisplayBodyWithDefaults() *CardDisplayBody`

NewCardDisplayBodyWithDefaults instantiates a new CardDisplayBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *CardDisplayBody) GetProperties() []CardDisplayProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CardDisplayBody) GetPropertiesOk() (*[]CardDisplayProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CardDisplayBody) SetProperties(v []CardDisplayProperty)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


