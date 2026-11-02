# \AdvancedAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmObjects202609ContactsGdprDeleteCrmObjects202609BetaObjectTypeGdprDelete**](AdvancedAPI.md#PostCrmObjects202609ContactsGdprDeleteCrmObjects202609BetaObjectTypeGdprDelete) | **Post** /crm/objects/2026-09/contacts/gdpr-delete | Permanently delete a contact (GDPR-compliant)
[**PostCrmObjects202609ContactsMergeCrmObjects202609BetaObjectTypeMerge**](AdvancedAPI.md#PostCrmObjects202609ContactsMergeCrmObjects202609BetaObjectTypeMerge) | **Post** /crm/objects/2026-09/contacts/merge | Merge two contacts



## PostCrmObjects202609ContactsGdprDeleteCrmObjects202609BetaObjectTypeGdprDelete

> PostCrmObjects202609ContactsGdprDeleteCrmObjects202609BetaObjectTypeGdprDelete(ctx).PublicGdprDeleteInput(publicGdprDeleteInput).Execute()

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
	r, err := apiClient.AdvancedAPI.PostCrmObjects202609ContactsGdprDeleteCrmObjects202609BetaObjectTypeGdprDelete(context.Background()).PublicGdprDeleteInput(publicGdprDeleteInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedAPI.PostCrmObjects202609ContactsGdprDeleteCrmObjects202609BetaObjectTypeGdprDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmObjects202609ContactsGdprDeleteCrmObjects202609BetaObjectTypeGdprDeleteRequest struct via the builder pattern


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


## PostCrmObjects202609ContactsMergeCrmObjects202609BetaObjectTypeMerge

> SimplePublicObject PostCrmObjects202609ContactsMergeCrmObjects202609BetaObjectTypeMerge(ctx).PublicMergeInput(publicMergeInput).Execute()

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
	resp, r, err := apiClient.AdvancedAPI.PostCrmObjects202609ContactsMergeCrmObjects202609BetaObjectTypeMerge(context.Background()).PublicMergeInput(publicMergeInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedAPI.PostCrmObjects202609ContactsMergeCrmObjects202609BetaObjectTypeMerge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmObjects202609ContactsMergeCrmObjects202609BetaObjectTypeMerge`: SimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `AdvancedAPI.PostCrmObjects202609ContactsMergeCrmObjects202609BetaObjectTypeMerge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmObjects202609ContactsMergeCrmObjects202609BetaObjectTypeMergeRequest struct via the builder pattern


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

