# MarketingEventEmailSubscriber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InteractionDateTime** | **int64** | The date and time at which the contact subscribed to the event. | 
**Properties** | Pointer to **map[string]string** |  | [optional] 
**Email** | **string** | The email address of the contact in HubSpot to associate with the event. Note that the contact must already exist in HubSpot; a contact will not be created. | 
**ContactProperties** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMarketingEventEmailSubscriber

`func NewMarketingEventEmailSubscriber(interactionDateTime int64, email string, ) *MarketingEventEmailSubscriber`

NewMarketingEventEmailSubscriber instantiates a new MarketingEventEmailSubscriber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingEventEmailSubscriberWithDefaults

`func NewMarketingEventEmailSubscriberWithDefaults() *MarketingEventEmailSubscriber`

NewMarketingEventEmailSubscriberWithDefaults instantiates a new MarketingEventEmailSubscriber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInteractionDateTime

`func (o *MarketingEventEmailSubscriber) GetInteractionDateTime() int64`

GetInteractionDateTime returns the InteractionDateTime field if non-nil, zero value otherwise.

### GetInteractionDateTimeOk

`func (o *MarketingEventEmailSubscriber) GetInteractionDateTimeOk() (*int64, bool)`

GetInteractionDateTimeOk returns a tuple with the InteractionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractionDateTime

`func (o *MarketingEventEmailSubscriber) SetInteractionDateTime(v int64)`

SetInteractionDateTime sets InteractionDateTime field to given value.


### GetProperties

`func (o *MarketingEventEmailSubscriber) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *MarketingEventEmailSubscriber) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *MarketingEventEmailSubscriber) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *MarketingEventEmailSubscriber) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetEmail

`func (o *MarketingEventEmailSubscriber) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MarketingEventEmailSubscriber) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MarketingEventEmailSubscriber) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetContactProperties

`func (o *MarketingEventEmailSubscriber) GetContactProperties() map[string]string`

GetContactProperties returns the ContactProperties field if non-nil, zero value otherwise.

### GetContactPropertiesOk

`func (o *MarketingEventEmailSubscriber) GetContactPropertiesOk() (*map[string]string, bool)`

GetContactPropertiesOk returns a tuple with the ContactProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactProperties

`func (o *MarketingEventEmailSubscriber) SetContactProperties(v map[string]string)`

SetContactProperties sets ContactProperties field to given value.

### HasContactProperties

`func (o *MarketingEventEmailSubscriber) HasContactProperties() bool`

HasContactProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


