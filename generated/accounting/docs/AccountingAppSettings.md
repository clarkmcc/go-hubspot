# AccountingAppSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **int32** | The ID of the accounting app. This is the identifier of the application created in your HubSpot developer portal. | 
**Urls** | [**AccountingAppUrls**](AccountingAppUrls.md) |  | 
**Features** | Pointer to [**AccountingFeatures**](AccountingFeatures.md) |  | [optional] 

## Methods

### NewAccountingAppSettings

`func NewAccountingAppSettings(appId int32, urls AccountingAppUrls, ) *AccountingAppSettings`

NewAccountingAppSettings instantiates a new AccountingAppSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountingAppSettingsWithDefaults

`func NewAccountingAppSettingsWithDefaults() *AccountingAppSettings`

NewAccountingAppSettingsWithDefaults instantiates a new AccountingAppSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *AccountingAppSettings) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AccountingAppSettings) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AccountingAppSettings) SetAppId(v int32)`

SetAppId sets AppId field to given value.


### GetUrls

`func (o *AccountingAppSettings) GetUrls() AccountingAppUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *AccountingAppSettings) GetUrlsOk() (*AccountingAppUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *AccountingAppSettings) SetUrls(v AccountingAppUrls)`

SetUrls sets Urls field to given value.


### GetFeatures

`func (o *AccountingAppSettings) GetFeatures() AccountingFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *AccountingAppSettings) GetFeaturesOk() (*AccountingFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *AccountingAppSettings) SetFeatures(v AccountingFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *AccountingAppSettings) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


