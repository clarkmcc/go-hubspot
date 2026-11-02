# \BatchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchArchive**](BatchAPI.md#BatchArchive) | **Post** /crm/objects/2026-09/contacts/batch/archive | Archive a batch of contacts
[**BatchCreate**](BatchAPI.md#BatchCreate) | **Post** /crm/objects/2026-09/contacts/batch/create | Create a batch of contacts
[**BatchRead**](BatchAPI.md#BatchRead) | **Post** /crm/objects/2026-09/contacts/batch/read | Retrieve a batch of contacts
[**BatchUpdate**](BatchAPI.md#BatchUpdate) | **Post** /crm/objects/2026-09/contacts/batch/update | Update a batch of contacts
[**PostContactsBatchUpsert**](BatchAPI.md#PostContactsBatchUpsert) | **Post** /crm/objects/2026-09/contacts/batch/upsert | Create or update a batch of contacts
[**PutDefault**](BatchAPI.md#PutDefault) | **Put** /crm/objects/2026-09/{fromObjectType}/{fromObjectId}/associations/default/{toObjectType}/{toObjectId} | Create Association



## BatchArchive

> BatchArchive(ctx).BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId).Execute()

Archive a batch of contacts



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
	r, err := apiClient.BatchAPI.BatchArchive(context.Background()).BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchArchiveRequest struct via the builder pattern


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


## BatchCreate

> BatchResponseSimplePublicObject BatchCreate(ctx).BatchInputSimplePublicObjectBatchInputForCreate(batchInputSimplePublicObjectBatchInputForCreate).Execute()

Create a batch of contacts



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
	resp, r, err := apiClient.BatchAPI.BatchCreate(context.Background()).BatchInputSimplePublicObjectBatchInputForCreate(batchInputSimplePublicObjectBatchInputForCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchCreate`: BatchResponseSimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchCreateRequest struct via the builder pattern


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


## BatchRead

> BatchResponseSimplePublicObject BatchRead(ctx).BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId).Archived(archived).Execute()

Retrieve a batch of contacts



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
	resp, r, err := apiClient.BatchAPI.BatchRead(context.Background()).BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchRead`: BatchResponseSimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchReadRequest struct via the builder pattern


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


## BatchUpdate

> BatchResponseSimplePublicObject BatchUpdate(ctx).BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput).Execute()

Update a batch of contacts



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
	resp, r, err := apiClient.BatchAPI.BatchUpdate(context.Background()).BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchUpdate`: BatchResponseSimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchUpdateRequest struct via the builder pattern


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


## PostContactsBatchUpsert

> BatchResponseSimplePublicUpsertObject PostContactsBatchUpsert(ctx).BatchInputSimplePublicObjectBatchInputUpsert(batchInputSimplePublicObjectBatchInputUpsert).Execute()

Create or update a batch of contacts



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
	resp, r, err := apiClient.BatchAPI.PostContactsBatchUpsert(context.Background()).BatchInputSimplePublicObjectBatchInputUpsert(batchInputSimplePublicObjectBatchInputUpsert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostContactsBatchUpsert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostContactsBatchUpsert`: BatchResponseSimplePublicUpsertObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostContactsBatchUpsert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostContactsBatchUpsertRequest struct via the builder pattern


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
	fromObjectId := "fromObjectId_example" // string | The unique identifier of the source object for the association.
	fromObjectType := "fromObjectType_example" // string | The type of the source object for the association.
	toObjectId := "toObjectId_example" // string | The unique identifier of the target object for the association.
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
**fromObjectId** | **string** | The unique identifier of the source object for the association. | 
**fromObjectType** | **string** | The type of the source object for the association. | 
**toObjectId** | **string** | The unique identifier of the target object for the association. | 
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

