# PublicSubscriptionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID for the subscription. | 
**Name** | **string** | The name of the subscription. | 
**Description** | **string** | A description of the subscription. | 
**Status** | **string** | Whether the contact is subscribed. | 
**SourceOfStatus** | **string** | Where the status is determined from e.g. PORTAL_WIDE_STATUS if the contact opted out from the portal. | 
**BrandId** | Pointer to **int64** | The ID of the brand that the subscription is associated with, if there is one. | [optional] 
**PreferenceGroupName** | Pointer to **string** | The name of the preferences group that the subscription is associated with. | [optional] 
**LegalBasis** | Pointer to **string** | The legal reason for the current status of the subscription. | [optional] 
**LegalBasisExplanation** | Pointer to **string** | A more detailed explanation to go with the legal basis. | [optional] 

## Methods

### NewPublicSubscriptionStatus

`func NewPublicSubscriptionStatus(id string, name string, description string, status string, sourceOfStatus string, ) *PublicSubscriptionStatus`

NewPublicSubscriptionStatus instantiates a new PublicSubscriptionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSubscriptionStatusWithDefaults

`func NewPublicSubscriptionStatusWithDefaults() *PublicSubscriptionStatus`

NewPublicSubscriptionStatusWithDefaults instantiates a new PublicSubscriptionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PublicSubscriptionStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicSubscriptionStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicSubscriptionStatus) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PublicSubscriptionStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicSubscriptionStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicSubscriptionStatus) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PublicSubscriptionStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PublicSubscriptionStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PublicSubscriptionStatus) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStatus

`func (o *PublicSubscriptionStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicSubscriptionStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicSubscriptionStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSourceOfStatus

`func (o *PublicSubscriptionStatus) GetSourceOfStatus() string`

GetSourceOfStatus returns the SourceOfStatus field if non-nil, zero value otherwise.

### GetSourceOfStatusOk

`func (o *PublicSubscriptionStatus) GetSourceOfStatusOk() (*string, bool)`

GetSourceOfStatusOk returns a tuple with the SourceOfStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfStatus

`func (o *PublicSubscriptionStatus) SetSourceOfStatus(v string)`

SetSourceOfStatus sets SourceOfStatus field to given value.


### GetBrandId

`func (o *PublicSubscriptionStatus) GetBrandId() int64`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *PublicSubscriptionStatus) GetBrandIdOk() (*int64, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *PublicSubscriptionStatus) SetBrandId(v int64)`

SetBrandId sets BrandId field to given value.

### HasBrandId

`func (o *PublicSubscriptionStatus) HasBrandId() bool`

HasBrandId returns a boolean if a field has been set.

### GetPreferenceGroupName

`func (o *PublicSubscriptionStatus) GetPreferenceGroupName() string`

GetPreferenceGroupName returns the PreferenceGroupName field if non-nil, zero value otherwise.

### GetPreferenceGroupNameOk

`func (o *PublicSubscriptionStatus) GetPreferenceGroupNameOk() (*string, bool)`

GetPreferenceGroupNameOk returns a tuple with the PreferenceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferenceGroupName

`func (o *PublicSubscriptionStatus) SetPreferenceGroupName(v string)`

SetPreferenceGroupName sets PreferenceGroupName field to given value.

### HasPreferenceGroupName

`func (o *PublicSubscriptionStatus) HasPreferenceGroupName() bool`

HasPreferenceGroupName returns a boolean if a field has been set.

### GetLegalBasis

`func (o *PublicSubscriptionStatus) GetLegalBasis() string`

GetLegalBasis returns the LegalBasis field if non-nil, zero value otherwise.

### GetLegalBasisOk

`func (o *PublicSubscriptionStatus) GetLegalBasisOk() (*string, bool)`

GetLegalBasisOk returns a tuple with the LegalBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalBasis

`func (o *PublicSubscriptionStatus) SetLegalBasis(v string)`

SetLegalBasis sets LegalBasis field to given value.

### HasLegalBasis

`func (o *PublicSubscriptionStatus) HasLegalBasis() bool`

HasLegalBasis returns a boolean if a field has been set.

### GetLegalBasisExplanation

`func (o *PublicSubscriptionStatus) GetLegalBasisExplanation() string`

GetLegalBasisExplanation returns the LegalBasisExplanation field if non-nil, zero value otherwise.

### GetLegalBasisExplanationOk

`func (o *PublicSubscriptionStatus) GetLegalBasisExplanationOk() (*string, bool)`

GetLegalBasisExplanationOk returns a tuple with the LegalBasisExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalBasisExplanation

`func (o *PublicSubscriptionStatus) SetLegalBasisExplanation(v string)`

SetLegalBasisExplanation sets LegalBasisExplanation field to given value.

### HasLegalBasisExplanation

`func (o *PublicSubscriptionStatus) HasLegalBasisExplanation() bool`

HasLegalBasisExplanation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


