# EmailSendStatusView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusId** | **string** | Identifier used to query the status of the send. | 
**SendResult** | Pointer to **string** | Result of the send. | [optional] 
**RequestedAt** | Pointer to **time.Time** | Time when the send was requested. | [optional] 
**StartedAt** | Pointer to **time.Time** | Time when the send began processing. | [optional] 
**CompletedAt** | Pointer to **time.Time** | Time when the send was completed. | [optional] 
**Status** | **string** | Status of the send request. | 
**EventId** | Pointer to [**EventIdView**](EventIdView.md) |  | [optional] 

## Methods

### NewEmailSendStatusView

`func NewEmailSendStatusView(statusId string, status string, ) *EmailSendStatusView`

NewEmailSendStatusView instantiates a new EmailSendStatusView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSendStatusViewWithDefaults

`func NewEmailSendStatusViewWithDefaults() *EmailSendStatusView`

NewEmailSendStatusViewWithDefaults instantiates a new EmailSendStatusView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusId

`func (o *EmailSendStatusView) GetStatusId() string`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *EmailSendStatusView) GetStatusIdOk() (*string, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *EmailSendStatusView) SetStatusId(v string)`

SetStatusId sets StatusId field to given value.


### GetSendResult

`func (o *EmailSendStatusView) GetSendResult() string`

GetSendResult returns the SendResult field if non-nil, zero value otherwise.

### GetSendResultOk

`func (o *EmailSendStatusView) GetSendResultOk() (*string, bool)`

GetSendResultOk returns a tuple with the SendResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResult

`func (o *EmailSendStatusView) SetSendResult(v string)`

SetSendResult sets SendResult field to given value.

### HasSendResult

`func (o *EmailSendStatusView) HasSendResult() bool`

HasSendResult returns a boolean if a field has been set.

### GetRequestedAt

`func (o *EmailSendStatusView) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *EmailSendStatusView) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *EmailSendStatusView) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *EmailSendStatusView) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *EmailSendStatusView) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *EmailSendStatusView) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *EmailSendStatusView) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *EmailSendStatusView) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *EmailSendStatusView) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *EmailSendStatusView) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *EmailSendStatusView) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *EmailSendStatusView) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetStatus

`func (o *EmailSendStatusView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailSendStatusView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailSendStatusView) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetEventId

`func (o *EmailSendStatusView) GetEventId() EventIdView`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *EmailSendStatusView) GetEventIdOk() (*EventIdView, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *EmailSendStatusView) SetEventId(v EventIdView)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *EmailSendStatusView) HasEventId() bool`

HasEventId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


