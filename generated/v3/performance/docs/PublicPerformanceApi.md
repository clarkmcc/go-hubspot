# \PublicPerformanceApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPage**](PublicPerformanceApi.md#GetPage) | **Get** /cms/v3/performance/ | View your website&#39;s performance.
[**GetUptime**](PublicPerformanceApi.md#GetUptime) | **Get** /cms/v3/performance/uptime | View your website&#39;s uptime.



## GetPage

> PublicPerformanceResponse GetPage(ctx).Domain(domain).Path(path).Pad(pad).Sum(sum).Period(period).Interval(interval).Start(start).End(end).Execute()

View your website's performance.



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
    domain := "domain_example" // string | The domain to search return data for. (optional)
    path := "path_example" // string | The url path of the domain to return data for. (optional)
    pad := true // bool | Specifies whether the time series data should have empty intervals if performance data is not present to create a continuous set. (optional)
    sum := true // bool | Specifies whether the time series data should be summated for the given period. Defaults to false. (optional)
    period := "period_example" // string | A relative period to return time series data for. This value is ignored if start and/or end are provided. Valid periods: [15m, 30m, 1h, 4h, 12h, 1d, 1w] (optional)
    interval := "interval_example" // string | The time series interval to group data by. Valid intervals: [1m, 5m, 15m, 30m, 1h, 4h, 12h, 1d, 1w] (optional)
    start := int64(789) // int64 | A timestamp in milliseconds that indicates the start of the time period. (optional)
    end := int64(789) // int64 | A timestamp in milliseconds that indicates the end of the time period. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicPerformanceApi.GetPage(context.Background()).Domain(domain).Path(path).Pad(pad).Sum(sum).Period(period).Interval(interval).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicPerformanceApi.GetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPage`: PublicPerformanceResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicPerformanceApi.GetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | The domain to search return data for. | 
 **path** | **string** | The url path of the domain to return data for. | 
 **pad** | **bool** | Specifies whether the time series data should have empty intervals if performance data is not present to create a continuous set. | 
 **sum** | **bool** | Specifies whether the time series data should be summated for the given period. Defaults to false. | 
 **period** | **string** | A relative period to return time series data for. This value is ignored if start and/or end are provided. Valid periods: [15m, 30m, 1h, 4h, 12h, 1d, 1w] | 
 **interval** | **string** | The time series interval to group data by. Valid intervals: [1m, 5m, 15m, 30m, 1h, 4h, 12h, 1d, 1w] | 
 **start** | **int64** | A timestamp in milliseconds that indicates the start of the time period. | 
 **end** | **int64** | A timestamp in milliseconds that indicates the end of the time period. | 

### Return type

[**PublicPerformanceResponse**](PublicPerformanceResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUptime

> PublicPerformanceResponse GetUptime(ctx).Domain(domain).Path(path).Pad(pad).Sum(sum).Period(period).Interval(interval).Start(start).End(end).Execute()

View your website's uptime.



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
    domain := "domain_example" // string | The domain to search return data for. (optional)
    path := "path_example" // string |  (optional)
    pad := true // bool | Specifies whether the time series data should have empty intervals if performance data is not present to create a continuous set. (optional)
    sum := true // bool | Specifies whether the time series data should be summated for the given period. Defaults to false. (optional)
    period := "period_example" // string | A relative period to return time series data for. This value is ignored if start and/or end are provided. Valid periods: [15m, 30m, 1h, 4h, 12h, 1d, 1w] (optional)
    interval := "interval_example" // string | The time series interval to group data by. Valid intervals: [1m, 5m, 15m, 30m, 1h, 4h, 12h, 1d, 1w] (optional)
    start := int64(789) // int64 | A timestamp in milliseconds that indicates the start of the time period. (optional)
    end := int64(789) // int64 | A timestamp in milliseconds that indicates the end of the time period. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicPerformanceApi.GetUptime(context.Background()).Domain(domain).Path(path).Pad(pad).Sum(sum).Period(period).Interval(interval).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicPerformanceApi.GetUptime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUptime`: PublicPerformanceResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicPerformanceApi.GetUptime`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUptimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | The domain to search return data for. | 
 **path** | **string** |  | 
 **pad** | **bool** | Specifies whether the time series data should have empty intervals if performance data is not present to create a continuous set. | 
 **sum** | **bool** | Specifies whether the time series data should be summated for the given period. Defaults to false. | 
 **period** | **string** | A relative period to return time series data for. This value is ignored if start and/or end are provided. Valid periods: [15m, 30m, 1h, 4h, 12h, 1d, 1w] | 
 **interval** | **string** | The time series interval to group data by. Valid intervals: [1m, 5m, 15m, 30m, 1h, 4h, 12h, 1d, 1w] | 
 **start** | **int64** | A timestamp in milliseconds that indicates the start of the time period. | 
 **end** | **int64** | A timestamp in milliseconds that indicates the end of the time period. | 

### Return type

[**PublicPerformanceResponse**](PublicPerformanceResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

