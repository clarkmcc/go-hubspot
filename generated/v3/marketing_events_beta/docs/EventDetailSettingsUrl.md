# EventDetailSettingsUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventDetailsUrl** | **string** | The url that will be used to fetch marketing event details by id. Must contain a &#x60;%s&#x60; character sequence that will be substituted with the event id. For example: &#x60;https://my.event.app/events/%s&#x60; | 

## Methods

### NewEventDetailSettingsUrl

`func NewEventDetailSettingsUrl(eventDetailsUrl string, ) *EventDetailSettingsUrl`

NewEventDetailSettingsUrl instantiates a new EventDetailSettingsUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDetailSettingsUrlWithDefaults

`func NewEventDetailSettingsUrlWithDefaults() *EventDetailSettingsUrl`

NewEventDetailSettingsUrlWithDefaults instantiates a new EventDetailSettingsUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventDetailsUrl

`func (o *EventDetailSettingsUrl) GetEventDetailsUrl() string`

GetEventDetailsUrl returns the EventDetailsUrl field if non-nil, zero value otherwise.

### GetEventDetailsUrlOk

`func (o *EventDetailSettingsUrl) GetEventDetailsUrlOk() (*string, bool)`

GetEventDetailsUrlOk returns a tuple with the EventDetailsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDetailsUrl

`func (o *EventDetailSettingsUrl) SetEventDetailsUrl(v string)`

SetEventDetailsUrl sets EventDetailsUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


