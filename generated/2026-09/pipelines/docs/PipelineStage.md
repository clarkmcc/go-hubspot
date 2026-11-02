# PipelineStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | **bool** | A boolean indicating whether the stage is archived. | 
**ArchivedAt** | Pointer to **time.Time** | The date and time when the stage was archived, in ISO 8601 format. | [optional] 
**CreatedAt** | **time.Time** | The date and time when the stage was created, in ISO 8601 format. | 
**DisplayOrder** | **int32** | An integer indicating the position of this stage within the pipeline, determining its order relative to other stages. | 
**Id** | **string** | A unique string identifier for the pipeline stage. | 
**Label** | **string** | The name or title of the pipeline stage, represented as a string. | 
**Metadata** | **map[string]string** | An object containing additional metadata for the stage, with string values for each property. | 
**UpdatedAt** | **time.Time** | The date and time when the stage was last updated, in ISO 8601 format. | 
**WritePermissions** | **string** | A string indicating the write permissions for the stage. Valid values include &#39;CRM_PERMISSIONS_ENFORCEMENT&#39;, &#39;READ_ONLY&#39;, and &#39;INTERNAL_ONLY&#39;. | 

## Methods

### NewPipelineStage

`func NewPipelineStage(archived bool, createdAt time.Time, displayOrder int32, id string, label string, metadata map[string]string, updatedAt time.Time, writePermissions string, ) *PipelineStage`

NewPipelineStage instantiates a new PipelineStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineStageWithDefaults

`func NewPipelineStageWithDefaults() *PipelineStage`

NewPipelineStageWithDefaults instantiates a new PipelineStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *PipelineStage) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *PipelineStage) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *PipelineStage) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetArchivedAt

`func (o *PipelineStage) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *PipelineStage) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *PipelineStage) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *PipelineStage) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PipelineStage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PipelineStage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PipelineStage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDisplayOrder

`func (o *PipelineStage) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *PipelineStage) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *PipelineStage) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.


### GetId

`func (o *PipelineStage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PipelineStage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PipelineStage) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *PipelineStage) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PipelineStage) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PipelineStage) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMetadata

`func (o *PipelineStage) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PipelineStage) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PipelineStage) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.


### GetUpdatedAt

`func (o *PipelineStage) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PipelineStage) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PipelineStage) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetWritePermissions

`func (o *PipelineStage) GetWritePermissions() string`

GetWritePermissions returns the WritePermissions field if non-nil, zero value otherwise.

### GetWritePermissionsOk

`func (o *PipelineStage) GetWritePermissionsOk() (*string, bool)`

GetWritePermissionsOk returns a tuple with the WritePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritePermissions

`func (o *PipelineStage) SetWritePermissions(v string)`

SetWritePermissions sets WritePermissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


