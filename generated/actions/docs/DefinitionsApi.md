# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deleteautomationv4actionsappIddefinitionIdArchive**](DefinitionsApi.md#Deleteautomationv4actionsappIddefinitionIdArchive) | **Delete** /automation/v4/actions/{appId}/{definitionId} | Archive a custom action
[**Getautomationv4actionsappIdGetPage**](DefinitionsApi.md#Getautomationv4actionsappIdGetPage) | **Get** /automation/v4/actions/{appId} | Get all custom actions
[**Getautomationv4actionsappIddefinitionIdGetById**](DefinitionsApi.md#Getautomationv4actionsappIddefinitionIdGetById) | **Get** /automation/v4/actions/{appId}/{definitionId} | Get a custom action
[**Patchautomationv4actionsappIddefinitionIdUpdate**](DefinitionsApi.md#Patchautomationv4actionsappIddefinitionIdUpdate) | **Patch** /automation/v4/actions/{appId}/{definitionId} | Update a custom action
[**Postautomationv4actionsappIdCreate**](DefinitionsApi.md#Postautomationv4actionsappIdCreate) | **Post** /automation/v4/actions/{appId} | Create new custom action

# **Deleteautomationv4actionsappIddefinitionIdArchive**
> Deleteautomationv4actionsappIddefinitionIdArchive(ctx, definitionId, appId)
Archive a custom action

Archives a single custom workflow action with the specified ID. Workflows that currently use this custom action will stop attempting to execute the action, and all future executions will be marked as a failure.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **definitionId** | **string**| The ID of the custom workflow action. | 
  **appId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getautomationv4actionsappIdGetPage**
> CollectionResponseExtensionActionDefinitionForwardPaging Getautomationv4actionsappIdGetPage(ctx, appId, optional)
Get all custom actions

Returns a list of all custom workflow actions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**|  | 
 **optional** | ***DefinitionsApiGetautomationv4actionsappIdGetPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefinitionsApiGetautomationv4actionsappIdGetPageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Maximum number of results per page. | 
 **after** | **optional.String**| The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **archived** | **optional.Bool**| Whether to include archived custom actions. | [default to false]

### Return type

[**CollectionResponseExtensionActionDefinitionForwardPaging**](CollectionResponseExtensionActionDefinitionForwardPaging.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getautomationv4actionsappIddefinitionIdGetById**
> ExtensionActionDefinition Getautomationv4actionsappIddefinitionIdGetById(ctx, definitionId, appId, optional)
Get a custom action

Returns a single custom workflow action with the specified ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **definitionId** | **string**| The ID of the custom workflow action. | 
  **appId** | **int32**|  | 
 **optional** | ***DefinitionsApiGetautomationv4actionsappIddefinitionIdGetByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefinitionsApiGetautomationv4actionsappIddefinitionIdGetByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **archived** | **optional.Bool**| Whether to include archived custom actions. | [default to false]

### Return type

[**ExtensionActionDefinition**](ExtensionActionDefinition.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patchautomationv4actionsappIddefinitionIdUpdate**
> ExtensionActionDefinition Patchautomationv4actionsappIddefinitionIdUpdate(ctx, body, definitionId, appId)
Update a custom action

Updates a custom workflow action with new values for the specified fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExtensionActionDefinitionPatch**](ExtensionActionDefinitionPatch.md)| The custom workflow action fields to be updated. | 
  **definitionId** | **string**| The ID of the custom workflow action. | 
  **appId** | **int32**|  | 

### Return type

[**ExtensionActionDefinition**](ExtensionActionDefinition.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postautomationv4actionsappIdCreate**
> ExtensionActionDefinition Postautomationv4actionsappIdCreate(ctx, body, appId)
Create new custom action

Creates a new custom workflow action.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExtensionActionDefinitionInput**](ExtensionActionDefinitionInput.md)| The custom workflow action to create. | 
  **appId** | **int32**|  | 

### Return type

[**ExtensionActionDefinition**](ExtensionActionDefinition.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

