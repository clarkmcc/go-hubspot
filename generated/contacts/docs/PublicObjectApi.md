# \PublicObjectApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmV3ObjectsContactsMergeMerge**](PublicObjectApi.md#PostCrmV3ObjectsContactsMergeMerge) | **Post** /crm/v3/objects/contacts/merge | Merge two contacts with same type



## PostCrmV3ObjectsContactsMergeMerge

> SimplePublicObject PostCrmV3ObjectsContactsMergeMerge(ctx).PublicMergeInput(publicMergeInput).Execute()

Merge two contacts with same type

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
    resp, r, err := apiClient.PublicObjectApi.PostCrmV3ObjectsContactsMergeMerge(context.Background()).PublicMergeInput(publicMergeInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicObjectApi.PostCrmV3ObjectsContactsMergeMerge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ObjectsContactsMergeMerge`: SimplePublicObject
    fmt.Fprintf(os.Stdout, "Response from `PublicObjectApi.PostCrmV3ObjectsContactsMergeMerge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsContactsMergeMergeRequest struct via the builder pattern


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

