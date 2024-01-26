# SimplePublicObjectInputForCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Associations** | Pointer to [**[]PublicAssociationsForObject**](PublicAssociationsForObject.md) |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSimplePublicObjectInputForCreate

`func NewSimplePublicObjectInputForCreate() *SimplePublicObjectInputForCreate`

NewSimplePublicObjectInputForCreate instantiates a new SimplePublicObjectInputForCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplePublicObjectInputForCreateWithDefaults

`func NewSimplePublicObjectInputForCreateWithDefaults() *SimplePublicObjectInputForCreate`

NewSimplePublicObjectInputForCreateWithDefaults instantiates a new SimplePublicObjectInputForCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociations

`func (o *SimplePublicObjectInputForCreate) GetAssociations() []PublicAssociationsForObject`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *SimplePublicObjectInputForCreate) GetAssociationsOk() (*[]PublicAssociationsForObject, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *SimplePublicObjectInputForCreate) SetAssociations(v []PublicAssociationsForObject)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *SimplePublicObjectInputForCreate) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetProperties

`func (o *SimplePublicObjectInputForCreate) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SimplePublicObjectInputForCreate) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SimplePublicObjectInputForCreate) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SimplePublicObjectInputForCreate) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


