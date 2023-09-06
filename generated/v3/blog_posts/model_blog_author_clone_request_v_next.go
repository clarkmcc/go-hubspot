/*
Blog Post endpoints

Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blog_posts

import (
	"encoding/json"
)

// BlogAuthorCloneRequestVNext Request body object for cloning blog authors.
type BlogAuthorCloneRequestVNext struct {
	// ID of the object to be cloned.
	Id string `json:"id"`
	// Language of newly cloned object.
	Language *string `json:"language,omitempty"`
	// Primary language in multi-language group.
	PrimaryLanguage *string    `json:"primaryLanguage,omitempty"`
	BlogAuthor      BlogAuthor `json:"blogAuthor"`
}

// NewBlogAuthorCloneRequestVNext instantiates a new BlogAuthorCloneRequestVNext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlogAuthorCloneRequestVNext(id string, blogAuthor BlogAuthor) *BlogAuthorCloneRequestVNext {
	this := BlogAuthorCloneRequestVNext{}
	this.Id = id
	this.BlogAuthor = blogAuthor
	return &this
}

// NewBlogAuthorCloneRequestVNextWithDefaults instantiates a new BlogAuthorCloneRequestVNext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlogAuthorCloneRequestVNextWithDefaults() *BlogAuthorCloneRequestVNext {
	this := BlogAuthorCloneRequestVNext{}
	return &this
}

// GetId returns the Id field value
func (o *BlogAuthorCloneRequestVNext) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BlogAuthorCloneRequestVNext) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BlogAuthorCloneRequestVNext) SetId(v string) {
	o.Id = v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *BlogAuthorCloneRequestVNext) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlogAuthorCloneRequestVNext) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *BlogAuthorCloneRequestVNext) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *BlogAuthorCloneRequestVNext) SetLanguage(v string) {
	o.Language = &v
}

// GetPrimaryLanguage returns the PrimaryLanguage field value if set, zero value otherwise.
func (o *BlogAuthorCloneRequestVNext) GetPrimaryLanguage() string {
	if o == nil || o.PrimaryLanguage == nil {
		var ret string
		return ret
	}
	return *o.PrimaryLanguage
}

// GetPrimaryLanguageOk returns a tuple with the PrimaryLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlogAuthorCloneRequestVNext) GetPrimaryLanguageOk() (*string, bool) {
	if o == nil || o.PrimaryLanguage == nil {
		return nil, false
	}
	return o.PrimaryLanguage, true
}

// HasPrimaryLanguage returns a boolean if a field has been set.
func (o *BlogAuthorCloneRequestVNext) HasPrimaryLanguage() bool {
	if o != nil && o.PrimaryLanguage != nil {
		return true
	}

	return false
}

// SetPrimaryLanguage gets a reference to the given string and assigns it to the PrimaryLanguage field.
func (o *BlogAuthorCloneRequestVNext) SetPrimaryLanguage(v string) {
	o.PrimaryLanguage = &v
}

// GetBlogAuthor returns the BlogAuthor field value
func (o *BlogAuthorCloneRequestVNext) GetBlogAuthor() BlogAuthor {
	if o == nil {
		var ret BlogAuthor
		return ret
	}

	return o.BlogAuthor
}

// GetBlogAuthorOk returns a tuple with the BlogAuthor field value
// and a boolean to check if the value has been set.
func (o *BlogAuthorCloneRequestVNext) GetBlogAuthorOk() (*BlogAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlogAuthor, true
}

// SetBlogAuthor sets field value
func (o *BlogAuthorCloneRequestVNext) SetBlogAuthor(v BlogAuthor) {
	o.BlogAuthor = v
}

func (o BlogAuthorCloneRequestVNext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.PrimaryLanguage != nil {
		toSerialize["primaryLanguage"] = o.PrimaryLanguage
	}
	if true {
		toSerialize["blogAuthor"] = o.BlogAuthor
	}
	return json.Marshal(toSerialize)
}

type NullableBlogAuthorCloneRequestVNext struct {
	value *BlogAuthorCloneRequestVNext
	isSet bool
}

func (v NullableBlogAuthorCloneRequestVNext) Get() *BlogAuthorCloneRequestVNext {
	return v.value
}

func (v *NullableBlogAuthorCloneRequestVNext) Set(val *BlogAuthorCloneRequestVNext) {
	v.value = val
	v.isSet = true
}

func (v NullableBlogAuthorCloneRequestVNext) IsSet() bool {
	return v.isSet
}

func (v *NullableBlogAuthorCloneRequestVNext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlogAuthorCloneRequestVNext(val *BlogAuthorCloneRequestVNext) *NullableBlogAuthorCloneRequestVNext {
	return &NullableBlogAuthorCloneRequestVNext{value: val, isSet: true}
}

func (v NullableBlogAuthorCloneRequestVNext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlogAuthorCloneRequestVNext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
