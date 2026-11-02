# PipelineReplaceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayOrder** | **int32** | An integer indicating the order in which the pipeline appears relative to other pipelines. | 
**Label** | **string** | A string representing the name of the pipeline. | 
**Stages** | [**[]PipelineStageInput**](PipelineStageInput.md) | An array of PipelineStageInput objects, each representing a stage within the pipeline. | 

## Methods

### NewPipelineReplaceInput

`func NewPipelineReplaceInput(displayOrder int32, label string, stages []PipelineStageInput, ) *PipelineReplaceInput`

NewPipelineReplaceInput instantiates a new PipelineReplaceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineReplaceInputWithDefaults

`func NewPipelineReplaceInputWithDefaults() *PipelineReplaceInput`

NewPipelineReplaceInputWithDefaults instantiates a new PipelineReplaceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayOrder

`func (o *PipelineReplaceInput) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *PipelineReplaceInput) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *PipelineReplaceInput) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.


### GetLabel

`func (o *PipelineReplaceInput) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PipelineReplaceInput) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PipelineReplaceInput) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetStages

`func (o *PipelineReplaceInput) GetStages() []PipelineStageInput`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *PipelineReplaceInput) GetStagesOk() (*[]PipelineStageInput, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *PipelineReplaceInput) SetStages(v []PipelineStageInput)`

SetStages sets Stages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


