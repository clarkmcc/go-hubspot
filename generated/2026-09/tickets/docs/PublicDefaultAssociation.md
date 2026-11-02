# PublicDefaultAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationSpec** | [**AssociationSpec**](AssociationSpec.md) |  | 
**From** | [**PublicObjectId**](PublicObjectId.md) |  | 
**To** | [**PublicObjectId**](PublicObjectId.md) |  | 

## Methods

### NewPublicDefaultAssociation

`func NewPublicDefaultAssociation(associationSpec AssociationSpec, from PublicObjectId, to PublicObjectId, ) *PublicDefaultAssociation`

NewPublicDefaultAssociation instantiates a new PublicDefaultAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicDefaultAssociationWithDefaults

`func NewPublicDefaultAssociationWithDefaults() *PublicDefaultAssociation`

NewPublicDefaultAssociationWithDefaults instantiates a new PublicDefaultAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationSpec

`func (o *PublicDefaultAssociation) GetAssociationSpec() AssociationSpec`

GetAssociationSpec returns the AssociationSpec field if non-nil, zero value otherwise.

### GetAssociationSpecOk

`func (o *PublicDefaultAssociation) GetAssociationSpecOk() (*AssociationSpec, bool)`

GetAssociationSpecOk returns a tuple with the AssociationSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationSpec

`func (o *PublicDefaultAssociation) SetAssociationSpec(v AssociationSpec)`

SetAssociationSpec sets AssociationSpec field to given value.


### GetFrom

`func (o *PublicDefaultAssociation) GetFrom() PublicObjectId`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PublicDefaultAssociation) GetFromOk() (*PublicObjectId, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PublicDefaultAssociation) SetFrom(v PublicObjectId)`

SetFrom sets From field to given value.


### GetTo

`func (o *PublicDefaultAssociation) GetTo() PublicObjectId`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PublicDefaultAssociation) GetToOk() (*PublicObjectId, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PublicDefaultAssociation) SetTo(v PublicObjectId)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


