/*
Custom Workflow Actions

Create custom workflow actions

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package actions

import (
	"encoding/json"
	"time"
)

// ActionRevision A revision of this custom action.
type ActionRevision struct {
	Definition ExtensionActionDefinition `json:"definition"`
	// The date the revision was created.
	CreatedAt time.Time `json:"createdAt"`
	Id        string    `json:"id"`
	// The version number of the custom action.
	RevisionId string `json:"revisionId"`
}

// NewActionRevision instantiates a new ActionRevision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionRevision(definition ExtensionActionDefinition, createdAt time.Time, id string, revisionId string) *ActionRevision {
	this := ActionRevision{}
	this.Definition = definition
	this.CreatedAt = createdAt
	this.Id = id
	this.RevisionId = revisionId
	return &this
}

// NewActionRevisionWithDefaults instantiates a new ActionRevision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionRevisionWithDefaults() *ActionRevision {
	this := ActionRevision{}
	return &this
}

// GetDefinition returns the Definition field value
func (o *ActionRevision) GetDefinition() ExtensionActionDefinition {
	if o == nil {
		var ret ExtensionActionDefinition
		return ret
	}

	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *ActionRevision) GetDefinitionOk() (*ExtensionActionDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value
func (o *ActionRevision) SetDefinition(v ExtensionActionDefinition) {
	o.Definition = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ActionRevision) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ActionRevision) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ActionRevision) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *ActionRevision) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ActionRevision) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ActionRevision) SetId(v string) {
	o.Id = v
}

// GetRevisionId returns the RevisionId field value
func (o *ActionRevision) GetRevisionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RevisionId
}

// GetRevisionIdOk returns a tuple with the RevisionId field value
// and a boolean to check if the value has been set.
func (o *ActionRevision) GetRevisionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevisionId, true
}

// SetRevisionId sets field value
func (o *ActionRevision) SetRevisionId(v string) {
	o.RevisionId = v
}

func (o ActionRevision) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["definition"] = o.Definition
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["revisionId"] = o.RevisionId
	}
	return json.Marshal(toSerialize)
}

type NullableActionRevision struct {
	value *ActionRevision
	isSet bool
}

func (v NullableActionRevision) Get() *ActionRevision {
	return v.value
}

func (v *NullableActionRevision) Set(val *ActionRevision) {
	v.value = val
	v.isSet = true
}

func (v NullableActionRevision) IsSet() bool {
	return v.isSet
}

func (v *NullableActionRevision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionRevision(val *ActionRevision) *NullableActionRevision {
	return &NullableActionRevision{value: val, isSet: true}
}

func (v NullableActionRevision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionRevision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}