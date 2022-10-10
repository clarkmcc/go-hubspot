# \RowsBatchApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchCloneDraftTableRows**](RowsBatchApi.md#BatchCloneDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/clone | Clone rows in batch
[**BatchCreateDraftTableRows**](RowsBatchApi.md#BatchCreateDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/create | Create rows in batch
[**BatchPurgeDraftTableRows**](RowsBatchApi.md#BatchPurgeDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/purge | Permanently deletes rows
[**BatchReadDraftTableRows**](RowsBatchApi.md#BatchReadDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/read | Get a set of rows from draft table
[**BatchReadTableRows**](RowsBatchApi.md#BatchReadTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/batch/read | Get a set of rows
[**BatchReplaceDraftTableRows**](RowsBatchApi.md#BatchReplaceDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/replace | Replace rows in batch in draft table
[**BatchUpdateDraftTableRows**](RowsBatchApi.md#BatchUpdateDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/update | Update rows in batch in draft table



## BatchCloneDraftTableRows

> BatchResponseHubDbTableRowV3 BatchCloneDraftTableRows(ctx, tableIdOrName).BatchInputString(batchInputString).Execute()

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
    resp, r, err := apiClient.RowsBatchApi.BatchCloneDraftTableRows(context.Background(), tableIdOrName).BatchInputString(batchInputString).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchApi.BatchCloneDraftTableRows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchCloneDraftTableRows`: BatchResponseHubDbTableRowV3
    fmt.Fprintf(os.Stdout, "Response from `RowsBatchApi.BatchCloneDraftTableRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchCloneDraftTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of row ids | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchCreateDraftTableRows

> BatchResponseHubDbTableRowV3 BatchCreateDraftTableRows(ctx, tableIdOrName).BatchInputHubDbTableRowV3Request(batchInputHubDbTableRowV3Request).Execute()

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
    resp, r, err := apiClient.RowsBatchApi.BatchCreateDraftTableRows(context.Background(), tableIdOrName).BatchInputHubDbTableRowV3Request(batchInputHubDbTableRowV3Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchApi.BatchCreateDraftTableRows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchCreateDraftTableRows`: BatchResponseHubDbTableRowV3
    fmt.Fprintf(os.Stdout, "Response from `RowsBatchApi.BatchCreateDraftTableRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchCreateDraftTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputHubDbTableRowV3Request** | [**BatchInputHubDbTableRowV3Request**](BatchInputHubDbTableRowV3Request.md) | JSON array of row objects | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchPurgeDraftTableRows

> BatchPurgeDraftTableRows(ctx, tableIdOrName).BatchInputString(batchInputString).Execute()

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
    resp, r, err := apiClient.RowsBatchApi.BatchPurgeDraftTableRows(context.Background(), tableIdOrName).BatchInputString(batchInputString).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchApi.BatchPurgeDraftTableRows``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBatchPurgeDraftTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputString** | [**BatchInputString**](BatchInputString.md) | JSON array of row ids. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchReadDraftTableRows

> BatchResponseHubDbTableRowV3 BatchReadDraftTableRows(ctx, tableIdOrName).BatchInputString(batchInputString).Execute()

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
    resp, r, err := apiClient.RowsBatchApi.BatchReadDraftTableRows(context.Background(), tableIdOrName).BatchInputString(batchInputString).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchApi.BatchReadDraftTableRows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchReadDraftTableRows`: BatchResponseHubDbTableRowV3
    fmt.Fprintf(os.Stdout, "Response from `RowsBatchApi.BatchReadDraftTableRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchReadDraftTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputString** | [**BatchInputString**](BatchInputString.md) | JSON array of row ids. | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchReadTableRows

> BatchResponseHubDbTableRowV3 BatchReadTableRows(ctx, tableIdOrName).BatchInputString(batchInputString).Execute()

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
    resp, r, err := apiClient.RowsBatchApi.BatchReadTableRows(context.Background(), tableIdOrName).BatchInputString(batchInputString).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchApi.BatchReadTableRows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchReadTableRows`: BatchResponseHubDbTableRowV3
    fmt.Fprintf(os.Stdout, "Response from `RowsBatchApi.BatchReadTableRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchReadTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of row ids | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchReplaceDraftTableRows

> BatchResponseHubDbTableRowV3 BatchReplaceDraftTableRows(ctx, tableIdOrName).BatchInputHubDbTableRowV3Request(batchInputHubDbTableRowV3Request).Execute()

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
    batchInputHubDbTableRowV3Request := *openapiclient.NewBatchInputHubDbTableRowV3Request([]openapiclient.HubDbTableRowV3Request{*openapiclient.NewHubDbTableRowV3Request(map[string]map[string]interface{}{"key": map[string]interface{}(123)})}) // BatchInputHubDbTableRowV3Request | JSON array of row objects.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RowsBatchApi.BatchReplaceDraftTableRows(context.Background(), tableIdOrName).BatchInputHubDbTableRowV3Request(batchInputHubDbTableRowV3Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchApi.BatchReplaceDraftTableRows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchReplaceDraftTableRows`: BatchResponseHubDbTableRowV3
    fmt.Fprintf(os.Stdout, "Response from `RowsBatchApi.BatchReplaceDraftTableRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchReplaceDraftTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputHubDbTableRowV3Request** | [**BatchInputHubDbTableRowV3Request**](BatchInputHubDbTableRowV3Request.md) | JSON array of row objects. | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchUpdateDraftTableRows

> BatchResponseHubDbTableRowV3 BatchUpdateDraftTableRows(ctx, tableIdOrName).BatchInputJsonNode(batchInputJsonNode).Execute()

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
    batchInputJsonNode := *openapiclient.NewBatchInputJsonNode([]map[string]interface{}{map[string]interface{}(123)}) // BatchInputJsonNode | JSON array of row objects.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RowsBatchApi.BatchUpdateDraftTableRows(context.Background(), tableIdOrName).BatchInputJsonNode(batchInputJsonNode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchApi.BatchUpdateDraftTableRows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchUpdateDraftTableRows`: BatchResponseHubDbTableRowV3
    fmt.Fprintf(os.Stdout, "Response from `RowsBatchApi.BatchUpdateDraftTableRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchUpdateDraftTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputJsonNode** | [**BatchInputJsonNode**](BatchInputJsonNode.md) | JSON array of row objects. | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

