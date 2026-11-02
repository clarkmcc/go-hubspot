# \AdvancedAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostContactsGdprDelete**](AdvancedAPI.md#PostContactsGdprDelete) | **Post** /crm/objects/2026-09/contacts/gdpr-delete | Permanently delete a contact (GDPR-compliant)
[**PostContactsMerge**](AdvancedAPI.md#PostContactsMerge) | **Post** /crm/objects/2026-09/contacts/merge | Merge two contacts



## PostContactsGdprDelete

> PostContactsGdprDelete(ctx).PublicGdprDeleteInput(publicGdprDeleteInput).Execute()

Permanently delete a contact (GDPR-compliant)



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
	r, err := apiClient.AdvancedAPI.PostContactsGdprDelete(context.Background()).PublicGdprDeleteInput(publicGdprDeleteInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedAPI.PostContactsGdprDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostContactsGdprDeleteRequest struct via the builder pattern


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


## PostContactsMerge

> SimplePublicObject PostContactsMerge(ctx).PublicMergeInput(publicMergeInput).Execute()

Merge two contacts



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
	resp, r, err := apiClient.AdvancedAPI.PostContactsMerge(context.Background()).PublicMergeInput(publicMergeInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedAPI.PostContactsMerge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostContactsMerge`: SimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `AdvancedAPI.PostContactsMerge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostContactsMergeRequest struct via the builder pattern


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

