# Property

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | When the property was created | [optional] 
**ArchivedAt** | Pointer to **time.Time** | When the property was archived. | [optional] 
**Name** | **string** | The internal property name, which must be used when referencing the property via the API. | 
**Label** | **string** | A human-readable property label that will be shown in HubSpot. | 
**Type** | **string** | The property data type. | 
**FieldType** | **string** | Controls how the property appears in HubSpot. | 
**Description** | **string** | A description of the property that will be shown as help text in HubSpot. | 
**GroupName** | **string** | The name of the property group the property belongs to. | 
**Options** | [**[]Option**](Option.md) | A list of valid options for the property. This field is required for enumerated properties, but will be empty for other property types. | 
**CreatedUserId** | Pointer to **string** | The internal ID of the user who created the property in HubSpot. This field may not exist if the property was created outside of HubSpot. | [optional] 
**UpdatedUserId** | Pointer to **string** | The internal user ID of the user who updated the property in HubSpot. This field may not exist if the property was updated outside of HubSpot. | [optional] 
**ReferencedObjectType** | Pointer to **string** | If this property is related to other object(s), they&#39;ll be listed here. | [optional] 
**DisplayOrder** | Pointer to **int32** | The order that this property should be displayed in the HubSpot UI relative to other properties for this object type. Properties are displayed in order starting with the lowest positive integer value. A value of -1 will cause the property to be displayed **after** any positive values. | [optional] 
**Calculated** | Pointer to **bool** | For default properties, true indicates that the property is calculated by a HubSpot process. It has no effect for custom properties. | [optional] 
**ExternalOptions** | Pointer to **bool** | For default properties, true indicates that the options are stored externally to the property settings. | [optional] 
**Archived** | Pointer to **bool** | Whether or not the property is archived. | [optional] 
**HasUniqueValue** | Pointer to **bool** | Whether or not the property&#39;s value must be unique. Once set, this can&#39;t be changed. | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**HubspotDefined** | Pointer to **bool** | This will be true for default object properties built into HubSpot. | [optional] 
**ShowCurrencySymbol** | Pointer to **bool** | Whether the property will display the currency symbol set in the account settings. | [optional] 
**ModificationMetadata** | Pointer to [**PropertyModificationMetadata**](PropertyModificationMetadata.md) |  | [optional] 
**FormField** | Pointer to **bool** | Whether or not the property can be used in a HubSpot form. | [optional] 

## Methods

### NewProperty

`func NewProperty(name string, label string, type_ string, fieldType string, description string, groupName string, options []Option, ) *Property`

NewProperty instantiates a new Property object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyWithDefaults

`func NewPropertyWithDefaults() *Property`

NewPropertyWithDefaults instantiates a new Property object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAt

`func (o *Property) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Property) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Property) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Property) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Property) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Property) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Property) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Property) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetArchivedAt

`func (o *Property) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *Property) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *Property) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *Property) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetName

`func (o *Property) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Property) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Property) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *Property) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Property) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Property) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *Property) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Property) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Property) SetType(v string)`

SetType sets Type field to given value.


### GetFieldType

`func (o *Property) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *Property) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *Property) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.


### GetDescription

`func (o *Property) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Property) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Property) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetGroupName

`func (o *Property) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *Property) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *Property) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetOptions

`func (o *Property) GetOptions() []Option`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Property) GetOptionsOk() (*[]Option, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Property) SetOptions(v []Option)`

SetOptions sets Options field to given value.


### GetCreatedUserId

`func (o *Property) GetCreatedUserId() string`

GetCreatedUserId returns the CreatedUserId field if non-nil, zero value otherwise.

### GetCreatedUserIdOk

`func (o *Property) GetCreatedUserIdOk() (*string, bool)`

GetCreatedUserIdOk returns a tuple with the CreatedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedUserId

`func (o *Property) SetCreatedUserId(v string)`

SetCreatedUserId sets CreatedUserId field to given value.

### HasCreatedUserId

`func (o *Property) HasCreatedUserId() bool`

HasCreatedUserId returns a boolean if a field has been set.

### GetUpdatedUserId

`func (o *Property) GetUpdatedUserId() string`

GetUpdatedUserId returns the UpdatedUserId field if non-nil, zero value otherwise.

### GetUpdatedUserIdOk

`func (o *Property) GetUpdatedUserIdOk() (*string, bool)`

GetUpdatedUserIdOk returns a tuple with the UpdatedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedUserId

`func (o *Property) SetUpdatedUserId(v string)`

SetUpdatedUserId sets UpdatedUserId field to given value.

### HasUpdatedUserId

`func (o *Property) HasUpdatedUserId() bool`

HasUpdatedUserId returns a boolean if a field has been set.

### GetReferencedObjectType

`func (o *Property) GetReferencedObjectType() string`

GetReferencedObjectType returns the ReferencedObjectType field if non-nil, zero value otherwise.

### GetReferencedObjectTypeOk

`func (o *Property) GetReferencedObjectTypeOk() (*string, bool)`

GetReferencedObjectTypeOk returns a tuple with the ReferencedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedObjectType

`func (o *Property) SetReferencedObjectType(v string)`

SetReferencedObjectType sets ReferencedObjectType field to given value.

### HasReferencedObjectType

`func (o *Property) HasReferencedObjectType() bool`

HasReferencedObjectType returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *Property) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *Property) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *Property) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *Property) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetCalculated

`func (o *Property) GetCalculated() bool`

GetCalculated returns the Calculated field if non-nil, zero value otherwise.

### GetCalculatedOk

`func (o *Property) GetCalculatedOk() (*bool, bool)`

GetCalculatedOk returns a tuple with the Calculated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculated

`func (o *Property) SetCalculated(v bool)`

SetCalculated sets Calculated field to given value.

### HasCalculated

`func (o *Property) HasCalculated() bool`

HasCalculated returns a boolean if a field has been set.

### GetExternalOptions

`func (o *Property) GetExternalOptions() bool`

GetExternalOptions returns the ExternalOptions field if non-nil, zero value otherwise.

### GetExternalOptionsOk

`func (o *Property) GetExternalOptionsOk() (*bool, bool)`

GetExternalOptionsOk returns a tuple with the ExternalOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOptions

`func (o *Property) SetExternalOptions(v bool)`

SetExternalOptions sets ExternalOptions field to given value.

### HasExternalOptions

`func (o *Property) HasExternalOptions() bool`

HasExternalOptions returns a boolean if a field has been set.

### GetArchived

`func (o *Property) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Property) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Property) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *Property) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetHasUniqueValue

`func (o *Property) GetHasUniqueValue() bool`

GetHasUniqueValue returns the HasUniqueValue field if non-nil, zero value otherwise.

### GetHasUniqueValueOk

`func (o *Property) GetHasUniqueValueOk() (*bool, bool)`

GetHasUniqueValueOk returns a tuple with the HasUniqueValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUniqueValue

`func (o *Property) SetHasUniqueValue(v bool)`

SetHasUniqueValue sets HasUniqueValue field to given value.

### HasHasUniqueValue

`func (o *Property) HasHasUniqueValue() bool`

HasHasUniqueValue returns a boolean if a field has been set.

### GetHidden

`func (o *Property) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *Property) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *Property) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *Property) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetHubspotDefined

`func (o *Property) GetHubspotDefined() bool`

GetHubspotDefined returns the HubspotDefined field if non-nil, zero value otherwise.

### GetHubspotDefinedOk

`func (o *Property) GetHubspotDefinedOk() (*bool, bool)`

GetHubspotDefinedOk returns a tuple with the HubspotDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubspotDefined

`func (o *Property) SetHubspotDefined(v bool)`

SetHubspotDefined sets HubspotDefined field to given value.

### HasHubspotDefined

`func (o *Property) HasHubspotDefined() bool`

HasHubspotDefined returns a boolean if a field has been set.

### GetShowCurrencySymbol

`func (o *Property) GetShowCurrencySymbol() bool`

GetShowCurrencySymbol returns the ShowCurrencySymbol field if non-nil, zero value otherwise.

### GetShowCurrencySymbolOk

`func (o *Property) GetShowCurrencySymbolOk() (*bool, bool)`

GetShowCurrencySymbolOk returns a tuple with the ShowCurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCurrencySymbol

`func (o *Property) SetShowCurrencySymbol(v bool)`

SetShowCurrencySymbol sets ShowCurrencySymbol field to given value.

### HasShowCurrencySymbol

`func (o *Property) HasShowCurrencySymbol() bool`

HasShowCurrencySymbol returns a boolean if a field has been set.

### GetModificationMetadata

`func (o *Property) GetModificationMetadata() PropertyModificationMetadata`

GetModificationMetadata returns the ModificationMetadata field if non-nil, zero value otherwise.

### GetModificationMetadataOk

`func (o *Property) GetModificationMetadataOk() (*PropertyModificationMetadata, bool)`

GetModificationMetadataOk returns a tuple with the ModificationMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationMetadata

`func (o *Property) SetModificationMetadata(v PropertyModificationMetadata)`

SetModificationMetadata sets ModificationMetadata field to given value.

### HasModificationMetadata

`func (o *Property) HasModificationMetadata() bool`

HasModificationMetadata returns a boolean if a field has been set.

### GetFormField

`func (o *Property) GetFormField() bool`

GetFormField returns the FormField field if non-nil, zero value otherwise.

### GetFormFieldOk

`func (o *Property) GetFormFieldOk() (*bool, bool)`

GetFormFieldOk returns a tuple with the FormField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormField

`func (o *Property) SetFormField(v bool)`

SetFormField sets FormField field to given value.

### HasFormField

`func (o *Property) HasFormField() bool`

HasFormField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


