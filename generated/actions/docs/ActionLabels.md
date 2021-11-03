# ActionLabels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputFieldLabels** | Pointer to **map[string]string** | A map of input field names to the user-facing labels. | [optional] 
**InputFieldDescriptions** | Pointer to **map[string]string** | A map of input field names to descriptions for the fields. These will show up as tooltips when users are editing your action. | [optional] 
**ActionName** | **string** | The name of this custom action. This is what will show up when users are selecting an action in the workflows app. | 
**ActionDescription** | Pointer to **string** | A description for this custom action. This will show up in the action editor along with the input fields. | [optional] 
**AppDisplayName** | Pointer to **string** | The name to be displayed at the top of the action editor in the workflows app. | [optional] 
**ActionCardContent** | Pointer to **string** | The label to be displayed in the action card of the workflow editor once this custom action has been added to a workflow. | [optional] 

## Methods

### NewActionLabels

`func NewActionLabels(actionName string, ) *ActionLabels`

NewActionLabels instantiates a new ActionLabels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionLabelsWithDefaults

`func NewActionLabelsWithDefaults() *ActionLabels`

NewActionLabelsWithDefaults instantiates a new ActionLabels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputFieldLabels

`func (o *ActionLabels) GetInputFieldLabels() map[string]string`

GetInputFieldLabels returns the InputFieldLabels field if non-nil, zero value otherwise.

### GetInputFieldLabelsOk

`func (o *ActionLabels) GetInputFieldLabelsOk() (*map[string]string, bool)`

GetInputFieldLabelsOk returns a tuple with the InputFieldLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFieldLabels

`func (o *ActionLabels) SetInputFieldLabels(v map[string]string)`

SetInputFieldLabels sets InputFieldLabels field to given value.

### HasInputFieldLabels

`func (o *ActionLabels) HasInputFieldLabels() bool`

HasInputFieldLabels returns a boolean if a field has been set.

### GetInputFieldDescriptions

`func (o *ActionLabels) GetInputFieldDescriptions() map[string]string`

GetInputFieldDescriptions returns the InputFieldDescriptions field if non-nil, zero value otherwise.

### GetInputFieldDescriptionsOk

`func (o *ActionLabels) GetInputFieldDescriptionsOk() (*map[string]string, bool)`

GetInputFieldDescriptionsOk returns a tuple with the InputFieldDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFieldDescriptions

`func (o *ActionLabels) SetInputFieldDescriptions(v map[string]string)`

SetInputFieldDescriptions sets InputFieldDescriptions field to given value.

### HasInputFieldDescriptions

`func (o *ActionLabels) HasInputFieldDescriptions() bool`

HasInputFieldDescriptions returns a boolean if a field has been set.

### GetActionName

`func (o *ActionLabels) GetActionName() string`

GetActionName returns the ActionName field if non-nil, zero value otherwise.

### GetActionNameOk

`func (o *ActionLabels) GetActionNameOk() (*string, bool)`

GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionName

`func (o *ActionLabels) SetActionName(v string)`

SetActionName sets ActionName field to given value.


### GetActionDescription

`func (o *ActionLabels) GetActionDescription() string`

GetActionDescription returns the ActionDescription field if non-nil, zero value otherwise.

### GetActionDescriptionOk

`func (o *ActionLabels) GetActionDescriptionOk() (*string, bool)`

GetActionDescriptionOk returns a tuple with the ActionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDescription

`func (o *ActionLabels) SetActionDescription(v string)`

SetActionDescription sets ActionDescription field to given value.

### HasActionDescription

`func (o *ActionLabels) HasActionDescription() bool`

HasActionDescription returns a boolean if a field has been set.

### GetAppDisplayName

`func (o *ActionLabels) GetAppDisplayName() string`

GetAppDisplayName returns the AppDisplayName field if non-nil, zero value otherwise.

### GetAppDisplayNameOk

`func (o *ActionLabels) GetAppDisplayNameOk() (*string, bool)`

GetAppDisplayNameOk returns a tuple with the AppDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDisplayName

`func (o *ActionLabels) SetAppDisplayName(v string)`

SetAppDisplayName sets AppDisplayName field to given value.

### HasAppDisplayName

`func (o *ActionLabels) HasAppDisplayName() bool`

HasAppDisplayName returns a boolean if a field has been set.

### GetActionCardContent

`func (o *ActionLabels) GetActionCardContent() string`

GetActionCardContent returns the ActionCardContent field if non-nil, zero value otherwise.

### GetActionCardContentOk

`func (o *ActionLabels) GetActionCardContentOk() (*string, bool)`

GetActionCardContentOk returns a tuple with the ActionCardContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionCardContent

`func (o *ActionLabels) SetActionCardContent(v string)`

SetActionCardContent sets ActionCardContent field to given value.

### HasActionCardContent

`func (o *ActionLabels) HasActionCardContent() bool`

HasActionCardContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


