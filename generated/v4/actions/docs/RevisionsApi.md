# \RevisionsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RevisionsGetByID**](RevisionsApi.md#RevisionsGetByID) | **Get** /automation/v4/actions/{appId}/{definitionId}/revisions/{revisionId} | Get a revision for a custom action
[**RevisionsGetPage**](RevisionsApi.md#RevisionsGetPage) | **Get** /automation/v4/actions/{appId}/{definitionId}/revisions | Get all revisions for a custom action



## RevisionsGetByID

> ActionRevision RevisionsGetByID(ctx, definitionId, revisionId, appId).Execute()

Get a revision for a custom action



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
    definitionId := "definitionId_example" // string | The ID of the custom workflow action.
    revisionId := "revisionId_example" // string | The version of the custom workflow action.
    appId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RevisionsApi.RevisionsGetByID(context.Background(), definitionId, revisionId, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionsApi.RevisionsGetByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevisionsGetByID`: ActionRevision
    fmt.Fprintf(os.Stdout, "Response from `RevisionsApi.RevisionsGetByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the custom workflow action. | 
**revisionId** | **string** | The version of the custom workflow action. | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevisionsGetByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ActionRevision**](ActionRevision.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevisionsGetPage

> CollectionResponseActionRevisionForwardPaging RevisionsGetPage(ctx, definitionId, appId).Limit(limit).After(after).Execute()

Get all revisions for a custom action



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
    definitionId := "definitionId_example" // string | The ID of the custom workflow action
    appId := int32(56) // int32 | 
    limit := int32(56) // int32 | Maximum number of results per page. (optional)
    after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RevisionsApi.RevisionsGetPage(context.Background(), definitionId, appId).Limit(limit).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionsApi.RevisionsGetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevisionsGetPage`: CollectionResponseActionRevisionForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `RevisionsApi.RevisionsGetPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the custom workflow action | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevisionsGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Maximum number of results per page. | 
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 

### Return type

[**CollectionResponseActionRevisionForwardPaging**](CollectionResponseActionRevisionForwardPaging.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

