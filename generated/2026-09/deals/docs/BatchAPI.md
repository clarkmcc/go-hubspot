# \BatchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Post03BatchArchive**](BatchAPI.md#Post03BatchArchive) | **Post** /crm/objects/2026-09/0-3/batch/archive | Archive a batch of deals by ID
[**Post03BatchCreate**](BatchAPI.md#Post03BatchCreate) | **Post** /crm/objects/2026-09/0-3/batch/create | Create a batch of deals
[**Post03BatchRead**](BatchAPI.md#Post03BatchRead) | **Post** /crm/objects/2026-09/0-3/batch/read | Read a batch of deals by internal ID, or unique property values
[**Post03BatchUpdate**](BatchAPI.md#Post03BatchUpdate) | **Post** /crm/objects/2026-09/0-3/batch/update | Update a batch of deals by internal ID, or unique property values
[**Post03BatchUpsert**](BatchAPI.md#Post03BatchUpsert) | **Post** /crm/objects/2026-09/0-3/batch/upsert | Create or update a batch of deals by unique property values
[**PutDefault**](BatchAPI.md#PutDefault) | **Put** /crm/objects/2026-09/{fromObjectType}/{fromObjectId}/associations/default/{toObjectType}/{toObjectId} | Create Association



## Post03BatchArchive

> Post03BatchArchive(ctx).BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId).Execute()

Archive a batch of deals by ID



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
	batchInputSimplePublicObjectId := *openapiclient.NewBatchInputSimplePublicObjectId([]openapiclient.SimplePublicObjectId{*openapiclient.NewSimplePublicObjectId("430001")}) // BatchInputSimplePublicObjectId | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchAPI.Post03BatchArchive(context.Background()).BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.Post03BatchArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPost03BatchArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputSimplePublicObjectId** | [**BatchInputSimplePublicObjectId**](BatchInputSimplePublicObjectId.md) |  | 

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


## Post03BatchCreate

> BatchResponseSimplePublicObject Post03BatchCreate(ctx).BatchInputSimplePublicObjectBatchInputForCreate(batchInputSimplePublicObjectBatchInputForCreate).Execute()

Create a batch of deals



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
	batchInputSimplePublicObjectBatchInputForCreate := *openapiclient.NewBatchInputSimplePublicObjectBatchInputForCreate([]openapiclient.SimplePublicObjectBatchInputForCreate{*openapiclient.NewSimplePublicObjectBatchInputForCreate([]openapiclient.PublicAssociationsForObject{*openapiclient.NewPublicAssociationsForObject(*openapiclient.NewPublicObjectId("Id_example"), []openapiclient.AssociationSpec{*openapiclient.NewAssociationSpec("AssociationCategory_example", int32(123))})}, map[string]string{"key": "Inner_example"})}) // BatchInputSimplePublicObjectBatchInputForCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.Post03BatchCreate(context.Background()).BatchInputSimplePublicObjectBatchInputForCreate(batchInputSimplePublicObjectBatchInputForCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.Post03BatchCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post03BatchCreate`: BatchResponseSimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.Post03BatchCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPost03BatchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputSimplePublicObjectBatchInputForCreate** | [**BatchInputSimplePublicObjectBatchInputForCreate**](BatchInputSimplePublicObjectBatchInputForCreate.md) |  | 

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post03BatchRead

> BatchResponseSimplePublicObject Post03BatchRead(ctx).BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId).Archived(archived).Execute()

Read a batch of deals by internal ID, or unique property values



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
	batchReadInputSimplePublicObjectId := *openapiclient.NewBatchReadInputSimplePublicObjectId([]openapiclient.SimplePublicObjectId{*openapiclient.NewSimplePublicObjectId("430001")}, []string{"Properties_example"}, []string{"PropertiesWithHistory_example"}) // BatchReadInputSimplePublicObjectId | 
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.Post03BatchRead(context.Background()).BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.Post03BatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post03BatchRead`: BatchResponseSimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.Post03BatchRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPost03BatchReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchReadInputSimplePublicObjectId** | [**BatchReadInputSimplePublicObjectId**](BatchReadInputSimplePublicObjectId.md) |  | 
 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post03BatchUpdate

> BatchResponseSimplePublicObject Post03BatchUpdate(ctx).BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput).Execute()

Update a batch of deals by internal ID, or unique property values



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
	batchInputSimplePublicObjectBatchInput := *openapiclient.NewBatchInputSimplePublicObjectBatchInput([]openapiclient.SimplePublicObjectBatchInput{*openapiclient.NewSimplePublicObjectBatchInput("Id_example", map[string]string{"key": "Inner_example"})}) // BatchInputSimplePublicObjectBatchInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.Post03BatchUpdate(context.Background()).BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.Post03BatchUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post03BatchUpdate`: BatchResponseSimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.Post03BatchUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPost03BatchUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputSimplePublicObjectBatchInput** | [**BatchInputSimplePublicObjectBatchInput**](BatchInputSimplePublicObjectBatchInput.md) |  | 

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post03BatchUpsert

> BatchResponseSimplePublicUpsertObject Post03BatchUpsert(ctx).BatchInputSimplePublicObjectBatchInputUpsert(batchInputSimplePublicObjectBatchInputUpsert).Execute()

Create or update a batch of deals by unique property values



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
	batchInputSimplePublicObjectBatchInputUpsert := *openapiclient.NewBatchInputSimplePublicObjectBatchInputUpsert([]openapiclient.SimplePublicObjectBatchInputUpsert{*openapiclient.NewSimplePublicObjectBatchInputUpsert("Id_example", map[string]string{"key": "Inner_example"})}) // BatchInputSimplePublicObjectBatchInputUpsert | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.Post03BatchUpsert(context.Background()).BatchInputSimplePublicObjectBatchInputUpsert(batchInputSimplePublicObjectBatchInputUpsert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.Post03BatchUpsert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post03BatchUpsert`: BatchResponseSimplePublicUpsertObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.Post03BatchUpsert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPost03BatchUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputSimplePublicObjectBatchInputUpsert** | [**BatchInputSimplePublicObjectBatchInputUpsert**](BatchInputSimplePublicObjectBatchInputUpsert.md) |  | 

### Return type

[**BatchResponseSimplePublicUpsertObject**](BatchResponseSimplePublicUpsertObject.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDefault

> BatchResponsePublicDefaultAssociation PutDefault(ctx, fromObjectId, fromObjectType, toObjectId, toObjectType).Execute()

Create Association



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
	fromObjectId := "fromObjectId_example" // string | The unique identifier of the source object.
	fromObjectType := "fromObjectType_example" // string | The type of the source object for the association.
	toObjectId := "toObjectId_example" // string | The unique identifier of the target object.
	toObjectType := "toObjectType_example" // string | The type of the target object for the association.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PutDefault(context.Background(), fromObjectId, fromObjectType, toObjectId, toObjectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PutDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDefault`: BatchResponsePublicDefaultAssociation
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PutDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectId** | **string** | The unique identifier of the source object. | 
**fromObjectType** | **string** | The type of the source object for the association. | 
**toObjectId** | **string** | The unique identifier of the target object. | 
**toObjectType** | **string** | The type of the target object for the association. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**BatchResponsePublicDefaultAssociation**](BatchResponsePublicDefaultAssociation.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

