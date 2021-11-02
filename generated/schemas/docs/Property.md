# Property

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | When the property was created | [optional] [default to null]
**ArchivedAt** | [**time.Time**](time.Time.md) | When the property was archived. | [optional] [default to null]
**Name** | **string** | The internal property name, which must be used when referencing the property via the API. | [default to null]
**Label** | **string** | A human-readable property label that will be shown in HubSpot. | [default to null]
**Type_** | **string** | The property data type. | [default to null]
**FieldType** | **string** | Controls how the property appears in HubSpot. | [default to null]
**Description** | **string** | A description of the property that will be shown as help text in HubSpot. | [default to null]
**GroupName** | **string** | The name of the property group the property belongs to. | [default to null]
**Options** | [**[]Option**](Option.md) | A list of valid options for the property. This field is required for enumerated properties, but will be empty for other property types. | [default to null]
**CreatedUserId** | **string** | The internal ID of the user who created the property in HubSpot. This field may not exist if the property was created outside of HubSpot. | [optional] [default to null]
**UpdatedUserId** | **string** | The internal user ID of the user who updated the property in HubSpot. This field may not exist if the property was updated outside of HubSpot. | [optional] [default to null]
**ReferencedObjectType** | **string** | If this property is related to other object(s), they&#x27;ll be listed here. | [optional] [default to null]
**DisplayOrder** | **int32** | The order that this property should be displayed in the HubSpot UI relative to other properties for this object type. Properties are displayed in order starting with the lowest positive integer value. A value of -1 will cause the property to be displayed **after** any positive values. | [optional] [default to null]
**Calculated** | **bool** | For default properties, true indicates that the property is calculated by a HubSpot process. It has no effect for custom properties. | [optional] [default to null]
**ExternalOptions** | **bool** | For default properties, true indicates that the options are stored externally to the property settings. | [optional] [default to null]
**Archived** | **bool** | Whether or not the property is archived. | [optional] [default to null]
**HasUniqueValue** | **bool** | Whether or not the property&#x27;s value must be unique. Once set, this can&#x27;t be changed. | [optional] [default to null]
**Hidden** | **bool** |  | [optional] [default to null]
**HubspotDefined** | **bool** | This will be true for default object properties built into HubSpot. | [optional] [default to null]
**ShowCurrencySymbol** | **bool** | Whether the property will display the currency symbol set in the account settings. | [optional] [default to null]
**ModificationMetadata** | [***PropertyModificationMetadata**](PropertyModificationMetadata.md) |  | [optional] [default to null]
**FormField** | **bool** | Whether or not the property can be used in a HubSpot form. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

