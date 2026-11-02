# \BatchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchArchive**](BatchAPI.md#BatchArchive) | **Post** /crm/associations/2026-09/{fromObjectType}/{toObjectType}/batch/archive | Remove associations
[**BatchCreate**](BatchAPI.md#BatchCreate) | **Post** /crm/associations/2026-09/{fromObjectType}/{toObjectType}/batch/create | Associate records (labelled)
[**PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreateCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreate**](BatchAPI.md#PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreateCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreate) | **Post** /crm/associations/2026-09/definitions/configurations/{fromObjectType}/{toObjectType}/batch/create | Create
[**PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurgeCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurge**](BatchAPI.md#PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurgeCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurge) | **Post** /crm/associations/2026-09/definitions/configurations/{fromObjectType}/{toObjectType}/batch/purge | Delete
[**PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdateCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdate**](BatchAPI.md#PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdateCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdate) | **Post** /crm/associations/2026-09/definitions/configurations/{fromObjectType}/{toObjectType}/batch/update | Update
[**PostCrmAssociations202609FromObjectTypeToObjectTypeBatchAssociateDefaultCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchAssociateDefault**](BatchAPI.md#PostCrmAssociations202609FromObjectTypeToObjectTypeBatchAssociateDefaultCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchAssociateDefault) | **Post** /crm/associations/2026-09/{fromObjectType}/{toObjectType}/batch/associate/default |  Create Default Associations
[**PostCrmAssociations202609FromObjectTypeToObjectTypeBatchLabelsArchiveCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchLabelsArchive**](BatchAPI.md#PostCrmAssociations202609FromObjectTypeToObjectTypeBatchLabelsArchiveCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchLabelsArchive) | **Post** /crm/associations/2026-09/{fromObjectType}/{toObjectType}/batch/labels/archive | Delete Specific Labels
[**PostCrmAssociations202609FromObjectTypeToObjectTypeBatchReadCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchRead**](BatchAPI.md#PostCrmAssociations202609FromObjectTypeToObjectTypeBatchReadCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchRead) | **Post** /crm/associations/2026-09/{fromObjectType}/{toObjectType}/batch/read | Retrieve associations
[**PostCrmObjects202609ObjectTypeBatchArchiveCrmObjects202609BetaObjectTypeBatchArchive**](BatchAPI.md#PostCrmObjects202609ObjectTypeBatchArchiveCrmObjects202609BetaObjectTypeBatchArchive) | **Post** /crm/objects/2026-09/{objectType}/batch/archive | Archive a batch of objects by ID
[**PostCrmObjects202609ObjectTypeBatchCreateCrmObjects202609BetaObjectTypeBatchCreate**](BatchAPI.md#PostCrmObjects202609ObjectTypeBatchCreateCrmObjects202609BetaObjectTypeBatchCreate) | **Post** /crm/objects/2026-09/{objectType}/batch/create | Create a batch of objects
[**PostCrmObjects202609ObjectTypeBatchReadCrmObjects202609BetaObjectTypeBatchRead**](BatchAPI.md#PostCrmObjects202609ObjectTypeBatchReadCrmObjects202609BetaObjectTypeBatchRead) | **Post** /crm/objects/2026-09/{objectType}/batch/read | Read a batch of objects by internal ID, or unique property values
[**PostCrmObjects202609ObjectTypeBatchUpdateCrmObjects202609BetaObjectTypeBatchUpdate**](BatchAPI.md#PostCrmObjects202609ObjectTypeBatchUpdateCrmObjects202609BetaObjectTypeBatchUpdate) | **Post** /crm/objects/2026-09/{objectType}/batch/update | Update a batch of objects by internal ID, or unique property values
[**PostCrmObjects202609ObjectTypeBatchUpsertCrmObjects202609BetaObjectTypeBatchUpsert**](BatchAPI.md#PostCrmObjects202609ObjectTypeBatchUpsertCrmObjects202609BetaObjectTypeBatchUpsert) | **Post** /crm/objects/2026-09/{objectType}/batch/upsert | Upsert a batch of CRM objects by their internal ID or unique property values.
[**PutCrmObjects202609FromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectIdCrmObjects202609BetaFromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectId**](BatchAPI.md#PutCrmObjects202609FromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectIdCrmObjects202609BetaFromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectId) | **Put** /crm/objects/2026-09/{fromObjectType}/{fromObjectId}/associations/default/{toObjectType}/{toObjectId} | Associate records (default)



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

> BatchResponseLabelsBetweenObjectPair BatchCreate(ctx, fromObjectType, toObjectType).BatchInputPublicAssociationMultiPost(batchInputPublicAssociationMultiPost).Execute()

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
	resp, r, err := apiClient.BatchAPI.BatchCreate(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociationMultiPost(batchInputPublicAssociationMultiPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchCreate`: BatchResponseLabelsBetweenObjectPair
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** | The type of the source object for the associations. | 
**toObjectType** | **string** | The type of the target object for the associations. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchCreateRequest struct via the builder pattern


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


## PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreateCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreate

> BatchResponsePublicAssociationDefinitionUserConfiguration PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreateCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreate(ctx, fromObjectType, toObjectType).BatchInputPublicAssociationDefinitionConfigurationCreateRequest(batchInputPublicAssociationDefinitionConfigurationCreateRequest).Execute()

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
	resp, r, err := apiClient.BatchAPI.PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreateCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreate(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociationDefinitionConfigurationCreateRequest(batchInputPublicAssociationDefinitionConfigurationCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreateCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreateCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreate`: BatchResponsePublicAssociationDefinitionUserConfiguration
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreateCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** | The type of the source object for the association configuration. | 
**toObjectType** | **string** | The type of the target object for the association configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreateCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchCreateRequest struct via the builder pattern


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


## PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurgeCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurge

> BatchResponse PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurgeCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurge(ctx, fromObjectType, toObjectType).BatchInputPublicAssociationSpec(batchInputPublicAssociationSpec).Execute()

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
	resp, r, err := apiClient.BatchAPI.PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurgeCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurge(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociationSpec(batchInputPublicAssociationSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurgeCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurgeCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurge`: BatchResponse
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurgeCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** | The type of the source object from which the association is being purged. | 
**toObjectType** | **string** | The type of the target object to which the association is being purged. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurgeCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchPurgeRequest struct via the builder pattern


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


## PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdateCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdate

> BatchResponsePublicAssociationDefinitionConfigurationUpdateResult PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdateCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdate(ctx, fromObjectType, toObjectType).BatchInputPublicAssociationDefinitionConfigurationUpdateRequest(batchInputPublicAssociationDefinitionConfigurationUpdateRequest).Execute()

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
	resp, r, err := apiClient.BatchAPI.PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdateCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdate(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociationDefinitionConfigurationUpdateRequest(batchInputPublicAssociationDefinitionConfigurationUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdateCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdateCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdate`: BatchResponsePublicAssociationDefinitionConfigurationUpdateResult
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdateCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** | The type of the source object for which the association configurations are being updated. | 
**toObjectType** | **string** | The type of the target object for which the association configurations are being updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdateCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeBatchUpdateRequest struct via the builder pattern


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


## PostCrmAssociations202609FromObjectTypeToObjectTypeBatchAssociateDefaultCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchAssociateDefault

> BatchResponsePublicDefaultAssociation PostCrmAssociations202609FromObjectTypeToObjectTypeBatchAssociateDefaultCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchAssociateDefault(ctx, fromObjectType, toObjectType).BatchInputPublicDefaultAssociationMultiPost(batchInputPublicDefaultAssociationMultiPost).Execute()

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
	resp, r, err := apiClient.BatchAPI.PostCrmAssociations202609FromObjectTypeToObjectTypeBatchAssociateDefaultCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchAssociateDefault(context.Background(), fromObjectType, toObjectType).BatchInputPublicDefaultAssociationMultiPost(batchInputPublicDefaultAssociationMultiPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostCrmAssociations202609FromObjectTypeToObjectTypeBatchAssociateDefaultCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchAssociateDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmAssociations202609FromObjectTypeToObjectTypeBatchAssociateDefaultCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchAssociateDefault`: BatchResponsePublicDefaultAssociation
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostCrmAssociations202609FromObjectTypeToObjectTypeBatchAssociateDefaultCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchAssociateDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** | The type of the source object for the association. | 
**toObjectType** | **string** | The type of the target object for the association. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmAssociations202609FromObjectTypeToObjectTypeBatchAssociateDefaultCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchAssociateDefaultRequest struct via the builder pattern


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


## PostCrmAssociations202609FromObjectTypeToObjectTypeBatchLabelsArchiveCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchLabelsArchive

> BatchResponse PostCrmAssociations202609FromObjectTypeToObjectTypeBatchLabelsArchiveCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchLabelsArchive(ctx, fromObjectType, toObjectType).BatchInputPublicAssociationMultiPost(batchInputPublicAssociationMultiPost).Execute()

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
	resp, r, err := apiClient.BatchAPI.PostCrmAssociations202609FromObjectTypeToObjectTypeBatchLabelsArchiveCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchLabelsArchive(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociationMultiPost(batchInputPublicAssociationMultiPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostCrmAssociations202609FromObjectTypeToObjectTypeBatchLabelsArchiveCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchLabelsArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmAssociations202609FromObjectTypeToObjectTypeBatchLabelsArchiveCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchLabelsArchive`: BatchResponse
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostCrmAssociations202609FromObjectTypeToObjectTypeBatchLabelsArchiveCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchLabelsArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** | The type of the source object in the association. | 
**toObjectType** | **string** | The type of the target object in the association. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmAssociations202609FromObjectTypeToObjectTypeBatchLabelsArchiveCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchLabelsArchiveRequest struct via the builder pattern


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


## PostCrmAssociations202609FromObjectTypeToObjectTypeBatchReadCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchRead

> BatchResponsePublicAssociationMultiWithLabel PostCrmAssociations202609FromObjectTypeToObjectTypeBatchReadCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchRead(ctx, fromObjectType, toObjectType).BatchInputPublicFetchAssociationsBatchRequest(batchInputPublicFetchAssociationsBatchRequest).Execute()

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
	resp, r, err := apiClient.BatchAPI.PostCrmAssociations202609FromObjectTypeToObjectTypeBatchReadCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchRead(context.Background(), fromObjectType, toObjectType).BatchInputPublicFetchAssociationsBatchRequest(batchInputPublicFetchAssociationsBatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostCrmAssociations202609FromObjectTypeToObjectTypeBatchReadCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmAssociations202609FromObjectTypeToObjectTypeBatchReadCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchRead`: BatchResponsePublicAssociationMultiWithLabel
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostCrmAssociations202609FromObjectTypeToObjectTypeBatchReadCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** | The type of the source object from which associations are being read. | 
**toObjectType** | **string** | The type of the target object to which associations are being read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmAssociations202609FromObjectTypeToObjectTypeBatchReadCrmAssociations202609BetaFromObjectTypeToObjectTypeBatchReadRequest struct via the builder pattern


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


## PostCrmObjects202609ObjectTypeBatchArchiveCrmObjects202609BetaObjectTypeBatchArchive

> PostCrmObjects202609ObjectTypeBatchArchiveCrmObjects202609BetaObjectTypeBatchArchive(ctx, objectType).BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId).Execute()

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
	r, err := apiClient.BatchAPI.PostCrmObjects202609ObjectTypeBatchArchiveCrmObjects202609BetaObjectTypeBatchArchive(context.Background(), objectType).BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostCrmObjects202609ObjectTypeBatchArchiveCrmObjects202609BetaObjectTypeBatchArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostCrmObjects202609ObjectTypeBatchArchiveCrmObjects202609BetaObjectTypeBatchArchiveRequest struct via the builder pattern


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


## PostCrmObjects202609ObjectTypeBatchCreateCrmObjects202609BetaObjectTypeBatchCreate

> BatchResponseSimplePublicObject PostCrmObjects202609ObjectTypeBatchCreateCrmObjects202609BetaObjectTypeBatchCreate(ctx, objectType).BatchInputSimplePublicObjectBatchInputForCreate(batchInputSimplePublicObjectBatchInputForCreate).Execute()

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
	resp, r, err := apiClient.BatchAPI.PostCrmObjects202609ObjectTypeBatchCreateCrmObjects202609BetaObjectTypeBatchCreate(context.Background(), objectType).BatchInputSimplePublicObjectBatchInputForCreate(batchInputSimplePublicObjectBatchInputForCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostCrmObjects202609ObjectTypeBatchCreateCrmObjects202609BetaObjectTypeBatchCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmObjects202609ObjectTypeBatchCreateCrmObjects202609BetaObjectTypeBatchCreate`: BatchResponseSimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostCrmObjects202609ObjectTypeBatchCreateCrmObjects202609BetaObjectTypeBatchCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of CRM object to create in batch, such as contacts, companies, or deals. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmObjects202609ObjectTypeBatchCreateCrmObjects202609BetaObjectTypeBatchCreateRequest struct via the builder pattern


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


## PostCrmObjects202609ObjectTypeBatchReadCrmObjects202609BetaObjectTypeBatchRead

> BatchResponseSimplePublicObject PostCrmObjects202609ObjectTypeBatchReadCrmObjects202609BetaObjectTypeBatchRead(ctx, objectType).BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId).Archived(archived).Execute()

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
	resp, r, err := apiClient.BatchAPI.PostCrmObjects202609ObjectTypeBatchReadCrmObjects202609BetaObjectTypeBatchRead(context.Background(), objectType).BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostCrmObjects202609ObjectTypeBatchReadCrmObjects202609BetaObjectTypeBatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmObjects202609ObjectTypeBatchReadCrmObjects202609BetaObjectTypeBatchRead`: BatchResponseSimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostCrmObjects202609ObjectTypeBatchReadCrmObjects202609BetaObjectTypeBatchRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of CRM object to read in batch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmObjects202609ObjectTypeBatchReadCrmObjects202609BetaObjectTypeBatchReadRequest struct via the builder pattern


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


## PostCrmObjects202609ObjectTypeBatchUpdateCrmObjects202609BetaObjectTypeBatchUpdate

> BatchResponseSimplePublicObject PostCrmObjects202609ObjectTypeBatchUpdateCrmObjects202609BetaObjectTypeBatchUpdate(ctx, objectType).BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput).Execute()

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
	resp, r, err := apiClient.BatchAPI.PostCrmObjects202609ObjectTypeBatchUpdateCrmObjects202609BetaObjectTypeBatchUpdate(context.Background(), objectType).BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostCrmObjects202609ObjectTypeBatchUpdateCrmObjects202609BetaObjectTypeBatchUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmObjects202609ObjectTypeBatchUpdateCrmObjects202609BetaObjectTypeBatchUpdate`: BatchResponseSimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostCrmObjects202609ObjectTypeBatchUpdateCrmObjects202609BetaObjectTypeBatchUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of CRM object to update in batch, such as contacts, companies, or deals. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmObjects202609ObjectTypeBatchUpdateCrmObjects202609BetaObjectTypeBatchUpdateRequest struct via the builder pattern


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


## PostCrmObjects202609ObjectTypeBatchUpsertCrmObjects202609BetaObjectTypeBatchUpsert

> BatchResponseSimplePublicUpsertObject PostCrmObjects202609ObjectTypeBatchUpsertCrmObjects202609BetaObjectTypeBatchUpsert(ctx, objectType).BatchInputSimplePublicObjectBatchInputUpsert(batchInputSimplePublicObjectBatchInputUpsert).Execute()

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
	resp, r, err := apiClient.BatchAPI.PostCrmObjects202609ObjectTypeBatchUpsertCrmObjects202609BetaObjectTypeBatchUpsert(context.Background(), objectType).BatchInputSimplePublicObjectBatchInputUpsert(batchInputSimplePublicObjectBatchInputUpsert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostCrmObjects202609ObjectTypeBatchUpsertCrmObjects202609BetaObjectTypeBatchUpsert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmObjects202609ObjectTypeBatchUpsertCrmObjects202609BetaObjectTypeBatchUpsert`: BatchResponseSimplePublicUpsertObject
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostCrmObjects202609ObjectTypeBatchUpsertCrmObjects202609BetaObjectTypeBatchUpsert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of CRM object to be upserted, such as contacts, companies, or deals. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmObjects202609ObjectTypeBatchUpsertCrmObjects202609BetaObjectTypeBatchUpsertRequest struct via the builder pattern


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


## PutCrmObjects202609FromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectIdCrmObjects202609BetaFromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectId

> BatchResponsePublicDefaultAssociation PutCrmObjects202609FromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectIdCrmObjects202609BetaFromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectId(ctx, fromObjectId, fromObjectType, toObjectId, toObjectType).Execute()

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
	resp, r, err := apiClient.BatchAPI.PutCrmObjects202609FromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectIdCrmObjects202609BetaFromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectId(context.Background(), fromObjectId, fromObjectType, toObjectId, toObjectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PutCrmObjects202609FromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectIdCrmObjects202609BetaFromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCrmObjects202609FromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectIdCrmObjects202609BetaFromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectId`: BatchResponsePublicDefaultAssociation
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PutCrmObjects202609FromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectIdCrmObjects202609BetaFromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectId`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPutCrmObjects202609FromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectIdCrmObjects202609BetaFromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectIdRequest struct via the builder pattern


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

