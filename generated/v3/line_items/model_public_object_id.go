/*
Line Items

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package line_items

import (
	"encoding/json"
)

// PublicObjectId struct for PublicObjectId
type PublicObjectId struct {
	Id string `json:"id"`
}

// NewPublicObjectId instantiates a new PublicObjectId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicObjectId(id string) *PublicObjectId {
	this := PublicObjectId{}
	this.Id = id
	return &this
}

// NewPublicObjectIdWithDefaults instantiates a new PublicObjectId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicObjectIdWithDefaults() *PublicObjectId {
	this := PublicObjectId{}
	return &this
}

// GetId returns the Id field value
func (o *PublicObjectId) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PublicObjectId) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PublicObjectId) SetId(v string) {
	o.Id = v
}

func (o PublicObjectId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullablePublicObjectId struct {
	value *PublicObjectId
	isSet bool
}

func (v NullablePublicObjectId) Get() *PublicObjectId {
	return v.value
}

func (v *NullablePublicObjectId) Set(val *PublicObjectId) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicObjectId) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicObjectId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicObjectId(val *PublicObjectId) *NullablePublicObjectId {
	return &NullablePublicObjectId{value: val, isSet: true}
}

func (v NullablePublicObjectId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicObjectId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
