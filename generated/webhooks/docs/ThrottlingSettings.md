# ThrottlingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxConcurrentRequests** | **int32** | The maximum number of HTTP requests HubSpot will attempt to make to your app in a given time frame determined by &#x60;period&#x60;. | 
**Period** | **string** | Time scale for this setting. Can be either &#x60;SECONDLY&#x60; (per second) or &#x60;ROLLING_MINUTE&#x60; (per minute). | 

## Methods

### NewThrottlingSettings

`func NewThrottlingSettings(maxConcurrentRequests int32, period string, ) *ThrottlingSettings`

NewThrottlingSettings instantiates a new ThrottlingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThrottlingSettingsWithDefaults

`func NewThrottlingSettingsWithDefaults() *ThrottlingSettings`

NewThrottlingSettingsWithDefaults instantiates a new ThrottlingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxConcurrentRequests

`func (o *ThrottlingSettings) GetMaxConcurrentRequests() int32`

GetMaxConcurrentRequests returns the MaxConcurrentRequests field if non-nil, zero value otherwise.

### GetMaxConcurrentRequestsOk

`func (o *ThrottlingSettings) GetMaxConcurrentRequestsOk() (*int32, bool)`

GetMaxConcurrentRequestsOk returns a tuple with the MaxConcurrentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentRequests

`func (o *ThrottlingSettings) SetMaxConcurrentRequests(v int32)`

SetMaxConcurrentRequests sets MaxConcurrentRequests field to given value.


### GetPeriod

`func (o *ThrottlingSettings) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *ThrottlingSettings) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *ThrottlingSettings) SetPeriod(v string)`

SetPeriod sets Period field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


