# SmtpApiTokenView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | User name to log into the HubSpot SMTP server. | 
**CreatedBy** | **string** | Email address of the user that sent the token creation request. | 
**Password** | Pointer to **string** | Password used to log into the HubSpot SMTP server. | [optional] 
**EmailCampaignId** | **string** | Identifier assigned to the campaign provided in the token creation request. | 
**CreatedAt** | **time.Time** | Timestamp generated when a token is created. | 
**CreateContact** | **bool** | Indicates whether a contact should be created for email recipients. | 
**CampaignName** | **string** | A name for the campaign tied to the token. | 

## Methods

### NewSmtpApiTokenView

`func NewSmtpApiTokenView(id string, createdBy string, emailCampaignId string, createdAt time.Time, createContact bool, campaignName string, ) *SmtpApiTokenView`

NewSmtpApiTokenView instantiates a new SmtpApiTokenView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpApiTokenViewWithDefaults

`func NewSmtpApiTokenViewWithDefaults() *SmtpApiTokenView`

NewSmtpApiTokenViewWithDefaults instantiates a new SmtpApiTokenView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SmtpApiTokenView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SmtpApiTokenView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SmtpApiTokenView) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *SmtpApiTokenView) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SmtpApiTokenView) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SmtpApiTokenView) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetPassword

`func (o *SmtpApiTokenView) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SmtpApiTokenView) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SmtpApiTokenView) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SmtpApiTokenView) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEmailCampaignId

`func (o *SmtpApiTokenView) GetEmailCampaignId() string`

GetEmailCampaignId returns the EmailCampaignId field if non-nil, zero value otherwise.

### GetEmailCampaignIdOk

`func (o *SmtpApiTokenView) GetEmailCampaignIdOk() (*string, bool)`

GetEmailCampaignIdOk returns a tuple with the EmailCampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCampaignId

`func (o *SmtpApiTokenView) SetEmailCampaignId(v string)`

SetEmailCampaignId sets EmailCampaignId field to given value.


### GetCreatedAt

`func (o *SmtpApiTokenView) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SmtpApiTokenView) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SmtpApiTokenView) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreateContact

`func (o *SmtpApiTokenView) GetCreateContact() bool`

GetCreateContact returns the CreateContact field if non-nil, zero value otherwise.

### GetCreateContactOk

`func (o *SmtpApiTokenView) GetCreateContactOk() (*bool, bool)`

GetCreateContactOk returns a tuple with the CreateContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateContact

`func (o *SmtpApiTokenView) SetCreateContact(v bool)`

SetCreateContact sets CreateContact field to given value.


### GetCampaignName

`func (o *SmtpApiTokenView) GetCampaignName() string`

GetCampaignName returns the CampaignName field if non-nil, zero value otherwise.

### GetCampaignNameOk

`func (o *SmtpApiTokenView) GetCampaignNameOk() (*string, bool)`

GetCampaignNameOk returns a tuple with the CampaignName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignName

`func (o *SmtpApiTokenView) SetCampaignName(v string)`

SetCampaignName sets CampaignName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


