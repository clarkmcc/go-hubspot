# SmtpApiTokenRequestEgg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateContact** | **bool** | Indicates whether a contact should be created for email recipients. | 
**CampaignName** | **string** | A name for the campaign tied to the SMTP API token. | 

## Methods

### NewSmtpApiTokenRequestEgg

`func NewSmtpApiTokenRequestEgg(createContact bool, campaignName string, ) *SmtpApiTokenRequestEgg`

NewSmtpApiTokenRequestEgg instantiates a new SmtpApiTokenRequestEgg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpApiTokenRequestEggWithDefaults

`func NewSmtpApiTokenRequestEggWithDefaults() *SmtpApiTokenRequestEgg`

NewSmtpApiTokenRequestEggWithDefaults instantiates a new SmtpApiTokenRequestEgg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateContact

`func (o *SmtpApiTokenRequestEgg) GetCreateContact() bool`

GetCreateContact returns the CreateContact field if non-nil, zero value otherwise.

### GetCreateContactOk

`func (o *SmtpApiTokenRequestEgg) GetCreateContactOk() (*bool, bool)`

GetCreateContactOk returns a tuple with the CreateContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateContact

`func (o *SmtpApiTokenRequestEgg) SetCreateContact(v bool)`

SetCreateContact sets CreateContact field to given value.


### GetCampaignName

`func (o *SmtpApiTokenRequestEgg) GetCampaignName() string`

GetCampaignName returns the CampaignName field if non-nil, zero value otherwise.

### GetCampaignNameOk

`func (o *SmtpApiTokenRequestEgg) GetCampaignNameOk() (*string, bool)`

GetCampaignNameOk returns a tuple with the CampaignName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignName

`func (o *SmtpApiTokenRequestEgg) SetCampaignName(v string)`

SetCampaignName sets CampaignName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


