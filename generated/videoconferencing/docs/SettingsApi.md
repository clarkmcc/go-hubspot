# \SettingsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchive**](SettingsApi.md#DeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchive) | **Delete** /crm/v3/extensions/videoconferencing/settings/{appId} | Delete settings
[**GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetById**](SettingsApi.md#GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetById) | **Get** /crm/v3/extensions/videoconferencing/settings/{appId} | Get settings
[**PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplace**](SettingsApi.md#PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplace) | **Put** /crm/v3/extensions/videoconferencing/settings/{appId} | Update settings



## DeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchive

> DeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchive(ctx, appId).Execute()

Delete settings



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
    appId := int32(56) // int32 | The ID of the video conference application. This is the identifier of the application created in your HubSpot developer portal.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SettingsApi.DeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchive(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.DeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the video conference application. This is the identifier of the application created in your HubSpot developer portal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetById

> ExternalSettings GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetById(ctx, appId).Execute()

Get settings



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
    appId := int32(56) // int32 | The ID of the video conference application. This is the identifier of the application created in your HubSpot developer portal.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SettingsApi.GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetById(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetById`: ExternalSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the video conference application. This is the identifier of the application created in your HubSpot developer portal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ExtensionsVideoconferencingSettingsAppIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalSettings**](ExternalSettings.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplace

> ExternalSettings PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplace(ctx, appId).ExternalSettings(externalSettings).Execute()

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
    appId := int32(56) // int32 | The ID of the video conference application. This is the identifier of the application created in your HubSpot developer portal.
    externalSettings := *openapiclient.NewExternalSettings("CreateMeetingUrl_example") // ExternalSettings | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SettingsApi.PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplace(context.Background(), appId).ExternalSettings(externalSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplace`: ExternalSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the video conference application. This is the identifier of the application created in your HubSpot developer portal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV3ExtensionsVideoconferencingSettingsAppIdReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalSettings** | [**ExternalSettings**](ExternalSettings.md) |  | 

### Return type

[**ExternalSettings**](ExternalSettings.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

