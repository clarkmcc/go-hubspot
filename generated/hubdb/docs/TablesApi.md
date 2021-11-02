# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletecmsv3hubdbtablestableIdOrNameArchiveTable**](TablesApi.md#Deletecmsv3hubdbtablestableIdOrNameArchiveTable) | **Delete** /cms/v3/hubdb/tables/{tableIdOrName} | Archive a table
[**Getcmsv3hubdbtablesGetAllTables**](TablesApi.md#Getcmsv3hubdbtablesGetAllTables) | **Get** /cms/v3/hubdb/tables | Get all published tables
[**Getcmsv3hubdbtablesdraftGetAllDraftTables**](TablesApi.md#Getcmsv3hubdbtablesdraftGetAllDraftTables) | **Get** /cms/v3/hubdb/tables/draft | Return all draft tables
[**Getcmsv3hubdbtablestableIdOrNameGetTableDetails**](TablesApi.md#Getcmsv3hubdbtablestableIdOrNameGetTableDetails) | **Get** /cms/v3/hubdb/tables/{tableIdOrName} | Get details for a published table
[**Getcmsv3hubdbtablestableIdOrNamedraftGetDraftTableDetailsById**](TablesApi.md#Getcmsv3hubdbtablestableIdOrNamedraftGetDraftTableDetailsById) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/draft | Get details for a draft table
[**Getcmsv3hubdbtablestableIdOrNamedraftexportExportDraftTable**](TablesApi.md#Getcmsv3hubdbtablestableIdOrNamedraftexportExportDraftTable) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/draft/export | Export a draft table
[**Getcmsv3hubdbtablestableIdOrNameexportExportTable**](TablesApi.md#Getcmsv3hubdbtablestableIdOrNameexportExportTable) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/export | Export a published version of a table
[**Patchcmsv3hubdbtablestableIdOrNamedraftUpdateDraftTable**](TablesApi.md#Patchcmsv3hubdbtablestableIdOrNamedraftUpdateDraftTable) | **Patch** /cms/v3/hubdb/tables/{tableIdOrName}/draft | Update an existing table
[**Postcmsv3hubdbtablesCreateTable**](TablesApi.md#Postcmsv3hubdbtablesCreateTable) | **Post** /cms/v3/hubdb/tables | Create a new table
[**Postcmsv3hubdbtablestableIdOrNamedraftcloneCloneDraftTable**](TablesApi.md#Postcmsv3hubdbtablestableIdOrNamedraftcloneCloneDraftTable) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/draft/clone | Clone a table
[**Postcmsv3hubdbtablestableIdOrNamedraftimportImportDraftTable**](TablesApi.md#Postcmsv3hubdbtablestableIdOrNamedraftimportImportDraftTable) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/draft/import | Import data into draft table
[**Postcmsv3hubdbtablestableIdOrNamedraftpublishPublishDraftTable**](TablesApi.md#Postcmsv3hubdbtablestableIdOrNamedraftpublishPublishDraftTable) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/draft/publish | Publish a table from draft
[**Postcmsv3hubdbtablestableIdOrNamedraftresetResetDraftTable**](TablesApi.md#Postcmsv3hubdbtablestableIdOrNamedraftresetResetDraftTable) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/draft/reset | Reset a draft table
[**Postcmsv3hubdbtablestableIdOrNameunpublishUnpublishTable**](TablesApi.md#Postcmsv3hubdbtablestableIdOrNameunpublishUnpublishTable) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/unpublish | Unpublish a table

# **Deletecmsv3hubdbtablestableIdOrNameArchiveTable**
> Deletecmsv3hubdbtablestableIdOrNameArchiveTable(ctx, tableIdOrName)
Archive a table

Archive (soft delete) an existing HubDB table. This archives both the published and draft versions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tableIdOrName** | **string**| The ID or name of the table to archive. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcmsv3hubdbtablesGetAllTables**
> CollectionResponseWithTotalHubDbTableV3ForwardPaging Getcmsv3hubdbtablesGetAllTables(ctx, optional)
Get all published tables

Returns the details for the `published` version of each table defined in an account, including column definitions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TablesApiGetcmsv3hubdbtablesGetAllTablesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TablesApiGetcmsv3hubdbtablesGetAllTablesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | [**optional.Interface of []string**](string.md)| Specifies which fields to use for sorting results. Valid fields are &#x60;name&#x60;, &#x60;createdAt&#x60;, &#x60;updatedAt&#x60;, &#x60;createdBy&#x60;, &#x60;updatedBy&#x60;. &#x60;createdAt&#x60; will be used by default. | 
 **after** | **optional.String**| The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **optional.Int32**| The maximum number of results to return. Default is 1000. | 
 **createdAt** | **optional.Time**| Only return tables created at exactly the specified time. | 
 **createdAfter** | **optional.Time**| Only return tables created after the specified time. | 
 **createdBefore** | **optional.Time**| Only return tables created before the specified time. | 
 **updatedAt** | **optional.Time**| Only return tables last updated at exactly the specified time. | 
 **updatedAfter** | **optional.Time**| Only return tables last updated after the specified time. | 
 **updatedBefore** | **optional.Time**| Only return tables last updated before the specified time. | 
 **archived** | **optional.Bool**| Specifies whether to return archived tables. Defaults to &#x60;false&#x60;. | 

### Return type

[**CollectionResponseWithTotalHubDbTableV3ForwardPaging**](CollectionResponseWithTotalHubDbTableV3ForwardPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcmsv3hubdbtablesdraftGetAllDraftTables**
> CollectionResponseWithTotalHubDbTableV3ForwardPaging Getcmsv3hubdbtablesdraftGetAllDraftTables(ctx, optional)
Return all draft tables

Returns the details for each draft table defined in the specified account, including column definitions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TablesApiGetcmsv3hubdbtablesdraftGetAllDraftTablesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TablesApiGetcmsv3hubdbtablesdraftGetAllDraftTablesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | [**optional.Interface of []string**](string.md)| Specifies which fields to use for sorting results. Valid fields are &#x60;name&#x60;, &#x60;createdAt&#x60;, &#x60;updatedAt&#x60;, &#x60;createdBy&#x60;, &#x60;updatedBy&#x60;. &#x60;createdAt&#x60; will be used by default. | 
 **after** | **optional.String**| The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **optional.Int32**| The maximum number of results to return. Default is 1000. | 
 **createdAt** | **optional.Time**| Only return tables created at exactly the specified time. | 
 **createdAfter** | **optional.Time**| Only return tables created after the specified time. | 
 **createdBefore** | **optional.Time**| Only return tables created before the specified time. | 
 **updatedAt** | **optional.Time**| Only return tables last updated at exactly the specified time. | 
 **updatedAfter** | **optional.Time**| Only return tables last updated after the specified time. | 
 **updatedBefore** | **optional.Time**| Only return tables last updated before the specified time. | 
 **archived** | **optional.Bool**| Specifies whether to return archived tables. Defaults to &#x60;false&#x60;. | 

### Return type

[**CollectionResponseWithTotalHubDbTableV3ForwardPaging**](CollectionResponseWithTotalHubDbTableV3ForwardPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcmsv3hubdbtablestableIdOrNameGetTableDetails**
> HubDbTableV3 Getcmsv3hubdbtablestableIdOrNameGetTableDetails(ctx, tableIdOrName, optional)
Get details for a published table

Returns the details for the `published` version of the specified table. This will include the definitions for the columns in the table and the number of rows in the table.  **Note:** This endpoint can be accessed without any authentication if the table is set to be allowed for public access.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tableIdOrName** | **string**| The ID or name of the table to return. | 
 **optional** | ***TablesApiGetcmsv3hubdbtablestableIdOrNameGetTableDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TablesApiGetcmsv3hubdbtablestableIdOrNameGetTableDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **optional.Bool**| Set this to &#x60;true&#x60; to return details for an archived table. Defaults to &#x60;false&#x60;. | 
 **includeForeignIds** | **optional.Bool**| Set this to &#x60;true&#x60; to populate foreign ID values in the result. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcmsv3hubdbtablestableIdOrNamedraftGetDraftTableDetailsById**
> HubDbTableV3 Getcmsv3hubdbtablestableIdOrNamedraftGetDraftTableDetailsById(ctx, tableIdOrName, optional)
Get details for a draft table

Get the details for the `draft` version of a specific HubDB table. This will include the definitions for the columns in the table and the number of rows in the table.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tableIdOrName** | **string**| The ID or name of the table to return. | 
 **optional** | ***TablesApiGetcmsv3hubdbtablestableIdOrNamedraftGetDraftTableDetailsByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TablesApiGetcmsv3hubdbtablestableIdOrNamedraftGetDraftTableDetailsByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **optional.Bool**| Set this to &#x60;true&#x60; to return an archived table. Defaults to &#x60;false&#x60;. | 
 **includeForeignIds** | **optional.Bool**| Set this to &#x60;true&#x60; to populate foreign ID values in the result. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcmsv3hubdbtablestableIdOrNamedraftexportExportDraftTable**
> *os.File Getcmsv3hubdbtablestableIdOrNamedraftexportExportDraftTable(ctx, tableIdOrName, optional)
Export a draft table

Exports the `draft` version of a table to CSV / EXCEL format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tableIdOrName** | **string**| The ID or name of the table to export. | 
 **optional** | ***TablesApiGetcmsv3hubdbtablestableIdOrNamedraftexportExportDraftTableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TablesApiGetcmsv3hubdbtablestableIdOrNamedraftexportExportDraftTableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**| The file format to export. Possible values include &#x60;CSV&#x60;, &#x60;XLSX&#x60;, and &#x60;XLS&#x60;. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.ms-excel, text/csv, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcmsv3hubdbtablestableIdOrNameexportExportTable**
> *os.File Getcmsv3hubdbtablestableIdOrNameexportExportTable(ctx, tableIdOrName, optional)
Export a published version of a table

Exports the `published` version of a table to CSV / EXCEL format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tableIdOrName** | **string**| The ID or name of the table to export. | 
 **optional** | ***TablesApiGetcmsv3hubdbtablestableIdOrNameexportExportTableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TablesApiGetcmsv3hubdbtablestableIdOrNameexportExportTableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**| The file format to export. Possible values include &#x60;CSV&#x60;, &#x60;XLSX&#x60;, and &#x60;XLS&#x60;. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.ms-excel, text/csv, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patchcmsv3hubdbtablestableIdOrNamedraftUpdateDraftTable**
> HubDbTableV3 Patchcmsv3hubdbtablestableIdOrNamedraftUpdateDraftTable(ctx, body, tableIdOrName, optional)
Update an existing table

Update an existing HubDB table. You can use this endpoint to add or remove columns to the table as well as restore an archived table. Tables updated using the endpoint will only modify the `draft` verion of the table. Use `publish` endpoint to push all the changes to the `published` version. To restore a table, include the query parameter `archived=true` and `\"archived\": false` in the json body. **Note:** You need to include all the columns in the input when you are adding/removing/updating a column. If you do not include an already existing column in the request, it will be deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HubDbTableV3Request**](HubDbTableV3Request.md)| The JSON schema for the table being updated. | 
  **tableIdOrName** | **string**| The ID or name of the table to update. | 
 **optional** | ***TablesApiPatchcmsv3hubdbtablestableIdOrNamedraftUpdateDraftTableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TablesApiPatchcmsv3hubdbtablestableIdOrNamedraftUpdateDraftTableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **archived** | **optional.**| Specifies whether to return archived tables. Defaults to &#x60;false&#x60;. | 
 **includeForeignIds** | **optional.**| Set this to &#x60;true&#x60; to populate foreign ID values in the result. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3hubdbtablesCreateTable**
> HubDbTableV3 Postcmsv3hubdbtablesCreateTable(ctx, body)
Create a new table

Creates a new draft HubDB table given a JSON schema. The table name and label should be unique for each account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HubDbTableV3Request**](HubDbTableV3Request.md)| The JSON schema for the table being created. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3hubdbtablestableIdOrNamedraftcloneCloneDraftTable**
> HubDbTableV3 Postcmsv3hubdbtablestableIdOrNamedraftcloneCloneDraftTable(ctx, body, tableIdOrName)
Clone a table

Clone an existing HubDB table. The `newName` and `newLabel` of the new table can be sent as JSON in the `body` parameter. This will create the cloned table as a `draft`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HubDbTableCloneRequest**](HubDbTableCloneRequest.md)| JSON object with the properties newName and newLabel. You can set copyRows to false to clone the table with copying rows and default is true. | 
  **tableIdOrName** | **string**| The ID or name of the table to clone. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3hubdbtablestableIdOrNamedraftimportImportDraftTable**
> ImportResult Postcmsv3hubdbtablestableIdOrNamedraftimportImportDraftTable(ctx, tableIdOrName, optional)
Import data into draft table

Import the contents of a CSV file into an existing HubDB table. The data will always be imported into the `draft` version of the table. Use `/publish` endpoint to push these changes to `published` version. This endpoint takes a multi-part POST request. The first part will be a set of JSON-formatted options for the import and you can specify this with the name as `config`.  The second part will be the CSV file you want to import and you can specify this with the name as `file`. Refer the [overview section](https://developers.hubspot.com/docs/api/cms/hubdb#importing-tables) to check the details and format of the JSON-formatted options for the import.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tableIdOrName** | **string**| The ID of the destination table where data will be imported. | 
 **optional** | ***TablesApiPostcmsv3hubdbtablestableIdOrNamedraftimportImportDraftTableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TablesApiPostcmsv3hubdbtablestableIdOrNamedraftimportImportDraftTableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **config** | **optional.**|  | 
 **file** | **optional.Interface of *os.File****optional.**|  | 

### Return type

[**ImportResult**](ImportResult.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3hubdbtablestableIdOrNamedraftpublishPublishDraftTable**
> HubDbTableV3 Postcmsv3hubdbtablestableIdOrNamedraftpublishPublishDraftTable(ctx, tableIdOrName, optional)
Publish a table from draft

Publishes the table by copying the data and table schema changes from draft version to the published version, meaning any website pages using data from the table will be updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tableIdOrName** | **string**| The ID or name of the table to publish. | 
 **optional** | ***TablesApiPostcmsv3hubdbtablestableIdOrNamedraftpublishPublishDraftTableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TablesApiPostcmsv3hubdbtablestableIdOrNamedraftpublishPublishDraftTableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeForeignIds** | **optional.Bool**| Set this to &#x60;true&#x60; to populate foreign ID values in the response. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3hubdbtablestableIdOrNamedraftresetResetDraftTable**
> HubDbTableV3 Postcmsv3hubdbtablestableIdOrNamedraftresetResetDraftTable(ctx, tableIdOrName, optional)
Reset a draft table

Replaces the data in the `draft` version of the table with values from the `published` version. Any unpublished changes in the `draft` will be lost after this call is made.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tableIdOrName** | **string**| The ID or name of the table to reset. | 
 **optional** | ***TablesApiPostcmsv3hubdbtablestableIdOrNamedraftresetResetDraftTableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TablesApiPostcmsv3hubdbtablestableIdOrNamedraftresetResetDraftTableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeForeignIds** | **optional.Bool**| Set this to &#x60;true&#x60; to populate foreign ID values in the response. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcmsv3hubdbtablestableIdOrNameunpublishUnpublishTable**
> HubDbTableV3 Postcmsv3hubdbtablestableIdOrNameunpublishUnpublishTable(ctx, tableIdOrName, optional)
Unpublish a table

Unpublishes the table, meaning any website pages using data from the table will not render any data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tableIdOrName** | **string**| The ID or name of the table to publish. | 
 **optional** | ***TablesApiPostcmsv3hubdbtablestableIdOrNameunpublishUnpublishTableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TablesApiPostcmsv3hubdbtablestableIdOrNameunpublishUnpublishTableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeForeignIds** | **optional.Bool**| Set this to &#x60;true&#x60; to populate foreign ID values in the response. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

