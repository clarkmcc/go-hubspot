# BatchResponseTimelineEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The status of the batch response. Should always be COMPLETED if processed. | 
**Results** | [**[]TimelineEventResponse**](TimelineEventResponse.md) | Successfully created events. | 
**RequestedAt** | Pointer to **time.Time** | The time the request occurred. | [optional] 
**StartedAt** | **time.Time** | The time the request began processing. | 
**CompletedAt** | **time.Time** | The time the request was completed. | 
**Links** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBatchResponseTimelineEventResponse

`func NewBatchResponseTimelineEventResponse(status string, results []TimelineEventResponse, startedAt time.Time, completedAt time.Time, ) *BatchResponseTimelineEventResponse`

NewBatchResponseTimelineEventResponse instantiates a new BatchResponseTimelineEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResponseTimelineEventResponseWithDefaults

`func NewBatchResponseTimelineEventResponseWithDefaults() *BatchResponseTimelineEventResponse`

NewBatchResponseTimelineEventResponseWithDefaults instantiates a new BatchResponseTimelineEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BatchResponseTimelineEventResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchResponseTimelineEventResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchResponseTimelineEventResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetResults

`func (o *BatchResponseTimelineEventResponse) GetResults() []TimelineEventResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchResponseTimelineEventResponse) GetResultsOk() (*[]TimelineEventResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchResponseTimelineEventResponse) SetResults(v []TimelineEventResponse)`

SetResults sets Results field to given value.


### GetRequestedAt

`func (o *BatchResponseTimelineEventResponse) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *BatchResponseTimelineEventResponse) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *BatchResponseTimelineEventResponse) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *BatchResponseTimelineEventResponse) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *BatchResponseTimelineEventResponse) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BatchResponseTimelineEventResponse) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BatchResponseTimelineEventResponse) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetCompletedAt

`func (o *BatchResponseTimelineEventResponse) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BatchResponseTimelineEventResponse) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BatchResponseTimelineEventResponse) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetLinks

`func (o *BatchResponseTimelineEventResponse) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BatchResponseTimelineEventResponse) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BatchResponseTimelineEventResponse) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BatchResponseTimelineEventResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


