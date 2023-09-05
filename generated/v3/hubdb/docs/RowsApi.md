# \RowsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneDraftTableRow**](RowsApi.md#CloneDraftTableRow) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft/clone | Clone a row
[**CreateTableRow**](RowsApi.md#CreateTableRow) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows | Add a new row to a table
[**GetDraftTableRowByID**](RowsApi.md#GetDraftTableRowByID) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft | Get a row from the draft table
[**GetTableRow**](RowsApi.md#GetTableRow) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId} | Get a table row
[**GetTableRows**](RowsApi.md#GetTableRows) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/rows | Get rows for a table
[**PurgeDraftTableRow**](RowsApi.md#PurgeDraftTableRow) | **Delete** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft | Permanently deletes a row
[**ReadDraftTableRows**](RowsApi.md#ReadDraftTableRows) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft | Get rows from draft table
[**ReplaceDraftTableRow**](RowsApi.md#ReplaceDraftTableRow) | **Put** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft | Replaces an existing row
[**UpdateDraftTableRow**](RowsApi.md#UpdateDraftTableRow) | **Patch** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft | Updates an existing row



## CloneDraftTableRow

> HubDbTableRowV3 CloneDraftTableRow(ctx, tableIdOrName, rowId).Execute()

Clone a row



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
    tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table
    rowId := "rowId_example" // string | The ID of the row

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RowsApi.CloneDraftTableRow(context.Background(), tableIdOrName, rowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsApi.CloneDraftTableRow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneDraftTableRow`: HubDbTableRowV3
    fmt.Fprintf(os.Stdout, "Response from `RowsApi.CloneDraftTableRow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 
**rowId** | **string** | The ID of the row | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneDraftTableRowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTableRow

> HubDbTableRowV3 CreateTableRow(ctx, tableIdOrName).HubDbTableRowV3Request(hubDbTableRowV3Request).Execute()

Add a new row to a table



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
    tableIdOrName := "tableIdOrName_example" // string | The ID or name of the target table.
    hubDbTableRowV3Request := *openapiclient.NewHubDbTableRowV3Request(map[string]map[string]interface{}{"key": map[string]interface{}(123)}) // HubDbTableRowV3Request | The row definition JSON, formatted as described above.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RowsApi.CreateTableRow(context.Background(), tableIdOrName).HubDbTableRowV3Request(hubDbTableRowV3Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsApi.CreateTableRow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTableRow`: HubDbTableRowV3
    fmt.Fprintf(os.Stdout, "Response from `RowsApi.CreateTableRow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the target table. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTableRowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hubDbTableRowV3Request** | [**HubDbTableRowV3Request**](HubDbTableRowV3Request.md) | The row definition JSON, formatted as described above. | 

### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDraftTableRowByID

> HubDbTableRowV3 GetDraftTableRowByID(ctx, tableIdOrName, rowId).Execute()

Get a row from the draft table



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
    tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table
    rowId := "rowId_example" // string | The ID of the row

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RowsApi.GetDraftTableRowByID(context.Background(), tableIdOrName, rowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsApi.GetDraftTableRowByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDraftTableRowByID`: HubDbTableRowV3
    fmt.Fprintf(os.Stdout, "Response from `RowsApi.GetDraftTableRowByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 
**rowId** | **string** | The ID of the row | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDraftTableRowByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTableRow

> HubDbTableRowV3 GetTableRow(ctx, tableIdOrName, rowId).Execute()

Get a table row



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
    tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table
    rowId := "rowId_example" // string | The ID of the row

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RowsApi.GetTableRow(context.Background(), tableIdOrName, rowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsApi.GetTableRow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTableRow`: HubDbTableRowV3
    fmt.Fprintf(os.Stdout, "Response from `RowsApi.GetTableRow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 
**rowId** | **string** | The ID of the row | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTableRowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTableRows

> CollectionResponseWithTotalHubDbTableRowV3ForwardPaging GetTableRows(ctx, tableIdOrName).Sort(sort).After(after).Limit(limit).Properties(properties).Execute()

Get rows for a table



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
    tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to query.
    sort := []string{"Inner_example"} // []string | Specifies the column names to sort the results by. See the above description for more details. (optional)
    after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
    limit := int32(56) // int32 | The maximum number of results to return. Default is `1000`. (optional)
    properties := []string{"Inner_example"} // []string | Specify the column names to get results containing only the required columns instead of all column details. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RowsApi.GetTableRows(context.Background(), tableIdOrName).Sort(sort).After(after).Limit(limit).Properties(properties).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsApi.GetTableRows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTableRows`: CollectionResponseWithTotalHubDbTableRowV3ForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `RowsApi.GetTableRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **[]string** | Specifies the column names to sort the results by. See the above description for more details. | 
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is &#x60;1000&#x60;. | 
 **properties** | **[]string** | Specify the column names to get results containing only the required columns instead of all column details. | 

### Return type

[**CollectionResponseWithTotalHubDbTableRowV3ForwardPaging**](CollectionResponseWithTotalHubDbTableRowV3ForwardPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurgeDraftTableRow

> PurgeDraftTableRow(ctx, tableIdOrName, rowId).Execute()

Permanently deletes a row



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
    tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table
    rowId := "rowId_example" // string | The ID of the row

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RowsApi.PurgeDraftTableRow(context.Background(), tableIdOrName, rowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsApi.PurgeDraftTableRow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 
**rowId** | **string** | The ID of the row | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurgeDraftTableRowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadDraftTableRows

> CollectionResponseWithTotalHubDbTableRowV3ForwardPaging ReadDraftTableRows(ctx, tableIdOrName).Sort(sort).After(after).Limit(limit).Properties(properties).Execute()

Get rows from draft table



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
    tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to query.
    sort := []string{"Inner_example"} // []string | Specifies the column names to sort the results by. (optional)
    after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
    limit := int32(56) // int32 | The maximum number of results to return. Default is `1000`. (optional)
    properties := []string{"Inner_example"} // []string | Specify the column names to get results containing only the required columns instead of all column details. If you want to include multiple columns in the result, use this query param as many times.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RowsApi.ReadDraftTableRows(context.Background(), tableIdOrName).Sort(sort).After(after).Limit(limit).Properties(properties).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsApi.ReadDraftTableRows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadDraftTableRows`: CollectionResponseWithTotalHubDbTableRowV3ForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `RowsApi.ReadDraftTableRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadDraftTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **[]string** | Specifies the column names to sort the results by. | 
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is &#x60;1000&#x60;. | 
 **properties** | **[]string** | Specify the column names to get results containing only the required columns instead of all column details. If you want to include multiple columns in the result, use this query param as many times.  | 

### Return type

[**CollectionResponseWithTotalHubDbTableRowV3ForwardPaging**](CollectionResponseWithTotalHubDbTableRowV3ForwardPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceDraftTableRow

> HubDbTableRowV3 ReplaceDraftTableRow(ctx, tableIdOrName, rowId).HubDbTableRowV3Request(hubDbTableRowV3Request).Execute()

Replaces an existing row



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
    tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table
    rowId := "rowId_example" // string | The ID of the row
    hubDbTableRowV3Request := *openapiclient.NewHubDbTableRowV3Request(map[string]map[string]interface{}{"key": map[string]interface{}(123)}) // HubDbTableRowV3Request | The JSON object of the row

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RowsApi.ReplaceDraftTableRow(context.Background(), tableIdOrName, rowId).HubDbTableRowV3Request(hubDbTableRowV3Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsApi.ReplaceDraftTableRow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceDraftTableRow`: HubDbTableRowV3
    fmt.Fprintf(os.Stdout, "Response from `RowsApi.ReplaceDraftTableRow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 
**rowId** | **string** | The ID of the row | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceDraftTableRowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hubDbTableRowV3Request** | [**HubDbTableRowV3Request**](HubDbTableRowV3Request.md) | The JSON object of the row | 

### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDraftTableRow

> HubDbTableRowV3 UpdateDraftTableRow(ctx, tableIdOrName, rowId).HubDbTableRowV3Request(hubDbTableRowV3Request).Execute()

Updates an existing row



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
    tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table
    rowId := "rowId_example" // string | The ID of the row
    hubDbTableRowV3Request := *openapiclient.NewHubDbTableRowV3Request(map[string]map[string]interface{}{"key": map[string]interface{}(123)}) // HubDbTableRowV3Request | The JSON object of the row with necessary fields that needs to be updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RowsApi.UpdateDraftTableRow(context.Background(), tableIdOrName, rowId).HubDbTableRowV3Request(hubDbTableRowV3Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsApi.UpdateDraftTableRow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDraftTableRow`: HubDbTableRowV3
    fmt.Fprintf(os.Stdout, "Response from `RowsApi.UpdateDraftTableRow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 
**rowId** | **string** | The ID of the row | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDraftTableRowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hubDbTableRowV3Request** | [**HubDbTableRowV3Request**](HubDbTableRowV3Request.md) | The JSON object of the row with necessary fields that needs to be updated. | 

### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

