# PublicActionDefinitionPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputFields** | Pointer to [**[]InputFieldDefinition**](InputFieldDefinition.md) |  | [optional] 
**OutputFields** | Pointer to [**[]OutputFieldDefinition**](OutputFieldDefinition.md) |  | [optional] 
**ActionUrl** | Pointer to **string** |  | [optional] 
**InputFieldDependencies** | Pointer to [**[]PublicActionDefinitionInputFieldDependenciesInner**](PublicActionDefinitionInputFieldDependenciesInner.md) |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**ExecutionRules** | Pointer to [**[]PublicExecutionTranslationRule**](PublicExecutionTranslationRule.md) |  | [optional] 
**ObjectTypes** | Pointer to **[]string** |  | [optional] 
**ObjectRequestOptions** | Pointer to [**PublicObjectRequestOptions**](PublicObjectRequestOptions.md) |  | [optional] 
**Labels** | Pointer to [**map[string]PublicActionLabels**](PublicActionLabels.md) |  | [optional] 

## Methods

### NewPublicActionDefinitionPatch

`func NewPublicActionDefinitionPatch() *PublicActionDefinitionPatch`

NewPublicActionDefinitionPatch instantiates a new PublicActionDefinitionPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicActionDefinitionPatchWithDefaults

`func NewPublicActionDefinitionPatchWithDefaults() *PublicActionDefinitionPatch`

NewPublicActionDefinitionPatchWithDefaults instantiates a new PublicActionDefinitionPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputFields

`func (o *PublicActionDefinitionPatch) GetInputFields() []InputFieldDefinition`

GetInputFields returns the InputFields field if non-nil, zero value otherwise.

### GetInputFieldsOk

`func (o *PublicActionDefinitionPatch) GetInputFieldsOk() (*[]InputFieldDefinition, bool)`

GetInputFieldsOk returns a tuple with the InputFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFields

`func (o *PublicActionDefinitionPatch) SetInputFields(v []InputFieldDefinition)`

SetInputFields sets InputFields field to given value.

### HasInputFields

`func (o *PublicActionDefinitionPatch) HasInputFields() bool`

HasInputFields returns a boolean if a field has been set.

### GetOutputFields

`func (o *PublicActionDefinitionPatch) GetOutputFields() []OutputFieldDefinition`

GetOutputFields returns the OutputFields field if non-nil, zero value otherwise.

### GetOutputFieldsOk

`func (o *PublicActionDefinitionPatch) GetOutputFieldsOk() (*[]OutputFieldDefinition, bool)`

GetOutputFieldsOk returns a tuple with the OutputFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFields

`func (o *PublicActionDefinitionPatch) SetOutputFields(v []OutputFieldDefinition)`

SetOutputFields sets OutputFields field to given value.

### HasOutputFields

`func (o *PublicActionDefinitionPatch) HasOutputFields() bool`

HasOutputFields returns a boolean if a field has been set.

### GetActionUrl

`func (o *PublicActionDefinitionPatch) GetActionUrl() string`

GetActionUrl returns the ActionUrl field if non-nil, zero value otherwise.

### GetActionUrlOk

`func (o *PublicActionDefinitionPatch) GetActionUrlOk() (*string, bool)`

GetActionUrlOk returns a tuple with the ActionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionUrl

`func (o *PublicActionDefinitionPatch) SetActionUrl(v string)`

SetActionUrl sets ActionUrl field to given value.

### HasActionUrl

`func (o *PublicActionDefinitionPatch) HasActionUrl() bool`

HasActionUrl returns a boolean if a field has been set.

### GetInputFieldDependencies

`func (o *PublicActionDefinitionPatch) GetInputFieldDependencies() []PublicActionDefinitionInputFieldDependenciesInner`

GetInputFieldDependencies returns the InputFieldDependencies field if non-nil, zero value otherwise.

### GetInputFieldDependenciesOk

`func (o *PublicActionDefinitionPatch) GetInputFieldDependenciesOk() (*[]PublicActionDefinitionInputFieldDependenciesInner, bool)`

GetInputFieldDependenciesOk returns a tuple with the InputFieldDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFieldDependencies

`func (o *PublicActionDefinitionPatch) SetInputFieldDependencies(v []PublicActionDefinitionInputFieldDependenciesInner)`

SetInputFieldDependencies sets InputFieldDependencies field to given value.

### HasInputFieldDependencies

`func (o *PublicActionDefinitionPatch) HasInputFieldDependencies() bool`

HasInputFieldDependencies returns a boolean if a field has been set.

### GetPublished

`func (o *PublicActionDefinitionPatch) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *PublicActionDefinitionPatch) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *PublicActionDefinitionPatch) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *PublicActionDefinitionPatch) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetExecutionRules

`func (o *PublicActionDefinitionPatch) GetExecutionRules() []PublicExecutionTranslationRule`

GetExecutionRules returns the ExecutionRules field if non-nil, zero value otherwise.

### GetExecutionRulesOk

`func (o *PublicActionDefinitionPatch) GetExecutionRulesOk() (*[]PublicExecutionTranslationRule, bool)`

GetExecutionRulesOk returns a tuple with the ExecutionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionRules

`func (o *PublicActionDefinitionPatch) SetExecutionRules(v []PublicExecutionTranslationRule)`

SetExecutionRules sets ExecutionRules field to given value.

### HasExecutionRules

`func (o *PublicActionDefinitionPatch) HasExecutionRules() bool`

HasExecutionRules returns a boolean if a field has been set.

### GetObjectTypes

`func (o *PublicActionDefinitionPatch) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *PublicActionDefinitionPatch) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *PublicActionDefinitionPatch) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.

### HasObjectTypes

`func (o *PublicActionDefinitionPatch) HasObjectTypes() bool`

HasObjectTypes returns a boolean if a field has been set.

### GetObjectRequestOptions

`func (o *PublicActionDefinitionPatch) GetObjectRequestOptions() PublicObjectRequestOptions`

GetObjectRequestOptions returns the ObjectRequestOptions field if non-nil, zero value otherwise.

### GetObjectRequestOptionsOk

`func (o *PublicActionDefinitionPatch) GetObjectRequestOptionsOk() (*PublicObjectRequestOptions, bool)`

GetObjectRequestOptionsOk returns a tuple with the ObjectRequestOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectRequestOptions

`func (o *PublicActionDefinitionPatch) SetObjectRequestOptions(v PublicObjectRequestOptions)`

SetObjectRequestOptions sets ObjectRequestOptions field to given value.

### HasObjectRequestOptions

`func (o *PublicActionDefinitionPatch) HasObjectRequestOptions() bool`

HasObjectRequestOptions returns a boolean if a field has been set.

### GetLabels

`func (o *PublicActionDefinitionPatch) GetLabels() map[string]PublicActionLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PublicActionDefinitionPatch) GetLabelsOk() (*map[string]PublicActionLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PublicActionDefinitionPatch) SetLabels(v map[string]PublicActionLabels)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PublicActionDefinitionPatch) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


