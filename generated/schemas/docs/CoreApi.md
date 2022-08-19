# \CoreApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmObjectSchemasV3SchemasObjectTypeArchive**](CoreApi.md#DeleteCrmObjectSchemasV3SchemasObjectTypeArchive) | **Delete** /crm/v3/schemas/{objectType} | Delete a schema
[**DeleteCrmObjectSchemasV3SchemasObjectTypeAssociationsAssociationIdentifierArchiveAssociation**](CoreApi.md#DeleteCrmObjectSchemasV3SchemasObjectTypeAssociationsAssociationIdentifierArchiveAssociation) | **Delete** /crm/v3/schemas/{objectType}/associations/{associationIdentifier} | Remove an association
[**GetCrmObjectSchemasV3SchemasGetAll**](CoreApi.md#GetCrmObjectSchemasV3SchemasGetAll) | **Get** /crm/v3/schemas | Get all schemas
[**GetCrmObjectSchemasV3SchemasObjectTypeGetById**](CoreApi.md#GetCrmObjectSchemasV3SchemasObjectTypeGetById) | **Get** /crm/v3/schemas/{objectType} | Get an existing schema
[**PatchCrmObjectSchemasV3SchemasObjectTypeUpdate**](CoreApi.md#PatchCrmObjectSchemasV3SchemasObjectTypeUpdate) | **Patch** /crm/v3/schemas/{objectType} | Update a schema
[**PostCrmObjectSchemasV3SchemasCreate**](CoreApi.md#PostCrmObjectSchemasV3SchemasCreate) | **Post** /crm/v3/schemas | Create a new schema
[**PostCrmObjectSchemasV3SchemasObjectTypeAssociationsCreateAssociation**](CoreApi.md#PostCrmObjectSchemasV3SchemasObjectTypeAssociationsCreateAssociation) | **Post** /crm/v3/schemas/{objectType}/associations | Create an association



## DeleteCrmObjectSchemasV3SchemasObjectTypeArchive

> DeleteCrmObjectSchemasV3SchemasObjectTypeArchive(ctx, objectType).Archived(archived).Execute()

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
    resp, r, err := apiClient.CoreApi.DeleteCrmObjectSchemasV3SchemasObjectTypeArchive(context.Background(), objectType).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.DeleteCrmObjectSchemasV3SchemasObjectTypeArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCrmObjectSchemasV3SchemasObjectTypeArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCrmObjectSchemasV3SchemasObjectTypeAssociationsAssociationIdentifierArchiveAssociation

> DeleteCrmObjectSchemasV3SchemasObjectTypeAssociationsAssociationIdentifierArchiveAssociation(ctx, objectType, associationIdentifier).Execute()

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
    resp, r, err := apiClient.CoreApi.DeleteCrmObjectSchemasV3SchemasObjectTypeAssociationsAssociationIdentifierArchiveAssociation(context.Background(), objectType, associationIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.DeleteCrmObjectSchemasV3SchemasObjectTypeAssociationsAssociationIdentifierArchiveAssociation``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCrmObjectSchemasV3SchemasObjectTypeAssociationsAssociationIdentifierArchiveAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmObjectSchemasV3SchemasGetAll

> CollectionResponseObjectSchemaNoPaging GetCrmObjectSchemasV3SchemasGetAll(ctx).Archived(archived).Execute()

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
    resp, r, err := apiClient.CoreApi.GetCrmObjectSchemasV3SchemasGetAll(context.Background()).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.GetCrmObjectSchemasV3SchemasGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmObjectSchemasV3SchemasGetAll`: CollectionResponseObjectSchemaNoPaging
    fmt.Fprintf(os.Stdout, "Response from `CoreApi.GetCrmObjectSchemasV3SchemasGetAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmObjectSchemasV3SchemasGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

[**CollectionResponseObjectSchemaNoPaging**](CollectionResponseObjectSchemaNoPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmObjectSchemasV3SchemasObjectTypeGetById

> ObjectSchema GetCrmObjectSchemasV3SchemasObjectTypeGetById(ctx, objectType).Execute()

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
    resp, r, err := apiClient.CoreApi.GetCrmObjectSchemasV3SchemasObjectTypeGetById(context.Background(), objectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.GetCrmObjectSchemasV3SchemasObjectTypeGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmObjectSchemasV3SchemasObjectTypeGetById`: ObjectSchema
    fmt.Fprintf(os.Stdout, "Response from `CoreApi.GetCrmObjectSchemasV3SchemasObjectTypeGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | Fully qualified name or object type ID of your schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmObjectSchemasV3SchemasObjectTypeGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectSchema**](ObjectSchema.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCrmObjectSchemasV3SchemasObjectTypeUpdate

> ObjectTypeDefinition PatchCrmObjectSchemasV3SchemasObjectTypeUpdate(ctx, objectType).ObjectTypeDefinitionPatch(objectTypeDefinitionPatch).Execute()

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
    resp, r, err := apiClient.CoreApi.PatchCrmObjectSchemasV3SchemasObjectTypeUpdate(context.Background(), objectType).ObjectTypeDefinitionPatch(objectTypeDefinitionPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.PatchCrmObjectSchemasV3SchemasObjectTypeUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCrmObjectSchemasV3SchemasObjectTypeUpdate`: ObjectTypeDefinition
    fmt.Fprintf(os.Stdout, "Response from `CoreApi.PatchCrmObjectSchemasV3SchemasObjectTypeUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | Fully qualified name or object type ID of your schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCrmObjectSchemasV3SchemasObjectTypeUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectTypeDefinitionPatch** | [**ObjectTypeDefinitionPatch**](ObjectTypeDefinitionPatch.md) | Attributes to update in your schema. | 

### Return type

[**ObjectTypeDefinition**](ObjectTypeDefinition.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmObjectSchemasV3SchemasCreate

> ObjectSchema PostCrmObjectSchemasV3SchemasCreate(ctx).ObjectSchemaEgg(objectSchemaEgg).Execute()

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
    objectSchemaEgg := *openapiclient.NewObjectSchemaEgg(*openapiclient.NewObjectTypeDefinitionLabels(), []string{"RequiredProperties_example"}, []string{"SearchableProperties_example"}, []string{"SecondaryDisplayProperties_example"}, []openapiclient.ObjectTypePropertyCreate{*openapiclient.NewObjectTypePropertyCreate("Name_example", "Label_example", "Type_example", "FieldType_example")}, []string{"AssociatedObjects_example"}, "Name_example") // ObjectSchemaEgg | Object schema definition, including properties and associations.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreApi.PostCrmObjectSchemasV3SchemasCreate(context.Background()).ObjectSchemaEgg(objectSchemaEgg).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.PostCrmObjectSchemasV3SchemasCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmObjectSchemasV3SchemasCreate`: ObjectSchema
    fmt.Fprintf(os.Stdout, "Response from `CoreApi.PostCrmObjectSchemasV3SchemasCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmObjectSchemasV3SchemasCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectSchemaEgg** | [**ObjectSchemaEgg**](ObjectSchemaEgg.md) | Object schema definition, including properties and associations. | 

### Return type

[**ObjectSchema**](ObjectSchema.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmObjectSchemasV3SchemasObjectTypeAssociationsCreateAssociation

> AssociationDefinition PostCrmObjectSchemasV3SchemasObjectTypeAssociationsCreateAssociation(ctx, objectType).AssociationDefinitionEgg(associationDefinitionEgg).Execute()

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
    associationDefinitionEgg := *openapiclient.NewAssociationDefinitionEgg("FromObjectTypeId_example", "ToObjectTypeId_example") // AssociationDefinitionEgg | Attributes that define the association.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreApi.PostCrmObjectSchemasV3SchemasObjectTypeAssociationsCreateAssociation(context.Background(), objectType).AssociationDefinitionEgg(associationDefinitionEgg).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.PostCrmObjectSchemasV3SchemasObjectTypeAssociationsCreateAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmObjectSchemasV3SchemasObjectTypeAssociationsCreateAssociation`: AssociationDefinition
    fmt.Fprintf(os.Stdout, "Response from `CoreApi.PostCrmObjectSchemasV3SchemasObjectTypeAssociationsCreateAssociation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | Fully qualified name or object type ID of your schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmObjectSchemasV3SchemasObjectTypeAssociationsCreateAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **associationDefinitionEgg** | [**AssociationDefinitionEgg**](AssociationDefinitionEgg.md) | Attributes that define the association. | 

### Return type

[**AssociationDefinition**](AssociationDefinition.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

