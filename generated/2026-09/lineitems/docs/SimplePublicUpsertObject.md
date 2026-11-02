# SimplePublicUpsertObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | **bool** | A boolean indicating whether this object is archived. | 
**ArchivedAt** | Pointer to **time.Time** | The date and time when this object was archived, in ISO 8601 format. | [optional] 
**CreatedAt** | **time.Time** | The date and time when this object was created, in ISO 8601 format. | 
**Id** | **string** | The unique identifier for this object. | 
**New** | **bool** | A boolean indicating whether this object is newly created. | 
**ObjectWriteTraceId** | Pointer to **string** | A string used to trace the write operations for this object. | [optional] 
**Properties** | **map[string]string** | A map of property names to their values for this object. | 
**PropertiesWithHistory** | Pointer to [**map[string][]ValueWithTimestamp**](array.md) | A map of property names to arrays of values with timestamps, showing the history of changes for each property. | [optional] 
**UpdatedAt** | **time.Time** | The date and time when this object was last updated, in ISO 8601 format. | 
**Url** | Pointer to **string** | The URL for accessing this object. | [optional] 
**Warnings** | Pointer to [**[]PublicObjectWarning**](PublicObjectWarning.md) | An array of warnings related to this object, each containing a category, message, and context. | [optional] 

## Methods

### NewSimplePublicUpsertObject

`func NewSimplePublicUpsertObject(archived bool, createdAt time.Time, id string, new bool, properties map[string]string, updatedAt time.Time, ) *SimplePublicUpsertObject`

NewSimplePublicUpsertObject instantiates a new SimplePublicUpsertObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplePublicUpsertObjectWithDefaults

`func NewSimplePublicUpsertObjectWithDefaults() *SimplePublicUpsertObject`

NewSimplePublicUpsertObjectWithDefaults instantiates a new SimplePublicUpsertObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *SimplePublicUpsertObject) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *SimplePublicUpsertObject) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *SimplePublicUpsertObject) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetArchivedAt

`func (o *SimplePublicUpsertObject) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *SimplePublicUpsertObject) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *SimplePublicUpsertObject) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *SimplePublicUpsertObject) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SimplePublicUpsertObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SimplePublicUpsertObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SimplePublicUpsertObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *SimplePublicUpsertObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimplePublicUpsertObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimplePublicUpsertObject) SetId(v string)`

SetId sets Id field to given value.


### GetNew

`func (o *SimplePublicUpsertObject) GetNew() bool`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *SimplePublicUpsertObject) GetNewOk() (*bool, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *SimplePublicUpsertObject) SetNew(v bool)`

SetNew sets New field to given value.


### GetObjectWriteTraceId

`func (o *SimplePublicUpsertObject) GetObjectWriteTraceId() string`

GetObjectWriteTraceId returns the ObjectWriteTraceId field if non-nil, zero value otherwise.

### GetObjectWriteTraceIdOk

`func (o *SimplePublicUpsertObject) GetObjectWriteTraceIdOk() (*string, bool)`

GetObjectWriteTraceIdOk returns a tuple with the ObjectWriteTraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectWriteTraceId

`func (o *SimplePublicUpsertObject) SetObjectWriteTraceId(v string)`

SetObjectWriteTraceId sets ObjectWriteTraceId field to given value.

### HasObjectWriteTraceId

`func (o *SimplePublicUpsertObject) HasObjectWriteTraceId() bool`

HasObjectWriteTraceId returns a boolean if a field has been set.

### GetProperties

`func (o *SimplePublicUpsertObject) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SimplePublicUpsertObject) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SimplePublicUpsertObject) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.


### GetPropertiesWithHistory

`func (o *SimplePublicUpsertObject) GetPropertiesWithHistory() map[string][]ValueWithTimestamp`

GetPropertiesWithHistory returns the PropertiesWithHistory field if non-nil, zero value otherwise.

### GetPropertiesWithHistoryOk

`func (o *SimplePublicUpsertObject) GetPropertiesWithHistoryOk() (*map[string][]ValueWithTimestamp, bool)`

GetPropertiesWithHistoryOk returns a tuple with the PropertiesWithHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesWithHistory

`func (o *SimplePublicUpsertObject) SetPropertiesWithHistory(v map[string][]ValueWithTimestamp)`

SetPropertiesWithHistory sets PropertiesWithHistory field to given value.

### HasPropertiesWithHistory

`func (o *SimplePublicUpsertObject) HasPropertiesWithHistory() bool`

HasPropertiesWithHistory returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SimplePublicUpsertObject) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SimplePublicUpsertObject) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SimplePublicUpsertObject) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUrl

`func (o *SimplePublicUpsertObject) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SimplePublicUpsertObject) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SimplePublicUpsertObject) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SimplePublicUpsertObject) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWarnings

`func (o *SimplePublicUpsertObject) GetWarnings() []PublicObjectWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SimplePublicUpsertObject) GetWarningsOk() (*[]PublicObjectWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SimplePublicUpsertObject) SetWarnings(v []PublicObjectWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SimplePublicUpsertObject) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


