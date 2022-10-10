# TimelineEventTemplateToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Used for list segmentation and reporting. | 
**ObjectPropertyName** | Pointer to **string** | The name of the CRM object property. This will populate the CRM object property associated with the event. With enough of these, you can fully build CRM objects via the Timeline API. | [optional] 
**Options** | [**[]TimelineEventTemplateTokenOption**](TimelineEventTemplateTokenOption.md) | If type is &#x60;enumeration&#x60;, we should have a list of options to choose from. | 
**Name** | **string** | The name of the token referenced in the templates. This must be unique for the specific template. It may only contain alphanumeric characters, periods, dashes, or underscores (. - _). | 
**Type** | **string** | The data type of the token. You can currently choose from [string, number, date, enumeration]. | 
**CreatedAt** | Pointer to **time.Time** | The date and time that the Event Template Token was created, as an ISO 8601 timestamp. Will be null if the template was created before Feb 18th, 2020. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time that the Event Template Token was last updated, as an ISO 8601 timestamp. Will be null if the template was created before Feb 18th, 2020. | [optional] 

## Methods

### NewTimelineEventTemplateToken

`func NewTimelineEventTemplateToken(label string, options []TimelineEventTemplateTokenOption, name string, type_ string, ) *TimelineEventTemplateToken`

NewTimelineEventTemplateToken instantiates a new TimelineEventTemplateToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineEventTemplateTokenWithDefaults

`func NewTimelineEventTemplateTokenWithDefaults() *TimelineEventTemplateToken`

NewTimelineEventTemplateTokenWithDefaults instantiates a new TimelineEventTemplateToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *TimelineEventTemplateToken) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TimelineEventTemplateToken) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TimelineEventTemplateToken) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetObjectPropertyName

`func (o *TimelineEventTemplateToken) GetObjectPropertyName() string`

GetObjectPropertyName returns the ObjectPropertyName field if non-nil, zero value otherwise.

### GetObjectPropertyNameOk

`func (o *TimelineEventTemplateToken) GetObjectPropertyNameOk() (*string, bool)`

GetObjectPropertyNameOk returns a tuple with the ObjectPropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectPropertyName

`func (o *TimelineEventTemplateToken) SetObjectPropertyName(v string)`

SetObjectPropertyName sets ObjectPropertyName field to given value.

### HasObjectPropertyName

`func (o *TimelineEventTemplateToken) HasObjectPropertyName() bool`

HasObjectPropertyName returns a boolean if a field has been set.

### GetOptions

`func (o *TimelineEventTemplateToken) GetOptions() []TimelineEventTemplateTokenOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *TimelineEventTemplateToken) GetOptionsOk() (*[]TimelineEventTemplateTokenOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *TimelineEventTemplateToken) SetOptions(v []TimelineEventTemplateTokenOption)`

SetOptions sets Options field to given value.


### GetName

`func (o *TimelineEventTemplateToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TimelineEventTemplateToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TimelineEventTemplateToken) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *TimelineEventTemplateToken) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimelineEventTemplateToken) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimelineEventTemplateToken) SetType(v string)`

SetType sets Type field to given value.


### GetCreatedAt

`func (o *TimelineEventTemplateToken) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TimelineEventTemplateToken) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TimelineEventTemplateToken) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TimelineEventTemplateToken) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TimelineEventTemplateToken) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TimelineEventTemplateToken) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TimelineEventTemplateToken) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TimelineEventTemplateToken) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


