# \PublicObjectApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmV3ObjectsObjectTypeMergeMerge**](PublicObjectApi.md#PostCrmV3ObjectsObjectTypeMergeMerge) | **Post** /crm/v3/objects/{objectType}/merge | Merge two objects with same type



## PostCrmV3ObjectsObjectTypeMergeMerge

> SimplePublicObject PostCrmV3ObjectsObjectTypeMergeMerge(ctx, objectType).PublicMergeInput(publicMergeInput).Execute()

Merge two objects with same type

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
    publicMergeInput := *openapiclient.NewPublicMergeInput("PrimaryObjectId_example", "ObjectIdToMerge_example") // PublicMergeInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicObjectApi.PostCrmV3ObjectsObjectTypeMergeMerge(context.Background(), objectType).PublicMergeInput(publicMergeInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicObjectApi.PostCrmV3ObjectsObjectTypeMergeMerge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ObjectsObjectTypeMergeMerge`: SimplePublicObject
    fmt.Fprintf(os.Stdout, "Response from `PublicObjectApi.PostCrmV3ObjectsObjectTypeMergeMerge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsObjectTypeMergeMergeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicMergeInput** | [**PublicMergeInput**](PublicMergeInput.md) |  | 

### Return type

[**SimplePublicObject**](SimplePublicObject.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

