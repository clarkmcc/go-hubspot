# \DefinitionsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Archive**](DefinitionsApi.md#Archive) | **Delete** /automation/v4/actions/{appId}/{definitionId} | Archive an extension definition
[**Create**](DefinitionsApi.md#Create) | **Post** /automation/v4/actions/{appId} | Create a new extension definition
[**GetByID**](DefinitionsApi.md#GetByID) | **Get** /automation/v4/actions/{appId}/{definitionId} | Get extension definition by Id
[**GetPage**](DefinitionsApi.md#GetPage) | **Get** /automation/v4/actions/{appId} | Get paged extension definitions
[**Update**](DefinitionsApi.md#Update) | **Patch** /automation/v4/actions/{appId}/{definitionId} | Patch an existing extension definition



## Archive

> Archive(ctx, definitionId, appId).Execute()

Archive an extension definition

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
    definitionId := "definitionId_example" // string | 
    appId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefinitionsApi.Archive(context.Background(), definitionId, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefinitionsApi.Archive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> PublicActionDefinition Create(ctx, appId).PublicActionDefinitionEgg(publicActionDefinitionEgg).Execute()

Create a new extension definition

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
    appId := int32(56) // int32 | 
    publicActionDefinitionEgg := *openapiclient.NewPublicActionDefinitionEgg([]openapiclient.InputFieldDefinition{*openapiclient.NewInputFieldDefinition(false, *openapiclient.NewFieldTypeDefinition("Name_example", []openapiclient.Option{*openapiclient.NewOption(false, int32(123), float32(123), "Description_example", false, "Label_example", "Value_example")}, "Type_example", false))}, []openapiclient.PublicActionFunction{*openapiclient.NewPublicActionFunction("FunctionSource_example", "FunctionType_example")}, "ActionUrl_example", false, []string{"ObjectTypes_example"}, map[string]PublicActionLabels{"key": *openapiclient.NewPublicActionLabels("ActionName_example")}) // PublicActionDefinitionEgg | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefinitionsApi.Create(context.Background(), appId).PublicActionDefinitionEgg(publicActionDefinitionEgg).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefinitionsApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: PublicActionDefinition
    fmt.Fprintf(os.Stdout, "Response from `DefinitionsApi.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicActionDefinitionEgg** | [**PublicActionDefinitionEgg**](PublicActionDefinitionEgg.md) |  | 

### Return type

[**PublicActionDefinition**](PublicActionDefinition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByID

> PublicActionDefinition GetByID(ctx, definitionId, appId).Archived(archived).Execute()

Get extension definition by Id

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
    definitionId := "definitionId_example" // string | 
    appId := int32(56) // int32 | 
    archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefinitionsApi.GetByID(context.Background(), definitionId, appId).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefinitionsApi.GetByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetByID`: PublicActionDefinition
    fmt.Fprintf(os.Stdout, "Response from `DefinitionsApi.GetByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

[**PublicActionDefinition**](PublicActionDefinition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPage

> CollectionResponsePublicActionDefinitionForwardPaging GetPage(ctx, appId).Limit(limit).After(after).Archived(archived).Execute()

Get paged extension definitions

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
    appId := int32(56) // int32 | 
    limit := int32(56) // int32 | The maximum number of results to display per page. (optional)
    after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
    archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefinitionsApi.GetPage(context.Background(), appId).Limit(limit).After(after).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefinitionsApi.GetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPage`: CollectionResponsePublicActionDefinitionForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `DefinitionsApi.GetPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of results to display per page. | 
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

[**CollectionResponsePublicActionDefinitionForwardPaging**](CollectionResponsePublicActionDefinitionForwardPaging.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> PublicActionDefinition Update(ctx, definitionId, appId).PublicActionDefinitionPatch(publicActionDefinitionPatch).Execute()

Patch an existing extension definition

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
    definitionId := "definitionId_example" // string | 
    appId := int32(56) // int32 | 
    publicActionDefinitionPatch := *openapiclient.NewPublicActionDefinitionPatch() // PublicActionDefinitionPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefinitionsApi.Update(context.Background(), definitionId, appId).PublicActionDefinitionPatch(publicActionDefinitionPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefinitionsApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: PublicActionDefinition
    fmt.Fprintf(os.Stdout, "Response from `DefinitionsApi.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **publicActionDefinitionPatch** | [**PublicActionDefinitionPatch**](PublicActionDefinitionPatch.md) |  | 

### Return type

[**PublicActionDefinition**](PublicActionDefinition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

