# PropertyGroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | A human-readable label that will be shown in HubSpot. | [optional] 
**DisplayOrder** | Pointer to **int32** | Property groups are displayed in order starting with the lowest positive integer value. Values of -1 will cause the property group to be displayed after any positive values. | [optional] 

## Methods

### NewPropertyGroupUpdate

`func NewPropertyGroupUpdate() *PropertyGroupUpdate`

NewPropertyGroupUpdate instantiates a new PropertyGroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyGroupUpdateWithDefaults

`func NewPropertyGroupUpdateWithDefaults() *PropertyGroupUpdate`

NewPropertyGroupUpdateWithDefaults instantiates a new PropertyGroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *PropertyGroupUpdate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PropertyGroupUpdate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PropertyGroupUpdate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PropertyGroupUpdate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *PropertyGroupUpdate) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *PropertyGroupUpdate) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *PropertyGroupUpdate) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *PropertyGroupUpdate) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


