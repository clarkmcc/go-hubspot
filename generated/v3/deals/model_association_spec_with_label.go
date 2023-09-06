/*
Deals

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package deals

import (
	"encoding/json"
)

// AssociationSpecWithLabel struct for AssociationSpecWithLabel
type AssociationSpecWithLabel struct {
	Category string  `json:"category"`
	TypeId   int32   `json:"typeId"`
	Label    *string `json:"label,omitempty"`
}

// NewAssociationSpecWithLabel instantiates a new AssociationSpecWithLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociationSpecWithLabel(category string, typeId int32) *AssociationSpecWithLabel {
	this := AssociationSpecWithLabel{}
	this.Category = category
	this.TypeId = typeId
	return &this
}

// NewAssociationSpecWithLabelWithDefaults instantiates a new AssociationSpecWithLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociationSpecWithLabelWithDefaults() *AssociationSpecWithLabel {
	this := AssociationSpecWithLabel{}
	return &this
}

// GetCategory returns the Category field value
func (o *AssociationSpecWithLabel) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *AssociationSpecWithLabel) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *AssociationSpecWithLabel) SetCategory(v string) {
	o.Category = v
}

// GetTypeId returns the TypeId field value
func (o *AssociationSpecWithLabel) GetTypeId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value
// and a boolean to check if the value has been set.
func (o *AssociationSpecWithLabel) GetTypeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeId, true
}

// SetTypeId sets field value
func (o *AssociationSpecWithLabel) SetTypeId(v int32) {
	o.TypeId = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *AssociationSpecWithLabel) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociationSpecWithLabel) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *AssociationSpecWithLabel) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *AssociationSpecWithLabel) SetLabel(v string) {
	o.Label = &v
}

func (o AssociationSpecWithLabel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["category"] = o.Category
	}
	if true {
		toSerialize["typeId"] = o.TypeId
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	return json.Marshal(toSerialize)
}

type NullableAssociationSpecWithLabel struct {
	value *AssociationSpecWithLabel
	isSet bool
}

func (v NullableAssociationSpecWithLabel) Get() *AssociationSpecWithLabel {
	return v.value
}

func (v *NullableAssociationSpecWithLabel) Set(val *AssociationSpecWithLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociationSpecWithLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociationSpecWithLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociationSpecWithLabel(val *AssociationSpecWithLabel) *NullableAssociationSpecWithLabel {
	return &NullableAssociationSpecWithLabel{value: val, isSet: true}
}

func (v NullableAssociationSpecWithLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssociationSpecWithLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
