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

// DomainCdnConfig struct for DomainCdnConfig
type DomainCdnConfig struct {
	PortalId                int32  `json:"portalId"`
	Created                 int64  `json:"created"`
	Updated                 int64  `json:"updated"`
	DomainId                int64  `json:"domainId"`
	CertId                  int64  `json:"certId"`
	CertExpirationDate      int64  `json:"certExpirationDate"`
	CdnId                   string `json:"cdnId"`
	CdnGroupId              string `json:"cdnGroupId"`
	SslStatus               string `json:"sslStatus"`
	ValidationMethod        string `json:"validationMethod"`
	HttpBody                string `json:"httpBody"`
	HttpUrlPath             string `json:"httpUrlPath"`
	Cname                   string `json:"cname"`
	CnameTarget             string `json:"cnameTarget"`
	MinimumTlsVersion       string `json:"minimumTlsVersion"`
	AlternateOriginHostname string `json:"alternateOriginHostname"`
	Id                      int64  `json:"id"`
}

// NewDomainCdnConfig instantiates a new DomainCdnConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainCdnConfig(portalId int32, created int64, updated int64, domainId int64, certId int64, certExpirationDate int64, cdnId string, cdnGroupId string, sslStatus string, validationMethod string, httpBody string, httpUrlPath string, cname string, cnameTarget string, minimumTlsVersion string, alternateOriginHostname string, id int64) *DomainCdnConfig {
	this := DomainCdnConfig{}
	this.PortalId = portalId
	this.Created = created
	this.Updated = updated
	this.DomainId = domainId
	this.CertId = certId
	this.CertExpirationDate = certExpirationDate
	this.CdnId = cdnId
	this.CdnGroupId = cdnGroupId
	this.SslStatus = sslStatus
	this.ValidationMethod = validationMethod
	this.HttpBody = httpBody
	this.HttpUrlPath = httpUrlPath
	this.Cname = cname
	this.CnameTarget = cnameTarget
	this.MinimumTlsVersion = minimumTlsVersion
	this.AlternateOriginHostname = alternateOriginHostname
	this.Id = id
	return &this
}

// NewDomainCdnConfigWithDefaults instantiates a new DomainCdnConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainCdnConfigWithDefaults() *DomainCdnConfig {
	this := DomainCdnConfig{}
	return &this
}

// GetPortalId returns the PortalId field value
func (o *DomainCdnConfig) GetPortalId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PortalId
}

// GetPortalIdOk returns a tuple with the PortalId field value
// and a boolean to check if the value has been set.
func (o *DomainCdnConfig) GetPortalIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PortalId, true
}

// SetPortalId sets field value
func (o *DomainCdnConfig) SetPortalId(v int32) {
	o.PortalId = v
}

// GetCreated returns the Created field value
func (o *DomainCdnConfig) GetCreated() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *DomainCdnConfig) GetCreatedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *DomainCdnConfig) SetCreated(v int64) {
	o.Created = v
}

// GetUpdated returns the Updated field value
func (o *DomainCdnConfig) GetUpdated() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *DomainCdnConfig) GetUpdatedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *DomainCdnConfig) SetUpdated(v int64) {
	o.Updated = v
}

// GetDomainId returns the DomainId field value
func (o *DomainCdnConfig) GetDomainId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value
// and a boolean to check if the value has been set.
func (o *DomainCdnConfig) GetDomainIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DomainId, true
}

// SetDomainId sets field value
func (o *DomainCdnConfig) SetDomainId(v int64) {
	o.DomainId = v
}

// GetCertId returns the CertId field value
func (o *DomainCdnConfig) GetCertId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CertId
}

// GetCertIdOk returns a tuple with the CertId field value
// and a boolean to check if the value has been set.
func (o *DomainCdnConfig) GetCertIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertId, true
}

// SetCertId sets field value
func (o *DomainCdnConfig) SetCertId(v int64) {
	o.CertId = v
}

// GetCertExpirationDate returns the CertExpirationDate field value
func (o *DomainCdnConfig) GetCertExpirationDate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CertExpirationDate
}

// GetCertExpirationDateOk returns a tuple with the CertExpirationDate field value
// and a boolean to check if the value has been set.
func (o *DomainCdnConfig) GetCertExpirationDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertExpirationDate, true
}

// SetCertExpirationDate sets field value
func (o *DomainCdnConfig) SetCertExpirationDate(v int64) {
	o.CertExpirationDate = v
}

// GetCdnId returns the CdnId field value
func (o *DomainCdnConfig) GetCdnId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CdnId
}

// GetCdnIdOk returns a tuple with the CdnId field value
// and a boolean to check if the value has been set.
func (o *DomainCdnConfig) GetCdnIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CdnId, true
}

// SetCdnId sets field value
func (o *DomainCdnConfig) SetCdnId(v string) {
	o.CdnId = v
}

// GetCdnGroupId returns the CdnGroupId field value
func (o *DomainCdnConfig) GetCdnGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CdnGroupId
}

// GetCdnGroupIdOk returns a tuple with the CdnGroupId field value
// and a boolean to check if the value has been set.
func (o *DomainCdnConfig) GetCdnGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CdnGroupId, true
}

// SetCdnGroupId sets field value
func (o *DomainCdnConfig) SetCdnGroupId(v string) {
	o.CdnGroupId = v
}

// GetSslStatus returns the SslStatus field value
func (o *DomainCdnConfig) GetSslStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SslStatus
}

// GetSslStatusOk returns a tuple with the SslStatus field value
// and a boolean to check if the value has been set.
func (o *DomainCdnConfig) GetSslStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SslStatus, true
}

// SetSslStatus sets field value
func (o *DomainCdnConfig) SetSslStatus(v string) {
	o.SslStatus = v
}

// GetValidationMethod returns the ValidationMethod field value
func (o *DomainCdnConfig) GetValidationMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidationMethod
}

// GetValidationMethodOk returns a tuple with the ValidationMethod field value
// and a boolean to check if the value has been set.
func (o *DomainCdnConfig) GetValidationMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidationMethod, true
}

// SetValidationMethod sets field value
func (o *DomainCdnConfig) SetValidationMethod(v string) {
	o.ValidationMethod = v
}

// GetHttpBody returns the HttpBody field value
func (o *DomainCdnConfig) GetHttpBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HttpBody
}

// GetHttpBodyOk returns a tuple with the HttpBody field value
// and a boolean to check if the value has been set.
func (o *DomainCdnConfig) GetHttpBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpBody, true
}

// SetHttpBody sets field value
func (o *DomainCdnConfig) SetHttpBody(v string) {
	o.HttpBody = v
}

// GetHttpUrlPath returns the HttpUrlPath field value
func (o *DomainCdnConfig) GetHttpUrlPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HttpUrlPath
}

// GetHttpUrlPathOk returns a tuple with the HttpUrlPath field value
// and a boolean to check if the value has been set.
func (o *DomainCdnConfig) GetHttpUrlPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpUrlPath, true
}

// SetHttpUrlPath sets field value
func (o *DomainCdnConfig) SetHttpUrlPath(v string) {
	o.HttpUrlPath = v
}

// GetCname returns the Cname field value
func (o *DomainCdnConfig) GetCname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cname
}

// GetCnameOk returns a tuple with the Cname field value
// and a boolean to check if the value has been set.
func (o *DomainCdnConfig) GetCnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cname, true
}

// SetCname sets field value
func (o *DomainCdnConfig) SetCname(v string) {
	o.Cname = v
}

// GetCnameTarget returns the CnameTarget field value
func (o *DomainCdnConfig) GetCnameTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CnameTarget
}

// GetCnameTargetOk returns a tuple with the CnameTarget field value
// and a boolean to check if the value has been set.
func (o *DomainCdnConfig) GetCnameTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CnameTarget, true
}

// SetCnameTarget sets field value
func (o *DomainCdnConfig) SetCnameTarget(v string) {
	o.CnameTarget = v
}

// GetMinimumTlsVersion returns the MinimumTlsVersion field value
func (o *DomainCdnConfig) GetMinimumTlsVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinimumTlsVersion
}

// GetMinimumTlsVersionOk returns a tuple with the MinimumTlsVersion field value
// and a boolean to check if the value has been set.
func (o *DomainCdnConfig) GetMinimumTlsVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinimumTlsVersion, true
}

// SetMinimumTlsVersion sets field value
func (o *DomainCdnConfig) SetMinimumTlsVersion(v string) {
	o.MinimumTlsVersion = v
}

// GetAlternateOriginHostname returns the AlternateOriginHostname field value
func (o *DomainCdnConfig) GetAlternateOriginHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlternateOriginHostname
}

// GetAlternateOriginHostnameOk returns a tuple with the AlternateOriginHostname field value
// and a boolean to check if the value has been set.
func (o *DomainCdnConfig) GetAlternateOriginHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlternateOriginHostname, true
}

// SetAlternateOriginHostname sets field value
func (o *DomainCdnConfig) SetAlternateOriginHostname(v string) {
	o.AlternateOriginHostname = v
}

// GetId returns the Id field value
func (o *DomainCdnConfig) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DomainCdnConfig) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DomainCdnConfig) SetId(v int64) {
	o.Id = v
}

func (o DomainCdnConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["portalId"] = o.PortalId
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["updated"] = o.Updated
	}
	if true {
		toSerialize["domainId"] = o.DomainId
	}
	if true {
		toSerialize["certId"] = o.CertId
	}
	if true {
		toSerialize["certExpirationDate"] = o.CertExpirationDate
	}
	if true {
		toSerialize["cdnId"] = o.CdnId
	}
	if true {
		toSerialize["cdnGroupId"] = o.CdnGroupId
	}
	if true {
		toSerialize["sslStatus"] = o.SslStatus
	}
	if true {
		toSerialize["validationMethod"] = o.ValidationMethod
	}
	if true {
		toSerialize["httpBody"] = o.HttpBody
	}
	if true {
		toSerialize["httpUrlPath"] = o.HttpUrlPath
	}
	if true {
		toSerialize["cname"] = o.Cname
	}
	if true {
		toSerialize["cnameTarget"] = o.CnameTarget
	}
	if true {
		toSerialize["minimumTlsVersion"] = o.MinimumTlsVersion
	}
	if true {
		toSerialize["alternateOriginHostname"] = o.AlternateOriginHostname
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableDomainCdnConfig struct {
	value *DomainCdnConfig
	isSet bool
}

func (v NullableDomainCdnConfig) Get() *DomainCdnConfig {
	return v.value
}

func (v *NullableDomainCdnConfig) Set(val *DomainCdnConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainCdnConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainCdnConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainCdnConfig(val *DomainCdnConfig) *NullableDomainCdnConfig {
	return &NullableDomainCdnConfig{value: val, isSet: true}
}

func (v NullableDomainCdnConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainCdnConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}