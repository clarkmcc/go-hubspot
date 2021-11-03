# TimelineEventTemplateTokenUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Used for list segmentation and reporting. | 
**ObjectPropertyName** | Pointer to **string** | The name of the CRM object property. This will populate the CRM object property associated with the event. With enough of these, you can fully build CRM objects via the Timeline API. | [optional] 
**Options** | [**[]TimelineEventTemplateTokenOption**](TimelineEventTemplateTokenOption.md) | If type is &#x60;enumeration&#x60;, we should have a list of options to choose from. | 

## Methods

### NewTimelineEventTemplateTokenUpdateRequest

`func NewTimelineEventTemplateTokenUpdateRequest(label string, options []TimelineEventTemplateTokenOption, ) *TimelineEventTemplateTokenUpdateRequest`

NewTimelineEventTemplateTokenUpdateRequest instantiates a new TimelineEventTemplateTokenUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineEventTemplateTokenUpdateRequestWithDefaults

`func NewTimelineEventTemplateTokenUpdateRequestWithDefaults() *TimelineEventTemplateTokenUpdateRequest`

NewTimelineEventTemplateTokenUpdateRequestWithDefaults instantiates a new TimelineEventTemplateTokenUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *TimelineEventTemplateTokenUpdateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TimelineEventTemplateTokenUpdateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TimelineEventTemplateTokenUpdateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetObjectPropertyName

`func (o *TimelineEventTemplateTokenUpdateRequest) GetObjectPropertyName() string`

GetObjectPropertyName returns the ObjectPropertyName field if non-nil, zero value otherwise.

### GetObjectPropertyNameOk

`func (o *TimelineEventTemplateTokenUpdateRequest) GetObjectPropertyNameOk() (*string, bool)`

GetObjectPropertyNameOk returns a tuple with the ObjectPropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectPropertyName

`func (o *TimelineEventTemplateTokenUpdateRequest) SetObjectPropertyName(v string)`

SetObjectPropertyName sets ObjectPropertyName field to given value.

### HasObjectPropertyName

`func (o *TimelineEventTemplateTokenUpdateRequest) HasObjectPropertyName() bool`

HasObjectPropertyName returns a boolean if a field has been set.

### GetOptions

`func (o *TimelineEventTemplateTokenUpdateRequest) GetOptions() []TimelineEventTemplateTokenOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *TimelineEventTemplateTokenUpdateRequest) GetOptionsOk() (*[]TimelineEventTemplateTokenOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *TimelineEventTemplateTokenUpdateRequest) SetOptions(v []TimelineEventTemplateTokenOption)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


