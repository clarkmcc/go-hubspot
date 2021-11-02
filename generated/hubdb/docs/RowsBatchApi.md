# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Postcmsv3hubdbtablestableIdOrNamerowsbatchreadBatchReadTableRows**](RowsBatchApi.md#Postcmsv3hubdbtablestableIdOrNamerowsbatchreadBatchReadTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/batch/read | Get a set of rows
[**Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchcloneBatchCloneDraftTableRows**](RowsBatchApi.md#Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchcloneBatchCloneDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/clone | Clone rows in batch
[**Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchcreateBatchCreateDraftTableRows**](RowsBatchApi.md#Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchcreateBatchCreateDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/create | Create rows in batch
[**Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchpurgeBatchPurgeDraftTableRows**](RowsBatchApi.md#Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchpurgeBatchPurgeDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/purge | Permanently deletes rows
[**Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchreadBatchReadDraftTableRows**](RowsBatchApi.md#Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchreadBatchReadDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/read | Get a set of rows from draft table
[**Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchreplaceBatchReplaceDraftTableRows**](RowsBatchApi.md#Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchreplaceBatchReplaceDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/replace | Replace rows in batch in draft table
[**Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchupdateBatchUpdateDraftTableRows**](RowsBatchApi.md#Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchupdateBatchUpdateDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/update | Update rows in batch in draft table

# **Postcmsv3hubdbtablestableIdOrNamerowsbatchreadBatchReadTableRows**
> BatchResponseHubDbTableRowV3 Postcmsv3hubdbtablestableIdOrNamerowsbatchreadBatchReadTableRows(ctx, body, tableIdOrName)
Get a set of rows

Returns rows in the `published` version of the specified table, given a set of row ids. **Note:** This endpoint can be accessed without any authentication if the table is set to be allowed for public access.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputString**](BatchInputString.md)| The JSON array of row ids | 
  **tableIdOrName** | **string**| The ID or name of the table to query. | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchcloneBatchCloneDraftTableRows**
> BatchResponseHubDbTableRowV3 Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchcloneBatchCloneDraftTableRows(ctx, body, tableIdOrName)
Clone rows in batch

Clones rows in the `draft` version of the specified table, given a set of row ids.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputString**](BatchInputString.md)| The JSON array of row ids | 
  **tableIdOrName** | **string**| The ID or name of the table | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchcreateBatchCreateDraftTableRows**
> BatchResponseHubDbTableRowV3 Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchcreateBatchCreateDraftTableRows(ctx, body, tableIdOrName)
Create rows in batch

Creates rows in the `draft` version of the specified table, given an array of row objects. See the overview section for more details with an example.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputHubDbTableRowV3Request**](BatchInputHubDbTableRowV3Request.md)| JSON array of row objects | 
  **tableIdOrName** | **string**| The ID or name of the table | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchpurgeBatchPurgeDraftTableRows**
> Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchpurgeBatchPurgeDraftTableRows(ctx, body, tableIdOrName)
Permanently deletes rows

Permanently deletes rows from the `draft` version of the table, given a set of row ids.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputString**](BatchInputString.md)| JSON array of row ids. | 
  **tableIdOrName** | **string**| The ID or name of the table | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchreadBatchReadDraftTableRows**
> BatchResponseHubDbTableRowV3 Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchreadBatchReadDraftTableRows(ctx, body, tableIdOrName)
Get a set of rows from draft table

Returns rows in the `draft` version of the specified table, given a set of row ids.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputString**](BatchInputString.md)| JSON array of row ids. | 
  **tableIdOrName** | **string**| The ID or name of the table | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchreplaceBatchReplaceDraftTableRows**
> BatchResponseHubDbTableRowV3 Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchreplaceBatchReplaceDraftTableRows(ctx, body, tableIdOrName)
Replace rows in batch in draft table

Replaces multiple rows as a batch in the `draft` version of the table. See the endpoint `PUT /tables/{tableIdOrName}/rows/{rowId}/draft` for details on updating a single row.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputHubDbTableRowV3Request**](BatchInputHubDbTableRowV3Request.md)| JSON array of row objects. | 
  **tableIdOrName** | **string**| The ID or name of the table | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchupdateBatchUpdateDraftTableRows**
> BatchResponseHubDbTableRowV3 Postcmsv3hubdbtablestableIdOrNamerowsdraftbatchupdateBatchUpdateDraftTableRows(ctx, body, tableIdOrName)
Update rows in batch in draft table

Updates multiple rows as a batch in the `draft` version of the table. See the endpoint `PATCH /tables/{tableIdOrName}/rows/{rowId}/draft` for details on updating a single row.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchInputJsonNode**](BatchInputJsonNode.md)| JSON array of row objects. | 
  **tableIdOrName** | **string**| The ID or name of the table | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

