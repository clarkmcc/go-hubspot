# PublicUpdateSubscriptionStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | **string** | Contact&#39;s email address. | 
**SubscriptionId** | **string** | ID of the subscription being updated for the contact. | 
**LegalBasis** | Pointer to **string** | Legal basis for updating the contact&#39;s status (required for GDPR enabled portals). | [optional] 
**LegalBasisExplanation** | Pointer to **string** | A more detailed explanation to go with the legal basis (required for GDPR enabled portals). | [optional] 

## Methods

### NewPublicUpdateSubscriptionStatusRequest

`func NewPublicUpdateSubscriptionStatusRequest(emailAddress string, subscriptionId string, ) *PublicUpdateSubscriptionStatusRequest`

NewPublicUpdateSubscriptionStatusRequest instantiates a new PublicUpdateSubscriptionStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicUpdateSubscriptionStatusRequestWithDefaults

`func NewPublicUpdateSubscriptionStatusRequestWithDefaults() *PublicUpdateSubscriptionStatusRequest`

NewPublicUpdateSubscriptionStatusRequestWithDefaults instantiates a new PublicUpdateSubscriptionStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *PublicUpdateSubscriptionStatusRequest) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *PublicUpdateSubscriptionStatusRequest) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *PublicUpdateSubscriptionStatusRequest) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetSubscriptionId

`func (o *PublicUpdateSubscriptionStatusRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PublicUpdateSubscriptionStatusRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PublicUpdateSubscriptionStatusRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetLegalBasis

`func (o *PublicUpdateSubscriptionStatusRequest) GetLegalBasis() string`

GetLegalBasis returns the LegalBasis field if non-nil, zero value otherwise.

### GetLegalBasisOk

`func (o *PublicUpdateSubscriptionStatusRequest) GetLegalBasisOk() (*string, bool)`

GetLegalBasisOk returns a tuple with the LegalBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalBasis

`func (o *PublicUpdateSubscriptionStatusRequest) SetLegalBasis(v string)`

SetLegalBasis sets LegalBasis field to given value.

### HasLegalBasis

`func (o *PublicUpdateSubscriptionStatusRequest) HasLegalBasis() bool`

HasLegalBasis returns a boolean if a field has been set.

### GetLegalBasisExplanation

`func (o *PublicUpdateSubscriptionStatusRequest) GetLegalBasisExplanation() string`

GetLegalBasisExplanation returns the LegalBasisExplanation field if non-nil, zero value otherwise.

### GetLegalBasisExplanationOk

`func (o *PublicUpdateSubscriptionStatusRequest) GetLegalBasisExplanationOk() (*string, bool)`

GetLegalBasisExplanationOk returns a tuple with the LegalBasisExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalBasisExplanation

`func (o *PublicUpdateSubscriptionStatusRequest) SetLegalBasisExplanation(v string)`

SetLegalBasisExplanation sets LegalBasisExplanation field to given value.

### HasLegalBasisExplanation

`func (o *PublicUpdateSubscriptionStatusRequest) HasLegalBasisExplanation() bool`

HasLegalBasisExplanation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


