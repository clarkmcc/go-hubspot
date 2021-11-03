# PublicAssociationMulti

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**PublicObjectId**](PublicObjectId.md) |  | 
**To** | [**[]AssociatedId**](AssociatedId.md) | The IDs of objects that are associated with the object identified by the ID in &#39;from&#39;. | 

## Methods

### NewPublicAssociationMulti

`func NewPublicAssociationMulti(from PublicObjectId, to []AssociatedId, ) *PublicAssociationMulti`

NewPublicAssociationMulti instantiates a new PublicAssociationMulti object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAssociationMultiWithDefaults

`func NewPublicAssociationMultiWithDefaults() *PublicAssociationMulti`

NewPublicAssociationMultiWithDefaults instantiates a new PublicAssociationMulti object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *PublicAssociationMulti) GetFrom() PublicObjectId`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PublicAssociationMulti) GetFromOk() (*PublicObjectId, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PublicAssociationMulti) SetFrom(v PublicObjectId)`

SetFrom sets From field to given value.


### GetTo

`func (o *PublicAssociationMulti) GetTo() []AssociatedId`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PublicAssociationMulti) GetToOk() (*[]AssociatedId, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PublicAssociationMulti) SetTo(v []AssociatedId)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


