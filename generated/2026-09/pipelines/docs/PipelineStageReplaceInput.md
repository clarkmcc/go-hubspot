# PipelineStageReplaceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayOrder** | **int32** | An integer indicating the position of the stage within the pipeline. | 
**Label** | **string** | The name of the pipeline stage, represented as a string. | 
**Metadata** | **map[string]string** | An object containing additional properties related to the stage, where each property is a string. | 

## Methods

### NewPipelineStageReplaceInput

`func NewPipelineStageReplaceInput(displayOrder int32, label string, metadata map[string]string, ) *PipelineStageReplaceInput`

NewPipelineStageReplaceInput instantiates a new PipelineStageReplaceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineStageReplaceInputWithDefaults

`func NewPipelineStageReplaceInputWithDefaults() *PipelineStageReplaceInput`

NewPipelineStageReplaceInputWithDefaults instantiates a new PipelineStageReplaceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayOrder

`func (o *PipelineStageReplaceInput) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *PipelineStageReplaceInput) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *PipelineStageReplaceInput) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.


### GetLabel

`func (o *PipelineStageReplaceInput) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PipelineStageReplaceInput) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PipelineStageReplaceInput) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMetadata

`func (o *PipelineStageReplaceInput) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PipelineStageReplaceInput) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PipelineStageReplaceInput) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


