# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletecrmObjectSchemasv3schemasobjectTypeArchive**](CoreApi.md#DeletecrmObjectSchemasv3schemasobjectTypeArchive) | **Delete** /crm/v3/schemas/{objectType} | Delete a schema
[**DeletecrmObjectSchemasv3schemasobjectTypeassociationsassociationIdentifierArchiveAssociation**](CoreApi.md#DeletecrmObjectSchemasv3schemasobjectTypeassociationsassociationIdentifierArchiveAssociation) | **Delete** /crm/v3/schemas/{objectType}/associations/{associationIdentifier} | Remove an association
[**GetcrmObjectSchemasv3schemasGetAll**](CoreApi.md#GetcrmObjectSchemasv3schemasGetAll) | **Get** /crm/v3/schemas | Get all schemas
[**GetcrmObjectSchemasv3schemasobjectTypeGetById**](CoreApi.md#GetcrmObjectSchemasv3schemasobjectTypeGetById) | **Get** /crm/v3/schemas/{objectType} | Get an existing schema
[**PatchcrmObjectSchemasv3schemasobjectTypeUpdate**](CoreApi.md#PatchcrmObjectSchemasv3schemasobjectTypeUpdate) | **Patch** /crm/v3/schemas/{objectType} | Update a schema
[**PostcrmObjectSchemasv3schemasCreate**](CoreApi.md#PostcrmObjectSchemasv3schemasCreate) | **Post** /crm/v3/schemas | Create a new schema
[**PostcrmObjectSchemasv3schemasobjectTypeassociationsCreateAssociation**](CoreApi.md#PostcrmObjectSchemasv3schemasobjectTypeassociationsCreateAssociation) | **Post** /crm/v3/schemas/{objectType}/associations | Create an association

# **DeletecrmObjectSchemasv3schemasobjectTypeArchive**
> DeletecrmObjectSchemasv3schemasobjectTypeArchive(ctx, objectType, optional)
Delete a schema

Deletes a schema. Any existing records of this schema must be deleted **first**. Otherwise this call will fail.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**| Fully qualified name or object type ID of your schema. | 
 **optional** | ***CoreApiDeletecrmObjectSchemasv3schemasobjectTypeArchiveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CoreApiDeletecrmObjectSchemasv3schemasobjectTypeArchiveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **optional.Bool**| Whether to return only results that have been archived. | [default to false]

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletecrmObjectSchemasv3schemasobjectTypeassociationsassociationIdentifierArchiveAssociation**
> DeletecrmObjectSchemasv3schemasobjectTypeassociationsassociationIdentifierArchiveAssociation(ctx, objectType, associationIdentifier)
Remove an association

Removes an existing association from a schema.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**| Fully qualified name or object type ID of your schema. | 
  **associationIdentifier** | **string**| Unique ID of the association to remove. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetcrmObjectSchemasv3schemasGetAll**
> CollectionResponseObjectSchemaNoPaging GetcrmObjectSchemasv3schemasGetAll(ctx, optional)
Get all schemas

Returns all object schemas that have been defined for your account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CoreApiGetcrmObjectSchemasv3schemasGetAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CoreApiGetcrmObjectSchemasv3schemasGetAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archived** | **optional.Bool**| Whether to return only results that have been archived. | [default to false]

### Return type

[**CollectionResponseObjectSchemaNoPaging**](CollectionResponseObjectSchemaNoPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetcrmObjectSchemasv3schemasobjectTypeGetById**
> ObjectSchema GetcrmObjectSchemasv3schemasobjectTypeGetById(ctx, objectType)
Get an existing schema

Returns an existing object schema.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**| Fully qualified name or object type ID of your schema. | 

### Return type

[**ObjectSchema**](ObjectSchema.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchcrmObjectSchemasv3schemasobjectTypeUpdate**
> ObjectTypeDefinition PatchcrmObjectSchemasv3schemasobjectTypeUpdate(ctx, body, objectType)
Update a schema

Update the details for an existing object schema.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ObjectTypeDefinitionPatch**](ObjectTypeDefinitionPatch.md)| Attributes to update in your schema. | 
  **objectType** | **string**| Fully qualified name or object type ID of your schema. | 

### Return type

[**ObjectTypeDefinition**](ObjectTypeDefinition.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostcrmObjectSchemasv3schemasCreate**
> ObjectSchema PostcrmObjectSchemasv3schemasCreate(ctx, body)
Create a new schema

Define a new object schema, along with custom properties and associations. The entire object schema, including its object type ID, properties, and associations will be returned in the response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ObjectSchemaEgg**](ObjectSchemaEgg.md)| Object schema definition, including properties and associations. | 

### Return type

[**ObjectSchema**](ObjectSchema.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostcrmObjectSchemasv3schemasobjectTypeassociationsCreateAssociation**
> AssociationDefinition PostcrmObjectSchemasv3schemasobjectTypeassociationsCreateAssociation(ctx, body, objectType)
Create an association

Defines a new association between the primary schema's object type and other object types.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AssociationDefinitionEgg**](AssociationDefinitionEgg.md)| Attributes that define the association. | 
  **objectType** | **string**| Fully qualified name or object type ID of your schema. | 

### Return type

[**AssociationDefinition**](AssociationDefinition.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

