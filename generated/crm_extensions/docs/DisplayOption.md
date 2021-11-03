# DisplayOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | JSON-friendly unique name for option. | 
**Label** | **string** | The text that will be displayed to users for this option. | 
**Type** | **string** | The type of status. | 

## Methods

### NewDisplayOption

`func NewDisplayOption(name string, label string, type_ string, ) *DisplayOption`

NewDisplayOption instantiates a new DisplayOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisplayOptionWithDefaults

`func NewDisplayOptionWithDefaults() *DisplayOption`

NewDisplayOptionWithDefaults instantiates a new DisplayOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DisplayOption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DisplayOption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DisplayOption) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *DisplayOption) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DisplayOption) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DisplayOption) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *DisplayOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DisplayOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DisplayOption) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


