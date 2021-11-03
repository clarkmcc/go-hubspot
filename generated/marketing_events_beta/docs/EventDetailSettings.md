# EventDetailSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **int32** | The id of the application the settings are for | 
**EventDetailsUrl** | **string** | The url that will be used to fetch marketing event details by id | 

## Methods

### NewEventDetailSettings

`func NewEventDetailSettings(appId int32, eventDetailsUrl string, ) *EventDetailSettings`

NewEventDetailSettings instantiates a new EventDetailSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDetailSettingsWithDefaults

`func NewEventDetailSettingsWithDefaults() *EventDetailSettings`

NewEventDetailSettingsWithDefaults instantiates a new EventDetailSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *EventDetailSettings) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *EventDetailSettings) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *EventDetailSettings) SetAppId(v int32)`

SetAppId sets AppId field to given value.


### GetEventDetailsUrl

`func (o *EventDetailSettings) GetEventDetailsUrl() string`

GetEventDetailsUrl returns the EventDetailsUrl field if non-nil, zero value otherwise.

### GetEventDetailsUrlOk

`func (o *EventDetailSettings) GetEventDetailsUrlOk() (*string, bool)`

GetEventDetailsUrlOk returns a tuple with the EventDetailsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDetailsUrl

`func (o *EventDetailSettings) SetEventDetailsUrl(v string)`

SetEventDetailsUrl sets EventDetailsUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


