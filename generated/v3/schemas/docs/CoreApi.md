# \CoreApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3SchemasObjectTypeArchive**](CoreApi.md#DeleteCrmV3SchemasObjectTypeArchive) | **Delete** /crm/v3/schemas/{objectType} | Delete a schema
[**DeleteCrmV3SchemasObjectTypeAssociationsAssociationIdentifierArchiveAssociation**](CoreApi.md#DeleteCrmV3SchemasObjectTypeAssociationsAssociationIdentifierArchiveAssociation) | **Delete** /crm/v3/schemas/{objectType}/associations/{associationIdentifier} | Remove an association
[**GetCrmV3SchemasGetAll**](CoreApi.md#GetCrmV3SchemasGetAll) | **Get** /crm/v3/schemas | Get all schemas
[**GetCrmV3SchemasObjectTypeGetById**](CoreApi.md#GetCrmV3SchemasObjectTypeGetById) | **Get** /crm/v3/schemas/{objectType} | Get an existing schema
[**PatchCrmV3SchemasObjectTypeUpdate**](CoreApi.md#PatchCrmV3SchemasObjectTypeUpdate) | **Patch** /crm/v3/schemas/{objectType} | Update a schema
[**PostCrmV3SchemasCreate**](CoreApi.md#PostCrmV3SchemasCreate) | **Post** /crm/v3/schemas | Create a new schema
[**PostCrmV3SchemasObjectTypeAssociationsCreateAssociation**](CoreApi.md#PostCrmV3SchemasObjectTypeAssociationsCreateAssociation) | **Post** /crm/v3/schemas/{objectType}/associations | Create an association



## DeleteCrmV3SchemasObjectTypeArchive

> DeleteCrmV3SchemasObjectTypeArchive(ctx, objectType).Archived(archived).Execute()

Delete a schema



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    objectType := "objectType_example" // string | Fully qualified name or object type ID of your schema.
    archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreApi.DeleteCrmV3SchemasObjectTypeArchive(context.Background(), objectType).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.DeleteCrmV3SchemasObjectTypeArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | Fully qualified name or object type ID of your schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3SchemasObjectTypeArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

 (empty response body)

### Authorization

[private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCrmV3SchemasObjectTypeAssociationsAssociationIdentifierArchiveAssociation

> DeleteCrmV3SchemasObjectTypeAssociationsAssociationIdentifierArchiveAssociation(ctx, objectType, associationIdentifier).Execute()

Remove an association



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    objectType := "objectType_example" // string | Fully qualified name or object type ID of your schema.
    associationIdentifier := "associationIdentifier_example" // string | Unique ID of the association to remove.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreApi.DeleteCrmV3SchemasObjectTypeAssociationsAssociationIdentifierArchiveAssociation(context.Background(), objectType, associationIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.DeleteCrmV3SchemasObjectTypeAssociationsAssociationIdentifierArchiveAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | Fully qualified name or object type ID of your schema. | 
**associationIdentifier** | **string** | Unique ID of the association to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3SchemasObjectTypeAssociationsAssociationIdentifierArchiveAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3SchemasGetAll

> CollectionResponseObjectSchemaNoPaging GetCrmV3SchemasGetAll(ctx).Archived(archived).Execute()

Get all schemas



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreApi.GetCrmV3SchemasGetAll(context.Background()).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.GetCrmV3SchemasGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3SchemasGetAll`: CollectionResponseObjectSchemaNoPaging
    fmt.Fprintf(os.Stdout, "Response from `CoreApi.GetCrmV3SchemasGetAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3SchemasGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

[**CollectionResponseObjectSchemaNoPaging**](CollectionResponseObjectSchemaNoPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3SchemasObjectTypeGetById

> ObjectSchema GetCrmV3SchemasObjectTypeGetById(ctx, objectType).Execute()

Get an existing schema



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    objectType := "objectType_example" // string | Fully qualified name or object type ID of your schema.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreApi.GetCrmV3SchemasObjectTypeGetById(context.Background(), objectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.GetCrmV3SchemasObjectTypeGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3SchemasObjectTypeGetById`: ObjectSchema
    fmt.Fprintf(os.Stdout, "Response from `CoreApi.GetCrmV3SchemasObjectTypeGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | Fully qualified name or object type ID of your schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3SchemasObjectTypeGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectSchema**](ObjectSchema.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCrmV3SchemasObjectTypeUpdate

> ObjectTypeDefinition PatchCrmV3SchemasObjectTypeUpdate(ctx, objectType).ObjectTypeDefinitionPatch(objectTypeDefinitionPatch).Execute()

Update a schema



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    objectType := "objectType_example" // string | Fully qualified name or object type ID of your schema.
    objectTypeDefinitionPatch := *openapiclient.NewObjectTypeDefinitionPatch() // ObjectTypeDefinitionPatch | Attributes to update in your schema.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreApi.PatchCrmV3SchemasObjectTypeUpdate(context.Background(), objectType).ObjectTypeDefinitionPatch(objectTypeDefinitionPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.PatchCrmV3SchemasObjectTypeUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCrmV3SchemasObjectTypeUpdate`: ObjectTypeDefinition
    fmt.Fprintf(os.Stdout, "Response from `CoreApi.PatchCrmV3SchemasObjectTypeUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | Fully qualified name or object type ID of your schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCrmV3SchemasObjectTypeUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectTypeDefinitionPatch** | [**ObjectTypeDefinitionPatch**](ObjectTypeDefinitionPatch.md) | Attributes to update in your schema. | 

### Return type

[**ObjectTypeDefinition**](ObjectTypeDefinition.md)

### Authorization

[private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3SchemasCreate

> ObjectSchema PostCrmV3SchemasCreate(ctx).ObjectSchemaEgg(objectSchemaEgg).Execute()

Create a new schema



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    objectSchemaEgg := *openapiclient.NewObjectSchemaEgg([]string{"RequiredProperties_example"}, "my_object", []string{"AssociatedObjects_example"}, []openapiclient.ObjectTypePropertyCreate{*openapiclient.NewObjectTypePropertyCreate("My object property", "enumeration", "Name_example", "select")}, *openapiclient.NewObjectTypeDefinitionLabels()) // ObjectSchemaEgg | Object schema definition, including properties and associations.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreApi.PostCrmV3SchemasCreate(context.Background()).ObjectSchemaEgg(objectSchemaEgg).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.PostCrmV3SchemasCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3SchemasCreate`: ObjectSchema
    fmt.Fprintf(os.Stdout, "Response from `CoreApi.PostCrmV3SchemasCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3SchemasCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectSchemaEgg** | [**ObjectSchemaEgg**](ObjectSchemaEgg.md) | Object schema definition, including properties and associations. | 

### Return type

[**ObjectSchema**](ObjectSchema.md)

### Authorization

[private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3SchemasObjectTypeAssociationsCreateAssociation

> AssociationDefinition PostCrmV3SchemasObjectTypeAssociationsCreateAssociation(ctx, objectType).AssociationDefinitionEgg(associationDefinitionEgg).Execute()

Create an association



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    objectType := "objectType_example" // string | Fully qualified name or object type ID of your schema.
    associationDefinitionEgg := *openapiclient.NewAssociationDefinitionEgg("2-123456", "contact") // AssociationDefinitionEgg | Attributes that define the association.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreApi.PostCrmV3SchemasObjectTypeAssociationsCreateAssociation(context.Background(), objectType).AssociationDefinitionEgg(associationDefinitionEgg).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.PostCrmV3SchemasObjectTypeAssociationsCreateAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3SchemasObjectTypeAssociationsCreateAssociation`: AssociationDefinition
    fmt.Fprintf(os.Stdout, "Response from `CoreApi.PostCrmV3SchemasObjectTypeAssociationsCreateAssociation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | Fully qualified name or object type ID of your schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3SchemasObjectTypeAssociationsCreateAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **associationDefinitionEgg** | [**AssociationDefinitionEgg**](AssociationDefinitionEgg.md) | Attributes that define the association. | 

### Return type

[**AssociationDefinition**](AssociationDefinition.md)

### Authorization

[private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

