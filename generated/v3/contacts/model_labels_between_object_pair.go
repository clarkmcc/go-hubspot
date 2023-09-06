/*
Contacts

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package contacts

import (
	"encoding/json"
)

// LabelsBetweenObjectPair struct for LabelsBetweenObjectPair
type LabelsBetweenObjectPair struct {
	FromObjectTypeId string   `json:"fromObjectTypeId"`
	FromObjectId     int32    `json:"fromObjectId"`
	ToObjectTypeId   string   `json:"toObjectTypeId"`
	ToObjectId       int32    `json:"toObjectId"`
	Labels           []string `json:"labels"`
}

// NewLabelsBetweenObjectPair instantiates a new LabelsBetweenObjectPair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelsBetweenObjectPair(fromObjectTypeId string, fromObjectId int32, toObjectTypeId string, toObjectId int32, labels []string) *LabelsBetweenObjectPair {
	this := LabelsBetweenObjectPair{}
	this.FromObjectTypeId = fromObjectTypeId
	this.FromObjectId = fromObjectId
	this.ToObjectTypeId = toObjectTypeId
	this.ToObjectId = toObjectId
	this.Labels = labels
	return &this
}

// NewLabelsBetweenObjectPairWithDefaults instantiates a new LabelsBetweenObjectPair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelsBetweenObjectPairWithDefaults() *LabelsBetweenObjectPair {
	this := LabelsBetweenObjectPair{}
	return &this
}

// GetFromObjectTypeId returns the FromObjectTypeId field value
func (o *LabelsBetweenObjectPair) GetFromObjectTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromObjectTypeId
}

// GetFromObjectTypeIdOk returns a tuple with the FromObjectTypeId field value
// and a boolean to check if the value has been set.
func (o *LabelsBetweenObjectPair) GetFromObjectTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromObjectTypeId, true
}

// SetFromObjectTypeId sets field value
func (o *LabelsBetweenObjectPair) SetFromObjectTypeId(v string) {
	o.FromObjectTypeId = v
}

// GetFromObjectId returns the FromObjectId field value
func (o *LabelsBetweenObjectPair) GetFromObjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FromObjectId
}

// GetFromObjectIdOk returns a tuple with the FromObjectId field value
// and a boolean to check if the value has been set.
func (o *LabelsBetweenObjectPair) GetFromObjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromObjectId, true
}

// SetFromObjectId sets field value
func (o *LabelsBetweenObjectPair) SetFromObjectId(v int32) {
	o.FromObjectId = v
}

// GetToObjectTypeId returns the ToObjectTypeId field value
func (o *LabelsBetweenObjectPair) GetToObjectTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToObjectTypeId
}

// GetToObjectTypeIdOk returns a tuple with the ToObjectTypeId field value
// and a boolean to check if the value has been set.
func (o *LabelsBetweenObjectPair) GetToObjectTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToObjectTypeId, true
}

// SetToObjectTypeId sets field value
func (o *LabelsBetweenObjectPair) SetToObjectTypeId(v string) {
	o.ToObjectTypeId = v
}

// GetToObjectId returns the ToObjectId field value
func (o *LabelsBetweenObjectPair) GetToObjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ToObjectId
}

// GetToObjectIdOk returns a tuple with the ToObjectId field value
// and a boolean to check if the value has been set.
func (o *LabelsBetweenObjectPair) GetToObjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToObjectId, true
}

// SetToObjectId sets field value
func (o *LabelsBetweenObjectPair) SetToObjectId(v int32) {
	o.ToObjectId = v
}

// GetLabels returns the Labels field value
func (o *LabelsBetweenObjectPair) GetLabels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *LabelsBetweenObjectPair) GetLabelsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *LabelsBetweenObjectPair) SetLabels(v []string) {
	o.Labels = v
}

func (o LabelsBetweenObjectPair) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fromObjectTypeId"] = o.FromObjectTypeId
	}
	if true {
		toSerialize["fromObjectId"] = o.FromObjectId
	}
	if true {
		toSerialize["toObjectTypeId"] = o.ToObjectTypeId
	}
	if true {
		toSerialize["toObjectId"] = o.ToObjectId
	}
	if true {
		toSerialize["labels"] = o.Labels
	}
	return json.Marshal(toSerialize)
}

type NullableLabelsBetweenObjectPair struct {
	value *LabelsBetweenObjectPair
	isSet bool
}

func (v NullableLabelsBetweenObjectPair) Get() *LabelsBetweenObjectPair {
	return v.value
}

func (v *NullableLabelsBetweenObjectPair) Set(val *LabelsBetweenObjectPair) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelsBetweenObjectPair) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelsBetweenObjectPair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelsBetweenObjectPair(val *LabelsBetweenObjectPair) *NullableLabelsBetweenObjectPair {
	return &NullableLabelsBetweenObjectPair{value: val, isSet: true}
}

func (v NullableLabelsBetweenObjectPair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelsBetweenObjectPair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
