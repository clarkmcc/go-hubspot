# PipelineStageInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | A label used to organize pipeline stages in HubSpot&#39;s UI. Each pipeline stage&#39;s label must be unique within that pipeline. | 
**DisplayOrder** | **int32** | The order for displaying this pipeline stage. If two pipeline stages have a matching &#x60;displayOrder&#x60;, they will be sorted alphabetically by label. | 
**Metadata** | **map[string]string** | A JSON object containing properties that are not present on all object pipelines.  For &#x60;deals&#x60; pipelines, the &#x60;probability&#x60; field is required (&#x60;{ \&quot;probability\&quot;: 0.5 }&#x60;), and represents the likelihood a deal will close. Possible values are between 0.0 and 1.0 in increments of 0.1.  For &#x60;tickets&#x60; pipelines, the &#x60;ticketState&#x60; field is optional (&#x60;{ \&quot;ticketState\&quot;: \&quot;OPEN\&quot; }&#x60;), and represents whether the ticket remains open or has been closed by a member of your Support team. Possible values are &#x60;OPEN&#x60; or &#x60;CLOSED&#x60;. | 

## Methods

### NewPipelineStageInput

`func NewPipelineStageInput(label string, displayOrder int32, metadata map[string]string, ) *PipelineStageInput`

NewPipelineStageInput instantiates a new PipelineStageInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineStageInputWithDefaults

`func NewPipelineStageInputWithDefaults() *PipelineStageInput`

NewPipelineStageInputWithDefaults instantiates a new PipelineStageInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *PipelineStageInput) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PipelineStageInput) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PipelineStageInput) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDisplayOrder

`func (o *PipelineStageInput) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *PipelineStageInput) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *PipelineStageInput) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.


### GetMetadata

`func (o *PipelineStageInput) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PipelineStageInput) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PipelineStageInput) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


