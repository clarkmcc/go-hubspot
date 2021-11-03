# SettingsChangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetUrl** | **string** | A publicly available URL for Hubspot to call where event payloads will be delivered. See [link-so-some-doc](#) for details about the format of these event payloads. | 
**Throttling** | [**ThrottlingSettings**](ThrottlingSettings.md) |  | 

## Methods

### NewSettingsChangeRequest

`func NewSettingsChangeRequest(targetUrl string, throttling ThrottlingSettings, ) *SettingsChangeRequest`

NewSettingsChangeRequest instantiates a new SettingsChangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsChangeRequestWithDefaults

`func NewSettingsChangeRequestWithDefaults() *SettingsChangeRequest`

NewSettingsChangeRequestWithDefaults instantiates a new SettingsChangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetUrl

`func (o *SettingsChangeRequest) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *SettingsChangeRequest) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *SettingsChangeRequest) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.


### GetThrottling

`func (o *SettingsChangeRequest) GetThrottling() ThrottlingSettings`

GetThrottling returns the Throttling field if non-nil, zero value otherwise.

### GetThrottlingOk

`func (o *SettingsChangeRequest) GetThrottlingOk() (*ThrottlingSettings, bool)`

GetThrottlingOk returns a tuple with the Throttling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottling

`func (o *SettingsChangeRequest) SetThrottling(v ThrottlingSettings)`

SetThrottling sets Throttling field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


