/*
Automation Actions V4

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package actions

import (
	"encoding/json"
)

// PublicActionDefinitionEgg struct for PublicActionDefinitionEgg
type PublicActionDefinitionEgg struct {
	InputFields            []InputFieldDefinition                              `json:"inputFields"`
	OutputFields           []OutputFieldDefinition                             `json:"outputFields,omitempty"`
	ArchivedAt             *int64                                              `json:"archivedAt,omitempty"`
	Functions              []PublicActionFunction                              `json:"functions"`
	ActionUrl              string                                              `json:"actionUrl"`
	InputFieldDependencies []PublicActionDefinitionInputFieldDependenciesInner `json:"inputFieldDependencies,omitempty"`
	Published              bool                                                `json:"published"`
	ExecutionRules         []PublicExecutionTranslationRule                    `json:"executionRules,omitempty"`
	ObjectTypes            []string                                            `json:"objectTypes"`
	ObjectRequestOptions   *PublicObjectRequestOptions                         `json:"objectRequestOptions,omitempty"`
	Labels                 map[string]PublicActionLabels                       `json:"labels"`
}

// NewPublicActionDefinitionEgg instantiates a new PublicActionDefinitionEgg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicActionDefinitionEgg(inputFields []InputFieldDefinition, functions []PublicActionFunction, actionUrl string, published bool, objectTypes []string, labels map[string]PublicActionLabels) *PublicActionDefinitionEgg {
	this := PublicActionDefinitionEgg{}
	this.InputFields = inputFields
	this.Functions = functions
	this.ActionUrl = actionUrl
	this.Published = published
	this.ObjectTypes = objectTypes
	this.Labels = labels
	return &this
}

// NewPublicActionDefinitionEggWithDefaults instantiates a new PublicActionDefinitionEgg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicActionDefinitionEggWithDefaults() *PublicActionDefinitionEgg {
	this := PublicActionDefinitionEgg{}
	return &this
}

// GetInputFields returns the InputFields field value
func (o *PublicActionDefinitionEgg) GetInputFields() []InputFieldDefinition {
	if o == nil {
		var ret []InputFieldDefinition
		return ret
	}

	return o.InputFields
}

// GetInputFieldsOk returns a tuple with the InputFields field value
// and a boolean to check if the value has been set.
func (o *PublicActionDefinitionEgg) GetInputFieldsOk() ([]InputFieldDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.InputFields, true
}

// SetInputFields sets field value
func (o *PublicActionDefinitionEgg) SetInputFields(v []InputFieldDefinition) {
	o.InputFields = v
}

// GetOutputFields returns the OutputFields field value if set, zero value otherwise.
func (o *PublicActionDefinitionEgg) GetOutputFields() []OutputFieldDefinition {
	if o == nil || o.OutputFields == nil {
		var ret []OutputFieldDefinition
		return ret
	}
	return o.OutputFields
}

// GetOutputFieldsOk returns a tuple with the OutputFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicActionDefinitionEgg) GetOutputFieldsOk() ([]OutputFieldDefinition, bool) {
	if o == nil || o.OutputFields == nil {
		return nil, false
	}
	return o.OutputFields, true
}

// HasOutputFields returns a boolean if a field has been set.
func (o *PublicActionDefinitionEgg) HasOutputFields() bool {
	if o != nil && o.OutputFields != nil {
		return true
	}

	return false
}

// SetOutputFields gets a reference to the given []OutputFieldDefinition and assigns it to the OutputFields field.
func (o *PublicActionDefinitionEgg) SetOutputFields(v []OutputFieldDefinition) {
	o.OutputFields = v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *PublicActionDefinitionEgg) GetArchivedAt() int64 {
	if o == nil || o.ArchivedAt == nil {
		var ret int64
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicActionDefinitionEgg) GetArchivedAtOk() (*int64, bool) {
	if o == nil || o.ArchivedAt == nil {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *PublicActionDefinitionEgg) HasArchivedAt() bool {
	if o != nil && o.ArchivedAt != nil {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given int64 and assigns it to the ArchivedAt field.
func (o *PublicActionDefinitionEgg) SetArchivedAt(v int64) {
	o.ArchivedAt = &v
}

// GetFunctions returns the Functions field value
func (o *PublicActionDefinitionEgg) GetFunctions() []PublicActionFunction {
	if o == nil {
		var ret []PublicActionFunction
		return ret
	}

	return o.Functions
}

// GetFunctionsOk returns a tuple with the Functions field value
// and a boolean to check if the value has been set.
func (o *PublicActionDefinitionEgg) GetFunctionsOk() ([]PublicActionFunction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Functions, true
}

// SetFunctions sets field value
func (o *PublicActionDefinitionEgg) SetFunctions(v []PublicActionFunction) {
	o.Functions = v
}

// GetActionUrl returns the ActionUrl field value
func (o *PublicActionDefinitionEgg) GetActionUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionUrl
}

// GetActionUrlOk returns a tuple with the ActionUrl field value
// and a boolean to check if the value has been set.
func (o *PublicActionDefinitionEgg) GetActionUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionUrl, true
}

// SetActionUrl sets field value
func (o *PublicActionDefinitionEgg) SetActionUrl(v string) {
	o.ActionUrl = v
}

// GetInputFieldDependencies returns the InputFieldDependencies field value if set, zero value otherwise.
func (o *PublicActionDefinitionEgg) GetInputFieldDependencies() []PublicActionDefinitionInputFieldDependenciesInner {
	if o == nil || o.InputFieldDependencies == nil {
		var ret []PublicActionDefinitionInputFieldDependenciesInner
		return ret
	}
	return o.InputFieldDependencies
}

// GetInputFieldDependenciesOk returns a tuple with the InputFieldDependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicActionDefinitionEgg) GetInputFieldDependenciesOk() ([]PublicActionDefinitionInputFieldDependenciesInner, bool) {
	if o == nil || o.InputFieldDependencies == nil {
		return nil, false
	}
	return o.InputFieldDependencies, true
}

// HasInputFieldDependencies returns a boolean if a field has been set.
func (o *PublicActionDefinitionEgg) HasInputFieldDependencies() bool {
	if o != nil && o.InputFieldDependencies != nil {
		return true
	}

	return false
}

// SetInputFieldDependencies gets a reference to the given []PublicActionDefinitionInputFieldDependenciesInner and assigns it to the InputFieldDependencies field.
func (o *PublicActionDefinitionEgg) SetInputFieldDependencies(v []PublicActionDefinitionInputFieldDependenciesInner) {
	o.InputFieldDependencies = v
}

// GetPublished returns the Published field value
func (o *PublicActionDefinitionEgg) GetPublished() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Published
}

// GetPublishedOk returns a tuple with the Published field value
// and a boolean to check if the value has been set.
func (o *PublicActionDefinitionEgg) GetPublishedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Published, true
}

// SetPublished sets field value
func (o *PublicActionDefinitionEgg) SetPublished(v bool) {
	o.Published = v
}

// GetExecutionRules returns the ExecutionRules field value if set, zero value otherwise.
func (o *PublicActionDefinitionEgg) GetExecutionRules() []PublicExecutionTranslationRule {
	if o == nil || o.ExecutionRules == nil {
		var ret []PublicExecutionTranslationRule
		return ret
	}
	return o.ExecutionRules
}

// GetExecutionRulesOk returns a tuple with the ExecutionRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicActionDefinitionEgg) GetExecutionRulesOk() ([]PublicExecutionTranslationRule, bool) {
	if o == nil || o.ExecutionRules == nil {
		return nil, false
	}
	return o.ExecutionRules, true
}

// HasExecutionRules returns a boolean if a field has been set.
func (o *PublicActionDefinitionEgg) HasExecutionRules() bool {
	if o != nil && o.ExecutionRules != nil {
		return true
	}

	return false
}

// SetExecutionRules gets a reference to the given []PublicExecutionTranslationRule and assigns it to the ExecutionRules field.
func (o *PublicActionDefinitionEgg) SetExecutionRules(v []PublicExecutionTranslationRule) {
	o.ExecutionRules = v
}

// GetObjectTypes returns the ObjectTypes field value
func (o *PublicActionDefinitionEgg) GetObjectTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ObjectTypes
}

// GetObjectTypesOk returns a tuple with the ObjectTypes field value
// and a boolean to check if the value has been set.
func (o *PublicActionDefinitionEgg) GetObjectTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectTypes, true
}

// SetObjectTypes sets field value
func (o *PublicActionDefinitionEgg) SetObjectTypes(v []string) {
	o.ObjectTypes = v
}

// GetObjectRequestOptions returns the ObjectRequestOptions field value if set, zero value otherwise.
func (o *PublicActionDefinitionEgg) GetObjectRequestOptions() PublicObjectRequestOptions {
	if o == nil || o.ObjectRequestOptions == nil {
		var ret PublicObjectRequestOptions
		return ret
	}
	return *o.ObjectRequestOptions
}

// GetObjectRequestOptionsOk returns a tuple with the ObjectRequestOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicActionDefinitionEgg) GetObjectRequestOptionsOk() (*PublicObjectRequestOptions, bool) {
	if o == nil || o.ObjectRequestOptions == nil {
		return nil, false
	}
	return o.ObjectRequestOptions, true
}

// HasObjectRequestOptions returns a boolean if a field has been set.
func (o *PublicActionDefinitionEgg) HasObjectRequestOptions() bool {
	if o != nil && o.ObjectRequestOptions != nil {
		return true
	}

	return false
}

// SetObjectRequestOptions gets a reference to the given PublicObjectRequestOptions and assigns it to the ObjectRequestOptions field.
func (o *PublicActionDefinitionEgg) SetObjectRequestOptions(v PublicObjectRequestOptions) {
	o.ObjectRequestOptions = &v
}

// GetLabels returns the Labels field value
func (o *PublicActionDefinitionEgg) GetLabels() map[string]PublicActionLabels {
	if o == nil {
		var ret map[string]PublicActionLabels
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *PublicActionDefinitionEgg) GetLabelsOk() (*map[string]PublicActionLabels, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Labels, true
}

// SetLabels sets field value
func (o *PublicActionDefinitionEgg) SetLabels(v map[string]PublicActionLabels) {
	o.Labels = v
}

func (o PublicActionDefinitionEgg) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["inputFields"] = o.InputFields
	}
	if o.OutputFields != nil {
		toSerialize["outputFields"] = o.OutputFields
	}
	if o.ArchivedAt != nil {
		toSerialize["archivedAt"] = o.ArchivedAt
	}
	if true {
		toSerialize["functions"] = o.Functions
	}
	if true {
		toSerialize["actionUrl"] = o.ActionUrl
	}
	if o.InputFieldDependencies != nil {
		toSerialize["inputFieldDependencies"] = o.InputFieldDependencies
	}
	if true {
		toSerialize["published"] = o.Published
	}
	if o.ExecutionRules != nil {
		toSerialize["executionRules"] = o.ExecutionRules
	}
	if true {
		toSerialize["objectTypes"] = o.ObjectTypes
	}
	if o.ObjectRequestOptions != nil {
		toSerialize["objectRequestOptions"] = o.ObjectRequestOptions
	}
	if true {
		toSerialize["labels"] = o.Labels
	}
	return json.Marshal(toSerialize)
}

type NullablePublicActionDefinitionEgg struct {
	value *PublicActionDefinitionEgg
	isSet bool
}

func (v NullablePublicActionDefinitionEgg) Get() *PublicActionDefinitionEgg {
	return v.value
}

func (v *NullablePublicActionDefinitionEgg) Set(val *PublicActionDefinitionEgg) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicActionDefinitionEgg) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicActionDefinitionEgg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicActionDefinitionEgg(val *PublicActionDefinitionEgg) *NullablePublicActionDefinitionEgg {
	return &NullablePublicActionDefinitionEgg{value: val, isSet: true}
}

func (v NullablePublicActionDefinitionEgg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicActionDefinitionEgg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}