/*
Properties

All HubSpot objects store data in default and custom properties. These endpoints provide access to read and modify object properties in HubSpot.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package properties

import (
	"encoding/json"
	"time"
)

// Property struct for Property
type Property struct {
	// The internal user ID of the user who created the property in HubSpot. This field may not exist if the property was created outside of HubSpot.
	CreatedUserId *string `json:"createdUserId,omitempty"`
	// Whether or not the property will be hidden from the HubSpot UI. It's recommended this be set to false for custom properties.
	Hidden               *bool                         `json:"hidden,omitempty"`
	ModificationMetadata *PropertyModificationMetadata `json:"modificationMetadata,omitempty"`
	// Properties are shown in order, starting with the lowest positive integer value.
	DisplayOrder *int32 `json:"displayOrder,omitempty"`
	// A description of the property that will be shown as help text in HubSpot.
	Description string `json:"description"`
	// Whether or not the property will display the currency symbol set in the account settings.
	ShowCurrencySymbol *bool `json:"showCurrencySymbol,omitempty"`
	// A human-readable property label that will be shown in HubSpot.
	Label string `json:"label"`
	// The property data type.
	Type string `json:"type"`
	// This will be true for default object properties built into HubSpot.
	HubspotDefined *bool `json:"hubspotDefined,omitempty"`
	// Whether or not the property can be used in a HubSpot form.
	FormField *bool `json:"formField,omitempty"`
	//
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// When the property was archived.
	ArchivedAt *time.Time `json:"archivedAt,omitempty"`
	// Whether or not the property is archived.
	Archived *bool `json:"archived,omitempty"`
	// The name of the property group the property belongs to.
	GroupName string `json:"groupName"`
	// If this property is related to other object(s), they'll be listed here.
	ReferencedObjectType *string `json:"referencedObjectType,omitempty"`
	// The internal property name, which must be used when referencing the property via the API.
	Name string `json:"name"`
	// A list of valid options for the property. This field is required for enumerated properties, but will be empty for other property types.
	Options []Option `json:"options"`
	// Represents a formula that is used to compute a calculated property.
	CalculationFormula *string `json:"calculationFormula,omitempty"`
	// Whether or not the property's value must be unique. Once set, this can't be changed.
	HasUniqueValue *bool `json:"hasUniqueValue,omitempty"`
	// Controls how the property appears in HubSpot.
	FieldType string `json:"fieldType"`
	// The internal user ID of the user who updated the property in HubSpot. This field may not exist if the property was updated outside of HubSpot.
	UpdatedUserId *string `json:"updatedUserId,omitempty"`
	// For default properties, true indicates that the property is calculated by a HubSpot process. It has no effect for custom properties.
	Calculated *bool `json:"calculated,omitempty"`
	// For default properties, true indicates that the options are stored externally to the property settings.
	ExternalOptions *bool `json:"externalOptions,omitempty"`
	//
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewProperty instantiates a new Property object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProperty(description string, label string, type_ string, groupName string, name string, options []Option, fieldType string) *Property {
	this := Property{}
	this.Description = description
	this.Label = label
	this.Type = type_
	this.GroupName = groupName
	this.Name = name
	this.Options = options
	this.FieldType = fieldType
	return &this
}

// NewPropertyWithDefaults instantiates a new Property object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyWithDefaults() *Property {
	this := Property{}
	return &this
}

// GetCreatedUserId returns the CreatedUserId field value if set, zero value otherwise.
func (o *Property) GetCreatedUserId() string {
	if o == nil || o.CreatedUserId == nil {
		var ret string
		return ret
	}
	return *o.CreatedUserId
}

// GetCreatedUserIdOk returns a tuple with the CreatedUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetCreatedUserIdOk() (*string, bool) {
	if o == nil || o.CreatedUserId == nil {
		return nil, false
	}
	return o.CreatedUserId, true
}

// HasCreatedUserId returns a boolean if a field has been set.
func (o *Property) HasCreatedUserId() bool {
	if o != nil && o.CreatedUserId != nil {
		return true
	}

	return false
}

// SetCreatedUserId gets a reference to the given string and assigns it to the CreatedUserId field.
func (o *Property) SetCreatedUserId(v string) {
	o.CreatedUserId = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *Property) GetHidden() bool {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetHiddenOk() (*bool, bool) {
	if o == nil || o.Hidden == nil {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *Property) HasHidden() bool {
	if o != nil && o.Hidden != nil {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *Property) SetHidden(v bool) {
	o.Hidden = &v
}

// GetModificationMetadata returns the ModificationMetadata field value if set, zero value otherwise.
func (o *Property) GetModificationMetadata() PropertyModificationMetadata {
	if o == nil || o.ModificationMetadata == nil {
		var ret PropertyModificationMetadata
		return ret
	}
	return *o.ModificationMetadata
}

// GetModificationMetadataOk returns a tuple with the ModificationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetModificationMetadataOk() (*PropertyModificationMetadata, bool) {
	if o == nil || o.ModificationMetadata == nil {
		return nil, false
	}
	return o.ModificationMetadata, true
}

// HasModificationMetadata returns a boolean if a field has been set.
func (o *Property) HasModificationMetadata() bool {
	if o != nil && o.ModificationMetadata != nil {
		return true
	}

	return false
}

// SetModificationMetadata gets a reference to the given PropertyModificationMetadata and assigns it to the ModificationMetadata field.
func (o *Property) SetModificationMetadata(v PropertyModificationMetadata) {
	o.ModificationMetadata = &v
}

// GetDisplayOrder returns the DisplayOrder field value if set, zero value otherwise.
func (o *Property) GetDisplayOrder() int32 {
	if o == nil || o.DisplayOrder == nil {
		var ret int32
		return ret
	}
	return *o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetDisplayOrderOk() (*int32, bool) {
	if o == nil || o.DisplayOrder == nil {
		return nil, false
	}
	return o.DisplayOrder, true
}

// HasDisplayOrder returns a boolean if a field has been set.
func (o *Property) HasDisplayOrder() bool {
	if o != nil && o.DisplayOrder != nil {
		return true
	}

	return false
}

// SetDisplayOrder gets a reference to the given int32 and assigns it to the DisplayOrder field.
func (o *Property) SetDisplayOrder(v int32) {
	o.DisplayOrder = &v
}

// GetDescription returns the Description field value
func (o *Property) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Property) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Property) SetDescription(v string) {
	o.Description = v
}

// GetShowCurrencySymbol returns the ShowCurrencySymbol field value if set, zero value otherwise.
func (o *Property) GetShowCurrencySymbol() bool {
	if o == nil || o.ShowCurrencySymbol == nil {
		var ret bool
		return ret
	}
	return *o.ShowCurrencySymbol
}

// GetShowCurrencySymbolOk returns a tuple with the ShowCurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetShowCurrencySymbolOk() (*bool, bool) {
	if o == nil || o.ShowCurrencySymbol == nil {
		return nil, false
	}
	return o.ShowCurrencySymbol, true
}

// HasShowCurrencySymbol returns a boolean if a field has been set.
func (o *Property) HasShowCurrencySymbol() bool {
	if o != nil && o.ShowCurrencySymbol != nil {
		return true
	}

	return false
}

// SetShowCurrencySymbol gets a reference to the given bool and assigns it to the ShowCurrencySymbol field.
func (o *Property) SetShowCurrencySymbol(v bool) {
	o.ShowCurrencySymbol = &v
}

// GetLabel returns the Label field value
func (o *Property) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *Property) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *Property) SetLabel(v string) {
	o.Label = v
}

// GetType returns the Type field value
func (o *Property) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Property) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Property) SetType(v string) {
	o.Type = v
}

// GetHubspotDefined returns the HubspotDefined field value if set, zero value otherwise.
func (o *Property) GetHubspotDefined() bool {
	if o == nil || o.HubspotDefined == nil {
		var ret bool
		return ret
	}
	return *o.HubspotDefined
}

// GetHubspotDefinedOk returns a tuple with the HubspotDefined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetHubspotDefinedOk() (*bool, bool) {
	if o == nil || o.HubspotDefined == nil {
		return nil, false
	}
	return o.HubspotDefined, true
}

// HasHubspotDefined returns a boolean if a field has been set.
func (o *Property) HasHubspotDefined() bool {
	if o != nil && o.HubspotDefined != nil {
		return true
	}

	return false
}

// SetHubspotDefined gets a reference to the given bool and assigns it to the HubspotDefined field.
func (o *Property) SetHubspotDefined(v bool) {
	o.HubspotDefined = &v
}

// GetFormField returns the FormField field value if set, zero value otherwise.
func (o *Property) GetFormField() bool {
	if o == nil || o.FormField == nil {
		var ret bool
		return ret
	}
	return *o.FormField
}

// GetFormFieldOk returns a tuple with the FormField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetFormFieldOk() (*bool, bool) {
	if o == nil || o.FormField == nil {
		return nil, false
	}
	return o.FormField, true
}

// HasFormField returns a boolean if a field has been set.
func (o *Property) HasFormField() bool {
	if o != nil && o.FormField != nil {
		return true
	}

	return false
}

// SetFormField gets a reference to the given bool and assigns it to the FormField field.
func (o *Property) SetFormField(v bool) {
	o.FormField = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Property) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Property) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Property) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *Property) GetArchivedAt() time.Time {
	if o == nil || o.ArchivedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetArchivedAtOk() (*time.Time, bool) {
	if o == nil || o.ArchivedAt == nil {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *Property) HasArchivedAt() bool {
	if o != nil && o.ArchivedAt != nil {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given time.Time and assigns it to the ArchivedAt field.
func (o *Property) SetArchivedAt(v time.Time) {
	o.ArchivedAt = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *Property) GetArchived() bool {
	if o == nil || o.Archived == nil {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetArchivedOk() (*bool, bool) {
	if o == nil || o.Archived == nil {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *Property) HasArchived() bool {
	if o != nil && o.Archived != nil {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *Property) SetArchived(v bool) {
	o.Archived = &v
}

// GetGroupName returns the GroupName field value
func (o *Property) GetGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value
// and a boolean to check if the value has been set.
func (o *Property) GetGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupName, true
}

// SetGroupName sets field value
func (o *Property) SetGroupName(v string) {
	o.GroupName = v
}

// GetReferencedObjectType returns the ReferencedObjectType field value if set, zero value otherwise.
func (o *Property) GetReferencedObjectType() string {
	if o == nil || o.ReferencedObjectType == nil {
		var ret string
		return ret
	}
	return *o.ReferencedObjectType
}

// GetReferencedObjectTypeOk returns a tuple with the ReferencedObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetReferencedObjectTypeOk() (*string, bool) {
	if o == nil || o.ReferencedObjectType == nil {
		return nil, false
	}
	return o.ReferencedObjectType, true
}

// HasReferencedObjectType returns a boolean if a field has been set.
func (o *Property) HasReferencedObjectType() bool {
	if o != nil && o.ReferencedObjectType != nil {
		return true
	}

	return false
}

// SetReferencedObjectType gets a reference to the given string and assigns it to the ReferencedObjectType field.
func (o *Property) SetReferencedObjectType(v string) {
	o.ReferencedObjectType = &v
}

// GetName returns the Name field value
func (o *Property) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Property) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Property) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *Property) GetOptions() []Option {
	if o == nil {
		var ret []Option
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *Property) GetOptionsOk() ([]Option, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *Property) SetOptions(v []Option) {
	o.Options = v
}

// GetCalculationFormula returns the CalculationFormula field value if set, zero value otherwise.
func (o *Property) GetCalculationFormula() string {
	if o == nil || o.CalculationFormula == nil {
		var ret string
		return ret
	}
	return *o.CalculationFormula
}

// GetCalculationFormulaOk returns a tuple with the CalculationFormula field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetCalculationFormulaOk() (*string, bool) {
	if o == nil || o.CalculationFormula == nil {
		return nil, false
	}
	return o.CalculationFormula, true
}

// HasCalculationFormula returns a boolean if a field has been set.
func (o *Property) HasCalculationFormula() bool {
	if o != nil && o.CalculationFormula != nil {
		return true
	}

	return false
}

// SetCalculationFormula gets a reference to the given string and assigns it to the CalculationFormula field.
func (o *Property) SetCalculationFormula(v string) {
	o.CalculationFormula = &v
}

// GetHasUniqueValue returns the HasUniqueValue field value if set, zero value otherwise.
func (o *Property) GetHasUniqueValue() bool {
	if o == nil || o.HasUniqueValue == nil {
		var ret bool
		return ret
	}
	return *o.HasUniqueValue
}

// GetHasUniqueValueOk returns a tuple with the HasUniqueValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetHasUniqueValueOk() (*bool, bool) {
	if o == nil || o.HasUniqueValue == nil {
		return nil, false
	}
	return o.HasUniqueValue, true
}

// HasHasUniqueValue returns a boolean if a field has been set.
func (o *Property) HasHasUniqueValue() bool {
	if o != nil && o.HasUniqueValue != nil {
		return true
	}

	return false
}

// SetHasUniqueValue gets a reference to the given bool and assigns it to the HasUniqueValue field.
func (o *Property) SetHasUniqueValue(v bool) {
	o.HasUniqueValue = &v
}

// GetFieldType returns the FieldType field value
func (o *Property) GetFieldType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldType
}

// GetFieldTypeOk returns a tuple with the FieldType field value
// and a boolean to check if the value has been set.
func (o *Property) GetFieldTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldType, true
}

// SetFieldType sets field value
func (o *Property) SetFieldType(v string) {
	o.FieldType = v
}

// GetUpdatedUserId returns the UpdatedUserId field value if set, zero value otherwise.
func (o *Property) GetUpdatedUserId() string {
	if o == nil || o.UpdatedUserId == nil {
		var ret string
		return ret
	}
	return *o.UpdatedUserId
}

// GetUpdatedUserIdOk returns a tuple with the UpdatedUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetUpdatedUserIdOk() (*string, bool) {
	if o == nil || o.UpdatedUserId == nil {
		return nil, false
	}
	return o.UpdatedUserId, true
}

// HasUpdatedUserId returns a boolean if a field has been set.
func (o *Property) HasUpdatedUserId() bool {
	if o != nil && o.UpdatedUserId != nil {
		return true
	}

	return false
}

// SetUpdatedUserId gets a reference to the given string and assigns it to the UpdatedUserId field.
func (o *Property) SetUpdatedUserId(v string) {
	o.UpdatedUserId = &v
}

// GetCalculated returns the Calculated field value if set, zero value otherwise.
func (o *Property) GetCalculated() bool {
	if o == nil || o.Calculated == nil {
		var ret bool
		return ret
	}
	return *o.Calculated
}

// GetCalculatedOk returns a tuple with the Calculated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetCalculatedOk() (*bool, bool) {
	if o == nil || o.Calculated == nil {
		return nil, false
	}
	return o.Calculated, true
}

// HasCalculated returns a boolean if a field has been set.
func (o *Property) HasCalculated() bool {
	if o != nil && o.Calculated != nil {
		return true
	}

	return false
}

// SetCalculated gets a reference to the given bool and assigns it to the Calculated field.
func (o *Property) SetCalculated(v bool) {
	o.Calculated = &v
}

// GetExternalOptions returns the ExternalOptions field value if set, zero value otherwise.
func (o *Property) GetExternalOptions() bool {
	if o == nil || o.ExternalOptions == nil {
		var ret bool
		return ret
	}
	return *o.ExternalOptions
}

// GetExternalOptionsOk returns a tuple with the ExternalOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetExternalOptionsOk() (*bool, bool) {
	if o == nil || o.ExternalOptions == nil {
		return nil, false
	}
	return o.ExternalOptions, true
}

// HasExternalOptions returns a boolean if a field has been set.
func (o *Property) HasExternalOptions() bool {
	if o != nil && o.ExternalOptions != nil {
		return true
	}

	return false
}

// SetExternalOptions gets a reference to the given bool and assigns it to the ExternalOptions field.
func (o *Property) SetExternalOptions(v bool) {
	o.ExternalOptions = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Property) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Property) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Property) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Property) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedUserId != nil {
		toSerialize["createdUserId"] = o.CreatedUserId
	}
	if o.Hidden != nil {
		toSerialize["hidden"] = o.Hidden
	}
	if o.ModificationMetadata != nil {
		toSerialize["modificationMetadata"] = o.ModificationMetadata
	}
	if o.DisplayOrder != nil {
		toSerialize["displayOrder"] = o.DisplayOrder
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.ShowCurrencySymbol != nil {
		toSerialize["showCurrencySymbol"] = o.ShowCurrencySymbol
	}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.HubspotDefined != nil {
		toSerialize["hubspotDefined"] = o.HubspotDefined
	}
	if o.FormField != nil {
		toSerialize["formField"] = o.FormField
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.ArchivedAt != nil {
		toSerialize["archivedAt"] = o.ArchivedAt
	}
	if o.Archived != nil {
		toSerialize["archived"] = o.Archived
	}
	if true {
		toSerialize["groupName"] = o.GroupName
	}
	if o.ReferencedObjectType != nil {
		toSerialize["referencedObjectType"] = o.ReferencedObjectType
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["options"] = o.Options
	}
	if o.CalculationFormula != nil {
		toSerialize["calculationFormula"] = o.CalculationFormula
	}
	if o.HasUniqueValue != nil {
		toSerialize["hasUniqueValue"] = o.HasUniqueValue
	}
	if true {
		toSerialize["fieldType"] = o.FieldType
	}
	if o.UpdatedUserId != nil {
		toSerialize["updatedUserId"] = o.UpdatedUserId
	}
	if o.Calculated != nil {
		toSerialize["calculated"] = o.Calculated
	}
	if o.ExternalOptions != nil {
		toSerialize["externalOptions"] = o.ExternalOptions
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableProperty struct {
	value *Property
	isSet bool
}

func (v NullableProperty) Get() *Property {
	return v.value
}

func (v *NullableProperty) Set(val *Property) {
	v.value = val
	v.isSet = true
}

func (v NullableProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProperty(val *Property) *NullableProperty {
	return &NullableProperty{value: val, isSet: true}
}

func (v NullableProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
