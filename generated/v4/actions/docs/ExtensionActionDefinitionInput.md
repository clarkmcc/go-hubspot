# ExtensionActionDefinitionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Functions** | [**[]ActionFunction**](ActionFunction.md) | A list of functions associated with the custom workflow action. | 
**ActionUrl** | **string** | The URL that will accept an HTTPS request each time workflows executes the custom action. | 
**Published** | **bool** | Whether this custom action is published to customers. | 
**ArchivedAt** | Pointer to **int64** | The date that this custom action was archived, if the custom action is archived. | [optional] 
**InputFields** | [**[]InputFieldDefinition**](InputFieldDefinition.md) | The list of input fields to display in this custom action. | 
**ObjectRequestOptions** | Pointer to [**ObjectRequestOptions**](ObjectRequestOptions.md) |  | [optional] 
**InputFieldDependencies** | Pointer to [**[]ExtensionActionDefinitionPatchInputFieldDependenciesInner**](ExtensionActionDefinitionPatchInputFieldDependenciesInner.md) | A list of dependencies between the input fields. These configure when the input fields should be visible. | [optional] 
**Labels** | [**map[string]ActionLabels**](ActionLabels.md) | The user-facing labels for the custom action. | 
**ObjectTypes** | **[]string** | The object types that this custom action supports. | 

## Methods

### NewExtensionActionDefinitionInput

`func NewExtensionActionDefinitionInput(functions []ActionFunction, actionUrl string, published bool, inputFields []InputFieldDefinition, labels map[string]ActionLabels, objectTypes []string, ) *ExtensionActionDefinitionInput`

NewExtensionActionDefinitionInput instantiates a new ExtensionActionDefinitionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionActionDefinitionInputWithDefaults

`func NewExtensionActionDefinitionInputWithDefaults() *ExtensionActionDefinitionInput`

NewExtensionActionDefinitionInputWithDefaults instantiates a new ExtensionActionDefinitionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctions

`func (o *ExtensionActionDefinitionInput) GetFunctions() []ActionFunction`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *ExtensionActionDefinitionInput) GetFunctionsOk() (*[]ActionFunction, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *ExtensionActionDefinitionInput) SetFunctions(v []ActionFunction)`

SetFunctions sets Functions field to given value.


### GetActionUrl

`func (o *ExtensionActionDefinitionInput) GetActionUrl() string`

GetActionUrl returns the ActionUrl field if non-nil, zero value otherwise.

### GetActionUrlOk

`func (o *ExtensionActionDefinitionInput) GetActionUrlOk() (*string, bool)`

GetActionUrlOk returns a tuple with the ActionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionUrl

`func (o *ExtensionActionDefinitionInput) SetActionUrl(v string)`

SetActionUrl sets ActionUrl field to given value.


### GetPublished

`func (o *ExtensionActionDefinitionInput) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *ExtensionActionDefinitionInput) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *ExtensionActionDefinitionInput) SetPublished(v bool)`

SetPublished sets Published field to given value.


### GetArchivedAt

`func (o *ExtensionActionDefinitionInput) GetArchivedAt() int64`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *ExtensionActionDefinitionInput) GetArchivedAtOk() (*int64, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *ExtensionActionDefinitionInput) SetArchivedAt(v int64)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *ExtensionActionDefinitionInput) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetInputFields

`func (o *ExtensionActionDefinitionInput) GetInputFields() []InputFieldDefinition`

GetInputFields returns the InputFields field if non-nil, zero value otherwise.

### GetInputFieldsOk

`func (o *ExtensionActionDefinitionInput) GetInputFieldsOk() (*[]InputFieldDefinition, bool)`

GetInputFieldsOk returns a tuple with the InputFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFields

`func (o *ExtensionActionDefinitionInput) SetInputFields(v []InputFieldDefinition)`

SetInputFields sets InputFields field to given value.


### GetObjectRequestOptions

`func (o *ExtensionActionDefinitionInput) GetObjectRequestOptions() ObjectRequestOptions`

GetObjectRequestOptions returns the ObjectRequestOptions field if non-nil, zero value otherwise.

### GetObjectRequestOptionsOk

`func (o *ExtensionActionDefinitionInput) GetObjectRequestOptionsOk() (*ObjectRequestOptions, bool)`

GetObjectRequestOptionsOk returns a tuple with the ObjectRequestOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectRequestOptions

`func (o *ExtensionActionDefinitionInput) SetObjectRequestOptions(v ObjectRequestOptions)`

SetObjectRequestOptions sets ObjectRequestOptions field to given value.

### HasObjectRequestOptions

`func (o *ExtensionActionDefinitionInput) HasObjectRequestOptions() bool`

HasObjectRequestOptions returns a boolean if a field has been set.

### GetInputFieldDependencies

`func (o *ExtensionActionDefinitionInput) GetInputFieldDependencies() []ExtensionActionDefinitionPatchInputFieldDependenciesInner`

GetInputFieldDependencies returns the InputFieldDependencies field if non-nil, zero value otherwise.

### GetInputFieldDependenciesOk

`func (o *ExtensionActionDefinitionInput) GetInputFieldDependenciesOk() (*[]ExtensionActionDefinitionPatchInputFieldDependenciesInner, bool)`

GetInputFieldDependenciesOk returns a tuple with the InputFieldDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFieldDependencies

`func (o *ExtensionActionDefinitionInput) SetInputFieldDependencies(v []ExtensionActionDefinitionPatchInputFieldDependenciesInner)`

SetInputFieldDependencies sets InputFieldDependencies field to given value.

### HasInputFieldDependencies

`func (o *ExtensionActionDefinitionInput) HasInputFieldDependencies() bool`

HasInputFieldDependencies returns a boolean if a field has been set.

### GetLabels

`func (o *ExtensionActionDefinitionInput) GetLabels() map[string]ActionLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ExtensionActionDefinitionInput) GetLabelsOk() (*map[string]ActionLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ExtensionActionDefinitionInput) SetLabels(v map[string]ActionLabels)`

SetLabels sets Labels field to given value.


### GetObjectTypes

`func (o *ExtensionActionDefinitionInput) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *ExtensionActionDefinitionInput) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *ExtensionActionDefinitionInput) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


