# ObjectSchemaEgg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecondaryDisplayProperties** | Pointer to **[]string** | The names of secondary properties for this object. These will be displayed as secondary on the HubSpot record page for this object type. | [optional] 
**RequiredProperties** | **[]string** | The names of properties that should be **required** when creating an object of this type. | 
**SearchableProperties** | Pointer to **[]string** | Names of properties that will be indexed for this object type in by HubSpot&#39;s product search. | [optional] 
**PrimaryDisplayProperty** | Pointer to **string** | The name of the primary property for this object. This will be displayed as primary on the HubSpot record page for this object type. | [optional] 
**Name** | **string** | A unique name for this object. For internal use only. | 
**Description** | Pointer to **string** |  | [optional] 
**AssociatedObjects** | **[]string** | Associations defined for this object type. | 
**Properties** | [**[]ObjectTypePropertyCreate**](ObjectTypePropertyCreate.md) | Properties defined for this object type. | 
**Labels** | [**ObjectTypeDefinitionLabels**](ObjectTypeDefinitionLabels.md) |  | 

## Methods

### NewObjectSchemaEgg

`func NewObjectSchemaEgg(requiredProperties []string, name string, associatedObjects []string, properties []ObjectTypePropertyCreate, labels ObjectTypeDefinitionLabels, ) *ObjectSchemaEgg`

NewObjectSchemaEgg instantiates a new ObjectSchemaEgg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectSchemaEggWithDefaults

`func NewObjectSchemaEggWithDefaults() *ObjectSchemaEgg`

NewObjectSchemaEggWithDefaults instantiates a new ObjectSchemaEgg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecondaryDisplayProperties

`func (o *ObjectSchemaEgg) GetSecondaryDisplayProperties() []string`

GetSecondaryDisplayProperties returns the SecondaryDisplayProperties field if non-nil, zero value otherwise.

### GetSecondaryDisplayPropertiesOk

`func (o *ObjectSchemaEgg) GetSecondaryDisplayPropertiesOk() (*[]string, bool)`

GetSecondaryDisplayPropertiesOk returns a tuple with the SecondaryDisplayProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDisplayProperties

`func (o *ObjectSchemaEgg) SetSecondaryDisplayProperties(v []string)`

SetSecondaryDisplayProperties sets SecondaryDisplayProperties field to given value.

### HasSecondaryDisplayProperties

`func (o *ObjectSchemaEgg) HasSecondaryDisplayProperties() bool`

HasSecondaryDisplayProperties returns a boolean if a field has been set.

### GetRequiredProperties

`func (o *ObjectSchemaEgg) GetRequiredProperties() []string`

GetRequiredProperties returns the RequiredProperties field if non-nil, zero value otherwise.

### GetRequiredPropertiesOk

`func (o *ObjectSchemaEgg) GetRequiredPropertiesOk() (*[]string, bool)`

GetRequiredPropertiesOk returns a tuple with the RequiredProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredProperties

`func (o *ObjectSchemaEgg) SetRequiredProperties(v []string)`

SetRequiredProperties sets RequiredProperties field to given value.


### GetSearchableProperties

`func (o *ObjectSchemaEgg) GetSearchableProperties() []string`

GetSearchableProperties returns the SearchableProperties field if non-nil, zero value otherwise.

### GetSearchablePropertiesOk

`func (o *ObjectSchemaEgg) GetSearchablePropertiesOk() (*[]string, bool)`

GetSearchablePropertiesOk returns a tuple with the SearchableProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchableProperties

`func (o *ObjectSchemaEgg) SetSearchableProperties(v []string)`

SetSearchableProperties sets SearchableProperties field to given value.

### HasSearchableProperties

`func (o *ObjectSchemaEgg) HasSearchableProperties() bool`

HasSearchableProperties returns a boolean if a field has been set.

### GetPrimaryDisplayProperty

`func (o *ObjectSchemaEgg) GetPrimaryDisplayProperty() string`

GetPrimaryDisplayProperty returns the PrimaryDisplayProperty field if non-nil, zero value otherwise.

### GetPrimaryDisplayPropertyOk

`func (o *ObjectSchemaEgg) GetPrimaryDisplayPropertyOk() (*string, bool)`

GetPrimaryDisplayPropertyOk returns a tuple with the PrimaryDisplayProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDisplayProperty

`func (o *ObjectSchemaEgg) SetPrimaryDisplayProperty(v string)`

SetPrimaryDisplayProperty sets PrimaryDisplayProperty field to given value.

### HasPrimaryDisplayProperty

`func (o *ObjectSchemaEgg) HasPrimaryDisplayProperty() bool`

HasPrimaryDisplayProperty returns a boolean if a field has been set.

### GetName

`func (o *ObjectSchemaEgg) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectSchemaEgg) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectSchemaEgg) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ObjectSchemaEgg) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectSchemaEgg) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectSchemaEgg) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectSchemaEgg) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAssociatedObjects

`func (o *ObjectSchemaEgg) GetAssociatedObjects() []string`

GetAssociatedObjects returns the AssociatedObjects field if non-nil, zero value otherwise.

### GetAssociatedObjectsOk

`func (o *ObjectSchemaEgg) GetAssociatedObjectsOk() (*[]string, bool)`

GetAssociatedObjectsOk returns a tuple with the AssociatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObjects

`func (o *ObjectSchemaEgg) SetAssociatedObjects(v []string)`

SetAssociatedObjects sets AssociatedObjects field to given value.


### GetProperties

`func (o *ObjectSchemaEgg) GetProperties() []ObjectTypePropertyCreate`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ObjectSchemaEgg) GetPropertiesOk() (*[]ObjectTypePropertyCreate, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ObjectSchemaEgg) SetProperties(v []ObjectTypePropertyCreate)`

SetProperties sets Properties field to given value.


### GetLabels

`func (o *ObjectSchemaEgg) GetLabels() ObjectTypeDefinitionLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ObjectSchemaEgg) GetLabelsOk() (*ObjectTypeDefinitionLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ObjectSchemaEgg) SetLabels(v ObjectTypeDefinitionLabels)`

SetLabels sets Labels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


