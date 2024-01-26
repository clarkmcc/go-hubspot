# FieldTypeDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HelpText** | Pointer to **string** |  | [optional] 
**ReferencedObjectType** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Options** | [**[]Option**](Option.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**ExternalOptionsReferenceType** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**FieldType** | Pointer to **string** |  | [optional] 
**OptionsUrl** | Pointer to **string** |  | [optional] 
**ExternalOptions** | **bool** |  | 

## Methods

### NewFieldTypeDefinition

`func NewFieldTypeDefinition(name string, options []Option, type_ string, externalOptions bool, ) *FieldTypeDefinition`

NewFieldTypeDefinition instantiates a new FieldTypeDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldTypeDefinitionWithDefaults

`func NewFieldTypeDefinitionWithDefaults() *FieldTypeDefinition`

NewFieldTypeDefinitionWithDefaults instantiates a new FieldTypeDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHelpText

`func (o *FieldTypeDefinition) GetHelpText() string`

GetHelpText returns the HelpText field if non-nil, zero value otherwise.

### GetHelpTextOk

`func (o *FieldTypeDefinition) GetHelpTextOk() (*string, bool)`

GetHelpTextOk returns a tuple with the HelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpText

`func (o *FieldTypeDefinition) SetHelpText(v string)`

SetHelpText sets HelpText field to given value.

### HasHelpText

`func (o *FieldTypeDefinition) HasHelpText() bool`

HasHelpText returns a boolean if a field has been set.

### GetReferencedObjectType

`func (o *FieldTypeDefinition) GetReferencedObjectType() string`

GetReferencedObjectType returns the ReferencedObjectType field if non-nil, zero value otherwise.

### GetReferencedObjectTypeOk

`func (o *FieldTypeDefinition) GetReferencedObjectTypeOk() (*string, bool)`

GetReferencedObjectTypeOk returns a tuple with the ReferencedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedObjectType

`func (o *FieldTypeDefinition) SetReferencedObjectType(v string)`

SetReferencedObjectType sets ReferencedObjectType field to given value.

### HasReferencedObjectType

`func (o *FieldTypeDefinition) HasReferencedObjectType() bool`

HasReferencedObjectType returns a boolean if a field has been set.

### GetName

`func (o *FieldTypeDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldTypeDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldTypeDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *FieldTypeDefinition) GetOptions() []Option`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FieldTypeDefinition) GetOptionsOk() (*[]Option, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FieldTypeDefinition) SetOptions(v []Option)`

SetOptions sets Options field to given value.


### GetDescription

`func (o *FieldTypeDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FieldTypeDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FieldTypeDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FieldTypeDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalOptionsReferenceType

`func (o *FieldTypeDefinition) GetExternalOptionsReferenceType() string`

GetExternalOptionsReferenceType returns the ExternalOptionsReferenceType field if non-nil, zero value otherwise.

### GetExternalOptionsReferenceTypeOk

`func (o *FieldTypeDefinition) GetExternalOptionsReferenceTypeOk() (*string, bool)`

GetExternalOptionsReferenceTypeOk returns a tuple with the ExternalOptionsReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOptionsReferenceType

`func (o *FieldTypeDefinition) SetExternalOptionsReferenceType(v string)`

SetExternalOptionsReferenceType sets ExternalOptionsReferenceType field to given value.

### HasExternalOptionsReferenceType

`func (o *FieldTypeDefinition) HasExternalOptionsReferenceType() bool`

HasExternalOptionsReferenceType returns a boolean if a field has been set.

### GetLabel

`func (o *FieldTypeDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FieldTypeDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FieldTypeDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FieldTypeDefinition) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *FieldTypeDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FieldTypeDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FieldTypeDefinition) SetType(v string)`

SetType sets Type field to given value.


### GetFieldType

`func (o *FieldTypeDefinition) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *FieldTypeDefinition) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *FieldTypeDefinition) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *FieldTypeDefinition) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### GetOptionsUrl

`func (o *FieldTypeDefinition) GetOptionsUrl() string`

GetOptionsUrl returns the OptionsUrl field if non-nil, zero value otherwise.

### GetOptionsUrlOk

`func (o *FieldTypeDefinition) GetOptionsUrlOk() (*string, bool)`

GetOptionsUrlOk returns a tuple with the OptionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsUrl

`func (o *FieldTypeDefinition) SetOptionsUrl(v string)`

SetOptionsUrl sets OptionsUrl field to given value.

### HasOptionsUrl

`func (o *FieldTypeDefinition) HasOptionsUrl() bool`

HasOptionsUrl returns a boolean if a field has been set.

### GetExternalOptions

`func (o *FieldTypeDefinition) GetExternalOptions() bool`

GetExternalOptions returns the ExternalOptions field if non-nil, zero value otherwise.

### GetExternalOptionsOk

`func (o *FieldTypeDefinition) GetExternalOptionsOk() (*bool, bool)`

GetExternalOptionsOk returns a tuple with the ExternalOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOptions

`func (o *FieldTypeDefinition) SetExternalOptions(v bool)`

SetExternalOptions sets ExternalOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


