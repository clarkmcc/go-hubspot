# TimelineEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTemplateId** | **string** | The event template ID. | 
**Email** | Pointer to **string** | The email address used for contact-specific events. This can be used to identify existing contacts, create new ones, or change the email for an existing contact (if paired with the &#x60;objectId&#x60;). | [optional] 
**ObjectId** | Pointer to **string** | The CRM object identifier. This is required for every event other than contacts (where utk or email can be used). | [optional] 
**Utk** | Pointer to **string** | Use the &#x60;utk&#x60; parameter to associate an event with a contact by &#x60;usertoken&#x60;. This is recommended if you don&#39;t know a user&#39;s email, but have an identifying user token in your cookie. | [optional] 
**Domain** | Pointer to **string** | The event domain (often paired with utk). | [optional] 
**Timestamp** | Pointer to **time.Time** | The time the event occurred. If not passed in, the curren time will be assumed. This is used to determine where an event is shown on a CRM object&#39;s timeline. | [optional] 
**Tokens** | **map[string]string** | A collection of token keys and values associated with the template tokens. | 
**ExtraData** | Pointer to **map[string]interface{}** | Additional event-specific data that can be interpreted by the template&#39;s markdown. | [optional] 
**TimelineIFrame** | Pointer to [**TimelineEventIFrame**](TimelineEventIFrame.md) |  | [optional] 
**Id** | Pointer to **string** | Identifier for the event. This is optional, and we recommend you do not pass this in. We will create one for you if you omit this. You can also use &#x60;{{uuid}}&#x60; anywhere in the ID to generate a unique string, guaranteeing uniqueness. | [optional] 

## Methods

### NewTimelineEvent

`func NewTimelineEvent(eventTemplateId string, tokens map[string]string, ) *TimelineEvent`

NewTimelineEvent instantiates a new TimelineEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineEventWithDefaults

`func NewTimelineEventWithDefaults() *TimelineEvent`

NewTimelineEventWithDefaults instantiates a new TimelineEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventTemplateId

`func (o *TimelineEvent) GetEventTemplateId() string`

GetEventTemplateId returns the EventTemplateId field if non-nil, zero value otherwise.

### GetEventTemplateIdOk

`func (o *TimelineEvent) GetEventTemplateIdOk() (*string, bool)`

GetEventTemplateIdOk returns a tuple with the EventTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTemplateId

`func (o *TimelineEvent) SetEventTemplateId(v string)`

SetEventTemplateId sets EventTemplateId field to given value.


### GetEmail

`func (o *TimelineEvent) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TimelineEvent) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TimelineEvent) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TimelineEvent) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetObjectId

`func (o *TimelineEvent) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *TimelineEvent) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *TimelineEvent) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *TimelineEvent) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetUtk

`func (o *TimelineEvent) GetUtk() string`

GetUtk returns the Utk field if non-nil, zero value otherwise.

### GetUtkOk

`func (o *TimelineEvent) GetUtkOk() (*string, bool)`

GetUtkOk returns a tuple with the Utk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtk

`func (o *TimelineEvent) SetUtk(v string)`

SetUtk sets Utk field to given value.

### HasUtk

`func (o *TimelineEvent) HasUtk() bool`

HasUtk returns a boolean if a field has been set.

### GetDomain

`func (o *TimelineEvent) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *TimelineEvent) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *TimelineEvent) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *TimelineEvent) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetTimestamp

`func (o *TimelineEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TimelineEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TimelineEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TimelineEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTokens

`func (o *TimelineEvent) GetTokens() map[string]string`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *TimelineEvent) GetTokensOk() (*map[string]string, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *TimelineEvent) SetTokens(v map[string]string)`

SetTokens sets Tokens field to given value.


### GetExtraData

`func (o *TimelineEvent) GetExtraData() map[string]interface{}`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *TimelineEvent) GetExtraDataOk() (*map[string]interface{}, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *TimelineEvent) SetExtraData(v map[string]interface{})`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *TimelineEvent) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### GetTimelineIFrame

`func (o *TimelineEvent) GetTimelineIFrame() TimelineEventIFrame`

GetTimelineIFrame returns the TimelineIFrame field if non-nil, zero value otherwise.

### GetTimelineIFrameOk

`func (o *TimelineEvent) GetTimelineIFrameOk() (*TimelineEventIFrame, bool)`

GetTimelineIFrameOk returns a tuple with the TimelineIFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineIFrame

`func (o *TimelineEvent) SetTimelineIFrame(v TimelineEventIFrame)`

SetTimelineIFrame sets TimelineIFrame field to given value.

### HasTimelineIFrame

`func (o *TimelineEvent) HasTimelineIFrame() bool`

HasTimelineIFrame returns a boolean if a field has been set.

### GetId

`func (o *TimelineEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimelineEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimelineEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TimelineEvent) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


