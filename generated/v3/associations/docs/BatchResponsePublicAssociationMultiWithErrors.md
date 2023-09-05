# BatchResponsePublicAssociationMultiWithErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Results** | [**[]PublicAssociationMulti**](PublicAssociationMulti.md) |  | 
**NumErrors** | Pointer to **int32** |  | [optional] 
**Errors** | Pointer to [**[]StandardError**](StandardError.md) |  | [optional] 
**RequestedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | **time.Time** |  | 
**CompletedAt** | **time.Time** |  | 
**Links** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBatchResponsePublicAssociationMultiWithErrors

`func NewBatchResponsePublicAssociationMultiWithErrors(status string, results []PublicAssociationMulti, startedAt time.Time, completedAt time.Time, ) *BatchResponsePublicAssociationMultiWithErrors`

NewBatchResponsePublicAssociationMultiWithErrors instantiates a new BatchResponsePublicAssociationMultiWithErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResponsePublicAssociationMultiWithErrorsWithDefaults

`func NewBatchResponsePublicAssociationMultiWithErrorsWithDefaults() *BatchResponsePublicAssociationMultiWithErrors`

NewBatchResponsePublicAssociationMultiWithErrorsWithDefaults instantiates a new BatchResponsePublicAssociationMultiWithErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BatchResponsePublicAssociationMultiWithErrors) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchResponsePublicAssociationMultiWithErrors) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchResponsePublicAssociationMultiWithErrors) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetResults

`func (o *BatchResponsePublicAssociationMultiWithErrors) GetResults() []PublicAssociationMulti`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchResponsePublicAssociationMultiWithErrors) GetResultsOk() (*[]PublicAssociationMulti, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchResponsePublicAssociationMultiWithErrors) SetResults(v []PublicAssociationMulti)`

SetResults sets Results field to given value.


### GetNumErrors

`func (o *BatchResponsePublicAssociationMultiWithErrors) GetNumErrors() int32`

GetNumErrors returns the NumErrors field if non-nil, zero value otherwise.

### GetNumErrorsOk

`func (o *BatchResponsePublicAssociationMultiWithErrors) GetNumErrorsOk() (*int32, bool)`

GetNumErrorsOk returns a tuple with the NumErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumErrors

`func (o *BatchResponsePublicAssociationMultiWithErrors) SetNumErrors(v int32)`

SetNumErrors sets NumErrors field to given value.

### HasNumErrors

`func (o *BatchResponsePublicAssociationMultiWithErrors) HasNumErrors() bool`

HasNumErrors returns a boolean if a field has been set.

### GetErrors

`func (o *BatchResponsePublicAssociationMultiWithErrors) GetErrors() []StandardError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BatchResponsePublicAssociationMultiWithErrors) GetErrorsOk() (*[]StandardError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BatchResponsePublicAssociationMultiWithErrors) SetErrors(v []StandardError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BatchResponsePublicAssociationMultiWithErrors) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetRequestedAt

`func (o *BatchResponsePublicAssociationMultiWithErrors) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *BatchResponsePublicAssociationMultiWithErrors) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *BatchResponsePublicAssociationMultiWithErrors) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *BatchResponsePublicAssociationMultiWithErrors) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *BatchResponsePublicAssociationMultiWithErrors) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BatchResponsePublicAssociationMultiWithErrors) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BatchResponsePublicAssociationMultiWithErrors) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetCompletedAt

`func (o *BatchResponsePublicAssociationMultiWithErrors) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BatchResponsePublicAssociationMultiWithErrors) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BatchResponsePublicAssociationMultiWithErrors) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetLinks

`func (o *BatchResponsePublicAssociationMultiWithErrors) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BatchResponsePublicAssociationMultiWithErrors) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BatchResponsePublicAssociationMultiWithErrors) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BatchResponsePublicAssociationMultiWithErrors) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


