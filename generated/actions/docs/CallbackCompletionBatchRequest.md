# CallbackCompletionBatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackId** | **string** | The ID of the callback to complete. | 
**OutputFields** | **map[string]string** | A map of action output names and values. | 

## Methods

### NewCallbackCompletionBatchRequest

`func NewCallbackCompletionBatchRequest(callbackId string, outputFields map[string]string, ) *CallbackCompletionBatchRequest`

NewCallbackCompletionBatchRequest instantiates a new CallbackCompletionBatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallbackCompletionBatchRequestWithDefaults

`func NewCallbackCompletionBatchRequestWithDefaults() *CallbackCompletionBatchRequest`

NewCallbackCompletionBatchRequestWithDefaults instantiates a new CallbackCompletionBatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackId

`func (o *CallbackCompletionBatchRequest) GetCallbackId() string`

GetCallbackId returns the CallbackId field if non-nil, zero value otherwise.

### GetCallbackIdOk

`func (o *CallbackCompletionBatchRequest) GetCallbackIdOk() (*string, bool)`

GetCallbackIdOk returns a tuple with the CallbackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackId

`func (o *CallbackCompletionBatchRequest) SetCallbackId(v string)`

SetCallbackId sets CallbackId field to given value.


### GetOutputFields

`func (o *CallbackCompletionBatchRequest) GetOutputFields() map[string]string`

GetOutputFields returns the OutputFields field if non-nil, zero value otherwise.

### GetOutputFieldsOk

`func (o *CallbackCompletionBatchRequest) GetOutputFieldsOk() (*map[string]string, bool)`

GetOutputFieldsOk returns a tuple with the OutputFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFields

`func (o *CallbackCompletionBatchRequest) SetOutputFields(v map[string]string)`

SetOutputFields sets OutputFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


