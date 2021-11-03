# PropertyUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | A human-readable property label that will be shown in HubSpot. | [optional] 
**Type** | Pointer to **string** | The data type of the property. | [optional] 
**FieldType** | Pointer to **string** | Controls how the property appears in HubSpot. | [optional] 
**GroupName** | Pointer to **string** | The name of the property group the property belongs to. | [optional] 
**Description** | Pointer to **string** | A description of the property that will be shown as help text in HubSpot. | [optional] 
**Options** | Pointer to [**[]OptionInput**](OptionInput.md) | A list of valid options for the property. | [optional] 
**DisplayOrder** | Pointer to **int32** | Properties are displayed in order starting with the lowest positive integer value. Values of -1 will cause the Property to be displayed after any positive values. | [optional] 
**Hidden** | Pointer to **bool** | If true, the property won&#39;t be visible and can&#39;t be used in HubSpot. | [optional] 
**FormField** | Pointer to **bool** | Whether or not the property can be used in a HubSpot form. | [optional] 

## Methods

### NewPropertyUpdate

`func NewPropertyUpdate() *PropertyUpdate`

NewPropertyUpdate instantiates a new PropertyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyUpdateWithDefaults

`func NewPropertyUpdateWithDefaults() *PropertyUpdate`

NewPropertyUpdateWithDefaults instantiates a new PropertyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *PropertyUpdate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PropertyUpdate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PropertyUpdate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PropertyUpdate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *PropertyUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PropertyUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PropertyUpdate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PropertyUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFieldType

`func (o *PropertyUpdate) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *PropertyUpdate) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *PropertyUpdate) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *PropertyUpdate) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### GetGroupName

`func (o *PropertyUpdate) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *PropertyUpdate) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *PropertyUpdate) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *PropertyUpdate) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetDescription

`func (o *PropertyUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PropertyUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PropertyUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PropertyUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOptions

`func (o *PropertyUpdate) GetOptions() []OptionInput`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PropertyUpdate) GetOptionsOk() (*[]OptionInput, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PropertyUpdate) SetOptions(v []OptionInput)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PropertyUpdate) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *PropertyUpdate) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *PropertyUpdate) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *PropertyUpdate) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *PropertyUpdate) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetHidden

`func (o *PropertyUpdate) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *PropertyUpdate) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *PropertyUpdate) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *PropertyUpdate) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetFormField

`func (o *PropertyUpdate) GetFormField() bool`

GetFormField returns the FormField field if non-nil, zero value otherwise.

### GetFormFieldOk

`func (o *PropertyUpdate) GetFormFieldOk() (*bool, bool)`

GetFormFieldOk returns a tuple with the FormField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormField

`func (o *PropertyUpdate) SetFormField(v bool)`

SetFormField sets FormField field to given value.

### HasFormField

`func (o *PropertyUpdate) HasFormField() bool`

HasFormField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


