# PublicPerformanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PerformanceView**](PerformanceView.md) |  | 
**Domain** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**StartInterval** | **int64** |  | 
**EndInterval** | **int64** |  | 
**Interval** | **string** |  | 
**Period** | Pointer to **string** |  | [optional] 

## Methods

### NewPublicPerformanceResponse

`func NewPublicPerformanceResponse(data []PerformanceView, startInterval int64, endInterval int64, interval string, ) *PublicPerformanceResponse`

NewPublicPerformanceResponse instantiates a new PublicPerformanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicPerformanceResponseWithDefaults

`func NewPublicPerformanceResponseWithDefaults() *PublicPerformanceResponse`

NewPublicPerformanceResponseWithDefaults instantiates a new PublicPerformanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PublicPerformanceResponse) GetData() []PerformanceView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PublicPerformanceResponse) GetDataOk() (*[]PerformanceView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PublicPerformanceResponse) SetData(v []PerformanceView)`

SetData sets Data field to given value.


### GetDomain

`func (o *PublicPerformanceResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PublicPerformanceResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PublicPerformanceResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PublicPerformanceResponse) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetPath

`func (o *PublicPerformanceResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PublicPerformanceResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PublicPerformanceResponse) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PublicPerformanceResponse) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetStartInterval

`func (o *PublicPerformanceResponse) GetStartInterval() int64`

GetStartInterval returns the StartInterval field if non-nil, zero value otherwise.

### GetStartIntervalOk

`func (o *PublicPerformanceResponse) GetStartIntervalOk() (*int64, bool)`

GetStartIntervalOk returns a tuple with the StartInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartInterval

`func (o *PublicPerformanceResponse) SetStartInterval(v int64)`

SetStartInterval sets StartInterval field to given value.


### GetEndInterval

`func (o *PublicPerformanceResponse) GetEndInterval() int64`

GetEndInterval returns the EndInterval field if non-nil, zero value otherwise.

### GetEndIntervalOk

`func (o *PublicPerformanceResponse) GetEndIntervalOk() (*int64, bool)`

GetEndIntervalOk returns a tuple with the EndInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndInterval

`func (o *PublicPerformanceResponse) SetEndInterval(v int64)`

SetEndInterval sets EndInterval field to given value.


### GetInterval

`func (o *PublicPerformanceResponse) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PublicPerformanceResponse) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PublicPerformanceResponse) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetPeriod

`func (o *PublicPerformanceResponse) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *PublicPerformanceResponse) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *PublicPerformanceResponse) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *PublicPerformanceResponse) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


