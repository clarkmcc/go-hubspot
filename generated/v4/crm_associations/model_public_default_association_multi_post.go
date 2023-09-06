/*
CrmPublicAssociationsServiceV4

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crm_associations

import (
	"encoding/json"
)

// PublicDefaultAssociationMultiPost struct for PublicDefaultAssociationMultiPost
type PublicDefaultAssociationMultiPost struct {
	From PublicObjectId `json:"from"`
	To   PublicObjectId `json:"to"`
}

// NewPublicDefaultAssociationMultiPost instantiates a new PublicDefaultAssociationMultiPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicDefaultAssociationMultiPost(from PublicObjectId, to PublicObjectId) *PublicDefaultAssociationMultiPost {
	this := PublicDefaultAssociationMultiPost{}
	this.From = from
	this.To = to
	return &this
}

// NewPublicDefaultAssociationMultiPostWithDefaults instantiates a new PublicDefaultAssociationMultiPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicDefaultAssociationMultiPostWithDefaults() *PublicDefaultAssociationMultiPost {
	this := PublicDefaultAssociationMultiPost{}
	return &this
}

// GetFrom returns the From field value
func (o *PublicDefaultAssociationMultiPost) GetFrom() PublicObjectId {
	if o == nil {
		var ret PublicObjectId
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *PublicDefaultAssociationMultiPost) GetFromOk() (*PublicObjectId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *PublicDefaultAssociationMultiPost) SetFrom(v PublicObjectId) {
	o.From = v
}

// GetTo returns the To field value
func (o *PublicDefaultAssociationMultiPost) GetTo() PublicObjectId {
	if o == nil {
		var ret PublicObjectId
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *PublicDefaultAssociationMultiPost) GetToOk() (*PublicObjectId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *PublicDefaultAssociationMultiPost) SetTo(v PublicObjectId) {
	o.To = v
}

func (o PublicDefaultAssociationMultiPost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["from"] = o.From
	}
	if true {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullablePublicDefaultAssociationMultiPost struct {
	value *PublicDefaultAssociationMultiPost
	isSet bool
}

func (v NullablePublicDefaultAssociationMultiPost) Get() *PublicDefaultAssociationMultiPost {
	return v.value
}

func (v *NullablePublicDefaultAssociationMultiPost) Set(val *PublicDefaultAssociationMultiPost) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicDefaultAssociationMultiPost) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicDefaultAssociationMultiPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicDefaultAssociationMultiPost(val *PublicDefaultAssociationMultiPost) *NullablePublicDefaultAssociationMultiPost {
	return &NullablePublicDefaultAssociationMultiPost{value: val, isSet: true}
}

func (v NullablePublicDefaultAssociationMultiPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicDefaultAssociationMultiPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
