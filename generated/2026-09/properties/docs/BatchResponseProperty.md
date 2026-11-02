# BatchResponseProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | **time.Time** | The timestamp indicating when the batch operation was completed. | 
**Links** | Pointer to **map[string]string** | A collection of URLs linking to documentation or resources related to the batch operation. | [optional] 
**RequestedAt** | Pointer to **time.Time** | The timestamp indicating when the batch operation was requested. | [optional] 
**Results** | [**[]Property**](Property.md) | An array of Property objects representing the results of the batch operation. | 
**StartedAt** | **time.Time** | The timestamp indicating when the batch operation began processing. | 
**Status** | **string** | The current status of the batch operation, with possible values being CANCELED, COMPLETE, PENDING, or PROCESSING. | 

## Methods

### NewBatchResponseProperty

`func NewBatchResponseProperty(completedAt time.Time, results []Property, startedAt time.Time, status string, ) *BatchResponseProperty`

NewBatchResponseProperty instantiates a new BatchResponseProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResponsePropertyWithDefaults

`func NewBatchResponsePropertyWithDefaults() *BatchResponseProperty`

NewBatchResponsePropertyWithDefaults instantiates a new BatchResponseProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *BatchResponseProperty) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BatchResponseProperty) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BatchResponseProperty) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetLinks

`func (o *BatchResponseProperty) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BatchResponseProperty) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BatchResponseProperty) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BatchResponseProperty) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRequestedAt

`func (o *BatchResponseProperty) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *BatchResponseProperty) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *BatchResponseProperty) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *BatchResponseProperty) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetResults

`func (o *BatchResponseProperty) GetResults() []Property`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchResponseProperty) GetResultsOk() (*[]Property, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchResponseProperty) SetResults(v []Property)`

SetResults sets Results field to given value.


### GetStartedAt

`func (o *BatchResponseProperty) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BatchResponseProperty) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BatchResponseProperty) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetStatus

`func (o *BatchResponseProperty) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchResponseProperty) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchResponseProperty) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


