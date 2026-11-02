# PipelineStagePatchInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | Pointer to **bool** | A boolean indicating whether the pipeline stage is archived. | [optional] 
**DisplayOrder** | Pointer to **int32** | An integer specifying the order in which the stage appears relative to other stages. | [optional] 
**Label** | Pointer to **string** | A string representing the name or title of the pipeline stage. | [optional] 
**Metadata** | **map[string]string** | An object containing additional properties related to the pipeline stage, where each property is a string. | 

## Methods

### NewPipelineStagePatchInput

`func NewPipelineStagePatchInput(metadata map[string]string, ) *PipelineStagePatchInput`

NewPipelineStagePatchInput instantiates a new PipelineStagePatchInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineStagePatchInputWithDefaults

`func NewPipelineStagePatchInputWithDefaults() *PipelineStagePatchInput`

NewPipelineStagePatchInputWithDefaults instantiates a new PipelineStagePatchInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *PipelineStagePatchInput) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *PipelineStagePatchInput) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *PipelineStagePatchInput) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *PipelineStagePatchInput) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *PipelineStagePatchInput) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *PipelineStagePatchInput) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *PipelineStagePatchInput) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *PipelineStagePatchInput) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetLabel

`func (o *PipelineStagePatchInput) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PipelineStagePatchInput) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PipelineStagePatchInput) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PipelineStagePatchInput) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMetadata

`func (o *PipelineStagePatchInput) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PipelineStagePatchInput) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PipelineStagePatchInput) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


