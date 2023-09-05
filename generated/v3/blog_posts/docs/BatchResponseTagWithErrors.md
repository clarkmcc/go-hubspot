# BatchResponseTagWithErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Status of batch operation. | 
**Results** | [**[]Tag**](Tag.md) | Results of batch operation. | 
**NumErrors** | Pointer to **int32** | Number of errors. | [optional] 
**Errors** | Pointer to [**[]StandardError**](StandardError.md) | Errors in batch operation. | [optional] 
**RequestedAt** | Pointer to **time.Time** | Time of batch operation request. | [optional] 
**StartedAt** | **time.Time** | Time of batch operation start. | 
**CompletedAt** | **time.Time** | Time of batch operation completion. | 
**Links** | Pointer to **map[string]string** | Links associated with batch operation. | [optional] 

## Methods

### NewBatchResponseTagWithErrors

`func NewBatchResponseTagWithErrors(status string, results []Tag, startedAt time.Time, completedAt time.Time, ) *BatchResponseTagWithErrors`

NewBatchResponseTagWithErrors instantiates a new BatchResponseTagWithErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResponseTagWithErrorsWithDefaults

`func NewBatchResponseTagWithErrorsWithDefaults() *BatchResponseTagWithErrors`

NewBatchResponseTagWithErrorsWithDefaults instantiates a new BatchResponseTagWithErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BatchResponseTagWithErrors) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchResponseTagWithErrors) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchResponseTagWithErrors) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetResults

`func (o *BatchResponseTagWithErrors) GetResults() []Tag`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchResponseTagWithErrors) GetResultsOk() (*[]Tag, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchResponseTagWithErrors) SetResults(v []Tag)`

SetResults sets Results field to given value.


### GetNumErrors

`func (o *BatchResponseTagWithErrors) GetNumErrors() int32`

GetNumErrors returns the NumErrors field if non-nil, zero value otherwise.

### GetNumErrorsOk

`func (o *BatchResponseTagWithErrors) GetNumErrorsOk() (*int32, bool)`

GetNumErrorsOk returns a tuple with the NumErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumErrors

`func (o *BatchResponseTagWithErrors) SetNumErrors(v int32)`

SetNumErrors sets NumErrors field to given value.

### HasNumErrors

`func (o *BatchResponseTagWithErrors) HasNumErrors() bool`

HasNumErrors returns a boolean if a field has been set.

### GetErrors

`func (o *BatchResponseTagWithErrors) GetErrors() []StandardError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BatchResponseTagWithErrors) GetErrorsOk() (*[]StandardError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BatchResponseTagWithErrors) SetErrors(v []StandardError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BatchResponseTagWithErrors) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetRequestedAt

`func (o *BatchResponseTagWithErrors) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *BatchResponseTagWithErrors) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *BatchResponseTagWithErrors) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *BatchResponseTagWithErrors) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *BatchResponseTagWithErrors) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BatchResponseTagWithErrors) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BatchResponseTagWithErrors) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetCompletedAt

`func (o *BatchResponseTagWithErrors) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BatchResponseTagWithErrors) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BatchResponseTagWithErrors) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetLinks

`func (o *BatchResponseTagWithErrors) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BatchResponseTagWithErrors) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BatchResponseTagWithErrors) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BatchResponseTagWithErrors) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


