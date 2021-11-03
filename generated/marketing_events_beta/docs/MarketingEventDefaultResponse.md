# MarketingEventDefaultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventName** | **string** | The name of the marketing event. | 
**EventType** | Pointer to **string** | The type of the marketing event. | [optional] 
**StartDateTime** | Pointer to **time.Time** | The start date and time of the marketing event. | [optional] 
**EndDateTime** | Pointer to **time.Time** | The end date and time of the marketing event. | [optional] 
**EventOrganizer** | **string** | The name of the organizer of the marketing event. | 
**EventDescription** | Pointer to **string** | The description of the marketing event. | [optional] 
**EventUrl** | Pointer to **string** | The URL in the external event application where the marketing event can be managed. | [optional] 
**EventCancelled** | Pointer to **bool** | Indicates if the marketing event has been cancelled. | [optional] 
**CustomProperties** | Pointer to [**[]PropertyValue**](PropertyValue.md) | A list of PropertyValues. These can be whatever kind of property names and values you want. However, they must already exist on the HubSpot account&#39;s definition of the MarketingEvent Object. If they don&#39;t they will be filtered out and not set. In order to do this you&#39;ll need to create a new PropertyGroup on the HubSpot account&#39;s MarketingEvent object for your specific app and create the Custom Property you want to track on that HubSpot account. Do not create any new default properties on the MarketingEvent object as that will apply to all HubSpot accounts.  | [optional] 

## Methods

### NewMarketingEventDefaultResponse

`func NewMarketingEventDefaultResponse(eventName string, eventOrganizer string, ) *MarketingEventDefaultResponse`

NewMarketingEventDefaultResponse instantiates a new MarketingEventDefaultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingEventDefaultResponseWithDefaults

`func NewMarketingEventDefaultResponseWithDefaults() *MarketingEventDefaultResponse`

NewMarketingEventDefaultResponseWithDefaults instantiates a new MarketingEventDefaultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventName

`func (o *MarketingEventDefaultResponse) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *MarketingEventDefaultResponse) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *MarketingEventDefaultResponse) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetEventType

`func (o *MarketingEventDefaultResponse) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MarketingEventDefaultResponse) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MarketingEventDefaultResponse) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *MarketingEventDefaultResponse) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetStartDateTime

`func (o *MarketingEventDefaultResponse) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MarketingEventDefaultResponse) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MarketingEventDefaultResponse) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MarketingEventDefaultResponse) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetEndDateTime

`func (o *MarketingEventDefaultResponse) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MarketingEventDefaultResponse) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MarketingEventDefaultResponse) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MarketingEventDefaultResponse) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetEventOrganizer

`func (o *MarketingEventDefaultResponse) GetEventOrganizer() string`

GetEventOrganizer returns the EventOrganizer field if non-nil, zero value otherwise.

### GetEventOrganizerOk

`func (o *MarketingEventDefaultResponse) GetEventOrganizerOk() (*string, bool)`

GetEventOrganizerOk returns a tuple with the EventOrganizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventOrganizer

`func (o *MarketingEventDefaultResponse) SetEventOrganizer(v string)`

SetEventOrganizer sets EventOrganizer field to given value.


### GetEventDescription

`func (o *MarketingEventDefaultResponse) GetEventDescription() string`

GetEventDescription returns the EventDescription field if non-nil, zero value otherwise.

### GetEventDescriptionOk

`func (o *MarketingEventDefaultResponse) GetEventDescriptionOk() (*string, bool)`

GetEventDescriptionOk returns a tuple with the EventDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDescription

`func (o *MarketingEventDefaultResponse) SetEventDescription(v string)`

SetEventDescription sets EventDescription field to given value.

### HasEventDescription

`func (o *MarketingEventDefaultResponse) HasEventDescription() bool`

HasEventDescription returns a boolean if a field has been set.

### GetEventUrl

`func (o *MarketingEventDefaultResponse) GetEventUrl() string`

GetEventUrl returns the EventUrl field if non-nil, zero value otherwise.

### GetEventUrlOk

`func (o *MarketingEventDefaultResponse) GetEventUrlOk() (*string, bool)`

GetEventUrlOk returns a tuple with the EventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventUrl

`func (o *MarketingEventDefaultResponse) SetEventUrl(v string)`

SetEventUrl sets EventUrl field to given value.

### HasEventUrl

`func (o *MarketingEventDefaultResponse) HasEventUrl() bool`

HasEventUrl returns a boolean if a field has been set.

### GetEventCancelled

`func (o *MarketingEventDefaultResponse) GetEventCancelled() bool`

GetEventCancelled returns the EventCancelled field if non-nil, zero value otherwise.

### GetEventCancelledOk

`func (o *MarketingEventDefaultResponse) GetEventCancelledOk() (*bool, bool)`

GetEventCancelledOk returns a tuple with the EventCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCancelled

`func (o *MarketingEventDefaultResponse) SetEventCancelled(v bool)`

SetEventCancelled sets EventCancelled field to given value.

### HasEventCancelled

`func (o *MarketingEventDefaultResponse) HasEventCancelled() bool`

HasEventCancelled returns a boolean if a field has been set.

### GetCustomProperties

`func (o *MarketingEventDefaultResponse) GetCustomProperties() []PropertyValue`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *MarketingEventDefaultResponse) GetCustomPropertiesOk() (*[]PropertyValue, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *MarketingEventDefaultResponse) SetCustomProperties(v []PropertyValue)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *MarketingEventDefaultResponse) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


