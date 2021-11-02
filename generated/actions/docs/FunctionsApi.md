# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deleteautomationv4actionsappIddefinitionIdfunctionsfunctionTypeArchiveByFunctionType**](FunctionsApi.md#Deleteautomationv4actionsappIddefinitionIdfunctionsfunctionTypeArchiveByFunctionType) | **Delete** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType} | Delete a custom action function
[**Deleteautomationv4actionsappIddefinitionIdfunctionsfunctionTypefunctionIdArchive**](FunctionsApi.md#Deleteautomationv4actionsappIddefinitionIdfunctionsfunctionTypefunctionIdArchive) | **Delete** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId} | Delete a custom action function
[**Getautomationv4actionsappIddefinitionIdfunctionsGetPage**](FunctionsApi.md#Getautomationv4actionsappIddefinitionIdfunctionsGetPage) | **Get** /automation/v4/actions/{appId}/{definitionId}/functions | Get all custom action functions
[**Getautomationv4actionsappIddefinitionIdfunctionsfunctionTypeGetByFunctionType**](FunctionsApi.md#Getautomationv4actionsappIddefinitionIdfunctionsfunctionTypeGetByFunctionType) | **Get** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType} | Get a custom action function
[**Getautomationv4actionsappIddefinitionIdfunctionsfunctionTypefunctionIdGetById**](FunctionsApi.md#Getautomationv4actionsappIddefinitionIdfunctionsfunctionTypefunctionIdGetById) | **Get** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId} | Get a custom action function
[**Putautomationv4actionsappIddefinitionIdfunctionsfunctionTypeCreateOrReplaceByFunctionType**](FunctionsApi.md#Putautomationv4actionsappIddefinitionIdfunctionsfunctionTypeCreateOrReplaceByFunctionType) | **Put** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType} | Create or replace a custom action function
[**Putautomationv4actionsappIddefinitionIdfunctionsfunctionTypefunctionIdCreateOrReplace**](FunctionsApi.md#Putautomationv4actionsappIddefinitionIdfunctionsfunctionTypefunctionIdCreateOrReplace) | **Put** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId} | Create or replace a custom action function

# **Deleteautomationv4actionsappIddefinitionIdfunctionsfunctionTypeArchiveByFunctionType**
> Deleteautomationv4actionsappIddefinitionIdfunctionsfunctionTypeArchiveByFunctionType(ctx, definitionId, functionType, appId)
Delete a custom action function

Delete a function for a custom workflow action. This will remove the function itself as well as removing the association between the function and the custom action. This can't be undone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **definitionId** | **string**| The ID of the custom workflow action. | 
  **functionType** | **string**| The type of function. This determines when the function will be called. | 
  **appId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Deleteautomationv4actionsappIddefinitionIdfunctionsfunctionTypefunctionIdArchive**
> Deleteautomationv4actionsappIddefinitionIdfunctionsfunctionTypefunctionIdArchive(ctx, definitionId, functionType, functionId, appId)
Delete a custom action function

Delete a function for a custom workflow action. This will remove the function itself as well as removing the association between the function and the custom action. This can't be undone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **definitionId** | **string**| The ID of the custom workflow action | 
  **functionType** | **string**| The type of function. This determines when the function will be called. | 
  **functionId** | **string**| The ID qualifier for the function. This is used to specify which input field a function is associated with for &#x60;PRE_FETCH_OPTIONS&#x60; and &#x60;POST_FETCH_OPTIONS&#x60; function types. | 
  **appId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getautomationv4actionsappIddefinitionIdfunctionsGetPage**
> CollectionResponseActionFunctionIdentifierNoPaging Getautomationv4actionsappIddefinitionIdfunctionsGetPage(ctx, definitionId, appId)
Get all custom action functions

Returns a list of all functions that are associated with the given custom workflow action.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **definitionId** | **string**| The ID of the custom workflow action. | 
  **appId** | **int32**|  | 

### Return type

[**CollectionResponseActionFunctionIdentifierNoPaging**](CollectionResponseActionFunctionIdentifierNoPaging.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getautomationv4actionsappIddefinitionIdfunctionsfunctionTypeGetByFunctionType**
> ActionFunction Getautomationv4actionsappIddefinitionIdfunctionsfunctionTypeGetByFunctionType(ctx, definitionId, functionType, appId)
Get a custom action function

Returns the given function for a custom workflow action.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **definitionId** | **string**| The ID of the custom workflow action. | 
  **functionType** | **string**| The type of function. This determines when the function will be called. | 
  **appId** | **int32**|  | 

### Return type

[**ActionFunction**](ActionFunction.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getautomationv4actionsappIddefinitionIdfunctionsfunctionTypefunctionIdGetById**
> ActionFunction Getautomationv4actionsappIddefinitionIdfunctionsfunctionTypefunctionIdGetById(ctx, definitionId, functionType, functionId, appId)
Get a custom action function

Returns the given function for a custom workflow action.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **definitionId** | **string**| The ID of the custom workflow action. | 
  **functionType** | **string**| The type of function. This determines when the function will be called. | 
  **functionId** | **string**| The ID qualifier for the function. This is used to specify which input field a function is associated with for &#x60;PRE_FETCH_OPTIONS&#x60; and &#x60;POST_FETCH_OPTIONS&#x60; function types. | 
  **appId** | **int32**|  | 

### Return type

[**ActionFunction**](ActionFunction.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Putautomationv4actionsappIddefinitionIdfunctionsfunctionTypeCreateOrReplaceByFunctionType**
> ActionFunctionIdentifier Putautomationv4actionsappIddefinitionIdfunctionsfunctionTypeCreateOrReplaceByFunctionType(ctx, body, definitionId, functionType, appId)
Create or replace a custom action function

Creates or replaces a function for a custom workflow action.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| The function source code. Must be valid JavaScript code. | 
  **definitionId** | **string**| The ID of the custom workflow action. | 
  **functionType** | **string**| The type of function. This determines when the function will be called. | 
  **appId** | **int32**|  | 

### Return type

[**ActionFunctionIdentifier**](ActionFunctionIdentifier.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Putautomationv4actionsappIddefinitionIdfunctionsfunctionTypefunctionIdCreateOrReplace**
> ActionFunctionIdentifier Putautomationv4actionsappIddefinitionIdfunctionsfunctionTypefunctionIdCreateOrReplace(ctx, body, definitionId, functionType, functionId, appId)
Create or replace a custom action function

Creates or replaces a function for a custom workflow action.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| The function source code. Must be valid JavaScript code. | 
  **definitionId** | **string**| The ID of the custom workflow action. | 
  **functionType** | **string**| The type of function. This determines when the function will be called. | 
  **functionId** | **string**| The ID qualifier for the function. This is used to specify which input field a function is associated with for &#x60;PRE_FETCH_OPTIONS&#x60; and &#x60;POST_FETCH_OPTIONS&#x60; function types. | 
  **appId** | **int32**|  | 

### Return type

[**ActionFunctionIdentifier**](ActionFunctionIdentifier.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

