# ExtensionActionDefinitionPatch

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionUrl** | **string** | The URL that will accept an HTTPS request each time workflows executes the custom action. | [optional] [default to null]
**Published** | **bool** | Whether this custom action is published to customers. | [optional] [default to null]
**InputFields** | [**[]InputFieldDefinition**](InputFieldDefinition.md) | The list of input fields to display in this custom action. | [optional] [default to null]
**ObjectRequestOptions** | [***ObjectRequestOptions**](ObjectRequestOptions.md) |  | [optional] [default to null]
**InputFieldDependencies** | [**[]OneOfExtensionActionDefinitionPatchInputFieldDependenciesItems**](.md) | A list of dependencies between the input fields. These configure when the input fields should be visible. | [optional] [default to null]
**Labels** | [**map[string]ActionLabels**](ActionLabels.md) | The user-facing labels for the custom action. | [optional] [default to null]
**ObjectTypes** | **[]string** | The object types that this custom action supports. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

