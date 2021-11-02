# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletecrmv3pipelinesobjectTypepipelineIdstagesstageIdArchive**](PipelineStagesApi.md#Deletecrmv3pipelinesobjectTypepipelineIdstagesstageIdArchive) | **Delete** /crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId} | Archive a pipeline stage
[**Getcrmv3pipelinesobjectTypepipelineIdstagesGetAll**](PipelineStagesApi.md#Getcrmv3pipelinesobjectTypepipelineIdstagesGetAll) | **Get** /crm/v3/pipelines/{objectType}/{pipelineId}/stages | Return all stages of a pipeline
[**Getcrmv3pipelinesobjectTypepipelineIdstagesstageIdGetById**](PipelineStagesApi.md#Getcrmv3pipelinesobjectTypepipelineIdstagesstageIdGetById) | **Get** /crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId} | Return a pipeline stage by ID
[**Patchcrmv3pipelinesobjectTypepipelineIdstagesstageIdUpdate**](PipelineStagesApi.md#Patchcrmv3pipelinesobjectTypepipelineIdstagesstageIdUpdate) | **Patch** /crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId} | Update a pipeline stage
[**Postcrmv3pipelinesobjectTypepipelineIdstagesCreate**](PipelineStagesApi.md#Postcrmv3pipelinesobjectTypepipelineIdstagesCreate) | **Post** /crm/v3/pipelines/{objectType}/{pipelineId}/stages | Create a pipeline stage
[**Putcrmv3pipelinesobjectTypepipelineIdstagesstageIdReplace**](PipelineStagesApi.md#Putcrmv3pipelinesobjectTypepipelineIdstagesstageIdReplace) | **Put** /crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId} | Replace a pipeline stage

# **Deletecrmv3pipelinesobjectTypepipelineIdstagesstageIdArchive**
> Deletecrmv3pipelinesobjectTypepipelineIdstagesstageIdArchive(ctx, objectType, pipelineId, stageId)
Archive a pipeline stage

Archive the pipeline stage identified by `{stageId}` associated with the pipeline identified by `{pipelineId}`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **pipelineId** | **string**|  | 
  **stageId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcrmv3pipelinesobjectTypepipelineIdstagesGetAll**
> CollectionResponsePipelineStage Getcrmv3pipelinesobjectTypepipelineIdstagesGetAll(ctx, objectType, pipelineId, optional)
Return all stages of a pipeline

Return all the stages associated with the pipeline identified by `{pipelineId}`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **pipelineId** | **string**|  | 
 **optional** | ***PipelineStagesApiGetcrmv3pipelinesobjectTypepipelineIdstagesGetAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineStagesApiGetcrmv3pipelinesobjectTypepipelineIdstagesGetAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **archived** | **optional.Bool**| Whether to return only results that have been archived. | [default to false]

### Return type

[**CollectionResponsePipelineStage**](CollectionResponsePipelineStage.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcrmv3pipelinesobjectTypepipelineIdstagesstageIdGetById**
> PipelineStage Getcrmv3pipelinesobjectTypepipelineIdstagesstageIdGetById(ctx, objectType, pipelineId, stageId, optional)
Return a pipeline stage by ID

Return the stage identified by `{stageId}` associated with the pipeline identified by `{pipelineId}`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **pipelineId** | **string**|  | 
  **stageId** | **string**|  | 
 **optional** | ***PipelineStagesApiGetcrmv3pipelinesobjectTypepipelineIdstagesstageIdGetByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineStagesApiGetcrmv3pipelinesobjectTypepipelineIdstagesstageIdGetByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **archived** | **optional.Bool**| Whether to return only results that have been archived. | [default to false]

### Return type

[**PipelineStage**](PipelineStage.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patchcrmv3pipelinesobjectTypepipelineIdstagesstageIdUpdate**
> PipelineStage Patchcrmv3pipelinesobjectTypepipelineIdstagesstageIdUpdate(ctx, objectType, pipelineId, stageId, optional)
Update a pipeline stage

Perform a partial update of the pipeline stage identified by `{stageId}` associated with the pipeline identified by `{pipelineId}`. Any properties not included in this update will keep their existing values. The updated stage will be returned in the response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **pipelineId** | **string**|  | 
  **stageId** | **string**|  | 
 **optional** | ***PipelineStagesApiPatchcrmv3pipelinesobjectTypepipelineIdstagesstageIdUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineStagesApiPatchcrmv3pipelinesobjectTypepipelineIdstagesstageIdUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of PipelineStagePatchInput**](PipelineStagePatchInput.md)|  | 
 **archived** | **optional.**| Whether to return only results that have been archived. | [default to false]

### Return type

[**PipelineStage**](PipelineStage.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3pipelinesobjectTypepipelineIdstagesCreate**
> PipelineStage Postcrmv3pipelinesobjectTypepipelineIdstagesCreate(ctx, objectType, pipelineId, optional)
Create a pipeline stage

Create a new stage associated with the pipeline identified by `{pipelineId}`. The entire stage object, including its unique ID, will be returned in the response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **pipelineId** | **string**|  | 
 **optional** | ***PipelineStagesApiPostcrmv3pipelinesobjectTypepipelineIdstagesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineStagesApiPostcrmv3pipelinesobjectTypepipelineIdstagesCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PipelineStageInput**](PipelineStageInput.md)|  | 

### Return type

[**PipelineStage**](PipelineStage.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Putcrmv3pipelinesobjectTypepipelineIdstagesstageIdReplace**
> PipelineStage Putcrmv3pipelinesobjectTypepipelineIdstagesstageIdReplace(ctx, objectType, pipelineId, stageId, optional)
Replace a pipeline stage

Replace all the properties of an existing pipeline stage with the values provided. The updated stage will be returned in the response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **pipelineId** | **string**|  | 
  **stageId** | **string**|  | 
 **optional** | ***PipelineStagesApiPutcrmv3pipelinesobjectTypepipelineIdstagesstageIdReplaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineStagesApiPutcrmv3pipelinesobjectTypepipelineIdstagesstageIdReplaceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of PipelineStageInput**](PipelineStageInput.md)|  | 

### Return type

[**PipelineStage**](PipelineStage.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

