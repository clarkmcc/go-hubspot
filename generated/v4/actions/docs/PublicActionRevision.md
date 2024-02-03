# PublicActionRevision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Definition** | [**PublicActionDefinition**](PublicActionDefinition.md) |  | 
**Id** | **string** |  | 

## Methods

### NewPublicActionRevision

`func NewPublicActionRevision(revisionId string, createdAt time.Time, definition PublicActionDefinition, id string, ) *PublicActionRevision`

NewPublicActionRevision instantiates a new PublicActionRevision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicActionRevisionWithDefaults

`func NewPublicActionRevisionWithDefaults() *PublicActionRevision`

NewPublicActionRevisionWithDefaults instantiates a new PublicActionRevision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionId

`func (o *PublicActionRevision) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *PublicActionRevision) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *PublicActionRevision) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.


### GetCreatedAt

`func (o *PublicActionRevision) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicActionRevision) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicActionRevision) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDefinition

`func (o *PublicActionRevision) GetDefinition() PublicActionDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *PublicActionRevision) GetDefinitionOk() (*PublicActionDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *PublicActionRevision) SetDefinition(v PublicActionDefinition)`

SetDefinition sets Definition field to given value.


### GetId

`func (o *PublicActionRevision) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicActionRevision) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicActionRevision) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


