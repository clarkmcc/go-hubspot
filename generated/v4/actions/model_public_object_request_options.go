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

// PublicObjectRequestOptions struct for PublicObjectRequestOptions
type PublicObjectRequestOptions struct {
	Properties []string `json:"properties"`
}

// NewPublicObjectRequestOptions instantiates a new PublicObjectRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicObjectRequestOptions(properties []string) *PublicObjectRequestOptions {
	this := PublicObjectRequestOptions{}
	this.Properties = properties
	return &this
}

// NewPublicObjectRequestOptionsWithDefaults instantiates a new PublicObjectRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicObjectRequestOptionsWithDefaults() *PublicObjectRequestOptions {
	this := PublicObjectRequestOptions{}
	return &this
}

// GetProperties returns the Properties field value
func (o *PublicObjectRequestOptions) GetProperties() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *PublicObjectRequestOptions) GetPropertiesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties, true
}

// SetProperties sets field value
func (o *PublicObjectRequestOptions) SetProperties(v []string) {
	o.Properties = v
}

func (o PublicObjectRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullablePublicObjectRequestOptions struct {
	value *PublicObjectRequestOptions
	isSet bool
}

func (v NullablePublicObjectRequestOptions) Get() *PublicObjectRequestOptions {
	return v.value
}

func (v *NullablePublicObjectRequestOptions) Set(val *PublicObjectRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicObjectRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicObjectRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicObjectRequestOptions(val *PublicObjectRequestOptions) *NullablePublicObjectRequestOptions {
	return &NullablePublicObjectRequestOptions{value: val, isSet: true}
}

func (v NullablePublicObjectRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicObjectRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
