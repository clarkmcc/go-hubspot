# \BasicAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePipelines**](BasicAPI.md#DeletePipelines) | **Delete** /crm/pipelines/2026-09/{objectType}/{pipelineId} | Delete pipeline
[**DeletePipelinesStages**](BasicAPI.md#DeletePipelinesStages) | **Delete** /crm/pipelines/2026-09/{objectType}/{pipelineId}/stages/{stageId} | Delete stage
[**GetPipelines**](BasicAPI.md#GetPipelines) | **Get** /crm/pipelines/2026-09/{objectType} | Retrieve pipelines
[**GetPipelinesAudit**](BasicAPI.md#GetPipelinesAudit) | **Get** /crm/pipelines/2026-09/{objectType}/{pipelineId}/audit | Retrieve audit logs
[**GetPipelinesByObjectTypeByPipelineId**](BasicAPI.md#GetPipelinesByObjectTypeByPipelineId) | **Get** /crm/pipelines/2026-09/{objectType}/{pipelineId} | Retrieve pipeline
[**GetPipelinesByObjectTypeByPipelineIdStagesByStageId**](BasicAPI.md#GetPipelinesByObjectTypeByPipelineIdStagesByStageId) | **Get** /crm/pipelines/2026-09/{objectType}/{pipelineId}/stages/{stageId} | Retrieve stage
[**GetPipelinesStages**](BasicAPI.md#GetPipelinesStages) | **Get** /crm/pipelines/2026-09/{objectType}/{pipelineId}/stages | List stages
[**GetPipelinesStagesAudit**](BasicAPI.md#GetPipelinesStagesAudit) | **Get** /crm/pipelines/2026-09/{objectType}/{pipelineId}/stages/{stageId}/audit | Retrieve stage audit
[**PatchPipelines**](BasicAPI.md#PatchPipelines) | **Patch** /crm/pipelines/2026-09/{objectType}/{pipelineId} | Update pipeline (Partial)
[**PatchPipelinesStages**](BasicAPI.md#PatchPipelinesStages) | **Patch** /crm/pipelines/2026-09/{objectType}/{pipelineId}/stages/{stageId} | Update stage (Partial)
[**PostPipelines**](BasicAPI.md#PostPipelines) | **Post** /crm/pipelines/2026-09/{objectType} | Create pipeline
[**PostPipelinesStages**](BasicAPI.md#PostPipelinesStages) | **Post** /crm/pipelines/2026-09/{objectType}/{pipelineId}/stages | Create Stage
[**PutPipelines**](BasicAPI.md#PutPipelines) | **Put** /crm/pipelines/2026-09/{objectType}/{pipelineId} | Update pipeline (Replace)
[**PutPipelinesStages**](BasicAPI.md#PutPipelinesStages) | **Put** /crm/pipelines/2026-09/{objectType}/{pipelineId}/stages/{stageId} | Update stage



## DeletePipelines

> DeletePipelines(ctx, objectType, pipelineId).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).Execute()

Delete pipeline



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
	objectType := "objectType_example" // string | The type of object the pipeline belongs to, such as deals or tickets.
	pipelineId := "pipelineId_example" // string | The unique identifier of the pipeline to delete.
	validateDealStageUsagesBeforeDelete := true // bool | A boolean indicating whether to validate deal stage usages before deleting the pipeline. Defaults to true. (optional) (default to true)
	validateReferencesBeforeDelete := true // bool | A boolean indicating whether to validate references before deleting the pipeline. Defaults to true. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicAPI.DeletePipelines(context.Background(), objectType, pipelineId).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.DeletePipelines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of object the pipeline belongs to, such as deals or tickets. | 
**pipelineId** | **string** | The unique identifier of the pipeline to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePipelinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateDealStageUsagesBeforeDelete** | **bool** | A boolean indicating whether to validate deal stage usages before deleting the pipeline. Defaults to true. | [default to true]
 **validateReferencesBeforeDelete** | **bool** | A boolean indicating whether to validate references before deleting the pipeline. Defaults to true. | [default to true]

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePipelinesStages

> DeletePipelinesStages(ctx, objectType, pipelineId, stageId).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).Execute()

Delete stage



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
	objectType := "objectType_example" // string | The type of object the pipeline belongs to, such as deals or tickets.
	pipelineId := "pipelineId_example" // string | The unique identifier of the pipeline from which the stage will be deleted.
	stageId := "stageId_example" // string | The unique identifier of the stage to delete.
	validateReferencesBeforeDelete := true // bool | A boolean flag indicating whether to validate references before deleting the stage. Defaults to true. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicAPI.DeletePipelinesStages(context.Background(), objectType, pipelineId, stageId).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.DeletePipelinesStages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of object the pipeline belongs to, such as deals or tickets. | 
**pipelineId** | **string** | The unique identifier of the pipeline from which the stage will be deleted. | 
**stageId** | **string** | The unique identifier of the stage to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePipelinesStagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **validateReferencesBeforeDelete** | **bool** | A boolean flag indicating whether to validate references before deleting the stage. Defaults to true. | [default to true]

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelines

> CollectionResponsePipelineNoPaging GetPipelines(ctx, objectType).Execute()

Retrieve pipelines



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
	objectType := "objectType_example" // string | The type of object for which pipelines are being retrieved. This is a required parameter and must be specified in the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetPipelines(context.Background(), objectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetPipelines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPipelines`: CollectionResponsePipelineNoPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetPipelines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of object for which pipelines are being retrieved. This is a required parameter and must be specified in the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionResponsePipelineNoPaging**](CollectionResponsePipelineNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelinesAudit

> CollectionResponsePublicAuditInfoNoPaging GetPipelinesAudit(ctx, objectType, pipelineId).Execute()

Retrieve audit logs



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
	objectType := "objectType_example" // string | The type of object the pipeline is associated with, such as deals or tickets.
	pipelineId := "pipelineId_example" // string | The unique identifier of the pipeline for which audit logs are being retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetPipelinesAudit(context.Background(), objectType, pipelineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetPipelinesAudit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPipelinesAudit`: CollectionResponsePublicAuditInfoNoPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetPipelinesAudit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of object the pipeline is associated with, such as deals or tickets. | 
**pipelineId** | **string** | The unique identifier of the pipeline for which audit logs are being retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelinesAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CollectionResponsePublicAuditInfoNoPaging**](CollectionResponsePublicAuditInfoNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelinesByObjectTypeByPipelineId

> Pipeline GetPipelinesByObjectTypeByPipelineId(ctx, objectType, pipelineId).Execute()

Retrieve pipeline



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
	objectType := "objectType_example" // string | The type of object the pipeline is associated with, such as deals or tickets.
	pipelineId := "pipelineId_example" // string | The unique identifier of the pipeline to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetPipelinesByObjectTypeByPipelineId(context.Background(), objectType, pipelineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetPipelinesByObjectTypeByPipelineId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPipelinesByObjectTypeByPipelineId`: Pipeline
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetPipelinesByObjectTypeByPipelineId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of object the pipeline is associated with, such as deals or tickets. | 
**pipelineId** | **string** | The unique identifier of the pipeline to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelinesByObjectTypeByPipelineIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelinesByObjectTypeByPipelineIdStagesByStageId

> PipelineStage GetPipelinesByObjectTypeByPipelineIdStagesByStageId(ctx, objectType, pipelineId, stageId).Execute()

Retrieve stage



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
	objectType := "objectType_example" // string | The type of object for which the pipeline stage is being retrieved.
	pipelineId := "pipelineId_example" // string | The unique identifier of the pipeline that contains the stage.
	stageId := "stageId_example" // string | The unique identifier of the stage to retrieve within the pipeline.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetPipelinesByObjectTypeByPipelineIdStagesByStageId(context.Background(), objectType, pipelineId, stageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetPipelinesByObjectTypeByPipelineIdStagesByStageId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPipelinesByObjectTypeByPipelineIdStagesByStageId`: PipelineStage
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetPipelinesByObjectTypeByPipelineIdStagesByStageId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of object for which the pipeline stage is being retrieved. | 
**pipelineId** | **string** | The unique identifier of the pipeline that contains the stage. | 
**stageId** | **string** | The unique identifier of the stage to retrieve within the pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelinesByObjectTypeByPipelineIdStagesByStageIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PipelineStage**](PipelineStage.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelinesStages

> CollectionResponsePipelineStageNoPaging GetPipelinesStages(ctx, objectType, pipelineId).Execute()

List stages



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
	objectType := "objectType_example" // string | The type of object for which the pipeline stages are being retrieved, such as deals or tickets.
	pipelineId := "pipelineId_example" // string | The unique identifier of the pipeline whose stages are being retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetPipelinesStages(context.Background(), objectType, pipelineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetPipelinesStages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPipelinesStages`: CollectionResponsePipelineStageNoPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetPipelinesStages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of object for which the pipeline stages are being retrieved, such as deals or tickets. | 
**pipelineId** | **string** | The unique identifier of the pipeline whose stages are being retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelinesStagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CollectionResponsePipelineStageNoPaging**](CollectionResponsePipelineStageNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelinesStagesAudit

> CollectionResponsePublicAuditInfoNoPaging GetPipelinesStagesAudit(ctx, objectType, pipelineId, stageId).Execute()

Retrieve stage audit



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
	objectType := "objectType_example" // string | The type of CRM object that the pipeline belongs to, such as deals or tickets.
	pipelineId := "pipelineId_example" // string | 
	stageId := "stageId_example" // string | The unique identifier of the stage within the pipeline for which audit information is being retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetPipelinesStagesAudit(context.Background(), objectType, pipelineId, stageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetPipelinesStagesAudit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPipelinesStagesAudit`: CollectionResponsePublicAuditInfoNoPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetPipelinesStagesAudit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of CRM object that the pipeline belongs to, such as deals or tickets. | 
**pipelineId** | **string** |  | 
**stageId** | **string** | The unique identifier of the stage within the pipeline for which audit information is being retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelinesStagesAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CollectionResponsePublicAuditInfoNoPaging**](CollectionResponsePublicAuditInfoNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPipelines

> Pipeline PatchPipelines(ctx, objectType, pipelineId).PipelinePatchInput(pipelinePatchInput).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).Execute()

Update pipeline (Partial)



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
	objectType := "objectType_example" // string | The type of object associated with the pipeline, such as deals or tickets.
	pipelineId := "pipelineId_example" // string | The unique identifier of the pipeline to update.
	pipelinePatchInput := *openapiclient.NewPipelinePatchInput() // PipelinePatchInput | 
	validateDealStageUsagesBeforeDelete := true // bool | A boolean indicating whether to validate deal stage usages before deleting. Defaults to true. (optional) (default to true)
	validateReferencesBeforeDelete := true // bool | A boolean indicating whether to validate references before deleting. Defaults to true. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PatchPipelines(context.Background(), objectType, pipelineId).PipelinePatchInput(pipelinePatchInput).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PatchPipelines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchPipelines`: Pipeline
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PatchPipelines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of object associated with the pipeline, such as deals or tickets. | 
**pipelineId** | **string** | The unique identifier of the pipeline to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPipelinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pipelinePatchInput** | [**PipelinePatchInput**](PipelinePatchInput.md) |  | 
 **validateDealStageUsagesBeforeDelete** | **bool** | A boolean indicating whether to validate deal stage usages before deleting. Defaults to true. | [default to true]
 **validateReferencesBeforeDelete** | **bool** | A boolean indicating whether to validate references before deleting. Defaults to true. | [default to true]

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPipelinesStages

> PipelineStage PatchPipelinesStages(ctx, objectType, pipelineId, stageId).PipelineStagePatchInput(pipelineStagePatchInput).Execute()

Update stage (Partial)



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
	objectType := "objectType_example" // string | The type of object that the pipeline belongs to, such as deals or tickets.
	pipelineId := "pipelineId_example" // string | The unique identifier of the pipeline that contains the stage to be updated.
	stageId := "stageId_example" // string | The unique identifier of the stage to update within the specified pipeline.
	pipelineStagePatchInput := *openapiclient.NewPipelineStagePatchInput(map[string]string{"key": "Inner_example"}) // PipelineStagePatchInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PatchPipelinesStages(context.Background(), objectType, pipelineId, stageId).PipelineStagePatchInput(pipelineStagePatchInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PatchPipelinesStages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchPipelinesStages`: PipelineStage
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PatchPipelinesStages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of object that the pipeline belongs to, such as deals or tickets. | 
**pipelineId** | **string** | The unique identifier of the pipeline that contains the stage to be updated. | 
**stageId** | **string** | The unique identifier of the stage to update within the specified pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPipelinesStagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pipelineStagePatchInput** | [**PipelineStagePatchInput**](PipelineStagePatchInput.md) |  | 

### Return type

[**PipelineStage**](PipelineStage.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPipelines

> Pipeline PostPipelines(ctx, objectType).PipelineInput(pipelineInput).Execute()

Create pipeline



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
	objectType := "objectType_example" // string | The type of object for which the pipeline is being created, such as deals or tickets.
	pipelineInput := *openapiclient.NewPipelineInput(int32(123), "Label_example", []openapiclient.PipelineStageInput{*openapiclient.NewPipelineStageInput(int32(123), "Label_example", map[string]string{"key": "Inner_example"})}) // PipelineInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PostPipelines(context.Background(), objectType).PipelineInput(pipelineInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostPipelines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPipelines`: Pipeline
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PostPipelines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of object for which the pipeline is being created, such as deals or tickets. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPipelinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pipelineInput** | [**PipelineInput**](PipelineInput.md) |  | 

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPipelinesStages

> PipelineStage PostPipelinesStages(ctx, objectType, pipelineId).PipelineStageInput(pipelineStageInput).Execute()

Create Stage



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
	objectType := "objectType_example" // string | The type of object for which the pipeline stage is being created, such as deals or tickets.
	pipelineId := "pipelineId_example" // string | The unique identifier of the pipeline to which the new stage will be added.
	pipelineStageInput := *openapiclient.NewPipelineStageInput(int32(123), "Label_example", map[string]string{"key": "Inner_example"}) // PipelineStageInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PostPipelinesStages(context.Background(), objectType, pipelineId).PipelineStageInput(pipelineStageInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostPipelinesStages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPipelinesStages`: PipelineStage
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PostPipelinesStages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of object for which the pipeline stage is being created, such as deals or tickets. | 
**pipelineId** | **string** | The unique identifier of the pipeline to which the new stage will be added. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPipelinesStagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pipelineStageInput** | [**PipelineStageInput**](PipelineStageInput.md) |  | 

### Return type

[**PipelineStage**](PipelineStage.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPipelines

> Pipeline PutPipelines(ctx, objectType, pipelineId).PipelineReplaceInput(pipelineReplaceInput).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).Execute()

Update pipeline (Replace)



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
	objectType := "objectType_example" // string | The type of object the pipeline is associated with, such as deals or tickets.
	pipelineId := "pipelineId_example" // string | The unique identifier of the pipeline to update.
	pipelineReplaceInput := *openapiclient.NewPipelineReplaceInput(int32(123), "Label_example", []openapiclient.PipelineStageInput{*openapiclient.NewPipelineStageInput(int32(123), "Label_example", map[string]string{"key": "Inner_example"})}) // PipelineReplaceInput | 
	validateDealStageUsagesBeforeDelete := true // bool | A boolean indicating whether to validate deal stage usages before deleting. Defaults to true. (optional) (default to true)
	validateReferencesBeforeDelete := true // bool | A boolean indicating whether to validate references before deleting. Defaults to true. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PutPipelines(context.Background(), objectType, pipelineId).PipelineReplaceInput(pipelineReplaceInput).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PutPipelines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPipelines`: Pipeline
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PutPipelines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of object the pipeline is associated with, such as deals or tickets. | 
**pipelineId** | **string** | The unique identifier of the pipeline to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPipelinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pipelineReplaceInput** | [**PipelineReplaceInput**](PipelineReplaceInput.md) |  | 
 **validateDealStageUsagesBeforeDelete** | **bool** | A boolean indicating whether to validate deal stage usages before deleting. Defaults to true. | [default to true]
 **validateReferencesBeforeDelete** | **bool** | A boolean indicating whether to validate references before deleting. Defaults to true. | [default to true]

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPipelinesStages

> PipelineStage PutPipelinesStages(ctx, objectType, pipelineId, stageId).PipelineStageReplaceInput(pipelineStageReplaceInput).Execute()

Update stage



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
	objectType := "objectType_example" // string | The type of object the pipeline belongs to, such as deals or tickets.
	pipelineId := "pipelineId_example" // string | The unique identifier of the pipeline containing the stage.
	stageId := "stageId_example" // string | The unique identifier of the stage to update.
	pipelineStageReplaceInput := *openapiclient.NewPipelineStageReplaceInput(int32(123), "Label_example", map[string]string{"key": "Inner_example"}) // PipelineStageReplaceInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PutPipelinesStages(context.Background(), objectType, pipelineId, stageId).PipelineStageReplaceInput(pipelineStageReplaceInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PutPipelinesStages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPipelinesStages`: PipelineStage
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PutPipelinesStages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of object the pipeline belongs to, such as deals or tickets. | 
**pipelineId** | **string** | The unique identifier of the pipeline containing the stage. | 
**stageId** | **string** | The unique identifier of the stage to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPipelinesStagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pipelineStageReplaceInput** | [**PipelineStageReplaceInput**](PipelineStageReplaceInput.md) |  | 

### Return type

[**PipelineStage**](PipelineStage.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

