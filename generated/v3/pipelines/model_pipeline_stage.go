/*
Pipelines

Pipelines represent distinct stages in a workflow, like closing a deal or servicing a support ticket. These endpoints provide access to read and modify pipelines in HubSpot. Pipelines support `deals` and `tickets` object types.  ## Pipeline ID validation  When calling endpoints that take pipelineId as a parameter, that ID must correspond to an existing, un-archived pipeline. Otherwise the request will fail with a `404 Not Found` response.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipelines

import (
	"encoding/json"
	"time"
)

// PipelineStage A pipeline stage definition.
type PipelineStage struct {
	// The date the pipeline stage was created. The stages on default pipelines will have createdAt = 0.
	CreatedAt time.Time `json:"createdAt"`
	// The date the pipeline was archived. `archivedAt` will only be present if the pipeline is archived.
	ArchivedAt *time.Time `json:"archivedAt,omitempty"`
	// Whether the pipeline is archived.
	Archived bool `json:"archived"`
	// A JSON object containing properties that are not present on all object pipelines.  For `deals` pipelines, the `probability` field is required (`{ \"probability\": 0.5 }`), and represents the likelihood a deal will close. Possible values are between 0.0 and 1.0 in increments of 0.1.  For `tickets` pipelines, the `ticketState` field is optional (`{ \"ticketState\": \"OPEN\" }`), and represents whether the ticket remains open or has been closed by a member of your Support team. Possible values are `OPEN` or `CLOSED`.
	Metadata map[string]string `json:"metadata"`
	// The order for displaying this pipeline stage. If two pipeline stages have a matching `displayOrder`, they will be sorted alphabetically by label.
	DisplayOrder     int32   `json:"displayOrder"`
	WritePermissions *string `json:"writePermissions,omitempty"`
	// A label used to organize pipeline stages in HubSpot's UI. Each pipeline stage's label must be unique within that pipeline.
	Label string `json:"label"`
	// A unique identifier generated by HubSpot that can be used to retrieve and update the pipeline stage.
	Id string `json:"id"`
	// The date the pipeline stage was last updated.
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewPipelineStage instantiates a new PipelineStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineStage(createdAt time.Time, archived bool, metadata map[string]string, displayOrder int32, label string, id string, updatedAt time.Time) *PipelineStage {
	this := PipelineStage{}
	this.CreatedAt = createdAt
	this.Archived = archived
	this.Metadata = metadata
	this.DisplayOrder = displayOrder
	this.Label = label
	this.Id = id
	this.UpdatedAt = updatedAt
	return &this
}

// NewPipelineStageWithDefaults instantiates a new PipelineStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineStageWithDefaults() *PipelineStage {
	this := PipelineStage{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *PipelineStage) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PipelineStage) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PipelineStage) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *PipelineStage) GetArchivedAt() time.Time {
	if o == nil || o.ArchivedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineStage) GetArchivedAtOk() (*time.Time, bool) {
	if o == nil || o.ArchivedAt == nil {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *PipelineStage) HasArchivedAt() bool {
	if o != nil && o.ArchivedAt != nil {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given time.Time and assigns it to the ArchivedAt field.
func (o *PipelineStage) SetArchivedAt(v time.Time) {
	o.ArchivedAt = &v
}

// GetArchived returns the Archived field value
func (o *PipelineStage) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *PipelineStage) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *PipelineStage) SetArchived(v bool) {
	o.Archived = v
}

// GetMetadata returns the Metadata field value
func (o *PipelineStage) GetMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *PipelineStage) GetMetadataOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *PipelineStage) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetDisplayOrder returns the DisplayOrder field value
func (o *PipelineStage) GetDisplayOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value
// and a boolean to check if the value has been set.
func (o *PipelineStage) GetDisplayOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayOrder, true
}

// SetDisplayOrder sets field value
func (o *PipelineStage) SetDisplayOrder(v int32) {
	o.DisplayOrder = v
}

// GetWritePermissions returns the WritePermissions field value if set, zero value otherwise.
func (o *PipelineStage) GetWritePermissions() string {
	if o == nil || o.WritePermissions == nil {
		var ret string
		return ret
	}
	return *o.WritePermissions
}

// GetWritePermissionsOk returns a tuple with the WritePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineStage) GetWritePermissionsOk() (*string, bool) {
	if o == nil || o.WritePermissions == nil {
		return nil, false
	}
	return o.WritePermissions, true
}

// HasWritePermissions returns a boolean if a field has been set.
func (o *PipelineStage) HasWritePermissions() bool {
	if o != nil && o.WritePermissions != nil {
		return true
	}

	return false
}

// SetWritePermissions gets a reference to the given string and assigns it to the WritePermissions field.
func (o *PipelineStage) SetWritePermissions(v string) {
	o.WritePermissions = &v
}

// GetLabel returns the Label field value
func (o *PipelineStage) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *PipelineStage) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *PipelineStage) SetLabel(v string) {
	o.Label = v
}

// GetId returns the Id field value
func (o *PipelineStage) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PipelineStage) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PipelineStage) SetId(v string) {
	o.Id = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PipelineStage) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PipelineStage) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PipelineStage) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o PipelineStage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.ArchivedAt != nil {
		toSerialize["archivedAt"] = o.ArchivedAt
	}
	if true {
		toSerialize["archived"] = o.Archived
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["displayOrder"] = o.DisplayOrder
	}
	if o.WritePermissions != nil {
		toSerialize["writePermissions"] = o.WritePermissions
	}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineStage struct {
	value *PipelineStage
	isSet bool
}

func (v NullablePipelineStage) Get() *PipelineStage {
	return v.value
}

func (v *NullablePipelineStage) Set(val *PipelineStage) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineStage) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineStage(val *PipelineStage) *NullablePipelineStage {
	return &NullablePipelineStage{value: val, isSet: true}
}

func (v NullablePipelineStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
