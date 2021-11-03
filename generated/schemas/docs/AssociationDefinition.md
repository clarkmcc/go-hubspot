# AssociationDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromObjectTypeId** | **string** | ID of the primary object type to link from. | 
**ToObjectTypeId** | **string** | ID of the target object type ID to link to. | 
**Name** | Pointer to **string** | A unique name for this association. | [optional] 
**Id** | **string** | A unique ID for this association. | 
**CreatedAt** | Pointer to **time.Time** | When the association was defined. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | When the association was last updated. | [optional] 

## Methods

### NewAssociationDefinition

`func NewAssociationDefinition(fromObjectTypeId string, toObjectTypeId string, id string, ) *AssociationDefinition`

NewAssociationDefinition instantiates a new AssociationDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociationDefinitionWithDefaults

`func NewAssociationDefinitionWithDefaults() *AssociationDefinition`

NewAssociationDefinitionWithDefaults instantiates a new AssociationDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromObjectTypeId

`func (o *AssociationDefinition) GetFromObjectTypeId() string`

GetFromObjectTypeId returns the FromObjectTypeId field if non-nil, zero value otherwise.

### GetFromObjectTypeIdOk

`func (o *AssociationDefinition) GetFromObjectTypeIdOk() (*string, bool)`

GetFromObjectTypeIdOk returns a tuple with the FromObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromObjectTypeId

`func (o *AssociationDefinition) SetFromObjectTypeId(v string)`

SetFromObjectTypeId sets FromObjectTypeId field to given value.


### GetToObjectTypeId

`func (o *AssociationDefinition) GetToObjectTypeId() string`

GetToObjectTypeId returns the ToObjectTypeId field if non-nil, zero value otherwise.

### GetToObjectTypeIdOk

`func (o *AssociationDefinition) GetToObjectTypeIdOk() (*string, bool)`

GetToObjectTypeIdOk returns a tuple with the ToObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToObjectTypeId

`func (o *AssociationDefinition) SetToObjectTypeId(v string)`

SetToObjectTypeId sets ToObjectTypeId field to given value.


### GetName

`func (o *AssociationDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssociationDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssociationDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssociationDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *AssociationDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssociationDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssociationDefinition) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *AssociationDefinition) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AssociationDefinition) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AssociationDefinition) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AssociationDefinition) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AssociationDefinition) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AssociationDefinition) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AssociationDefinition) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AssociationDefinition) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


