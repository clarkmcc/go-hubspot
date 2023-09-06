/*
Blog Post endpoints

Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authors

import (
	"encoding/json"
	"time"
)

// BatchResponseBlogPostWithErrors Response object for batch operations on blog posts with errors.
type BatchResponseBlogPostWithErrors struct {
	// Status of batch operation.
	Status string `json:"status"`
	// Results of batch operation.
	Results []BlogPost `json:"results"`
	// Number of errors.
	NumErrors *int32 `json:"numErrors,omitempty"`
	// Errors in batch operation.
	Errors []StandardError `json:"errors,omitempty"`
	// Time of batch operation request.
	RequestedAt *time.Time `json:"requestedAt,omitempty"`
	// Time of batch operation start.
	StartedAt time.Time `json:"startedAt"`
	// Time of batch operation completion.
	CompletedAt time.Time `json:"completedAt"`
	// Links associated with batch operation.
	Links *map[string]string `json:"links,omitempty"`
}

// NewBatchResponseBlogPostWithErrors instantiates a new BatchResponseBlogPostWithErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchResponseBlogPostWithErrors(status string, results []BlogPost, startedAt time.Time, completedAt time.Time) *BatchResponseBlogPostWithErrors {
	this := BatchResponseBlogPostWithErrors{}
	this.Status = status
	this.Results = results
	this.StartedAt = startedAt
	this.CompletedAt = completedAt
	return &this
}

// NewBatchResponseBlogPostWithErrorsWithDefaults instantiates a new BatchResponseBlogPostWithErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchResponseBlogPostWithErrorsWithDefaults() *BatchResponseBlogPostWithErrors {
	this := BatchResponseBlogPostWithErrors{}
	return &this
}

// GetStatus returns the Status field value
func (o *BatchResponseBlogPostWithErrors) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BatchResponseBlogPostWithErrors) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BatchResponseBlogPostWithErrors) SetStatus(v string) {
	o.Status = v
}

// GetResults returns the Results field value
func (o *BatchResponseBlogPostWithErrors) GetResults() []BlogPost {
	if o == nil {
		var ret []BlogPost
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *BatchResponseBlogPostWithErrors) GetResultsOk() ([]BlogPost, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *BatchResponseBlogPostWithErrors) SetResults(v []BlogPost) {
	o.Results = v
}

// GetNumErrors returns the NumErrors field value if set, zero value otherwise.
func (o *BatchResponseBlogPostWithErrors) GetNumErrors() int32 {
	if o == nil || o.NumErrors == nil {
		var ret int32
		return ret
	}
	return *o.NumErrors
}

// GetNumErrorsOk returns a tuple with the NumErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseBlogPostWithErrors) GetNumErrorsOk() (*int32, bool) {
	if o == nil || o.NumErrors == nil {
		return nil, false
	}
	return o.NumErrors, true
}

// HasNumErrors returns a boolean if a field has been set.
func (o *BatchResponseBlogPostWithErrors) HasNumErrors() bool {
	if o != nil && o.NumErrors != nil {
		return true
	}

	return false
}

// SetNumErrors gets a reference to the given int32 and assigns it to the NumErrors field.
func (o *BatchResponseBlogPostWithErrors) SetNumErrors(v int32) {
	o.NumErrors = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *BatchResponseBlogPostWithErrors) GetErrors() []StandardError {
	if o == nil || o.Errors == nil {
		var ret []StandardError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseBlogPostWithErrors) GetErrorsOk() ([]StandardError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *BatchResponseBlogPostWithErrors) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []StandardError and assigns it to the Errors field.
func (o *BatchResponseBlogPostWithErrors) SetErrors(v []StandardError) {
	o.Errors = v
}

// GetRequestedAt returns the RequestedAt field value if set, zero value otherwise.
func (o *BatchResponseBlogPostWithErrors) GetRequestedAt() time.Time {
	if o == nil || o.RequestedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.RequestedAt
}

// GetRequestedAtOk returns a tuple with the RequestedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseBlogPostWithErrors) GetRequestedAtOk() (*time.Time, bool) {
	if o == nil || o.RequestedAt == nil {
		return nil, false
	}
	return o.RequestedAt, true
}

// HasRequestedAt returns a boolean if a field has been set.
func (o *BatchResponseBlogPostWithErrors) HasRequestedAt() bool {
	if o != nil && o.RequestedAt != nil {
		return true
	}

	return false
}

// SetRequestedAt gets a reference to the given time.Time and assigns it to the RequestedAt field.
func (o *BatchResponseBlogPostWithErrors) SetRequestedAt(v time.Time) {
	o.RequestedAt = &v
}

// GetStartedAt returns the StartedAt field value
func (o *BatchResponseBlogPostWithErrors) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *BatchResponseBlogPostWithErrors) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *BatchResponseBlogPostWithErrors) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetCompletedAt returns the CompletedAt field value
func (o *BatchResponseBlogPostWithErrors) GetCompletedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value
// and a boolean to check if the value has been set.
func (o *BatchResponseBlogPostWithErrors) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletedAt, true
}

// SetCompletedAt sets field value
func (o *BatchResponseBlogPostWithErrors) SetCompletedAt(v time.Time) {
	o.CompletedAt = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BatchResponseBlogPostWithErrors) GetLinks() map[string]string {
	if o == nil || o.Links == nil {
		var ret map[string]string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseBlogPostWithErrors) GetLinksOk() (*map[string]string, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BatchResponseBlogPostWithErrors) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]string and assigns it to the Links field.
func (o *BatchResponseBlogPostWithErrors) SetLinks(v map[string]string) {
	o.Links = &v
}

func (o BatchResponseBlogPostWithErrors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["results"] = o.Results
	}
	if o.NumErrors != nil {
		toSerialize["numErrors"] = o.NumErrors
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.RequestedAt != nil {
		toSerialize["requestedAt"] = o.RequestedAt
	}
	if true {
		toSerialize["startedAt"] = o.StartedAt
	}
	if true {
		toSerialize["completedAt"] = o.CompletedAt
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableBatchResponseBlogPostWithErrors struct {
	value *BatchResponseBlogPostWithErrors
	isSet bool
}

func (v NullableBatchResponseBlogPostWithErrors) Get() *BatchResponseBlogPostWithErrors {
	return v.value
}

func (v *NullableBatchResponseBlogPostWithErrors) Set(val *BatchResponseBlogPostWithErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchResponseBlogPostWithErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchResponseBlogPostWithErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchResponseBlogPostWithErrors(val *BatchResponseBlogPostWithErrors) *NullableBatchResponseBlogPostWithErrors {
	return &NullableBatchResponseBlogPostWithErrors{value: val, isSet: true}
}

func (v NullableBatchResponseBlogPostWithErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchResponseBlogPostWithErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
