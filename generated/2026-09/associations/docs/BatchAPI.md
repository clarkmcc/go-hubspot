# \BatchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchArchive**](BatchAPI.md#BatchArchive) | **Post** /crm/associations/2026-09/{fromObjectType}/{toObjectType}/batch/archive | Remove associations
[**BatchCreate**](BatchAPI.md#BatchCreate) | **Post** /crm/associations/2026-09/definitions/configurations/{fromObjectType}/{toObjectType}/batch/create | Create
[**BatchCreate_0**](BatchAPI.md#BatchCreate_0) | **Post** /crm/associations/2026-09/{fromObjectType}/{toObjectType}/batch/create | Associate records (labelled)
[**PostBatchArchive**](BatchAPI.md#PostBatchArchive) | **Post** /crm/objects/2026-09/{objectType}/batch/archive | Archive a batch of objects by ID
[**PostBatchAssociateDefault**](BatchAPI.md#PostBatchAssociateDefault) | **Post** /crm/associations/2026-09/{fromObjectType}/{toObjectType}/batch/associate/default |  Create Default Associations
[**PostBatchCreate**](BatchAPI.md#PostBatchCreate) | **Post** /crm/objects/2026-09/{objectType}/batch/create | Create a batch of objects
[**PostBatchLabelsArchive**](BatchAPI.md#PostBatchLabelsArchive) | **Post** /crm/associations/2026-09/{fromObjectType}/{toObjectType}/batch/labels/archive | Delete Specific Labels
[**PostBatchRead**](BatchAPI.md#PostBatchRead) | **Post** /crm/associations/2026-09/{fromObjectType}/{toObjectType}/batch/read | Retrieve associations
[**PostBatchUpdate**](BatchAPI.md#PostBatchUpdate) | **Post** /crm/objects/2026-09/{objectType}/batch/update | Update a batch of objects by internal ID, or unique property values
[**PostBatchUpsert**](BatchAPI.md#PostBatchUpsert) | **Post** /crm/objects/2026-09/{objectType}/batch/upsert | Upsert a batch of CRM objects by their internal ID or unique property values.
[**PostByObjectTypeBatchRead**](BatchAPI.md#PostByObjectTypeBatchRead) | **Post** /crm/objects/2026-09/{objectType}/batch/read | Read a batch of objects by internal ID, or unique property values
[**PostDefinitionsConfigurationsBatchPurge**](BatchAPI.md#PostDefinitionsConfigurationsBatchPurge) | **Post** /crm/associations/2026-09/definitions/configurations/{fromObjectType}/{toObjectType}/batch/purge | Delete
[**PostDefinitionsConfigurationsBatchUpdate**](BatchAPI.md#PostDefinitionsConfigurationsBatchUpdate) | **Post** /crm/associations/2026-09/definitions/configurations/{fromObjectType}/{toObjectType}/batch/update | Update
[**PutDefault**](BatchAPI.md#PutDefault) | **Put** /crm/objects/2026-09/{fromObjectType}/{fromObjectId}/associations/default/{toObjectType}/{toObjectId} | Associate records (default)



## BatchArchive

> BatchResponse BatchArchive(ctx, fromObjectType, toObjectType).BatchInputPublicAssociationMultiArchive(batchInputPublicAssociationMultiArchive).Execute()

Remove associations



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
	fromObjectType := "fromObjectType_example" // string | The type of the source object from which associations will be deleted.
	toObjectType := "toObjectType_example" // string | The type of the target object to which associations will be deleted.
	batchInputPublicAssociationMultiArchive := *openapiclient.NewBatchInputPublicAssociationMultiArchive([]openapiclient.PublicAssociationMultiArchive{*openapiclient.NewPublicAssociationMultiArchive(*openapiclient.NewPublicObjectId("Id_example"), []openapiclient.PublicObjectId{*openapiclient.NewPublicObjectId("Id_example")})}) // BatchInputPublicAssociationMultiArchive | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.BatchArchive(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociationMultiArchive(batchInputPublicAssociationMultiArchive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchArchive`: BatchResponse
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** | The type of the source object from which associations will be deleted. | 
**toObjectType** | **string** | The type of the target object to which associations will be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicAssociationMultiArchive** | [**BatchInputPublicAssociationMultiArchive**](BatchInputPublicAssociationMultiArchive.md) |  | 

### Return type

[**BatchResponse**](BatchResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchCreate

> BatchResponsePublicAssociationDefinitionUserConfiguration BatchCreate(ctx, fromObjectType, toObjectType).BatchInputPublicAssociationDefinitionConfigurationCreateRequest(batchInputPublicAssociationDefinitionConfigurationCreateRequest).Execute()

Create



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
	fromObjectType := "fromObjectType_example" // string | The type of the source object for the association configuration.
	toObjectType := "toObjectType_example" // string | The type of the target object for the association configuration.
	batchInputPublicAssociationDefinitionConfigurationCreateRequest := *openapiclient.NewBatchInputPublicAssociationDefinitionConfigurationCreateRequest([]openapiclient.PublicAssociationDefinitionConfigurationCreateRequest{*openapiclient.NewPublicAssociationDefinitionConfigurationCreateRequest("Category_example", int32(123), int32(123))}) // BatchInputPublicAssociationDefinitionConfigurationCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.BatchCreate(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociationDefinitionConfigurationCreateRequest(batchInputPublicAssociationDefinitionConfigurationCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchCreate`: BatchResponsePublicAssociationDefinitionUserConfiguration
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** | The type of the source object for the association configuration. | 
**toObjectType** | **string** | The type of the target object for the association configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicAssociationDefinitionConfigurationCreateRequest** | [**BatchInputPublicAssociationDefinitionConfigurationCreateRequest**](BatchInputPublicAssociationDefinitionConfigurationCreateRequest.md) |  | 

### Return type

[**BatchResponsePublicAssociationDefinitionUserConfiguration**](BatchResponsePublicAssociationDefinitionUserConfiguration.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchCreate_0

> BatchResponseLabelsBetweenObjectPair BatchCreate_0(ctx, fromObjectType, toObjectType).BatchInputPublicAssociationMultiPost(batchInputPublicAssociationMultiPost).Execute()

Associate records (labelled)



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
	fromObjectType := "fromObjectType_example" // string | The type of the source object for the associations.
	toObjectType := "toObjectType_example" // string | The type of the target object for the associations.
	batchInputPublicAssociationMultiPost := *openapiclient.NewBatchInputPublicAssociationMultiPost([]openapiclient.PublicAssociationMultiPost{*openapiclient.NewPublicAssociationMultiPost(*openapiclient.NewPublicObjectId("Id_example"), *openapiclient.NewPublicObjectId("Id_example"), []openapiclient.AssociationSpec{*openapiclient.NewAssociationSpec("AssociationCategory_example", int32(123))})}) // BatchInputPublicAssociationMultiPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.BatchCreate_0(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociationMultiPost(batchInputPublicAssociationMultiPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchCreate_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchCreate_0`: BatchResponseLabelsBetweenObjectPair
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchCreate_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** | The type of the source object for the associations. | 
**toObjectType** | **string** | The type of the target object for the associations. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchCreate_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicAssociationMultiPost** | [**BatchInputPublicAssociationMultiPost**](BatchInputPublicAssociationMultiPost.md) |  | 

### Return type

[**BatchResponseLabelsBetweenObjectPair**](BatchResponseLabelsBetweenObjectPair.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBatchArchive

> PostBatchArchive(ctx, objectType).BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId).Execute()

Archive a batch of objects by ID



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
	objectType := "objectType_example" // string | The type of CRM object to archive in batch, such as contacts, companies, deals, etc.
	batchInputSimplePublicObjectId := *openapiclient.NewBatchInputSimplePublicObjectId([]openapiclient.SimplePublicObjectId{*openapiclient.NewSimplePublicObjectId("430001")}) // BatchInputSimplePublicObjectId | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchAPI.PostBatchArchive(context.Background(), objectType).BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostBatchArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of CRM object to archive in batch, such as contacts, companies, deals, etc. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBatchArchiveRequest struct via the builder pattern


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


## PostBatchAssociateDefault

> BatchResponsePublicDefaultAssociation PostBatchAssociateDefault(ctx, fromObjectType, toObjectType).BatchInputPublicDefaultAssociationMultiPost(batchInputPublicDefaultAssociationMultiPost).Execute()

 Create Default Associations



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
	fromObjectType := "fromObjectType_example" // string | The type of the source object for the association.
	toObjectType := "toObjectType_example" // string | The type of the target object for the association.
	batchInputPublicDefaultAssociationMultiPost := *openapiclient.NewBatchInputPublicDefaultAssociationMultiPost([]openapiclient.PublicDefaultAssociationMultiPost{*openapiclient.NewPublicDefaultAssociationMultiPost(*openapiclient.NewPublicObjectId("Id_example"), *openapiclient.NewPublicObjectId("Id_example"))}) // BatchInputPublicDefaultAssociationMultiPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostBatchAssociateDefault(context.Background(), fromObjectType, toObjectType).BatchInputPublicDefaultAssociationMultiPost(batchInputPublicDefaultAssociationMultiPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostBatchAssociateDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBatchAssociateDefault`: BatchResponsePublicDefaultAssociation
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostBatchAssociateDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** | The type of the source object for the association. | 
**toObjectType** | **string** | The type of the target object for the association. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBatchAssociateDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicDefaultAssociationMultiPost** | [**BatchInputPublicDefaultAssociationMultiPost**](BatchInputPublicDefaultAssociationMultiPost.md) |  | 

### Return type

[**BatchResponsePublicDefaultAssociation**](BatchResponsePublicDefaultAssociation.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBatchCreate

> BatchResponseSimplePublicObject PostBatchCreate(ctx, objectType).BatchInputSimplePublicObjectBatchInputForCreate(batchInputSimplePublicObjectBatchInputForCreate).Execute()

Create a batch of objects



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
	objectType := "objectType_example" // string | The type of CRM object to create in batch, such as contacts, companies, or deals.
	batchInputSimplePublicObjectBatchInputForCreate := *openapiclient.NewBatchInputSimplePublicObjectBatchInputForCreate([]openapiclient.SimplePublicObjectBatchInputForCreate{*openapiclient.NewSimplePublicObjectBatchInputForCreate([]openapiclient.PublicAssociationsForObject{*openapiclient.NewPublicAssociationsForObject(*openapiclient.NewPublicObjectId("Id_example"), []openapiclient.AssociationSpec{*openapiclient.NewAssociationSpec("AssociationCategory_example", int32(123))})}, map[string]string{"key": "Inner_example"})}) // BatchInputSimplePublicObjectBatchInputForCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostBatchCreate(context.Background(), objectType).BatchInputSimplePublicObjectBatchInputForCreate(batchInputSimplePublicObjectBatchInputForCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostBatchCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBatchCreate`: BatchResponseSimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostBatchCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of CRM object to create in batch, such as contacts, companies, or deals. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBatchCreateRequest struct via the builder pattern


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


## PostBatchLabelsArchive

> BatchResponse PostBatchLabelsArchive(ctx, fromObjectType, toObjectType).BatchInputPublicAssociationMultiPost(batchInputPublicAssociationMultiPost).Execute()

Delete Specific Labels



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
	fromObjectType := "fromObjectType_example" // string | The type of the source object in the association.
	toObjectType := "toObjectType_example" // string | The type of the target object in the association.
	batchInputPublicAssociationMultiPost := *openapiclient.NewBatchInputPublicAssociationMultiPost([]openapiclient.PublicAssociationMultiPost{*openapiclient.NewPublicAssociationMultiPost(*openapiclient.NewPublicObjectId("Id_example"), *openapiclient.NewPublicObjectId("Id_example"), []openapiclient.AssociationSpec{*openapiclient.NewAssociationSpec("AssociationCategory_example", int32(123))})}) // BatchInputPublicAssociationMultiPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostBatchLabelsArchive(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociationMultiPost(batchInputPublicAssociationMultiPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostBatchLabelsArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBatchLabelsArchive`: BatchResponse
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostBatchLabelsArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** | The type of the source object in the association. | 
**toObjectType** | **string** | The type of the target object in the association. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBatchLabelsArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicAssociationMultiPost** | [**BatchInputPublicAssociationMultiPost**](BatchInputPublicAssociationMultiPost.md) |  | 

### Return type

[**BatchResponse**](BatchResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBatchRead

> BatchResponsePublicAssociationMultiWithLabel PostBatchRead(ctx, fromObjectType, toObjectType).BatchInputPublicFetchAssociationsBatchRequest(batchInputPublicFetchAssociationsBatchRequest).Execute()

Retrieve associations



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
	fromObjectType := "fromObjectType_example" // string | The type of the source object from which associations are being read.
	toObjectType := "toObjectType_example" // string | The type of the target object to which associations are being read.
	batchInputPublicFetchAssociationsBatchRequest := *openapiclient.NewBatchInputPublicFetchAssociationsBatchRequest([]openapiclient.PublicFetchAssociationsBatchRequest{*openapiclient.NewPublicFetchAssociationsBatchRequest("Id_example")}) // BatchInputPublicFetchAssociationsBatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostBatchRead(context.Background(), fromObjectType, toObjectType).BatchInputPublicFetchAssociationsBatchRequest(batchInputPublicFetchAssociationsBatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostBatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBatchRead`: BatchResponsePublicAssociationMultiWithLabel
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostBatchRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** | The type of the source object from which associations are being read. | 
**toObjectType** | **string** | The type of the target object to which associations are being read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBatchReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicFetchAssociationsBatchRequest** | [**BatchInputPublicFetchAssociationsBatchRequest**](BatchInputPublicFetchAssociationsBatchRequest.md) |  | 

### Return type

[**BatchResponsePublicAssociationMultiWithLabel**](BatchResponsePublicAssociationMultiWithLabel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBatchUpdate

> BatchResponseSimplePublicObject PostBatchUpdate(ctx, objectType).BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput).Execute()

Update a batch of objects by internal ID, or unique property values



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
	objectType := "objectType_example" // string | The type of CRM object to update in batch, such as contacts, companies, or deals.
	batchInputSimplePublicObjectBatchInput := *openapiclient.NewBatchInputSimplePublicObjectBatchInput([]openapiclient.SimplePublicObjectBatchInput{*openapiclient.NewSimplePublicObjectBatchInput("Id_example", map[string]string{"key": "Inner_example"})}) // BatchInputSimplePublicObjectBatchInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostBatchUpdate(context.Background(), objectType).BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostBatchUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBatchUpdate`: BatchResponseSimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostBatchUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of CRM object to update in batch, such as contacts, companies, or deals. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBatchUpdateRequest struct via the builder pattern


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


## PostBatchUpsert

> BatchResponseSimplePublicUpsertObject PostBatchUpsert(ctx, objectType).BatchInputSimplePublicObjectBatchInputUpsert(batchInputSimplePublicObjectBatchInputUpsert).Execute()

Upsert a batch of CRM objects by their internal ID or unique property values.



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
	objectType := "objectType_example" // string | The type of CRM object to be upserted, such as contacts, companies, or deals.
	batchInputSimplePublicObjectBatchInputUpsert := *openapiclient.NewBatchInputSimplePublicObjectBatchInputUpsert([]openapiclient.SimplePublicObjectBatchInputUpsert{*openapiclient.NewSimplePublicObjectBatchInputUpsert("Id_example", map[string]string{"key": "Inner_example"})}) // BatchInputSimplePublicObjectBatchInputUpsert | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostBatchUpsert(context.Background(), objectType).BatchInputSimplePublicObjectBatchInputUpsert(batchInputSimplePublicObjectBatchInputUpsert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostBatchUpsert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBatchUpsert`: BatchResponseSimplePublicUpsertObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostBatchUpsert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of CRM object to be upserted, such as contacts, companies, or deals. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBatchUpsertRequest struct via the builder pattern


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


## PostByObjectTypeBatchRead

> BatchResponseSimplePublicObject PostByObjectTypeBatchRead(ctx, objectType).BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId).Archived(archived).Execute()

Read a batch of objects by internal ID, or unique property values



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
	objectType := "objectType_example" // string | The type of CRM object to read in batch.
	batchReadInputSimplePublicObjectId := *openapiclient.NewBatchReadInputSimplePublicObjectId([]openapiclient.SimplePublicObjectId{*openapiclient.NewSimplePublicObjectId("430001")}, []string{"Properties_example"}, []string{"PropertiesWithHistory_example"}) // BatchReadInputSimplePublicObjectId | 
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostByObjectTypeBatchRead(context.Background(), objectType).BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostByObjectTypeBatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostByObjectTypeBatchRead`: BatchResponseSimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostByObjectTypeBatchRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of CRM object to read in batch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostByObjectTypeBatchReadRequest struct via the builder pattern


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


## PostDefinitionsConfigurationsBatchPurge

> BatchResponse PostDefinitionsConfigurationsBatchPurge(ctx, fromObjectType, toObjectType).BatchInputPublicAssociationSpec(batchInputPublicAssociationSpec).Execute()

Delete



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
	fromObjectType := "fromObjectType_example" // string | The type of the source object from which the association is being purged.
	toObjectType := "toObjectType_example" // string | The type of the target object to which the association is being purged.
	batchInputPublicAssociationSpec := *openapiclient.NewBatchInputPublicAssociationSpec([]openapiclient.PublicAssociationSpec{*openapiclient.NewPublicAssociationSpec("Category_example", int32(123))}) // BatchInputPublicAssociationSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostDefinitionsConfigurationsBatchPurge(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociationSpec(batchInputPublicAssociationSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostDefinitionsConfigurationsBatchPurge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDefinitionsConfigurationsBatchPurge`: BatchResponse
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostDefinitionsConfigurationsBatchPurge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** | The type of the source object from which the association is being purged. | 
**toObjectType** | **string** | The type of the target object to which the association is being purged. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDefinitionsConfigurationsBatchPurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicAssociationSpec** | [**BatchInputPublicAssociationSpec**](BatchInputPublicAssociationSpec.md) |  | 

### Return type

[**BatchResponse**](BatchResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDefinitionsConfigurationsBatchUpdate

> BatchResponsePublicAssociationDefinitionConfigurationUpdateResult PostDefinitionsConfigurationsBatchUpdate(ctx, fromObjectType, toObjectType).BatchInputPublicAssociationDefinitionConfigurationUpdateRequest(batchInputPublicAssociationDefinitionConfigurationUpdateRequest).Execute()

Update



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
	fromObjectType := "fromObjectType_example" // string | The type of the source object for which the association configurations are being updated.
	toObjectType := "toObjectType_example" // string | The type of the target object for which the association configurations are being updated.
	batchInputPublicAssociationDefinitionConfigurationUpdateRequest := *openapiclient.NewBatchInputPublicAssociationDefinitionConfigurationUpdateRequest([]openapiclient.PublicAssociationDefinitionConfigurationUpdateRequest{*openapiclient.NewPublicAssociationDefinitionConfigurationUpdateRequest("Category_example", int32(123), int32(123))}) // BatchInputPublicAssociationDefinitionConfigurationUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostDefinitionsConfigurationsBatchUpdate(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociationDefinitionConfigurationUpdateRequest(batchInputPublicAssociationDefinitionConfigurationUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostDefinitionsConfigurationsBatchUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDefinitionsConfigurationsBatchUpdate`: BatchResponsePublicAssociationDefinitionConfigurationUpdateResult
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostDefinitionsConfigurationsBatchUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** | The type of the source object for which the association configurations are being updated. | 
**toObjectType** | **string** | The type of the target object for which the association configurations are being updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDefinitionsConfigurationsBatchUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicAssociationDefinitionConfigurationUpdateRequest** | [**BatchInputPublicAssociationDefinitionConfigurationUpdateRequest**](BatchInputPublicAssociationDefinitionConfigurationUpdateRequest.md) |  | 

### Return type

[**BatchResponsePublicAssociationDefinitionConfigurationUpdateResult**](BatchResponsePublicAssociationDefinitionConfigurationUpdateResult.md)

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

Associate records (default)



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

