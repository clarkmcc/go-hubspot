# ObjectSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | [**ObjectTypeDefinitionLabels**](ObjectTypeDefinitionLabels.md) |  | 
**RequiredProperties** | **[]string** | The names of properties that should be **required** when creating an object of this type. | 
**SearchableProperties** | **[]string** | Names of properties that will be indexed for this object type in by HubSpot&#39;s product search. | 
**PrimaryDisplayProperty** | Pointer to **string** | The name of the primary property for this object. This will be displayed as primary on the HubSpot record page for this object type. | [optional] 
**SecondaryDisplayProperties** | **[]string** | The names of secondary properties for this object. These will be displayed as secondary on the HubSpot record page for this object type. | 
**Archived** | **bool** |  | 
**Id** | **string** | A unique ID for this schema&#39;s object type. Will be defined as {meta-type}-{unique ID}. | 
**FullyQualifiedName** | **string** | An assigned unique ID for the object, including portal ID and object name. | 
**CreatedAt** | Pointer to **time.Time** | When the object schema was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | When the object schema was last updated. | [optional] 
**ObjectTypeId** | **string** |  | 
**Properties** | [**[]Property**](Property.md) | Properties defined for this object type. | 
**Associations** | [**[]AssociationDefinition**](AssociationDefinition.md) | Associations defined for a given object type. | 
**Name** | **string** | A unique name for the schema&#39;s object type. | 

## Methods

### NewObjectSchema

`func NewObjectSchema(labels ObjectTypeDefinitionLabels, requiredProperties []string, searchableProperties []string, secondaryDisplayProperties []string, archived bool, id string, fullyQualifiedName string, objectTypeId string, properties []Property, associations []AssociationDefinition, name string, ) *ObjectSchema`

NewObjectSchema instantiates a new ObjectSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectSchemaWithDefaults

`func NewObjectSchemaWithDefaults() *ObjectSchema`

NewObjectSchemaWithDefaults instantiates a new ObjectSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *ObjectSchema) GetLabels() ObjectTypeDefinitionLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ObjectSchema) GetLabelsOk() (*ObjectTypeDefinitionLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ObjectSchema) SetLabels(v ObjectTypeDefinitionLabels)`

SetLabels sets Labels field to given value.


### GetRequiredProperties

`func (o *ObjectSchema) GetRequiredProperties() []string`

GetRequiredProperties returns the RequiredProperties field if non-nil, zero value otherwise.

### GetRequiredPropertiesOk

`func (o *ObjectSchema) GetRequiredPropertiesOk() (*[]string, bool)`

GetRequiredPropertiesOk returns a tuple with the RequiredProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredProperties

`func (o *ObjectSchema) SetRequiredProperties(v []string)`

SetRequiredProperties sets RequiredProperties field to given value.


### GetSearchableProperties

`func (o *ObjectSchema) GetSearchableProperties() []string`

GetSearchableProperties returns the SearchableProperties field if non-nil, zero value otherwise.

### GetSearchablePropertiesOk

`func (o *ObjectSchema) GetSearchablePropertiesOk() (*[]string, bool)`

GetSearchablePropertiesOk returns a tuple with the SearchableProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchableProperties

`func (o *ObjectSchema) SetSearchableProperties(v []string)`

SetSearchableProperties sets SearchableProperties field to given value.


### GetPrimaryDisplayProperty

`func (o *ObjectSchema) GetPrimaryDisplayProperty() string`

GetPrimaryDisplayProperty returns the PrimaryDisplayProperty field if non-nil, zero value otherwise.

### GetPrimaryDisplayPropertyOk

`func (o *ObjectSchema) GetPrimaryDisplayPropertyOk() (*string, bool)`

GetPrimaryDisplayPropertyOk returns a tuple with the PrimaryDisplayProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDisplayProperty

`func (o *ObjectSchema) SetPrimaryDisplayProperty(v string)`

SetPrimaryDisplayProperty sets PrimaryDisplayProperty field to given value.

### HasPrimaryDisplayProperty

`func (o *ObjectSchema) HasPrimaryDisplayProperty() bool`

HasPrimaryDisplayProperty returns a boolean if a field has been set.

### GetSecondaryDisplayProperties

`func (o *ObjectSchema) GetSecondaryDisplayProperties() []string`

GetSecondaryDisplayProperties returns the SecondaryDisplayProperties field if non-nil, zero value otherwise.

### GetSecondaryDisplayPropertiesOk

`func (o *ObjectSchema) GetSecondaryDisplayPropertiesOk() (*[]string, bool)`

GetSecondaryDisplayPropertiesOk returns a tuple with the SecondaryDisplayProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDisplayProperties

`func (o *ObjectSchema) SetSecondaryDisplayProperties(v []string)`

SetSecondaryDisplayProperties sets SecondaryDisplayProperties field to given value.


### GetArchived

`func (o *ObjectSchema) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ObjectSchema) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ObjectSchema) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetId

`func (o *ObjectSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectSchema) SetId(v string)`

SetId sets Id field to given value.


### GetFullyQualifiedName

`func (o *ObjectSchema) GetFullyQualifiedName() string`

GetFullyQualifiedName returns the FullyQualifiedName field if non-nil, zero value otherwise.

### GetFullyQualifiedNameOk

`func (o *ObjectSchema) GetFullyQualifiedNameOk() (*string, bool)`

GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedName

`func (o *ObjectSchema) SetFullyQualifiedName(v string)`

SetFullyQualifiedName sets FullyQualifiedName field to given value.


### GetCreatedAt

`func (o *ObjectSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ObjectSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ObjectSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ObjectSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ObjectSchema) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ObjectSchema) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ObjectSchema) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ObjectSchema) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetObjectTypeId

`func (o *ObjectSchema) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *ObjectSchema) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *ObjectSchema) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetProperties

`func (o *ObjectSchema) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ObjectSchema) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ObjectSchema) SetProperties(v []Property)`

SetProperties sets Properties field to given value.


### GetAssociations

`func (o *ObjectSchema) GetAssociations() []AssociationDefinition`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *ObjectSchema) GetAssociationsOk() (*[]AssociationDefinition, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *ObjectSchema) SetAssociations(v []AssociationDefinition)`

SetAssociations sets Associations field to given value.


### GetName

`func (o *ObjectSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectSchema) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


