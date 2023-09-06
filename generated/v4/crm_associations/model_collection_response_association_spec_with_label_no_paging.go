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

// CollectionResponseAssociationSpecWithLabelNoPaging struct for CollectionResponseAssociationSpecWithLabelNoPaging
type CollectionResponseAssociationSpecWithLabelNoPaging struct {
	Results []AssociationSpecWithLabel `json:"results"`
}

// NewCollectionResponseAssociationSpecWithLabelNoPaging instantiates a new CollectionResponseAssociationSpecWithLabelNoPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponseAssociationSpecWithLabelNoPaging(results []AssociationSpecWithLabel) *CollectionResponseAssociationSpecWithLabelNoPaging {
	this := CollectionResponseAssociationSpecWithLabelNoPaging{}
	this.Results = results
	return &this
}

// NewCollectionResponseAssociationSpecWithLabelNoPagingWithDefaults instantiates a new CollectionResponseAssociationSpecWithLabelNoPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponseAssociationSpecWithLabelNoPagingWithDefaults() *CollectionResponseAssociationSpecWithLabelNoPaging {
	this := CollectionResponseAssociationSpecWithLabelNoPaging{}
	return &this
}

// GetResults returns the Results field value
func (o *CollectionResponseAssociationSpecWithLabelNoPaging) GetResults() []AssociationSpecWithLabel {
	if o == nil {
		var ret []AssociationSpecWithLabel
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseAssociationSpecWithLabelNoPaging) GetResultsOk() ([]AssociationSpecWithLabel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *CollectionResponseAssociationSpecWithLabelNoPaging) SetResults(v []AssociationSpecWithLabel) {
	o.Results = v
}

func (o CollectionResponseAssociationSpecWithLabelNoPaging) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionResponseAssociationSpecWithLabelNoPaging struct {
	value *CollectionResponseAssociationSpecWithLabelNoPaging
	isSet bool
}

func (v NullableCollectionResponseAssociationSpecWithLabelNoPaging) Get() *CollectionResponseAssociationSpecWithLabelNoPaging {
	return v.value
}

func (v *NullableCollectionResponseAssociationSpecWithLabelNoPaging) Set(val *CollectionResponseAssociationSpecWithLabelNoPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponseAssociationSpecWithLabelNoPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponseAssociationSpecWithLabelNoPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponseAssociationSpecWithLabelNoPaging(val *CollectionResponseAssociationSpecWithLabelNoPaging) *NullableCollectionResponseAssociationSpecWithLabelNoPaging {
	return &NullableCollectionResponseAssociationSpecWithLabelNoPaging{value: val, isSet: true}
}

func (v NullableCollectionResponseAssociationSpecWithLabelNoPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponseAssociationSpecWithLabelNoPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
