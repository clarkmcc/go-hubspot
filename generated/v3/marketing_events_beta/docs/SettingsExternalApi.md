# \SettingsExternalApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SettingsCreate**](SettingsExternalApi.md#SettingsCreate) | **Post** /marketing/v3/marketing-events/{appId}/settings | 
[**SettingsGetAll**](SettingsExternalApi.md#SettingsGetAll) | **Get** /marketing/v3/marketing-events/{appId}/settings | 



## SettingsCreate

> EventDetailSettings SettingsCreate(ctx, appId).EventDetailSettingsUrl(eventDetailSettingsUrl).Execute()



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
    eventDetailSettingsUrl := *openapiclient.NewEventDetailSettingsUrl("EventDetailsUrl_example") // EventDetailSettingsUrl | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsExternalApi.SettingsCreate(context.Background(), appId).EventDetailSettingsUrl(eventDetailSettingsUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsExternalApi.SettingsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SettingsCreate`: EventDetailSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsExternalApi.SettingsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventDetailSettingsUrl** | [**EventDetailSettingsUrl**](EventDetailSettingsUrl.md) |  | 

### Return type

[**EventDetailSettings**](EventDetailSettings.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey), [hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SettingsGetAll

> EventDetailSettings SettingsGetAll(ctx, appId).Execute()



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
    resp, r, err := apiClient.SettingsExternalApi.SettingsGetAll(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsExternalApi.SettingsGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SettingsGetAll`: EventDetailSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsExternalApi.SettingsGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventDetailSettings**](EventDetailSettings.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey), [hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

