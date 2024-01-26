/*
Tickets

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tickets

import (
	"encoding/json"
)

// SimplePublicObjectInputForCreate struct for SimplePublicObjectInputForCreate
type SimplePublicObjectInputForCreate struct {
	Associations []PublicAssociationsForObject `json:"associations"`
	Properties   map[string]string             `json:"properties"`
}

// NewSimplePublicObjectInputForCreate instantiates a new SimplePublicObjectInputForCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplePublicObjectInputForCreate(associations []PublicAssociationsForObject, properties map[string]string) *SimplePublicObjectInputForCreate {
	this := SimplePublicObjectInputForCreate{}
	this.Associations = associations
	this.Properties = properties
	return &this
}

// NewSimplePublicObjectInputForCreateWithDefaults instantiates a new SimplePublicObjectInputForCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplePublicObjectInputForCreateWithDefaults() *SimplePublicObjectInputForCreate {
	this := SimplePublicObjectInputForCreate{}
	return &this
}

// GetAssociations returns the Associations field value
func (o *SimplePublicObjectInputForCreate) GetAssociations() []PublicAssociationsForObject {
	if o == nil {
		var ret []PublicAssociationsForObject
		return ret
	}

	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectInputForCreate) GetAssociationsOk() ([]PublicAssociationsForObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Associations, true
}

// SetAssociations sets field value
func (o *SimplePublicObjectInputForCreate) SetAssociations(v []PublicAssociationsForObject) {
	o.Associations = v
}

// GetProperties returns the Properties field value
func (o *SimplePublicObjectInputForCreate) GetProperties() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectInputForCreate) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *SimplePublicObjectInputForCreate) SetProperties(v map[string]string) {
	o.Properties = v
}

func (o SimplePublicObjectInputForCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["associations"] = o.Associations
	}
	if true {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableSimplePublicObjectInputForCreate struct {
	value *SimplePublicObjectInputForCreate
	isSet bool
}

func (v NullableSimplePublicObjectInputForCreate) Get() *SimplePublicObjectInputForCreate {
	return v.value
}

func (v *NullableSimplePublicObjectInputForCreate) Set(val *SimplePublicObjectInputForCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplePublicObjectInputForCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplePublicObjectInputForCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplePublicObjectInputForCreate(val *SimplePublicObjectInputForCreate) *NullableSimplePublicObjectInputForCreate {
	return &NullableSimplePublicObjectInputForCreate{value: val, isSet: true}
}

func (v NullableSimplePublicObjectInputForCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplePublicObjectInputForCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
