# ReportCreationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnqueueTime** | [**DateTime**](DateTime.md) |  | 
**UserEmail** | **string** | Email of the user | 
**UserId** | **int32** | ID of the user | 

## Methods

### NewReportCreationResponse

`func NewReportCreationResponse(enqueueTime DateTime, userEmail string, userId int32, ) *ReportCreationResponse`

NewReportCreationResponse instantiates a new ReportCreationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportCreationResponseWithDefaults

`func NewReportCreationResponseWithDefaults() *ReportCreationResponse`

NewReportCreationResponseWithDefaults instantiates a new ReportCreationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnqueueTime

`func (o *ReportCreationResponse) GetEnqueueTime() DateTime`

GetEnqueueTime returns the EnqueueTime field if non-nil, zero value otherwise.

### GetEnqueueTimeOk

`func (o *ReportCreationResponse) GetEnqueueTimeOk() (*DateTime, bool)`

GetEnqueueTimeOk returns a tuple with the EnqueueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnqueueTime

`func (o *ReportCreationResponse) SetEnqueueTime(v DateTime)`

SetEnqueueTime sets EnqueueTime field to given value.


### GetUserEmail

`func (o *ReportCreationResponse) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *ReportCreationResponse) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *ReportCreationResponse) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.


### GetUserId

`func (o *ReportCreationResponse) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ReportCreationResponse) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ReportCreationResponse) SetUserId(v int32)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


