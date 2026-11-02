# ValueWithTimestamp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | Pointer to **string** | The identifier of the source from which the value was obtained. It is a string. | [optional] 
**SourceLabel** | Pointer to **string** | A label providing additional context about the source. It is a string. | [optional] 
**SourceType** | **string** | The type of source from which the value was obtained. It is a string and is required. | 
**Timestamp** | **time.Time** | The date and time when the value was recorded, in ISO 8601 format. | 
**UpdatedByUserId** | Pointer to **int32** | The ID of the user who last updated the value. It is an integer. | [optional] 
**Value** | **string** | The value associated with the timestamp. It is a string. | 

## Methods

### NewValueWithTimestamp

`func NewValueWithTimestamp(sourceType string, timestamp time.Time, value string, ) *ValueWithTimestamp`

NewValueWithTimestamp instantiates a new ValueWithTimestamp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueWithTimestampWithDefaults

`func NewValueWithTimestampWithDefaults() *ValueWithTimestamp`

NewValueWithTimestampWithDefaults instantiates a new ValueWithTimestamp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *ValueWithTimestamp) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ValueWithTimestamp) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ValueWithTimestamp) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ValueWithTimestamp) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceLabel

`func (o *ValueWithTimestamp) GetSourceLabel() string`

GetSourceLabel returns the SourceLabel field if non-nil, zero value otherwise.

### GetSourceLabelOk

`func (o *ValueWithTimestamp) GetSourceLabelOk() (*string, bool)`

GetSourceLabelOk returns a tuple with the SourceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLabel

`func (o *ValueWithTimestamp) SetSourceLabel(v string)`

SetSourceLabel sets SourceLabel field to given value.

### HasSourceLabel

`func (o *ValueWithTimestamp) HasSourceLabel() bool`

HasSourceLabel returns a boolean if a field has been set.

### GetSourceType

`func (o *ValueWithTimestamp) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ValueWithTimestamp) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ValueWithTimestamp) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetTimestamp

`func (o *ValueWithTimestamp) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ValueWithTimestamp) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ValueWithTimestamp) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetUpdatedByUserId

`func (o *ValueWithTimestamp) GetUpdatedByUserId() int32`

GetUpdatedByUserId returns the UpdatedByUserId field if non-nil, zero value otherwise.

### GetUpdatedByUserIdOk

`func (o *ValueWithTimestamp) GetUpdatedByUserIdOk() (*int32, bool)`

GetUpdatedByUserIdOk returns a tuple with the UpdatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUserId

`func (o *ValueWithTimestamp) SetUpdatedByUserId(v int32)`

SetUpdatedByUserId sets UpdatedByUserId field to given value.

### HasUpdatedByUserId

`func (o *ValueWithTimestamp) HasUpdatedByUserId() bool`

HasUpdatedByUserId returns a boolean if a field has been set.

### GetValue

`func (o *ValueWithTimestamp) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ValueWithTimestamp) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ValueWithTimestamp) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


