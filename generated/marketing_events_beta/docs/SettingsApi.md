# \SettingsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMarketingV3MarketingEventsAppIdSettingsGetAll**](SettingsApi.md#GetMarketingV3MarketingEventsAppIdSettingsGetAll) | **Get** /marketing/v3/marketing-events/{appId}/settings | Retrieve the application settings
[**PostMarketingV3MarketingEventsAppIdSettingsCreate**](SettingsApi.md#PostMarketingV3MarketingEventsAppIdSettingsCreate) | **Post** /marketing/v3/marketing-events/{appId}/settings | Update the application settings



## GetMarketingV3MarketingEventsAppIdSettingsGetAll

> EventDetailSettings GetMarketingV3MarketingEventsAppIdSettingsGetAll(ctx, appId).Execute()

Retrieve the application settings



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
    appId := int32(56) // int32 | The id of the application to retrieve the settings for.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.GetMarketingV3MarketingEventsAppIdSettingsGetAll(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetMarketingV3MarketingEventsAppIdSettingsGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketingV3MarketingEventsAppIdSettingsGetAll`: EventDetailSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetMarketingV3MarketingEventsAppIdSettingsGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The id of the application to retrieve the settings for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3MarketingEventsAppIdSettingsGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventDetailSettings**](EventDetailSettings.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsAppIdSettingsCreate

> EventDetailSettings PostMarketingV3MarketingEventsAppIdSettingsCreate(ctx, appId).EventDetailSettingsUrl(eventDetailSettingsUrl).Execute()

Update the application settings



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
    appId := int32(56) // int32 | The id of the application to update the settings for.
    eventDetailSettingsUrl := *openapiclient.NewEventDetailSettingsUrl("EventDetailsUrl_example") // EventDetailSettingsUrl | The new application settings

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.PostMarketingV3MarketingEventsAppIdSettingsCreate(context.Background(), appId).EventDetailSettingsUrl(eventDetailSettingsUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.PostMarketingV3MarketingEventsAppIdSettingsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMarketingV3MarketingEventsAppIdSettingsCreate`: EventDetailSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.PostMarketingV3MarketingEventsAppIdSettingsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The id of the application to update the settings for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsAppIdSettingsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventDetailSettingsUrl** | [**EventDetailSettingsUrl**](EventDetailSettingsUrl.md) | The new application settings | 

### Return type

[**EventDetailSettings**](EventDetailSettings.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

