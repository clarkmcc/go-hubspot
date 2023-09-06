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

// BlogPostLanguageCloneRequestVNext Request body object for creating new blog post language variant.
type BlogPostLanguageCloneRequestVNext struct {
	// ID of blog post to clone.
	Id string `json:"id"`
	// Target language of new variant.
	Language *string `json:"language,omitempty"`
}

// NewBlogPostLanguageCloneRequestVNext instantiates a new BlogPostLanguageCloneRequestVNext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlogPostLanguageCloneRequestVNext(id string) *BlogPostLanguageCloneRequestVNext {
	this := BlogPostLanguageCloneRequestVNext{}
	this.Id = id
	return &this
}

// NewBlogPostLanguageCloneRequestVNextWithDefaults instantiates a new BlogPostLanguageCloneRequestVNext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlogPostLanguageCloneRequestVNextWithDefaults() *BlogPostLanguageCloneRequestVNext {
	this := BlogPostLanguageCloneRequestVNext{}
	return &this
}

// GetId returns the Id field value
func (o *BlogPostLanguageCloneRequestVNext) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BlogPostLanguageCloneRequestVNext) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BlogPostLanguageCloneRequestVNext) SetId(v string) {
	o.Id = v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *BlogPostLanguageCloneRequestVNext) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlogPostLanguageCloneRequestVNext) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *BlogPostLanguageCloneRequestVNext) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *BlogPostLanguageCloneRequestVNext) SetLanguage(v string) {
	o.Language = &v
}

func (o BlogPostLanguageCloneRequestVNext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	return json.Marshal(toSerialize)
}

type NullableBlogPostLanguageCloneRequestVNext struct {
	value *BlogPostLanguageCloneRequestVNext
	isSet bool
}

func (v NullableBlogPostLanguageCloneRequestVNext) Get() *BlogPostLanguageCloneRequestVNext {
	return v.value
}

func (v *NullableBlogPostLanguageCloneRequestVNext) Set(val *BlogPostLanguageCloneRequestVNext) {
	v.value = val
	v.isSet = true
}

func (v NullableBlogPostLanguageCloneRequestVNext) IsSet() bool {
	return v.isSet
}

func (v *NullableBlogPostLanguageCloneRequestVNext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlogPostLanguageCloneRequestVNext(val *BlogPostLanguageCloneRequestVNext) *NullableBlogPostLanguageCloneRequestVNext {
	return &NullableBlogPostLanguageCloneRequestVNext{value: val, isSet: true}
}

func (v NullableBlogPostLanguageCloneRequestVNext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlogPostLanguageCloneRequestVNext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
