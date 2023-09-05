# PublicAssociationMultiPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**PublicObjectId**](PublicObjectId.md) |  | 
**To** | [**PublicObjectId**](PublicObjectId.md) |  | 
**Types** | [**[]AssociationSpec**](AssociationSpec.md) |  | 

## Methods

### NewPublicAssociationMultiPost

`func NewPublicAssociationMultiPost(from PublicObjectId, to PublicObjectId, types []AssociationSpec, ) *PublicAssociationMultiPost`

NewPublicAssociationMultiPost instantiates a new PublicAssociationMultiPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAssociationMultiPostWithDefaults

`func NewPublicAssociationMultiPostWithDefaults() *PublicAssociationMultiPost`

NewPublicAssociationMultiPostWithDefaults instantiates a new PublicAssociationMultiPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *PublicAssociationMultiPost) GetFrom() PublicObjectId`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PublicAssociationMultiPost) GetFromOk() (*PublicObjectId, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PublicAssociationMultiPost) SetFrom(v PublicObjectId)`

SetFrom sets From field to given value.


### GetTo

`func (o *PublicAssociationMultiPost) GetTo() PublicObjectId`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PublicAssociationMultiPost) GetToOk() (*PublicObjectId, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PublicAssociationMultiPost) SetTo(v PublicObjectId)`

SetTo sets To field to given value.


### GetTypes

`func (o *PublicAssociationMultiPost) GetTypes() []AssociationSpec`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *PublicAssociationMultiPost) GetTypesOk() (*[]AssociationSpec, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *PublicAssociationMultiPost) SetTypes(v []AssociationSpec)`

SetTypes sets Types field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


