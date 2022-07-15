# SimplePublicObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Properties** | **map[string]string** |  | 
**PropertiesWithHistory** | Pointer to [**map[string][]ValueWithTimestamp**](array.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Archived** | Pointer to **bool** |  | [optional] 
**ArchivedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSimplePublicObject

`func NewSimplePublicObject(id string, properties map[string]string, createdAt time.Time, updatedAt time.Time, ) *SimplePublicObject`

NewSimplePublicObject instantiates a new SimplePublicObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplePublicObjectWithDefaults

`func NewSimplePublicObjectWithDefaults() *SimplePublicObject`

NewSimplePublicObjectWithDefaults instantiates a new SimplePublicObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimplePublicObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimplePublicObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimplePublicObject) SetId(v string)`

SetId sets Id field to given value.


### GetProperties

`func (o *SimplePublicObject) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SimplePublicObject) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SimplePublicObject) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.


### GetPropertiesWithHistory

`func (o *SimplePublicObject) GetPropertiesWithHistory() map[string][]ValueWithTimestamp`

GetPropertiesWithHistory returns the PropertiesWithHistory field if non-nil, zero value otherwise.

### GetPropertiesWithHistoryOk

`func (o *SimplePublicObject) GetPropertiesWithHistoryOk() (*map[string][]ValueWithTimestamp, bool)`

GetPropertiesWithHistoryOk returns a tuple with the PropertiesWithHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesWithHistory

`func (o *SimplePublicObject) SetPropertiesWithHistory(v map[string][]ValueWithTimestamp)`

SetPropertiesWithHistory sets PropertiesWithHistory field to given value.

### HasPropertiesWithHistory

`func (o *SimplePublicObject) HasPropertiesWithHistory() bool`

HasPropertiesWithHistory returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SimplePublicObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SimplePublicObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SimplePublicObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SimplePublicObject) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SimplePublicObject) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SimplePublicObject) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetArchived

`func (o *SimplePublicObject) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *SimplePublicObject) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *SimplePublicObject) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *SimplePublicObject) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetArchivedAt

`func (o *SimplePublicObject) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *SimplePublicObject) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *SimplePublicObject) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *SimplePublicObject) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


