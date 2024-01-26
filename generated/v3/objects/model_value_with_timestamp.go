/*
CRM Objects

CRM objects such as companies, contacts, deals, line items, products, tickets, and quotes are standard objects in HubSpot’s CRM. These core building blocks support custom properties, store critical information, and play a central role in the HubSpot application.  ## Supported Object Types  This API provides access to collections of CRM objects, which return a map of property names to values. Each object type has its own set of default properties, which can be found by exploring the [CRM Object Properties API](https://developers.hubspot.com/docs/methods/crm-properties/crm-properties-overview).  |Object Type |Properties returned by default | |--|--| | `companies` | `name`, `domain` | | `contacts` | `firstname`, `lastname`, `email` | | `deals` | `dealname`, `amount`, `closedate`, `pipeline`, `dealstage` | | `products` | `name`, `description`, `price` | | `tickets` | `content`, `hs_pipeline`, `hs_pipeline_stage`, `hs_ticket_category`, `hs_ticket_priority`, `subject` |  Find a list of all properties for an object type using the [CRM Object Properties](https://developers.hubspot.com/docs/methods/crm-properties/get-properties) API. e.g. `GET https://api.hubapi.com/properties/v2/companies/properties`. Change the properties returned in the response using the `properties` array in the request body.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objects

import (
	"encoding/json"
	"time"
)

// ValueWithTimestamp struct for ValueWithTimestamp
type ValueWithTimestamp struct {
	SourceId        *string   `json:"sourceId,omitempty"`
	SourceType      string    `json:"sourceType"`
	SourceLabel     *string   `json:"sourceLabel,omitempty"`
	UpdatedByUserId *int32    `json:"updatedByUserId,omitempty"`
	Value           string    `json:"value"`
	Timestamp       time.Time `json:"timestamp"`
}

// NewValueWithTimestamp instantiates a new ValueWithTimestamp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValueWithTimestamp(sourceType string, value string, timestamp time.Time) *ValueWithTimestamp {
	this := ValueWithTimestamp{}
	this.SourceType = sourceType
	this.Value = value
	this.Timestamp = timestamp
	return &this
}

// NewValueWithTimestampWithDefaults instantiates a new ValueWithTimestamp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueWithTimestampWithDefaults() *ValueWithTimestamp {
	this := ValueWithTimestamp{}
	return &this
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *ValueWithTimestamp) GetSourceId() string {
	if o == nil || o.SourceId == nil {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueWithTimestamp) GetSourceIdOk() (*string, bool) {
	if o == nil || o.SourceId == nil {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *ValueWithTimestamp) HasSourceId() bool {
	if o != nil && o.SourceId != nil {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *ValueWithTimestamp) SetSourceId(v string) {
	o.SourceId = &v
}

// GetSourceType returns the SourceType field value
func (o *ValueWithTimestamp) GetSourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *ValueWithTimestamp) GetSourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *ValueWithTimestamp) SetSourceType(v string) {
	o.SourceType = v
}

// GetSourceLabel returns the SourceLabel field value if set, zero value otherwise.
func (o *ValueWithTimestamp) GetSourceLabel() string {
	if o == nil || o.SourceLabel == nil {
		var ret string
		return ret
	}
	return *o.SourceLabel
}

// GetSourceLabelOk returns a tuple with the SourceLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueWithTimestamp) GetSourceLabelOk() (*string, bool) {
	if o == nil || o.SourceLabel == nil {
		return nil, false
	}
	return o.SourceLabel, true
}

// HasSourceLabel returns a boolean if a field has been set.
func (o *ValueWithTimestamp) HasSourceLabel() bool {
	if o != nil && o.SourceLabel != nil {
		return true
	}

	return false
}

// SetSourceLabel gets a reference to the given string and assigns it to the SourceLabel field.
func (o *ValueWithTimestamp) SetSourceLabel(v string) {
	o.SourceLabel = &v
}

// GetUpdatedByUserId returns the UpdatedByUserId field value if set, zero value otherwise.
func (o *ValueWithTimestamp) GetUpdatedByUserId() int32 {
	if o == nil || o.UpdatedByUserId == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedByUserId
}

// GetUpdatedByUserIdOk returns a tuple with the UpdatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueWithTimestamp) GetUpdatedByUserIdOk() (*int32, bool) {
	if o == nil || o.UpdatedByUserId == nil {
		return nil, false
	}
	return o.UpdatedByUserId, true
}

// HasUpdatedByUserId returns a boolean if a field has been set.
func (o *ValueWithTimestamp) HasUpdatedByUserId() bool {
	if o != nil && o.UpdatedByUserId != nil {
		return true
	}

	return false
}

// SetUpdatedByUserId gets a reference to the given int32 and assigns it to the UpdatedByUserId field.
func (o *ValueWithTimestamp) SetUpdatedByUserId(v int32) {
	o.UpdatedByUserId = &v
}

// GetValue returns the Value field value
func (o *ValueWithTimestamp) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ValueWithTimestamp) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ValueWithTimestamp) SetValue(v string) {
	o.Value = v
}

// GetTimestamp returns the Timestamp field value
func (o *ValueWithTimestamp) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ValueWithTimestamp) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ValueWithTimestamp) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o ValueWithTimestamp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SourceId != nil {
		toSerialize["sourceId"] = o.SourceId
	}
	if true {
		toSerialize["sourceType"] = o.SourceType
	}
	if o.SourceLabel != nil {
		toSerialize["sourceLabel"] = o.SourceLabel
	}
	if o.UpdatedByUserId != nil {
		toSerialize["updatedByUserId"] = o.UpdatedByUserId
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableValueWithTimestamp struct {
	value *ValueWithTimestamp
	isSet bool
}

func (v NullableValueWithTimestamp) Get() *ValueWithTimestamp {
	return v.value
}

func (v *NullableValueWithTimestamp) Set(val *ValueWithTimestamp) {
	v.value = val
	v.isSet = true
}

func (v NullableValueWithTimestamp) IsSet() bool {
	return v.isSet
}

func (v *NullableValueWithTimestamp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueWithTimestamp(val *ValueWithTimestamp) *NullableValueWithTimestamp {
	return &NullableValueWithTimestamp{value: val, isSet: true}
}

func (v NullableValueWithTimestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValueWithTimestamp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
