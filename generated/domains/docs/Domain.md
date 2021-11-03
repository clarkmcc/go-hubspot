# Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortalId** | **int32** |  | 
**Id** | **int64** |  | 
**Created** | **int64** |  | 
**Updated** | **int64** |  | 
**Domain** | **string** |  | 
**PrimaryLandingPage** | **bool** |  | 
**PrimaryEmail** | **bool** |  | 
**PrimaryBlog** | **bool** |  | 
**PrimaryBlogPost** | **bool** |  | 
**PrimarySitePage** | **bool** |  | 
**PrimaryKnowledge** | **bool** |  | 
**PrimaryLegacyPage** | **bool** |  | 
**PrimaryClickTracking** | **bool** |  | 
**FullCategoryKey** | **string** |  | 
**SecondaryToDomain** | **string** |  | 
**IsResolving** | **bool** |  | 
**IsDnsCorrect** | **bool** |  | 
**ManuallyMarkedAsResolving** | **bool** |  | 
**ConsecutiveNonResolvingCount** | **int32** |  | 
**SslCname** | **string** |  | 
**IsSslEnabled** | **bool** |  | 
**IsSslOnly** | **bool** |  | 
**CertificateId** | **int64** |  | 
**SslRequestId** | **int64** |  | 
**IsUsedForBlogPost** | **bool** |  | 
**IsUsedForSitePage** | **bool** |  | 
**IsUsedForLandingPage** | **bool** |  | 
**IsUsedForEmail** | **bool** |  | 
**IsUsedForKnowledge** | **bool** |  | 
**SetupTaskId** | **int64** |  | 
**IsSetupComplete** | **bool** |  | 
**SetUpLanguage** | **string** |  | 
**TeamIds** | **[]int64** |  | 
**ActualCname** | **string** |  | 
**CorrectCname** | **string** |  | 
**ActualIp** | **string** |  | 
**ApexResolutionStatus** | **string** |  | 
**ApexDomain** | **string** |  | 
**PublicSuffix** | **string** |  | 
**ApexIpAddresses** | **[]string** |  | 
**SiteId** | **int64** |  | 
**BrandId** | **int64** |  | 
**Deletable** | **bool** |  | 
**DomainCdnConfig** | [**DomainCdnConfig**](DomainCdnConfig.md) |  | 
**SetupInfo** | [**DomainSetupInfo**](DomainSetupInfo.md) |  | 
**DerivedBrandName** | **string** |  | 
**CreatedById** | **int32** |  | 
**UpdatedById** | **int32** |  | 
**Label** | **string** |  | 
**IsAnyPrimary** | **bool** |  | 
**IsLegacyDomain** | **bool** |  | 
**IsInternalDomain** | **bool** |  | 
**IsResolvingInternalProperty** | **bool** |  | 
**IsResolvingIgnoringManuallyMarkedAsResolving** | **bool** |  | 
**IsUsedForAnyContentType** | **bool** |  | 
**IsLegacy** | **bool** |  | 
**AuthorAt** | **int64** |  | 
**CosObjectType** | **string** |  | 
**CdnPurgeEmbargoTime** | **int64** |  | 
**IsStagingDomain** | **bool** |  | 

## Methods

### NewDomain

`func NewDomain(portalId int32, id int64, created int64, updated int64, domain string, primaryLandingPage bool, primaryEmail bool, primaryBlog bool, primaryBlogPost bool, primarySitePage bool, primaryKnowledge bool, primaryLegacyPage bool, primaryClickTracking bool, fullCategoryKey string, secondaryToDomain string, isResolving bool, isDnsCorrect bool, manuallyMarkedAsResolving bool, consecutiveNonResolvingCount int32, sslCname string, isSslEnabled bool, isSslOnly bool, certificateId int64, sslRequestId int64, isUsedForBlogPost bool, isUsedForSitePage bool, isUsedForLandingPage bool, isUsedForEmail bool, isUsedForKnowledge bool, setupTaskId int64, isSetupComplete bool, setUpLanguage string, teamIds []int64, actualCname string, correctCname string, actualIp string, apexResolutionStatus string, apexDomain string, publicSuffix string, apexIpAddresses []string, siteId int64, brandId int64, deletable bool, domainCdnConfig DomainCdnConfig, setupInfo DomainSetupInfo, derivedBrandName string, createdById int32, updatedById int32, label string, isAnyPrimary bool, isLegacyDomain bool, isInternalDomain bool, isResolvingInternalProperty bool, isResolvingIgnoringManuallyMarkedAsResolving bool, isUsedForAnyContentType bool, isLegacy bool, authorAt int64, cosObjectType string, cdnPurgeEmbargoTime int64, isStagingDomain bool, ) *Domain`

NewDomain instantiates a new Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWithDefaults

`func NewDomainWithDefaults() *Domain`

NewDomainWithDefaults instantiates a new Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortalId

`func (o *Domain) GetPortalId() int32`

GetPortalId returns the PortalId field if non-nil, zero value otherwise.

### GetPortalIdOk

`func (o *Domain) GetPortalIdOk() (*int32, bool)`

GetPortalIdOk returns a tuple with the PortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalId

`func (o *Domain) SetPortalId(v int32)`

SetPortalId sets PortalId field to given value.


### GetId

`func (o *Domain) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Domain) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Domain) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *Domain) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Domain) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Domain) SetCreated(v int64)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *Domain) GetUpdated() int64`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Domain) GetUpdatedOk() (*int64, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Domain) SetUpdated(v int64)`

SetUpdated sets Updated field to given value.


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


### GetPrimaryBlog

`func (o *Domain) GetPrimaryBlog() bool`

GetPrimaryBlog returns the PrimaryBlog field if non-nil, zero value otherwise.

### GetPrimaryBlogOk

`func (o *Domain) GetPrimaryBlogOk() (*bool, bool)`

GetPrimaryBlogOk returns a tuple with the PrimaryBlog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryBlog

`func (o *Domain) SetPrimaryBlog(v bool)`

SetPrimaryBlog sets PrimaryBlog field to given value.


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


### GetPrimaryLegacyPage

`func (o *Domain) GetPrimaryLegacyPage() bool`

GetPrimaryLegacyPage returns the PrimaryLegacyPage field if non-nil, zero value otherwise.

### GetPrimaryLegacyPageOk

`func (o *Domain) GetPrimaryLegacyPageOk() (*bool, bool)`

GetPrimaryLegacyPageOk returns a tuple with the PrimaryLegacyPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLegacyPage

`func (o *Domain) SetPrimaryLegacyPage(v bool)`

SetPrimaryLegacyPage sets PrimaryLegacyPage field to given value.


### GetPrimaryClickTracking

`func (o *Domain) GetPrimaryClickTracking() bool`

GetPrimaryClickTracking returns the PrimaryClickTracking field if non-nil, zero value otherwise.

### GetPrimaryClickTrackingOk

`func (o *Domain) GetPrimaryClickTrackingOk() (*bool, bool)`

GetPrimaryClickTrackingOk returns a tuple with the PrimaryClickTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryClickTracking

`func (o *Domain) SetPrimaryClickTracking(v bool)`

SetPrimaryClickTracking sets PrimaryClickTracking field to given value.


### GetFullCategoryKey

`func (o *Domain) GetFullCategoryKey() string`

GetFullCategoryKey returns the FullCategoryKey field if non-nil, zero value otherwise.

### GetFullCategoryKeyOk

`func (o *Domain) GetFullCategoryKeyOk() (*string, bool)`

GetFullCategoryKeyOk returns a tuple with the FullCategoryKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullCategoryKey

`func (o *Domain) SetFullCategoryKey(v string)`

SetFullCategoryKey sets FullCategoryKey field to given value.


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


### GetIsDnsCorrect

`func (o *Domain) GetIsDnsCorrect() bool`

GetIsDnsCorrect returns the IsDnsCorrect field if non-nil, zero value otherwise.

### GetIsDnsCorrectOk

`func (o *Domain) GetIsDnsCorrectOk() (*bool, bool)`

GetIsDnsCorrectOk returns a tuple with the IsDnsCorrect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDnsCorrect

`func (o *Domain) SetIsDnsCorrect(v bool)`

SetIsDnsCorrect sets IsDnsCorrect field to given value.


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


### GetConsecutiveNonResolvingCount

`func (o *Domain) GetConsecutiveNonResolvingCount() int32`

GetConsecutiveNonResolvingCount returns the ConsecutiveNonResolvingCount field if non-nil, zero value otherwise.

### GetConsecutiveNonResolvingCountOk

`func (o *Domain) GetConsecutiveNonResolvingCountOk() (*int32, bool)`

GetConsecutiveNonResolvingCountOk returns a tuple with the ConsecutiveNonResolvingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveNonResolvingCount

`func (o *Domain) SetConsecutiveNonResolvingCount(v int32)`

SetConsecutiveNonResolvingCount sets ConsecutiveNonResolvingCount field to given value.


### GetSslCname

`func (o *Domain) GetSslCname() string`

GetSslCname returns the SslCname field if non-nil, zero value otherwise.

### GetSslCnameOk

`func (o *Domain) GetSslCnameOk() (*string, bool)`

GetSslCnameOk returns a tuple with the SslCname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCname

`func (o *Domain) SetSslCname(v string)`

SetSslCname sets SslCname field to given value.


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


### GetCertificateId

`func (o *Domain) GetCertificateId() int64`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *Domain) GetCertificateIdOk() (*int64, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *Domain) SetCertificateId(v int64)`

SetCertificateId sets CertificateId field to given value.


### GetSslRequestId

`func (o *Domain) GetSslRequestId() int64`

GetSslRequestId returns the SslRequestId field if non-nil, zero value otherwise.

### GetSslRequestIdOk

`func (o *Domain) GetSslRequestIdOk() (*int64, bool)`

GetSslRequestIdOk returns a tuple with the SslRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslRequestId

`func (o *Domain) SetSslRequestId(v int64)`

SetSslRequestId sets SslRequestId field to given value.


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


### GetSetupTaskId

`func (o *Domain) GetSetupTaskId() int64`

GetSetupTaskId returns the SetupTaskId field if non-nil, zero value otherwise.

### GetSetupTaskIdOk

`func (o *Domain) GetSetupTaskIdOk() (*int64, bool)`

GetSetupTaskIdOk returns a tuple with the SetupTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupTaskId

`func (o *Domain) SetSetupTaskId(v int64)`

SetSetupTaskId sets SetupTaskId field to given value.


### GetIsSetupComplete

`func (o *Domain) GetIsSetupComplete() bool`

GetIsSetupComplete returns the IsSetupComplete field if non-nil, zero value otherwise.

### GetIsSetupCompleteOk

`func (o *Domain) GetIsSetupCompleteOk() (*bool, bool)`

GetIsSetupCompleteOk returns a tuple with the IsSetupComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSetupComplete

`func (o *Domain) SetIsSetupComplete(v bool)`

SetIsSetupComplete sets IsSetupComplete field to given value.


### GetSetUpLanguage

`func (o *Domain) GetSetUpLanguage() string`

GetSetUpLanguage returns the SetUpLanguage field if non-nil, zero value otherwise.

### GetSetUpLanguageOk

`func (o *Domain) GetSetUpLanguageOk() (*string, bool)`

GetSetUpLanguageOk returns a tuple with the SetUpLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetUpLanguage

`func (o *Domain) SetSetUpLanguage(v string)`

SetSetUpLanguage sets SetUpLanguage field to given value.


### GetTeamIds

`func (o *Domain) GetTeamIds() []int64`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *Domain) GetTeamIdsOk() (*[]int64, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *Domain) SetTeamIds(v []int64)`

SetTeamIds sets TeamIds field to given value.


### GetActualCname

`func (o *Domain) GetActualCname() string`

GetActualCname returns the ActualCname field if non-nil, zero value otherwise.

### GetActualCnameOk

`func (o *Domain) GetActualCnameOk() (*string, bool)`

GetActualCnameOk returns a tuple with the ActualCname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualCname

`func (o *Domain) SetActualCname(v string)`

SetActualCname sets ActualCname field to given value.


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


### GetActualIp

`func (o *Domain) GetActualIp() string`

GetActualIp returns the ActualIp field if non-nil, zero value otherwise.

### GetActualIpOk

`func (o *Domain) GetActualIpOk() (*string, bool)`

GetActualIpOk returns a tuple with the ActualIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualIp

`func (o *Domain) SetActualIp(v string)`

SetActualIp sets ActualIp field to given value.


### GetApexResolutionStatus

`func (o *Domain) GetApexResolutionStatus() string`

GetApexResolutionStatus returns the ApexResolutionStatus field if non-nil, zero value otherwise.

### GetApexResolutionStatusOk

`func (o *Domain) GetApexResolutionStatusOk() (*string, bool)`

GetApexResolutionStatusOk returns a tuple with the ApexResolutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApexResolutionStatus

`func (o *Domain) SetApexResolutionStatus(v string)`

SetApexResolutionStatus sets ApexResolutionStatus field to given value.


### GetApexDomain

`func (o *Domain) GetApexDomain() string`

GetApexDomain returns the ApexDomain field if non-nil, zero value otherwise.

### GetApexDomainOk

`func (o *Domain) GetApexDomainOk() (*string, bool)`

GetApexDomainOk returns a tuple with the ApexDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApexDomain

`func (o *Domain) SetApexDomain(v string)`

SetApexDomain sets ApexDomain field to given value.


### GetPublicSuffix

`func (o *Domain) GetPublicSuffix() string`

GetPublicSuffix returns the PublicSuffix field if non-nil, zero value otherwise.

### GetPublicSuffixOk

`func (o *Domain) GetPublicSuffixOk() (*string, bool)`

GetPublicSuffixOk returns a tuple with the PublicSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicSuffix

`func (o *Domain) SetPublicSuffix(v string)`

SetPublicSuffix sets PublicSuffix field to given value.


### GetApexIpAddresses

`func (o *Domain) GetApexIpAddresses() []string`

GetApexIpAddresses returns the ApexIpAddresses field if non-nil, zero value otherwise.

### GetApexIpAddressesOk

`func (o *Domain) GetApexIpAddressesOk() (*[]string, bool)`

GetApexIpAddressesOk returns a tuple with the ApexIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApexIpAddresses

`func (o *Domain) SetApexIpAddresses(v []string)`

SetApexIpAddresses sets ApexIpAddresses field to given value.


### GetSiteId

`func (o *Domain) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Domain) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Domain) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.


### GetBrandId

`func (o *Domain) GetBrandId() int64`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *Domain) GetBrandIdOk() (*int64, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *Domain) SetBrandId(v int64)`

SetBrandId sets BrandId field to given value.


### GetDeletable

`func (o *Domain) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *Domain) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *Domain) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.


### GetDomainCdnConfig

`func (o *Domain) GetDomainCdnConfig() DomainCdnConfig`

GetDomainCdnConfig returns the DomainCdnConfig field if non-nil, zero value otherwise.

### GetDomainCdnConfigOk

`func (o *Domain) GetDomainCdnConfigOk() (*DomainCdnConfig, bool)`

GetDomainCdnConfigOk returns a tuple with the DomainCdnConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainCdnConfig

`func (o *Domain) SetDomainCdnConfig(v DomainCdnConfig)`

SetDomainCdnConfig sets DomainCdnConfig field to given value.


### GetSetupInfo

`func (o *Domain) GetSetupInfo() DomainSetupInfo`

GetSetupInfo returns the SetupInfo field if non-nil, zero value otherwise.

### GetSetupInfoOk

`func (o *Domain) GetSetupInfoOk() (*DomainSetupInfo, bool)`

GetSetupInfoOk returns a tuple with the SetupInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupInfo

`func (o *Domain) SetSetupInfo(v DomainSetupInfo)`

SetSetupInfo sets SetupInfo field to given value.


### GetDerivedBrandName

`func (o *Domain) GetDerivedBrandName() string`

GetDerivedBrandName returns the DerivedBrandName field if non-nil, zero value otherwise.

### GetDerivedBrandNameOk

`func (o *Domain) GetDerivedBrandNameOk() (*string, bool)`

GetDerivedBrandNameOk returns a tuple with the DerivedBrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedBrandName

`func (o *Domain) SetDerivedBrandName(v string)`

SetDerivedBrandName sets DerivedBrandName field to given value.


### GetCreatedById

`func (o *Domain) GetCreatedById() int32`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *Domain) GetCreatedByIdOk() (*int32, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *Domain) SetCreatedById(v int32)`

SetCreatedById sets CreatedById field to given value.


### GetUpdatedById

`func (o *Domain) GetUpdatedById() int32`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *Domain) GetUpdatedByIdOk() (*int32, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *Domain) SetUpdatedById(v int32)`

SetUpdatedById sets UpdatedById field to given value.


### GetLabel

`func (o *Domain) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Domain) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Domain) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetIsAnyPrimary

`func (o *Domain) GetIsAnyPrimary() bool`

GetIsAnyPrimary returns the IsAnyPrimary field if non-nil, zero value otherwise.

### GetIsAnyPrimaryOk

`func (o *Domain) GetIsAnyPrimaryOk() (*bool, bool)`

GetIsAnyPrimaryOk returns a tuple with the IsAnyPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnyPrimary

`func (o *Domain) SetIsAnyPrimary(v bool)`

SetIsAnyPrimary sets IsAnyPrimary field to given value.


### GetIsLegacyDomain

`func (o *Domain) GetIsLegacyDomain() bool`

GetIsLegacyDomain returns the IsLegacyDomain field if non-nil, zero value otherwise.

### GetIsLegacyDomainOk

`func (o *Domain) GetIsLegacyDomainOk() (*bool, bool)`

GetIsLegacyDomainOk returns a tuple with the IsLegacyDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLegacyDomain

`func (o *Domain) SetIsLegacyDomain(v bool)`

SetIsLegacyDomain sets IsLegacyDomain field to given value.


### GetIsInternalDomain

`func (o *Domain) GetIsInternalDomain() bool`

GetIsInternalDomain returns the IsInternalDomain field if non-nil, zero value otherwise.

### GetIsInternalDomainOk

`func (o *Domain) GetIsInternalDomainOk() (*bool, bool)`

GetIsInternalDomainOk returns a tuple with the IsInternalDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternalDomain

`func (o *Domain) SetIsInternalDomain(v bool)`

SetIsInternalDomain sets IsInternalDomain field to given value.


### GetIsResolvingInternalProperty

`func (o *Domain) GetIsResolvingInternalProperty() bool`

GetIsResolvingInternalProperty returns the IsResolvingInternalProperty field if non-nil, zero value otherwise.

### GetIsResolvingInternalPropertyOk

`func (o *Domain) GetIsResolvingInternalPropertyOk() (*bool, bool)`

GetIsResolvingInternalPropertyOk returns a tuple with the IsResolvingInternalProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsResolvingInternalProperty

`func (o *Domain) SetIsResolvingInternalProperty(v bool)`

SetIsResolvingInternalProperty sets IsResolvingInternalProperty field to given value.


### GetIsResolvingIgnoringManuallyMarkedAsResolving

`func (o *Domain) GetIsResolvingIgnoringManuallyMarkedAsResolving() bool`

GetIsResolvingIgnoringManuallyMarkedAsResolving returns the IsResolvingIgnoringManuallyMarkedAsResolving field if non-nil, zero value otherwise.

### GetIsResolvingIgnoringManuallyMarkedAsResolvingOk

`func (o *Domain) GetIsResolvingIgnoringManuallyMarkedAsResolvingOk() (*bool, bool)`

GetIsResolvingIgnoringManuallyMarkedAsResolvingOk returns a tuple with the IsResolvingIgnoringManuallyMarkedAsResolving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsResolvingIgnoringManuallyMarkedAsResolving

`func (o *Domain) SetIsResolvingIgnoringManuallyMarkedAsResolving(v bool)`

SetIsResolvingIgnoringManuallyMarkedAsResolving sets IsResolvingIgnoringManuallyMarkedAsResolving field to given value.


### GetIsUsedForAnyContentType

`func (o *Domain) GetIsUsedForAnyContentType() bool`

GetIsUsedForAnyContentType returns the IsUsedForAnyContentType field if non-nil, zero value otherwise.

### GetIsUsedForAnyContentTypeOk

`func (o *Domain) GetIsUsedForAnyContentTypeOk() (*bool, bool)`

GetIsUsedForAnyContentTypeOk returns a tuple with the IsUsedForAnyContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsedForAnyContentType

`func (o *Domain) SetIsUsedForAnyContentType(v bool)`

SetIsUsedForAnyContentType sets IsUsedForAnyContentType field to given value.


### GetIsLegacy

`func (o *Domain) GetIsLegacy() bool`

GetIsLegacy returns the IsLegacy field if non-nil, zero value otherwise.

### GetIsLegacyOk

`func (o *Domain) GetIsLegacyOk() (*bool, bool)`

GetIsLegacyOk returns a tuple with the IsLegacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLegacy

`func (o *Domain) SetIsLegacy(v bool)`

SetIsLegacy sets IsLegacy field to given value.


### GetAuthorAt

`func (o *Domain) GetAuthorAt() int64`

GetAuthorAt returns the AuthorAt field if non-nil, zero value otherwise.

### GetAuthorAtOk

`func (o *Domain) GetAuthorAtOk() (*int64, bool)`

GetAuthorAtOk returns a tuple with the AuthorAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorAt

`func (o *Domain) SetAuthorAt(v int64)`

SetAuthorAt sets AuthorAt field to given value.


### GetCosObjectType

`func (o *Domain) GetCosObjectType() string`

GetCosObjectType returns the CosObjectType field if non-nil, zero value otherwise.

### GetCosObjectTypeOk

`func (o *Domain) GetCosObjectTypeOk() (*string, bool)`

GetCosObjectTypeOk returns a tuple with the CosObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosObjectType

`func (o *Domain) SetCosObjectType(v string)`

SetCosObjectType sets CosObjectType field to given value.


### GetCdnPurgeEmbargoTime

`func (o *Domain) GetCdnPurgeEmbargoTime() int64`

GetCdnPurgeEmbargoTime returns the CdnPurgeEmbargoTime field if non-nil, zero value otherwise.

### GetCdnPurgeEmbargoTimeOk

`func (o *Domain) GetCdnPurgeEmbargoTimeOk() (*int64, bool)`

GetCdnPurgeEmbargoTimeOk returns a tuple with the CdnPurgeEmbargoTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnPurgeEmbargoTime

`func (o *Domain) SetCdnPurgeEmbargoTime(v int64)`

SetCdnPurgeEmbargoTime sets CdnPurgeEmbargoTime field to given value.


### GetIsStagingDomain

`func (o *Domain) GetIsStagingDomain() bool`

GetIsStagingDomain returns the IsStagingDomain field if non-nil, zero value otherwise.

### GetIsStagingDomainOk

`func (o *Domain) GetIsStagingDomainOk() (*bool, bool)`

GetIsStagingDomainOk returns a tuple with the IsStagingDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStagingDomain

`func (o *Domain) SetIsStagingDomain(v bool)`

SetIsStagingDomain sets IsStagingDomain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


