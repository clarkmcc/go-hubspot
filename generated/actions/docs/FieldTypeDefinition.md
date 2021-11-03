# FieldTypeDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The input field name. | 
**Type** | **string** | The data type of the field. | 
**FieldType** | Pointer to **string** | Controls how the field appears in HubSpot. | [optional] 
**Options** | [**[]Option**](Option.md) | A list of valid options for the field value. | 
**OptionsUrl** | Pointer to **string** | A URL that will accept HTTPS requests when the valid options for the field are fetched. | [optional] 
**ReferencedObjectType** | Pointer to **string** | This can be set to &#x60;OWNER&#x60; if the field should contain a HubSpot owner value. | [optional] 

## Methods

### NewFieldTypeDefinition

`func NewFieldTypeDefinition(name string, type_ string, options []Option, ) *FieldTypeDefinition`

NewFieldTypeDefinition instantiates a new FieldTypeDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldTypeDefinitionWithDefaults

`func NewFieldTypeDefinitionWithDefaults() *FieldTypeDefinition`

NewFieldTypeDefinitionWithDefaults instantiates a new FieldTypeDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


