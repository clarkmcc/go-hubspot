# PropertyCreate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The internal property name, which must be used when referencing the property via the API. | [default to null]
**Label** | **string** | A human-readable property label that will be shown in HubSpot. | [default to null]
**Type_** | **string** | The data type of the property. | [default to null]
**FieldType** | **string** | Controls how the property appears in HubSpot. | [default to null]
**GroupName** | **string** | The name of the property group the property belongs to. | [default to null]
**Description** | **string** | A description of the property that will be shown as help text in HubSpot. | [optional] [default to null]
**Options** | [**[]OptionInput**](OptionInput.md) | A list of valid options for the property. This field is required for enumerated properties. | [optional] [default to null]
**DisplayOrder** | **int32** | Properties are displayed in order starting with the lowest positive integer value. Values of -1 will cause the property to be displayed after any positive values. | [optional] [default to null]
**HasUniqueValue** | **bool** | Whether or not the property&#x27;s value must be unique. Once set, this can&#x27;t be changed. | [optional] [default to null]
**Hidden** | **bool** | If true, the property won&#x27;t be visible and can&#x27;t be used in HubSpot. | [optional] [default to null]
**FormField** | **bool** | Whether or not the property can be used in a HubSpot form. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

