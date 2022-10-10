# ExternalSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateMeetingUrl** | **string** | The URL that HubSpot will send requests to create a new video conference. | 
**UpdateMeetingUrl** | Pointer to **string** | The URL that HubSpot will send updates to existing meetings. Typically called when the user changes the topic or times of a meeting. | [optional] 
**DeleteMeetingUrl** | Pointer to **string** | The URL that HubSpot will send notifications of meetings that have been deleted in HubSpot. | [optional] 
**UserVerifyUrl** | Pointer to **string** | The URL that HubSpot will use to verify that a user exists in the video conference application. | [optional] 

## Methods

### NewExternalSettings

`func NewExternalSettings(createMeetingUrl string, ) *ExternalSettings`

NewExternalSettings instantiates a new ExternalSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalSettingsWithDefaults

`func NewExternalSettingsWithDefaults() *ExternalSettings`

NewExternalSettingsWithDefaults instantiates a new ExternalSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateMeetingUrl

`func (o *ExternalSettings) GetCreateMeetingUrl() string`

GetCreateMeetingUrl returns the CreateMeetingUrl field if non-nil, zero value otherwise.

### GetCreateMeetingUrlOk

`func (o *ExternalSettings) GetCreateMeetingUrlOk() (*string, bool)`

GetCreateMeetingUrlOk returns a tuple with the CreateMeetingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateMeetingUrl

`func (o *ExternalSettings) SetCreateMeetingUrl(v string)`

SetCreateMeetingUrl sets CreateMeetingUrl field to given value.


### GetUpdateMeetingUrl

`func (o *ExternalSettings) GetUpdateMeetingUrl() string`

GetUpdateMeetingUrl returns the UpdateMeetingUrl field if non-nil, zero value otherwise.

### GetUpdateMeetingUrlOk

`func (o *ExternalSettings) GetUpdateMeetingUrlOk() (*string, bool)`

GetUpdateMeetingUrlOk returns a tuple with the UpdateMeetingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateMeetingUrl

`func (o *ExternalSettings) SetUpdateMeetingUrl(v string)`

SetUpdateMeetingUrl sets UpdateMeetingUrl field to given value.

### HasUpdateMeetingUrl

`func (o *ExternalSettings) HasUpdateMeetingUrl() bool`

HasUpdateMeetingUrl returns a boolean if a field has been set.

### GetDeleteMeetingUrl

`func (o *ExternalSettings) GetDeleteMeetingUrl() string`

GetDeleteMeetingUrl returns the DeleteMeetingUrl field if non-nil, zero value otherwise.

### GetDeleteMeetingUrlOk

`func (o *ExternalSettings) GetDeleteMeetingUrlOk() (*string, bool)`

GetDeleteMeetingUrlOk returns a tuple with the DeleteMeetingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteMeetingUrl

`func (o *ExternalSettings) SetDeleteMeetingUrl(v string)`

SetDeleteMeetingUrl sets DeleteMeetingUrl field to given value.

### HasDeleteMeetingUrl

`func (o *ExternalSettings) HasDeleteMeetingUrl() bool`

HasDeleteMeetingUrl returns a boolean if a field has been set.

### GetUserVerifyUrl

`func (o *ExternalSettings) GetUserVerifyUrl() string`

GetUserVerifyUrl returns the UserVerifyUrl field if non-nil, zero value otherwise.

### GetUserVerifyUrlOk

`func (o *ExternalSettings) GetUserVerifyUrlOk() (*string, bool)`

GetUserVerifyUrlOk returns a tuple with the UserVerifyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVerifyUrl

`func (o *ExternalSettings) SetUserVerifyUrl(v string)`

SetUserVerifyUrl sets UserVerifyUrl field to given value.

### HasUserVerifyUrl

`func (o *ExternalSettings) HasUserVerifyUrl() bool`

HasUserVerifyUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


