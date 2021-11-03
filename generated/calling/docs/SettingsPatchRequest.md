# SettingsPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of your calling service to display to users. | [optional] 
**Url** | Pointer to **string** | The URL to your phone/calling UI, built with the [Calling SDK](#). | [optional] 
**Height** | Pointer to **int32** | The target height of the iframe that will contain your phone/calling UI. | [optional] 
**Width** | Pointer to **int32** | The target width of the iframe that will contain your phone/calling UI. | [optional] 
**IsReady** | Pointer to **bool** | When true, your service will appear as an option under the *Call* action in contact records of connected accounts. | [optional] 
**SupportsCustomObjects** | Pointer to **bool** | When true, you are indicating that your service is compatible with engagement v2 service and can be used with custom objects. | [optional] 

## Methods

### NewSettingsPatchRequest

`func NewSettingsPatchRequest() *SettingsPatchRequest`

NewSettingsPatchRequest instantiates a new SettingsPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsPatchRequestWithDefaults

`func NewSettingsPatchRequestWithDefaults() *SettingsPatchRequest`

NewSettingsPatchRequestWithDefaults instantiates a new SettingsPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SettingsPatchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SettingsPatchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SettingsPatchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SettingsPatchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *SettingsPatchRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SettingsPatchRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SettingsPatchRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SettingsPatchRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHeight

`func (o *SettingsPatchRequest) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *SettingsPatchRequest) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *SettingsPatchRequest) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *SettingsPatchRequest) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetWidth

`func (o *SettingsPatchRequest) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *SettingsPatchRequest) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *SettingsPatchRequest) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *SettingsPatchRequest) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetIsReady

`func (o *SettingsPatchRequest) GetIsReady() bool`

GetIsReady returns the IsReady field if non-nil, zero value otherwise.

### GetIsReadyOk

`func (o *SettingsPatchRequest) GetIsReadyOk() (*bool, bool)`

GetIsReadyOk returns a tuple with the IsReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReady

`func (o *SettingsPatchRequest) SetIsReady(v bool)`

SetIsReady sets IsReady field to given value.

### HasIsReady

`func (o *SettingsPatchRequest) HasIsReady() bool`

HasIsReady returns a boolean if a field has been set.

### GetSupportsCustomObjects

`func (o *SettingsPatchRequest) GetSupportsCustomObjects() bool`

GetSupportsCustomObjects returns the SupportsCustomObjects field if non-nil, zero value otherwise.

### GetSupportsCustomObjectsOk

`func (o *SettingsPatchRequest) GetSupportsCustomObjectsOk() (*bool, bool)`

GetSupportsCustomObjectsOk returns a tuple with the SupportsCustomObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsCustomObjects

`func (o *SettingsPatchRequest) SetSupportsCustomObjects(v bool)`

SetSupportsCustomObjects sets SupportsCustomObjects field to given value.

### HasSupportsCustomObjects

`func (o *SettingsPatchRequest) HasSupportsCustomObjects() bool`

HasSupportsCustomObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


