# ExtensionActionDefinitionPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionUrl** | Pointer to **string** | The URL that will accept an HTTPS request each time workflows executes the custom action. | [optional] 
**Published** | Pointer to **bool** | Whether this custom action is published to customers. | [optional] 
**InputFields** | Pointer to [**[]InputFieldDefinition**](InputFieldDefinition.md) | The list of input fields to display in this custom action. | [optional] 
**ObjectRequestOptions** | Pointer to [**ObjectRequestOptions**](ObjectRequestOptions.md) |  | [optional] 
**InputFieldDependencies** | Pointer to [**[]ExtensionActionDefinitionPatchInputFieldDependenciesInner**](ExtensionActionDefinitionPatchInputFieldDependenciesInner.md) | A list of dependencies between the input fields. These configure when the input fields should be visible. | [optional] 
**Labels** | Pointer to [**map[string]ActionLabels**](ActionLabels.md) | The user-facing labels for the custom action. | [optional] 
**ObjectTypes** | Pointer to **[]string** | The object types that this custom action supports. | [optional] 

## Methods

### NewExtensionActionDefinitionPatch

`func NewExtensionActionDefinitionPatch() *ExtensionActionDefinitionPatch`

NewExtensionActionDefinitionPatch instantiates a new ExtensionActionDefinitionPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionActionDefinitionPatchWithDefaults

`func NewExtensionActionDefinitionPatchWithDefaults() *ExtensionActionDefinitionPatch`

NewExtensionActionDefinitionPatchWithDefaults instantiates a new ExtensionActionDefinitionPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionUrl

`func (o *ExtensionActionDefinitionPatch) GetActionUrl() string`

GetActionUrl returns the ActionUrl field if non-nil, zero value otherwise.

### GetActionUrlOk

`func (o *ExtensionActionDefinitionPatch) GetActionUrlOk() (*string, bool)`

GetActionUrlOk returns a tuple with the ActionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionUrl

`func (o *ExtensionActionDefinitionPatch) SetActionUrl(v string)`

SetActionUrl sets ActionUrl field to given value.

### HasActionUrl

`func (o *ExtensionActionDefinitionPatch) HasActionUrl() bool`

HasActionUrl returns a boolean if a field has been set.

### GetPublished

`func (o *ExtensionActionDefinitionPatch) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *ExtensionActionDefinitionPatch) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *ExtensionActionDefinitionPatch) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *ExtensionActionDefinitionPatch) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetInputFields

`func (o *ExtensionActionDefinitionPatch) GetInputFields() []InputFieldDefinition`

GetInputFields returns the InputFields field if non-nil, zero value otherwise.

### GetInputFieldsOk

`func (o *ExtensionActionDefinitionPatch) GetInputFieldsOk() (*[]InputFieldDefinition, bool)`

GetInputFieldsOk returns a tuple with the InputFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFields

`func (o *ExtensionActionDefinitionPatch) SetInputFields(v []InputFieldDefinition)`

SetInputFields sets InputFields field to given value.

### HasInputFields

`func (o *ExtensionActionDefinitionPatch) HasInputFields() bool`

HasInputFields returns a boolean if a field has been set.

### GetObjectRequestOptions

`func (o *ExtensionActionDefinitionPatch) GetObjectRequestOptions() ObjectRequestOptions`

GetObjectRequestOptions returns the ObjectRequestOptions field if non-nil, zero value otherwise.

### GetObjectRequestOptionsOk

`func (o *ExtensionActionDefinitionPatch) GetObjectRequestOptionsOk() (*ObjectRequestOptions, bool)`

GetObjectRequestOptionsOk returns a tuple with the ObjectRequestOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectRequestOptions

`func (o *ExtensionActionDefinitionPatch) SetObjectRequestOptions(v ObjectRequestOptions)`

SetObjectRequestOptions sets ObjectRequestOptions field to given value.

### HasObjectRequestOptions

`func (o *ExtensionActionDefinitionPatch) HasObjectRequestOptions() bool`

HasObjectRequestOptions returns a boolean if a field has been set.

### GetInputFieldDependencies

`func (o *ExtensionActionDefinitionPatch) GetInputFieldDependencies() []ExtensionActionDefinitionPatchInputFieldDependenciesInner`

GetInputFieldDependencies returns the InputFieldDependencies field if non-nil, zero value otherwise.

### GetInputFieldDependenciesOk

`func (o *ExtensionActionDefinitionPatch) GetInputFieldDependenciesOk() (*[]ExtensionActionDefinitionPatchInputFieldDependenciesInner, bool)`

GetInputFieldDependenciesOk returns a tuple with the InputFieldDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFieldDependencies

`func (o *ExtensionActionDefinitionPatch) SetInputFieldDependencies(v []ExtensionActionDefinitionPatchInputFieldDependenciesInner)`

SetInputFieldDependencies sets InputFieldDependencies field to given value.

### HasInputFieldDependencies

`func (o *ExtensionActionDefinitionPatch) HasInputFieldDependencies() bool`

HasInputFieldDependencies returns a boolean if a field has been set.

### GetLabels

`func (o *ExtensionActionDefinitionPatch) GetLabels() map[string]ActionLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ExtensionActionDefinitionPatch) GetLabelsOk() (*map[string]ActionLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ExtensionActionDefinitionPatch) SetLabels(v map[string]ActionLabels)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ExtensionActionDefinitionPatch) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetObjectTypes

`func (o *ExtensionActionDefinitionPatch) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *ExtensionActionDefinitionPatch) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *ExtensionActionDefinitionPatch) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.

### HasObjectTypes

`func (o *ExtensionActionDefinitionPatch) HasObjectTypes() bool`

HasObjectTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


