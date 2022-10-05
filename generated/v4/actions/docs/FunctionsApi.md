# \FunctionsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FunctionsArchive**](FunctionsApi.md#FunctionsArchive) | **Delete** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId} | Delete a custom action function
[**FunctionsArchiveByType**](FunctionsApi.md#FunctionsArchiveByType) | **Delete** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType} | Delete a custom action function
[**FunctionsCreateOrReplace**](FunctionsApi.md#FunctionsCreateOrReplace) | **Put** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId} | Create or replace a custom action function
[**FunctionsCreateOrReplaceByType**](FunctionsApi.md#FunctionsCreateOrReplaceByType) | **Put** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType} | Create or replace a custom action function
[**FunctionsGetByID**](FunctionsApi.md#FunctionsGetByID) | **Get** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId} | Get a custom action function
[**FunctionsGetByType**](FunctionsApi.md#FunctionsGetByType) | **Get** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType} | Get a custom action function
[**FunctionsGetPage**](FunctionsApi.md#FunctionsGetPage) | **Get** /automation/v4/actions/{appId}/{definitionId}/functions | Get all custom action functions



## FunctionsArchive

> FunctionsArchive(ctx, definitionId, functionType, functionId, appId).Execute()

Delete a custom action function



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
    functionType := "functionType_example" // string | The type of function. This determines when the function will be called.
    functionId := "functionId_example" // string | The ID qualifier for the function. This is used to specify which input field a function is associated with for `PRE_FETCH_OPTIONS` and `POST_FETCH_OPTIONS` function types.
    appId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsApi.FunctionsArchive(context.Background(), definitionId, functionType, functionId, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.FunctionsArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the custom workflow action | 
**functionType** | **string** | The type of function. This determines when the function will be called. | 
**functionId** | **string** | The ID qualifier for the function. This is used to specify which input field a function is associated with for &#x60;PRE_FETCH_OPTIONS&#x60; and &#x60;POST_FETCH_OPTIONS&#x60; function types. | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsArchiveByType

> FunctionsArchiveByType(ctx, definitionId, functionType, appId).Execute()

Delete a custom action function



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
    functionType := "functionType_example" // string | The type of function. This determines when the function will be called.
    appId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsApi.FunctionsArchiveByType(context.Background(), definitionId, functionType, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.FunctionsArchiveByType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the custom workflow action. | 
**functionType** | **string** | The type of function. This determines when the function will be called. | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsArchiveByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsCreateOrReplace

> ActionFunctionIdentifier FunctionsCreateOrReplace(ctx, definitionId, functionType, functionId, appId).Body(body).Execute()

Create or replace a custom action function



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
    functionType := "functionType_example" // string | The type of function. This determines when the function will be called.
    functionId := "functionId_example" // string | The ID qualifier for the function. This is used to specify which input field a function is associated with for `PRE_FETCH_OPTIONS` and `POST_FETCH_OPTIONS` function types.
    appId := int32(56) // int32 | 
    body := "body_example" // string | The function source code. Must be valid JavaScript code.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsApi.FunctionsCreateOrReplace(context.Background(), definitionId, functionType, functionId, appId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.FunctionsCreateOrReplace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FunctionsCreateOrReplace`: ActionFunctionIdentifier
    fmt.Fprintf(os.Stdout, "Response from `FunctionsApi.FunctionsCreateOrReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the custom workflow action. | 
**functionType** | **string** | The type of function. This determines when the function will be called. | 
**functionId** | **string** | The ID qualifier for the function. This is used to specify which input field a function is associated with for &#x60;PRE_FETCH_OPTIONS&#x60; and &#x60;POST_FETCH_OPTIONS&#x60; function types. | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsCreateOrReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **string** | The function source code. Must be valid JavaScript code. | 

### Return type

[**ActionFunctionIdentifier**](ActionFunctionIdentifier.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsCreateOrReplaceByType

> ActionFunctionIdentifier FunctionsCreateOrReplaceByType(ctx, definitionId, functionType, appId).Body(body).Execute()

Create or replace a custom action function



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
    functionType := "functionType_example" // string | The type of function. This determines when the function will be called.
    appId := int32(56) // int32 | 
    body := "body_example" // string | The function source code. Must be valid JavaScript code.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsApi.FunctionsCreateOrReplaceByType(context.Background(), definitionId, functionType, appId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.FunctionsCreateOrReplaceByType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FunctionsCreateOrReplaceByType`: ActionFunctionIdentifier
    fmt.Fprintf(os.Stdout, "Response from `FunctionsApi.FunctionsCreateOrReplaceByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the custom workflow action. | 
**functionType** | **string** | The type of function. This determines when the function will be called. | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsCreateOrReplaceByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **string** | The function source code. Must be valid JavaScript code. | 

### Return type

[**ActionFunctionIdentifier**](ActionFunctionIdentifier.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsGetByID

> ActionFunction FunctionsGetByID(ctx, definitionId, functionType, functionId, appId).Execute()

Get a custom action function



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
    functionType := "functionType_example" // string | The type of function. This determines when the function will be called.
    functionId := "functionId_example" // string | The ID qualifier for the function. This is used to specify which input field a function is associated with for `PRE_FETCH_OPTIONS` and `POST_FETCH_OPTIONS` function types.
    appId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsApi.FunctionsGetByID(context.Background(), definitionId, functionType, functionId, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.FunctionsGetByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FunctionsGetByID`: ActionFunction
    fmt.Fprintf(os.Stdout, "Response from `FunctionsApi.FunctionsGetByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the custom workflow action. | 
**functionType** | **string** | The type of function. This determines when the function will be called. | 
**functionId** | **string** | The ID qualifier for the function. This is used to specify which input field a function is associated with for &#x60;PRE_FETCH_OPTIONS&#x60; and &#x60;POST_FETCH_OPTIONS&#x60; function types. | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsGetByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ActionFunction**](ActionFunction.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsGetByType

> ActionFunction FunctionsGetByType(ctx, definitionId, functionType, appId).Execute()

Get a custom action function



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
    functionType := "functionType_example" // string | The type of function. This determines when the function will be called.
    appId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsApi.FunctionsGetByType(context.Background(), definitionId, functionType, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.FunctionsGetByType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FunctionsGetByType`: ActionFunction
    fmt.Fprintf(os.Stdout, "Response from `FunctionsApi.FunctionsGetByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the custom workflow action. | 
**functionType** | **string** | The type of function. This determines when the function will be called. | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsGetByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ActionFunction**](ActionFunction.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsGetPage

> CollectionResponseActionFunctionIdentifierNoPaging FunctionsGetPage(ctx, definitionId, appId).Execute()

Get all custom action functions



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
    appId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsApi.FunctionsGetPage(context.Background(), definitionId, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.FunctionsGetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FunctionsGetPage`: CollectionResponseActionFunctionIdentifierNoPaging
    fmt.Fprintf(os.Stdout, "Response from `FunctionsApi.FunctionsGetPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the custom workflow action. | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CollectionResponseActionFunctionIdentifierNoPaging**](CollectionResponseActionFunctionIdentifierNoPaging.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

