# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getcmsv3performanceGetPage**](PublicPerformanceApi.md#Getcmsv3performanceGetPage) | **Get** /cms/v3/performance/ | View your website&#x27;s performance.
[**Getcmsv3performanceuptimeGetUptime**](PublicPerformanceApi.md#Getcmsv3performanceuptimeGetUptime) | **Get** /cms/v3/performance/uptime | View your website&#x27;s uptime.

# **Getcmsv3performanceGetPage**
> PublicPerformanceResponse Getcmsv3performanceGetPage(ctx, optional)
View your website's performance.

Returns time series data website performance data for the given domain and/or path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PublicPerformanceApiGetcmsv3performanceGetPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicPerformanceApiGetcmsv3performanceGetPageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **optional.String**| The domain to search return data for. | 
 **path** | **optional.String**| The url path of the domain to return data for. | 
 **pad** | **optional.Bool**| Specifies whether the time series data should have empty intervals if performance data is not present to create a continuous set. | 
 **sum** | **optional.Bool**| Specifies whether the time series data should be summated for the given period. Defaults to false. | 
 **period** | **optional.String**| A relative period to return time series data for. This value is ignored if start and/or end are provided. Valid periods: [15m, 30m, 1h, 4h, 12h, 1d, 1w] | 
 **interval** | **optional.String**| The time series interval to group data by. Valid intervals: [1m, 5m, 15m, 30m, 1h, 4h, 12h, 1d, 1w] | 
 **start** | **optional.Int64**| A timestamp in milliseconds that indicates the start of the time period. | 
 **end** | **optional.Int64**| A timestamp in milliseconds that indicates the end of the time period. | 

### Return type

[**PublicPerformanceResponse**](PublicPerformanceResponse.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcmsv3performanceuptimeGetUptime**
> PublicPerformanceResponse Getcmsv3performanceuptimeGetUptime(ctx, optional)
View your website's uptime.

Returns uptime time series website performance data for the given domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PublicPerformanceApiGetcmsv3performanceuptimeGetUptimeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicPerformanceApiGetcmsv3performanceuptimeGetUptimeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **optional.String**| The domain to search return data for. | 
 **path** | **optional.String**|  | 
 **pad** | **optional.Bool**| Specifies whether the time series data should have empty intervals if performance data is not present to create a continuous set. | 
 **sum** | **optional.Bool**| Specifies whether the time series data should be summated for the given period. Defaults to false. | 
 **period** | **optional.String**| A relative period to return time series data for. This value is ignored if start and/or end are provided. Valid periods: [15m, 30m, 1h, 4h, 12h, 1d, 1w] | 
 **interval** | **optional.String**| The time series interval to group data by. Valid intervals: [1m, 5m, 15m, 30m, 1h, 4h, 12h, 1d, 1w] | 
 **start** | **optional.Int64**| A timestamp in milliseconds that indicates the start of the time period. | 
 **end** | **optional.Int64**| A timestamp in milliseconds that indicates the end of the time period. | 

### Return type

[**PublicPerformanceResponse**](PublicPerformanceResponse.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

