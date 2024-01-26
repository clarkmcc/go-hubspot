/*
Marketing Events

These APIs allow you to interact with HubSpot's Marketing Events Extension. It allows you to: * Create, Read or update Marketing Event information in HubSpot * Specify whether a HubSpot contact has registered, attended or cancelled a registration to a Marketing Event. * Specify a URL that can be called to get the details of a Marketing Event.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_events_beta

import (
	"encoding/json"
)

// BatchInputMarketingEventSubscriber List of HubSpot contacts to subscribe to the marketing event
type BatchInputMarketingEventSubscriber struct {
	// List of HubSpot contacts to subscribe to the marketing event
	Inputs []MarketingEventSubscriber `json:"inputs"`
}

// NewBatchInputMarketingEventSubscriber instantiates a new BatchInputMarketingEventSubscriber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputMarketingEventSubscriber(inputs []MarketingEventSubscriber) *BatchInputMarketingEventSubscriber {
	this := BatchInputMarketingEventSubscriber{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputMarketingEventSubscriberWithDefaults instantiates a new BatchInputMarketingEventSubscriber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputMarketingEventSubscriberWithDefaults() *BatchInputMarketingEventSubscriber {
	this := BatchInputMarketingEventSubscriber{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputMarketingEventSubscriber) GetInputs() []MarketingEventSubscriber {
	if o == nil {
		var ret []MarketingEventSubscriber
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputMarketingEventSubscriber) GetInputsOk() ([]MarketingEventSubscriber, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputMarketingEventSubscriber) SetInputs(v []MarketingEventSubscriber) {
	o.Inputs = v
}

func (o BatchInputMarketingEventSubscriber) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["inputs"] = o.Inputs
	}
	return json.Marshal(toSerialize)
}

type NullableBatchInputMarketingEventSubscriber struct {
	value *BatchInputMarketingEventSubscriber
	isSet bool
}

func (v NullableBatchInputMarketingEventSubscriber) Get() *BatchInputMarketingEventSubscriber {
	return v.value
}

func (v *NullableBatchInputMarketingEventSubscriber) Set(val *BatchInputMarketingEventSubscriber) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputMarketingEventSubscriber) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputMarketingEventSubscriber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputMarketingEventSubscriber(val *BatchInputMarketingEventSubscriber) *NullableBatchInputMarketingEventSubscriber {
	return &NullableBatchInputMarketingEventSubscriber{value: val, isSet: true}
}

func (v NullableBatchInputMarketingEventSubscriber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputMarketingEventSubscriber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
