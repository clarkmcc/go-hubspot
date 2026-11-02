# \BatchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostSchemasBatchRead**](BatchAPI.md#PostSchemasBatchRead) | **Post** /crm-object-schemas/2026-09/schemas/batch/read | Retrieve multiple custom object schemas in a batch request.



## PostSchemasBatchRead

> CollectionResponseObjectSchemaNoPaging PostSchemasBatchRead(ctx).ObjectSchemaBatchReadRequest(objectSchemaBatchReadRequest).Execute()

Retrieve multiple custom object schemas in a batch request.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	objectSchemaBatchReadRequest := *openapiclient.NewObjectSchemaBatchReadRequest(false, false, false, []string{"Inputs_example"}) // ObjectSchemaBatchReadRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostSchemasBatchRead(context.Background()).ObjectSchemaBatchReadRequest(objectSchemaBatchReadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostSchemasBatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSchemasBatchRead`: CollectionResponseObjectSchemaNoPaging
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostSchemasBatchRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSchemasBatchReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectSchemaBatchReadRequest** | [**ObjectSchemaBatchReadRequest**](ObjectSchemaBatchReadRequest.md) |  | 

### Return type

[**CollectionResponseObjectSchemaNoPaging**](CollectionResponseObjectSchemaNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

