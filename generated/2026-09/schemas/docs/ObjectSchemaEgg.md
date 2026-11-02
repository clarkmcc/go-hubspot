# ObjectSchemaEgg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowsSensitiveProperties** | **bool** | Determines if the object type can include properties that are marked as sensitive. | 
**AssociatedObjects** | **[]string** | Associations defined for this object type. | 
**Description** | Pointer to **string** | A brief explanation of the object type. | [optional] 
**Labels** | [**ObjectTypeDefinitionLabels**](ObjectTypeDefinitionLabels.md) |  | 
**Name** | **string** | A unique name for this object. For internal use only. | 
**PrimaryDisplayProperty** | Pointer to **string** | The name of the primary property for this object. This will be displayed as primary on the HubSpot record page for this object type. | [optional] 
**Properties** | [**[]ObjectTypePropertyCreate**](ObjectTypePropertyCreate.md) | Properties defined for this object type. | 
**RequiredProperties** | **[]string** | The names of properties that should be **required** when creating an object of this type. | 
**SearchableProperties** | **[]string** | Names of properties that will be indexed for this object type in by HubSpot&#39;s product search. | 
**SecondaryDisplayProperties** | **[]string** | The names of secondary properties for this object. These will be displayed as secondary on the HubSpot record page for this object type. | 
**ShouldCreateSameObjectAssociation** | **bool** | A boolean indicating whether associations should be created between objects of the same type. | 

## Methods

### NewObjectSchemaEgg

`func NewObjectSchemaEgg(allowsSensitiveProperties bool, associatedObjects []string, labels ObjectTypeDefinitionLabels, name string, properties []ObjectTypePropertyCreate, requiredProperties []string, searchableProperties []string, secondaryDisplayProperties []string, shouldCreateSameObjectAssociation bool, ) *ObjectSchemaEgg`

NewObjectSchemaEgg instantiates a new ObjectSchemaEgg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectSchemaEggWithDefaults

`func NewObjectSchemaEggWithDefaults() *ObjectSchemaEgg`

NewObjectSchemaEggWithDefaults instantiates a new ObjectSchemaEgg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowsSensitiveProperties

`func (o *ObjectSchemaEgg) GetAllowsSensitiveProperties() bool`

GetAllowsSensitiveProperties returns the AllowsSensitiveProperties field if non-nil, zero value otherwise.

### GetAllowsSensitivePropertiesOk

`func (o *ObjectSchemaEgg) GetAllowsSensitivePropertiesOk() (*bool, bool)`

GetAllowsSensitivePropertiesOk returns a tuple with the AllowsSensitiveProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowsSensitiveProperties

`func (o *ObjectSchemaEgg) SetAllowsSensitiveProperties(v bool)`

SetAllowsSensitiveProperties sets AllowsSensitiveProperties field to given value.


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


### GetShouldCreateSameObjectAssociation

`func (o *ObjectSchemaEgg) GetShouldCreateSameObjectAssociation() bool`

GetShouldCreateSameObjectAssociation returns the ShouldCreateSameObjectAssociation field if non-nil, zero value otherwise.

### GetShouldCreateSameObjectAssociationOk

`func (o *ObjectSchemaEgg) GetShouldCreateSameObjectAssociationOk() (*bool, bool)`

GetShouldCreateSameObjectAssociationOk returns a tuple with the ShouldCreateSameObjectAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCreateSameObjectAssociation

`func (o *ObjectSchemaEgg) SetShouldCreateSameObjectAssociation(v bool)`

SetShouldCreateSameObjectAssociation sets ShouldCreateSameObjectAssociation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


