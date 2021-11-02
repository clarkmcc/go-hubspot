# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletecrmv3pipelinesobjectTypepipelineIdArchive**](PipelinesApi.md#Deletecrmv3pipelinesobjectTypepipelineIdArchive) | **Delete** /crm/v3/pipelines/{objectType}/{pipelineId} | Archive a pipeline
[**Getcrmv3pipelinesobjectTypeGetAll**](PipelinesApi.md#Getcrmv3pipelinesobjectTypeGetAll) | **Get** /crm/v3/pipelines/{objectType} | Retrieve all pipelines
[**Getcrmv3pipelinesobjectTypepipelineIdGetById**](PipelinesApi.md#Getcrmv3pipelinesobjectTypepipelineIdGetById) | **Get** /crm/v3/pipelines/{objectType}/{pipelineId} | Return a pipeline by ID
[**Patchcrmv3pipelinesobjectTypepipelineIdUpdate**](PipelinesApi.md#Patchcrmv3pipelinesobjectTypepipelineIdUpdate) | **Patch** /crm/v3/pipelines/{objectType}/{pipelineId} | Update a pipeline
[**Postcrmv3pipelinesobjectTypeCreate**](PipelinesApi.md#Postcrmv3pipelinesobjectTypeCreate) | **Post** /crm/v3/pipelines/{objectType} | Create a pipeline
[**Putcrmv3pipelinesobjectTypepipelineIdReplace**](PipelinesApi.md#Putcrmv3pipelinesobjectTypepipelineIdReplace) | **Put** /crm/v3/pipelines/{objectType}/{pipelineId} | Replace a pipeline

# **Deletecrmv3pipelinesobjectTypepipelineIdArchive**
> Deletecrmv3pipelinesobjectTypepipelineIdArchive(ctx, objectType, pipelineId)
Archive a pipeline

Archive the pipeline identified by `{pipelineId}`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **pipelineId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcrmv3pipelinesobjectTypeGetAll**
> CollectionResponsePipeline Getcrmv3pipelinesobjectTypeGetAll(ctx, objectType, optional)
Retrieve all pipelines

Return all pipelines for the object type specified by `{objectType}`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
 **optional** | ***PipelinesApiGetcrmv3pipelinesobjectTypeGetAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelinesApiGetcrmv3pipelinesobjectTypeGetAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **optional.Bool**| Whether to return only results that have been archived. | [default to false]

### Return type

[**CollectionResponsePipeline**](CollectionResponsePipeline.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcrmv3pipelinesobjectTypepipelineIdGetById**
> Pipeline Getcrmv3pipelinesobjectTypepipelineIdGetById(ctx, objectType, pipelineId, optional)
Return a pipeline by ID

Return a single pipeline object identified by its unique `{pipelineId}`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **pipelineId** | **string**|  | 
 **optional** | ***PipelinesApiGetcrmv3pipelinesobjectTypepipelineIdGetByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelinesApiGetcrmv3pipelinesobjectTypepipelineIdGetByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **archived** | **optional.Bool**| Whether to return only results that have been archived. | [default to false]

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patchcrmv3pipelinesobjectTypepipelineIdUpdate**
> Pipeline Patchcrmv3pipelinesobjectTypepipelineIdUpdate(ctx, objectType, pipelineId, optional)
Update a pipeline

Perform a partial update of the pipeline identified by `{pipelineId}`. The updated pipeline will be returned in the response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **pipelineId** | **string**|  | 
 **optional** | ***PipelinesApiPatchcrmv3pipelinesobjectTypepipelineIdUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelinesApiPatchcrmv3pipelinesobjectTypepipelineIdUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PipelinePatchInput**](PipelinePatchInput.md)|  | 
 **archived** | **optional.**| Whether to return only results that have been archived. | [default to false]

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Postcrmv3pipelinesobjectTypeCreate**
> Pipeline Postcrmv3pipelinesobjectTypeCreate(ctx, objectType, optional)
Create a pipeline

Create a new pipeline with the provided property values. The entire pipeline object, including its unique ID, will be returned in the response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
 **optional** | ***PipelinesApiPostcrmv3pipelinesobjectTypeCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelinesApiPostcrmv3pipelinesobjectTypeCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PipelineInput**](PipelineInput.md)|  | 

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Putcrmv3pipelinesobjectTypepipelineIdReplace**
> Pipeline Putcrmv3pipelinesobjectTypepipelineIdReplace(ctx, objectType, pipelineId, optional)
Replace a pipeline

Replace all the properties of an existing pipeline with the values provided. This will overwrite any existing pipeline stages. The updated pipeline will be returned in the response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **pipelineId** | **string**|  | 
 **optional** | ***PipelinesApiPutcrmv3pipelinesobjectTypepipelineIdReplaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelinesApiPutcrmv3pipelinesobjectTypepipelineIdReplaceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PipelineInput**](PipelineInput.md)|  | 

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

