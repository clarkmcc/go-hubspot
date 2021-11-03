# MarketingEventPublicReadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventName** | **string** | The name of the marketing event. | 
**EventType** | Pointer to **string** | The type of the marketing event. | [optional] 
**StartDateTime** | Pointer to **time.Time** | The start date and time of the marketing event. | [optional] 
**EndDateTime** | Pointer to **time.Time** | The end date and time of the marketing event. | [optional] 
**EventOrganizer** | **string** | The name of the organizer of the marketing event. | 
**EventDescription** | Pointer to **string** | The description of the marketing event. | [optional] 
**EventUrl** | Pointer to **string** | A URL in the external event application where the marketing event can be managed. | [optional] 
**EventCancelled** | Pointer to **bool** | Indicates if the marketing event has been cancelled. | [optional] 
**CustomProperties** | Pointer to [**[]PropertyValue**](PropertyValue.md) | A list of PropertyValues. These can be whatever kind of property names and values you want. However, they must already exist on the HubSpot account&#39;s definition of the MarketingEvent Object. If they don&#39;t they will be filtered out and not set. In order to do this you&#39;ll need to create a new PropertyGroup on the HubSpot account&#39;s MarketingEvent object for your specific app and create the Custom Property you want to track on that HubSpot account. Do not create any new default properties on the MarketingEvent object as that will apply to all HubSpot accounts.  | [optional] 
**ExternalEventId** | **string** | The id of the marketing event in the external event application. | 
**Registrants** | **int32** | The number of HubSpot contacts that registered for this marketing event. | 
**Attendees** | **int32** | The number of HubSpot contacts that attended this marketing event. | 
**Cancellations** | **int32** | The number of HubSpot contacts that registered for this marketing event, but later cancelled their registration. | 
**NoShows** | **int32** | The number of HubSpot contacts that registered for this marketing event, but did not attend. This field only had a value when the event is over. | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Id** | **string** |  | 

## Methods

### NewMarketingEventPublicReadResponse

`func NewMarketingEventPublicReadResponse(eventName string, eventOrganizer string, externalEventId string, registrants int32, attendees int32, cancellations int32, noShows int32, createdAt time.Time, updatedAt time.Time, id string, ) *MarketingEventPublicReadResponse`

NewMarketingEventPublicReadResponse instantiates a new MarketingEventPublicReadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingEventPublicReadResponseWithDefaults

`func NewMarketingEventPublicReadResponseWithDefaults() *MarketingEventPublicReadResponse`

NewMarketingEventPublicReadResponseWithDefaults instantiates a new MarketingEventPublicReadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventName

`func (o *MarketingEventPublicReadResponse) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *MarketingEventPublicReadResponse) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *MarketingEventPublicReadResponse) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetEventType

`func (o *MarketingEventPublicReadResponse) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MarketingEventPublicReadResponse) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MarketingEventPublicReadResponse) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *MarketingEventPublicReadResponse) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetStartDateTime

`func (o *MarketingEventPublicReadResponse) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MarketingEventPublicReadResponse) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MarketingEventPublicReadResponse) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MarketingEventPublicReadResponse) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetEndDateTime

`func (o *MarketingEventPublicReadResponse) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MarketingEventPublicReadResponse) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MarketingEventPublicReadResponse) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MarketingEventPublicReadResponse) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetEventOrganizer

`func (o *MarketingEventPublicReadResponse) GetEventOrganizer() string`

GetEventOrganizer returns the EventOrganizer field if non-nil, zero value otherwise.

### GetEventOrganizerOk

`func (o *MarketingEventPublicReadResponse) GetEventOrganizerOk() (*string, bool)`

GetEventOrganizerOk returns a tuple with the EventOrganizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventOrganizer

`func (o *MarketingEventPublicReadResponse) SetEventOrganizer(v string)`

SetEventOrganizer sets EventOrganizer field to given value.


### GetEventDescription

`func (o *MarketingEventPublicReadResponse) GetEventDescription() string`

GetEventDescription returns the EventDescription field if non-nil, zero value otherwise.

### GetEventDescriptionOk

`func (o *MarketingEventPublicReadResponse) GetEventDescriptionOk() (*string, bool)`

GetEventDescriptionOk returns a tuple with the EventDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDescription

`func (o *MarketingEventPublicReadResponse) SetEventDescription(v string)`

SetEventDescription sets EventDescription field to given value.

### HasEventDescription

`func (o *MarketingEventPublicReadResponse) HasEventDescription() bool`

HasEventDescription returns a boolean if a field has been set.

### GetEventUrl

`func (o *MarketingEventPublicReadResponse) GetEventUrl() string`

GetEventUrl returns the EventUrl field if non-nil, zero value otherwise.

### GetEventUrlOk

`func (o *MarketingEventPublicReadResponse) GetEventUrlOk() (*string, bool)`

GetEventUrlOk returns a tuple with the EventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventUrl

`func (o *MarketingEventPublicReadResponse) SetEventUrl(v string)`

SetEventUrl sets EventUrl field to given value.

### HasEventUrl

`func (o *MarketingEventPublicReadResponse) HasEventUrl() bool`

HasEventUrl returns a boolean if a field has been set.

### GetEventCancelled

`func (o *MarketingEventPublicReadResponse) GetEventCancelled() bool`

GetEventCancelled returns the EventCancelled field if non-nil, zero value otherwise.

### GetEventCancelledOk

`func (o *MarketingEventPublicReadResponse) GetEventCancelledOk() (*bool, bool)`

GetEventCancelledOk returns a tuple with the EventCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCancelled

`func (o *MarketingEventPublicReadResponse) SetEventCancelled(v bool)`

SetEventCancelled sets EventCancelled field to given value.

### HasEventCancelled

`func (o *MarketingEventPublicReadResponse) HasEventCancelled() bool`

HasEventCancelled returns a boolean if a field has been set.

### GetCustomProperties

`func (o *MarketingEventPublicReadResponse) GetCustomProperties() []PropertyValue`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *MarketingEventPublicReadResponse) GetCustomPropertiesOk() (*[]PropertyValue, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *MarketingEventPublicReadResponse) SetCustomProperties(v []PropertyValue)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *MarketingEventPublicReadResponse) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetExternalEventId

`func (o *MarketingEventPublicReadResponse) GetExternalEventId() string`

GetExternalEventId returns the ExternalEventId field if non-nil, zero value otherwise.

### GetExternalEventIdOk

`func (o *MarketingEventPublicReadResponse) GetExternalEventIdOk() (*string, bool)`

GetExternalEventIdOk returns a tuple with the ExternalEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEventId

`func (o *MarketingEventPublicReadResponse) SetExternalEventId(v string)`

SetExternalEventId sets ExternalEventId field to given value.


### GetRegistrants

`func (o *MarketingEventPublicReadResponse) GetRegistrants() int32`

GetRegistrants returns the Registrants field if non-nil, zero value otherwise.

### GetRegistrantsOk

`func (o *MarketingEventPublicReadResponse) GetRegistrantsOk() (*int32, bool)`

GetRegistrantsOk returns a tuple with the Registrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrants

`func (o *MarketingEventPublicReadResponse) SetRegistrants(v int32)`

SetRegistrants sets Registrants field to given value.


### GetAttendees

`func (o *MarketingEventPublicReadResponse) GetAttendees() int32`

GetAttendees returns the Attendees field if non-nil, zero value otherwise.

### GetAttendeesOk

`func (o *MarketingEventPublicReadResponse) GetAttendeesOk() (*int32, bool)`

GetAttendeesOk returns a tuple with the Attendees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendees

`func (o *MarketingEventPublicReadResponse) SetAttendees(v int32)`

SetAttendees sets Attendees field to given value.


### GetCancellations

`func (o *MarketingEventPublicReadResponse) GetCancellations() int32`

GetCancellations returns the Cancellations field if non-nil, zero value otherwise.

### GetCancellationsOk

`func (o *MarketingEventPublicReadResponse) GetCancellationsOk() (*int32, bool)`

GetCancellationsOk returns a tuple with the Cancellations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellations

`func (o *MarketingEventPublicReadResponse) SetCancellations(v int32)`

SetCancellations sets Cancellations field to given value.


### GetNoShows

`func (o *MarketingEventPublicReadResponse) GetNoShows() int32`

GetNoShows returns the NoShows field if non-nil, zero value otherwise.

### GetNoShowsOk

`func (o *MarketingEventPublicReadResponse) GetNoShowsOk() (*int32, bool)`

GetNoShowsOk returns a tuple with the NoShows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoShows

`func (o *MarketingEventPublicReadResponse) SetNoShows(v int32)`

SetNoShows sets NoShows field to given value.


### GetCreatedAt

`func (o *MarketingEventPublicReadResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MarketingEventPublicReadResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MarketingEventPublicReadResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *MarketingEventPublicReadResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MarketingEventPublicReadResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MarketingEventPublicReadResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetId

`func (o *MarketingEventPublicReadResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketingEventPublicReadResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketingEventPublicReadResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


