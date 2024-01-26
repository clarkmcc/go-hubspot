/*
Posts

Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blog_posts

import (
	"encoding/json"
)

// RowMetaData struct for RowMetaData
type RowMetaData struct {
	CssClass string `json:"cssClass"`
	Styles   Styles `json:"styles"`
}

// NewRowMetaData instantiates a new RowMetaData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRowMetaData(cssClass string, styles Styles) *RowMetaData {
	this := RowMetaData{}
	this.CssClass = cssClass
	this.Styles = styles
	return &this
}

// NewRowMetaDataWithDefaults instantiates a new RowMetaData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRowMetaDataWithDefaults() *RowMetaData {
	this := RowMetaData{}
	return &this
}

// GetCssClass returns the CssClass field value
func (o *RowMetaData) GetCssClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CssClass
}

// GetCssClassOk returns a tuple with the CssClass field value
// and a boolean to check if the value has been set.
func (o *RowMetaData) GetCssClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CssClass, true
}

// SetCssClass sets field value
func (o *RowMetaData) SetCssClass(v string) {
	o.CssClass = v
}

// GetStyles returns the Styles field value
func (o *RowMetaData) GetStyles() Styles {
	if o == nil {
		var ret Styles
		return ret
	}

	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value
// and a boolean to check if the value has been set.
func (o *RowMetaData) GetStylesOk() (*Styles, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Styles, true
}

// SetStyles sets field value
func (o *RowMetaData) SetStyles(v Styles) {
	o.Styles = v
}

func (o RowMetaData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cssClass"] = o.CssClass
	}
	if true {
		toSerialize["styles"] = o.Styles
	}
	return json.Marshal(toSerialize)
}

type NullableRowMetaData struct {
	value *RowMetaData
	isSet bool
}

func (v NullableRowMetaData) Get() *RowMetaData {
	return v.value
}

func (v *NullableRowMetaData) Set(val *RowMetaData) {
	v.value = val
	v.isSet = true
}

func (v NullableRowMetaData) IsSet() bool {
	return v.isSet
}

func (v *NullableRowMetaData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRowMetaData(val *RowMetaData) *NullableRowMetaData {
	return &NullableRowMetaData{value: val, isSet: true}
}

func (v NullableRowMetaData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRowMetaData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
