# TopLevelActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to [**IFrameActionBody**](IFrameActionBody.md) |  | [optional] 
**Primary** | Pointer to [**IntegratorObjectResultActionsInner**](IntegratorObjectResultActionsInner.md) |  | [optional] 
**Secondary** | [**[]IntegratorObjectResultActionsInner**](IntegratorObjectResultActionsInner.md) |  | 

## Methods

### NewTopLevelActions

`func NewTopLevelActions(secondary []IntegratorObjectResultActionsInner, ) *TopLevelActions`

NewTopLevelActions instantiates a new TopLevelActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopLevelActionsWithDefaults

`func NewTopLevelActionsWithDefaults() *TopLevelActions`

NewTopLevelActionsWithDefaults instantiates a new TopLevelActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *TopLevelActions) GetSettings() IFrameActionBody`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *TopLevelActions) GetSettingsOk() (*IFrameActionBody, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *TopLevelActions) SetSettings(v IFrameActionBody)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *TopLevelActions) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetPrimary

`func (o *TopLevelActions) GetPrimary() IntegratorObjectResultActionsInner`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *TopLevelActions) GetPrimaryOk() (*IntegratorObjectResultActionsInner, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *TopLevelActions) SetPrimary(v IntegratorObjectResultActionsInner)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *TopLevelActions) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetSecondary

`func (o *TopLevelActions) GetSecondary() []IntegratorObjectResultActionsInner`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *TopLevelActions) GetSecondaryOk() (*[]IntegratorObjectResultActionsInner, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *TopLevelActions) SetSecondary(v []IntegratorObjectResultActionsInner)`

SetSecondary sets Secondary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


