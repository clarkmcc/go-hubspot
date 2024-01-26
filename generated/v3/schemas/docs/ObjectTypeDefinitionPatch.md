# ObjectTypeDefinitionPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecondaryDisplayProperties** | Pointer to **[]string** | The names of secondary properties for this object. These will be displayed as secondary on the HubSpot record page for this object type. | [optional] 
**RequiredProperties** | Pointer to **[]string** | The names of properties that should be **required** when creating an object of this type. | [optional] 
**SearchableProperties** | Pointer to **[]string** | Names of properties that will be indexed for this object type in by HubSpot&#39;s product search. | [optional] 
**PrimaryDisplayProperty** | Pointer to **string** | The name of the primary property for this object. This will be displayed as primary on the HubSpot record page for this object type. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Restorable** | Pointer to **bool** |  | [optional] 
**Labels** | Pointer to [**ObjectTypeDefinitionLabels**](ObjectTypeDefinitionLabels.md) |  | [optional] 

## Methods

### NewObjectTypeDefinitionPatch

`func NewObjectTypeDefinitionPatch() *ObjectTypeDefinitionPatch`

NewObjectTypeDefinitionPatch instantiates a new ObjectTypeDefinitionPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectTypeDefinitionPatchWithDefaults

`func NewObjectTypeDefinitionPatchWithDefaults() *ObjectTypeDefinitionPatch`

NewObjectTypeDefinitionPatchWithDefaults instantiates a new ObjectTypeDefinitionPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecondaryDisplayProperties

`func (o *ObjectTypeDefinitionPatch) GetSecondaryDisplayProperties() []string`

GetSecondaryDisplayProperties returns the SecondaryDisplayProperties field if non-nil, zero value otherwise.

### GetSecondaryDisplayPropertiesOk

`func (o *ObjectTypeDefinitionPatch) GetSecondaryDisplayPropertiesOk() (*[]string, bool)`

GetSecondaryDisplayPropertiesOk returns a tuple with the SecondaryDisplayProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDisplayProperties

`func (o *ObjectTypeDefinitionPatch) SetSecondaryDisplayProperties(v []string)`

SetSecondaryDisplayProperties sets SecondaryDisplayProperties field to given value.

### HasSecondaryDisplayProperties

`func (o *ObjectTypeDefinitionPatch) HasSecondaryDisplayProperties() bool`

HasSecondaryDisplayProperties returns a boolean if a field has been set.

### GetRequiredProperties

`func (o *ObjectTypeDefinitionPatch) GetRequiredProperties() []string`

GetRequiredProperties returns the RequiredProperties field if non-nil, zero value otherwise.

### GetRequiredPropertiesOk

`func (o *ObjectTypeDefinitionPatch) GetRequiredPropertiesOk() (*[]string, bool)`

GetRequiredPropertiesOk returns a tuple with the RequiredProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredProperties

`func (o *ObjectTypeDefinitionPatch) SetRequiredProperties(v []string)`

SetRequiredProperties sets RequiredProperties field to given value.

### HasRequiredProperties

`func (o *ObjectTypeDefinitionPatch) HasRequiredProperties() bool`

HasRequiredProperties returns a boolean if a field has been set.

### GetSearchableProperties

`func (o *ObjectTypeDefinitionPatch) GetSearchableProperties() []string`

GetSearchableProperties returns the SearchableProperties field if non-nil, zero value otherwise.

### GetSearchablePropertiesOk

`func (o *ObjectTypeDefinitionPatch) GetSearchablePropertiesOk() (*[]string, bool)`

GetSearchablePropertiesOk returns a tuple with the SearchableProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchableProperties

`func (o *ObjectTypeDefinitionPatch) SetSearchableProperties(v []string)`

SetSearchableProperties sets SearchableProperties field to given value.

### HasSearchableProperties

`func (o *ObjectTypeDefinitionPatch) HasSearchableProperties() bool`

HasSearchableProperties returns a boolean if a field has been set.

### GetPrimaryDisplayProperty

`func (o *ObjectTypeDefinitionPatch) GetPrimaryDisplayProperty() string`

GetPrimaryDisplayProperty returns the PrimaryDisplayProperty field if non-nil, zero value otherwise.

### GetPrimaryDisplayPropertyOk

`func (o *ObjectTypeDefinitionPatch) GetPrimaryDisplayPropertyOk() (*string, bool)`

GetPrimaryDisplayPropertyOk returns a tuple with the PrimaryDisplayProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDisplayProperty

`func (o *ObjectTypeDefinitionPatch) SetPrimaryDisplayProperty(v string)`

SetPrimaryDisplayProperty sets PrimaryDisplayProperty field to given value.

### HasPrimaryDisplayProperty

`func (o *ObjectTypeDefinitionPatch) HasPrimaryDisplayProperty() bool`

HasPrimaryDisplayProperty returns a boolean if a field has been set.

### GetDescription

`func (o *ObjectTypeDefinitionPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectTypeDefinitionPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectTypeDefinitionPatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectTypeDefinitionPatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRestorable

`func (o *ObjectTypeDefinitionPatch) GetRestorable() bool`

GetRestorable returns the Restorable field if non-nil, zero value otherwise.

### GetRestorableOk

`func (o *ObjectTypeDefinitionPatch) GetRestorableOk() (*bool, bool)`

GetRestorableOk returns a tuple with the Restorable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestorable

`func (o *ObjectTypeDefinitionPatch) SetRestorable(v bool)`

SetRestorable sets Restorable field to given value.

### HasRestorable

`func (o *ObjectTypeDefinitionPatch) HasRestorable() bool`

HasRestorable returns a boolean if a field has been set.

### GetLabels

`func (o *ObjectTypeDefinitionPatch) GetLabels() ObjectTypeDefinitionLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ObjectTypeDefinitionPatch) GetLabelsOk() (*ObjectTypeDefinitionLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ObjectTypeDefinitionPatch) SetLabels(v ObjectTypeDefinitionLabels)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ObjectTypeDefinitionPatch) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


