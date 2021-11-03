# PipelineStagePatchInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | A label used to organize pipeline stages in HubSpot&#39;s UI. Each pipeline stage&#39;s label must be unique within that pipeline. | [optional] 
**Archived** | Pointer to **bool** | Whether the pipeline is archived. | [optional] 
**DisplayOrder** | Pointer to **int32** | The order for displaying this pipeline stage. If two pipeline stages have a matching &#x60;displayOrder&#x60;, they will be sorted alphabetically by label. | [optional] 
**Metadata** | **map[string]string** | A JSON object containing properties that are not present on all object pipelines.  For &#x60;deals&#x60; pipelines, the &#x60;probability&#x60; field is required (&#x60;{ \&quot;probability\&quot;: 0.5 }&#x60;), and represents the likelihood a deal will close. Possible values are between 0.0 and 1.0 in increments of 0.1.  For &#x60;tickets&#x60; pipelines, the &#x60;ticketState&#x60; field is optional (&#x60;{ \&quot;ticketState\&quot;: \&quot;OPEN\&quot; }&#x60;), and represents whether the ticket remains open or has been closed by a member of your Support team. Possible values are &#x60;OPEN&#x60; or &#x60;CLOSED&#x60;. | 

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


