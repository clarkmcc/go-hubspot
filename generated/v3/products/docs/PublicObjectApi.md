# \PublicObjectApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmV3ObjectsProductsMergeMerge**](PublicObjectApi.md#PostCrmV3ObjectsProductsMergeMerge) | **Post** /crm/v3/objects/products/merge | Merge two products with same type



## PostCrmV3ObjectsProductsMergeMerge

> SimplePublicObject PostCrmV3ObjectsProductsMergeMerge(ctx).PublicMergeInput(publicMergeInput).Execute()

Merge two products with same type

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
    publicMergeInput := *openapiclient.NewPublicMergeInput("PrimaryObjectId_example", "ObjectIdToMerge_example") // PublicMergeInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicObjectApi.PostCrmV3ObjectsProductsMergeMerge(context.Background()).PublicMergeInput(publicMergeInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicObjectApi.PostCrmV3ObjectsProductsMergeMerge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ObjectsProductsMergeMerge`: SimplePublicObject
    fmt.Fprintf(os.Stdout, "Response from `PublicObjectApi.PostCrmV3ObjectsProductsMergeMerge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsProductsMergeMergeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicMergeInput** | [**PublicMergeInput**](PublicMergeInput.md) |  | 

### Return type

[**SimplePublicObject**](SimplePublicObject.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

