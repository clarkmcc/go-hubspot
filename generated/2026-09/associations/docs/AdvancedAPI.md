# \AdvancedAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostGdprDelete**](AdvancedAPI.md#PostGdprDelete) | **Post** /crm/objects/2026-09/{objectType}/gdpr-delete | GDPR DELETE
[**PostMerge**](AdvancedAPI.md#PostMerge) | **Post** /crm/objects/2026-09/{objectType}/merge | Merge two CRM objects of the specified type, combining their properties and associations into a single object.



## PostGdprDelete

> PostGdprDelete(ctx, objectType).PublicGdprDeleteInput(publicGdprDeleteInput).Execute()

GDPR DELETE



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
	r, err := apiClient.AdvancedAPI.PostGdprDelete(context.Background(), objectType).PublicGdprDeleteInput(publicGdprDeleteInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedAPI.PostGdprDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostGdprDeleteRequest struct via the builder pattern


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


## PostMerge

> SimplePublicObject PostMerge(ctx, objectType).PublicMergeInput(publicMergeInput).Execute()

Merge two CRM objects of the specified type, combining their properties and associations into a single object.



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
	objectType := "objectType_example" // string | The type of the objects to merge, such as contacts, companies, or deals.
	publicMergeInput := *openapiclient.NewPublicMergeInput("ObjectIdToMerge_example", "PrimaryObjectId_example") // PublicMergeInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedAPI.PostMerge(context.Background(), objectType).PublicMergeInput(publicMergeInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedAPI.PostMerge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMerge`: SimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `AdvancedAPI.PostMerge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of the objects to merge, such as contacts, companies, or deals. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMergeRequest struct via the builder pattern


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

