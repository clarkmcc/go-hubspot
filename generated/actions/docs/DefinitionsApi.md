# \DefinitionsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAutomationV4ActionsAppIdDefinitionIdArchive**](DefinitionsApi.md#DeleteAutomationV4ActionsAppIdDefinitionIdArchive) | **Delete** /automation/v4/actions/{appId}/{definitionId} | Archive a custom action
[**GetAutomationV4ActionsAppIdDefinitionIdGetById**](DefinitionsApi.md#GetAutomationV4ActionsAppIdDefinitionIdGetById) | **Get** /automation/v4/actions/{appId}/{definitionId} | Get a custom action
[**GetAutomationV4ActionsAppIdGetPage**](DefinitionsApi.md#GetAutomationV4ActionsAppIdGetPage) | **Get** /automation/v4/actions/{appId} | Get all custom actions
[**PatchAutomationV4ActionsAppIdDefinitionIdUpdate**](DefinitionsApi.md#PatchAutomationV4ActionsAppIdDefinitionIdUpdate) | **Patch** /automation/v4/actions/{appId}/{definitionId} | Update a custom action
[**PostAutomationV4ActionsAppIdCreate**](DefinitionsApi.md#PostAutomationV4ActionsAppIdCreate) | **Post** /automation/v4/actions/{appId} | Create new custom action



## DeleteAutomationV4ActionsAppIdDefinitionIdArchive

> DeleteAutomationV4ActionsAppIdDefinitionIdArchive(ctx, definitionId, appId).Execute()

Archive a custom action



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
    resp, r, err := apiClient.DefinitionsApi.DeleteAutomationV4ActionsAppIdDefinitionIdArchive(context.Background(), definitionId, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefinitionsApi.DeleteAutomationV4ActionsAppIdDefinitionIdArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the custom workflow action. | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutomationV4ActionsAppIdDefinitionIdArchiveRequest struct via the builder pattern


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


## GetAutomationV4ActionsAppIdDefinitionIdGetById

> ExtensionActionDefinition GetAutomationV4ActionsAppIdDefinitionIdGetById(ctx, definitionId, appId).Archived(archived).Execute()

Get a custom action



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
    archived := true // bool | Whether to include archived custom actions. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefinitionsApi.GetAutomationV4ActionsAppIdDefinitionIdGetById(context.Background(), definitionId, appId).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefinitionsApi.GetAutomationV4ActionsAppIdDefinitionIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutomationV4ActionsAppIdDefinitionIdGetById`: ExtensionActionDefinition
    fmt.Fprintf(os.Stdout, "Response from `DefinitionsApi.GetAutomationV4ActionsAppIdDefinitionIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the custom workflow action. | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationV4ActionsAppIdDefinitionIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **archived** | **bool** | Whether to include archived custom actions. | [default to false]

### Return type

[**ExtensionActionDefinition**](ExtensionActionDefinition.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomationV4ActionsAppIdGetPage

> CollectionResponseExtensionActionDefinitionForwardPaging GetAutomationV4ActionsAppIdGetPage(ctx, appId).Limit(limit).After(after).Archived(archived).Execute()

Get all custom actions



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
    limit := int32(56) // int32 | Maximum number of results per page. (optional)
    after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
    archived := true // bool | Whether to include archived custom actions. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefinitionsApi.GetAutomationV4ActionsAppIdGetPage(context.Background(), appId).Limit(limit).After(after).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefinitionsApi.GetAutomationV4ActionsAppIdGetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutomationV4ActionsAppIdGetPage`: CollectionResponseExtensionActionDefinitionForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `DefinitionsApi.GetAutomationV4ActionsAppIdGetPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationV4ActionsAppIdGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of results per page. | 
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **archived** | **bool** | Whether to include archived custom actions. | [default to false]

### Return type

[**CollectionResponseExtensionActionDefinitionForwardPaging**](CollectionResponseExtensionActionDefinitionForwardPaging.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAutomationV4ActionsAppIdDefinitionIdUpdate

> ExtensionActionDefinition PatchAutomationV4ActionsAppIdDefinitionIdUpdate(ctx, definitionId, appId).ExtensionActionDefinitionPatch(extensionActionDefinitionPatch).Execute()

Update a custom action



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
    extensionActionDefinitionPatch := *openapiclient.NewExtensionActionDefinitionPatch() // ExtensionActionDefinitionPatch | The custom workflow action fields to be updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefinitionsApi.PatchAutomationV4ActionsAppIdDefinitionIdUpdate(context.Background(), definitionId, appId).ExtensionActionDefinitionPatch(extensionActionDefinitionPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefinitionsApi.PatchAutomationV4ActionsAppIdDefinitionIdUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchAutomationV4ActionsAppIdDefinitionIdUpdate`: ExtensionActionDefinition
    fmt.Fprintf(os.Stdout, "Response from `DefinitionsApi.PatchAutomationV4ActionsAppIdDefinitionIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The ID of the custom workflow action. | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAutomationV4ActionsAppIdDefinitionIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **extensionActionDefinitionPatch** | [**ExtensionActionDefinitionPatch**](ExtensionActionDefinitionPatch.md) | The custom workflow action fields to be updated. | 

### Return type

[**ExtensionActionDefinition**](ExtensionActionDefinition.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutomationV4ActionsAppIdCreate

> ExtensionActionDefinition PostAutomationV4ActionsAppIdCreate(ctx, appId).ExtensionActionDefinitionInput(extensionActionDefinitionInput).Execute()

Create new custom action



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
    extensionActionDefinitionInput := *openapiclient.NewExtensionActionDefinitionInput([]openapiclient.ActionFunction{*openapiclient.NewActionFunction("FunctionSource_example", "FunctionType_example")}, "ActionUrl_example", false, []openapiclient.InputFieldDefinition{*openapiclient.NewInputFieldDefinition(*openapiclient.NewFieldTypeDefinition("Name_example", "Type_example", []openapiclient.Option{*openapiclient.NewOption("Label_example", "Value_example", int32(123), float32(123), false, "Description_example", false)}), false)}, map[string]ActionLabels{"key": *openapiclient.NewActionLabels("ActionName_example")}, []string{"ObjectTypes_example"}) // ExtensionActionDefinitionInput | The custom workflow action to create.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefinitionsApi.PostAutomationV4ActionsAppIdCreate(context.Background(), appId).ExtensionActionDefinitionInput(extensionActionDefinitionInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefinitionsApi.PostAutomationV4ActionsAppIdCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAutomationV4ActionsAppIdCreate`: ExtensionActionDefinition
    fmt.Fprintf(os.Stdout, "Response from `DefinitionsApi.PostAutomationV4ActionsAppIdCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutomationV4ActionsAppIdCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extensionActionDefinitionInput** | [**ExtensionActionDefinitionInput**](ExtensionActionDefinitionInput.md) | The custom workflow action to create. | 

### Return type

[**ExtensionActionDefinition**](ExtensionActionDefinition.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

