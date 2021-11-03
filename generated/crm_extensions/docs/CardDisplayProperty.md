# CardDisplayProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | An internal identifier for this property. This value must be unique TODO. | 
**Label** | **string** | The label for this property as you&#39;d like it displayed to users. | 
**DataType** | **string** | Type of data represented by this property. | 
**Options** | [**[]DisplayOption**](DisplayOption.md) | An array of available options that can be displayed. Only used in when &#x60;dataType&#x60; is &#x60;STATUS&#x60;. | 

## Methods

### NewCardDisplayProperty

`func NewCardDisplayProperty(name string, label string, dataType string, options []DisplayOption, ) *CardDisplayProperty`

NewCardDisplayProperty instantiates a new CardDisplayProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardDisplayPropertyWithDefaults

`func NewCardDisplayPropertyWithDefaults() *CardDisplayProperty`

NewCardDisplayPropertyWithDefaults instantiates a new CardDisplayProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CardDisplayProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardDisplayProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CardDisplayProperty) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *CardDisplayProperty) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CardDisplayProperty) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CardDisplayProperty) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDataType

`func (o *CardDisplayProperty) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *CardDisplayProperty) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *CardDisplayProperty) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetOptions

`func (o *CardDisplayProperty) GetOptions() []DisplayOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CardDisplayProperty) GetOptionsOk() (*[]DisplayOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CardDisplayProperty) SetOptions(v []DisplayOption)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


