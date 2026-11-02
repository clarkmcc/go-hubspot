# BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | **time.Time** | The date and time when the batch operation was completed, in ISO 8601 format. | 
**Errors** | Pointer to [**[]StandardError**](StandardError.md) | An array of errors encountered during the batch operation, each represented by a StandardError object. | [optional] 
**Links** | Pointer to **map[string]string** | An object containing key-value pairs of link names to their associated URIs, providing additional resources or documentation. | [optional] 
**NumErrors** | Pointer to **int32** | The total number of errors encountered during the batch operation. | [optional] 
**RequestedAt** | Pointer to **time.Time** | The date and time when the batch operation was requested, in ISO 8601 format. | [optional] 
**Results** | [**[]PublicAssociationDefinitionUserConfiguration**](PublicAssociationDefinitionUserConfiguration.md) | An array of results from the batch operation, each representing a user configuration of a public association definition. | 
**StartedAt** | **time.Time** | The date and time when the batch operation started, in ISO 8601 format. | 
**Status** | **string** | The current status of the batch operation. Valid values include &#39;PENDING&#39;, &#39;PROCESSING&#39;, &#39;CANCELED&#39;, and &#39;COMPLETE&#39;. | 

## Methods

### NewBatchResponsePublicAssociationDefinitionUserConfigurationWithErrors

`func NewBatchResponsePublicAssociationDefinitionUserConfigurationWithErrors(completedAt time.Time, results []PublicAssociationDefinitionUserConfiguration, startedAt time.Time, status string, ) *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors`

NewBatchResponsePublicAssociationDefinitionUserConfigurationWithErrors instantiates a new BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResponsePublicAssociationDefinitionUserConfigurationWithErrorsWithDefaults

`func NewBatchResponsePublicAssociationDefinitionUserConfigurationWithErrorsWithDefaults() *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors`

NewBatchResponsePublicAssociationDefinitionUserConfigurationWithErrorsWithDefaults instantiates a new BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetErrors

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) GetErrors() []StandardError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) GetErrorsOk() (*[]StandardError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) SetErrors(v []StandardError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetLinks

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetNumErrors

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) GetNumErrors() int32`

GetNumErrors returns the NumErrors field if non-nil, zero value otherwise.

### GetNumErrorsOk

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) GetNumErrorsOk() (*int32, bool)`

GetNumErrorsOk returns a tuple with the NumErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumErrors

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) SetNumErrors(v int32)`

SetNumErrors sets NumErrors field to given value.

### HasNumErrors

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) HasNumErrors() bool`

HasNumErrors returns a boolean if a field has been set.

### GetRequestedAt

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetResults

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) GetResults() []PublicAssociationDefinitionUserConfiguration`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) GetResultsOk() (*[]PublicAssociationDefinitionUserConfiguration, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) SetResults(v []PublicAssociationDefinitionUserConfiguration)`

SetResults sets Results field to given value.


### GetStartedAt

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetStatus

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchResponsePublicAssociationDefinitionUserConfigurationWithErrors) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


