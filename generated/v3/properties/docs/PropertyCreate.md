# PropertyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hidden** | Pointer to **bool** | If true, the property won&#39;t be visible and can&#39;t be used in HubSpot. | [optional] 
**DisplayOrder** | Pointer to **int32** | Properties are displayed in order starting with the lowest positive integer value. Values of -1 will cause the property to be displayed after any positive values. | [optional] 
**Description** | Pointer to **string** | A description of the property that will be shown as help text in HubSpot. | [optional] 
**Label** | **string** | A human-readable property label that will be shown in HubSpot. | 
**Type** | **string** | The data type of the property. | 
**FormField** | Pointer to **bool** | Whether or not the property can be used in a HubSpot form. | [optional] 
**GroupName** | **string** | The name of the property group the property belongs to. | 
**ReferencedObjectType** | Pointer to **string** | Should be set to &#39;OWNER&#39; when &#39;externalOptions&#39; is true, which causes the property to dynamically pull option values from the current HubSpot users. | [optional] 
**Name** | **string** | The internal property name, which must be used when referencing the property via the API. | 
**Options** | Pointer to [**[]OptionInput**](OptionInput.md) | A list of valid options for the property. This field is required for enumerated properties. | [optional] 
**CalculationFormula** | Pointer to **string** | Represents a formula that is used to compute a calculated property. | [optional] 
**HasUniqueValue** | Pointer to **bool** | Whether or not the property&#39;s value must be unique. Once set, this can&#39;t be changed. | [optional] 
**FieldType** | **string** | Controls how the property appears in HubSpot. | 
**ExternalOptions** | Pointer to **bool** | Applicable only for &#39;enumeration&#39; type properties.  Should be set to true in conjunction with a &#39;referencedObjectType&#39; of &#39;OWNER&#39;.  Otherwise false. | [optional] 

## Methods

### NewPropertyCreate

`func NewPropertyCreate(label string, type_ string, groupName string, name string, fieldType string, ) *PropertyCreate`

NewPropertyCreate instantiates a new PropertyCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyCreateWithDefaults

`func NewPropertyCreateWithDefaults() *PropertyCreate`

NewPropertyCreateWithDefaults instantiates a new PropertyCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHidden

`func (o *PropertyCreate) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *PropertyCreate) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *PropertyCreate) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *PropertyCreate) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *PropertyCreate) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *PropertyCreate) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *PropertyCreate) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *PropertyCreate) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetDescription

`func (o *PropertyCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PropertyCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PropertyCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PropertyCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *PropertyCreate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PropertyCreate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PropertyCreate) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *PropertyCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PropertyCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PropertyCreate) SetType(v string)`

SetType sets Type field to given value.


### GetFormField

`func (o *PropertyCreate) GetFormField() bool`

GetFormField returns the FormField field if non-nil, zero value otherwise.

### GetFormFieldOk

`func (o *PropertyCreate) GetFormFieldOk() (*bool, bool)`

GetFormFieldOk returns a tuple with the FormField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormField

`func (o *PropertyCreate) SetFormField(v bool)`

SetFormField sets FormField field to given value.

### HasFormField

`func (o *PropertyCreate) HasFormField() bool`

HasFormField returns a boolean if a field has been set.

### GetGroupName

`func (o *PropertyCreate) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *PropertyCreate) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *PropertyCreate) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetReferencedObjectType

`func (o *PropertyCreate) GetReferencedObjectType() string`

GetReferencedObjectType returns the ReferencedObjectType field if non-nil, zero value otherwise.

### GetReferencedObjectTypeOk

`func (o *PropertyCreate) GetReferencedObjectTypeOk() (*string, bool)`

GetReferencedObjectTypeOk returns a tuple with the ReferencedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedObjectType

`func (o *PropertyCreate) SetReferencedObjectType(v string)`

SetReferencedObjectType sets ReferencedObjectType field to given value.

### HasReferencedObjectType

`func (o *PropertyCreate) HasReferencedObjectType() bool`

HasReferencedObjectType returns a boolean if a field has been set.

### GetName

`func (o *PropertyCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PropertyCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PropertyCreate) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *PropertyCreate) GetOptions() []OptionInput`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PropertyCreate) GetOptionsOk() (*[]OptionInput, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PropertyCreate) SetOptions(v []OptionInput)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PropertyCreate) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetCalculationFormula

`func (o *PropertyCreate) GetCalculationFormula() string`

GetCalculationFormula returns the CalculationFormula field if non-nil, zero value otherwise.

### GetCalculationFormulaOk

`func (o *PropertyCreate) GetCalculationFormulaOk() (*string, bool)`

GetCalculationFormulaOk returns a tuple with the CalculationFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationFormula

`func (o *PropertyCreate) SetCalculationFormula(v string)`

SetCalculationFormula sets CalculationFormula field to given value.

### HasCalculationFormula

`func (o *PropertyCreate) HasCalculationFormula() bool`

HasCalculationFormula returns a boolean if a field has been set.

### GetHasUniqueValue

`func (o *PropertyCreate) GetHasUniqueValue() bool`

GetHasUniqueValue returns the HasUniqueValue field if non-nil, zero value otherwise.

### GetHasUniqueValueOk

`func (o *PropertyCreate) GetHasUniqueValueOk() (*bool, bool)`

GetHasUniqueValueOk returns a tuple with the HasUniqueValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUniqueValue

`func (o *PropertyCreate) SetHasUniqueValue(v bool)`

SetHasUniqueValue sets HasUniqueValue field to given value.

### HasHasUniqueValue

`func (o *PropertyCreate) HasHasUniqueValue() bool`

HasHasUniqueValue returns a boolean if a field has been set.

### GetFieldType

`func (o *PropertyCreate) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *PropertyCreate) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *PropertyCreate) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.


### GetExternalOptions

`func (o *PropertyCreate) GetExternalOptions() bool`

GetExternalOptions returns the ExternalOptions field if non-nil, zero value otherwise.

### GetExternalOptionsOk

`func (o *PropertyCreate) GetExternalOptionsOk() (*bool, bool)`

GetExternalOptionsOk returns a tuple with the ExternalOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOptions

`func (o *PropertyCreate) SetExternalOptions(v bool)`

SetExternalOptions sets ExternalOptions field to given value.

### HasExternalOptions

`func (o *PropertyCreate) HasExternalOptions() bool`

HasExternalOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


