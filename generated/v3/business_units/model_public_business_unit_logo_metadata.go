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

// PublicBusinessUnitLogoMetadata A Business Unit's logo metadata
type PublicBusinessUnitLogoMetadata struct {
	// The logo's alt text
	LogoAltText *string `json:"logoAltText,omitempty"`
	// The logo's resized url
	ResizedUrl *string `json:"resizedUrl,omitempty"`
	// The logo's url
	LogoUrl *string `json:"logoUrl,omitempty"`
}

// NewPublicBusinessUnitLogoMetadata instantiates a new PublicBusinessUnitLogoMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicBusinessUnitLogoMetadata() *PublicBusinessUnitLogoMetadata {
	this := PublicBusinessUnitLogoMetadata{}
	return &this
}

// NewPublicBusinessUnitLogoMetadataWithDefaults instantiates a new PublicBusinessUnitLogoMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicBusinessUnitLogoMetadataWithDefaults() *PublicBusinessUnitLogoMetadata {
	this := PublicBusinessUnitLogoMetadata{}
	return &this
}

// GetLogoAltText returns the LogoAltText field value if set, zero value otherwise.
func (o *PublicBusinessUnitLogoMetadata) GetLogoAltText() string {
	if o == nil || o.LogoAltText == nil {
		var ret string
		return ret
	}
	return *o.LogoAltText
}

// GetLogoAltTextOk returns a tuple with the LogoAltText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicBusinessUnitLogoMetadata) GetLogoAltTextOk() (*string, bool) {
	if o == nil || o.LogoAltText == nil {
		return nil, false
	}
	return o.LogoAltText, true
}

// HasLogoAltText returns a boolean if a field has been set.
func (o *PublicBusinessUnitLogoMetadata) HasLogoAltText() bool {
	if o != nil && o.LogoAltText != nil {
		return true
	}

	return false
}

// SetLogoAltText gets a reference to the given string and assigns it to the LogoAltText field.
func (o *PublicBusinessUnitLogoMetadata) SetLogoAltText(v string) {
	o.LogoAltText = &v
}

// GetResizedUrl returns the ResizedUrl field value if set, zero value otherwise.
func (o *PublicBusinessUnitLogoMetadata) GetResizedUrl() string {
	if o == nil || o.ResizedUrl == nil {
		var ret string
		return ret
	}
	return *o.ResizedUrl
}

// GetResizedUrlOk returns a tuple with the ResizedUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicBusinessUnitLogoMetadata) GetResizedUrlOk() (*string, bool) {
	if o == nil || o.ResizedUrl == nil {
		return nil, false
	}
	return o.ResizedUrl, true
}

// HasResizedUrl returns a boolean if a field has been set.
func (o *PublicBusinessUnitLogoMetadata) HasResizedUrl() bool {
	if o != nil && o.ResizedUrl != nil {
		return true
	}

	return false
}

// SetResizedUrl gets a reference to the given string and assigns it to the ResizedUrl field.
func (o *PublicBusinessUnitLogoMetadata) SetResizedUrl(v string) {
	o.ResizedUrl = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *PublicBusinessUnitLogoMetadata) GetLogoUrl() string {
	if o == nil || o.LogoUrl == nil {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicBusinessUnitLogoMetadata) GetLogoUrlOk() (*string, bool) {
	if o == nil || o.LogoUrl == nil {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *PublicBusinessUnitLogoMetadata) HasLogoUrl() bool {
	if o != nil && o.LogoUrl != nil {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *PublicBusinessUnitLogoMetadata) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

func (o PublicBusinessUnitLogoMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LogoAltText != nil {
		toSerialize["logoAltText"] = o.LogoAltText
	}
	if o.ResizedUrl != nil {
		toSerialize["resizedUrl"] = o.ResizedUrl
	}
	if o.LogoUrl != nil {
		toSerialize["logoUrl"] = o.LogoUrl
	}
	return json.Marshal(toSerialize)
}

type NullablePublicBusinessUnitLogoMetadata struct {
	value *PublicBusinessUnitLogoMetadata
	isSet bool
}

func (v NullablePublicBusinessUnitLogoMetadata) Get() *PublicBusinessUnitLogoMetadata {
	return v.value
}

func (v *NullablePublicBusinessUnitLogoMetadata) Set(val *PublicBusinessUnitLogoMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicBusinessUnitLogoMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicBusinessUnitLogoMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicBusinessUnitLogoMetadata(val *PublicBusinessUnitLogoMetadata) *NullablePublicBusinessUnitLogoMetadata {
	return &NullablePublicBusinessUnitLogoMetadata{value: val, isSet: true}
}

func (v NullablePublicBusinessUnitLogoMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicBusinessUnitLogoMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
