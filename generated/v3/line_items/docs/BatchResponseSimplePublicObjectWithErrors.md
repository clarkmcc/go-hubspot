# BatchResponseSimplePublicObjectWithErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Results** | [**[]SimplePublicObject**](SimplePublicObject.md) |  | 
**NumErrors** | Pointer to **int32** |  | [optional] 
**Errors** | Pointer to [**[]StandardError**](StandardError.md) |  | [optional] 
**RequestedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | **time.Time** |  | 
**CompletedAt** | **time.Time** |  | 
**Links** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBatchResponseSimplePublicObjectWithErrors

`func NewBatchResponseSimplePublicObjectWithErrors(status string, results []SimplePublicObject, startedAt time.Time, completedAt time.Time, ) *BatchResponseSimplePublicObjectWithErrors`

NewBatchResponseSimplePublicObjectWithErrors instantiates a new BatchResponseSimplePublicObjectWithErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResponseSimplePublicObjectWithErrorsWithDefaults

`func NewBatchResponseSimplePublicObjectWithErrorsWithDefaults() *BatchResponseSimplePublicObjectWithErrors`

NewBatchResponseSimplePublicObjectWithErrorsWithDefaults instantiates a new BatchResponseSimplePublicObjectWithErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BatchResponseSimplePublicObjectWithErrors) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchResponseSimplePublicObjectWithErrors) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchResponseSimplePublicObjectWithErrors) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetResults

`func (o *BatchResponseSimplePublicObjectWithErrors) GetResults() []SimplePublicObject`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchResponseSimplePublicObjectWithErrors) GetResultsOk() (*[]SimplePublicObject, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchResponseSimplePublicObjectWithErrors) SetResults(v []SimplePublicObject)`

SetResults sets Results field to given value.


### GetNumErrors

`func (o *BatchResponseSimplePublicObjectWithErrors) GetNumErrors() int32`

GetNumErrors returns the NumErrors field if non-nil, zero value otherwise.

### GetNumErrorsOk

`func (o *BatchResponseSimplePublicObjectWithErrors) GetNumErrorsOk() (*int32, bool)`

GetNumErrorsOk returns a tuple with the NumErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumErrors

`func (o *BatchResponseSimplePublicObjectWithErrors) SetNumErrors(v int32)`

SetNumErrors sets NumErrors field to given value.

### HasNumErrors

`func (o *BatchResponseSimplePublicObjectWithErrors) HasNumErrors() bool`

HasNumErrors returns a boolean if a field has been set.

### GetErrors

`func (o *BatchResponseSimplePublicObjectWithErrors) GetErrors() []StandardError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BatchResponseSimplePublicObjectWithErrors) GetErrorsOk() (*[]StandardError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BatchResponseSimplePublicObjectWithErrors) SetErrors(v []StandardError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BatchResponseSimplePublicObjectWithErrors) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetRequestedAt

`func (o *BatchResponseSimplePublicObjectWithErrors) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *BatchResponseSimplePublicObjectWithErrors) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *BatchResponseSimplePublicObjectWithErrors) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *BatchResponseSimplePublicObjectWithErrors) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *BatchResponseSimplePublicObjectWithErrors) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BatchResponseSimplePublicObjectWithErrors) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BatchResponseSimplePublicObjectWithErrors) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetCompletedAt

`func (o *BatchResponseSimplePublicObjectWithErrors) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BatchResponseSimplePublicObjectWithErrors) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BatchResponseSimplePublicObjectWithErrors) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetLinks

`func (o *BatchResponseSimplePublicObjectWithErrors) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BatchResponseSimplePublicObjectWithErrors) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BatchResponseSimplePublicObjectWithErrors) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BatchResponseSimplePublicObjectWithErrors) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


