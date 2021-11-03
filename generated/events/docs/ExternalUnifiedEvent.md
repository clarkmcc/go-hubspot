# ExternalUnifiedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** | The objectType for the object which did the event. | 
**ObjectId** | **string** | The objectId of the object which did the event. | 
**EventType** | **string** | The format of the &#x60;eventType&#x60; string is &#x60;ae{appId}_{eventTypeLabel}&#x60;, &#x60;pe{portalId}_{eventTypeLabel}&#x60;, or just &#x60;e_{eventTypeLabel}&#x60; for HubSpot events. | 
**OccurredAt** | **time.Time** | An ISO 8601 timestamp when the event occurred. | 
**Id** | **string** | A unique identifier for the event. | 
**Properties** | **map[string]string** |  | 

## Methods

### NewExternalUnifiedEvent

`func NewExternalUnifiedEvent(objectType string, objectId string, eventType string, occurredAt time.Time, id string, properties map[string]string, ) *ExternalUnifiedEvent`

NewExternalUnifiedEvent instantiates a new ExternalUnifiedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalUnifiedEventWithDefaults

`func NewExternalUnifiedEventWithDefaults() *ExternalUnifiedEvent`

NewExternalUnifiedEventWithDefaults instantiates a new ExternalUnifiedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *ExternalUnifiedEvent) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ExternalUnifiedEvent) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ExternalUnifiedEvent) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *ExternalUnifiedEvent) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ExternalUnifiedEvent) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ExternalUnifiedEvent) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetEventType

`func (o *ExternalUnifiedEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ExternalUnifiedEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ExternalUnifiedEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetOccurredAt

`func (o *ExternalUnifiedEvent) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *ExternalUnifiedEvent) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *ExternalUnifiedEvent) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.


### GetId

`func (o *ExternalUnifiedEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalUnifiedEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalUnifiedEvent) SetId(v string)`

SetId sets Id field to given value.


### GetProperties

`func (o *ExternalUnifiedEvent) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ExternalUnifiedEvent) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ExternalUnifiedEvent) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


