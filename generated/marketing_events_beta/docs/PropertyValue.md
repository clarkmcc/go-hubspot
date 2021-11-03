# PropertyValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | **string** |  | 
**Timestamp** | **int64** |  | 
**SourceId** | **string** |  | 
**SourceLabel** | **string** |  | 
**Source** | **string** |  | 
**SelectedByUser** | **bool** |  | 
**SelectedByUserTimestamp** | **int64** |  | 
**SourceVid** | **[]int64** |  | 
**SourceMetadata** | **string** | Source metadata encoded as a base64 string. For example: &#x60;ZXhhbXBsZSBzdHJpbmc&#x3D;&#x60; | 
**RequestId** | **string** |  | 
**UpdatedByUserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewPropertyValue

`func NewPropertyValue(name string, value string, timestamp int64, sourceId string, sourceLabel string, source string, selectedByUser bool, selectedByUserTimestamp int64, sourceVid []int64, sourceMetadata string, requestId string, ) *PropertyValue`

NewPropertyValue instantiates a new PropertyValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyValueWithDefaults

`func NewPropertyValueWithDefaults() *PropertyValue`

NewPropertyValueWithDefaults instantiates a new PropertyValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PropertyValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PropertyValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PropertyValue) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *PropertyValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PropertyValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PropertyValue) SetValue(v string)`

SetValue sets Value field to given value.


### GetTimestamp

`func (o *PropertyValue) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PropertyValue) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PropertyValue) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetSourceId

`func (o *PropertyValue) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *PropertyValue) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *PropertyValue) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceLabel

`func (o *PropertyValue) GetSourceLabel() string`

GetSourceLabel returns the SourceLabel field if non-nil, zero value otherwise.

### GetSourceLabelOk

`func (o *PropertyValue) GetSourceLabelOk() (*string, bool)`

GetSourceLabelOk returns a tuple with the SourceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLabel

`func (o *PropertyValue) SetSourceLabel(v string)`

SetSourceLabel sets SourceLabel field to given value.


### GetSource

`func (o *PropertyValue) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PropertyValue) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PropertyValue) SetSource(v string)`

SetSource sets Source field to given value.


### GetSelectedByUser

`func (o *PropertyValue) GetSelectedByUser() bool`

GetSelectedByUser returns the SelectedByUser field if non-nil, zero value otherwise.

### GetSelectedByUserOk

`func (o *PropertyValue) GetSelectedByUserOk() (*bool, bool)`

GetSelectedByUserOk returns a tuple with the SelectedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedByUser

`func (o *PropertyValue) SetSelectedByUser(v bool)`

SetSelectedByUser sets SelectedByUser field to given value.


### GetSelectedByUserTimestamp

`func (o *PropertyValue) GetSelectedByUserTimestamp() int64`

GetSelectedByUserTimestamp returns the SelectedByUserTimestamp field if non-nil, zero value otherwise.

### GetSelectedByUserTimestampOk

`func (o *PropertyValue) GetSelectedByUserTimestampOk() (*int64, bool)`

GetSelectedByUserTimestampOk returns a tuple with the SelectedByUserTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedByUserTimestamp

`func (o *PropertyValue) SetSelectedByUserTimestamp(v int64)`

SetSelectedByUserTimestamp sets SelectedByUserTimestamp field to given value.


### GetSourceVid

`func (o *PropertyValue) GetSourceVid() []int64`

GetSourceVid returns the SourceVid field if non-nil, zero value otherwise.

### GetSourceVidOk

`func (o *PropertyValue) GetSourceVidOk() (*[]int64, bool)`

GetSourceVidOk returns a tuple with the SourceVid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVid

`func (o *PropertyValue) SetSourceVid(v []int64)`

SetSourceVid sets SourceVid field to given value.


### GetSourceMetadata

`func (o *PropertyValue) GetSourceMetadata() string`

GetSourceMetadata returns the SourceMetadata field if non-nil, zero value otherwise.

### GetSourceMetadataOk

`func (o *PropertyValue) GetSourceMetadataOk() (*string, bool)`

GetSourceMetadataOk returns a tuple with the SourceMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMetadata

`func (o *PropertyValue) SetSourceMetadata(v string)`

SetSourceMetadata sets SourceMetadata field to given value.


### GetRequestId

`func (o *PropertyValue) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PropertyValue) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PropertyValue) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetUpdatedByUserId

`func (o *PropertyValue) GetUpdatedByUserId() int32`

GetUpdatedByUserId returns the UpdatedByUserId field if non-nil, zero value otherwise.

### GetUpdatedByUserIdOk

`func (o *PropertyValue) GetUpdatedByUserIdOk() (*int32, bool)`

GetUpdatedByUserIdOk returns a tuple with the UpdatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUserId

`func (o *PropertyValue) SetUpdatedByUserId(v int32)`

SetUpdatedByUserId sets UpdatedByUserId field to given value.

### HasUpdatedByUserId

`func (o *PropertyValue) HasUpdatedByUserId() bool`

HasUpdatedByUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


