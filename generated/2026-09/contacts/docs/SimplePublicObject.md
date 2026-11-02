# SimplePublicObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | **bool** | A boolean indicating whether this object is archived. | 
**ArchivedAt** | Pointer to **time.Time** | The date and time when this object was archived, in ISO 8601 format. | [optional] 
**CreatedAt** | **time.Time** | The date and time when this object was created, in ISO 8601 format. | 
**Id** | **string** | The unique identifier for this object, represented as a string. | 
**ObjectWriteTraceId** | Pointer to **string** | A string representing the trace ID for write operations related to this object. | [optional] 
**Properties** | **map[string]string** | A map of property names to their values, where each value is a string. | 
**PropertiesWithHistory** | Pointer to [**map[string][]ValueWithTimestamp**](array.md) | A map of property names to their historical values, each represented as an array of ValueWithTimestamp objects. | [optional] 
**UpdatedAt** | **time.Time** | The date and time when this object was last updated, in ISO 8601 format. | 
**Url** | Pointer to **string** | A string containing the URL of this object. | [optional] 
**Warnings** | Pointer to [**[]PublicObjectWarning**](PublicObjectWarning.md) | An array of PublicObjectWarning objects, each containing warnings related to this object. | [optional] 

## Methods

### NewSimplePublicObject

`func NewSimplePublicObject(archived bool, createdAt time.Time, id string, properties map[string]string, updatedAt time.Time, ) *SimplePublicObject`

NewSimplePublicObject instantiates a new SimplePublicObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplePublicObjectWithDefaults

`func NewSimplePublicObjectWithDefaults() *SimplePublicObject`

NewSimplePublicObjectWithDefaults instantiates a new SimplePublicObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetObjectWriteTraceId

`func (o *SimplePublicObject) GetObjectWriteTraceId() string`

GetObjectWriteTraceId returns the ObjectWriteTraceId field if non-nil, zero value otherwise.

### GetObjectWriteTraceIdOk

`func (o *SimplePublicObject) GetObjectWriteTraceIdOk() (*string, bool)`

GetObjectWriteTraceIdOk returns a tuple with the ObjectWriteTraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectWriteTraceId

`func (o *SimplePublicObject) SetObjectWriteTraceId(v string)`

SetObjectWriteTraceId sets ObjectWriteTraceId field to given value.

### HasObjectWriteTraceId

`func (o *SimplePublicObject) HasObjectWriteTraceId() bool`

HasObjectWriteTraceId returns a boolean if a field has been set.

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


### GetUrl

`func (o *SimplePublicObject) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SimplePublicObject) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SimplePublicObject) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SimplePublicObject) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWarnings

`func (o *SimplePublicObject) GetWarnings() []PublicObjectWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SimplePublicObject) GetWarningsOk() (*[]PublicObjectWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SimplePublicObject) SetWarnings(v []PublicObjectWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SimplePublicObject) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


