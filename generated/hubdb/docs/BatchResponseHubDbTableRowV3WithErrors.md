# BatchResponseHubDbTableRowV3WithErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Results** | [**[]HubDbTableRowV3**](HubDbTableRowV3.md) |  | 
**NumErrors** | Pointer to **int32** |  | [optional] 
**Errors** | Pointer to [**[]StandardError**](StandardError.md) |  | [optional] 
**RequestedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | **time.Time** |  | 
**CompletedAt** | **time.Time** |  | 
**Links** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBatchResponseHubDbTableRowV3WithErrors

`func NewBatchResponseHubDbTableRowV3WithErrors(status string, results []HubDbTableRowV3, startedAt time.Time, completedAt time.Time, ) *BatchResponseHubDbTableRowV3WithErrors`

NewBatchResponseHubDbTableRowV3WithErrors instantiates a new BatchResponseHubDbTableRowV3WithErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResponseHubDbTableRowV3WithErrorsWithDefaults

`func NewBatchResponseHubDbTableRowV3WithErrorsWithDefaults() *BatchResponseHubDbTableRowV3WithErrors`

NewBatchResponseHubDbTableRowV3WithErrorsWithDefaults instantiates a new BatchResponseHubDbTableRowV3WithErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BatchResponseHubDbTableRowV3WithErrors) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchResponseHubDbTableRowV3WithErrors) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchResponseHubDbTableRowV3WithErrors) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetResults

`func (o *BatchResponseHubDbTableRowV3WithErrors) GetResults() []HubDbTableRowV3`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchResponseHubDbTableRowV3WithErrors) GetResultsOk() (*[]HubDbTableRowV3, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchResponseHubDbTableRowV3WithErrors) SetResults(v []HubDbTableRowV3)`

SetResults sets Results field to given value.


### GetNumErrors

`func (o *BatchResponseHubDbTableRowV3WithErrors) GetNumErrors() int32`

GetNumErrors returns the NumErrors field if non-nil, zero value otherwise.

### GetNumErrorsOk

`func (o *BatchResponseHubDbTableRowV3WithErrors) GetNumErrorsOk() (*int32, bool)`

GetNumErrorsOk returns a tuple with the NumErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumErrors

`func (o *BatchResponseHubDbTableRowV3WithErrors) SetNumErrors(v int32)`

SetNumErrors sets NumErrors field to given value.

### HasNumErrors

`func (o *BatchResponseHubDbTableRowV3WithErrors) HasNumErrors() bool`

HasNumErrors returns a boolean if a field has been set.

### GetErrors

`func (o *BatchResponseHubDbTableRowV3WithErrors) GetErrors() []StandardError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BatchResponseHubDbTableRowV3WithErrors) GetErrorsOk() (*[]StandardError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BatchResponseHubDbTableRowV3WithErrors) SetErrors(v []StandardError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BatchResponseHubDbTableRowV3WithErrors) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetRequestedAt

`func (o *BatchResponseHubDbTableRowV3WithErrors) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *BatchResponseHubDbTableRowV3WithErrors) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *BatchResponseHubDbTableRowV3WithErrors) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *BatchResponseHubDbTableRowV3WithErrors) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *BatchResponseHubDbTableRowV3WithErrors) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BatchResponseHubDbTableRowV3WithErrors) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BatchResponseHubDbTableRowV3WithErrors) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetCompletedAt

`func (o *BatchResponseHubDbTableRowV3WithErrors) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BatchResponseHubDbTableRowV3WithErrors) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BatchResponseHubDbTableRowV3WithErrors) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetLinks

`func (o *BatchResponseHubDbTableRowV3WithErrors) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BatchResponseHubDbTableRowV3WithErrors) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BatchResponseHubDbTableRowV3WithErrors) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BatchResponseHubDbTableRowV3WithErrors) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


