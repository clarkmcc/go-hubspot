# SimplePublicObjectWithAssociations

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
**Associations** | Pointer to [**map[string]CollectionResponseAssociatedId**](CollectionResponseAssociatedId.md) |  | [optional] 

## Methods

### NewSimplePublicObjectWithAssociations

`func NewSimplePublicObjectWithAssociations(id string, properties map[string]string, createdAt time.Time, updatedAt time.Time, ) *SimplePublicObjectWithAssociations`

NewSimplePublicObjectWithAssociations instantiates a new SimplePublicObjectWithAssociations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplePublicObjectWithAssociationsWithDefaults

`func NewSimplePublicObjectWithAssociationsWithDefaults() *SimplePublicObjectWithAssociations`

NewSimplePublicObjectWithAssociationsWithDefaults instantiates a new SimplePublicObjectWithAssociations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimplePublicObjectWithAssociations) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimplePublicObjectWithAssociations) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimplePublicObjectWithAssociations) SetId(v string)`

SetId sets Id field to given value.


### GetProperties

`func (o *SimplePublicObjectWithAssociations) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SimplePublicObjectWithAssociations) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SimplePublicObjectWithAssociations) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.


### GetPropertiesWithHistory

`func (o *SimplePublicObjectWithAssociations) GetPropertiesWithHistory() map[string][]ValueWithTimestamp`

GetPropertiesWithHistory returns the PropertiesWithHistory field if non-nil, zero value otherwise.

### GetPropertiesWithHistoryOk

`func (o *SimplePublicObjectWithAssociations) GetPropertiesWithHistoryOk() (*map[string][]ValueWithTimestamp, bool)`

GetPropertiesWithHistoryOk returns a tuple with the PropertiesWithHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesWithHistory

`func (o *SimplePublicObjectWithAssociations) SetPropertiesWithHistory(v map[string][]ValueWithTimestamp)`

SetPropertiesWithHistory sets PropertiesWithHistory field to given value.

### HasPropertiesWithHistory

`func (o *SimplePublicObjectWithAssociations) HasPropertiesWithHistory() bool`

HasPropertiesWithHistory returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SimplePublicObjectWithAssociations) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SimplePublicObjectWithAssociations) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SimplePublicObjectWithAssociations) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SimplePublicObjectWithAssociations) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SimplePublicObjectWithAssociations) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SimplePublicObjectWithAssociations) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetArchived

`func (o *SimplePublicObjectWithAssociations) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *SimplePublicObjectWithAssociations) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *SimplePublicObjectWithAssociations) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *SimplePublicObjectWithAssociations) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetArchivedAt

`func (o *SimplePublicObjectWithAssociations) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *SimplePublicObjectWithAssociations) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *SimplePublicObjectWithAssociations) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *SimplePublicObjectWithAssociations) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetAssociations

`func (o *SimplePublicObjectWithAssociations) GetAssociations() map[string]CollectionResponseAssociatedId`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *SimplePublicObjectWithAssociations) GetAssociationsOk() (*map[string]CollectionResponseAssociatedId, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *SimplePublicObjectWithAssociations) SetAssociations(v map[string]CollectionResponseAssociatedId)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *SimplePublicObjectWithAssociations) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


