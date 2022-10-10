# TimelineEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier for the event. This should be unique to the app and event template. If you use the same ID for different CRM objects, the last to be processed will win and the first will not have a record. You can also use &#x60;{{uuid}}&#x60; anywhere in the ID to generate a unique string, guaranteeing uniqueness. | 
**EventTemplateId** | **string** | The event template ID. | 
**Email** | Pointer to **string** | The email address used for contact-specific events. This can be used to identify existing contacts, create new ones, or change the email for an existing contact (if paired with the &#x60;objectId&#x60;). | [optional] 
**ObjectId** | Pointer to **string** | The CRM object identifier. This is required for every event other than contacts (where utk or email can be used). | [optional] 
**Utk** | Pointer to **string** | Use the &#x60;utk&#x60; parameter to associate an event with a contact by &#x60;usertoken&#x60;. This is recommended if you don&#39;t know a user&#39;s email, but have an identifying user token in your cookie. | [optional] 
**Domain** | Pointer to **string** | The event domain (often paired with utk). | [optional] 
**Timestamp** | Pointer to **time.Time** | The time the event occurred. If not passed in, the curren time will be assumed. This is used to determine where an event is shown on a CRM object&#39;s timeline. | [optional] 
**Tokens** | **map[string]string** | A collection of token keys and values associated with the template tokens. | 
**ExtraData** | Pointer to **map[string]interface{}** | Additional event-specific data that can be interpreted by the template&#39;s markdown. | [optional] 
**TimelineIFrame** | Pointer to [**TimelineEventIFrame**](TimelineEventIFrame.md) |  | [optional] 
**ObjectType** | **string** | The ObjectType associated with the EventTemplate. | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTimelineEventResponse

`func NewTimelineEventResponse(id string, eventTemplateId string, tokens map[string]string, objectType string, ) *TimelineEventResponse`

NewTimelineEventResponse instantiates a new TimelineEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineEventResponseWithDefaults

`func NewTimelineEventResponseWithDefaults() *TimelineEventResponse`

NewTimelineEventResponseWithDefaults instantiates a new TimelineEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimelineEventResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimelineEventResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimelineEventResponse) SetId(v string)`

SetId sets Id field to given value.


### GetEventTemplateId

`func (o *TimelineEventResponse) GetEventTemplateId() string`

GetEventTemplateId returns the EventTemplateId field if non-nil, zero value otherwise.

### GetEventTemplateIdOk

`func (o *TimelineEventResponse) GetEventTemplateIdOk() (*string, bool)`

GetEventTemplateIdOk returns a tuple with the EventTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTemplateId

`func (o *TimelineEventResponse) SetEventTemplateId(v string)`

SetEventTemplateId sets EventTemplateId field to given value.


### GetEmail

`func (o *TimelineEventResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TimelineEventResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TimelineEventResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TimelineEventResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetObjectId

`func (o *TimelineEventResponse) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *TimelineEventResponse) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *TimelineEventResponse) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *TimelineEventResponse) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetUtk

`func (o *TimelineEventResponse) GetUtk() string`

GetUtk returns the Utk field if non-nil, zero value otherwise.

### GetUtkOk

`func (o *TimelineEventResponse) GetUtkOk() (*string, bool)`

GetUtkOk returns a tuple with the Utk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtk

`func (o *TimelineEventResponse) SetUtk(v string)`

SetUtk sets Utk field to given value.

### HasUtk

`func (o *TimelineEventResponse) HasUtk() bool`

HasUtk returns a boolean if a field has been set.

### GetDomain

`func (o *TimelineEventResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *TimelineEventResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *TimelineEventResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *TimelineEventResponse) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetTimestamp

`func (o *TimelineEventResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TimelineEventResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TimelineEventResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TimelineEventResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTokens

`func (o *TimelineEventResponse) GetTokens() map[string]string`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *TimelineEventResponse) GetTokensOk() (*map[string]string, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *TimelineEventResponse) SetTokens(v map[string]string)`

SetTokens sets Tokens field to given value.


### GetExtraData

`func (o *TimelineEventResponse) GetExtraData() map[string]interface{}`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *TimelineEventResponse) GetExtraDataOk() (*map[string]interface{}, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *TimelineEventResponse) SetExtraData(v map[string]interface{})`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *TimelineEventResponse) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### GetTimelineIFrame

`func (o *TimelineEventResponse) GetTimelineIFrame() TimelineEventIFrame`

GetTimelineIFrame returns the TimelineIFrame field if non-nil, zero value otherwise.

### GetTimelineIFrameOk

`func (o *TimelineEventResponse) GetTimelineIFrameOk() (*TimelineEventIFrame, bool)`

GetTimelineIFrameOk returns a tuple with the TimelineIFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineIFrame

`func (o *TimelineEventResponse) SetTimelineIFrame(v TimelineEventIFrame)`

SetTimelineIFrame sets TimelineIFrame field to given value.

### HasTimelineIFrame

`func (o *TimelineEventResponse) HasTimelineIFrame() bool`

HasTimelineIFrame returns a boolean if a field has been set.

### GetObjectType

`func (o *TimelineEventResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TimelineEventResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TimelineEventResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *TimelineEventResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TimelineEventResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TimelineEventResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TimelineEventResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


