# PipelineInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayOrder** | **int32** | An integer indicating the order in which the pipeline appears relative to other pipelines. | 
**Label** | **string** | A string representing the name or title of the pipeline. | 
**PipelineId** | Pointer to **string** | A string that uniquely identifies the pipeline. This property is optional when creating a new pipeline. | [optional] 
**Stages** | [**[]PipelineStageInput**](PipelineStageInput.md) | An array of PipelineStageInput objects, each representing a stage within the pipeline. | 

## Methods

### NewPipelineInput

`func NewPipelineInput(displayOrder int32, label string, stages []PipelineStageInput, ) *PipelineInput`

NewPipelineInput instantiates a new PipelineInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineInputWithDefaults

`func NewPipelineInputWithDefaults() *PipelineInput`

NewPipelineInputWithDefaults instantiates a new PipelineInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetPipelineId

`func (o *PipelineInput) GetPipelineId() string`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *PipelineInput) GetPipelineIdOk() (*string, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *PipelineInput) SetPipelineId(v string)`

SetPipelineId sets PipelineId field to given value.

### HasPipelineId

`func (o *PipelineInput) HasPipelineId() bool`

HasPipelineId returns a boolean if a field has been set.

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


