# \FunctionsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeArchiveByFunctionType**](FunctionsApi.md#DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeArchiveByFunctionType) | **Delete** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType} | Delete a custom action function
[**DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdArchive**](FunctionsApi.md#DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdArchive) | **Delete** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId} | Delete a custom action function
[**GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdGetById**](FunctionsApi.md#GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdGetById) | **Get** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId} | Get a custom action function
[**GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeGetByFunctionType**](FunctionsApi.md#GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeGetByFunctionType) | **Get** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType} | Get a custom action function
[**GetAutomationV4ActionsAppIdDefinitionIdFunctionsGetPage**](FunctionsApi.md#GetAutomationV4ActionsAppIdDefinitionIdFunctionsGetPage) | **Get** /automation/v4/actions/{appId}/{definitionId}/functions | Get all custom action functions
[**PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeCreateOrReplaceByFunctionType**](FunctionsApi.md#PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeCreateOrReplaceByFunctionType) | **Put** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType} | Create or replace a custom action function
[**PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdCreateOrReplace**](FunctionsApi.md#PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdCreateOrReplace) | **Put** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId} | Create or replace a custom action function



## DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeArchiveByFunctionType

> DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeArchiveByFunctionType(ctx, definitionId, functionType, appId).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FunctionsApi.DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeArchiveByFunctionType(context.Background(), definitionId, functionType, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeArchiveByFunctionType``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeArchiveByFunctionTypeRequest struct via the builder pattern


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


## DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdArchive

> DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdArchive(ctx, definitionId, functionType, functionId, appId).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FunctionsApi.DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdArchive(context.Background(), definitionId, functionType, functionId, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdArchiveRequest struct via the builder pattern


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


## GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdGetById

> ActionFunction GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdGetById(ctx, definitionId, functionType, functionId, appId).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FunctionsApi.GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdGetById(context.Background(), definitionId, functionType, functionId, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdGetById`: ActionFunction
    fmt.Fprintf(os.Stdout, "Response from `FunctionsApi.GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdGetById`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdGetByIdRequest struct via the builder pattern


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


## GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeGetByFunctionType

> ActionFunction GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeGetByFunctionType(ctx, definitionId, functionType, appId).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FunctionsApi.GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeGetByFunctionType(context.Background(), definitionId, functionType, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeGetByFunctionType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeGetByFunctionType`: ActionFunction
    fmt.Fprintf(os.Stdout, "Response from `FunctionsApi.GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeGetByFunctionType`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeGetByFunctionTypeRequest struct via the builder pattern


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


## GetAutomationV4ActionsAppIdDefinitionIdFunctionsGetPage

> CollectionResponseActionFunctionIdentifierNoPaging GetAutomationV4ActionsAppIdDefinitionIdFunctionsGetPage(ctx, definitionId, appId).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FunctionsApi.GetAutomationV4ActionsAppIdDefinitionIdFunctionsGetPage(context.Background(), definitionId, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.GetAutomationV4ActionsAppIdDefinitionIdFunctionsGetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutomationV4ActionsAppIdDefinitionIdFunctionsGetPage`: CollectionResponseActionFunctionIdentifierNoPaging
    fmt.Fprintf(os.Stdout, "Response from `FunctionsApi.GetAutomationV4ActionsAppIdDefinitionIdFunctionsGetPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the custom workflow action. | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationV4ActionsAppIdDefinitionIdFunctionsGetPageRequest struct via the builder pattern


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


## PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeCreateOrReplaceByFunctionType

> ActionFunctionIdentifier PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeCreateOrReplaceByFunctionType(ctx, definitionId, functionType, appId).Body(body).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FunctionsApi.PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeCreateOrReplaceByFunctionType(context.Background(), definitionId, functionType, appId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeCreateOrReplaceByFunctionType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeCreateOrReplaceByFunctionType`: ActionFunctionIdentifier
    fmt.Fprintf(os.Stdout, "Response from `FunctionsApi.PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeCreateOrReplaceByFunctionType`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeCreateOrReplaceByFunctionTypeRequest struct via the builder pattern


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


## PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdCreateOrReplace

> ActionFunctionIdentifier PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdCreateOrReplace(ctx, definitionId, functionType, functionId, appId).Body(body).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FunctionsApi.PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdCreateOrReplace(context.Background(), definitionId, functionType, functionId, appId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdCreateOrReplace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdCreateOrReplace`: ActionFunctionIdentifier
    fmt.Fprintf(os.Stdout, "Response from `FunctionsApi.PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdCreateOrReplace`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdCreateOrReplaceRequest struct via the builder pattern


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

