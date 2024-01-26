# \RowsBatchApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCmsV3HubdbTablesTableIdOrNameRowsBatchReadReadTableRows**](RowsBatchApi.md#PostCmsV3HubdbTablesTableIdOrNameRowsBatchReadReadTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/batch/read | Get a set of rows
[**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneCloneDraftTableRows**](RowsBatchApi.md#PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneCloneDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/clone | Clone rows in batch
[**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateCreateDraftTableRows**](RowsBatchApi.md#PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateCreateDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/create | Create rows in batch
[**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgePurgeDraftTableRows**](RowsBatchApi.md#PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgePurgeDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/purge | Permanently deletes rows
[**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadReadDraftTableRows**](RowsBatchApi.md#PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadReadDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/read | Get a set of rows from draft table
[**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceReplaceDraftTableRows**](RowsBatchApi.md#PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceReplaceDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/replace | Replace rows in batch in draft table
[**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateUpdateDraftTableRows**](RowsBatchApi.md#PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateUpdateDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/update | Update rows in batch in draft table



## PostCmsV3HubdbTablesTableIdOrNameRowsBatchReadReadTableRows

> BatchResponseHubDbTableRowV3 PostCmsV3HubdbTablesTableIdOrNameRowsBatchReadReadTableRows(ctx, tableIdOrName).BatchInputString(batchInputString).Execute()

Get a set of rows



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
    batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of row ids

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsBatchReadReadTableRows(context.Background(), tableIdOrName).BatchInputString(batchInputString).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsBatchReadReadTableRows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3HubdbTablesTableIdOrNameRowsBatchReadReadTableRows`: BatchResponseHubDbTableRowV3
    fmt.Fprintf(os.Stdout, "Response from `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsBatchReadReadTableRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsBatchReadReadTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of row ids | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneCloneDraftTableRows

> BatchResponseHubDbTableRowV3 PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneCloneDraftTableRows(ctx, tableIdOrName).BatchInputString(batchInputString).Execute()

Clone rows in batch



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
    batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of row ids

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneCloneDraftTableRows(context.Background(), tableIdOrName).BatchInputString(batchInputString).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneCloneDraftTableRows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneCloneDraftTableRows`: BatchResponseHubDbTableRowV3
    fmt.Fprintf(os.Stdout, "Response from `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneCloneDraftTableRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneCloneDraftTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of row ids | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateCreateDraftTableRows

> BatchResponseHubDbTableRowV3 PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateCreateDraftTableRows(ctx, tableIdOrName).BatchInputHubDbTableRowV3Request(batchInputHubDbTableRowV3Request).Execute()

Create rows in batch



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
    batchInputHubDbTableRowV3Request := *openapiclient.NewBatchInputHubDbTableRowV3Request([]openapiclient.HubDbTableRowV3Request{*openapiclient.NewHubDbTableRowV3Request(map[string]map[string]interface{}{"key": map[string]interface{}(123)})}) // BatchInputHubDbTableRowV3Request | JSON array of row objects

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateCreateDraftTableRows(context.Background(), tableIdOrName).BatchInputHubDbTableRowV3Request(batchInputHubDbTableRowV3Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateCreateDraftTableRows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateCreateDraftTableRows`: BatchResponseHubDbTableRowV3
    fmt.Fprintf(os.Stdout, "Response from `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateCreateDraftTableRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateCreateDraftTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputHubDbTableRowV3Request** | [**BatchInputHubDbTableRowV3Request**](BatchInputHubDbTableRowV3Request.md) | JSON array of row objects | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgePurgeDraftTableRows

> PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgePurgeDraftTableRows(ctx, tableIdOrName).BatchInputString(batchInputString).Execute()

Permanently deletes rows



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
    batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | JSON array of row ids.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgePurgeDraftTableRows(context.Background(), tableIdOrName).BatchInputString(batchInputString).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgePurgeDraftTableRows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgePurgeDraftTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputString** | [**BatchInputString**](BatchInputString.md) | JSON array of row ids. | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadReadDraftTableRows

> BatchResponseHubDbTableRowV3 PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadReadDraftTableRows(ctx, tableIdOrName).BatchInputString(batchInputString).Execute()

Get a set of rows from draft table



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
    batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | JSON array of row ids.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadReadDraftTableRows(context.Background(), tableIdOrName).BatchInputString(batchInputString).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadReadDraftTableRows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadReadDraftTableRows`: BatchResponseHubDbTableRowV3
    fmt.Fprintf(os.Stdout, "Response from `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadReadDraftTableRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadReadDraftTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputString** | [**BatchInputString**](BatchInputString.md) | JSON array of row ids. | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceReplaceDraftTableRows

> BatchResponseHubDbTableRowV3 PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceReplaceDraftTableRows(ctx, tableIdOrName).BatchInputHubDbTableRowV3BatchUpdateRequest(batchInputHubDbTableRowV3BatchUpdateRequest).Execute()

Replace rows in batch in draft table



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
    batchInputHubDbTableRowV3BatchUpdateRequest := *openapiclient.NewBatchInputHubDbTableRowV3BatchUpdateRequest([]openapiclient.HubDbTableRowV3BatchUpdateRequest{*openapiclient.NewHubDbTableRowV3BatchUpdateRequest(map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "4099275931")}) // BatchInputHubDbTableRowV3BatchUpdateRequest | JSON array of row objects.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceReplaceDraftTableRows(context.Background(), tableIdOrName).BatchInputHubDbTableRowV3BatchUpdateRequest(batchInputHubDbTableRowV3BatchUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceReplaceDraftTableRows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceReplaceDraftTableRows`: BatchResponseHubDbTableRowV3
    fmt.Fprintf(os.Stdout, "Response from `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceReplaceDraftTableRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceReplaceDraftTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputHubDbTableRowV3BatchUpdateRequest** | [**BatchInputHubDbTableRowV3BatchUpdateRequest**](BatchInputHubDbTableRowV3BatchUpdateRequest.md) | JSON array of row objects. | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateUpdateDraftTableRows

> BatchResponseHubDbTableRowV3 PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateUpdateDraftTableRows(ctx, tableIdOrName).BatchInputHubDbTableRowV3BatchUpdateRequest(batchInputHubDbTableRowV3BatchUpdateRequest).Execute()

Update rows in batch in draft table



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
    batchInputHubDbTableRowV3BatchUpdateRequest := *openapiclient.NewBatchInputHubDbTableRowV3BatchUpdateRequest([]openapiclient.HubDbTableRowV3BatchUpdateRequest{*openapiclient.NewHubDbTableRowV3BatchUpdateRequest(map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "4099275931")}) // BatchInputHubDbTableRowV3BatchUpdateRequest | JSON array of row objects.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateUpdateDraftTableRows(context.Background(), tableIdOrName).BatchInputHubDbTableRowV3BatchUpdateRequest(batchInputHubDbTableRowV3BatchUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateUpdateDraftTableRows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateUpdateDraftTableRows`: BatchResponseHubDbTableRowV3
    fmt.Fprintf(os.Stdout, "Response from `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateUpdateDraftTableRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateUpdateDraftTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputHubDbTableRowV3BatchUpdateRequest** | [**BatchInputHubDbTableRowV3BatchUpdateRequest**](BatchInputHubDbTableRowV3BatchUpdateRequest.md) | JSON array of row objects. | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

