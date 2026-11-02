# \AdvancedAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmObjects202609ObjectTypeGdprDeleteCrmObjects202609BetaObjectTypeGdprDelete**](AdvancedAPI.md#PostCrmObjects202609ObjectTypeGdprDeleteCrmObjects202609BetaObjectTypeGdprDelete) | **Post** /crm/objects/2026-09/{objectType}/gdpr-delete | GDPR Delete
[**PostCrmObjects202609ObjectTypeMergeCrmObjects202609BetaObjectTypeMerge**](AdvancedAPI.md#PostCrmObjects202609ObjectTypeMergeCrmObjects202609BetaObjectTypeMerge) | **Post** /crm/objects/2026-09/{objectType}/merge | Merge objects



## PostCrmObjects202609ObjectTypeGdprDeleteCrmObjects202609BetaObjectTypeGdprDelete

> PostCrmObjects202609ObjectTypeGdprDeleteCrmObjects202609BetaObjectTypeGdprDelete(ctx, objectType).PublicGdprDeleteInput(publicGdprDeleteInput).Execute()

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
	objectType := "objectType_example" // string | The type of CRM object to be deleted, such as 'contacts'.
	publicGdprDeleteInput := *openapiclient.NewPublicGdprDeleteInput("ObjectId_example") // PublicGdprDeleteInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdvancedAPI.PostCrmObjects202609ObjectTypeGdprDeleteCrmObjects202609BetaObjectTypeGdprDelete(context.Background(), objectType).PublicGdprDeleteInput(publicGdprDeleteInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedAPI.PostCrmObjects202609ObjectTypeGdprDeleteCrmObjects202609BetaObjectTypeGdprDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of CRM object to be deleted, such as &#39;contacts&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmObjects202609ObjectTypeGdprDeleteCrmObjects202609BetaObjectTypeGdprDeleteRequest struct via the builder pattern


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


## PostCrmObjects202609ObjectTypeMergeCrmObjects202609BetaObjectTypeMerge

> SimplePublicObject PostCrmObjects202609ObjectTypeMergeCrmObjects202609BetaObjectTypeMerge(ctx, objectType).PublicMergeInput(publicMergeInput).Execute()

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
	objectType := "objectType_example" // string | The type of the CRM object to be merged, such as contacts, companies, or deals.
	publicMergeInput := *openapiclient.NewPublicMergeInput("ObjectIdToMerge_example", "PrimaryObjectId_example") // PublicMergeInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedAPI.PostCrmObjects202609ObjectTypeMergeCrmObjects202609BetaObjectTypeMerge(context.Background(), objectType).PublicMergeInput(publicMergeInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedAPI.PostCrmObjects202609ObjectTypeMergeCrmObjects202609BetaObjectTypeMerge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmObjects202609ObjectTypeMergeCrmObjects202609BetaObjectTypeMerge`: SimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `AdvancedAPI.PostCrmObjects202609ObjectTypeMergeCrmObjects202609BetaObjectTypeMerge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of the CRM object to be merged, such as contacts, companies, or deals. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmObjects202609ObjectTypeMergeCrmObjects202609BetaObjectTypeMergeRequest struct via the builder pattern


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

