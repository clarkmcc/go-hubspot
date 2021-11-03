# PropertyGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The internal property group name, which must be used when referencing the property group via the API. | 
**Label** | **string** | A human-readable label that will be shown in HubSpot. | 
**DisplayOrder** | **int32** | Property groups are displayed in order starting with the lowest positive integer value. Values of -1 will cause the property group to be displayed after any positive values. | 
**Archived** | **bool** |  | 

## Methods

### NewPropertyGroup

`func NewPropertyGroup(name string, label string, displayOrder int32, archived bool, ) *PropertyGroup`

NewPropertyGroup instantiates a new PropertyGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyGroupWithDefaults

`func NewPropertyGroupWithDefaults() *PropertyGroup`

NewPropertyGroupWithDefaults instantiates a new PropertyGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PropertyGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PropertyGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PropertyGroup) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *PropertyGroup) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PropertyGroup) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PropertyGroup) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDisplayOrder

`func (o *PropertyGroup) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *PropertyGroup) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *PropertyGroup) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.


### GetArchived

`func (o *PropertyGroup) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *PropertyGroup) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *PropertyGroup) SetArchived(v bool)`

SetArchived sets Archived field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


