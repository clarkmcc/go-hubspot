# MarketingEventSubscriber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InteractionDateTime** | **int64** | The date and time at which the contact subscribed to the event. | 
**Properties** | Pointer to **map[string]string** |  | [optional] 
**Vid** | Pointer to **int32** |  | [optional] 

## Methods

### NewMarketingEventSubscriber

`func NewMarketingEventSubscriber(interactionDateTime int64, ) *MarketingEventSubscriber`

NewMarketingEventSubscriber instantiates a new MarketingEventSubscriber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingEventSubscriberWithDefaults

`func NewMarketingEventSubscriberWithDefaults() *MarketingEventSubscriber`

NewMarketingEventSubscriberWithDefaults instantiates a new MarketingEventSubscriber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInteractionDateTime

`func (o *MarketingEventSubscriber) GetInteractionDateTime() int64`

GetInteractionDateTime returns the InteractionDateTime field if non-nil, zero value otherwise.

### GetInteractionDateTimeOk

`func (o *MarketingEventSubscriber) GetInteractionDateTimeOk() (*int64, bool)`

GetInteractionDateTimeOk returns a tuple with the InteractionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractionDateTime

`func (o *MarketingEventSubscriber) SetInteractionDateTime(v int64)`

SetInteractionDateTime sets InteractionDateTime field to given value.


### GetProperties

`func (o *MarketingEventSubscriber) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *MarketingEventSubscriber) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *MarketingEventSubscriber) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *MarketingEventSubscriber) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetVid

`func (o *MarketingEventSubscriber) GetVid() int32`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *MarketingEventSubscriber) GetVidOk() (*int32, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *MarketingEventSubscriber) SetVid(v int32)`

SetVid sets Vid field to given value.

### HasVid

`func (o *MarketingEventSubscriber) HasVid() bool`

HasVid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


