# MarketingEventEmailSubscriber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InteractionDateTime** | **int64** | The date and time at which the contact subscribed to the event. | 
**Email** | **string** | The email address of the contact in HubSpot to associate with the event. Note that the contact must already exist in HubSpot; a contact will not be created. | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


