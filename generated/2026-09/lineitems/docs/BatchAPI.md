# \BatchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostLineItemsBatchArchive**](BatchAPI.md#PostLineItemsBatchArchive) | **Post** /crm/objects/2026-09/line_items/batch/archive | Batch archive objects
[**PostLineItemsBatchCreate**](BatchAPI.md#PostLineItemsBatchCreate) | **Post** /crm/objects/2026-09/line_items/batch/create | Batch create objects
[**PostLineItemsBatchRead**](BatchAPI.md#PostLineItemsBatchRead) | **Post** /crm/objects/2026-09/line_items/batch/read | Read batch objects
[**PostLineItemsBatchUpdate**](BatchAPI.md#PostLineItemsBatchUpdate) | **Post** /crm/objects/2026-09/line_items/batch/update | Update batch objects
[**PostLineItemsBatchUpsert**](BatchAPI.md#PostLineItemsBatchUpsert) | **Post** /crm/objects/2026-09/line_items/batch/upsert | Upsert batch objects
[**PutDefault**](BatchAPI.md#PutDefault) | **Put** /crm/objects/2026-09/{fromObjectType}/{fromObjectId}/associations/default/{toObjectType}/{toObjectId} | Create Association



## PostLineItemsBatchArchive

> PostLineItemsBatchArchive(ctx).BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId).Execute()

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
	r, err := apiClient.BatchAPI.PostLineItemsBatchArchive(context.Background()).BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostLineItemsBatchArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLineItemsBatchArchiveRequest struct via the builder pattern


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


## PostLineItemsBatchCreate

> BatchResponseSimplePublicObject PostLineItemsBatchCreate(ctx).BatchInputSimplePublicObjectBatchInputForCreate(batchInputSimplePublicObjectBatchInputForCreate).Execute()

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
	resp, r, err := apiClient.BatchAPI.PostLineItemsBatchCreate(context.Background()).BatchInputSimplePublicObjectBatchInputForCreate(batchInputSimplePublicObjectBatchInputForCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostLineItemsBatchCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLineItemsBatchCreate`: BatchResponseSimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostLineItemsBatchCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLineItemsBatchCreateRequest struct via the builder pattern


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


## PostLineItemsBatchRead

> BatchResponseSimplePublicObject PostLineItemsBatchRead(ctx).BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId).Archived(archived).Execute()

Read batch objects



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
	resp, r, err := apiClient.BatchAPI.PostLineItemsBatchRead(context.Background()).BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostLineItemsBatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLineItemsBatchRead`: BatchResponseSimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostLineItemsBatchRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLineItemsBatchReadRequest struct via the builder pattern


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


## PostLineItemsBatchUpdate

> BatchResponseSimplePublicObject PostLineItemsBatchUpdate(ctx).BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput).Execute()

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
	resp, r, err := apiClient.BatchAPI.PostLineItemsBatchUpdate(context.Background()).BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostLineItemsBatchUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLineItemsBatchUpdate`: BatchResponseSimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostLineItemsBatchUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLineItemsBatchUpdateRequest struct via the builder pattern


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


## PostLineItemsBatchUpsert

> BatchResponseSimplePublicUpsertObject PostLineItemsBatchUpsert(ctx).BatchInputSimplePublicObjectBatchInputUpsert(batchInputSimplePublicObjectBatchInputUpsert).Execute()

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
	resp, r, err := apiClient.BatchAPI.PostLineItemsBatchUpsert(context.Background()).BatchInputSimplePublicObjectBatchInputUpsert(batchInputSimplePublicObjectBatchInputUpsert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostLineItemsBatchUpsert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLineItemsBatchUpsert`: BatchResponseSimplePublicUpsertObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostLineItemsBatchUpsert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLineItemsBatchUpsertRequest struct via the builder pattern


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
	fromObjectId := "fromObjectId_example" // string | The unique identifier of the source object in the association.
	fromObjectType := "fromObjectType_example" // string | The type of the source object in the association.
	toObjectId := "toObjectId_example" // string | The unique identifier of the target object in the association.
	toObjectType := "toObjectType_example" // string | The type of the target object in the association.

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
**fromObjectId** | **string** | The unique identifier of the source object in the association. | 
**fromObjectType** | **string** | The type of the source object in the association. | 
**toObjectId** | **string** | The unique identifier of the target object in the association. | 
**toObjectType** | **string** | The type of the target object in the association. | 

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

