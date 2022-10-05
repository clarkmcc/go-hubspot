# \SettingsExternalApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMarketingV3MarketingEventsAppIdSettingsGetAll**](SettingsExternalApi.md#GetMarketingV3MarketingEventsAppIdSettingsGetAll) | **Get** /marketing/v3/marketing-events/{appId}/settings | 
[**PostMarketingV3MarketingEventsAppIdSettingsCreate**](SettingsExternalApi.md#PostMarketingV3MarketingEventsAppIdSettingsCreate) | **Post** /marketing/v3/marketing-events/{appId}/settings | 



## GetMarketingV3MarketingEventsAppIdSettingsGetAll

> EventDetailSettings GetMarketingV3MarketingEventsAppIdSettingsGetAll(ctx, appId).Execute()



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
    resp, r, err := apiClient.SettingsExternalApi.GetMarketingV3MarketingEventsAppIdSettingsGetAll(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsExternalApi.GetMarketingV3MarketingEventsAppIdSettingsGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketingV3MarketingEventsAppIdSettingsGetAll`: EventDetailSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsExternalApi.GetMarketingV3MarketingEventsAppIdSettingsGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3MarketingEventsAppIdSettingsGetAllRequest struct via the builder pattern


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


## PostMarketingV3MarketingEventsAppIdSettingsCreate

> EventDetailSettings PostMarketingV3MarketingEventsAppIdSettingsCreate(ctx, appId).EventDetailSettingsUrl(eventDetailSettingsUrl).Execute()



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
    resp, r, err := apiClient.SettingsExternalApi.PostMarketingV3MarketingEventsAppIdSettingsCreate(context.Background(), appId).EventDetailSettingsUrl(eventDetailSettingsUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsExternalApi.PostMarketingV3MarketingEventsAppIdSettingsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMarketingV3MarketingEventsAppIdSettingsCreate`: EventDetailSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsExternalApi.PostMarketingV3MarketingEventsAppIdSettingsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsAppIdSettingsCreateRequest struct via the builder pattern


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

