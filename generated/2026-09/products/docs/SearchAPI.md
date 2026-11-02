# \SearchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostProductsSearch**](SearchAPI.md#PostProductsSearch) | **Post** /crm/objects/2026-09/products/search | Search



## PostProductsSearch

> CollectionResponseWithTotalSimplePublicObject PostProductsSearch(ctx).PublicObjectSearchRequest(publicObjectSearchRequest).Execute()

Search



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
	publicObjectSearchRequest := *openapiclient.NewPublicObjectSearchRequest("After_example", []openapiclient.FilterGroup{*openapiclient.NewFilterGroup([]openapiclient.Filter{*openapiclient.NewFilter("Operator_example", "PropertyName_example")})}, int32(123), []string{"Properties_example"}, []string{"Sorts_example"}) // PublicObjectSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.PostProductsSearch(context.Background()).PublicObjectSearchRequest(publicObjectSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.PostProductsSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProductsSearch`: CollectionResponseWithTotalSimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.PostProductsSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProductsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicObjectSearchRequest** | [**PublicObjectSearchRequest**](PublicObjectSearchRequest.md) |  | 

### Return type

[**CollectionResponseWithTotalSimplePublicObject**](CollectionResponseWithTotalSimplePublicObject.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

