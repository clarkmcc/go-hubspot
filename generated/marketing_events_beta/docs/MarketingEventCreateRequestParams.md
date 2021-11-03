# MarketingEventCreateRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventName** | **string** | The name of the marketing event. | 
**EventType** | Pointer to **string** | Describes what type of event this is.  For example: &#x60;WEBINAR&#x60;, &#x60;CONFERENCE&#x60;, &#x60;WORKSHOP&#x60; | [optional] 
**StartDateTime** | Pointer to **time.Time** | The start date and time of the marketing event. | [optional] 
**EndDateTime** | Pointer to **time.Time** | The end date and time of the marketing event. | [optional] 
**EventOrganizer** | **string** | The name of the organizer of the marketing event. | 
**EventDescription** | Pointer to **string** | The description of the marketing event. | [optional] 
**EventUrl** | Pointer to **string** | A URL in the external event application where the marketing event can be managed. | [optional] 
**EventCancelled** | Pointer to **bool** | Indicates if the marketing event has been cancelled.  Defaults to &#x60;false&#x60; | [optional] 
**CustomProperties** | Pointer to [**[]PropertyValue**](PropertyValue.md) | A list of PropertyValues. These can be whatever kind of property names and values you want. However, they must already exist on the HubSpot account&#39;s definition of the MarketingEvent Object. If they don&#39;t they will be filtered out and not set. In order to do this you&#39;ll need to create a new PropertyGroup on the HubSpot account&#39;s MarketingEvent object for your specific app and create the Custom Property you want to track on that HubSpot account. Do not create any new default properties on the MarketingEvent object as that will apply to all HubSpot accounts.  | [optional] 
**ExternalAccountId** | **string** | The accountId that is associated with this marketing event in the external event application. | 
**ExternalEventId** | **string** | The id of the marketing event in the external event application. | 

## Methods

### NewMarketingEventCreateRequestParams

`func NewMarketingEventCreateRequestParams(eventName string, eventOrganizer string, externalAccountId string, externalEventId string, ) *MarketingEventCreateRequestParams`

NewMarketingEventCreateRequestParams instantiates a new MarketingEventCreateRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingEventCreateRequestParamsWithDefaults

`func NewMarketingEventCreateRequestParamsWithDefaults() *MarketingEventCreateRequestParams`

NewMarketingEventCreateRequestParamsWithDefaults instantiates a new MarketingEventCreateRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventName

`func (o *MarketingEventCreateRequestParams) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *MarketingEventCreateRequestParams) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *MarketingEventCreateRequestParams) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetEventType

`func (o *MarketingEventCreateRequestParams) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MarketingEventCreateRequestParams) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MarketingEventCreateRequestParams) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *MarketingEventCreateRequestParams) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetStartDateTime

`func (o *MarketingEventCreateRequestParams) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MarketingEventCreateRequestParams) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MarketingEventCreateRequestParams) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MarketingEventCreateRequestParams) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetEndDateTime

`func (o *MarketingEventCreateRequestParams) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MarketingEventCreateRequestParams) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MarketingEventCreateRequestParams) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MarketingEventCreateRequestParams) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetEventOrganizer

`func (o *MarketingEventCreateRequestParams) GetEventOrganizer() string`

GetEventOrganizer returns the EventOrganizer field if non-nil, zero value otherwise.

### GetEventOrganizerOk

`func (o *MarketingEventCreateRequestParams) GetEventOrganizerOk() (*string, bool)`

GetEventOrganizerOk returns a tuple with the EventOrganizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventOrganizer

`func (o *MarketingEventCreateRequestParams) SetEventOrganizer(v string)`

SetEventOrganizer sets EventOrganizer field to given value.


### GetEventDescription

`func (o *MarketingEventCreateRequestParams) GetEventDescription() string`

GetEventDescription returns the EventDescription field if non-nil, zero value otherwise.

### GetEventDescriptionOk

`func (o *MarketingEventCreateRequestParams) GetEventDescriptionOk() (*string, bool)`

GetEventDescriptionOk returns a tuple with the EventDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDescription

`func (o *MarketingEventCreateRequestParams) SetEventDescription(v string)`

SetEventDescription sets EventDescription field to given value.

### HasEventDescription

`func (o *MarketingEventCreateRequestParams) HasEventDescription() bool`

HasEventDescription returns a boolean if a field has been set.

### GetEventUrl

`func (o *MarketingEventCreateRequestParams) GetEventUrl() string`

GetEventUrl returns the EventUrl field if non-nil, zero value otherwise.

### GetEventUrlOk

`func (o *MarketingEventCreateRequestParams) GetEventUrlOk() (*string, bool)`

GetEventUrlOk returns a tuple with the EventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventUrl

`func (o *MarketingEventCreateRequestParams) SetEventUrl(v string)`

SetEventUrl sets EventUrl field to given value.

### HasEventUrl

`func (o *MarketingEventCreateRequestParams) HasEventUrl() bool`

HasEventUrl returns a boolean if a field has been set.

### GetEventCancelled

`func (o *MarketingEventCreateRequestParams) GetEventCancelled() bool`

GetEventCancelled returns the EventCancelled field if non-nil, zero value otherwise.

### GetEventCancelledOk

`func (o *MarketingEventCreateRequestParams) GetEventCancelledOk() (*bool, bool)`

GetEventCancelledOk returns a tuple with the EventCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCancelled

`func (o *MarketingEventCreateRequestParams) SetEventCancelled(v bool)`

SetEventCancelled sets EventCancelled field to given value.

### HasEventCancelled

`func (o *MarketingEventCreateRequestParams) HasEventCancelled() bool`

HasEventCancelled returns a boolean if a field has been set.

### GetCustomProperties

`func (o *MarketingEventCreateRequestParams) GetCustomProperties() []PropertyValue`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *MarketingEventCreateRequestParams) GetCustomPropertiesOk() (*[]PropertyValue, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *MarketingEventCreateRequestParams) SetCustomProperties(v []PropertyValue)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *MarketingEventCreateRequestParams) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetExternalAccountId

`func (o *MarketingEventCreateRequestParams) GetExternalAccountId() string`

GetExternalAccountId returns the ExternalAccountId field if non-nil, zero value otherwise.

### GetExternalAccountIdOk

`func (o *MarketingEventCreateRequestParams) GetExternalAccountIdOk() (*string, bool)`

GetExternalAccountIdOk returns a tuple with the ExternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAccountId

`func (o *MarketingEventCreateRequestParams) SetExternalAccountId(v string)`

SetExternalAccountId sets ExternalAccountId field to given value.


### GetExternalEventId

`func (o *MarketingEventCreateRequestParams) GetExternalEventId() string`

GetExternalEventId returns the ExternalEventId field if non-nil, zero value otherwise.

### GetExternalEventIdOk

`func (o *MarketingEventCreateRequestParams) GetExternalEventIdOk() (*string, bool)`

GetExternalEventIdOk returns a tuple with the ExternalEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEventId

`func (o *MarketingEventCreateRequestParams) SetExternalEventId(v string)`

SetExternalEventId sets ExternalEventId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


