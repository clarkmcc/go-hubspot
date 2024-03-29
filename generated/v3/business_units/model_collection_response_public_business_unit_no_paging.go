/*
Business Units

Retrieve Business Unit information.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package business_units

import (
	"encoding/json"
)

// CollectionResponsePublicBusinessUnitNoPaging A response object containing a collection of Business Units
type CollectionResponsePublicBusinessUnitNoPaging struct {
	// The collection of Business Units
	Results []PublicBusinessUnit `json:"results"`
}

// NewCollectionResponsePublicBusinessUnitNoPaging instantiates a new CollectionResponsePublicBusinessUnitNoPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponsePublicBusinessUnitNoPaging(results []PublicBusinessUnit) *CollectionResponsePublicBusinessUnitNoPaging {
	this := CollectionResponsePublicBusinessUnitNoPaging{}
	this.Results = results
	return &this
}

// NewCollectionResponsePublicBusinessUnitNoPagingWithDefaults instantiates a new CollectionResponsePublicBusinessUnitNoPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponsePublicBusinessUnitNoPagingWithDefaults() *CollectionResponsePublicBusinessUnitNoPaging {
	this := CollectionResponsePublicBusinessUnitNoPaging{}
	return &this
}

// GetResults returns the Results field value
func (o *CollectionResponsePublicBusinessUnitNoPaging) GetResults() []PublicBusinessUnit {
	if o == nil {
		var ret []PublicBusinessUnit
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponsePublicBusinessUnitNoPaging) GetResultsOk() ([]PublicBusinessUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *CollectionResponsePublicBusinessUnitNoPaging) SetResults(v []PublicBusinessUnit) {
	o.Results = v
}

func (o CollectionResponsePublicBusinessUnitNoPaging) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionResponsePublicBusinessUnitNoPaging struct {
	value *CollectionResponsePublicBusinessUnitNoPaging
	isSet bool
}

func (v NullableCollectionResponsePublicBusinessUnitNoPaging) Get() *CollectionResponsePublicBusinessUnitNoPaging {
	return v.value
}

func (v *NullableCollectionResponsePublicBusinessUnitNoPaging) Set(val *CollectionResponsePublicBusinessUnitNoPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponsePublicBusinessUnitNoPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponsePublicBusinessUnitNoPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponsePublicBusinessUnitNoPaging(val *CollectionResponsePublicBusinessUnitNoPaging) *NullableCollectionResponsePublicBusinessUnitNoPaging {
	return &NullableCollectionResponsePublicBusinessUnitNoPaging{value: val, isSet: true}
}

func (v NullableCollectionResponsePublicBusinessUnitNoPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponsePublicBusinessUnitNoPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
