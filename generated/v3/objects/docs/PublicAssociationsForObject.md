# PublicAssociationsForObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Types** | [**[]AssociationSpec**](AssociationSpec.md) |  | 
**To** | [**PublicObjectId**](PublicObjectId.md) |  | 

## Methods

### NewPublicAssociationsForObject

`func NewPublicAssociationsForObject(types []AssociationSpec, to PublicObjectId, ) *PublicAssociationsForObject`

NewPublicAssociationsForObject instantiates a new PublicAssociationsForObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAssociationsForObjectWithDefaults

`func NewPublicAssociationsForObjectWithDefaults() *PublicAssociationsForObject`

NewPublicAssociationsForObjectWithDefaults instantiates a new PublicAssociationsForObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypes

`func (o *PublicAssociationsForObject) GetTypes() []AssociationSpec`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *PublicAssociationsForObject) GetTypesOk() (*[]AssociationSpec, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *PublicAssociationsForObject) SetTypes(v []AssociationSpec)`

SetTypes sets Types field to given value.


### GetTo

`func (o *PublicAssociationsForObject) GetTo() PublicObjectId`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PublicAssociationsForObject) GetToOk() (*PublicObjectId, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PublicAssociationsForObject) SetTo(v PublicObjectId)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


