# InputFieldDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeDefinition** | [**FieldTypeDefinition**](FieldTypeDefinition.md) |  | 
**SupportedValueTypes** | Pointer to **[]string** | Controls what kind of input a customer can use to specify the field value. Must contain exactly one of &#x60;STATIC_VALUE&#x60; or &#x60;OBJECT_PROPERTY&#x60;. If &#x60;STATIC_VALUE&#x60;, the customer will be able to choose a value when configuring the custom action; if &#x60;OBJECT_PROPERTY&#x60;, the customer will be able to choose a property from the enrolled workflow object that the field value will be copied from. In the future we may support more than one input control type here. | [optional] 
**IsRequired** | **bool** | Whether the field is required for the custom action to be valid | 

## Methods

### NewInputFieldDefinition

`func NewInputFieldDefinition(typeDefinition FieldTypeDefinition, isRequired bool, ) *InputFieldDefinition`

NewInputFieldDefinition instantiates a new InputFieldDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputFieldDefinitionWithDefaults

`func NewInputFieldDefinitionWithDefaults() *InputFieldDefinition`

NewInputFieldDefinitionWithDefaults instantiates a new InputFieldDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


