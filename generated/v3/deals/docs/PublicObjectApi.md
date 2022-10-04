# \PublicObjectApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmV3ObjectsDealsMergeMerge**](PublicObjectApi.md#PostCrmV3ObjectsDealsMergeMerge) | **Post** /crm/v3/objects/deals/merge | Merge two deals with same type



## PostCrmV3ObjectsDealsMergeMerge

> SimplePublicObject PostCrmV3ObjectsDealsMergeMerge(ctx).PublicMergeInput(publicMergeInput).Execute()

Merge two deals with same type

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
    resp, r, err := apiClient.PublicObjectApi.PostCrmV3ObjectsDealsMergeMerge(context.Background()).PublicMergeInput(publicMergeInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicObjectApi.PostCrmV3ObjectsDealsMergeMerge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ObjectsDealsMergeMerge`: SimplePublicObject
    fmt.Fprintf(os.Stdout, "Response from `PublicObjectApi.PostCrmV3ObjectsDealsMergeMerge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsDealsMergeMergeRequest struct via the builder pattern


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

