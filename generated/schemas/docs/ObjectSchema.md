# ObjectSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | [***ObjectTypeDefinitionLabels**](ObjectTypeDefinitionLabels.md) |  | [default to null]
**RequiredProperties** | **[]string** | The names of properties that should be **required** when creating an object of this type. | [default to null]
**SearchableProperties** | **[]string** | Names of properties that will be indexed for this object type in by HubSpot&#x27;s product search. | [default to null]
**PrimaryDisplayProperty** | **string** | The name of the primary property for this object. This will be displayed as primary on the HubSpot record page for this object type. | [optional] [default to null]
**SecondaryDisplayProperties** | **[]string** | The names of secondary properties for this object. These will be displayed as secondary on the HubSpot record page for this object type. | [default to null]
**Archived** | **bool** |  | [default to null]
**Id** | **string** | A unique ID for this schema&#x27;s object type. Will be defined as {meta-type}-{unique ID}. | [default to null]
**FullyQualifiedName** | **string** | An assigned unique ID for the object, including portal ID and object name. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | When the object schema was created. | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | When the object schema was last updated. | [optional] [default to null]
**ObjectTypeId** | **string** |  | [default to null]
**Properties** | [**[]Property**](Property.md) | Properties defined for this object type. | [default to null]
**Associations** | [**[]AssociationDefinition**](AssociationDefinition.md) | Associations defined for a given object type. | [default to null]
**Name** | **string** | A unique name for the schema&#x27;s object type. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

