# Pipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | **bool** | A boolean indicating whether the pipeline is archived. | 
**ArchivedAt** | Pointer to **time.Time** | The date and time when the pipeline was archived, in ISO 8601 format. | [optional] 
**CreatedAt** | **time.Time** | The date and time when the pipeline was created, in ISO 8601 format. | 
**DisplayOrder** | **int32** | An integer indicating the order in which the pipeline appears relative to other pipelines. | 
**Id** | **string** | A unique string identifier for the pipeline. | 
**Label** | **string** | The name or title of the pipeline, represented as a string. | 
**Stages** | [**[]PipelineStage**](PipelineStage.md) | An array of PipelineStage objects that define the stages within this pipeline. | 
**UpdatedAt** | **time.Time** | The date and time when the pipeline was last updated, in ISO 8601 format. | 

## Methods

### NewPipeline

`func NewPipeline(archived bool, createdAt time.Time, displayOrder int32, id string, label string, stages []PipelineStage, updatedAt time.Time, ) *Pipeline`

NewPipeline instantiates a new Pipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineWithDefaults

`func NewPipelineWithDefaults() *Pipeline`

NewPipelineWithDefaults instantiates a new Pipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *Pipeline) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Pipeline) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Pipeline) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetArchivedAt

`func (o *Pipeline) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *Pipeline) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *Pipeline) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *Pipeline) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Pipeline) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Pipeline) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Pipeline) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDisplayOrder

`func (o *Pipeline) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *Pipeline) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *Pipeline) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.


### GetId

`func (o *Pipeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Pipeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Pipeline) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *Pipeline) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Pipeline) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Pipeline) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetStages

`func (o *Pipeline) GetStages() []PipelineStage`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *Pipeline) GetStagesOk() (*[]PipelineStage, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *Pipeline) SetStages(v []PipelineStage)`

SetStages sets Stages field to given value.


### GetUpdatedAt

`func (o *Pipeline) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Pipeline) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Pipeline) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


