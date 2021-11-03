# MarketingEventExternalUniqueIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **int32** | The id of the application that created the marketing event in HubSpot. | 
**ExternalAccountId** | **string** | The accountId that is associated with this marketing event in the external event application. | 
**ExternalEventId** | **string** | The id of the marketing event in the external event application. | 

## Methods

### NewMarketingEventExternalUniqueIdentifier

`func NewMarketingEventExternalUniqueIdentifier(appId int32, externalAccountId string, externalEventId string, ) *MarketingEventExternalUniqueIdentifier`

NewMarketingEventExternalUniqueIdentifier instantiates a new MarketingEventExternalUniqueIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingEventExternalUniqueIdentifierWithDefaults

`func NewMarketingEventExternalUniqueIdentifierWithDefaults() *MarketingEventExternalUniqueIdentifier`

NewMarketingEventExternalUniqueIdentifierWithDefaults instantiates a new MarketingEventExternalUniqueIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *MarketingEventExternalUniqueIdentifier) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *MarketingEventExternalUniqueIdentifier) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *MarketingEventExternalUniqueIdentifier) SetAppId(v int32)`

SetAppId sets AppId field to given value.


### GetExternalAccountId

`func (o *MarketingEventExternalUniqueIdentifier) GetExternalAccountId() string`

GetExternalAccountId returns the ExternalAccountId field if non-nil, zero value otherwise.

### GetExternalAccountIdOk

`func (o *MarketingEventExternalUniqueIdentifier) GetExternalAccountIdOk() (*string, bool)`

GetExternalAccountIdOk returns a tuple with the ExternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAccountId

`func (o *MarketingEventExternalUniqueIdentifier) SetExternalAccountId(v string)`

SetExternalAccountId sets ExternalAccountId field to given value.


### GetExternalEventId

`func (o *MarketingEventExternalUniqueIdentifier) GetExternalEventId() string`

GetExternalEventId returns the ExternalEventId field if non-nil, zero value otherwise.

### GetExternalEventIdOk

`func (o *MarketingEventExternalUniqueIdentifier) GetExternalEventIdOk() (*string, bool)`

GetExternalEventIdOk returns a tuple with the ExternalEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEventId

`func (o *MarketingEventExternalUniqueIdentifier) SetExternalEventId(v string)`

SetExternalEventId sets ExternalEventId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


