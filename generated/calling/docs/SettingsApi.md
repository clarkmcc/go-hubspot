# \SettingsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3ExtensionsCallingAppIdSettingsArchive**](SettingsApi.md#DeleteCrmV3ExtensionsCallingAppIdSettingsArchive) | **Delete** /crm/v3/extensions/calling/{appId}/settings | Delete calling settings
[**GetCrmV3ExtensionsCallingAppIdSettingsGetById**](SettingsApi.md#GetCrmV3ExtensionsCallingAppIdSettingsGetById) | **Get** /crm/v3/extensions/calling/{appId}/settings | Get calling settings
[**PatchCrmV3ExtensionsCallingAppIdSettingsUpdate**](SettingsApi.md#PatchCrmV3ExtensionsCallingAppIdSettingsUpdate) | **Patch** /crm/v3/extensions/calling/{appId}/settings | Update settings
[**PostCrmV3ExtensionsCallingAppIdSettingsCreate**](SettingsApi.md#PostCrmV3ExtensionsCallingAppIdSettingsCreate) | **Post** /crm/v3/extensions/calling/{appId}/settings | Configure a calling extension



## DeleteCrmV3ExtensionsCallingAppIdSettingsArchive

> DeleteCrmV3ExtensionsCallingAppIdSettingsArchive(ctx, appId).Execute()

Delete calling settings



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
    appId := int32(56) // int32 | The ID of the target app.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.DeleteCrmV3ExtensionsCallingAppIdSettingsArchive(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.DeleteCrmV3ExtensionsCallingAppIdSettingsArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3ExtensionsCallingAppIdSettingsArchiveRequest struct via the builder pattern


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


## GetCrmV3ExtensionsCallingAppIdSettingsGetById

> SettingsResponse GetCrmV3ExtensionsCallingAppIdSettingsGetById(ctx, appId).Execute()

Get calling settings



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
    appId := int32(56) // int32 | The ID of the target app.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.GetCrmV3ExtensionsCallingAppIdSettingsGetById(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetCrmV3ExtensionsCallingAppIdSettingsGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3ExtensionsCallingAppIdSettingsGetById`: SettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetCrmV3ExtensionsCallingAppIdSettingsGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ExtensionsCallingAppIdSettingsGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SettingsResponse**](SettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCrmV3ExtensionsCallingAppIdSettingsUpdate

> SettingsResponse PatchCrmV3ExtensionsCallingAppIdSettingsUpdate(ctx, appId).SettingsPatchRequest(settingsPatchRequest).Execute()

Update settings



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
    appId := int32(56) // int32 | The ID of the target app.
    settingsPatchRequest := *openapiclient.NewSettingsPatchRequest() // SettingsPatchRequest | Updated details for the settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.PatchCrmV3ExtensionsCallingAppIdSettingsUpdate(context.Background(), appId).SettingsPatchRequest(settingsPatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.PatchCrmV3ExtensionsCallingAppIdSettingsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCrmV3ExtensionsCallingAppIdSettingsUpdate`: SettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.PatchCrmV3ExtensionsCallingAppIdSettingsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCrmV3ExtensionsCallingAppIdSettingsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **settingsPatchRequest** | [**SettingsPatchRequest**](SettingsPatchRequest.md) | Updated details for the settings. | 

### Return type

[**SettingsResponse**](SettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ExtensionsCallingAppIdSettingsCreate

> SettingsResponse PostCrmV3ExtensionsCallingAppIdSettingsCreate(ctx, appId).SettingsRequest(settingsRequest).Execute()

Configure a calling extension



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
    appId := int32(56) // int32 | The ID of the target app.
    settingsRequest := *openapiclient.NewSettingsRequest("Name_example", "Url_example") // SettingsRequest | Settings state to create with.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.PostCrmV3ExtensionsCallingAppIdSettingsCreate(context.Background(), appId).SettingsRequest(settingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.PostCrmV3ExtensionsCallingAppIdSettingsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ExtensionsCallingAppIdSettingsCreate`: SettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.PostCrmV3ExtensionsCallingAppIdSettingsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExtensionsCallingAppIdSettingsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **settingsRequest** | [**SettingsRequest**](SettingsRequest.md) | Settings state to create with. | 

### Return type

[**SettingsResponse**](SettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

