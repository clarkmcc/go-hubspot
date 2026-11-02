# \SearchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmObjects202609ObjectTypeSearchCrmObjects202609BetaObjectTypeSearch**](SearchAPI.md#PostCrmObjects202609ObjectTypeSearchCrmObjects202609BetaObjectTypeSearch) | **Post** /crm/objects/2026-09/{objectType}/search | Search objects



## PostCrmObjects202609ObjectTypeSearchCrmObjects202609BetaObjectTypeSearch

> CollectionResponseWithTotalSimplePublicObject PostCrmObjects202609ObjectTypeSearchCrmObjects202609BetaObjectTypeSearch(ctx, objectType).PublicObjectSearchRequest(publicObjectSearchRequest).Execute()

Search objects



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
	objectType := "objectType_example" // string | The type of CRM object to search, such as contacts, companies, or deals.
	publicObjectSearchRequest := *openapiclient.NewPublicObjectSearchRequest("After_example", []openapiclient.FilterGroup{*openapiclient.NewFilterGroup([]openapiclient.Filter{*openapiclient.NewFilter("Operator_example", "PropertyName_example")})}, int32(123), []string{"Properties_example"}, []string{"Sorts_example"}) // PublicObjectSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.PostCrmObjects202609ObjectTypeSearchCrmObjects202609BetaObjectTypeSearch(context.Background(), objectType).PublicObjectSearchRequest(publicObjectSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.PostCrmObjects202609ObjectTypeSearchCrmObjects202609BetaObjectTypeSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmObjects202609ObjectTypeSearchCrmObjects202609BetaObjectTypeSearch`: CollectionResponseWithTotalSimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.PostCrmObjects202609ObjectTypeSearchCrmObjects202609BetaObjectTypeSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of CRM object to search, such as contacts, companies, or deals. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmObjects202609ObjectTypeSearchCrmObjects202609BetaObjectTypeSearchRequest struct via the builder pattern


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

