# \TablesApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveTable**](TablesApi.md#ArchiveTable) | **Delete** /cms/v3/hubdb/tables/{tableIdOrName} | Archive a table
[**CloneDraftTable**](TablesApi.md#CloneDraftTable) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/draft/clone | Clone a table
[**CreateTable**](TablesApi.md#CreateTable) | **Post** /cms/v3/hubdb/tables | Create a new table
[**ExportDraftTable**](TablesApi.md#ExportDraftTable) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/draft/export | Export a draft table
[**ExportTable**](TablesApi.md#ExportTable) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/export | Export a published version of a table
[**GetAllDraftTables**](TablesApi.md#GetAllDraftTables) | **Get** /cms/v3/hubdb/tables/draft | Return all draft tables
[**GetAllTables**](TablesApi.md#GetAllTables) | **Get** /cms/v3/hubdb/tables | Get all published tables
[**GetDraftTableDetailsByID**](TablesApi.md#GetDraftTableDetailsByID) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/draft | Get details for a draft table
[**GetTableDetails**](TablesApi.md#GetTableDetails) | **Get** /cms/v3/hubdb/tables/{tableIdOrName} | Get details for a published table
[**ImportDraftTable**](TablesApi.md#ImportDraftTable) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/draft/import | Import data into draft table
[**PublishDraftTable**](TablesApi.md#PublishDraftTable) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/draft/publish | Publish a table from draft
[**ResetDraftTable**](TablesApi.md#ResetDraftTable) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/draft/reset | Reset a draft table
[**UnpublishTable**](TablesApi.md#UnpublishTable) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/unpublish | Unpublish a table
[**UpdateDraftTable**](TablesApi.md#UpdateDraftTable) | **Patch** /cms/v3/hubdb/tables/{tableIdOrName}/draft | Update an existing table



## ArchiveTable

> ArchiveTable(ctx, tableIdOrName).Execute()

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
    resp, r, err := apiClient.TablesApi.ArchiveTable(context.Background(), tableIdOrName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.ArchiveTable``: %v\n", err)
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

Other parameters are passed through a pointer to a apiArchiveTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloneDraftTable

> HubDbTableV3 CloneDraftTable(ctx, tableIdOrName).HubDbTableCloneRequest(hubDbTableCloneRequest).Execute()

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
    resp, r, err := apiClient.TablesApi.CloneDraftTable(context.Background(), tableIdOrName).HubDbTableCloneRequest(hubDbTableCloneRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.CloneDraftTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneDraftTable`: HubDbTableV3
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.CloneDraftTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to clone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneDraftTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hubDbTableCloneRequest** | [**HubDbTableCloneRequest**](HubDbTableCloneRequest.md) | JSON object with the properties newName and newLabel. You can set copyRows to false to clone the table with copying rows and default is true. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTable

> HubDbTableV3 CreateTable(ctx).HubDbTableV3Request(hubDbTableV3Request).Execute()

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
    resp, r, err := apiClient.TablesApi.CreateTable(context.Background()).HubDbTableV3Request(hubDbTableV3Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.CreateTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTable`: HubDbTableV3
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.CreateTable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hubDbTableV3Request** | [**HubDbTableV3Request**](HubDbTableV3Request.md) | The JSON schema for the table being created. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportDraftTable

> *os.File ExportDraftTable(ctx, tableIdOrName).Format(format).Execute()

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
    resp, r, err := apiClient.TablesApi.ExportDraftTable(context.Background(), tableIdOrName).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.ExportDraftTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportDraftTable`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.ExportDraftTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to export. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportDraftTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** | The file format to export. Possible values include &#x60;CSV&#x60;, &#x60;XLSX&#x60;, and &#x60;XLS&#x60;. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ms-excel, text/csv, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportTable

> *os.File ExportTable(ctx, tableIdOrName).Format(format).Execute()

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
    resp, r, err := apiClient.TablesApi.ExportTable(context.Background(), tableIdOrName).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.ExportTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportTable`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.ExportTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to export. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** | The file format to export. Possible values include &#x60;CSV&#x60;, &#x60;XLSX&#x60;, and &#x60;XLS&#x60;. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ms-excel, text/csv, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllDraftTables

> CollectionResponseWithTotalHubDbTableV3ForwardPaging GetAllDraftTables(ctx).Sort(sort).After(after).Limit(limit).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Archived(archived).Execute()

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
    resp, r, err := apiClient.TablesApi.GetAllDraftTables(context.Background()).Sort(sort).After(after).Limit(limit).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.GetAllDraftTables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllDraftTables`: CollectionResponseWithTotalHubDbTableV3ForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.GetAllDraftTables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllDraftTablesRequest struct via the builder pattern


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

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllTables

> CollectionResponseWithTotalHubDbTableV3ForwardPaging GetAllTables(ctx).Sort(sort).After(after).Limit(limit).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Archived(archived).Execute()

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
    resp, r, err := apiClient.TablesApi.GetAllTables(context.Background()).Sort(sort).After(after).Limit(limit).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.GetAllTables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllTables`: CollectionResponseWithTotalHubDbTableV3ForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.GetAllTables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTablesRequest struct via the builder pattern


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

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDraftTableDetailsByID

> HubDbTableV3 GetDraftTableDetailsByID(ctx, tableIdOrName).Archived(archived).IncludeForeignIds(includeForeignIds).Execute()

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
    resp, r, err := apiClient.TablesApi.GetDraftTableDetailsByID(context.Background(), tableIdOrName).Archived(archived).IncludeForeignIds(includeForeignIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.GetDraftTableDetailsByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDraftTableDetailsByID`: HubDbTableV3
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.GetDraftTableDetailsByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDraftTableDetailsByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Set this to &#x60;true&#x60; to return an archived table. Defaults to &#x60;false&#x60;. | 
 **includeForeignIds** | **bool** | Set this to &#x60;true&#x60; to populate foreign ID values in the result. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTableDetails

> HubDbTableV3 GetTableDetails(ctx, tableIdOrName).Archived(archived).IncludeForeignIds(includeForeignIds).Execute()

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
    resp, r, err := apiClient.TablesApi.GetTableDetails(context.Background(), tableIdOrName).Archived(archived).IncludeForeignIds(includeForeignIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.GetTableDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTableDetails`: HubDbTableV3
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.GetTableDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTableDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Set this to &#x60;true&#x60; to return details for an archived table. Defaults to &#x60;false&#x60;. | 
 **includeForeignIds** | **bool** | Set this to &#x60;true&#x60; to populate foreign ID values in the result. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportDraftTable

> ImportResult ImportDraftTable(ctx, tableIdOrName).Config(config).File(file).Execute()

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
    resp, r, err := apiClient.TablesApi.ImportDraftTable(context.Background(), tableIdOrName).Config(config).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.ImportDraftTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportDraftTable`: ImportResult
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.ImportDraftTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID of the destination table where data will be imported. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportDraftTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **config** | **string** | Configuration for the import in JSON format as described above. | 
 **file** | ***os.File** | The source CSV file to be imported. | 

### Return type

[**ImportResult**](ImportResult.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishDraftTable

> HubDbTableV3 PublishDraftTable(ctx, tableIdOrName).IncludeForeignIds(includeForeignIds).Execute()

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
    resp, r, err := apiClient.TablesApi.PublishDraftTable(context.Background(), tableIdOrName).IncludeForeignIds(includeForeignIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.PublishDraftTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublishDraftTable`: HubDbTableV3
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.PublishDraftTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to publish. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishDraftTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeForeignIds** | **bool** | Set this to &#x60;true&#x60; to populate foreign ID values in the response. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetDraftTable

> HubDbTableV3 ResetDraftTable(ctx, tableIdOrName).IncludeForeignIds(includeForeignIds).Execute()

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
    resp, r, err := apiClient.TablesApi.ResetDraftTable(context.Background(), tableIdOrName).IncludeForeignIds(includeForeignIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.ResetDraftTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetDraftTable`: HubDbTableV3
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.ResetDraftTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to reset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetDraftTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeForeignIds** | **bool** | Set this to &#x60;true&#x60; to populate foreign ID values in the response. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnpublishTable

> HubDbTableV3 UnpublishTable(ctx, tableIdOrName).IncludeForeignIds(includeForeignIds).Execute()

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
    resp, r, err := apiClient.TablesApi.UnpublishTable(context.Background(), tableIdOrName).IncludeForeignIds(includeForeignIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.UnpublishTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnpublishTable`: HubDbTableV3
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.UnpublishTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to publish. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnpublishTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeForeignIds** | **bool** | Set this to &#x60;true&#x60; to populate foreign ID values in the response. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDraftTable

> HubDbTableV3 UpdateDraftTable(ctx, tableIdOrName).HubDbTableV3Request(hubDbTableV3Request).Archived(archived).IncludeForeignIds(includeForeignIds).Execute()

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
    resp, r, err := apiClient.TablesApi.UpdateDraftTable(context.Background(), tableIdOrName).HubDbTableV3Request(hubDbTableV3Request).Archived(archived).IncludeForeignIds(includeForeignIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.UpdateDraftTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDraftTable`: HubDbTableV3
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.UpdateDraftTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDraftTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hubDbTableV3Request** | [**HubDbTableV3Request**](HubDbTableV3Request.md) | The JSON schema for the table being updated. | 
 **archived** | **bool** | Specifies whether to return archived tables. Defaults to &#x60;false&#x60;. | 
 **includeForeignIds** | **bool** | Set this to &#x60;true&#x60; to populate foreign ID values in the result. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

