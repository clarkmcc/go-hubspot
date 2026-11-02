# BatchResponsePublicAssociationDefinitionUserConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | **time.Time** | The date and time when the batch operation was completed, in ISO 8601 format. | 
**Links** | Pointer to **map[string]string** | A map of link names to associated URIs, providing additional information related to the batch operation. | [optional] 
**RequestedAt** | Pointer to **time.Time** | The date and time when the batch operation was requested, in ISO 8601 format. | [optional] 
**Results** | [**[]PublicAssociationDefinitionUserConfiguration**](PublicAssociationDefinitionUserConfiguration.md) | An array containing the results of the batch operation, each item representing a public association definition user configuration. | 
**StartedAt** | **time.Time** | The date and time when the batch operation started, in ISO 8601 format. | 
**Status** | **string** | The current status of the batch operation. Valid values include &#39;PENDING&#39;, &#39;PROCESSING&#39;, &#39;CANCELED&#39;, and &#39;COMPLETE&#39;. | 

## Methods

### NewBatchResponsePublicAssociationDefinitionUserConfiguration

`func NewBatchResponsePublicAssociationDefinitionUserConfiguration(completedAt time.Time, results []PublicAssociationDefinitionUserConfiguration, startedAt time.Time, status string, ) *BatchResponsePublicAssociationDefinitionUserConfiguration`

NewBatchResponsePublicAssociationDefinitionUserConfiguration instantiates a new BatchResponsePublicAssociationDefinitionUserConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResponsePublicAssociationDefinitionUserConfigurationWithDefaults

`func NewBatchResponsePublicAssociationDefinitionUserConfigurationWithDefaults() *BatchResponsePublicAssociationDefinitionUserConfiguration`

NewBatchResponsePublicAssociationDefinitionUserConfigurationWithDefaults instantiates a new BatchResponsePublicAssociationDefinitionUserConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *BatchResponsePublicAssociationDefinitionUserConfiguration) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BatchResponsePublicAssociationDefinitionUserConfiguration) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BatchResponsePublicAssociationDefinitionUserConfiguration) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetLinks

`func (o *BatchResponsePublicAssociationDefinitionUserConfiguration) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BatchResponsePublicAssociationDefinitionUserConfiguration) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BatchResponsePublicAssociationDefinitionUserConfiguration) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BatchResponsePublicAssociationDefinitionUserConfiguration) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRequestedAt

`func (o *BatchResponsePublicAssociationDefinitionUserConfiguration) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *BatchResponsePublicAssociationDefinitionUserConfiguration) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *BatchResponsePublicAssociationDefinitionUserConfiguration) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *BatchResponsePublicAssociationDefinitionUserConfiguration) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetResults

`func (o *BatchResponsePublicAssociationDefinitionUserConfiguration) GetResults() []PublicAssociationDefinitionUserConfiguration`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchResponsePublicAssociationDefinitionUserConfiguration) GetResultsOk() (*[]PublicAssociationDefinitionUserConfiguration, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchResponsePublicAssociationDefinitionUserConfiguration) SetResults(v []PublicAssociationDefinitionUserConfiguration)`

SetResults sets Results field to given value.


### GetStartedAt

`func (o *BatchResponsePublicAssociationDefinitionUserConfiguration) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BatchResponsePublicAssociationDefinitionUserConfiguration) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BatchResponsePublicAssociationDefinitionUserConfiguration) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetStatus

`func (o *BatchResponsePublicAssociationDefinitionUserConfiguration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchResponsePublicAssociationDefinitionUserConfiguration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchResponsePublicAssociationDefinitionUserConfiguration) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


