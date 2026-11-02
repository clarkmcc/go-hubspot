# BatchResponseSimplePublicUpsertObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | **time.Time** | The timestamp when the batch process was completed, in ISO 8601 format. | 
**Links** | Pointer to **map[string]string** | An object containing relevant links related to the batch request. | [optional] 
**RequestedAt** | Pointer to **time.Time** | The timestamp when the batch process was initiated, in ISO 8601 format. | [optional] 
**Results** | [**[]SimplePublicUpsertObject**](SimplePublicUpsertObject.md) | An array of objects representing the results of the batch upsert operation. Each object includes details about the created or updated CRM object. | 
**StartedAt** | **time.Time** | The timestamp when the batch process began execution, in ISO 8601 format. | 
**Status** | **string** | The status of the batch processing request. The expected value is \&quot;COMPLETE\&quot;. | 

## Methods

### NewBatchResponseSimplePublicUpsertObject

`func NewBatchResponseSimplePublicUpsertObject(completedAt time.Time, results []SimplePublicUpsertObject, startedAt time.Time, status string, ) *BatchResponseSimplePublicUpsertObject`

NewBatchResponseSimplePublicUpsertObject instantiates a new BatchResponseSimplePublicUpsertObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResponseSimplePublicUpsertObjectWithDefaults

`func NewBatchResponseSimplePublicUpsertObjectWithDefaults() *BatchResponseSimplePublicUpsertObject`

NewBatchResponseSimplePublicUpsertObjectWithDefaults instantiates a new BatchResponseSimplePublicUpsertObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *BatchResponseSimplePublicUpsertObject) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BatchResponseSimplePublicUpsertObject) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BatchResponseSimplePublicUpsertObject) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetLinks

`func (o *BatchResponseSimplePublicUpsertObject) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BatchResponseSimplePublicUpsertObject) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BatchResponseSimplePublicUpsertObject) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BatchResponseSimplePublicUpsertObject) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRequestedAt

`func (o *BatchResponseSimplePublicUpsertObject) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *BatchResponseSimplePublicUpsertObject) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *BatchResponseSimplePublicUpsertObject) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *BatchResponseSimplePublicUpsertObject) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetResults

`func (o *BatchResponseSimplePublicUpsertObject) GetResults() []SimplePublicUpsertObject`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchResponseSimplePublicUpsertObject) GetResultsOk() (*[]SimplePublicUpsertObject, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchResponseSimplePublicUpsertObject) SetResults(v []SimplePublicUpsertObject)`

SetResults sets Results field to given value.


### GetStartedAt

`func (o *BatchResponseSimplePublicUpsertObject) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BatchResponseSimplePublicUpsertObject) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BatchResponseSimplePublicUpsertObject) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetStatus

`func (o *BatchResponseSimplePublicUpsertObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchResponseSimplePublicUpsertObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchResponseSimplePublicUpsertObject) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


