# PublicObjectListRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListId** | **string** | The ID of the list containing the imported objects. | 
**ObjectType** | **string** | The type of object contained in the list. | 

## Methods

### NewPublicObjectListRecord

`func NewPublicObjectListRecord(listId string, objectType string, ) *PublicObjectListRecord`

NewPublicObjectListRecord instantiates a new PublicObjectListRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicObjectListRecordWithDefaults

`func NewPublicObjectListRecordWithDefaults() *PublicObjectListRecord`

NewPublicObjectListRecordWithDefaults instantiates a new PublicObjectListRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListId

`func (o *PublicObjectListRecord) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *PublicObjectListRecord) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *PublicObjectListRecord) SetListId(v string)`

SetListId sets ListId field to given value.


### GetObjectType

`func (o *PublicObjectListRecord) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PublicObjectListRecord) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PublicObjectListRecord) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


