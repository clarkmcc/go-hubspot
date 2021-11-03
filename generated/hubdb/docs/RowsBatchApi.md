# \RowsBatchApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCmsV3HubdbTablesTableIdOrNameRowsBatchReadBatchReadTableRows**](RowsBatchApi.md#PostCmsV3HubdbTablesTableIdOrNameRowsBatchReadBatchReadTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/batch/read | Get a set of rows
[**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneBatchCloneDraftTableRows**](RowsBatchApi.md#PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneBatchCloneDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/clone | Clone rows in batch
[**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateBatchCreateDraftTableRows**](RowsBatchApi.md#PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateBatchCreateDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/create | Create rows in batch
[**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgeBatchPurgeDraftTableRows**](RowsBatchApi.md#PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgeBatchPurgeDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/purge | Permanently deletes rows
[**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadBatchReadDraftTableRows**](RowsBatchApi.md#PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadBatchReadDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/read | Get a set of rows from draft table
[**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceBatchReplaceDraftTableRows**](RowsBatchApi.md#PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceBatchReplaceDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/replace | Replace rows in batch in draft table
[**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateBatchUpdateDraftTableRows**](RowsBatchApi.md#PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateBatchUpdateDraftTableRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/update | Update rows in batch in draft table



## PostCmsV3HubdbTablesTableIdOrNameRowsBatchReadBatchReadTableRows

> map[string]interface{} PostCmsV3HubdbTablesTableIdOrNameRowsBatchReadBatchReadTableRows(ctx, tableIdOrName).BatchInputString(batchInputString).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsBatchReadBatchReadTableRows(context.Background(), tableIdOrName).BatchInputString(batchInputString).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsBatchReadBatchReadTableRows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3HubdbTablesTableIdOrNameRowsBatchReadBatchReadTableRows`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsBatchReadBatchReadTableRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsBatchReadBatchReadTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of row ids | 

### Return type

**map[string]interface{}**

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneBatchCloneDraftTableRows

> map[string]interface{} PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneBatchCloneDraftTableRows(ctx, tableIdOrName).BatchInputString(batchInputString).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneBatchCloneDraftTableRows(context.Background(), tableIdOrName).BatchInputString(batchInputString).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneBatchCloneDraftTableRows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneBatchCloneDraftTableRows`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneBatchCloneDraftTableRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneBatchCloneDraftTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of row ids | 

### Return type

**map[string]interface{}**

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateBatchCreateDraftTableRows

> map[string]interface{} PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateBatchCreateDraftTableRows(ctx, tableIdOrName).BatchInputHubDbTableRowV3Request(batchInputHubDbTableRowV3Request).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateBatchCreateDraftTableRows(context.Background(), tableIdOrName).BatchInputHubDbTableRowV3Request(batchInputHubDbTableRowV3Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateBatchCreateDraftTableRows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateBatchCreateDraftTableRows`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateBatchCreateDraftTableRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateBatchCreateDraftTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputHubDbTableRowV3Request** | [**BatchInputHubDbTableRowV3Request**](BatchInputHubDbTableRowV3Request.md) | JSON array of row objects | 

### Return type

**map[string]interface{}**

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgeBatchPurgeDraftTableRows

> PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgeBatchPurgeDraftTableRows(ctx, tableIdOrName).BatchInputString(batchInputString).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgeBatchPurgeDraftTableRows(context.Background(), tableIdOrName).BatchInputString(batchInputString).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgeBatchPurgeDraftTableRows``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgeBatchPurgeDraftTableRowsRequest struct via the builder pattern


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


## PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadBatchReadDraftTableRows

> map[string]interface{} PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadBatchReadDraftTableRows(ctx, tableIdOrName).BatchInputString(batchInputString).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadBatchReadDraftTableRows(context.Background(), tableIdOrName).BatchInputString(batchInputString).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadBatchReadDraftTableRows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadBatchReadDraftTableRows`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadBatchReadDraftTableRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadBatchReadDraftTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputString** | [**BatchInputString**](BatchInputString.md) | JSON array of row ids. | 

### Return type

**map[string]interface{}**

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceBatchReplaceDraftTableRows

> map[string]interface{} PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceBatchReplaceDraftTableRows(ctx, tableIdOrName).BatchInputHubDbTableRowV3Request(batchInputHubDbTableRowV3Request).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceBatchReplaceDraftTableRows(context.Background(), tableIdOrName).BatchInputHubDbTableRowV3Request(batchInputHubDbTableRowV3Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceBatchReplaceDraftTableRows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceBatchReplaceDraftTableRows`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceBatchReplaceDraftTableRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceBatchReplaceDraftTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputHubDbTableRowV3Request** | [**BatchInputHubDbTableRowV3Request**](BatchInputHubDbTableRowV3Request.md) | JSON array of row objects. | 

### Return type

**map[string]interface{}**

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateBatchUpdateDraftTableRows

> map[string]interface{} PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateBatchUpdateDraftTableRows(ctx, tableIdOrName).BatchInputJsonNode(batchInputJsonNode).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateBatchUpdateDraftTableRows(context.Background(), tableIdOrName).BatchInputJsonNode(batchInputJsonNode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateBatchUpdateDraftTableRows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateBatchUpdateDraftTableRows`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RowsBatchApi.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateBatchUpdateDraftTableRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateBatchUpdateDraftTableRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputJsonNode** | [**BatchInputJsonNode**](BatchInputJsonNode.md) | JSON array of row objects. | 

### Return type

**map[string]interface{}**

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

