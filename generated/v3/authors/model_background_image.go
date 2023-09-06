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

// BackgroundImage struct for BackgroundImage
type BackgroundImage struct {
	ImageUrl           string `json:"imageUrl"`
	BackgroundSize     string `json:"backgroundSize"`
	BackgroundPosition string `json:"backgroundPosition"`
}

// NewBackgroundImage instantiates a new BackgroundImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackgroundImage(imageUrl string, backgroundSize string, backgroundPosition string) *BackgroundImage {
	this := BackgroundImage{}
	this.ImageUrl = imageUrl
	this.BackgroundSize = backgroundSize
	this.BackgroundPosition = backgroundPosition
	return &this
}

// NewBackgroundImageWithDefaults instantiates a new BackgroundImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackgroundImageWithDefaults() *BackgroundImage {
	this := BackgroundImage{}
	return &this
}

// GetImageUrl returns the ImageUrl field value
func (o *BackgroundImage) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *BackgroundImage) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *BackgroundImage) SetImageUrl(v string) {
	o.ImageUrl = v
}

// GetBackgroundSize returns the BackgroundSize field value
func (o *BackgroundImage) GetBackgroundSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BackgroundSize
}

// GetBackgroundSizeOk returns a tuple with the BackgroundSize field value
// and a boolean to check if the value has been set.
func (o *BackgroundImage) GetBackgroundSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackgroundSize, true
}

// SetBackgroundSize sets field value
func (o *BackgroundImage) SetBackgroundSize(v string) {
	o.BackgroundSize = v
}

// GetBackgroundPosition returns the BackgroundPosition field value
func (o *BackgroundImage) GetBackgroundPosition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BackgroundPosition
}

// GetBackgroundPositionOk returns a tuple with the BackgroundPosition field value
// and a boolean to check if the value has been set.
func (o *BackgroundImage) GetBackgroundPositionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackgroundPosition, true
}

// SetBackgroundPosition sets field value
func (o *BackgroundImage) SetBackgroundPosition(v string) {
	o.BackgroundPosition = v
}

func (o BackgroundImage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if true {
		toSerialize["backgroundSize"] = o.BackgroundSize
	}
	if true {
		toSerialize["backgroundPosition"] = o.BackgroundPosition
	}
	return json.Marshal(toSerialize)
}

type NullableBackgroundImage struct {
	value *BackgroundImage
	isSet bool
}

func (v NullableBackgroundImage) Get() *BackgroundImage {
	return v.value
}

func (v *NullableBackgroundImage) Set(val *BackgroundImage) {
	v.value = val
	v.isSet = true
}

func (v NullableBackgroundImage) IsSet() bool {
	return v.isSet
}

func (v *NullableBackgroundImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackgroundImage(val *BackgroundImage) *NullableBackgroundImage {
	return &NullableBackgroundImage{value: val, isSet: true}
}

func (v NullableBackgroundImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackgroundImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
