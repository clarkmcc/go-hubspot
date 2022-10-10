# SettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of your calling service to display to users. | 
**Url** | **string** | The URL to your phone/calling UI, built with the [Calling SDK](#). | 
**Height** | **int32** | The target height of the iframe that will contain your phone/calling UI. | 
**Width** | **int32** | The target width of the iframe that will contain your phone/calling UI. | 
**IsReady** | **bool** | When true, your service will appear as an option under the *Call* action in contact records of connected accounts. | 
**SupportsCustomObjects** | **bool** | When true, you are indicating that your service is compatible with engagement v2 service and can be used with custom objects. | 
**CreatedAt** | **time.Time** | When this calling extension was created. | 
**UpdatedAt** | **time.Time** | The last time the settings for this calling extension were modified. | 

## Methods

### NewSettingsResponse

`func NewSettingsResponse(name string, url string, height int32, width int32, isReady bool, supportsCustomObjects bool, createdAt time.Time, updatedAt time.Time, ) *SettingsResponse`

NewSettingsResponse instantiates a new SettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsResponseWithDefaults

`func NewSettingsResponseWithDefaults() *SettingsResponse`

NewSettingsResponseWithDefaults instantiates a new SettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SettingsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SettingsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SettingsResponse) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *SettingsResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SettingsResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SettingsResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeight

`func (o *SettingsResponse) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *SettingsResponse) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *SettingsResponse) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetWidth

`func (o *SettingsResponse) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *SettingsResponse) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *SettingsResponse) SetWidth(v int32)`

SetWidth sets Width field to given value.


### GetIsReady

`func (o *SettingsResponse) GetIsReady() bool`

GetIsReady returns the IsReady field if non-nil, zero value otherwise.

### GetIsReadyOk

`func (o *SettingsResponse) GetIsReadyOk() (*bool, bool)`

GetIsReadyOk returns a tuple with the IsReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReady

`func (o *SettingsResponse) SetIsReady(v bool)`

SetIsReady sets IsReady field to given value.


### GetSupportsCustomObjects

`func (o *SettingsResponse) GetSupportsCustomObjects() bool`

GetSupportsCustomObjects returns the SupportsCustomObjects field if non-nil, zero value otherwise.

### GetSupportsCustomObjectsOk

`func (o *SettingsResponse) GetSupportsCustomObjectsOk() (*bool, bool)`

GetSupportsCustomObjectsOk returns a tuple with the SupportsCustomObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsCustomObjects

`func (o *SettingsResponse) SetSupportsCustomObjects(v bool)`

SetSupportsCustomObjects sets SupportsCustomObjects field to given value.


### GetCreatedAt

`func (o *SettingsResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SettingsResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SettingsResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SettingsResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SettingsResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SettingsResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


