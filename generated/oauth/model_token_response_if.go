/*
OAuthService

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oauth

import (
	"encoding/json"
)

// TokenResponseIF struct for TokenResponseIF
type TokenResponseIF struct {
	RefreshToken string  `json:"refresh_token"`
	ExpiresIn    int32   `json:"expires_in"`
	AccessToken  string  `json:"access_token"`
	IdToken      *string `json:"id_token,omitempty"`
	TokenType    string  `json:"token_type"`
}

// NewTokenResponseIF instantiates a new TokenResponseIF object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenResponseIF(refreshToken string, expiresIn int32, accessToken string, tokenType string) *TokenResponseIF {
	this := TokenResponseIF{}
	this.RefreshToken = refreshToken
	this.ExpiresIn = expiresIn
	this.AccessToken = accessToken
	this.TokenType = tokenType
	return &this
}

// NewTokenResponseIFWithDefaults instantiates a new TokenResponseIF object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenResponseIFWithDefaults() *TokenResponseIF {
	this := TokenResponseIF{}
	return &this
}

// GetRefreshToken returns the RefreshToken field value
func (o *TokenResponseIF) GetRefreshToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value
// and a boolean to check if the value has been set.
func (o *TokenResponseIF) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefreshToken, true
}

// SetRefreshToken sets field value
func (o *TokenResponseIF) SetRefreshToken(v string) {
	o.RefreshToken = v
}

// GetExpiresIn returns the ExpiresIn field value
func (o *TokenResponseIF) GetExpiresIn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value
// and a boolean to check if the value has been set.
func (o *TokenResponseIF) GetExpiresInOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresIn, true
}

// SetExpiresIn sets field value
func (o *TokenResponseIF) SetExpiresIn(v int32) {
	o.ExpiresIn = v
}

// GetAccessToken returns the AccessToken field value
func (o *TokenResponseIF) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *TokenResponseIF) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *TokenResponseIF) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetIdToken returns the IdToken field value if set, zero value otherwise.
func (o *TokenResponseIF) GetIdToken() string {
	if o == nil || o.IdToken == nil {
		var ret string
		return ret
	}
	return *o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseIF) GetIdTokenOk() (*string, bool) {
	if o == nil || o.IdToken == nil {
		return nil, false
	}
	return o.IdToken, true
}

// HasIdToken returns a boolean if a field has been set.
func (o *TokenResponseIF) HasIdToken() bool {
	if o != nil && o.IdToken != nil {
		return true
	}

	return false
}

// SetIdToken gets a reference to the given string and assigns it to the IdToken field.
func (o *TokenResponseIF) SetIdToken(v string) {
	o.IdToken = &v
}

// GetTokenType returns the TokenType field value
func (o *TokenResponseIF) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *TokenResponseIF) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *TokenResponseIF) SetTokenType(v string) {
	o.TokenType = v
}

func (o TokenResponseIF) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	if true {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.IdToken != nil {
		toSerialize["id_token"] = o.IdToken
	}
	if true {
		toSerialize["token_type"] = o.TokenType
	}
	return json.Marshal(toSerialize)
}

type NullableTokenResponseIF struct {
	value *TokenResponseIF
	isSet bool
}

func (v NullableTokenResponseIF) Get() *TokenResponseIF {
	return v.value
}

func (v *NullableTokenResponseIF) Set(val *TokenResponseIF) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenResponseIF) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenResponseIF) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenResponseIF(val *TokenResponseIF) *NullableTokenResponseIF {
	return &NullableTokenResponseIF{value: val, isSet: true}
}

func (v NullableTokenResponseIF) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenResponseIF) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}