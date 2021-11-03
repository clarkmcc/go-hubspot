# PerformanceView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTimestamp** | **int64** | The timestamp in milliseconds of the start of this interval. | 
**EndTimestamp** | **int64** | The timestamp in milliseconds of the end of this interval. | 
**StartDatetime** | **string** |  | 
**EndDatetime** | **string** |  | 
**TotalRequests** | **int32** | The total number of requests received in this period. | 
**CacheHits** | **int32** | The total number of requests that were served cached responses. | 
**CacheHitRate** | **float32** | The percentage of requests that were served cached responses. | 
**TotalRequestTime** | **int32** |  | 
**AvgOriginResponseTime** | **int32** | The average response time in milliseconds from the origin to the edge. | 
**ResponseTimeMs** | **int32** | The average response time in milliseconds. | 
**Var100X** | **int32** | The number of responses that had an http status code between 1000-1999. | 
**Var20X** | **int32** | The number of responses that had an http status code between 200-299. | 
**Var30X** | **int32** | The number of responses that had an http status code between 300-399. | 
**Var40X** | **int32** | The number of responses that had an http status code between 400-499. | 
**Var50X** | **int32** | The number of responses that had an http status code between 500-599. | 
**Var403** | **int32** | The number of responses that had an http status code of 403. | 
**Var404** | **int32** | The number of responses that had an http status code of 404. | 
**Var500** | **int32** | The number of responses that had an http status code of 500. | 
**Var504** | **int32** | The number of responses that had an http status code of 504. | 
**Var50th** | **int32** | The 50th percentile response time. | 
**Var95th** | **int32** | The 95th percentile response time. | 
**Var99th** | **int32** | The 99th percentile response time. | 

## Methods

### NewPerformanceView

`func NewPerformanceView(startTimestamp int64, endTimestamp int64, startDatetime string, endDatetime string, totalRequests int32, cacheHits int32, cacheHitRate float32, totalRequestTime int32, avgOriginResponseTime int32, responseTimeMs int32, var100X int32, var20X int32, var30X int32, var40X int32, var50X int32, var403 int32, var404 int32, var500 int32, var504 int32, var50th int32, var95th int32, var99th int32, ) *PerformanceView`

NewPerformanceView instantiates a new PerformanceView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformanceViewWithDefaults

`func NewPerformanceViewWithDefaults() *PerformanceView`

NewPerformanceViewWithDefaults instantiates a new PerformanceView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTimestamp

`func (o *PerformanceView) GetStartTimestamp() int64`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *PerformanceView) GetStartTimestampOk() (*int64, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *PerformanceView) SetStartTimestamp(v int64)`

SetStartTimestamp sets StartTimestamp field to given value.


### GetEndTimestamp

`func (o *PerformanceView) GetEndTimestamp() int64`

GetEndTimestamp returns the EndTimestamp field if non-nil, zero value otherwise.

### GetEndTimestampOk

`func (o *PerformanceView) GetEndTimestampOk() (*int64, bool)`

GetEndTimestampOk returns a tuple with the EndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestamp

`func (o *PerformanceView) SetEndTimestamp(v int64)`

SetEndTimestamp sets EndTimestamp field to given value.


### GetStartDatetime

`func (o *PerformanceView) GetStartDatetime() string`

GetStartDatetime returns the StartDatetime field if non-nil, zero value otherwise.

### GetStartDatetimeOk

`func (o *PerformanceView) GetStartDatetimeOk() (*string, bool)`

GetStartDatetimeOk returns a tuple with the StartDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDatetime

`func (o *PerformanceView) SetStartDatetime(v string)`

SetStartDatetime sets StartDatetime field to given value.


### GetEndDatetime

`func (o *PerformanceView) GetEndDatetime() string`

GetEndDatetime returns the EndDatetime field if non-nil, zero value otherwise.

### GetEndDatetimeOk

`func (o *PerformanceView) GetEndDatetimeOk() (*string, bool)`

GetEndDatetimeOk returns a tuple with the EndDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDatetime

`func (o *PerformanceView) SetEndDatetime(v string)`

SetEndDatetime sets EndDatetime field to given value.


### GetTotalRequests

`func (o *PerformanceView) GetTotalRequests() int32`

GetTotalRequests returns the TotalRequests field if non-nil, zero value otherwise.

### GetTotalRequestsOk

`func (o *PerformanceView) GetTotalRequestsOk() (*int32, bool)`

GetTotalRequestsOk returns a tuple with the TotalRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRequests

`func (o *PerformanceView) SetTotalRequests(v int32)`

SetTotalRequests sets TotalRequests field to given value.


### GetCacheHits

`func (o *PerformanceView) GetCacheHits() int32`

GetCacheHits returns the CacheHits field if non-nil, zero value otherwise.

### GetCacheHitsOk

`func (o *PerformanceView) GetCacheHitsOk() (*int32, bool)`

GetCacheHitsOk returns a tuple with the CacheHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheHits

`func (o *PerformanceView) SetCacheHits(v int32)`

SetCacheHits sets CacheHits field to given value.


### GetCacheHitRate

`func (o *PerformanceView) GetCacheHitRate() float32`

GetCacheHitRate returns the CacheHitRate field if non-nil, zero value otherwise.

### GetCacheHitRateOk

`func (o *PerformanceView) GetCacheHitRateOk() (*float32, bool)`

GetCacheHitRateOk returns a tuple with the CacheHitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheHitRate

`func (o *PerformanceView) SetCacheHitRate(v float32)`

SetCacheHitRate sets CacheHitRate field to given value.


### GetTotalRequestTime

`func (o *PerformanceView) GetTotalRequestTime() int32`

GetTotalRequestTime returns the TotalRequestTime field if non-nil, zero value otherwise.

### GetTotalRequestTimeOk

`func (o *PerformanceView) GetTotalRequestTimeOk() (*int32, bool)`

GetTotalRequestTimeOk returns a tuple with the TotalRequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRequestTime

`func (o *PerformanceView) SetTotalRequestTime(v int32)`

SetTotalRequestTime sets TotalRequestTime field to given value.


### GetAvgOriginResponseTime

`func (o *PerformanceView) GetAvgOriginResponseTime() int32`

GetAvgOriginResponseTime returns the AvgOriginResponseTime field if non-nil, zero value otherwise.

### GetAvgOriginResponseTimeOk

`func (o *PerformanceView) GetAvgOriginResponseTimeOk() (*int32, bool)`

GetAvgOriginResponseTimeOk returns a tuple with the AvgOriginResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgOriginResponseTime

`func (o *PerformanceView) SetAvgOriginResponseTime(v int32)`

SetAvgOriginResponseTime sets AvgOriginResponseTime field to given value.


### GetResponseTimeMs

`func (o *PerformanceView) GetResponseTimeMs() int32`

GetResponseTimeMs returns the ResponseTimeMs field if non-nil, zero value otherwise.

### GetResponseTimeMsOk

`func (o *PerformanceView) GetResponseTimeMsOk() (*int32, bool)`

GetResponseTimeMsOk returns a tuple with the ResponseTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeMs

`func (o *PerformanceView) SetResponseTimeMs(v int32)`

SetResponseTimeMs sets ResponseTimeMs field to given value.


### GetVar100X

`func (o *PerformanceView) GetVar100X() int32`

GetVar100X returns the Var100X field if non-nil, zero value otherwise.

### GetVar100XOk

`func (o *PerformanceView) GetVar100XOk() (*int32, bool)`

GetVar100XOk returns a tuple with the Var100X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar100X

`func (o *PerformanceView) SetVar100X(v int32)`

SetVar100X sets Var100X field to given value.


### GetVar20X

`func (o *PerformanceView) GetVar20X() int32`

GetVar20X returns the Var20X field if non-nil, zero value otherwise.

### GetVar20XOk

`func (o *PerformanceView) GetVar20XOk() (*int32, bool)`

GetVar20XOk returns a tuple with the Var20X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar20X

`func (o *PerformanceView) SetVar20X(v int32)`

SetVar20X sets Var20X field to given value.


### GetVar30X

`func (o *PerformanceView) GetVar30X() int32`

GetVar30X returns the Var30X field if non-nil, zero value otherwise.

### GetVar30XOk

`func (o *PerformanceView) GetVar30XOk() (*int32, bool)`

GetVar30XOk returns a tuple with the Var30X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar30X

`func (o *PerformanceView) SetVar30X(v int32)`

SetVar30X sets Var30X field to given value.


### GetVar40X

`func (o *PerformanceView) GetVar40X() int32`

GetVar40X returns the Var40X field if non-nil, zero value otherwise.

### GetVar40XOk

`func (o *PerformanceView) GetVar40XOk() (*int32, bool)`

GetVar40XOk returns a tuple with the Var40X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar40X

`func (o *PerformanceView) SetVar40X(v int32)`

SetVar40X sets Var40X field to given value.


### GetVar50X

`func (o *PerformanceView) GetVar50X() int32`

GetVar50X returns the Var50X field if non-nil, zero value otherwise.

### GetVar50XOk

`func (o *PerformanceView) GetVar50XOk() (*int32, bool)`

GetVar50XOk returns a tuple with the Var50X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar50X

`func (o *PerformanceView) SetVar50X(v int32)`

SetVar50X sets Var50X field to given value.


### GetVar403

`func (o *PerformanceView) GetVar403() int32`

GetVar403 returns the Var403 field if non-nil, zero value otherwise.

### GetVar403Ok

`func (o *PerformanceView) GetVar403Ok() (*int32, bool)`

GetVar403Ok returns a tuple with the Var403 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar403

`func (o *PerformanceView) SetVar403(v int32)`

SetVar403 sets Var403 field to given value.


### GetVar404

`func (o *PerformanceView) GetVar404() int32`

GetVar404 returns the Var404 field if non-nil, zero value otherwise.

### GetVar404Ok

`func (o *PerformanceView) GetVar404Ok() (*int32, bool)`

GetVar404Ok returns a tuple with the Var404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar404

`func (o *PerformanceView) SetVar404(v int32)`

SetVar404 sets Var404 field to given value.


### GetVar500

`func (o *PerformanceView) GetVar500() int32`

GetVar500 returns the Var500 field if non-nil, zero value otherwise.

### GetVar500Ok

`func (o *PerformanceView) GetVar500Ok() (*int32, bool)`

GetVar500Ok returns a tuple with the Var500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar500

`func (o *PerformanceView) SetVar500(v int32)`

SetVar500 sets Var500 field to given value.


### GetVar504

`func (o *PerformanceView) GetVar504() int32`

GetVar504 returns the Var504 field if non-nil, zero value otherwise.

### GetVar504Ok

`func (o *PerformanceView) GetVar504Ok() (*int32, bool)`

GetVar504Ok returns a tuple with the Var504 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar504

`func (o *PerformanceView) SetVar504(v int32)`

SetVar504 sets Var504 field to given value.


### GetVar50th

`func (o *PerformanceView) GetVar50th() int32`

GetVar50th returns the Var50th field if non-nil, zero value otherwise.

### GetVar50thOk

`func (o *PerformanceView) GetVar50thOk() (*int32, bool)`

GetVar50thOk returns a tuple with the Var50th field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar50th

`func (o *PerformanceView) SetVar50th(v int32)`

SetVar50th sets Var50th field to given value.


### GetVar95th

`func (o *PerformanceView) GetVar95th() int32`

GetVar95th returns the Var95th field if non-nil, zero value otherwise.

### GetVar95thOk

`func (o *PerformanceView) GetVar95thOk() (*int32, bool)`

GetVar95thOk returns a tuple with the Var95th field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar95th

`func (o *PerformanceView) SetVar95th(v int32)`

SetVar95th sets Var95th field to given value.


### GetVar99th

`func (o *PerformanceView) GetVar99th() int32`

GetVar99th returns the Var99th field if non-nil, zero value otherwise.

### GetVar99thOk

`func (o *PerformanceView) GetVar99thOk() (*int32, bool)`

GetVar99thOk returns a tuple with the Var99th field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar99th

`func (o *PerformanceView) SetVar99th(v int32)`

SetVar99th sets Var99th field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


