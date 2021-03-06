/*
Domains

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package domains

import (
	"encoding/json"
)

// DomainSetupInfo struct for DomainSetupInfo
type DomainSetupInfo struct {
	SupportsSslExternally bool     `json:"supportsSslExternally"`
	WhoIsEmailAddresses   []string `json:"whoIsEmailAddresses"`
}

// NewDomainSetupInfo instantiates a new DomainSetupInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainSetupInfo(supportsSslExternally bool, whoIsEmailAddresses []string) *DomainSetupInfo {
	this := DomainSetupInfo{}
	this.SupportsSslExternally = supportsSslExternally
	this.WhoIsEmailAddresses = whoIsEmailAddresses
	return &this
}

// NewDomainSetupInfoWithDefaults instantiates a new DomainSetupInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainSetupInfoWithDefaults() *DomainSetupInfo {
	this := DomainSetupInfo{}
	return &this
}

// GetSupportsSslExternally returns the SupportsSslExternally field value
func (o *DomainSetupInfo) GetSupportsSslExternally() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SupportsSslExternally
}

// GetSupportsSslExternallyOk returns a tuple with the SupportsSslExternally field value
// and a boolean to check if the value has been set.
func (o *DomainSetupInfo) GetSupportsSslExternallyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportsSslExternally, true
}

// SetSupportsSslExternally sets field value
func (o *DomainSetupInfo) SetSupportsSslExternally(v bool) {
	o.SupportsSslExternally = v
}

// GetWhoIsEmailAddresses returns the WhoIsEmailAddresses field value
func (o *DomainSetupInfo) GetWhoIsEmailAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.WhoIsEmailAddresses
}

// GetWhoIsEmailAddressesOk returns a tuple with the WhoIsEmailAddresses field value
// and a boolean to check if the value has been set.
func (o *DomainSetupInfo) GetWhoIsEmailAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WhoIsEmailAddresses, true
}

// SetWhoIsEmailAddresses sets field value
func (o *DomainSetupInfo) SetWhoIsEmailAddresses(v []string) {
	o.WhoIsEmailAddresses = v
}

func (o DomainSetupInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["supportsSslExternally"] = o.SupportsSslExternally
	}
	if true {
		toSerialize["whoIsEmailAddresses"] = o.WhoIsEmailAddresses
	}
	return json.Marshal(toSerialize)
}

type NullableDomainSetupInfo struct {
	value *DomainSetupInfo
	isSet bool
}

func (v NullableDomainSetupInfo) Get() *DomainSetupInfo {
	return v.value
}

func (v *NullableDomainSetupInfo) Set(val *DomainSetupInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainSetupInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainSetupInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainSetupInfo(val *DomainSetupInfo) *NullableDomainSetupInfo {
	return &NullableDomainSetupInfo{value: val, isSet: true}
}

func (v NullableDomainSetupInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainSetupInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
