# BatchResponsePublicAssociationMultiWithLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Results** | [**[]PublicAssociationMultiWithLabel**](PublicAssociationMultiWithLabel.md) |  | 
**RequestedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | **time.Time** |  | 
**CompletedAt** | **time.Time** |  | 
**Links** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBatchResponsePublicAssociationMultiWithLabel

`func NewBatchResponsePublicAssociationMultiWithLabel(status string, results []PublicAssociationMultiWithLabel, startedAt time.Time, completedAt time.Time, ) *BatchResponsePublicAssociationMultiWithLabel`

NewBatchResponsePublicAssociationMultiWithLabel instantiates a new BatchResponsePublicAssociationMultiWithLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResponsePublicAssociationMultiWithLabelWithDefaults

`func NewBatchResponsePublicAssociationMultiWithLabelWithDefaults() *BatchResponsePublicAssociationMultiWithLabel`

NewBatchResponsePublicAssociationMultiWithLabelWithDefaults instantiates a new BatchResponsePublicAssociationMultiWithLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BatchResponsePublicAssociationMultiWithLabel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchResponsePublicAssociationMultiWithLabel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchResponsePublicAssociationMultiWithLabel) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetResults

`func (o *BatchResponsePublicAssociationMultiWithLabel) GetResults() []PublicAssociationMultiWithLabel`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchResponsePublicAssociationMultiWithLabel) GetResultsOk() (*[]PublicAssociationMultiWithLabel, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchResponsePublicAssociationMultiWithLabel) SetResults(v []PublicAssociationMultiWithLabel)`

SetResults sets Results field to given value.


### GetRequestedAt

`func (o *BatchResponsePublicAssociationMultiWithLabel) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *BatchResponsePublicAssociationMultiWithLabel) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *BatchResponsePublicAssociationMultiWithLabel) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *BatchResponsePublicAssociationMultiWithLabel) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *BatchResponsePublicAssociationMultiWithLabel) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BatchResponsePublicAssociationMultiWithLabel) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BatchResponsePublicAssociationMultiWithLabel) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetCompletedAt

`func (o *BatchResponsePublicAssociationMultiWithLabel) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BatchResponsePublicAssociationMultiWithLabel) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BatchResponsePublicAssociationMultiWithLabel) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetLinks

`func (o *BatchResponsePublicAssociationMultiWithLabel) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BatchResponsePublicAssociationMultiWithLabel) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BatchResponsePublicAssociationMultiWithLabel) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BatchResponsePublicAssociationMultiWithLabel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


