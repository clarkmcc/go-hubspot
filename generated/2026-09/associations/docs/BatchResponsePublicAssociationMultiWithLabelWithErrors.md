# BatchResponsePublicAssociationMultiWithLabelWithErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | **time.Time** | The date and time when the batch operation was completed, in ISO 8601 format. | 
**Errors** | Pointer to [**[]StandardError**](StandardError.md) | An array of errors that occurred during the batch operation, with each error represented by a standard error object. | [optional] 
**Links** | Pointer to **map[string]string** | An object containing additional links related to the batch operation, with each link represented as a string. | [optional] 
**NumErrors** | Pointer to **int32** | The number of errors that occurred during the batch operation, represented as an integer. | [optional] 
**RequestedAt** | Pointer to **time.Time** | The date and time when the batch operation was requested, in ISO 8601 format. | [optional] 
**Results** | [**[]PublicAssociationMultiWithLabel**](PublicAssociationMultiWithLabel.md) | An array of results from the batch operation, each item representing a public association with labels. | 
**StartedAt** | **time.Time** | The date and time when the batch operation started, in ISO 8601 format. | 
**Status** | **string** | The current status of the batch operation. Valid values include &#39;PENDING&#39;, &#39;PROCESSING&#39;, &#39;CANCELED&#39;, and &#39;COMPLETE&#39;. | 

## Methods

### NewBatchResponsePublicAssociationMultiWithLabelWithErrors

`func NewBatchResponsePublicAssociationMultiWithLabelWithErrors(completedAt time.Time, results []PublicAssociationMultiWithLabel, startedAt time.Time, status string, ) *BatchResponsePublicAssociationMultiWithLabelWithErrors`

NewBatchResponsePublicAssociationMultiWithLabelWithErrors instantiates a new BatchResponsePublicAssociationMultiWithLabelWithErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResponsePublicAssociationMultiWithLabelWithErrorsWithDefaults

`func NewBatchResponsePublicAssociationMultiWithLabelWithErrorsWithDefaults() *BatchResponsePublicAssociationMultiWithLabelWithErrors`

NewBatchResponsePublicAssociationMultiWithLabelWithErrorsWithDefaults instantiates a new BatchResponsePublicAssociationMultiWithLabelWithErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetErrors

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) GetErrors() []StandardError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) GetErrorsOk() (*[]StandardError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) SetErrors(v []StandardError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetLinks

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetNumErrors

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) GetNumErrors() int32`

GetNumErrors returns the NumErrors field if non-nil, zero value otherwise.

### GetNumErrorsOk

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) GetNumErrorsOk() (*int32, bool)`

GetNumErrorsOk returns a tuple with the NumErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumErrors

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) SetNumErrors(v int32)`

SetNumErrors sets NumErrors field to given value.

### HasNumErrors

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) HasNumErrors() bool`

HasNumErrors returns a boolean if a field has been set.

### GetRequestedAt

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetResults

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) GetResults() []PublicAssociationMultiWithLabel`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) GetResultsOk() (*[]PublicAssociationMultiWithLabel, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) SetResults(v []PublicAssociationMultiWithLabel)`

SetResults sets Results field to given value.


### GetStartedAt

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetStatus

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchResponsePublicAssociationMultiWithLabelWithErrors) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


