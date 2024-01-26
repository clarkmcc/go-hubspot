# \BatchApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmV3ObjectsFeedbackSubmissionsBatchArchiveArchive**](BatchApi.md#PostCrmV3ObjectsFeedbackSubmissionsBatchArchiveArchive) | **Post** /crm/v3/objects/feedback_submissions/batch/archive | Archive a batch of feedback submissions by ID
[**PostCrmV3ObjectsFeedbackSubmissionsBatchCreateCreate**](BatchApi.md#PostCrmV3ObjectsFeedbackSubmissionsBatchCreateCreate) | **Post** /crm/v3/objects/feedback_submissions/batch/create | Create a batch of feedback submissions
[**PostCrmV3ObjectsFeedbackSubmissionsBatchReadRead**](BatchApi.md#PostCrmV3ObjectsFeedbackSubmissionsBatchReadRead) | **Post** /crm/v3/objects/feedback_submissions/batch/read | Read a batch of feedback submissions by internal ID, or unique property values
[**PostCrmV3ObjectsFeedbackSubmissionsBatchUpdateUpdate**](BatchApi.md#PostCrmV3ObjectsFeedbackSubmissionsBatchUpdateUpdate) | **Post** /crm/v3/objects/feedback_submissions/batch/update | Update a batch of feedback submissions



## PostCrmV3ObjectsFeedbackSubmissionsBatchArchiveArchive

> PostCrmV3ObjectsFeedbackSubmissionsBatchArchiveArchive(ctx).BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId).Execute()

Archive a batch of feedback submissions by ID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    batchInputSimplePublicObjectId := *openapiclient.NewBatchInputSimplePublicObjectId([]openapiclient.SimplePublicObjectId{*openapiclient.NewSimplePublicObjectId("Id_example")}) // BatchInputSimplePublicObjectId | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostCrmV3ObjectsFeedbackSubmissionsBatchArchiveArchive(context.Background()).BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV3ObjectsFeedbackSubmissionsBatchArchiveArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsFeedbackSubmissionsBatchArchiveArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputSimplePublicObjectId** | [**BatchInputSimplePublicObjectId**](BatchInputSimplePublicObjectId.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ObjectsFeedbackSubmissionsBatchCreateCreate

> BatchResponseSimplePublicObject PostCrmV3ObjectsFeedbackSubmissionsBatchCreateCreate(ctx).BatchInputSimplePublicObjectInputForCreate(batchInputSimplePublicObjectInputForCreate).Execute()

Create a batch of feedback submissions

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    batchInputSimplePublicObjectInputForCreate := *openapiclient.NewBatchInputSimplePublicObjectInputForCreate([]openapiclient.SimplePublicObjectInputForCreate{*openapiclient.NewSimplePublicObjectInputForCreate([]openapiclient.PublicAssociationsForObject{*openapiclient.NewPublicAssociationsForObject([]openapiclient.AssociationSpec{*openapiclient.NewAssociationSpec("AssociationCategory_example", int32(123))}, *openapiclient.NewPublicObjectId("Id_example"))}, map[string]string{"key": "Inner_example"})}) // BatchInputSimplePublicObjectInputForCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostCrmV3ObjectsFeedbackSubmissionsBatchCreateCreate(context.Background()).BatchInputSimplePublicObjectInputForCreate(batchInputSimplePublicObjectInputForCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV3ObjectsFeedbackSubmissionsBatchCreateCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ObjectsFeedbackSubmissionsBatchCreateCreate`: BatchResponseSimplePublicObject
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.PostCrmV3ObjectsFeedbackSubmissionsBatchCreateCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsFeedbackSubmissionsBatchCreateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputSimplePublicObjectInputForCreate** | [**BatchInputSimplePublicObjectInputForCreate**](BatchInputSimplePublicObjectInputForCreate.md) |  | 

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ObjectsFeedbackSubmissionsBatchReadRead

> BatchResponseSimplePublicObject PostCrmV3ObjectsFeedbackSubmissionsBatchReadRead(ctx).BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId).Archived(archived).Execute()

Read a batch of feedback submissions by internal ID, or unique property values

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    batchReadInputSimplePublicObjectId := *openapiclient.NewBatchReadInputSimplePublicObjectId([]string{"PropertiesWithHistory_example"}, []openapiclient.SimplePublicObjectId{*openapiclient.NewSimplePublicObjectId("Id_example")}, []string{"Properties_example"}) // BatchReadInputSimplePublicObjectId | 
    archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostCrmV3ObjectsFeedbackSubmissionsBatchReadRead(context.Background()).BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV3ObjectsFeedbackSubmissionsBatchReadRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ObjectsFeedbackSubmissionsBatchReadRead`: BatchResponseSimplePublicObject
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.PostCrmV3ObjectsFeedbackSubmissionsBatchReadRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsFeedbackSubmissionsBatchReadReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchReadInputSimplePublicObjectId** | [**BatchReadInputSimplePublicObjectId**](BatchReadInputSimplePublicObjectId.md) |  | 
 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ObjectsFeedbackSubmissionsBatchUpdateUpdate

> BatchResponseSimplePublicObject PostCrmV3ObjectsFeedbackSubmissionsBatchUpdateUpdate(ctx).BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput).Execute()

Update a batch of feedback submissions

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    batchInputSimplePublicObjectBatchInput := *openapiclient.NewBatchInputSimplePublicObjectBatchInput([]openapiclient.SimplePublicObjectBatchInput{*openapiclient.NewSimplePublicObjectBatchInput("Id_example", map[string]string{"key": "Inner_example"})}) // BatchInputSimplePublicObjectBatchInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostCrmV3ObjectsFeedbackSubmissionsBatchUpdateUpdate(context.Background()).BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV3ObjectsFeedbackSubmissionsBatchUpdateUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ObjectsFeedbackSubmissionsBatchUpdateUpdate`: BatchResponseSimplePublicObject
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.PostCrmV3ObjectsFeedbackSubmissionsBatchUpdateUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsFeedbackSubmissionsBatchUpdateUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputSimplePublicObjectBatchInput** | [**BatchInputSimplePublicObjectBatchInput**](BatchInputSimplePublicObjectBatchInput.md) |  | 

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

