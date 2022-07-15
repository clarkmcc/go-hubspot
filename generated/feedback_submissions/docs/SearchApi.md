# \SearchApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmV3ObjectsFeedbackSubmissionsSearchDoSearch**](SearchApi.md#PostCrmV3ObjectsFeedbackSubmissionsSearchDoSearch) | **Post** /crm/v3/objects/feedback_submissions/search | 



## PostCrmV3ObjectsFeedbackSubmissionsSearchDoSearch

> CollectionResponseWithTotalSimplePublicObjectForwardPaging PostCrmV3ObjectsFeedbackSubmissionsSearchDoSearch(ctx).PublicObjectSearchRequest(publicObjectSearchRequest).Execute()



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
    publicObjectSearchRequest := *openapiclient.NewPublicObjectSearchRequest([]openapiclient.FilterGroup{*openapiclient.NewFilterGroup([]openapiclient.Filter{*openapiclient.NewFilter("PropertyName_example", "Operator_example")})}, []string{"Sorts_example"}, []string{"Properties_example"}, int32(123), int32(123)) // PublicObjectSearchRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.PostCrmV3ObjectsFeedbackSubmissionsSearchDoSearch(context.Background()).PublicObjectSearchRequest(publicObjectSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.PostCrmV3ObjectsFeedbackSubmissionsSearchDoSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ObjectsFeedbackSubmissionsSearchDoSearch`: CollectionResponseWithTotalSimplePublicObjectForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.PostCrmV3ObjectsFeedbackSubmissionsSearchDoSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsFeedbackSubmissionsSearchDoSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicObjectSearchRequest** | [**PublicObjectSearchRequest**](PublicObjectSearchRequest.md) |  | 

### Return type

[**CollectionResponseWithTotalSimplePublicObjectForwardPaging**](CollectionResponseWithTotalSimplePublicObjectForwardPaging.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

