# BatchResponseHubDbTableRowV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**[]HubDbTableRowV3**](HubDbTableRowV3.md) |  | [optional] 
**RequestedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**Links** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBatchResponseHubDbTableRowV3

`func NewBatchResponseHubDbTableRowV3() *BatchResponseHubDbTableRowV3`

NewBatchResponseHubDbTableRowV3 instantiates a new BatchResponseHubDbTableRowV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResponseHubDbTableRowV3WithDefaults

`func NewBatchResponseHubDbTableRowV3WithDefaults() *BatchResponseHubDbTableRowV3`

NewBatchResponseHubDbTableRowV3WithDefaults instantiates a new BatchResponseHubDbTableRowV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BatchResponseHubDbTableRowV3) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchResponseHubDbTableRowV3) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchResponseHubDbTableRowV3) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BatchResponseHubDbTableRowV3) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResults

`func (o *BatchResponseHubDbTableRowV3) GetResults() []HubDbTableRowV3`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchResponseHubDbTableRowV3) GetResultsOk() (*[]HubDbTableRowV3, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchResponseHubDbTableRowV3) SetResults(v []HubDbTableRowV3)`

SetResults sets Results field to given value.

### HasResults

`func (o *BatchResponseHubDbTableRowV3) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetRequestedAt

`func (o *BatchResponseHubDbTableRowV3) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *BatchResponseHubDbTableRowV3) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *BatchResponseHubDbTableRowV3) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *BatchResponseHubDbTableRowV3) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *BatchResponseHubDbTableRowV3) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BatchResponseHubDbTableRowV3) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BatchResponseHubDbTableRowV3) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *BatchResponseHubDbTableRowV3) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *BatchResponseHubDbTableRowV3) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BatchResponseHubDbTableRowV3) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BatchResponseHubDbTableRowV3) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *BatchResponseHubDbTableRowV3) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetLinks

`func (o *BatchResponseHubDbTableRowV3) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BatchResponseHubDbTableRowV3) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BatchResponseHubDbTableRowV3) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BatchResponseHubDbTableRowV3) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


