# PublicAssociationMultiWithLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**PublicObjectId**](PublicObjectId.md) |  | 
**To** | [**[]MultiAssociatedObjectWithLabel**](MultiAssociatedObjectWithLabel.md) |  | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 

## Methods

### NewPublicAssociationMultiWithLabel

`func NewPublicAssociationMultiWithLabel(from PublicObjectId, to []MultiAssociatedObjectWithLabel, ) *PublicAssociationMultiWithLabel`

NewPublicAssociationMultiWithLabel instantiates a new PublicAssociationMultiWithLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAssociationMultiWithLabelWithDefaults

`func NewPublicAssociationMultiWithLabelWithDefaults() *PublicAssociationMultiWithLabel`

NewPublicAssociationMultiWithLabelWithDefaults instantiates a new PublicAssociationMultiWithLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *PublicAssociationMultiWithLabel) GetFrom() PublicObjectId`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PublicAssociationMultiWithLabel) GetFromOk() (*PublicObjectId, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PublicAssociationMultiWithLabel) SetFrom(v PublicObjectId)`

SetFrom sets From field to given value.


### GetTo

`func (o *PublicAssociationMultiWithLabel) GetTo() []MultiAssociatedObjectWithLabel`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PublicAssociationMultiWithLabel) GetToOk() (*[]MultiAssociatedObjectWithLabel, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PublicAssociationMultiWithLabel) SetTo(v []MultiAssociatedObjectWithLabel)`

SetTo sets To field to given value.


### GetPaging

`func (o *PublicAssociationMultiWithLabel) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *PublicAssociationMultiWithLabel) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *PublicAssociationMultiWithLabel) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *PublicAssociationMultiWithLabel) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


