/*
Marketing Events Extension

These APIs allow you to interact with HubSpot's Marketing Events Extension. It allows you to: * Create, Read or update Marketing Event information in HubSpot * Specify whether a HubSpot contact has registered, attended or cancelled a registration to a Marketing Event. * Specify a URL that can be called to get the details of a Marketing Event.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_events_beta

import (
	"encoding/json"
	"time"
)

// MarketingEventUpdateRequestParams struct for MarketingEventUpdateRequestParams
type MarketingEventUpdateRequestParams struct {
	// The start date and time of the marketing event.
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// A list of PropertyValues. These can be whatever kind of property names and values you want. However, they must already exist on the HubSpot account's definition of the MarketingEvent Object. If they don't they will be filtered out and not set. In order to do this you'll need to create a new PropertyGroup on the HubSpot account's MarketingEvent object for your specific app and create the Custom Property you want to track on that HubSpot account. Do not create any new default properties on the MarketingEvent object as that will apply to all HubSpot accounts.
	CustomProperties []PropertyValue `json:"customProperties,omitempty"`
	// Indicates if the marketing event has been cancelled. Defaults to `false`
	EventCancelled *bool `json:"eventCancelled,omitempty"`
	// The name of the organizer of the marketing event.
	EventOrganizer *string `json:"eventOrganizer,omitempty"`
	// A URL in the external event application where the marketing event can be managed.
	EventUrl *string `json:"eventUrl,omitempty"`
	// The description of the marketing event.
	EventDescription *string `json:"eventDescription,omitempty"`
	// The name of the marketing event.
	EventName *string `json:"eventName,omitempty"`
	// Describes what type of event this is.  For example: `WEBINAR`, `CONFERENCE`, `WORKSHOP`
	EventType *string `json:"eventType,omitempty"`
	// The end date and time of the marketing event.
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
}

// NewMarketingEventUpdateRequestParams instantiates a new MarketingEventUpdateRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketingEventUpdateRequestParams() *MarketingEventUpdateRequestParams {
	this := MarketingEventUpdateRequestParams{}
	return &this
}

// NewMarketingEventUpdateRequestParamsWithDefaults instantiates a new MarketingEventUpdateRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketingEventUpdateRequestParamsWithDefaults() *MarketingEventUpdateRequestParams {
	this := MarketingEventUpdateRequestParams{}
	return &this
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise.
func (o *MarketingEventUpdateRequestParams) GetStartDateTime() time.Time {
	if o == nil || o.StartDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventUpdateRequestParams) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil || o.StartDateTime == nil {
		return nil, false
	}
	return o.StartDateTime, true
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *MarketingEventUpdateRequestParams) HasStartDateTime() bool {
	if o != nil && o.StartDateTime != nil {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given time.Time and assigns it to the StartDateTime field.
func (o *MarketingEventUpdateRequestParams) SetStartDateTime(v time.Time) {
	o.StartDateTime = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *MarketingEventUpdateRequestParams) GetCustomProperties() []PropertyValue {
	if o == nil || o.CustomProperties == nil {
		var ret []PropertyValue
		return ret
	}
	return o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventUpdateRequestParams) GetCustomPropertiesOk() ([]PropertyValue, bool) {
	if o == nil || o.CustomProperties == nil {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *MarketingEventUpdateRequestParams) HasCustomProperties() bool {
	if o != nil && o.CustomProperties != nil {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given []PropertyValue and assigns it to the CustomProperties field.
func (o *MarketingEventUpdateRequestParams) SetCustomProperties(v []PropertyValue) {
	o.CustomProperties = v
}

// GetEventCancelled returns the EventCancelled field value if set, zero value otherwise.
func (o *MarketingEventUpdateRequestParams) GetEventCancelled() bool {
	if o == nil || o.EventCancelled == nil {
		var ret bool
		return ret
	}
	return *o.EventCancelled
}

// GetEventCancelledOk returns a tuple with the EventCancelled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventUpdateRequestParams) GetEventCancelledOk() (*bool, bool) {
	if o == nil || o.EventCancelled == nil {
		return nil, false
	}
	return o.EventCancelled, true
}

// HasEventCancelled returns a boolean if a field has been set.
func (o *MarketingEventUpdateRequestParams) HasEventCancelled() bool {
	if o != nil && o.EventCancelled != nil {
		return true
	}

	return false
}

// SetEventCancelled gets a reference to the given bool and assigns it to the EventCancelled field.
func (o *MarketingEventUpdateRequestParams) SetEventCancelled(v bool) {
	o.EventCancelled = &v
}

// GetEventOrganizer returns the EventOrganizer field value if set, zero value otherwise.
func (o *MarketingEventUpdateRequestParams) GetEventOrganizer() string {
	if o == nil || o.EventOrganizer == nil {
		var ret string
		return ret
	}
	return *o.EventOrganizer
}

// GetEventOrganizerOk returns a tuple with the EventOrganizer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventUpdateRequestParams) GetEventOrganizerOk() (*string, bool) {
	if o == nil || o.EventOrganizer == nil {
		return nil, false
	}
	return o.EventOrganizer, true
}

// HasEventOrganizer returns a boolean if a field has been set.
func (o *MarketingEventUpdateRequestParams) HasEventOrganizer() bool {
	if o != nil && o.EventOrganizer != nil {
		return true
	}

	return false
}

// SetEventOrganizer gets a reference to the given string and assigns it to the EventOrganizer field.
func (o *MarketingEventUpdateRequestParams) SetEventOrganizer(v string) {
	o.EventOrganizer = &v
}

// GetEventUrl returns the EventUrl field value if set, zero value otherwise.
func (o *MarketingEventUpdateRequestParams) GetEventUrl() string {
	if o == nil || o.EventUrl == nil {
		var ret string
		return ret
	}
	return *o.EventUrl
}

// GetEventUrlOk returns a tuple with the EventUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventUpdateRequestParams) GetEventUrlOk() (*string, bool) {
	if o == nil || o.EventUrl == nil {
		return nil, false
	}
	return o.EventUrl, true
}

// HasEventUrl returns a boolean if a field has been set.
func (o *MarketingEventUpdateRequestParams) HasEventUrl() bool {
	if o != nil && o.EventUrl != nil {
		return true
	}

	return false
}

// SetEventUrl gets a reference to the given string and assigns it to the EventUrl field.
func (o *MarketingEventUpdateRequestParams) SetEventUrl(v string) {
	o.EventUrl = &v
}

// GetEventDescription returns the EventDescription field value if set, zero value otherwise.
func (o *MarketingEventUpdateRequestParams) GetEventDescription() string {
	if o == nil || o.EventDescription == nil {
		var ret string
		return ret
	}
	return *o.EventDescription
}

// GetEventDescriptionOk returns a tuple with the EventDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventUpdateRequestParams) GetEventDescriptionOk() (*string, bool) {
	if o == nil || o.EventDescription == nil {
		return nil, false
	}
	return o.EventDescription, true
}

// HasEventDescription returns a boolean if a field has been set.
func (o *MarketingEventUpdateRequestParams) HasEventDescription() bool {
	if o != nil && o.EventDescription != nil {
		return true
	}

	return false
}

// SetEventDescription gets a reference to the given string and assigns it to the EventDescription field.
func (o *MarketingEventUpdateRequestParams) SetEventDescription(v string) {
	o.EventDescription = &v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *MarketingEventUpdateRequestParams) GetEventName() string {
	if o == nil || o.EventName == nil {
		var ret string
		return ret
	}
	return *o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventUpdateRequestParams) GetEventNameOk() (*string, bool) {
	if o == nil || o.EventName == nil {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *MarketingEventUpdateRequestParams) HasEventName() bool {
	if o != nil && o.EventName != nil {
		return true
	}

	return false
}

// SetEventName gets a reference to the given string and assigns it to the EventName field.
func (o *MarketingEventUpdateRequestParams) SetEventName(v string) {
	o.EventName = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *MarketingEventUpdateRequestParams) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventUpdateRequestParams) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *MarketingEventUpdateRequestParams) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *MarketingEventUpdateRequestParams) SetEventType(v string) {
	o.EventType = &v
}

// GetEndDateTime returns the EndDateTime field value if set, zero value otherwise.
func (o *MarketingEventUpdateRequestParams) GetEndDateTime() time.Time {
	if o == nil || o.EndDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDateTime
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventUpdateRequestParams) GetEndDateTimeOk() (*time.Time, bool) {
	if o == nil || o.EndDateTime == nil {
		return nil, false
	}
	return o.EndDateTime, true
}

// HasEndDateTime returns a boolean if a field has been set.
func (o *MarketingEventUpdateRequestParams) HasEndDateTime() bool {
	if o != nil && o.EndDateTime != nil {
		return true
	}

	return false
}

// SetEndDateTime gets a reference to the given time.Time and assigns it to the EndDateTime field.
func (o *MarketingEventUpdateRequestParams) SetEndDateTime(v time.Time) {
	o.EndDateTime = &v
}

func (o MarketingEventUpdateRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StartDateTime != nil {
		toSerialize["startDateTime"] = o.StartDateTime
	}
	if o.CustomProperties != nil {
		toSerialize["customProperties"] = o.CustomProperties
	}
	if o.EventCancelled != nil {
		toSerialize["eventCancelled"] = o.EventCancelled
	}
	if o.EventOrganizer != nil {
		toSerialize["eventOrganizer"] = o.EventOrganizer
	}
	if o.EventUrl != nil {
		toSerialize["eventUrl"] = o.EventUrl
	}
	if o.EventDescription != nil {
		toSerialize["eventDescription"] = o.EventDescription
	}
	if o.EventName != nil {
		toSerialize["eventName"] = o.EventName
	}
	if o.EventType != nil {
		toSerialize["eventType"] = o.EventType
	}
	if o.EndDateTime != nil {
		toSerialize["endDateTime"] = o.EndDateTime
	}
	return json.Marshal(toSerialize)
}

type NullableMarketingEventUpdateRequestParams struct {
	value *MarketingEventUpdateRequestParams
	isSet bool
}

func (v NullableMarketingEventUpdateRequestParams) Get() *MarketingEventUpdateRequestParams {
	return v.value
}

func (v *NullableMarketingEventUpdateRequestParams) Set(val *MarketingEventUpdateRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketingEventUpdateRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketingEventUpdateRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketingEventUpdateRequestParams(val *MarketingEventUpdateRequestParams) *NullableMarketingEventUpdateRequestParams {
	return &NullableMarketingEventUpdateRequestParams{value: val, isSet: true}
}

func (v NullableMarketingEventUpdateRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketingEventUpdateRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
