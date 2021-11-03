# \SettingsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWebhooksV3AppIdSettingsClear**](SettingsApi.md#DeleteWebhooksV3AppIdSettingsClear) | **Delete** /webhooks/v3/{appId}/settings | Clear webhook settings
[**GetWebhooksV3AppIdSettingsGetAll**](SettingsApi.md#GetWebhooksV3AppIdSettingsGetAll) | **Get** /webhooks/v3/{appId}/settings | Get webhook settings
[**PutWebhooksV3AppIdSettingsConfigure**](SettingsApi.md#PutWebhooksV3AppIdSettingsConfigure) | **Put** /webhooks/v3/{appId}/settings | Configure webhook settings



## DeleteWebhooksV3AppIdSettingsClear

> DeleteWebhooksV3AppIdSettingsClear(ctx, appId).Execute()

Clear webhook settings



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SettingsApi.DeleteWebhooksV3AppIdSettingsClear(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.DeleteWebhooksV3AppIdSettingsClear``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteWebhooksV3AppIdSettingsClearRequest struct via the builder pattern


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


## GetWebhooksV3AppIdSettingsGetAll

> SettingsResponse GetWebhooksV3AppIdSettingsGetAll(ctx, appId).Execute()

Get webhook settings



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SettingsApi.GetWebhooksV3AppIdSettingsGetAll(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetWebhooksV3AppIdSettingsGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhooksV3AppIdSettingsGetAll`: SettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetWebhooksV3AppIdSettingsGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksV3AppIdSettingsGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SettingsResponse**](SettingsResponse.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutWebhooksV3AppIdSettingsConfigure

> SettingsResponse PutWebhooksV3AppIdSettingsConfigure(ctx, appId).SettingsChangeRequest(settingsChangeRequest).Execute()

Configure webhook settings



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
    settingsChangeRequest := *openapiclient.NewSettingsChangeRequest("TargetUrl_example", *openapiclient.NewThrottlingSettings(int32(123), "Period_example")) // SettingsChangeRequest | Settings state to create new with or replace existing settings with.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SettingsApi.PutWebhooksV3AppIdSettingsConfigure(context.Background(), appId).SettingsChangeRequest(settingsChangeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.PutWebhooksV3AppIdSettingsConfigure``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutWebhooksV3AppIdSettingsConfigure`: SettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.PutWebhooksV3AppIdSettingsConfigure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutWebhooksV3AppIdSettingsConfigureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **settingsChangeRequest** | [**SettingsChangeRequest**](SettingsChangeRequest.md) | Settings state to create new with or replace existing settings with. | 

### Return type

[**SettingsResponse**](SettingsResponse.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

