# \RecordingSettingsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormat**](RecordingSettingsApi.md#GetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormat) | **Get** /crm/v3/extensions/calling/{appId}/settings/recording | 
[**PatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormat**](RecordingSettingsApi.md#PatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormat) | **Patch** /crm/v3/extensions/calling/{appId}/settings/recording | 
[**PostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormat**](RecordingSettingsApi.md#PostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormat) | **Post** /crm/v3/extensions/calling/{appId}/settings/recording | 



## GetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormat

> RecordingSettingsResponse GetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormat(ctx, appId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordingSettingsApi.GetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormat(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordingSettingsApi.GetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormat`: RecordingSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `RecordingSettingsApi.GetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RecordingSettingsResponse**](RecordingSettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormat

> RecordingSettingsResponse PatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormat(ctx, appId).RecordingSettingsPatchRequest(recordingSettingsPatchRequest).Execute()



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
    recordingSettingsPatchRequest := *openapiclient.NewRecordingSettingsPatchRequest() // RecordingSettingsPatchRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordingSettingsApi.PatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormat(context.Background(), appId).RecordingSettingsPatchRequest(recordingSettingsPatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordingSettingsApi.PatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormat`: RecordingSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `RecordingSettingsApi.PatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recordingSettingsPatchRequest** | [**RecordingSettingsPatchRequest**](RecordingSettingsPatchRequest.md) |  | 

### Return type

[**RecordingSettingsResponse**](RecordingSettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormat

> RecordingSettingsResponse PostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormat(ctx, appId).RecordingSettingsRequest(recordingSettingsRequest).Execute()



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
    recordingSettingsRequest := *openapiclient.NewRecordingSettingsRequest("UrlToRetrieveAuthedRecording_example") // RecordingSettingsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordingSettingsApi.PostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormat(context.Background(), appId).RecordingSettingsRequest(recordingSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordingSettingsApi.PostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormat`: RecordingSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `RecordingSettingsApi.PostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recordingSettingsRequest** | [**RecordingSettingsRequest**](RecordingSettingsRequest.md) |  | 

### Return type

[**RecordingSettingsResponse**](RecordingSettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

