# DateTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateOnly** | **bool** | Indicates whether the DateTime value represents only a date without a time component. | 
**TimeZoneShift** | **int32** | The integer value representing the shift in minutes from UTC for the DateTime value. | 
**Value** | **int64** | The integer value representing a specific point in time. | 

## Methods

### NewDateTime

`func NewDateTime(dateOnly bool, timeZoneShift int32, value int64, ) *DateTime`

NewDateTime instantiates a new DateTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateTimeWithDefaults

`func NewDateTimeWithDefaults() *DateTime`

NewDateTimeWithDefaults instantiates a new DateTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateOnly

`func (o *DateTime) GetDateOnly() bool`

GetDateOnly returns the DateOnly field if non-nil, zero value otherwise.

### GetDateOnlyOk

`func (o *DateTime) GetDateOnlyOk() (*bool, bool)`

GetDateOnlyOk returns a tuple with the DateOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOnly

`func (o *DateTime) SetDateOnly(v bool)`

SetDateOnly sets DateOnly field to given value.


### GetTimeZoneShift

`func (o *DateTime) GetTimeZoneShift() int32`

GetTimeZoneShift returns the TimeZoneShift field if non-nil, zero value otherwise.

### GetTimeZoneShiftOk

`func (o *DateTime) GetTimeZoneShiftOk() (*int32, bool)`

GetTimeZoneShiftOk returns a tuple with the TimeZoneShift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneShift

`func (o *DateTime) SetTimeZoneShift(v int32)`

SetTimeZoneShift sets TimeZoneShift field to given value.


### GetValue

`func (o *DateTime) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DateTime) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DateTime) SetValue(v int64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


