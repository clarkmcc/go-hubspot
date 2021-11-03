# PublicImportMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectLists** | [**[]PublicObjectListRecord**](PublicObjectListRecord.md) | The lists containing the imported objects. | 
**Counters** | **map[string]int32** | Summarized outcomes of each row a developer attempted to import into HubSpot. | 
**FileIds** | **[]string** | The IDs of files uploaded in the File Manager API. | 

## Methods

### NewPublicImportMetadata

`func NewPublicImportMetadata(objectLists []PublicObjectListRecord, counters map[string]int32, fileIds []string, ) *PublicImportMetadata`

NewPublicImportMetadata instantiates a new PublicImportMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicImportMetadataWithDefaults

`func NewPublicImportMetadataWithDefaults() *PublicImportMetadata`

NewPublicImportMetadataWithDefaults instantiates a new PublicImportMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectLists

`func (o *PublicImportMetadata) GetObjectLists() []PublicObjectListRecord`

GetObjectLists returns the ObjectLists field if non-nil, zero value otherwise.

### GetObjectListsOk

`func (o *PublicImportMetadata) GetObjectListsOk() (*[]PublicObjectListRecord, bool)`

GetObjectListsOk returns a tuple with the ObjectLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectLists

`func (o *PublicImportMetadata) SetObjectLists(v []PublicObjectListRecord)`

SetObjectLists sets ObjectLists field to given value.


### GetCounters

`func (o *PublicImportMetadata) GetCounters() map[string]int32`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *PublicImportMetadata) GetCountersOk() (*map[string]int32, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *PublicImportMetadata) SetCounters(v map[string]int32)`

SetCounters sets Counters field to given value.


### GetFileIds

`func (o *PublicImportMetadata) GetFileIds() []string`

GetFileIds returns the FileIds field if non-nil, zero value otherwise.

### GetFileIdsOk

`func (o *PublicImportMetadata) GetFileIdsOk() (*[]string, bool)`

GetFileIdsOk returns a tuple with the FileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileIds

`func (o *PublicImportMetadata) SetFileIds(v []string)`

SetFileIds sets FileIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


