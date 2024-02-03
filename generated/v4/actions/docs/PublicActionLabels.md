# PublicActionLabels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputFieldDescriptions** | Pointer to **map[string]string** |  | [optional] 
**AppDisplayName** | Pointer to **string** |  | [optional] 
**OutputFieldLabels** | Pointer to **map[string]string** |  | [optional] 
**InputFieldOptionLabels** | Pointer to **map[string]map[string]string** |  | [optional] 
**ActionDescription** | Pointer to **string** |  | [optional] 
**ExecutionRules** | Pointer to **map[string]string** |  | [optional] 
**InputFieldLabels** | Pointer to **map[string]string** |  | [optional] 
**ActionName** | **string** |  | 
**ActionCardContent** | Pointer to **string** |  | [optional] 

## Methods

### NewPublicActionLabels

`func NewPublicActionLabels(actionName string, ) *PublicActionLabels`

NewPublicActionLabels instantiates a new PublicActionLabels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicActionLabelsWithDefaults

`func NewPublicActionLabelsWithDefaults() *PublicActionLabels`

NewPublicActionLabelsWithDefaults instantiates a new PublicActionLabels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputFieldDescriptions

`func (o *PublicActionLabels) GetInputFieldDescriptions() map[string]string`

GetInputFieldDescriptions returns the InputFieldDescriptions field if non-nil, zero value otherwise.

### GetInputFieldDescriptionsOk

`func (o *PublicActionLabels) GetInputFieldDescriptionsOk() (*map[string]string, bool)`

GetInputFieldDescriptionsOk returns a tuple with the InputFieldDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFieldDescriptions

`func (o *PublicActionLabels) SetInputFieldDescriptions(v map[string]string)`

SetInputFieldDescriptions sets InputFieldDescriptions field to given value.

### HasInputFieldDescriptions

`func (o *PublicActionLabels) HasInputFieldDescriptions() bool`

HasInputFieldDescriptions returns a boolean if a field has been set.

### GetAppDisplayName

`func (o *PublicActionLabels) GetAppDisplayName() string`

GetAppDisplayName returns the AppDisplayName field if non-nil, zero value otherwise.

### GetAppDisplayNameOk

`func (o *PublicActionLabels) GetAppDisplayNameOk() (*string, bool)`

GetAppDisplayNameOk returns a tuple with the AppDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDisplayName

`func (o *PublicActionLabels) SetAppDisplayName(v string)`

SetAppDisplayName sets AppDisplayName field to given value.

### HasAppDisplayName

`func (o *PublicActionLabels) HasAppDisplayName() bool`

HasAppDisplayName returns a boolean if a field has been set.

### GetOutputFieldLabels

`func (o *PublicActionLabels) GetOutputFieldLabels() map[string]string`

GetOutputFieldLabels returns the OutputFieldLabels field if non-nil, zero value otherwise.

### GetOutputFieldLabelsOk

`func (o *PublicActionLabels) GetOutputFieldLabelsOk() (*map[string]string, bool)`

GetOutputFieldLabelsOk returns a tuple with the OutputFieldLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFieldLabels

`func (o *PublicActionLabels) SetOutputFieldLabels(v map[string]string)`

SetOutputFieldLabels sets OutputFieldLabels field to given value.

### HasOutputFieldLabels

`func (o *PublicActionLabels) HasOutputFieldLabels() bool`

HasOutputFieldLabels returns a boolean if a field has been set.

### GetInputFieldOptionLabels

`func (o *PublicActionLabels) GetInputFieldOptionLabels() map[string]map[string]string`

GetInputFieldOptionLabels returns the InputFieldOptionLabels field if non-nil, zero value otherwise.

### GetInputFieldOptionLabelsOk

`func (o *PublicActionLabels) GetInputFieldOptionLabelsOk() (*map[string]map[string]string, bool)`

GetInputFieldOptionLabelsOk returns a tuple with the InputFieldOptionLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFieldOptionLabels

`func (o *PublicActionLabels) SetInputFieldOptionLabels(v map[string]map[string]string)`

SetInputFieldOptionLabels sets InputFieldOptionLabels field to given value.

### HasInputFieldOptionLabels

`func (o *PublicActionLabels) HasInputFieldOptionLabels() bool`

HasInputFieldOptionLabels returns a boolean if a field has been set.

### GetActionDescription

`func (o *PublicActionLabels) GetActionDescription() string`

GetActionDescription returns the ActionDescription field if non-nil, zero value otherwise.

### GetActionDescriptionOk

`func (o *PublicActionLabels) GetActionDescriptionOk() (*string, bool)`

GetActionDescriptionOk returns a tuple with the ActionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDescription

`func (o *PublicActionLabels) SetActionDescription(v string)`

SetActionDescription sets ActionDescription field to given value.

### HasActionDescription

`func (o *PublicActionLabels) HasActionDescription() bool`

HasActionDescription returns a boolean if a field has been set.

### GetExecutionRules

`func (o *PublicActionLabels) GetExecutionRules() map[string]string`

GetExecutionRules returns the ExecutionRules field if non-nil, zero value otherwise.

### GetExecutionRulesOk

`func (o *PublicActionLabels) GetExecutionRulesOk() (*map[string]string, bool)`

GetExecutionRulesOk returns a tuple with the ExecutionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionRules

`func (o *PublicActionLabels) SetExecutionRules(v map[string]string)`

SetExecutionRules sets ExecutionRules field to given value.

### HasExecutionRules

`func (o *PublicActionLabels) HasExecutionRules() bool`

HasExecutionRules returns a boolean if a field has been set.

### GetInputFieldLabels

`func (o *PublicActionLabels) GetInputFieldLabels() map[string]string`

GetInputFieldLabels returns the InputFieldLabels field if non-nil, zero value otherwise.

### GetInputFieldLabelsOk

`func (o *PublicActionLabels) GetInputFieldLabelsOk() (*map[string]string, bool)`

GetInputFieldLabelsOk returns a tuple with the InputFieldLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFieldLabels

`func (o *PublicActionLabels) SetInputFieldLabels(v map[string]string)`

SetInputFieldLabels sets InputFieldLabels field to given value.

### HasInputFieldLabels

`func (o *PublicActionLabels) HasInputFieldLabels() bool`

HasInputFieldLabels returns a boolean if a field has been set.

### GetActionName

`func (o *PublicActionLabels) GetActionName() string`

GetActionName returns the ActionName field if non-nil, zero value otherwise.

### GetActionNameOk

`func (o *PublicActionLabels) GetActionNameOk() (*string, bool)`

GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionName

`func (o *PublicActionLabels) SetActionName(v string)`

SetActionName sets ActionName field to given value.


### GetActionCardContent

`func (o *PublicActionLabels) GetActionCardContent() string`

GetActionCardContent returns the ActionCardContent field if non-nil, zero value otherwise.

### GetActionCardContentOk

`func (o *PublicActionLabels) GetActionCardContentOk() (*string, bool)`

GetActionCardContentOk returns a tuple with the ActionCardContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionCardContent

`func (o *PublicActionLabels) SetActionCardContent(v string)`

SetActionCardContent sets ActionCardContent field to given value.

### HasActionCardContent

`func (o *PublicActionLabels) HasActionCardContent() bool`

HasActionCardContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


