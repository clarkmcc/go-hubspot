# Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of this domain. | 
**Domain** | **string** | The actual domain or sub-domain. e.g. www.hubspot.com | 
**PrimaryLandingPage** | Pointer to **bool** |  | [optional] 
**PrimaryEmail** | Pointer to **bool** |  | [optional] 
**PrimaryBlogPost** | Pointer to **bool** |  | [optional] 
**PrimarySitePage** | Pointer to **bool** |  | [optional] 
**PrimaryKnowledge** | Pointer to **bool** |  | [optional] 
**SecondaryToDomain** | Pointer to **string** |  | [optional] 
**IsResolving** | **bool** | Whether the DNS for this domain is optimally configured for use with HubSpot. | 
**ManuallyMarkedAsResolving** | Pointer to **bool** |  | [optional] 
**IsSslEnabled** | Pointer to **bool** |  | [optional] 
**IsSslOnly** | Pointer to **bool** |  | [optional] 
**IsUsedForBlogPost** | **bool** | Whether the domain is used for CMS blog posts. | 
**IsUsedForSitePage** | **bool** | Whether the domain is used for CMS site pages. | 
**IsUsedForLandingPage** | **bool** | Whether the domain is used for CMS landing pages. | 
**IsUsedForEmail** | **bool** | Whether the domain is used for CMS email web pages. | 
**IsUsedForKnowledge** | **bool** | Whether the domain is used for CMS knowledge pages. | 
**CorrectCname** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDomain

`func NewDomain(id string, domain string, isResolving bool, isUsedForBlogPost bool, isUsedForSitePage bool, isUsedForLandingPage bool, isUsedForEmail bool, isUsedForKnowledge bool, ) *Domain`

NewDomain instantiates a new Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWithDefaults

`func NewDomainWithDefaults() *Domain`

NewDomainWithDefaults instantiates a new Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Domain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Domain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Domain) SetId(v string)`

SetId sets Id field to given value.


### GetDomain

`func (o *Domain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Domain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Domain) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetPrimaryLandingPage

`func (o *Domain) GetPrimaryLandingPage() bool`

GetPrimaryLandingPage returns the PrimaryLandingPage field if non-nil, zero value otherwise.

### GetPrimaryLandingPageOk

`func (o *Domain) GetPrimaryLandingPageOk() (*bool, bool)`

GetPrimaryLandingPageOk returns a tuple with the PrimaryLandingPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLandingPage

`func (o *Domain) SetPrimaryLandingPage(v bool)`

SetPrimaryLandingPage sets PrimaryLandingPage field to given value.

### HasPrimaryLandingPage

`func (o *Domain) HasPrimaryLandingPage() bool`

HasPrimaryLandingPage returns a boolean if a field has been set.

### GetPrimaryEmail

`func (o *Domain) GetPrimaryEmail() bool`

GetPrimaryEmail returns the PrimaryEmail field if non-nil, zero value otherwise.

### GetPrimaryEmailOk

`func (o *Domain) GetPrimaryEmailOk() (*bool, bool)`

GetPrimaryEmailOk returns a tuple with the PrimaryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEmail

`func (o *Domain) SetPrimaryEmail(v bool)`

SetPrimaryEmail sets PrimaryEmail field to given value.

### HasPrimaryEmail

`func (o *Domain) HasPrimaryEmail() bool`

HasPrimaryEmail returns a boolean if a field has been set.

### GetPrimaryBlogPost

`func (o *Domain) GetPrimaryBlogPost() bool`

GetPrimaryBlogPost returns the PrimaryBlogPost field if non-nil, zero value otherwise.

### GetPrimaryBlogPostOk

`func (o *Domain) GetPrimaryBlogPostOk() (*bool, bool)`

GetPrimaryBlogPostOk returns a tuple with the PrimaryBlogPost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryBlogPost

`func (o *Domain) SetPrimaryBlogPost(v bool)`

SetPrimaryBlogPost sets PrimaryBlogPost field to given value.

### HasPrimaryBlogPost

`func (o *Domain) HasPrimaryBlogPost() bool`

HasPrimaryBlogPost returns a boolean if a field has been set.

### GetPrimarySitePage

`func (o *Domain) GetPrimarySitePage() bool`

GetPrimarySitePage returns the PrimarySitePage field if non-nil, zero value otherwise.

### GetPrimarySitePageOk

`func (o *Domain) GetPrimarySitePageOk() (*bool, bool)`

GetPrimarySitePageOk returns a tuple with the PrimarySitePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySitePage

`func (o *Domain) SetPrimarySitePage(v bool)`

SetPrimarySitePage sets PrimarySitePage field to given value.

### HasPrimarySitePage

`func (o *Domain) HasPrimarySitePage() bool`

HasPrimarySitePage returns a boolean if a field has been set.

### GetPrimaryKnowledge

`func (o *Domain) GetPrimaryKnowledge() bool`

GetPrimaryKnowledge returns the PrimaryKnowledge field if non-nil, zero value otherwise.

### GetPrimaryKnowledgeOk

`func (o *Domain) GetPrimaryKnowledgeOk() (*bool, bool)`

GetPrimaryKnowledgeOk returns a tuple with the PrimaryKnowledge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKnowledge

`func (o *Domain) SetPrimaryKnowledge(v bool)`

SetPrimaryKnowledge sets PrimaryKnowledge field to given value.

### HasPrimaryKnowledge

`func (o *Domain) HasPrimaryKnowledge() bool`

HasPrimaryKnowledge returns a boolean if a field has been set.

### GetSecondaryToDomain

`func (o *Domain) GetSecondaryToDomain() string`

GetSecondaryToDomain returns the SecondaryToDomain field if non-nil, zero value otherwise.

### GetSecondaryToDomainOk

`func (o *Domain) GetSecondaryToDomainOk() (*string, bool)`

GetSecondaryToDomainOk returns a tuple with the SecondaryToDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryToDomain

`func (o *Domain) SetSecondaryToDomain(v string)`

SetSecondaryToDomain sets SecondaryToDomain field to given value.

### HasSecondaryToDomain

`func (o *Domain) HasSecondaryToDomain() bool`

HasSecondaryToDomain returns a boolean if a field has been set.

### GetIsResolving

`func (o *Domain) GetIsResolving() bool`

GetIsResolving returns the IsResolving field if non-nil, zero value otherwise.

### GetIsResolvingOk

`func (o *Domain) GetIsResolvingOk() (*bool, bool)`

GetIsResolvingOk returns a tuple with the IsResolving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsResolving

`func (o *Domain) SetIsResolving(v bool)`

SetIsResolving sets IsResolving field to given value.


### GetManuallyMarkedAsResolving

`func (o *Domain) GetManuallyMarkedAsResolving() bool`

GetManuallyMarkedAsResolving returns the ManuallyMarkedAsResolving field if non-nil, zero value otherwise.

### GetManuallyMarkedAsResolvingOk

`func (o *Domain) GetManuallyMarkedAsResolvingOk() (*bool, bool)`

GetManuallyMarkedAsResolvingOk returns a tuple with the ManuallyMarkedAsResolving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyMarkedAsResolving

`func (o *Domain) SetManuallyMarkedAsResolving(v bool)`

SetManuallyMarkedAsResolving sets ManuallyMarkedAsResolving field to given value.

### HasManuallyMarkedAsResolving

`func (o *Domain) HasManuallyMarkedAsResolving() bool`

HasManuallyMarkedAsResolving returns a boolean if a field has been set.

### GetIsSslEnabled

`func (o *Domain) GetIsSslEnabled() bool`

GetIsSslEnabled returns the IsSslEnabled field if non-nil, zero value otherwise.

### GetIsSslEnabledOk

`func (o *Domain) GetIsSslEnabledOk() (*bool, bool)`

GetIsSslEnabledOk returns a tuple with the IsSslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSslEnabled

`func (o *Domain) SetIsSslEnabled(v bool)`

SetIsSslEnabled sets IsSslEnabled field to given value.

### HasIsSslEnabled

`func (o *Domain) HasIsSslEnabled() bool`

HasIsSslEnabled returns a boolean if a field has been set.

### GetIsSslOnly

`func (o *Domain) GetIsSslOnly() bool`

GetIsSslOnly returns the IsSslOnly field if non-nil, zero value otherwise.

### GetIsSslOnlyOk

`func (o *Domain) GetIsSslOnlyOk() (*bool, bool)`

GetIsSslOnlyOk returns a tuple with the IsSslOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSslOnly

`func (o *Domain) SetIsSslOnly(v bool)`

SetIsSslOnly sets IsSslOnly field to given value.

### HasIsSslOnly

`func (o *Domain) HasIsSslOnly() bool`

HasIsSslOnly returns a boolean if a field has been set.

### GetIsUsedForBlogPost

`func (o *Domain) GetIsUsedForBlogPost() bool`

GetIsUsedForBlogPost returns the IsUsedForBlogPost field if non-nil, zero value otherwise.

### GetIsUsedForBlogPostOk

`func (o *Domain) GetIsUsedForBlogPostOk() (*bool, bool)`

GetIsUsedForBlogPostOk returns a tuple with the IsUsedForBlogPost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsedForBlogPost

`func (o *Domain) SetIsUsedForBlogPost(v bool)`

SetIsUsedForBlogPost sets IsUsedForBlogPost field to given value.


### GetIsUsedForSitePage

`func (o *Domain) GetIsUsedForSitePage() bool`

GetIsUsedForSitePage returns the IsUsedForSitePage field if non-nil, zero value otherwise.

### GetIsUsedForSitePageOk

`func (o *Domain) GetIsUsedForSitePageOk() (*bool, bool)`

GetIsUsedForSitePageOk returns a tuple with the IsUsedForSitePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsedForSitePage

`func (o *Domain) SetIsUsedForSitePage(v bool)`

SetIsUsedForSitePage sets IsUsedForSitePage field to given value.


### GetIsUsedForLandingPage

`func (o *Domain) GetIsUsedForLandingPage() bool`

GetIsUsedForLandingPage returns the IsUsedForLandingPage field if non-nil, zero value otherwise.

### GetIsUsedForLandingPageOk

`func (o *Domain) GetIsUsedForLandingPageOk() (*bool, bool)`

GetIsUsedForLandingPageOk returns a tuple with the IsUsedForLandingPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsedForLandingPage

`func (o *Domain) SetIsUsedForLandingPage(v bool)`

SetIsUsedForLandingPage sets IsUsedForLandingPage field to given value.


### GetIsUsedForEmail

`func (o *Domain) GetIsUsedForEmail() bool`

GetIsUsedForEmail returns the IsUsedForEmail field if non-nil, zero value otherwise.

### GetIsUsedForEmailOk

`func (o *Domain) GetIsUsedForEmailOk() (*bool, bool)`

GetIsUsedForEmailOk returns a tuple with the IsUsedForEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsedForEmail

`func (o *Domain) SetIsUsedForEmail(v bool)`

SetIsUsedForEmail sets IsUsedForEmail field to given value.


### GetIsUsedForKnowledge

`func (o *Domain) GetIsUsedForKnowledge() bool`

GetIsUsedForKnowledge returns the IsUsedForKnowledge field if non-nil, zero value otherwise.

### GetIsUsedForKnowledgeOk

`func (o *Domain) GetIsUsedForKnowledgeOk() (*bool, bool)`

GetIsUsedForKnowledgeOk returns a tuple with the IsUsedForKnowledge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsedForKnowledge

`func (o *Domain) SetIsUsedForKnowledge(v bool)`

SetIsUsedForKnowledge sets IsUsedForKnowledge field to given value.


### GetCorrectCname

`func (o *Domain) GetCorrectCname() string`

GetCorrectCname returns the CorrectCname field if non-nil, zero value otherwise.

### GetCorrectCnameOk

`func (o *Domain) GetCorrectCnameOk() (*string, bool)`

GetCorrectCnameOk returns a tuple with the CorrectCname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrectCname

`func (o *Domain) SetCorrectCname(v string)`

SetCorrectCname sets CorrectCname field to given value.

### HasCorrectCname

`func (o *Domain) HasCorrectCname() bool`

HasCorrectCname returns a boolean if a field has been set.

### GetCreated

`func (o *Domain) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Domain) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Domain) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Domain) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *Domain) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Domain) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Domain) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Domain) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


