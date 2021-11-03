# TimelineEventTemplateUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The template name. | 
**HeaderTemplate** | Pointer to **string** | This uses Markdown syntax with Handlebars and event-specific data to render HTML on a timeline as a header. | [optional] 
**DetailTemplate** | Pointer to **string** | This uses Markdown syntax with Handlebars and event-specific data to render HTML on a timeline when you expand the details. | [optional] 
**Tokens** | [**[]TimelineEventTemplateToken**](TimelineEventTemplateToken.md) | A collection of tokens that can be used as custom properties on the event and to create fully fledged CRM objects. | 
**Id** | **string** | The template ID. | 

## Methods

### NewTimelineEventTemplateUpdateRequest

`func NewTimelineEventTemplateUpdateRequest(name string, tokens []TimelineEventTemplateToken, id string, ) *TimelineEventTemplateUpdateRequest`

NewTimelineEventTemplateUpdateRequest instantiates a new TimelineEventTemplateUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineEventTemplateUpdateRequestWithDefaults

`func NewTimelineEventTemplateUpdateRequestWithDefaults() *TimelineEventTemplateUpdateRequest`

NewTimelineEventTemplateUpdateRequestWithDefaults instantiates a new TimelineEventTemplateUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TimelineEventTemplateUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TimelineEventTemplateUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TimelineEventTemplateUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetHeaderTemplate

`func (o *TimelineEventTemplateUpdateRequest) GetHeaderTemplate() string`

GetHeaderTemplate returns the HeaderTemplate field if non-nil, zero value otherwise.

### GetHeaderTemplateOk

`func (o *TimelineEventTemplateUpdateRequest) GetHeaderTemplateOk() (*string, bool)`

GetHeaderTemplateOk returns a tuple with the HeaderTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderTemplate

`func (o *TimelineEventTemplateUpdateRequest) SetHeaderTemplate(v string)`

SetHeaderTemplate sets HeaderTemplate field to given value.

### HasHeaderTemplate

`func (o *TimelineEventTemplateUpdateRequest) HasHeaderTemplate() bool`

HasHeaderTemplate returns a boolean if a field has been set.

### GetDetailTemplate

`func (o *TimelineEventTemplateUpdateRequest) GetDetailTemplate() string`

GetDetailTemplate returns the DetailTemplate field if non-nil, zero value otherwise.

### GetDetailTemplateOk

`func (o *TimelineEventTemplateUpdateRequest) GetDetailTemplateOk() (*string, bool)`

GetDetailTemplateOk returns a tuple with the DetailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailTemplate

`func (o *TimelineEventTemplateUpdateRequest) SetDetailTemplate(v string)`

SetDetailTemplate sets DetailTemplate field to given value.

### HasDetailTemplate

`func (o *TimelineEventTemplateUpdateRequest) HasDetailTemplate() bool`

HasDetailTemplate returns a boolean if a field has been set.

### GetTokens

`func (o *TimelineEventTemplateUpdateRequest) GetTokens() []TimelineEventTemplateToken`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *TimelineEventTemplateUpdateRequest) GetTokensOk() (*[]TimelineEventTemplateToken, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *TimelineEventTemplateUpdateRequest) SetTokens(v []TimelineEventTemplateToken)`

SetTokens sets Tokens field to given value.


### GetId

`func (o *TimelineEventTemplateUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimelineEventTemplateUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimelineEventTemplateUpdateRequest) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


