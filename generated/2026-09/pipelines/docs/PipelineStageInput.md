# PipelineStageInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayOrder** | **int32** | An integer indicating the position of this stage within the pipeline. Lower numbers indicate earlier stages. | 
**Label** | **string** | The name of the pipeline stage. It is a string that provides a human-readable identifier for the stage. | 
**Metadata** | **map[string]string** | An object containing additional properties for the stage. Each property is a key-value pair where both key and value are strings. | 
**StageId** | Pointer to **string** | A string that uniquely identifies the stage within the pipeline. | [optional] 

## Methods

### NewPipelineStageInput

`func NewPipelineStageInput(displayOrder int32, label string, metadata map[string]string, ) *PipelineStageInput`

NewPipelineStageInput instantiates a new PipelineStageInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineStageInputWithDefaults

`func NewPipelineStageInputWithDefaults() *PipelineStageInput`

NewPipelineStageInputWithDefaults instantiates a new PipelineStageInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetStageId

`func (o *PipelineStageInput) GetStageId() string`

GetStageId returns the StageId field if non-nil, zero value otherwise.

### GetStageIdOk

`func (o *PipelineStageInput) GetStageIdOk() (*string, bool)`

GetStageIdOk returns a tuple with the StageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageId

`func (o *PipelineStageInput) SetStageId(v string)`

SetStageId sets StageId field to given value.

### HasStageId

`func (o *PipelineStageInput) HasStageId() bool`

HasStageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


