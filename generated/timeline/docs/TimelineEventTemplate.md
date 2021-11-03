# TimelineEventTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The template name. | 
**HeaderTemplate** | Pointer to **string** | This uses Markdown syntax with Handlebars and event-specific data to render HTML on a timeline as a header. | [optional] 
**DetailTemplate** | Pointer to **string** | This uses Markdown syntax with Handlebars and event-specific data to render HTML on a timeline when you expand the details. | [optional] 
**Tokens** | [**[]TimelineEventTemplateToken**](TimelineEventTemplateToken.md) | A collection of tokens that can be used as custom properties on the event and to create fully fledged CRM objects. | 
**Id** | **string** | The template ID. | 
**ObjectType** | **string** | The type of CRM object this template is for. [Contacts, companies, tickets, and deals] are supported. | 
**CreatedAt** | Pointer to **time.Time** | The date and time that the Event Template was created, as an ISO 8601 timestamp. Will be null if the template was created before Feb 18th, 2020. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time that the Event Template was last updated, as an ISO 8601 timestamp. Will be null if the template was created before Feb 18th, 2020. | [optional] 

## Methods

### NewTimelineEventTemplate

`func NewTimelineEventTemplate(name string, tokens []TimelineEventTemplateToken, id string, objectType string, ) *TimelineEventTemplate`

NewTimelineEventTemplate instantiates a new TimelineEventTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineEventTemplateWithDefaults

`func NewTimelineEventTemplateWithDefaults() *TimelineEventTemplate`

NewTimelineEventTemplateWithDefaults instantiates a new TimelineEventTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TimelineEventTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TimelineEventTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TimelineEventTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetHeaderTemplate

`func (o *TimelineEventTemplate) GetHeaderTemplate() string`

GetHeaderTemplate returns the HeaderTemplate field if non-nil, zero value otherwise.

### GetHeaderTemplateOk

`func (o *TimelineEventTemplate) GetHeaderTemplateOk() (*string, bool)`

GetHeaderTemplateOk returns a tuple with the HeaderTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderTemplate

`func (o *TimelineEventTemplate) SetHeaderTemplate(v string)`

SetHeaderTemplate sets HeaderTemplate field to given value.

### HasHeaderTemplate

`func (o *TimelineEventTemplate) HasHeaderTemplate() bool`

HasHeaderTemplate returns a boolean if a field has been set.

### GetDetailTemplate

`func (o *TimelineEventTemplate) GetDetailTemplate() string`

GetDetailTemplate returns the DetailTemplate field if non-nil, zero value otherwise.

### GetDetailTemplateOk

`func (o *TimelineEventTemplate) GetDetailTemplateOk() (*string, bool)`

GetDetailTemplateOk returns a tuple with the DetailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailTemplate

`func (o *TimelineEventTemplate) SetDetailTemplate(v string)`

SetDetailTemplate sets DetailTemplate field to given value.

### HasDetailTemplate

`func (o *TimelineEventTemplate) HasDetailTemplate() bool`

HasDetailTemplate returns a boolean if a field has been set.

### GetTokens

`func (o *TimelineEventTemplate) GetTokens() []TimelineEventTemplateToken`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *TimelineEventTemplate) GetTokensOk() (*[]TimelineEventTemplateToken, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *TimelineEventTemplate) SetTokens(v []TimelineEventTemplateToken)`

SetTokens sets Tokens field to given value.


### GetId

`func (o *TimelineEventTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimelineEventTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimelineEventTemplate) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *TimelineEventTemplate) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TimelineEventTemplate) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TimelineEventTemplate) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *TimelineEventTemplate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TimelineEventTemplate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TimelineEventTemplate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TimelineEventTemplate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TimelineEventTemplate) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TimelineEventTemplate) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TimelineEventTemplate) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TimelineEventTemplate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


