# \SearchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Post03Search**](SearchAPI.md#Post03Search) | **Post** /crm/objects/2026-09/0-3/search | Search for deals using various filters and criteria to retrieve specific records.



## Post03Search

> CollectionResponseWithTotalSimplePublicObject Post03Search(ctx).PublicObjectSearchRequest(publicObjectSearchRequest).Execute()

Search for deals using various filters and criteria to retrieve specific records.



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
	resp, r, err := apiClient.SearchAPI.Post03Search(context.Background()).PublicObjectSearchRequest(publicObjectSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.Post03Search``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post03Search`: CollectionResponseWithTotalSimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.Post03Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPost03SearchRequest struct via the builder pattern


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

