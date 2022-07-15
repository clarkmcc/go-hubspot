# \TablesApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCmsV3HubdbTablesTableIdOrNameArchiveTable**](TablesApi.md#DeleteCmsV3HubdbTablesTableIdOrNameArchiveTable) | **Delete** /cms/v3/hubdb/tables/{tableIdOrName} | Archive a table
[**GetCmsV3HubdbTablesDraftGetAllDraftTables**](TablesApi.md#GetCmsV3HubdbTablesDraftGetAllDraftTables) | **Get** /cms/v3/hubdb/tables/draft | Return all draft tables
[**GetCmsV3HubdbTablesGetAllTables**](TablesApi.md#GetCmsV3HubdbTablesGetAllTables) | **Get** /cms/v3/hubdb/tables | Get all published tables
[**GetCmsV3HubdbTablesTableIdOrNameDraftExportExportDraftTable**](TablesApi.md#GetCmsV3HubdbTablesTableIdOrNameDraftExportExportDraftTable) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/draft/export | Export a draft table
[**GetCmsV3HubdbTablesTableIdOrNameDraftGetDraftTableDetailsById**](TablesApi.md#GetCmsV3HubdbTablesTableIdOrNameDraftGetDraftTableDetailsById) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/draft | Get details for a draft table
[**GetCmsV3HubdbTablesTableIdOrNameExportExportTable**](TablesApi.md#GetCmsV3HubdbTablesTableIdOrNameExportExportTable) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/export | Export a published version of a table
[**GetCmsV3HubdbTablesTableIdOrNameGetTableDetails**](TablesApi.md#GetCmsV3HubdbTablesTableIdOrNameGetTableDetails) | **Get** /cms/v3/hubdb/tables/{tableIdOrName} | Get details for a published table
[**PatchCmsV3HubdbTablesTableIdOrNameDraftUpdateDraftTable**](TablesApi.md#PatchCmsV3HubdbTablesTableIdOrNameDraftUpdateDraftTable) | **Patch** /cms/v3/hubdb/tables/{tableIdOrName}/draft | Update an existing table
[**PostCmsV3HubdbTablesCreateTable**](TablesApi.md#PostCmsV3HubdbTablesCreateTable) | **Post** /cms/v3/hubdb/tables | Create a new table
[**PostCmsV3HubdbTablesTableIdOrNameDraftCloneCloneDraftTable**](TablesApi.md#PostCmsV3HubdbTablesTableIdOrNameDraftCloneCloneDraftTable) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/draft/clone | Clone a table
[**PostCmsV3HubdbTablesTableIdOrNameDraftImportImportDraftTable**](TablesApi.md#PostCmsV3HubdbTablesTableIdOrNameDraftImportImportDraftTable) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/draft/import | Import data into draft table
[**PostCmsV3HubdbTablesTableIdOrNameDraftPublishPublishDraftTable**](TablesApi.md#PostCmsV3HubdbTablesTableIdOrNameDraftPublishPublishDraftTable) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/draft/publish | Publish a table from draft
[**PostCmsV3HubdbTablesTableIdOrNameDraftResetResetDraftTable**](TablesApi.md#PostCmsV3HubdbTablesTableIdOrNameDraftResetResetDraftTable) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/draft/reset | Reset a draft table
[**PostCmsV3HubdbTablesTableIdOrNameUnpublishUnpublishTable**](TablesApi.md#PostCmsV3HubdbTablesTableIdOrNameUnpublishUnpublishTable) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/unpublish | Unpublish a table



## DeleteCmsV3HubdbTablesTableIdOrNameArchiveTable

> DeleteCmsV3HubdbTablesTableIdOrNameArchiveTable(ctx, tableIdOrName).Execute()

Archive a table



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
    tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to archive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.DeleteCmsV3HubdbTablesTableIdOrNameArchiveTable(context.Background(), tableIdOrName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.DeleteCmsV3HubdbTablesTableIdOrNameArchiveTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCmsV3HubdbTablesTableIdOrNameArchiveTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3HubdbTablesDraftGetAllDraftTables

> CollectionResponseWithTotalHubDbTableV3ForwardPaging GetCmsV3HubdbTablesDraftGetAllDraftTables(ctx).Sort(sort).After(after).Limit(limit).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Archived(archived).Execute()

Return all draft tables



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    sort := []string{"Inner_example"} // []string | Specifies which fields to use for sorting results. Valid fields are `name`, `createdAt`, `updatedAt`, `createdBy`, `updatedBy`. `createdAt` will be used by default. (optional)
    after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
    limit := int32(56) // int32 | The maximum number of results to return. Default is 1000. (optional)
    createdAt := time.Now() // time.Time | Only return tables created at exactly the specified time. (optional)
    createdAfter := time.Now() // time.Time | Only return tables created after the specified time. (optional)
    createdBefore := time.Now() // time.Time | Only return tables created before the specified time. (optional)
    updatedAt := time.Now() // time.Time | Only return tables last updated at exactly the specified time. (optional)
    updatedAfter := time.Now() // time.Time | Only return tables last updated after the specified time. (optional)
    updatedBefore := time.Now() // time.Time | Only return tables last updated before the specified time. (optional)
    archived := true // bool | Specifies whether to return archived tables. Defaults to `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.GetCmsV3HubdbTablesDraftGetAllDraftTables(context.Background()).Sort(sort).After(after).Limit(limit).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.GetCmsV3HubdbTablesDraftGetAllDraftTables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3HubdbTablesDraftGetAllDraftTables`: CollectionResponseWithTotalHubDbTableV3ForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.GetCmsV3HubdbTablesDraftGetAllDraftTables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3HubdbTablesDraftGetAllDraftTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **[]string** | Specifies which fields to use for sorting results. Valid fields are &#x60;name&#x60;, &#x60;createdAt&#x60;, &#x60;updatedAt&#x60;, &#x60;createdBy&#x60;, &#x60;updatedBy&#x60;. &#x60;createdAt&#x60; will be used by default. | 
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is 1000. | 
 **createdAt** | **time.Time** | Only return tables created at exactly the specified time. | 
 **createdAfter** | **time.Time** | Only return tables created after the specified time. | 
 **createdBefore** | **time.Time** | Only return tables created before the specified time. | 
 **updatedAt** | **time.Time** | Only return tables last updated at exactly the specified time. | 
 **updatedAfter** | **time.Time** | Only return tables last updated after the specified time. | 
 **updatedBefore** | **time.Time** | Only return tables last updated before the specified time. | 
 **archived** | **bool** | Specifies whether to return archived tables. Defaults to &#x60;false&#x60;. | 

### Return type

[**CollectionResponseWithTotalHubDbTableV3ForwardPaging**](CollectionResponseWithTotalHubDbTableV3ForwardPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3HubdbTablesGetAllTables

> CollectionResponseWithTotalHubDbTableV3ForwardPaging GetCmsV3HubdbTablesGetAllTables(ctx).Sort(sort).After(after).Limit(limit).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Archived(archived).Execute()

Get all published tables



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    sort := []string{"Inner_example"} // []string | Specifies which fields to use for sorting results. Valid fields are `name`, `createdAt`, `updatedAt`, `createdBy`, `updatedBy`. `createdAt` will be used by default. (optional)
    after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
    limit := int32(56) // int32 | The maximum number of results to return. Default is 1000. (optional)
    createdAt := time.Now() // time.Time | Only return tables created at exactly the specified time. (optional)
    createdAfter := time.Now() // time.Time | Only return tables created after the specified time. (optional)
    createdBefore := time.Now() // time.Time | Only return tables created before the specified time. (optional)
    updatedAt := time.Now() // time.Time | Only return tables last updated at exactly the specified time. (optional)
    updatedAfter := time.Now() // time.Time | Only return tables last updated after the specified time. (optional)
    updatedBefore := time.Now() // time.Time | Only return tables last updated before the specified time. (optional)
    archived := true // bool | Specifies whether to return archived tables. Defaults to `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.GetCmsV3HubdbTablesGetAllTables(context.Background()).Sort(sort).After(after).Limit(limit).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.GetCmsV3HubdbTablesGetAllTables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3HubdbTablesGetAllTables`: CollectionResponseWithTotalHubDbTableV3ForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.GetCmsV3HubdbTablesGetAllTables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3HubdbTablesGetAllTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **[]string** | Specifies which fields to use for sorting results. Valid fields are &#x60;name&#x60;, &#x60;createdAt&#x60;, &#x60;updatedAt&#x60;, &#x60;createdBy&#x60;, &#x60;updatedBy&#x60;. &#x60;createdAt&#x60; will be used by default. | 
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is 1000. | 
 **createdAt** | **time.Time** | Only return tables created at exactly the specified time. | 
 **createdAfter** | **time.Time** | Only return tables created after the specified time. | 
 **createdBefore** | **time.Time** | Only return tables created before the specified time. | 
 **updatedAt** | **time.Time** | Only return tables last updated at exactly the specified time. | 
 **updatedAfter** | **time.Time** | Only return tables last updated after the specified time. | 
 **updatedBefore** | **time.Time** | Only return tables last updated before the specified time. | 
 **archived** | **bool** | Specifies whether to return archived tables. Defaults to &#x60;false&#x60;. | 

### Return type

[**CollectionResponseWithTotalHubDbTableV3ForwardPaging**](CollectionResponseWithTotalHubDbTableV3ForwardPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3HubdbTablesTableIdOrNameDraftExportExportDraftTable

> *os.File GetCmsV3HubdbTablesTableIdOrNameDraftExportExportDraftTable(ctx, tableIdOrName).Format(format).Execute()

Export a draft table



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
    tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to export.
    format := "format_example" // string | The file format to export. Possible values include `CSV`, `XLSX`, and `XLS`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.GetCmsV3HubdbTablesTableIdOrNameDraftExportExportDraftTable(context.Background(), tableIdOrName).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.GetCmsV3HubdbTablesTableIdOrNameDraftExportExportDraftTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3HubdbTablesTableIdOrNameDraftExportExportDraftTable`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.GetCmsV3HubdbTablesTableIdOrNameDraftExportExportDraftTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to export. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3HubdbTablesTableIdOrNameDraftExportExportDraftTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** | The file format to export. Possible values include &#x60;CSV&#x60;, &#x60;XLSX&#x60;, and &#x60;XLS&#x60;. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ms-excel, text/csv, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3HubdbTablesTableIdOrNameDraftGetDraftTableDetailsById

> HubDbTableV3 GetCmsV3HubdbTablesTableIdOrNameDraftGetDraftTableDetailsById(ctx, tableIdOrName).Archived(archived).IncludeForeignIds(includeForeignIds).Execute()

Get details for a draft table



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
    tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to return.
    archived := true // bool | Set this to `true` to return an archived table. Defaults to `false`. (optional)
    includeForeignIds := true // bool | Set this to `true` to populate foreign ID values in the result. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.GetCmsV3HubdbTablesTableIdOrNameDraftGetDraftTableDetailsById(context.Background(), tableIdOrName).Archived(archived).IncludeForeignIds(includeForeignIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.GetCmsV3HubdbTablesTableIdOrNameDraftGetDraftTableDetailsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3HubdbTablesTableIdOrNameDraftGetDraftTableDetailsById`: HubDbTableV3
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.GetCmsV3HubdbTablesTableIdOrNameDraftGetDraftTableDetailsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3HubdbTablesTableIdOrNameDraftGetDraftTableDetailsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Set this to &#x60;true&#x60; to return an archived table. Defaults to &#x60;false&#x60;. | 
 **includeForeignIds** | **bool** | Set this to &#x60;true&#x60; to populate foreign ID values in the result. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3HubdbTablesTableIdOrNameExportExportTable

> *os.File GetCmsV3HubdbTablesTableIdOrNameExportExportTable(ctx, tableIdOrName).Format(format).Execute()

Export a published version of a table



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
    tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to export.
    format := "format_example" // string | The file format to export. Possible values include `CSV`, `XLSX`, and `XLS`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.GetCmsV3HubdbTablesTableIdOrNameExportExportTable(context.Background(), tableIdOrName).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.GetCmsV3HubdbTablesTableIdOrNameExportExportTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3HubdbTablesTableIdOrNameExportExportTable`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.GetCmsV3HubdbTablesTableIdOrNameExportExportTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to export. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3HubdbTablesTableIdOrNameExportExportTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** | The file format to export. Possible values include &#x60;CSV&#x60;, &#x60;XLSX&#x60;, and &#x60;XLS&#x60;. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ms-excel, text/csv, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3HubdbTablesTableIdOrNameGetTableDetails

> HubDbTableV3 GetCmsV3HubdbTablesTableIdOrNameGetTableDetails(ctx, tableIdOrName).Archived(archived).IncludeForeignIds(includeForeignIds).Execute()

Get details for a published table



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
    tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to return.
    archived := true // bool | Set this to `true` to return details for an archived table. Defaults to `false`. (optional)
    includeForeignIds := true // bool | Set this to `true` to populate foreign ID values in the result. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.GetCmsV3HubdbTablesTableIdOrNameGetTableDetails(context.Background(), tableIdOrName).Archived(archived).IncludeForeignIds(includeForeignIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.GetCmsV3HubdbTablesTableIdOrNameGetTableDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3HubdbTablesTableIdOrNameGetTableDetails`: HubDbTableV3
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.GetCmsV3HubdbTablesTableIdOrNameGetTableDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3HubdbTablesTableIdOrNameGetTableDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Set this to &#x60;true&#x60; to return details for an archived table. Defaults to &#x60;false&#x60;. | 
 **includeForeignIds** | **bool** | Set this to &#x60;true&#x60; to populate foreign ID values in the result. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCmsV3HubdbTablesTableIdOrNameDraftUpdateDraftTable

> HubDbTableV3 PatchCmsV3HubdbTablesTableIdOrNameDraftUpdateDraftTable(ctx, tableIdOrName).HubDbTableV3Request(hubDbTableV3Request).Archived(archived).IncludeForeignIds(includeForeignIds).Execute()

Update an existing table



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
    tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to update.
    hubDbTableV3Request := *openapiclient.NewHubDbTableV3Request("Name_example", "Label_example") // HubDbTableV3Request | The JSON schema for the table being updated.
    archived := true // bool | Specifies whether to return archived tables. Defaults to `false`. (optional)
    includeForeignIds := true // bool | Set this to `true` to populate foreign ID values in the result. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.PatchCmsV3HubdbTablesTableIdOrNameDraftUpdateDraftTable(context.Background(), tableIdOrName).HubDbTableV3Request(hubDbTableV3Request).Archived(archived).IncludeForeignIds(includeForeignIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.PatchCmsV3HubdbTablesTableIdOrNameDraftUpdateDraftTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCmsV3HubdbTablesTableIdOrNameDraftUpdateDraftTable`: HubDbTableV3
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.PatchCmsV3HubdbTablesTableIdOrNameDraftUpdateDraftTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCmsV3HubdbTablesTableIdOrNameDraftUpdateDraftTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hubDbTableV3Request** | [**HubDbTableV3Request**](HubDbTableV3Request.md) | The JSON schema for the table being updated. | 
 **archived** | **bool** | Specifies whether to return archived tables. Defaults to &#x60;false&#x60;. | 
 **includeForeignIds** | **bool** | Set this to &#x60;true&#x60; to populate foreign ID values in the result. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesCreateTable

> HubDbTableV3 PostCmsV3HubdbTablesCreateTable(ctx).HubDbTableV3Request(hubDbTableV3Request).Execute()

Create a new table



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
    hubDbTableV3Request := *openapiclient.NewHubDbTableV3Request("Name_example", "Label_example") // HubDbTableV3Request | The JSON schema for the table being created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.PostCmsV3HubdbTablesCreateTable(context.Background()).HubDbTableV3Request(hubDbTableV3Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.PostCmsV3HubdbTablesCreateTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3HubdbTablesCreateTable`: HubDbTableV3
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.PostCmsV3HubdbTablesCreateTable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesCreateTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hubDbTableV3Request** | [**HubDbTableV3Request**](HubDbTableV3Request.md) | The JSON schema for the table being created. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameDraftCloneCloneDraftTable

> HubDbTableV3 PostCmsV3HubdbTablesTableIdOrNameDraftCloneCloneDraftTable(ctx, tableIdOrName).HubDbTableCloneRequest(hubDbTableCloneRequest).Execute()

Clone a table



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
    tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to clone.
    hubDbTableCloneRequest := *openapiclient.NewHubDbTableCloneRequest(false) // HubDbTableCloneRequest | JSON object with the properties newName and newLabel. You can set copyRows to false to clone the table with copying rows and default is true.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.PostCmsV3HubdbTablesTableIdOrNameDraftCloneCloneDraftTable(context.Background(), tableIdOrName).HubDbTableCloneRequest(hubDbTableCloneRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.PostCmsV3HubdbTablesTableIdOrNameDraftCloneCloneDraftTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3HubdbTablesTableIdOrNameDraftCloneCloneDraftTable`: HubDbTableV3
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.PostCmsV3HubdbTablesTableIdOrNameDraftCloneCloneDraftTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to clone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameDraftCloneCloneDraftTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hubDbTableCloneRequest** | [**HubDbTableCloneRequest**](HubDbTableCloneRequest.md) | JSON object with the properties newName and newLabel. You can set copyRows to false to clone the table with copying rows and default is true. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameDraftImportImportDraftTable

> ImportResult PostCmsV3HubdbTablesTableIdOrNameDraftImportImportDraftTable(ctx, tableIdOrName).Config(config).File(file).Execute()

Import data into draft table



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
    tableIdOrName := "tableIdOrName_example" // string | The ID of the destination table where data will be imported.
    config := "config_example" // string | Configuration for the import in JSON format as described above. (optional)
    file := os.NewFile(1234, "some_file") // *os.File | The source CSV file to be imported. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.PostCmsV3HubdbTablesTableIdOrNameDraftImportImportDraftTable(context.Background(), tableIdOrName).Config(config).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.PostCmsV3HubdbTablesTableIdOrNameDraftImportImportDraftTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3HubdbTablesTableIdOrNameDraftImportImportDraftTable`: ImportResult
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.PostCmsV3HubdbTablesTableIdOrNameDraftImportImportDraftTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID of the destination table where data will be imported. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameDraftImportImportDraftTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **config** | **string** | Configuration for the import in JSON format as described above. | 
 **file** | ***os.File** | The source CSV file to be imported. | 

### Return type

[**ImportResult**](ImportResult.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameDraftPublishPublishDraftTable

> HubDbTableV3 PostCmsV3HubdbTablesTableIdOrNameDraftPublishPublishDraftTable(ctx, tableIdOrName).IncludeForeignIds(includeForeignIds).Execute()

Publish a table from draft



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
    tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to publish.
    includeForeignIds := true // bool | Set this to `true` to populate foreign ID values in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.PostCmsV3HubdbTablesTableIdOrNameDraftPublishPublishDraftTable(context.Background(), tableIdOrName).IncludeForeignIds(includeForeignIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.PostCmsV3HubdbTablesTableIdOrNameDraftPublishPublishDraftTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3HubdbTablesTableIdOrNameDraftPublishPublishDraftTable`: HubDbTableV3
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.PostCmsV3HubdbTablesTableIdOrNameDraftPublishPublishDraftTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to publish. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameDraftPublishPublishDraftTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeForeignIds** | **bool** | Set this to &#x60;true&#x60; to populate foreign ID values in the response. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameDraftResetResetDraftTable

> HubDbTableV3 PostCmsV3HubdbTablesTableIdOrNameDraftResetResetDraftTable(ctx, tableIdOrName).IncludeForeignIds(includeForeignIds).Execute()

Reset a draft table



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
    tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to reset.
    includeForeignIds := true // bool | Set this to `true` to populate foreign ID values in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.PostCmsV3HubdbTablesTableIdOrNameDraftResetResetDraftTable(context.Background(), tableIdOrName).IncludeForeignIds(includeForeignIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.PostCmsV3HubdbTablesTableIdOrNameDraftResetResetDraftTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3HubdbTablesTableIdOrNameDraftResetResetDraftTable`: HubDbTableV3
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.PostCmsV3HubdbTablesTableIdOrNameDraftResetResetDraftTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to reset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameDraftResetResetDraftTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeForeignIds** | **bool** | Set this to &#x60;true&#x60; to populate foreign ID values in the response. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameUnpublishUnpublishTable

> HubDbTableV3 PostCmsV3HubdbTablesTableIdOrNameUnpublishUnpublishTable(ctx, tableIdOrName).IncludeForeignIds(includeForeignIds).Execute()

Unpublish a table



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
    tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to publish.
    includeForeignIds := true // bool | Set this to `true` to populate foreign ID values in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.PostCmsV3HubdbTablesTableIdOrNameUnpublishUnpublishTable(context.Background(), tableIdOrName).IncludeForeignIds(includeForeignIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.PostCmsV3HubdbTablesTableIdOrNameUnpublishUnpublishTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3HubdbTablesTableIdOrNameUnpublishUnpublishTable`: HubDbTableV3
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.PostCmsV3HubdbTablesTableIdOrNameUnpublishUnpublishTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to publish. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameUnpublishUnpublishTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeForeignIds** | **bool** | Set this to &#x60;true&#x60; to populate foreign ID values in the response. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

