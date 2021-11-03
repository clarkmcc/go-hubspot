# AssociationDefinitionEgg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromObjectTypeId** | **string** | ID of the primary object type to link from. | 
**ToObjectTypeId** | **string** | ID of the target object type ID to link to. | 
**Name** | Pointer to **string** | A unique name for this association. | [optional] 

## Methods

### NewAssociationDefinitionEgg

`func NewAssociationDefinitionEgg(fromObjectTypeId string, toObjectTypeId string, ) *AssociationDefinitionEgg`

NewAssociationDefinitionEgg instantiates a new AssociationDefinitionEgg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociationDefinitionEggWithDefaults

`func NewAssociationDefinitionEggWithDefaults() *AssociationDefinitionEgg`

NewAssociationDefinitionEggWithDefaults instantiates a new AssociationDefinitionEgg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromObjectTypeId

`func (o *AssociationDefinitionEgg) GetFromObjectTypeId() string`

GetFromObjectTypeId returns the FromObjectTypeId field if non-nil, zero value otherwise.

### GetFromObjectTypeIdOk

`func (o *AssociationDefinitionEgg) GetFromObjectTypeIdOk() (*string, bool)`

GetFromObjectTypeIdOk returns a tuple with the FromObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromObjectTypeId

`func (o *AssociationDefinitionEgg) SetFromObjectTypeId(v string)`

SetFromObjectTypeId sets FromObjectTypeId field to given value.


### GetToObjectTypeId

`func (o *AssociationDefinitionEgg) GetToObjectTypeId() string`

GetToObjectTypeId returns the ToObjectTypeId field if non-nil, zero value otherwise.

### GetToObjectTypeIdOk

`func (o *AssociationDefinitionEgg) GetToObjectTypeIdOk() (*string, bool)`

GetToObjectTypeIdOk returns a tuple with the ToObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToObjectTypeId

`func (o *AssociationDefinitionEgg) SetToObjectTypeId(v string)`

SetToObjectTypeId sets ToObjectTypeId field to given value.


### GetName

`func (o *AssociationDefinitionEgg) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssociationDefinitionEgg) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssociationDefinitionEgg) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssociationDefinitionEgg) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


