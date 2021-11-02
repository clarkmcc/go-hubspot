# PerformanceView

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTimestamp** | **int64** | The timestamp in milliseconds of the start of this interval. | [default to null]
**EndTimestamp** | **int64** | The timestamp in milliseconds of the end of this interval. | [default to null]
**StartDatetime** | **string** |  | [default to null]
**EndDatetime** | **string** |  | [default to null]
**TotalRequests** | **int32** | The total number of requests received in this period. | [default to null]
**CacheHits** | **int32** | The total number of requests that were served cached responses. | [default to null]
**CacheHitRate** | **float64** | The percentage of requests that were served cached responses. | [default to null]
**TotalRequestTime** | **int32** |  | [default to null]
**AvgOriginResponseTime** | **int32** | The average response time in milliseconds from the origin to the edge. | [default to null]
**ResponseTimeMs** | **int32** | The average response time in milliseconds. | [default to null]
**Var100X** | **int32** | The number of responses that had an http status code between 1000-1999. | [default to null]
**Var20X** | **int32** | The number of responses that had an http status code between 200-299. | [default to null]
**Var30X** | **int32** | The number of responses that had an http status code between 300-399. | [default to null]
**Var40X** | **int32** | The number of responses that had an http status code between 400-499. | [default to null]
**Var50X** | **int32** | The number of responses that had an http status code between 500-599. | [default to null]
**Var403** | **int32** | The number of responses that had an http status code of 403. | [default to null]
**Var404** | **int32** | The number of responses that had an http status code of 404. | [default to null]
**Var500** | **int32** | The number of responses that had an http status code of 500. | [default to null]
**Var504** | **int32** | The number of responses that had an http status code of 504. | [default to null]
**Var50th** | **int32** | The 50th percentile response time. | [default to null]
**Var95th** | **int32** | The 95th percentile response time. | [default to null]
**Var99th** | **int32** | The 99th percentile response time. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

