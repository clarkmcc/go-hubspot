# \GroupsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3PropertiesObjectTypeGroupsGroupNameArchive**](GroupsApi.md#DeleteCrmV3PropertiesObjectTypeGroupsGroupNameArchive) | **Delete** /crm/v3/properties/{objectType}/groups/{groupName} | Archive a property group
[**GetCrmV3PropertiesObjectTypeGroupsGetAll**](GroupsApi.md#GetCrmV3PropertiesObjectTypeGroupsGetAll) | **Get** /crm/v3/properties/{objectType}/groups | Read all property groups
[**GetCrmV3PropertiesObjectTypeGroupsGroupNameGetByName**](GroupsApi.md#GetCrmV3PropertiesObjectTypeGroupsGroupNameGetByName) | **Get** /crm/v3/properties/{objectType}/groups/{groupName} | Read a property group
[**PatchCrmV3PropertiesObjectTypeGroupsGroupNameUpdate**](GroupsApi.md#PatchCrmV3PropertiesObjectTypeGroupsGroupNameUpdate) | **Patch** /crm/v3/properties/{objectType}/groups/{groupName} | Update a property group
[**PostCrmV3PropertiesObjectTypeGroupsCreate**](GroupsApi.md#PostCrmV3PropertiesObjectTypeGroupsCreate) | **Post** /crm/v3/properties/{objectType}/groups | Create a property group



## DeleteCrmV3PropertiesObjectTypeGroupsGroupNameArchive

> DeleteCrmV3PropertiesObjectTypeGroupsGroupNameArchive(ctx, objectType, groupName).Execute()

Archive a property group



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
    objectType := "objectType_example" // string | 
    groupName := "groupName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.DeleteCrmV3PropertiesObjectTypeGroupsGroupNameArchive(context.Background(), objectType, groupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.DeleteCrmV3PropertiesObjectTypeGroupsGroupNameArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**groupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3PropertiesObjectTypeGroupsGroupNameArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3PropertiesObjectTypeGroupsGetAll

> CollectionResponsePropertyGroup GetCrmV3PropertiesObjectTypeGroupsGetAll(ctx, objectType).Execute()

Read all property groups



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
    objectType := "objectType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.GetCrmV3PropertiesObjectTypeGroupsGetAll(context.Background(), objectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetCrmV3PropertiesObjectTypeGroupsGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3PropertiesObjectTypeGroupsGetAll`: CollectionResponsePropertyGroup
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetCrmV3PropertiesObjectTypeGroupsGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3PropertiesObjectTypeGroupsGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionResponsePropertyGroup**](CollectionResponsePropertyGroup.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3PropertiesObjectTypeGroupsGroupNameGetByName

> PropertyGroup GetCrmV3PropertiesObjectTypeGroupsGroupNameGetByName(ctx, objectType, groupName).Execute()

Read a property group



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
    objectType := "objectType_example" // string | 
    groupName := "groupName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.GetCrmV3PropertiesObjectTypeGroupsGroupNameGetByName(context.Background(), objectType, groupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetCrmV3PropertiesObjectTypeGroupsGroupNameGetByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3PropertiesObjectTypeGroupsGroupNameGetByName`: PropertyGroup
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetCrmV3PropertiesObjectTypeGroupsGroupNameGetByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**groupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3PropertiesObjectTypeGroupsGroupNameGetByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PropertyGroup**](PropertyGroup.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCrmV3PropertiesObjectTypeGroupsGroupNameUpdate

> PropertyGroup PatchCrmV3PropertiesObjectTypeGroupsGroupNameUpdate(ctx, objectType, groupName).PropertyGroupUpdate(propertyGroupUpdate).Execute()

Update a property group



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
    objectType := "objectType_example" // string | 
    groupName := "groupName_example" // string | 
    propertyGroupUpdate := *openapiclient.NewPropertyGroupUpdate() // PropertyGroupUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.PatchCrmV3PropertiesObjectTypeGroupsGroupNameUpdate(context.Background(), objectType, groupName).PropertyGroupUpdate(propertyGroupUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.PatchCrmV3PropertiesObjectTypeGroupsGroupNameUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCrmV3PropertiesObjectTypeGroupsGroupNameUpdate`: PropertyGroup
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.PatchCrmV3PropertiesObjectTypeGroupsGroupNameUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**groupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCrmV3PropertiesObjectTypeGroupsGroupNameUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **propertyGroupUpdate** | [**PropertyGroupUpdate**](PropertyGroupUpdate.md) |  | 

### Return type

[**PropertyGroup**](PropertyGroup.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3PropertiesObjectTypeGroupsCreate

> PropertyGroup PostCrmV3PropertiesObjectTypeGroupsCreate(ctx, objectType).PropertyGroupCreate(propertyGroupCreate).Execute()

Create a property group



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
    objectType := "objectType_example" // string | 
    propertyGroupCreate := *openapiclient.NewPropertyGroupCreate("Name_example", "Label_example") // PropertyGroupCreate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.PostCrmV3PropertiesObjectTypeGroupsCreate(context.Background(), objectType).PropertyGroupCreate(propertyGroupCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.PostCrmV3PropertiesObjectTypeGroupsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3PropertiesObjectTypeGroupsCreate`: PropertyGroup
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.PostCrmV3PropertiesObjectTypeGroupsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3PropertiesObjectTypeGroupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **propertyGroupCreate** | [**PropertyGroupCreate**](PropertyGroupCreate.md) |  | 

### Return type

[**PropertyGroup**](PropertyGroup.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

