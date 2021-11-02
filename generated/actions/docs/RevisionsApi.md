# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getautomationv4actionsappIddefinitionIdrevisionsGetPage**](RevisionsApi.md#Getautomationv4actionsappIddefinitionIdrevisionsGetPage) | **Get** /automation/v4/actions/{appId}/{definitionId}/revisions | Get all revisions for a custom action
[**Getautomationv4actionsappIddefinitionIdrevisionsrevisionIdGetById**](RevisionsApi.md#Getautomationv4actionsappIddefinitionIdrevisionsrevisionIdGetById) | **Get** /automation/v4/actions/{appId}/{definitionId}/revisions/{revisionId} | Get a revision for a custom action

# **Getautomationv4actionsappIddefinitionIdrevisionsGetPage**
> CollectionResponseActionRevisionForwardPaging Getautomationv4actionsappIddefinitionIdrevisionsGetPage(ctx, definitionId, appId, optional)
Get all revisions for a custom action

Returns a list of revisions for a custom workflow action.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **definitionId** | **string**| The ID of the custom workflow action | 
  **appId** | **int32**|  | 
 **optional** | ***RevisionsApiGetautomationv4actionsappIddefinitionIdrevisionsGetPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RevisionsApiGetautomationv4actionsappIddefinitionIdrevisionsGetPageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| Maximum number of results per page. | 
 **after** | **optional.String**| The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 

### Return type

[**CollectionResponseActionRevisionForwardPaging**](CollectionResponseActionRevisionForwardPaging.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getautomationv4actionsappIddefinitionIdrevisionsrevisionIdGetById**
> ActionRevision Getautomationv4actionsappIddefinitionIdrevisionsrevisionIdGetById(ctx, definitionId, revisionId, appId)
Get a revision for a custom action

Returns the given version of a custom workflow action.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **definitionId** | **string**| The ID of the custom workflow action. | 
  **revisionId** | **string**| The version of the custom workflow action. | 
  **appId** | **int32**|  | 

### Return type

[**ActionRevision**](ActionRevision.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

