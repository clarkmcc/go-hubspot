# PublicAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**PublicObjectId**](PublicObjectId.md) |  | 
**To** | [**PublicObjectId**](PublicObjectId.md) |  | 
**Type** | **string** |  | 

## Methods

### NewPublicAssociation

`func NewPublicAssociation(from PublicObjectId, to PublicObjectId, type_ string, ) *PublicAssociation`

NewPublicAssociation instantiates a new PublicAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAssociationWithDefaults

`func NewPublicAssociationWithDefaults() *PublicAssociation`

NewPublicAssociationWithDefaults instantiates a new PublicAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *PublicAssociation) GetFrom() PublicObjectId`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PublicAssociation) GetFromOk() (*PublicObjectId, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PublicAssociation) SetFrom(v PublicObjectId)`

SetFrom sets From field to given value.


### GetTo

`func (o *PublicAssociation) GetTo() PublicObjectId`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PublicAssociation) GetToOk() (*PublicObjectId, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PublicAssociation) SetTo(v PublicObjectId)`

SetTo sets To field to given value.


### GetType

`func (o *PublicAssociation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicAssociation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicAssociation) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


