/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oauth

import (
	"encoding/json"
)

// RefreshTokenInfoResponse struct for RefreshTokenInfoResponse
type RefreshTokenInfoResponse struct {
	Token     string   `json:"token"`
	User      *string  `json:"user,omitempty"`
	HubDomain *string  `json:"hub_domain,omitempty"`
	Scopes    []string `json:"scopes"`
	HubId     int32    `json:"hub_id"`
	ClientId  string   `json:"client_id"`
	UserId    int32    `json:"user_id"`
	TokenType string   `json:"token_type"`
}

// NewRefreshTokenInfoResponse instantiates a new RefreshTokenInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshTokenInfoResponse(token string, scopes []string, hubId int32, clientId string, userId int32, tokenType string) *RefreshTokenInfoResponse {
	this := RefreshTokenInfoResponse{}
	this.Token = token
	this.Scopes = scopes
	this.HubId = hubId
	this.ClientId = clientId
	this.UserId = userId
	this.TokenType = tokenType
	return &this
}

// NewRefreshTokenInfoResponseWithDefaults instantiates a new RefreshTokenInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshTokenInfoResponseWithDefaults() *RefreshTokenInfoResponse {
	this := RefreshTokenInfoResponse{}
	return &this
}

// GetToken returns the Token field value
func (o *RefreshTokenInfoResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *RefreshTokenInfoResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *RefreshTokenInfoResponse) SetToken(v string) {
	o.Token = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *RefreshTokenInfoResponse) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshTokenInfoResponse) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *RefreshTokenInfoResponse) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *RefreshTokenInfoResponse) SetUser(v string) {
	o.User = &v
}

// GetHubDomain returns the HubDomain field value if set, zero value otherwise.
func (o *RefreshTokenInfoResponse) GetHubDomain() string {
	if o == nil || o.HubDomain == nil {
		var ret string
		return ret
	}
	return *o.HubDomain
}

// GetHubDomainOk returns a tuple with the HubDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshTokenInfoResponse) GetHubDomainOk() (*string, bool) {
	if o == nil || o.HubDomain == nil {
		return nil, false
	}
	return o.HubDomain, true
}

// HasHubDomain returns a boolean if a field has been set.
func (o *RefreshTokenInfoResponse) HasHubDomain() bool {
	if o != nil && o.HubDomain != nil {
		return true
	}

	return false
}

// SetHubDomain gets a reference to the given string and assigns it to the HubDomain field.
func (o *RefreshTokenInfoResponse) SetHubDomain(v string) {
	o.HubDomain = &v
}

// GetScopes returns the Scopes field value
func (o *RefreshTokenInfoResponse) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *RefreshTokenInfoResponse) GetScopesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *RefreshTokenInfoResponse) SetScopes(v []string) {
	o.Scopes = v
}

// GetHubId returns the HubId field value
func (o *RefreshTokenInfoResponse) GetHubId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HubId
}

// GetHubIdOk returns a tuple with the HubId field value
// and a boolean to check if the value has been set.
func (o *RefreshTokenInfoResponse) GetHubIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HubId, true
}

// SetHubId sets field value
func (o *RefreshTokenInfoResponse) SetHubId(v int32) {
	o.HubId = v
}

// GetClientId returns the ClientId field value
func (o *RefreshTokenInfoResponse) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *RefreshTokenInfoResponse) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *RefreshTokenInfoResponse) SetClientId(v string) {
	o.ClientId = v
}

// GetUserId returns the UserId field value
func (o *RefreshTokenInfoResponse) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *RefreshTokenInfoResponse) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *RefreshTokenInfoResponse) SetUserId(v int32) {
	o.UserId = v
}

// GetTokenType returns the TokenType field value
func (o *RefreshTokenInfoResponse) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *RefreshTokenInfoResponse) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *RefreshTokenInfoResponse) SetTokenType(v string) {
	o.TokenType = v
}

func (o RefreshTokenInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["token"] = o.Token
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.HubDomain != nil {
		toSerialize["hub_domain"] = o.HubDomain
	}
	if true {
		toSerialize["scopes"] = o.Scopes
	}
	if true {
		toSerialize["hub_id"] = o.HubId
	}
	if true {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["user_id"] = o.UserId
	}
	if true {
		toSerialize["token_type"] = o.TokenType
	}
	return json.Marshal(toSerialize)
}

type NullableRefreshTokenInfoResponse struct {
	value *RefreshTokenInfoResponse
	isSet bool
}

func (v NullableRefreshTokenInfoResponse) Get() *RefreshTokenInfoResponse {
	return v.value
}

func (v *NullableRefreshTokenInfoResponse) Set(val *RefreshTokenInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshTokenInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshTokenInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshTokenInfoResponse(val *RefreshTokenInfoResponse) *NullableRefreshTokenInfoResponse {
	return &NullableRefreshTokenInfoResponse{value: val, isSet: true}
}

func (v NullableRefreshTokenInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshTokenInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
