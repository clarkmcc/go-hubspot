# \PublicObjectSchemasApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmObjectSchemasV3SchemasObjectTypePurgePurge**](PublicObjectSchemasApi.md#DeleteCrmObjectSchemasV3SchemasObjectTypePurgePurge) | **Delete** /crm/v3/schemas/{objectType}/purge | 



## DeleteCrmObjectSchemasV3SchemasObjectTypePurgePurge

> DeleteCrmObjectSchemasV3SchemasObjectTypePurgePurge(ctx, objectType).Execute()



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
    objectType := "objectType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicObjectSchemasApi.DeleteCrmObjectSchemasV3SchemasObjectTypePurgePurge(context.Background(), objectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicObjectSchemasApi.DeleteCrmObjectSchemasV3SchemasObjectTypePurgePurge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmObjectSchemasV3SchemasObjectTypePurgePurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

