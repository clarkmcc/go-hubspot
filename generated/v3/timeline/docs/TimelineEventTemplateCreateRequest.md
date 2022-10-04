# TimelineEventTemplateCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The template name. | 
**HeaderTemplate** | Pointer to **string** | This uses Markdown syntax with Handlebars and event-specific data to render HTML on a timeline as a header. | [optional] 
**DetailTemplate** | Pointer to **string** | This uses Markdown syntax with Handlebars and event-specific data to render HTML on a timeline when you expand the details. | [optional] 
**Tokens** | [**[]TimelineEventTemplateToken**](TimelineEventTemplateToken.md) | A collection of tokens that can be used as custom properties on the event and to create fully fledged CRM objects. | 
**ObjectType** | **string** | The type of CRM object this template is for. [Contacts, companies, tickets, and deals] are supported. | 

## Methods

### NewTimelineEventTemplateCreateRequest

`func NewTimelineEventTemplateCreateRequest(name string, tokens []TimelineEventTemplateToken, objectType string, ) *TimelineEventTemplateCreateRequest`

NewTimelineEventTemplateCreateRequest instantiates a new TimelineEventTemplateCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineEventTemplateCreateRequestWithDefaults

`func NewTimelineEventTemplateCreateRequestWithDefaults() *TimelineEventTemplateCreateRequest`

NewTimelineEventTemplateCreateRequestWithDefaults instantiates a new TimelineEventTemplateCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TimelineEventTemplateCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TimelineEventTemplateCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TimelineEventTemplateCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetHeaderTemplate

`func (o *TimelineEventTemplateCreateRequest) GetHeaderTemplate() string`

GetHeaderTemplate returns the HeaderTemplate field if non-nil, zero value otherwise.

### GetHeaderTemplateOk

`func (o *TimelineEventTemplateCreateRequest) GetHeaderTemplateOk() (*string, bool)`

GetHeaderTemplateOk returns a tuple with the HeaderTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderTemplate

`func (o *TimelineEventTemplateCreateRequest) SetHeaderTemplate(v string)`

SetHeaderTemplate sets HeaderTemplate field to given value.

### HasHeaderTemplate

`func (o *TimelineEventTemplateCreateRequest) HasHeaderTemplate() bool`

HasHeaderTemplate returns a boolean if a field has been set.

### GetDetailTemplate

`func (o *TimelineEventTemplateCreateRequest) GetDetailTemplate() string`

GetDetailTemplate returns the DetailTemplate field if non-nil, zero value otherwise.

### GetDetailTemplateOk

`func (o *TimelineEventTemplateCreateRequest) GetDetailTemplateOk() (*string, bool)`

GetDetailTemplateOk returns a tuple with the DetailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailTemplate

`func (o *TimelineEventTemplateCreateRequest) SetDetailTemplate(v string)`

SetDetailTemplate sets DetailTemplate field to given value.

### HasDetailTemplate

`func (o *TimelineEventTemplateCreateRequest) HasDetailTemplate() bool`

HasDetailTemplate returns a boolean if a field has been set.

### GetTokens

`func (o *TimelineEventTemplateCreateRequest) GetTokens() []TimelineEventTemplateToken`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *TimelineEventTemplateCreateRequest) GetTokensOk() (*[]TimelineEventTemplateToken, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *TimelineEventTemplateCreateRequest) SetTokens(v []TimelineEventTemplateToken)`

SetTokens sets Tokens field to given value.


### GetObjectType

`func (o *TimelineEventTemplateCreateRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TimelineEventTemplateCreateRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TimelineEventTemplateCreateRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


