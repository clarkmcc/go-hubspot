# BatchResponseSimplePublicUpsertObjectWithErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | **time.Time** | The date and time when the batch operation was completed, in ISO 8601 format. | 
**Errors** | Pointer to [**[]StandardError**](StandardError.md) | An array of StandardError objects, detailing any errors that occurred during the batch operation. | [optional] 
**Links** | Pointer to **map[string]string** | A map of link names to associated URIs, providing additional information or actions related to the batch operation. | [optional] 
**NumErrors** | Pointer to **int32** | The number of errors encountered during the batch operation. | [optional] 
**RequestedAt** | Pointer to **time.Time** | The date and time when the batch operation was requested, in ISO 8601 format. | [optional] 
**Results** | [**[]SimplePublicUpsertObject**](SimplePublicUpsertObject.md) | An array of SimplePublicUpsertObject items, representing the objects that were successfully upserted in the batch operation. | 
**StartedAt** | **time.Time** | The date and time when the batch operation started, in ISO 8601 format. | 
**Status** | **string** | The current status of the batch operation. Valid values include &#39;PENDING&#39;, &#39;PROCESSING&#39;, &#39;CANCELED&#39;, and &#39;COMPLETE&#39;. | 

## Methods

### NewBatchResponseSimplePublicUpsertObjectWithErrors

`func NewBatchResponseSimplePublicUpsertObjectWithErrors(completedAt time.Time, results []SimplePublicUpsertObject, startedAt time.Time, status string, ) *BatchResponseSimplePublicUpsertObjectWithErrors`

NewBatchResponseSimplePublicUpsertObjectWithErrors instantiates a new BatchResponseSimplePublicUpsertObjectWithErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResponseSimplePublicUpsertObjectWithErrorsWithDefaults

`func NewBatchResponseSimplePublicUpsertObjectWithErrorsWithDefaults() *BatchResponseSimplePublicUpsertObjectWithErrors`

NewBatchResponseSimplePublicUpsertObjectWithErrorsWithDefaults instantiates a new BatchResponseSimplePublicUpsertObjectWithErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetErrors

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) GetErrors() []StandardError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) GetErrorsOk() (*[]StandardError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) SetErrors(v []StandardError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetLinks

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetNumErrors

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) GetNumErrors() int32`

GetNumErrors returns the NumErrors field if non-nil, zero value otherwise.

### GetNumErrorsOk

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) GetNumErrorsOk() (*int32, bool)`

GetNumErrorsOk returns a tuple with the NumErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumErrors

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) SetNumErrors(v int32)`

SetNumErrors sets NumErrors field to given value.

### HasNumErrors

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) HasNumErrors() bool`

HasNumErrors returns a boolean if a field has been set.

### GetRequestedAt

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetResults

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) GetResults() []SimplePublicUpsertObject`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) GetResultsOk() (*[]SimplePublicUpsertObject, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) SetResults(v []SimplePublicUpsertObject)`

SetResults sets Results field to given value.


### GetStartedAt

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetStatus

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchResponseSimplePublicUpsertObjectWithErrors) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


