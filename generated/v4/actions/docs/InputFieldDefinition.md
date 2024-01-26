# InputFieldDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsRequired** | **bool** |  | 
**AutomationFieldType** | Pointer to **string** |  | [optional] 
**TypeDefinition** | [**FieldTypeDefinition**](FieldTypeDefinition.md) |  | 
**SupportedValueTypes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewInputFieldDefinition

`func NewInputFieldDefinition(isRequired bool, typeDefinition FieldTypeDefinition, ) *InputFieldDefinition`

NewInputFieldDefinition instantiates a new InputFieldDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputFieldDefinitionWithDefaults

`func NewInputFieldDefinitionWithDefaults() *InputFieldDefinition`

NewInputFieldDefinitionWithDefaults instantiates a new InputFieldDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsRequired

`func (o *InputFieldDefinition) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *InputFieldDefinition) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *InputFieldDefinition) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.


### GetAutomationFieldType

`func (o *InputFieldDefinition) GetAutomationFieldType() string`

GetAutomationFieldType returns the AutomationFieldType field if non-nil, zero value otherwise.

### GetAutomationFieldTypeOk

`func (o *InputFieldDefinition) GetAutomationFieldTypeOk() (*string, bool)`

GetAutomationFieldTypeOk returns a tuple with the AutomationFieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationFieldType

`func (o *InputFieldDefinition) SetAutomationFieldType(v string)`

SetAutomationFieldType sets AutomationFieldType field to given value.

### HasAutomationFieldType

`func (o *InputFieldDefinition) HasAutomationFieldType() bool`

HasAutomationFieldType returns a boolean if a field has been set.

### GetTypeDefinition

`func (o *InputFieldDefinition) GetTypeDefinition() FieldTypeDefinition`

GetTypeDefinition returns the TypeDefinition field if non-nil, zero value otherwise.

### GetTypeDefinitionOk

`func (o *InputFieldDefinition) GetTypeDefinitionOk() (*FieldTypeDefinition, bool)`

GetTypeDefinitionOk returns a tuple with the TypeDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDefinition

`func (o *InputFieldDefinition) SetTypeDefinition(v FieldTypeDefinition)`

SetTypeDefinition sets TypeDefinition field to given value.


### GetSupportedValueTypes

`func (o *InputFieldDefinition) GetSupportedValueTypes() []string`

GetSupportedValueTypes returns the SupportedValueTypes field if non-nil, zero value otherwise.

### GetSupportedValueTypesOk

`func (o *InputFieldDefinition) GetSupportedValueTypesOk() (*[]string, bool)`

GetSupportedValueTypesOk returns a tuple with the SupportedValueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedValueTypes

`func (o *InputFieldDefinition) SetSupportedValueTypes(v []string)`

SetSupportedValueTypes sets SupportedValueTypes field to given value.

### HasSupportedValueTypes

`func (o *InputFieldDefinition) HasSupportedValueTypes() bool`

HasSupportedValueTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


