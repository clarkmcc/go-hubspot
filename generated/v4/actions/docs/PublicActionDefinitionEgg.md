# PublicActionDefinitionEgg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputFields** | [**[]InputFieldDefinition**](InputFieldDefinition.md) |  | 
**OutputFields** | Pointer to [**[]OutputFieldDefinition**](OutputFieldDefinition.md) |  | [optional] 
**ArchivedAt** | Pointer to **int64** |  | [optional] 
**Functions** | [**[]PublicActionFunction**](PublicActionFunction.md) |  | 
**ActionUrl** | **string** |  | 
**InputFieldDependencies** | Pointer to [**[]PublicActionDefinitionInputFieldDependenciesInner**](PublicActionDefinitionInputFieldDependenciesInner.md) |  | [optional] 
**Published** | **bool** |  | 
**ExecutionRules** | Pointer to [**[]PublicExecutionTranslationRule**](PublicExecutionTranslationRule.md) |  | [optional] 
**ObjectTypes** | **[]string** |  | 
**ObjectRequestOptions** | Pointer to [**PublicObjectRequestOptions**](PublicObjectRequestOptions.md) |  | [optional] 
**Labels** | [**map[string]PublicActionLabels**](PublicActionLabels.md) |  | 

## Methods

### NewPublicActionDefinitionEgg

`func NewPublicActionDefinitionEgg(inputFields []InputFieldDefinition, functions []PublicActionFunction, actionUrl string, published bool, objectTypes []string, labels map[string]PublicActionLabels, ) *PublicActionDefinitionEgg`

NewPublicActionDefinitionEgg instantiates a new PublicActionDefinitionEgg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicActionDefinitionEggWithDefaults

`func NewPublicActionDefinitionEggWithDefaults() *PublicActionDefinitionEgg`

NewPublicActionDefinitionEggWithDefaults instantiates a new PublicActionDefinitionEgg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputFields

`func (o *PublicActionDefinitionEgg) GetInputFields() []InputFieldDefinition`

GetInputFields returns the InputFields field if non-nil, zero value otherwise.

### GetInputFieldsOk

`func (o *PublicActionDefinitionEgg) GetInputFieldsOk() (*[]InputFieldDefinition, bool)`

GetInputFieldsOk returns a tuple with the InputFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFields

`func (o *PublicActionDefinitionEgg) SetInputFields(v []InputFieldDefinition)`

SetInputFields sets InputFields field to given value.


### GetOutputFields

`func (o *PublicActionDefinitionEgg) GetOutputFields() []OutputFieldDefinition`

GetOutputFields returns the OutputFields field if non-nil, zero value otherwise.

### GetOutputFieldsOk

`func (o *PublicActionDefinitionEgg) GetOutputFieldsOk() (*[]OutputFieldDefinition, bool)`

GetOutputFieldsOk returns a tuple with the OutputFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFields

`func (o *PublicActionDefinitionEgg) SetOutputFields(v []OutputFieldDefinition)`

SetOutputFields sets OutputFields field to given value.

### HasOutputFields

`func (o *PublicActionDefinitionEgg) HasOutputFields() bool`

HasOutputFields returns a boolean if a field has been set.

### GetArchivedAt

`func (o *PublicActionDefinitionEgg) GetArchivedAt() int64`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *PublicActionDefinitionEgg) GetArchivedAtOk() (*int64, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *PublicActionDefinitionEgg) SetArchivedAt(v int64)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *PublicActionDefinitionEgg) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetFunctions

`func (o *PublicActionDefinitionEgg) GetFunctions() []PublicActionFunction`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *PublicActionDefinitionEgg) GetFunctionsOk() (*[]PublicActionFunction, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *PublicActionDefinitionEgg) SetFunctions(v []PublicActionFunction)`

SetFunctions sets Functions field to given value.


### GetActionUrl

`func (o *PublicActionDefinitionEgg) GetActionUrl() string`

GetActionUrl returns the ActionUrl field if non-nil, zero value otherwise.

### GetActionUrlOk

`func (o *PublicActionDefinitionEgg) GetActionUrlOk() (*string, bool)`

GetActionUrlOk returns a tuple with the ActionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionUrl

`func (o *PublicActionDefinitionEgg) SetActionUrl(v string)`

SetActionUrl sets ActionUrl field to given value.


### GetInputFieldDependencies

`func (o *PublicActionDefinitionEgg) GetInputFieldDependencies() []PublicActionDefinitionInputFieldDependenciesInner`

GetInputFieldDependencies returns the InputFieldDependencies field if non-nil, zero value otherwise.

### GetInputFieldDependenciesOk

`func (o *PublicActionDefinitionEgg) GetInputFieldDependenciesOk() (*[]PublicActionDefinitionInputFieldDependenciesInner, bool)`

GetInputFieldDependenciesOk returns a tuple with the InputFieldDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFieldDependencies

`func (o *PublicActionDefinitionEgg) SetInputFieldDependencies(v []PublicActionDefinitionInputFieldDependenciesInner)`

SetInputFieldDependencies sets InputFieldDependencies field to given value.

### HasInputFieldDependencies

`func (o *PublicActionDefinitionEgg) HasInputFieldDependencies() bool`

HasInputFieldDependencies returns a boolean if a field has been set.

### GetPublished

`func (o *PublicActionDefinitionEgg) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *PublicActionDefinitionEgg) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *PublicActionDefinitionEgg) SetPublished(v bool)`

SetPublished sets Published field to given value.


### GetExecutionRules

`func (o *PublicActionDefinitionEgg) GetExecutionRules() []PublicExecutionTranslationRule`

GetExecutionRules returns the ExecutionRules field if non-nil, zero value otherwise.

### GetExecutionRulesOk

`func (o *PublicActionDefinitionEgg) GetExecutionRulesOk() (*[]PublicExecutionTranslationRule, bool)`

GetExecutionRulesOk returns a tuple with the ExecutionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionRules

`func (o *PublicActionDefinitionEgg) SetExecutionRules(v []PublicExecutionTranslationRule)`

SetExecutionRules sets ExecutionRules field to given value.

### HasExecutionRules

`func (o *PublicActionDefinitionEgg) HasExecutionRules() bool`

HasExecutionRules returns a boolean if a field has been set.

### GetObjectTypes

`func (o *PublicActionDefinitionEgg) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *PublicActionDefinitionEgg) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *PublicActionDefinitionEgg) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.


### GetObjectRequestOptions

`func (o *PublicActionDefinitionEgg) GetObjectRequestOptions() PublicObjectRequestOptions`

GetObjectRequestOptions returns the ObjectRequestOptions field if non-nil, zero value otherwise.

### GetObjectRequestOptionsOk

`func (o *PublicActionDefinitionEgg) GetObjectRequestOptionsOk() (*PublicObjectRequestOptions, bool)`

GetObjectRequestOptionsOk returns a tuple with the ObjectRequestOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectRequestOptions

`func (o *PublicActionDefinitionEgg) SetObjectRequestOptions(v PublicObjectRequestOptions)`

SetObjectRequestOptions sets ObjectRequestOptions field to given value.

### HasObjectRequestOptions

`func (o *PublicActionDefinitionEgg) HasObjectRequestOptions() bool`

HasObjectRequestOptions returns a boolean if a field has been set.

### GetLabels

`func (o *PublicActionDefinitionEgg) GetLabels() map[string]PublicActionLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PublicActionDefinitionEgg) GetLabelsOk() (*map[string]PublicActionLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PublicActionDefinitionEgg) SetLabels(v map[string]PublicActionLabels)`

SetLabels sets Labels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


