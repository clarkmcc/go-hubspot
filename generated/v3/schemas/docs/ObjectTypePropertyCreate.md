# ObjectTypePropertyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hidden** | Pointer to **bool** |  | [optional] 
**OptionSortStrategy** | Pointer to **string** | Controls how the property options will be sorted in the HubSpot UI. | [optional] 
**DisplayOrder** | Pointer to **int32** | The order that this property should be displayed in the HubSpot UI relative to other properties for this object type. Properties are displayed in order starting with the lowest positive integer value. A value of -1 will cause the property to be displayed **after** any positive values. | [optional] 
**Description** | Pointer to **string** | A description of the property that will be shown as help text in HubSpot. | [optional] 
**ShowCurrencySymbol** | Pointer to **bool** | Whether the property will display the currency symbol in the HubSpot UI. | [optional] 
**Label** | **string** | A human-readable property label that will be shown in HubSpot. | 
**Type** | **string** | The data type of the property. | 
**FormField** | Pointer to **bool** | Whether the property can be used in a HubSpot form. | [optional] 
**GroupName** | Pointer to **string** | The name of the group this property belongs to. | [optional] 
**ReferencedObjectType** | Pointer to **string** | Defines the options this property will return, e.g. OWNER would return name of users on the portal. | [optional] 
**TextDisplayHint** | Pointer to **string** | Controls how text properties are formatted in the HubSpot UI | [optional] 
**Name** | **string** | The internal property name, which must be used when referencing the property from the API. | 
**Options** | Pointer to [**[]OptionInput**](OptionInput.md) | A list of available options for the property. This field is only required for enumerated properties. | [optional] 
**SearchableInGlobalSearch** | Pointer to **bool** | Allow users to search for information entered to this field (limited to 3 properties) | [optional] 
**NumberDisplayHint** | Pointer to **string** | Controls how numeric properties are formatted in the HubSpot UI | [optional] 
**HasUniqueValue** | Pointer to **bool** | Whether or not the property&#39;s value must be unique. Once set, this can&#39;t be changed. | [optional] 
**FieldType** | **string** | Controls how the property appears in HubSpot. | 

## Methods

### NewObjectTypePropertyCreate

`func NewObjectTypePropertyCreate(label string, type_ string, name string, fieldType string, ) *ObjectTypePropertyCreate`

NewObjectTypePropertyCreate instantiates a new ObjectTypePropertyCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectTypePropertyCreateWithDefaults

`func NewObjectTypePropertyCreateWithDefaults() *ObjectTypePropertyCreate`

NewObjectTypePropertyCreateWithDefaults instantiates a new ObjectTypePropertyCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHidden

`func (o *ObjectTypePropertyCreate) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ObjectTypePropertyCreate) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ObjectTypePropertyCreate) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ObjectTypePropertyCreate) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetOptionSortStrategy

`func (o *ObjectTypePropertyCreate) GetOptionSortStrategy() string`

GetOptionSortStrategy returns the OptionSortStrategy field if non-nil, zero value otherwise.

### GetOptionSortStrategyOk

`func (o *ObjectTypePropertyCreate) GetOptionSortStrategyOk() (*string, bool)`

GetOptionSortStrategyOk returns a tuple with the OptionSortStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSortStrategy

`func (o *ObjectTypePropertyCreate) SetOptionSortStrategy(v string)`

SetOptionSortStrategy sets OptionSortStrategy field to given value.

### HasOptionSortStrategy

`func (o *ObjectTypePropertyCreate) HasOptionSortStrategy() bool`

HasOptionSortStrategy returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *ObjectTypePropertyCreate) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ObjectTypePropertyCreate) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ObjectTypePropertyCreate) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ObjectTypePropertyCreate) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetDescription

`func (o *ObjectTypePropertyCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectTypePropertyCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectTypePropertyCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectTypePropertyCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetShowCurrencySymbol

`func (o *ObjectTypePropertyCreate) GetShowCurrencySymbol() bool`

GetShowCurrencySymbol returns the ShowCurrencySymbol field if non-nil, zero value otherwise.

### GetShowCurrencySymbolOk

`func (o *ObjectTypePropertyCreate) GetShowCurrencySymbolOk() (*bool, bool)`

GetShowCurrencySymbolOk returns a tuple with the ShowCurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCurrencySymbol

`func (o *ObjectTypePropertyCreate) SetShowCurrencySymbol(v bool)`

SetShowCurrencySymbol sets ShowCurrencySymbol field to given value.

### HasShowCurrencySymbol

`func (o *ObjectTypePropertyCreate) HasShowCurrencySymbol() bool`

HasShowCurrencySymbol returns a boolean if a field has been set.

### GetLabel

`func (o *ObjectTypePropertyCreate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ObjectTypePropertyCreate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ObjectTypePropertyCreate) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *ObjectTypePropertyCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ObjectTypePropertyCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ObjectTypePropertyCreate) SetType(v string)`

SetType sets Type field to given value.


### GetFormField

`func (o *ObjectTypePropertyCreate) GetFormField() bool`

GetFormField returns the FormField field if non-nil, zero value otherwise.

### GetFormFieldOk

`func (o *ObjectTypePropertyCreate) GetFormFieldOk() (*bool, bool)`

GetFormFieldOk returns a tuple with the FormField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormField

`func (o *ObjectTypePropertyCreate) SetFormField(v bool)`

SetFormField sets FormField field to given value.

### HasFormField

`func (o *ObjectTypePropertyCreate) HasFormField() bool`

HasFormField returns a boolean if a field has been set.

### GetGroupName

`func (o *ObjectTypePropertyCreate) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ObjectTypePropertyCreate) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ObjectTypePropertyCreate) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ObjectTypePropertyCreate) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetReferencedObjectType

`func (o *ObjectTypePropertyCreate) GetReferencedObjectType() string`

GetReferencedObjectType returns the ReferencedObjectType field if non-nil, zero value otherwise.

### GetReferencedObjectTypeOk

`func (o *ObjectTypePropertyCreate) GetReferencedObjectTypeOk() (*string, bool)`

GetReferencedObjectTypeOk returns a tuple with the ReferencedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedObjectType

`func (o *ObjectTypePropertyCreate) SetReferencedObjectType(v string)`

SetReferencedObjectType sets ReferencedObjectType field to given value.

### HasReferencedObjectType

`func (o *ObjectTypePropertyCreate) HasReferencedObjectType() bool`

HasReferencedObjectType returns a boolean if a field has been set.

### GetTextDisplayHint

`func (o *ObjectTypePropertyCreate) GetTextDisplayHint() string`

GetTextDisplayHint returns the TextDisplayHint field if non-nil, zero value otherwise.

### GetTextDisplayHintOk

`func (o *ObjectTypePropertyCreate) GetTextDisplayHintOk() (*string, bool)`

GetTextDisplayHintOk returns a tuple with the TextDisplayHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextDisplayHint

`func (o *ObjectTypePropertyCreate) SetTextDisplayHint(v string)`

SetTextDisplayHint sets TextDisplayHint field to given value.

### HasTextDisplayHint

`func (o *ObjectTypePropertyCreate) HasTextDisplayHint() bool`

HasTextDisplayHint returns a boolean if a field has been set.

### GetName

`func (o *ObjectTypePropertyCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectTypePropertyCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectTypePropertyCreate) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *ObjectTypePropertyCreate) GetOptions() []OptionInput`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ObjectTypePropertyCreate) GetOptionsOk() (*[]OptionInput, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ObjectTypePropertyCreate) SetOptions(v []OptionInput)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ObjectTypePropertyCreate) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetSearchableInGlobalSearch

`func (o *ObjectTypePropertyCreate) GetSearchableInGlobalSearch() bool`

GetSearchableInGlobalSearch returns the SearchableInGlobalSearch field if non-nil, zero value otherwise.

### GetSearchableInGlobalSearchOk

`func (o *ObjectTypePropertyCreate) GetSearchableInGlobalSearchOk() (*bool, bool)`

GetSearchableInGlobalSearchOk returns a tuple with the SearchableInGlobalSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchableInGlobalSearch

`func (o *ObjectTypePropertyCreate) SetSearchableInGlobalSearch(v bool)`

SetSearchableInGlobalSearch sets SearchableInGlobalSearch field to given value.

### HasSearchableInGlobalSearch

`func (o *ObjectTypePropertyCreate) HasSearchableInGlobalSearch() bool`

HasSearchableInGlobalSearch returns a boolean if a field has been set.

### GetNumberDisplayHint

`func (o *ObjectTypePropertyCreate) GetNumberDisplayHint() string`

GetNumberDisplayHint returns the NumberDisplayHint field if non-nil, zero value otherwise.

### GetNumberDisplayHintOk

`func (o *ObjectTypePropertyCreate) GetNumberDisplayHintOk() (*string, bool)`

GetNumberDisplayHintOk returns a tuple with the NumberDisplayHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberDisplayHint

`func (o *ObjectTypePropertyCreate) SetNumberDisplayHint(v string)`

SetNumberDisplayHint sets NumberDisplayHint field to given value.

### HasNumberDisplayHint

`func (o *ObjectTypePropertyCreate) HasNumberDisplayHint() bool`

HasNumberDisplayHint returns a boolean if a field has been set.

### GetHasUniqueValue

`func (o *ObjectTypePropertyCreate) GetHasUniqueValue() bool`

GetHasUniqueValue returns the HasUniqueValue field if non-nil, zero value otherwise.

### GetHasUniqueValueOk

`func (o *ObjectTypePropertyCreate) GetHasUniqueValueOk() (*bool, bool)`

GetHasUniqueValueOk returns a tuple with the HasUniqueValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUniqueValue

`func (o *ObjectTypePropertyCreate) SetHasUniqueValue(v bool)`

SetHasUniqueValue sets HasUniqueValue field to given value.

### HasHasUniqueValue

`func (o *ObjectTypePropertyCreate) HasHasUniqueValue() bool`

HasHasUniqueValue returns a boolean if a field has been set.

### GetFieldType

`func (o *ObjectTypePropertyCreate) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *ObjectTypePropertyCreate) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *ObjectTypePropertyCreate) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


