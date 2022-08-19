# \PublicPerformanceApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCmsV3PerformanceGetPage**](PublicPerformanceApi.md#GetCmsV3PerformanceGetPage) | **Get** /cms/v3/performance/ | View your website&#39;s performance.
[**GetCmsV3PerformanceUptimeGetUptime**](PublicPerformanceApi.md#GetCmsV3PerformanceUptimeGetUptime) | **Get** /cms/v3/performance/uptime | View your website&#39;s uptime.



## GetCmsV3PerformanceGetPage

> PublicPerformanceResponse GetCmsV3PerformanceGetPage(ctx).Domain(domain).Path(path).Pad(pad).Sum(sum).Period(period).Interval(interval).Start(start).End(end).Execute()

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
    resp, r, err := apiClient.PublicPerformanceApi.GetCmsV3PerformanceGetPage(context.Background()).Domain(domain).Path(path).Pad(pad).Sum(sum).Period(period).Interval(interval).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicPerformanceApi.GetCmsV3PerformanceGetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3PerformanceGetPage`: PublicPerformanceResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicPerformanceApi.GetCmsV3PerformanceGetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3PerformanceGetPageRequest struct via the builder pattern


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

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3PerformanceUptimeGetUptime

> PublicPerformanceResponse GetCmsV3PerformanceUptimeGetUptime(ctx).Domain(domain).Path(path).Pad(pad).Sum(sum).Period(period).Interval(interval).Start(start).End(end).Execute()

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
    resp, r, err := apiClient.PublicPerformanceApi.GetCmsV3PerformanceUptimeGetUptime(context.Background()).Domain(domain).Path(path).Pad(pad).Sum(sum).Period(period).Interval(interval).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicPerformanceApi.GetCmsV3PerformanceUptimeGetUptime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3PerformanceUptimeGetUptime`: PublicPerformanceResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicPerformanceApi.GetCmsV3PerformanceUptimeGetUptime`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3PerformanceUptimeGetUptimeRequest struct via the builder pattern


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

[hapikey](../README.md#hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

