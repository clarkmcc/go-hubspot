# \BasicAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOwners**](BasicAPI.md#GetOwners) | **Get** /crm/owners/2026-09 | List owners
[**GetOwnersByOwnerId**](BasicAPI.md#GetOwnersByOwnerId) | **Get** /crm/owners/2026-09/{ownerId} | Retrieve a specific owner by ID



## GetOwners

> CollectionResponsePublicOwnerForwardPaging GetOwners(ctx).After(after).Archived(archived).Email(email).Limit(limit).Execute()

List owners



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
	after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)
	email := "email_example" // string | Filter by email address (optional) (optional)
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetOwners(context.Background()).After(after).Archived(archived).Email(email).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetOwners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOwners`: CollectionResponsePublicOwnerForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetOwners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]
 **email** | **string** | Filter by email address (optional) | 
 **limit** | **int32** | The maximum number of results to display per page. | [default to 100]

### Return type

[**CollectionResponsePublicOwnerForwardPaging**](CollectionResponsePublicOwnerForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOwnersByOwnerId

> PublicOwner GetOwnersByOwnerId(ctx, ownerId).Archived(archived).IdProperty(idProperty).Execute()

Retrieve a specific owner by ID



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
	ownerId := int32(56) // int32 | The unique identifier of the owner to retrieve.
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)
	idProperty := "idProperty_example" // string | Specifies which property to use for identifying the owner. Valid values are 'id' or 'userId'. Defaults to 'id'. (optional) (default to "id")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetOwnersByOwnerId(context.Background(), ownerId).Archived(archived).IdProperty(idProperty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetOwnersByOwnerId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOwnersByOwnerId`: PublicOwner
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetOwnersByOwnerId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerId** | **int32** | The unique identifier of the owner to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnersByOwnerIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]
 **idProperty** | **string** | Specifies which property to use for identifying the owner. Valid values are &#39;id&#39; or &#39;userId&#39;. Defaults to &#39;id&#39;. | [default to &quot;id&quot;]

### Return type

[**PublicOwner**](PublicOwner.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

