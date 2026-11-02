# \AdvancedAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostProductsGdprDelete**](AdvancedAPI.md#PostProductsGdprDelete) | **Post** /crm/objects/2026-09/products/gdpr-delete | GDPR Delete
[**PostProductsMerge**](AdvancedAPI.md#PostProductsMerge) | **Post** /crm/objects/2026-09/products/merge | Merge objects



## PostProductsGdprDelete

> PostProductsGdprDelete(ctx).PublicGdprDeleteInput(publicGdprDeleteInput).Execute()

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
	r, err := apiClient.AdvancedAPI.PostProductsGdprDelete(context.Background()).PublicGdprDeleteInput(publicGdprDeleteInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedAPI.PostProductsGdprDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProductsGdprDeleteRequest struct via the builder pattern


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


## PostProductsMerge

> SimplePublicObject PostProductsMerge(ctx).PublicMergeInput(publicMergeInput).Execute()

Merge objects



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
	resp, r, err := apiClient.AdvancedAPI.PostProductsMerge(context.Background()).PublicMergeInput(publicMergeInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedAPI.PostProductsMerge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProductsMerge`: SimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `AdvancedAPI.PostProductsMerge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProductsMergeRequest struct via the builder pattern


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

