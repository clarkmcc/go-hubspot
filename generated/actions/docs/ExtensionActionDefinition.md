# ExtensionActionDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the custom action. | [default to null]
**RevisionId** | **string** |  | [default to null]
**Functions** | [**[]ActionFunctionIdentifier**](ActionFunctionIdentifier.md) | A list of functions associated with the custom workflow action. | [default to null]
**ActionUrl** | **string** | The URL that will accept an HTTPS request each time workflows executes the custom action. | [default to null]
**Published** | **bool** | Whether this custom action is published to customers. | [default to null]
**ArchivedAt** | **int64** | The date that this custom action was archived, if the custom action is archived. | [optional] [default to null]
**InputFields** | [**[]InputFieldDefinition**](InputFieldDefinition.md) | The list of input fields to display in this custom action. | [default to null]
**ObjectRequestOptions** | [***ObjectRequestOptions**](ObjectRequestOptions.md) |  | [optional] [default to null]
**InputFieldDependencies** | [**[]OneOfExtensionActionDefinitionInputFieldDependenciesItems**](.md) | A list of dependencies between the input fields. These configure when the input fields should be visible. | [optional] [default to null]
**Labels** | [**map[string]ActionLabels**](ActionLabels.md) | The user-facing labels for the custom action. | [default to null]
**ObjectTypes** | **[]string** | The object types that this custom action supports. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

