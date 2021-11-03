# PublicSubscriptionStatusesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipient** | **string** | Email address of the contact. | 
**SubscriptionStatuses** | [**[]PublicSubscriptionStatus**](PublicSubscriptionStatus.md) | A list of all of the contact&#39;s subscriptions statuses. | 

## Methods

### NewPublicSubscriptionStatusesResponse

`func NewPublicSubscriptionStatusesResponse(recipient string, subscriptionStatuses []PublicSubscriptionStatus, ) *PublicSubscriptionStatusesResponse`

NewPublicSubscriptionStatusesResponse instantiates a new PublicSubscriptionStatusesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSubscriptionStatusesResponseWithDefaults

`func NewPublicSubscriptionStatusesResponseWithDefaults() *PublicSubscriptionStatusesResponse`

NewPublicSubscriptionStatusesResponseWithDefaults instantiates a new PublicSubscriptionStatusesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipient

`func (o *PublicSubscriptionStatusesResponse) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *PublicSubscriptionStatusesResponse) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *PublicSubscriptionStatusesResponse) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetSubscriptionStatuses

`func (o *PublicSubscriptionStatusesResponse) GetSubscriptionStatuses() []PublicSubscriptionStatus`

GetSubscriptionStatuses returns the SubscriptionStatuses field if non-nil, zero value otherwise.

### GetSubscriptionStatusesOk

`func (o *PublicSubscriptionStatusesResponse) GetSubscriptionStatusesOk() (*[]PublicSubscriptionStatus, bool)`

GetSubscriptionStatusesOk returns a tuple with the SubscriptionStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatuses

`func (o *PublicSubscriptionStatusesResponse) SetSubscriptionStatuses(v []PublicSubscriptionStatus)`

SetSubscriptionStatuses sets SubscriptionStatuses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


