# PipelinePatchInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | A unique label used to organize pipelines in HubSpot&#39;s UI | [optional] 
**DisplayOrder** | Pointer to **int32** | The order for displaying this pipeline. If two pipelines have a matching &#x60;displayOrder&#x60;, they will be sorted alphabetically by label. | [optional] 
**Archived** | Pointer to **bool** | Whether the pipeline is archived. This property should only be provided when restoring an archived pipeline. If it&#39;s provided in any other call, the request will fail and a &#x60;400 Bad Request&#x60; will be returned. | [optional] 

## Methods

### NewPipelinePatchInput

`func NewPipelinePatchInput() *PipelinePatchInput`

NewPipelinePatchInput instantiates a new PipelinePatchInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelinePatchInputWithDefaults

`func NewPipelinePatchInputWithDefaults() *PipelinePatchInput`

NewPipelinePatchInputWithDefaults instantiates a new PipelinePatchInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *PipelinePatchInput) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PipelinePatchInput) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PipelinePatchInput) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PipelinePatchInput) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *PipelinePatchInput) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *PipelinePatchInput) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *PipelinePatchInput) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *PipelinePatchInput) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetArchived

`func (o *PipelinePatchInput) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *PipelinePatchInput) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *PipelinePatchInput) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *PipelinePatchInput) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


