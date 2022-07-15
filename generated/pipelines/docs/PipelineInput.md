# PipelineInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | A unique label used to organize pipelines in HubSpot&#39;s UI | 
**DisplayOrder** | **int32** | The order for displaying this pipeline. If two pipelines have a matching &#x60;displayOrder&#x60;, they will be sorted alphabetically by label. | 
**Stages** | [**[]PipelineStageInput**](PipelineStageInput.md) | Pipeline stage inputs used to create the new or replacement pipeline. | 

## Methods

### NewPipelineInput

`func NewPipelineInput(label string, displayOrder int32, stages []PipelineStageInput, ) *PipelineInput`

NewPipelineInput instantiates a new PipelineInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineInputWithDefaults

`func NewPipelineInputWithDefaults() *PipelineInput`

NewPipelineInputWithDefaults instantiates a new PipelineInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *PipelineInput) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PipelineInput) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PipelineInput) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDisplayOrder

`func (o *PipelineInput) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *PipelineInput) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *PipelineInput) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.


### GetStages

`func (o *PipelineInput) GetStages() []PipelineStageInput`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *PipelineInput) GetStagesOk() (*[]PipelineStageInput, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *PipelineInput) SetStages(v []PipelineStageInput)`

SetStages sets Stages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


