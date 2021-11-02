# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletecmsv3hubdbtablestableIdOrNamerowsrowIddraftPurgeDraftTableRow**](RowsApi.md#Deletecmsv3hubdbtablestableIdOrNamerowsrowIddraftPurgeDraftTableRow) | **Delete** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft | Permanently deletes a row
[**Getcmsv3hubdbtablestableIdOrNamerowsGetTableRows**](RowsApi.md#Getcmsv3hubdbtablestableIdOrNamerowsGetTableRows) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/rows | Get rows for a table
[**Getcmsv3hubdbtablestableIdOrNamerowsdraftReadDraftTableRows**](RowsApi.md#Getcmsv3hubdbtablestableIdOrNamerowsdraftReadDraftTableRows) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft | Get rows from draft table
[**Getcmsv3hubdbtablestableIdOrNamerowsrowIdGetTableRow**](RowsApi.md#Getcmsv3hubdbtablestableIdOrNamerowsrowIdGetTableRow) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId} | Get a table row
[**Getcmsv3hubdbtablestableIdOrNamerowsrowIddraftGetDraftTableRowById**](RowsApi.md#Getcmsv3hubdbtablestableIdOrNamerowsrowIddraftGetDraftTableRowById) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft | Get a row from the draft table
[**Patchcmsv3hubdbtablestableIdOrNamerowsrowIddraftUpdateDraftTableRow**](RowsApi.md#Patchcmsv3hubdbtablestableIdOrNamerowsrowIddraftUpdateDraftTableRow) | **Patch** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft | Updates an existing row
[**Postcmsv3hubdbtablestableIdOrNamerowsCreateTableRow**](RowsApi.md#Postcmsv3hubdbtablestableIdOrNamerowsCreateTableRow) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows | Add a new row to a table
[**Postcmsv3hubdbtablestableIdOrNamerowsrowIddraftcloneCloneDraftTableRow**](RowsApi.md#Postcmsv3hubdbtablestableIdOrNamerowsrowIddraftcloneCloneDraftTableRow) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft/clone | Clone a row
[**Putcmsv3hubdbtablestableIdOrNamerowsrowIddraftReplaceDraftTableRow**](RowsApi.md#Putcmsv3hubdbtablestableIdOrNamerowsrowIddraftReplaceDraftTableRow) | **Put** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft | Replaces an existing row

# **Deletecmsv3hubdbtablestableIdOrNamerowsrowIddraftPurgeDraftTableRow**
> Deletecmsv3hubdbtablestableIdOrNamerowsrowIddraftPurgeDraftTableRow(ctx, tableIdOrName, rowId)
Permanently deletes a row

Permanently deletes a row from a table's `draft` version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tableIdOrName** | **string**| The ID or name of the table | 
  **rowId** | **string**| The ID of the row | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcmsv3hubdbtablestableIdOrNamerowsGetTableRows**
> CollectionResponseWithTotalHubDbTableRowV3ForwardPaging Getcmsv3hubdbtablestableIdOrNamerowsGetTableRows(ctx, tableIdOrName, optional)
Get rows for a table

Returns a set of rows in the `published` version of the specified table. Row results can be filtered and sorted. Filtering and sorting options will be sent as query parameters to the API request. For example, by adding the query parameters `column1__gt=5&sort=-column1`, API returns the rows with values for column `column1` greater than 5 and in the descending order of `column1` values. Refer to the [overview section](https://developers.hubspot.com/docs/api/cms/hubdb#filtering-and-sorting-table-rows) for detailed filtering and sorting options. **Note:** This endpoint can be accessed without any authentication, if the table is set to be allowed for public access.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tableIdOrName** | **string**| The ID or name of the table to query. | 
 **optional** | ***RowsApiGetcmsv3hubdbtablestableIdOrNamerowsGetTableRowsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RowsApiGetcmsv3hubdbtablestableIdOrNamerowsGetTableRowsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | [**optional.Interface of []string**](string.md)| Specifies the column names to sort the results by. See the above description for more details. | 
 **after** | **optional.String**| The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **optional.Int32**| The maximum number of results to return. Default is &#x60;1000&#x60;. | 
 **properties** | [**optional.Interface of []string**](string.md)| Specify the column names to get results containing only the required columns instead of all column details. | 

### Return type

[**CollectionResponseWithTotalHubDbTableRowV3ForwardPaging**](CollectionResponseWithTotalHubDbTableRowV3ForwardPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcmsv3hubdbtablestableIdOrNamerowsdraftReadDraftTableRows**
> CollectionResponseWithTotalHubDbTableRowV3ForwardPaging Getcmsv3hubdbtablestableIdOrNamerowsdraftReadDraftTableRows(ctx, tableIdOrName, optional)
Get rows from draft table

Returns rows in the `draft` version of the specified table. Row results can be filtered and sorted. Filtering and sorting options will be sent as query parameters to the API request. For example, by adding the query parameters `column1__gt=5&sort=-column1`, API returns the rows with values for column `column1` greater than 5 and in the descending order of `column1` values. Refer to the [overview section](https://developers.hubspot.com/docs/api/cms/hubdb#filtering-and-sorting-table-rows) for detailed filtering and sorting options.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tableIdOrName** | **string**| The ID or name of the table to query. | 
 **optional** | ***RowsApiGetcmsv3hubdbtablestableIdOrNamerowsdraftReadDraftTableRowsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RowsApiGetcmsv3hubdbtablestableIdOrNamerowsdraftReadDraftTableRowsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | [**optional.Interface of []string**](string.md)| Specifies the column names to sort the results by. | 
 **after** | **optional.String**| The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **optional.Int32**| The maximum number of results to return. Default is &#x60;1000&#x60;. | 
 **properties** | [**optional.Interface of []string**](string.md)| Specify the column names to get results containing only the required columns instead of all column details. If you want to include multiple columns in the result, use this query param as many times.  | 

### Return type

[**CollectionResponseWithTotalHubDbTableRowV3ForwardPaging**](CollectionResponseWithTotalHubDbTableRowV3ForwardPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcmsv3hubdbtablestableIdOrNamerowsrowIdGetTableRow**
> HubDbTableRowV3 Getcmsv3hubdbtablestableIdOrNamerowsrowIdGetTableRow(ctx, tableIdOrName, rowId)
Get a table row

Get a single row by ID from a table's `published` version. **Note:** This endpoint can be accessed without any authentication, if the table is set to be allowed for public access.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tableIdOrName** | **string**| The ID or name of the table | 
  **rowId** | **string**| The ID of the row | 

### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcmsv3hubdbtablestableIdOrNamerowsrowIddraftGetDraftTableRowById**
> HubDbTableRowV3 Getcmsv3hubdbtablestableIdOrNamerowsrowIddraftGetDraftTableRowById(ctx, tableIdOrName, rowId)
Get a row from the draft table

Get a single row by ID from a table's `draft` version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tableIdOrName** | **string**| The ID or name of the table | 
  **rowId** | **string**| The ID of the row | 

### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patchcmsv3hubdbtablestableIdOrNamerowsrowIddraftUpdateDraftTableRow**
> HubDbTableRowV3 Patchcmsv3hubdbtablestableIdOrNamerowsrowIddraftUpdateDraftTableRow(ctx, body, tableIdOrName, rowId)
Updates an existing row

Sparse updates a single row in the table's `draft` version. All the column values need not be specified. Only the columns or fields that needs to be modified can be specified. See the `Create a row` endpoint for instructions on how to format the JSON row definitions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HubDbTableRowV3Request**](HubDbTableRowV3Request.md)| The JSON object of the row with necessary fields that needs to be updated. | 
  **tableIdOrName** | **string**| The ID or name of the table | 
  **rowId** | **string**| The ID of the row | 

### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3hubdbtablestableIdOrNamerowsCreateTableRow**
> HubDbTableRowV3 Postcmsv3hubdbtablestableIdOrNamerowsCreateTableRow(ctx, body, tableIdOrName)
Add a new row to a table

Add a new row to a HubDB table. New rows will be added to the `draft` version of the table. Use `publish` endpoint to push these changes to published version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HubDbTableRowV3Request**](HubDbTableRowV3Request.md)| The row definition JSON, formatted as described above. | 
  **tableIdOrName** | **string**| The ID or name of the target table. | 

### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3hubdbtablestableIdOrNamerowsrowIddraftcloneCloneDraftTableRow**
> HubDbTableRowV3 Postcmsv3hubdbtablestableIdOrNamerowsrowIddraftcloneCloneDraftTableRow(ctx, tableIdOrName, rowId)
Clone a row

Clones a single row in the `draft` version of the table.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tableIdOrName** | **string**| The ID or name of the table | 
  **rowId** | **string**| The ID of the row | 

### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Putcmsv3hubdbtablestableIdOrNamerowsrowIddraftReplaceDraftTableRow**
> HubDbTableRowV3 Putcmsv3hubdbtablestableIdOrNamerowsrowIddraftReplaceDraftTableRow(ctx, body, tableIdOrName, rowId)
Replaces an existing row

Replace a single row in the table's `draft` version. All the column values must be specified. If a column has a value in the target table and this request doesn't define that value, it will be deleted. See the `Create a row` endpoint for instructions on how to format the JSON row definitions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HubDbTableRowV3Request**](HubDbTableRowV3Request.md)| The JSON object of the row | 
  **tableIdOrName** | **string**| The ID or name of the table | 
  **rowId** | **string**| The ID of the row | 

### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

