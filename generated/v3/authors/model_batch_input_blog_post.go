/*
Blog Post endpoints

Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authors

import (
	"encoding/json"
)

// BatchInputBlogPost Wrapper for providing an array of blog posts as inputs.
type BatchInputBlogPost struct {
	// Blog posts to input.
	Inputs []BlogPost `json:"inputs"`
}

// NewBatchInputBlogPost instantiates a new BatchInputBlogPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputBlogPost(inputs []BlogPost) *BatchInputBlogPost {
	this := BatchInputBlogPost{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputBlogPostWithDefaults instantiates a new BatchInputBlogPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputBlogPostWithDefaults() *BatchInputBlogPost {
	this := BatchInputBlogPost{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputBlogPost) GetInputs() []BlogPost {
	if o == nil {
		var ret []BlogPost
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputBlogPost) GetInputsOk() ([]BlogPost, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputBlogPost) SetInputs(v []BlogPost) {
	o.Inputs = v
}

func (o BatchInputBlogPost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["inputs"] = o.Inputs
	}
	return json.Marshal(toSerialize)
}

type NullableBatchInputBlogPost struct {
	value *BatchInputBlogPost
	isSet bool
}

func (v NullableBatchInputBlogPost) Get() *BatchInputBlogPost {
	return v.value
}

func (v *NullableBatchInputBlogPost) Set(val *BatchInputBlogPost) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputBlogPost) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputBlogPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputBlogPost(val *BatchInputBlogPost) *NullableBatchInputBlogPost {
	return &NullableBatchInputBlogPost{value: val, isSet: true}
}

func (v NullableBatchInputBlogPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputBlogPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
