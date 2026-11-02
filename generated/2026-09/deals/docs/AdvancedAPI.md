# \AdvancedAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Post03GdprDelete**](AdvancedAPI.md#Post03GdprDelete) | **Post** /crm/objects/2026-09/0-3/gdpr-delete | GDPR Delete
[**Post03Merge**](AdvancedAPI.md#Post03Merge) | **Post** /crm/objects/2026-09/0-3/merge | Merge two deals with same type



## Post03GdprDelete

> Post03GdprDelete(ctx).PublicGdprDeleteInput(publicGdprDeleteInput).Execute()

GDPR Delete



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
	publicGdprDeleteInput := *openapiclient.NewPublicGdprDeleteInput("ObjectId_example") // PublicGdprDeleteInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdvancedAPI.Post03GdprDelete(context.Background()).PublicGdprDeleteInput(publicGdprDeleteInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedAPI.Post03GdprDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPost03GdprDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicGdprDeleteInput** | [**PublicGdprDeleteInput**](PublicGdprDeleteInput.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post03Merge

> SimplePublicObject Post03Merge(ctx).PublicMergeInput(publicMergeInput).Execute()

Merge two deals with same type



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
	publicMergeInput := *openapiclient.NewPublicMergeInput("ObjectIdToMerge_example", "PrimaryObjectId_example") // PublicMergeInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedAPI.Post03Merge(context.Background()).PublicMergeInput(publicMergeInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedAPI.Post03Merge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post03Merge`: SimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `AdvancedAPI.Post03Merge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPost03MergeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicMergeInput** | [**PublicMergeInput**](PublicMergeInput.md) |  | 

### Return type

[**SimplePublicObject**](SimplePublicObject.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

