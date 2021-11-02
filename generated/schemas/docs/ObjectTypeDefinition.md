# ObjectTypeDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | [***ObjectTypeDefinitionLabels**](ObjectTypeDefinitionLabels.md) |  | [default to null]
**RequiredProperties** | **[]string** | The names of properties that should be **required** when creating an object of this type. | [default to null]
**SearchableProperties** | **[]string** | Names of properties that will be indexed for this object type in by HubSpot&#x27;s product search. | [default to null]
**PrimaryDisplayProperty** | **string** | The name of the primary property for this object. This will be displayed as primary on the HubSpot record page for this object type. | [optional] [default to null]
**SecondaryDisplayProperties** | **[]string** | The names of secondary properties for this object. These will be displayed as secondary on the HubSpot record page for this object type. | [default to null]
**Archived** | **bool** |  | [default to null]
**Id** | **string** | A unique ID for this object type. Will be defined as {meta-type}-{unique ID}. | [default to null]
**FullyQualifiedName** | **string** |  | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | When the object type was created. | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | When the object type was last updated. | [optional] [default to null]
**ObjectTypeId** | **string** |  | [default to null]
**Name** | **string** | A unique name for this object. For internal use only. | [default to null]
**PortalId** | **int32** | The ID of the account that this object type is specific to. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

