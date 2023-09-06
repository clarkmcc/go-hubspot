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

// CollectionResponseWithTotalBlogPostForwardPaging Response object for collections of blog posts with pagination information.
type CollectionResponseWithTotalBlogPostForwardPaging struct {
	// Total number of blog posts.
	Total int32 `json:"total"`
	// Collection of blog posts.
	Results []BlogPost     `json:"results"`
	Paging  *ForwardPaging `json:"paging,omitempty"`
}

// NewCollectionResponseWithTotalBlogPostForwardPaging instantiates a new CollectionResponseWithTotalBlogPostForwardPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponseWithTotalBlogPostForwardPaging(total int32, results []BlogPost) *CollectionResponseWithTotalBlogPostForwardPaging {
	this := CollectionResponseWithTotalBlogPostForwardPaging{}
	this.Total = total
	this.Results = results
	return &this
}

// NewCollectionResponseWithTotalBlogPostForwardPagingWithDefaults instantiates a new CollectionResponseWithTotalBlogPostForwardPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponseWithTotalBlogPostForwardPagingWithDefaults() *CollectionResponseWithTotalBlogPostForwardPaging {
	this := CollectionResponseWithTotalBlogPostForwardPaging{}
	return &this
}

// GetTotal returns the Total field value
func (o *CollectionResponseWithTotalBlogPostForwardPaging) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseWithTotalBlogPostForwardPaging) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *CollectionResponseWithTotalBlogPostForwardPaging) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *CollectionResponseWithTotalBlogPostForwardPaging) GetResults() []BlogPost {
	if o == nil {
		var ret []BlogPost
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseWithTotalBlogPostForwardPaging) GetResultsOk() ([]BlogPost, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *CollectionResponseWithTotalBlogPostForwardPaging) SetResults(v []BlogPost) {
	o.Results = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *CollectionResponseWithTotalBlogPostForwardPaging) GetPaging() ForwardPaging {
	if o == nil || o.Paging == nil {
		var ret ForwardPaging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponseWithTotalBlogPostForwardPaging) GetPagingOk() (*ForwardPaging, bool) {
	if o == nil || o.Paging == nil {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *CollectionResponseWithTotalBlogPostForwardPaging) HasPaging() bool {
	if o != nil && o.Paging != nil {
		return true
	}

	return false
}

// SetPaging gets a reference to the given ForwardPaging and assigns it to the Paging field.
func (o *CollectionResponseWithTotalBlogPostForwardPaging) SetPaging(v ForwardPaging) {
	o.Paging = &v
}

func (o CollectionResponseWithTotalBlogPostForwardPaging) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["results"] = o.Results
	}
	if o.Paging != nil {
		toSerialize["paging"] = o.Paging
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionResponseWithTotalBlogPostForwardPaging struct {
	value *CollectionResponseWithTotalBlogPostForwardPaging
	isSet bool
}

func (v NullableCollectionResponseWithTotalBlogPostForwardPaging) Get() *CollectionResponseWithTotalBlogPostForwardPaging {
	return v.value
}

func (v *NullableCollectionResponseWithTotalBlogPostForwardPaging) Set(val *CollectionResponseWithTotalBlogPostForwardPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponseWithTotalBlogPostForwardPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponseWithTotalBlogPostForwardPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponseWithTotalBlogPostForwardPaging(val *CollectionResponseWithTotalBlogPostForwardPaging) *NullableCollectionResponseWithTotalBlogPostForwardPaging {
	return &NullableCollectionResponseWithTotalBlogPostForwardPaging{value: val, isSet: true}
}

func (v NullableCollectionResponseWithTotalBlogPostForwardPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponseWithTotalBlogPostForwardPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
