# PublicActionDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Functions** | [**[]PublicActionFunctionIdentifier**](PublicActionFunctionIdentifier.md) |  | 
**ActionUrl** | **string** |  | 
**Published** | **bool** |  | 
**Labels** | [**map[string]PublicActionLabels**](PublicActionLabels.md) |  | 
**InputFields** | [**[]InputFieldDefinition**](InputFieldDefinition.md) |  | 
**OutputFields** | Pointer to [**[]OutputFieldDefinition**](OutputFieldDefinition.md) |  | [optional] 
**RevisionId** | **string** |  | 
**ArchivedAt** | Pointer to **int64** |  | [optional] 
**InputFieldDependencies** | Pointer to [**[]PublicActionDefinitionInputFieldDependenciesInner**](PublicActionDefinitionInputFieldDependenciesInner.md) |  | [optional] 
**ExecutionRules** | Pointer to [**[]PublicExecutionTranslationRule**](PublicExecutionTranslationRule.md) |  | [optional] 
**Id** | **string** |  | 
**ObjectTypes** | **[]string** |  | 
**ObjectRequestOptions** | Pointer to [**PublicObjectRequestOptions**](PublicObjectRequestOptions.md) |  | [optional] 

## Methods

### NewPublicActionDefinition

`func NewPublicActionDefinition(functions []PublicActionFunctionIdentifier, actionUrl string, published bool, labels map[string]PublicActionLabels, inputFields []InputFieldDefinition, revisionId string, id string, objectTypes []string, ) *PublicActionDefinition`

NewPublicActionDefinition instantiates a new PublicActionDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicActionDefinitionWithDefaults

`func NewPublicActionDefinitionWithDefaults() *PublicActionDefinition`

NewPublicActionDefinitionWithDefaults instantiates a new PublicActionDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctions

`func (o *PublicActionDefinition) GetFunctions() []PublicActionFunctionIdentifier`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *PublicActionDefinition) GetFunctionsOk() (*[]PublicActionFunctionIdentifier, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *PublicActionDefinition) SetFunctions(v []PublicActionFunctionIdentifier)`

SetFunctions sets Functions field to given value.


### GetActionUrl

`func (o *PublicActionDefinition) GetActionUrl() string`

GetActionUrl returns the ActionUrl field if non-nil, zero value otherwise.

### GetActionUrlOk

`func (o *PublicActionDefinition) GetActionUrlOk() (*string, bool)`

GetActionUrlOk returns a tuple with the ActionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionUrl

`func (o *PublicActionDefinition) SetActionUrl(v string)`

SetActionUrl sets ActionUrl field to given value.


### GetPublished

`func (o *PublicActionDefinition) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *PublicActionDefinition) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *PublicActionDefinition) SetPublished(v bool)`

SetPublished sets Published field to given value.


### GetLabels

`func (o *PublicActionDefinition) GetLabels() map[string]PublicActionLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PublicActionDefinition) GetLabelsOk() (*map[string]PublicActionLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PublicActionDefinition) SetLabels(v map[string]PublicActionLabels)`

SetLabels sets Labels field to given value.


### GetInputFields

`func (o *PublicActionDefinition) GetInputFields() []InputFieldDefinition`

GetInputFields returns the InputFields field if non-nil, zero value otherwise.

### GetInputFieldsOk

`func (o *PublicActionDefinition) GetInputFieldsOk() (*[]InputFieldDefinition, bool)`

GetInputFieldsOk returns a tuple with the InputFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFields

`func (o *PublicActionDefinition) SetInputFields(v []InputFieldDefinition)`

SetInputFields sets InputFields field to given value.


### GetOutputFields

`func (o *PublicActionDefinition) GetOutputFields() []OutputFieldDefinition`

GetOutputFields returns the OutputFields field if non-nil, zero value otherwise.

### GetOutputFieldsOk

`func (o *PublicActionDefinition) GetOutputFieldsOk() (*[]OutputFieldDefinition, bool)`

GetOutputFieldsOk returns a tuple with the OutputFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFields

`func (o *PublicActionDefinition) SetOutputFields(v []OutputFieldDefinition)`

SetOutputFields sets OutputFields field to given value.

### HasOutputFields

`func (o *PublicActionDefinition) HasOutputFields() bool`

HasOutputFields returns a boolean if a field has been set.

### GetRevisionId

`func (o *PublicActionDefinition) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *PublicActionDefinition) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *PublicActionDefinition) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.


### GetArchivedAt

`func (o *PublicActionDefinition) GetArchivedAt() int64`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *PublicActionDefinition) GetArchivedAtOk() (*int64, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *PublicActionDefinition) SetArchivedAt(v int64)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *PublicActionDefinition) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetInputFieldDependencies

`func (o *PublicActionDefinition) GetInputFieldDependencies() []PublicActionDefinitionInputFieldDependenciesInner`

GetInputFieldDependencies returns the InputFieldDependencies field if non-nil, zero value otherwise.

### GetInputFieldDependenciesOk

`func (o *PublicActionDefinition) GetInputFieldDependenciesOk() (*[]PublicActionDefinitionInputFieldDependenciesInner, bool)`

GetInputFieldDependenciesOk returns a tuple with the InputFieldDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFieldDependencies

`func (o *PublicActionDefinition) SetInputFieldDependencies(v []PublicActionDefinitionInputFieldDependenciesInner)`

SetInputFieldDependencies sets InputFieldDependencies field to given value.

### HasInputFieldDependencies

`func (o *PublicActionDefinition) HasInputFieldDependencies() bool`

HasInputFieldDependencies returns a boolean if a field has been set.

### GetExecutionRules

`func (o *PublicActionDefinition) GetExecutionRules() []PublicExecutionTranslationRule`

GetExecutionRules returns the ExecutionRules field if non-nil, zero value otherwise.

### GetExecutionRulesOk

`func (o *PublicActionDefinition) GetExecutionRulesOk() (*[]PublicExecutionTranslationRule, bool)`

GetExecutionRulesOk returns a tuple with the ExecutionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionRules

`func (o *PublicActionDefinition) SetExecutionRules(v []PublicExecutionTranslationRule)`

SetExecutionRules sets ExecutionRules field to given value.

### HasExecutionRules

`func (o *PublicActionDefinition) HasExecutionRules() bool`

HasExecutionRules returns a boolean if a field has been set.

### GetId

`func (o *PublicActionDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicActionDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicActionDefinition) SetId(v string)`

SetId sets Id field to given value.


### GetObjectTypes

`func (o *PublicActionDefinition) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *PublicActionDefinition) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *PublicActionDefinition) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.


### GetObjectRequestOptions

`func (o *PublicActionDefinition) GetObjectRequestOptions() PublicObjectRequestOptions`

GetObjectRequestOptions returns the ObjectRequestOptions field if non-nil, zero value otherwise.

### GetObjectRequestOptionsOk

`func (o *PublicActionDefinition) GetObjectRequestOptionsOk() (*PublicObjectRequestOptions, bool)`

GetObjectRequestOptionsOk returns a tuple with the ObjectRequestOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectRequestOptions

`func (o *PublicActionDefinition) SetObjectRequestOptions(v PublicObjectRequestOptions)`

SetObjectRequestOptions sets ObjectRequestOptions field to given value.

### HasObjectRequestOptions

`func (o *PublicActionDefinition) HasObjectRequestOptions() bool`

HasObjectRequestOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


