# BatchResponseLabelsBetweenObjectPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | **time.Time** | The timestamp when the batch processing was completed, in ISO 8601 format. | 
**Links** | Pointer to **map[string]string** | An object containing relevant links related to the batch request. | [optional] 
**RequestedAt** | Pointer to **time.Time** | The timestamp when the batch request was initially made, in ISO 8601 format. | [optional] 
**Results** | [**[]LabelsBetweenObjectPair**](LabelsBetweenObjectPair.md) | An array of results for the batch operation, each representing a set of labels between an object pair. | 
**StartedAt** | **time.Time** | The timestamp when the batch processing began, in ISO 8601 format. | 
**Status** | **string** | The status of the batch processing request: \&quot;PENDING\&quot;, \&quot;PROCESSING\&quot;, \&quot;CANCELLED\&quot;, or \&quot;COMPLETE\&quot;. | 

## Methods

### NewBatchResponseLabelsBetweenObjectPair

`func NewBatchResponseLabelsBetweenObjectPair(completedAt time.Time, results []LabelsBetweenObjectPair, startedAt time.Time, status string, ) *BatchResponseLabelsBetweenObjectPair`

NewBatchResponseLabelsBetweenObjectPair instantiates a new BatchResponseLabelsBetweenObjectPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResponseLabelsBetweenObjectPairWithDefaults

`func NewBatchResponseLabelsBetweenObjectPairWithDefaults() *BatchResponseLabelsBetweenObjectPair`

NewBatchResponseLabelsBetweenObjectPairWithDefaults instantiates a new BatchResponseLabelsBetweenObjectPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *BatchResponseLabelsBetweenObjectPair) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BatchResponseLabelsBetweenObjectPair) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BatchResponseLabelsBetweenObjectPair) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetLinks

`func (o *BatchResponseLabelsBetweenObjectPair) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BatchResponseLabelsBetweenObjectPair) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BatchResponseLabelsBetweenObjectPair) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BatchResponseLabelsBetweenObjectPair) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRequestedAt

`func (o *BatchResponseLabelsBetweenObjectPair) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *BatchResponseLabelsBetweenObjectPair) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *BatchResponseLabelsBetweenObjectPair) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *BatchResponseLabelsBetweenObjectPair) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetResults

`func (o *BatchResponseLabelsBetweenObjectPair) GetResults() []LabelsBetweenObjectPair`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchResponseLabelsBetweenObjectPair) GetResultsOk() (*[]LabelsBetweenObjectPair, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchResponseLabelsBetweenObjectPair) SetResults(v []LabelsBetweenObjectPair)`

SetResults sets Results field to given value.


### GetStartedAt

`func (o *BatchResponseLabelsBetweenObjectPair) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BatchResponseLabelsBetweenObjectPair) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BatchResponseLabelsBetweenObjectPair) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetStatus

`func (o *BatchResponseLabelsBetweenObjectPair) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchResponseLabelsBetweenObjectPair) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchResponseLabelsBetweenObjectPair) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


