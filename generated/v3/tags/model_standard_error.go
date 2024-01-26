/*
Tags

Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tags

import (
	"encoding/json"
)

// StandardError Model definition for a standard error.
type StandardError struct {
	// Error subcategory.
	SubCategory map[string]interface{} `json:"subCategory,omitempty"`
	// Error context.
	Context map[string][]string `json:"context"`
	// Error links.
	Links map[string]string `json:"links"`
	// Error ID.
	Id *string `json:"id,omitempty"`
	// Error category.
	Category string `json:"category"`
	// Error message.
	Message string `json:"message"`
	// List of error details.
	Errors []ErrorDetail `json:"errors"`
	// Error status.
	Status string `json:"status"`
}

// NewStandardError instantiates a new StandardError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardError(context map[string][]string, links map[string]string, category string, message string, errors []ErrorDetail, status string) *StandardError {
	this := StandardError{}
	this.Context = context
	this.Links = links
	this.Category = category
	this.Message = message
	this.Errors = errors
	this.Status = status
	return &this
}

// NewStandardErrorWithDefaults instantiates a new StandardError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardErrorWithDefaults() *StandardError {
	this := StandardError{}
	return &this
}

// GetSubCategory returns the SubCategory field value if set, zero value otherwise.
func (o *StandardError) GetSubCategory() map[string]interface{} {
	if o == nil || o.SubCategory == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.SubCategory
}

// GetSubCategoryOk returns a tuple with the SubCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardError) GetSubCategoryOk() (map[string]interface{}, bool) {
	if o == nil || o.SubCategory == nil {
		return nil, false
	}
	return o.SubCategory, true
}

// HasSubCategory returns a boolean if a field has been set.
func (o *StandardError) HasSubCategory() bool {
	if o != nil && o.SubCategory != nil {
		return true
	}

	return false
}

// SetSubCategory gets a reference to the given map[string]interface{} and assigns it to the SubCategory field.
func (o *StandardError) SetSubCategory(v map[string]interface{}) {
	o.SubCategory = v
}

// GetContext returns the Context field value
func (o *StandardError) GetContext() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *StandardError) GetContextOk() (*map[string][]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *StandardError) SetContext(v map[string][]string) {
	o.Context = v
}

// GetLinks returns the Links field value
func (o *StandardError) GetLinks() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *StandardError) GetLinksOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *StandardError) SetLinks(v map[string]string) {
	o.Links = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StandardError) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardError) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StandardError) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StandardError) SetId(v string) {
	o.Id = &v
}

// GetCategory returns the Category field value
func (o *StandardError) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *StandardError) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *StandardError) SetCategory(v string) {
	o.Category = v
}

// GetMessage returns the Message field value
func (o *StandardError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *StandardError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *StandardError) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value
func (o *StandardError) GetErrors() []ErrorDetail {
	if o == nil {
		var ret []ErrorDetail
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *StandardError) GetErrorsOk() ([]ErrorDetail, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *StandardError) SetErrors(v []ErrorDetail) {
	o.Errors = v
}

// GetStatus returns the Status field value
func (o *StandardError) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *StandardError) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *StandardError) SetStatus(v string) {
	o.Status = v
}

func (o StandardError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SubCategory != nil {
		toSerialize["subCategory"] = o.SubCategory
	}
	if true {
		toSerialize["context"] = o.Context
	}
	if true {
		toSerialize["links"] = o.Links
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["category"] = o.Category
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["errors"] = o.Errors
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableStandardError struct {
	value *StandardError
	isSet bool
}

func (v NullableStandardError) Get() *StandardError {
	return v.value
}

func (v *NullableStandardError) Set(val *StandardError) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardError) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardError(val *StandardError) *NullableStandardError {
	return &NullableStandardError{value: val, isSet: true}
}

func (v NullableStandardError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandardError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
