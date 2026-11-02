# \BasicAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSchemas**](BasicAPI.md#DeleteSchemas) | **Delete** /crm-object-schemas/2026-09/schemas/{objectType} | Delete the schema of a specified custom object.
[**DeleteSchemasByObjectTypeByAssociationIdentifier**](BasicAPI.md#DeleteSchemasByObjectTypeByAssociationIdentifier) | **Delete** /crm-object-schemas/2026-09/schemas/{objectType}/associations/{associationIdentifier} | Delete an existing association for a specified object type.
[**GetSchemas**](BasicAPI.md#GetSchemas) | **Get** /crm-object-schemas/2026-09/schemas | Retrieve all custom object schemas.
[**GetSchemasByObjectType**](BasicAPI.md#GetSchemasByObjectType) | **Get** /crm-object-schemas/2026-09/schemas/{objectType} | Retrieve the schema of a specified custom object.
[**PatchSchemas**](BasicAPI.md#PatchSchemas) | **Patch** /crm-object-schemas/2026-09/schemas/{objectType} | Update the schema of a specified custom object.
[**PostSchemas**](BasicAPI.md#PostSchemas) | **Post** /crm-object-schemas/2026-09/schemas | Create a new custom object schema.



## DeleteSchemas

> DeleteSchemas(ctx, objectType).Archived(archived).Execute()

Delete the schema of a specified custom object.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	objectType := "objectType_example" // string | The unique identifier of the object schema to delete.
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicAPI.DeleteSchemas(context.Background(), objectType).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.DeleteSchemas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The unique identifier of the object schema to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSchemasByObjectTypeByAssociationIdentifier

> DeleteSchemasByObjectTypeByAssociationIdentifier(ctx, associationIdentifier, objectType).Execute()

Delete an existing association for a specified object type.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	associationIdentifier := "associationIdentifier_example" // string | The unique identifier of the association definition to delete.
	objectType := "objectType_example" // string | The type of the custom object from which the association is being deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicAPI.DeleteSchemasByObjectTypeByAssociationIdentifier(context.Background(), associationIdentifier, objectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.DeleteSchemasByObjectTypeByAssociationIdentifier``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**associationIdentifier** | **string** | The unique identifier of the association definition to delete. | 
**objectType** | **string** | The type of the custom object from which the association is being deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSchemasByObjectTypeByAssociationIdentifierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemas

> CollectionResponseObjectSchemaNoPaging GetSchemas(ctx).Archived(archived).IncludeAssociationDefinitions(includeAssociationDefinitions).IncludeAuditMetadata(includeAuditMetadata).IncludePropertyDefinitions(includePropertyDefinitions).Execute()

Retrieve all custom object schemas.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)
	includeAssociationDefinitions := true // bool | Whether to include association definitions in the response. Defaults to true. (optional) (default to true)
	includeAuditMetadata := true // bool | Whether to include audit metadata in the response. Defaults to true. (optional) (default to true)
	includePropertyDefinitions := true // bool | Whether to include property definitions in the response. Defaults to true. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetSchemas(context.Background()).Archived(archived).IncludeAssociationDefinitions(includeAssociationDefinitions).IncludeAuditMetadata(includeAuditMetadata).IncludePropertyDefinitions(includePropertyDefinitions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetSchemas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSchemas`: CollectionResponseObjectSchemaNoPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]
 **includeAssociationDefinitions** | **bool** | Whether to include association definitions in the response. Defaults to true. | [default to true]
 **includeAuditMetadata** | **bool** | Whether to include audit metadata in the response. Defaults to true. | [default to true]
 **includePropertyDefinitions** | **bool** | Whether to include property definitions in the response. Defaults to true. | [default to true]

### Return type

[**CollectionResponseObjectSchemaNoPaging**](CollectionResponseObjectSchemaNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemasByObjectType

> ObjectSchema GetSchemasByObjectType(ctx, objectType).IncludeAssociationDefinitions(includeAssociationDefinitions).IncludeAuditMetadata(includeAuditMetadata).IncludePropertyDefinitions(includePropertyDefinitions).Execute()

Retrieve the schema of a specified custom object.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	objectType := "objectType_example" // string | The unique identifier for the CRM object type whose schema is being retrieved.
	includeAssociationDefinitions := true // bool | A boolean indicating whether to include association definitions in the response. Defaults to true. (optional) (default to true)
	includeAuditMetadata := true // bool | A boolean indicating whether to include audit metadata in the response. Defaults to true. (optional) (default to true)
	includePropertyDefinitions := true // bool | A boolean indicating whether to include property definitions in the response. Defaults to true. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetSchemasByObjectType(context.Background(), objectType).IncludeAssociationDefinitions(includeAssociationDefinitions).IncludeAuditMetadata(includeAuditMetadata).IncludePropertyDefinitions(includePropertyDefinitions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetSchemasByObjectType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSchemasByObjectType`: ObjectSchema
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetSchemasByObjectType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The unique identifier for the CRM object type whose schema is being retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemasByObjectTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAssociationDefinitions** | **bool** | A boolean indicating whether to include association definitions in the response. Defaults to true. | [default to true]
 **includeAuditMetadata** | **bool** | A boolean indicating whether to include audit metadata in the response. Defaults to true. | [default to true]
 **includePropertyDefinitions** | **bool** | A boolean indicating whether to include property definitions in the response. Defaults to true. | [default to true]

### Return type

[**ObjectSchema**](ObjectSchema.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSchemas

> ObjectTypeDefinition PatchSchemas(ctx, objectType).ObjectTypeDefinitionPatch(objectTypeDefinitionPatch).Execute()

Update the schema of a specified custom object.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	objectType := "objectType_example" // string | The unique identifier of the CRM object type whose schema is to be updated.
	objectTypeDefinitionPatch := *openapiclient.NewObjectTypeDefinitionPatch(false) // ObjectTypeDefinitionPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PatchSchemas(context.Background(), objectType).ObjectTypeDefinitionPatch(objectTypeDefinitionPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PatchSchemas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSchemas`: ObjectTypeDefinition
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PatchSchemas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The unique identifier of the CRM object type whose schema is to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectTypeDefinitionPatch** | [**ObjectTypeDefinitionPatch**](ObjectTypeDefinitionPatch.md) |  | 

### Return type

[**ObjectTypeDefinition**](ObjectTypeDefinition.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSchemas

> ObjectSchema PostSchemas(ctx).ObjectSchemaEgg(objectSchemaEgg).Execute()

Create a new custom object schema.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	objectSchemaEgg := *openapiclient.NewObjectSchemaEgg(false, []string{"AssociatedObjects_example"}, *openapiclient.NewObjectTypeDefinitionLabels(), "Name_example", []openapiclient.ObjectTypePropertyCreate{*openapiclient.NewObjectTypePropertyCreate("FieldType_example", "Label_example", "Name_example", "Type_example")}, []string{"RequiredProperties_example"}, []string{"SearchableProperties_example"}, []string{"SecondaryDisplayProperties_example"}, false) // ObjectSchemaEgg | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PostSchemas(context.Background()).ObjectSchemaEgg(objectSchemaEgg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostSchemas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSchemas`: ObjectSchema
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PostSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectSchemaEgg** | [**ObjectSchemaEgg**](ObjectSchemaEgg.md) |  | 

### Return type

[**ObjectSchema**](ObjectSchema.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

