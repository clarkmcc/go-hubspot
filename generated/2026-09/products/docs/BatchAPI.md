# \BatchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostProductsBatchArchive**](BatchAPI.md#PostProductsBatchArchive) | **Post** /crm/objects/2026-09/products/batch/archive | Batch archive objects
[**PostProductsBatchCreate**](BatchAPI.md#PostProductsBatchCreate) | **Post** /crm/objects/2026-09/products/batch/create | Batch create objects
[**PostProductsBatchRead**](BatchAPI.md#PostProductsBatchRead) | **Post** /crm/objects/2026-09/products/batch/read | Batch read objects
[**PostProductsBatchUpdate**](BatchAPI.md#PostProductsBatchUpdate) | **Post** /crm/objects/2026-09/products/batch/update | Update batch objects
[**PostProductsBatchUpsert**](BatchAPI.md#PostProductsBatchUpsert) | **Post** /crm/objects/2026-09/products/batch/upsert | Upsert batch objects
[**PutDefault**](BatchAPI.md#PutDefault) | **Put** /crm/objects/2026-09/{fromObjectType}/{fromObjectId}/associations/default/{toObjectType}/{toObjectId} | Create Association



## PostProductsBatchArchive

> PostProductsBatchArchive(ctx).BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId).Execute()

Batch archive objects



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
	r, err := apiClient.BatchAPI.PostProductsBatchArchive(context.Background()).BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostProductsBatchArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProductsBatchArchiveRequest struct via the builder pattern


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


## PostProductsBatchCreate

> BatchResponseSimplePublicObject PostProductsBatchCreate(ctx).BatchInputSimplePublicObjectBatchInputForCreate(batchInputSimplePublicObjectBatchInputForCreate).Execute()

Batch create objects



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
	resp, r, err := apiClient.BatchAPI.PostProductsBatchCreate(context.Background()).BatchInputSimplePublicObjectBatchInputForCreate(batchInputSimplePublicObjectBatchInputForCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostProductsBatchCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProductsBatchCreate`: BatchResponseSimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostProductsBatchCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProductsBatchCreateRequest struct via the builder pattern


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


## PostProductsBatchRead

> BatchResponseSimplePublicObject PostProductsBatchRead(ctx).BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId).Archived(archived).Execute()

Batch read objects



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
	resp, r, err := apiClient.BatchAPI.PostProductsBatchRead(context.Background()).BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostProductsBatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProductsBatchRead`: BatchResponseSimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostProductsBatchRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProductsBatchReadRequest struct via the builder pattern


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


## PostProductsBatchUpdate

> BatchResponseSimplePublicObject PostProductsBatchUpdate(ctx).BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput).Execute()

Update batch objects



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
	resp, r, err := apiClient.BatchAPI.PostProductsBatchUpdate(context.Background()).BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostProductsBatchUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProductsBatchUpdate`: BatchResponseSimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostProductsBatchUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProductsBatchUpdateRequest struct via the builder pattern


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


## PostProductsBatchUpsert

> BatchResponseSimplePublicUpsertObject PostProductsBatchUpsert(ctx).BatchInputSimplePublicObjectBatchInputUpsert(batchInputSimplePublicObjectBatchInputUpsert).Execute()

Upsert batch objects



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
	resp, r, err := apiClient.BatchAPI.PostProductsBatchUpsert(context.Background()).BatchInputSimplePublicObjectBatchInputUpsert(batchInputSimplePublicObjectBatchInputUpsert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostProductsBatchUpsert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProductsBatchUpsert`: BatchResponseSimplePublicUpsertObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostProductsBatchUpsert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProductsBatchUpsertRequest struct via the builder pattern


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

