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

// PublicSingleFieldDependency struct for PublicSingleFieldDependency
type PublicSingleFieldDependency struct {
	DependencyType       string   `json:"dependencyType"`
	DependentFieldNames  []string `json:"dependentFieldNames"`
	ControllingFieldName string   `json:"controllingFieldName"`
}

// NewPublicSingleFieldDependency instantiates a new PublicSingleFieldDependency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicSingleFieldDependency(dependencyType string, dependentFieldNames []string, controllingFieldName string) *PublicSingleFieldDependency {
	this := PublicSingleFieldDependency{}
	this.DependencyType = dependencyType
	this.DependentFieldNames = dependentFieldNames
	this.ControllingFieldName = controllingFieldName
	return &this
}

// NewPublicSingleFieldDependencyWithDefaults instantiates a new PublicSingleFieldDependency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicSingleFieldDependencyWithDefaults() *PublicSingleFieldDependency {
	this := PublicSingleFieldDependency{}
	var dependencyType string = "SINGLE_FIELD"
	this.DependencyType = dependencyType
	return &this
}

// GetDependencyType returns the DependencyType field value
func (o *PublicSingleFieldDependency) GetDependencyType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DependencyType
}

// GetDependencyTypeOk returns a tuple with the DependencyType field value
// and a boolean to check if the value has been set.
func (o *PublicSingleFieldDependency) GetDependencyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DependencyType, true
}

// SetDependencyType sets field value
func (o *PublicSingleFieldDependency) SetDependencyType(v string) {
	o.DependencyType = v
}

// GetDependentFieldNames returns the DependentFieldNames field value
func (o *PublicSingleFieldDependency) GetDependentFieldNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DependentFieldNames
}

// GetDependentFieldNamesOk returns a tuple with the DependentFieldNames field value
// and a boolean to check if the value has been set.
func (o *PublicSingleFieldDependency) GetDependentFieldNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DependentFieldNames, true
}

// SetDependentFieldNames sets field value
func (o *PublicSingleFieldDependency) SetDependentFieldNames(v []string) {
	o.DependentFieldNames = v
}

// GetControllingFieldName returns the ControllingFieldName field value
func (o *PublicSingleFieldDependency) GetControllingFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ControllingFieldName
}

// GetControllingFieldNameOk returns a tuple with the ControllingFieldName field value
// and a boolean to check if the value has been set.
func (o *PublicSingleFieldDependency) GetControllingFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ControllingFieldName, true
}

// SetControllingFieldName sets field value
func (o *PublicSingleFieldDependency) SetControllingFieldName(v string) {
	o.ControllingFieldName = v
}

func (o PublicSingleFieldDependency) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dependencyType"] = o.DependencyType
	}
	if true {
		toSerialize["dependentFieldNames"] = o.DependentFieldNames
	}
	if true {
		toSerialize["controllingFieldName"] = o.ControllingFieldName
	}
	return json.Marshal(toSerialize)
}

type NullablePublicSingleFieldDependency struct {
	value *PublicSingleFieldDependency
	isSet bool
}

func (v NullablePublicSingleFieldDependency) Get() *PublicSingleFieldDependency {
	return v.value
}

func (v *NullablePublicSingleFieldDependency) Set(val *PublicSingleFieldDependency) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicSingleFieldDependency) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicSingleFieldDependency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicSingleFieldDependency(val *PublicSingleFieldDependency) *NullablePublicSingleFieldDependency {
	return &NullablePublicSingleFieldDependency{value: val, isSet: true}
}

func (v NullablePublicSingleFieldDependency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicSingleFieldDependency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
