# UrlMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The unique ID of this URL redirect. | 
**PortalId** | **int32** |  | 
**Created** | **int64** |  | 
**Updated** | **int64** |  | 
**CreatedById** | **int32** |  | 
**UpdatedById** | **int32** |  | 
**RoutePrefix** | **string** | The target incoming URL, path, or pattern to match for redirection. | 
**Destination** | **string** | The destination URL, where the target URL should be redirected if it matches the &#x60;routePrefix&#x60;. | 
**RedirectStyle** | **int32** | The type of redirect to create. Options include: 301 (permanent), 302 (temporary), or 305 (proxy). Find more details [here](https://knowledge.hubspot.com/cos-general/how-to-redirect-a-hubspot-page). | 
**ContentGroupId** | **int64** |  | 
**IsOnlyAfterNotFound** | **bool** | Whether the URL redirect mapping should apply only if a live page on the URL isn&#39;t found. If False, the URL redirect mapping will take precedence over any existing page. | 
**IsRegex** | **bool** |  | 
**IsMatchFullUrl** | **bool** | Whether the &#x60;routePrefix&#x60; should match on the entire URL, including the domain. | 
**IsMatchQueryString** | **bool** | Whether the &#x60;routePrefix&#x60; should match on the entire URL path, including the query string. | 
**IsPattern** | **bool** | Whether the &#x60;routePrefix&#x60; should match based on pattern. | 
**IsTrailingSlashOptional** | **bool** | Whether a trailing slash will be ignored. | 
**IsProtocolAgnostic** | **bool** | Whether the &#x60;routePrefix&#x60; should match both HTTP and HTTPS protocols. | 
**Name** | **string** |  | 
**Precedence** | **int32** | Used to prioritize URL redirection. If a given URL matches more than one redirect, the one with the **lower** precedence will be used. | 
**DeletedAt** | **int64** |  | 
**Note** | **string** |  | 
**Label** | **string** |  | 
**InternallyCreated** | **bool** |  | 
**CosObjectType** | **string** |  | 
**CdnPurgeEmbargoTime** | **int64** |  | 

## Methods

### NewUrlMapping

`func NewUrlMapping(id int64, portalId int32, created int64, updated int64, createdById int32, updatedById int32, routePrefix string, destination string, redirectStyle int32, contentGroupId int64, isOnlyAfterNotFound bool, isRegex bool, isMatchFullUrl bool, isMatchQueryString bool, isPattern bool, isTrailingSlashOptional bool, isProtocolAgnostic bool, name string, precedence int32, deletedAt int64, note string, label string, internallyCreated bool, cosObjectType string, cdnPurgeEmbargoTime int64, ) *UrlMapping`

NewUrlMapping instantiates a new UrlMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUrlMappingWithDefaults

`func NewUrlMappingWithDefaults() *UrlMapping`

NewUrlMappingWithDefaults instantiates a new UrlMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UrlMapping) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UrlMapping) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UrlMapping) SetId(v int64)`

SetId sets Id field to given value.


### GetPortalId

`func (o *UrlMapping) GetPortalId() int32`

GetPortalId returns the PortalId field if non-nil, zero value otherwise.

### GetPortalIdOk

`func (o *UrlMapping) GetPortalIdOk() (*int32, bool)`

GetPortalIdOk returns a tuple with the PortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalId

`func (o *UrlMapping) SetPortalId(v int32)`

SetPortalId sets PortalId field to given value.


### GetCreated

`func (o *UrlMapping) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UrlMapping) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UrlMapping) SetCreated(v int64)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *UrlMapping) GetUpdated() int64`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *UrlMapping) GetUpdatedOk() (*int64, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *UrlMapping) SetUpdated(v int64)`

SetUpdated sets Updated field to given value.


### GetCreatedById

`func (o *UrlMapping) GetCreatedById() int32`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *UrlMapping) GetCreatedByIdOk() (*int32, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *UrlMapping) SetCreatedById(v int32)`

SetCreatedById sets CreatedById field to given value.


### GetUpdatedById

`func (o *UrlMapping) GetUpdatedById() int32`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *UrlMapping) GetUpdatedByIdOk() (*int32, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *UrlMapping) SetUpdatedById(v int32)`

SetUpdatedById sets UpdatedById field to given value.


### GetRoutePrefix

`func (o *UrlMapping) GetRoutePrefix() string`

GetRoutePrefix returns the RoutePrefix field if non-nil, zero value otherwise.

### GetRoutePrefixOk

`func (o *UrlMapping) GetRoutePrefixOk() (*string, bool)`

GetRoutePrefixOk returns a tuple with the RoutePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutePrefix

`func (o *UrlMapping) SetRoutePrefix(v string)`

SetRoutePrefix sets RoutePrefix field to given value.


### GetDestination

`func (o *UrlMapping) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *UrlMapping) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *UrlMapping) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetRedirectStyle

`func (o *UrlMapping) GetRedirectStyle() int32`

GetRedirectStyle returns the RedirectStyle field if non-nil, zero value otherwise.

### GetRedirectStyleOk

`func (o *UrlMapping) GetRedirectStyleOk() (*int32, bool)`

GetRedirectStyleOk returns a tuple with the RedirectStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectStyle

`func (o *UrlMapping) SetRedirectStyle(v int32)`

SetRedirectStyle sets RedirectStyle field to given value.


### GetContentGroupId

`func (o *UrlMapping) GetContentGroupId() int64`

GetContentGroupId returns the ContentGroupId field if non-nil, zero value otherwise.

### GetContentGroupIdOk

`func (o *UrlMapping) GetContentGroupIdOk() (*int64, bool)`

GetContentGroupIdOk returns a tuple with the ContentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentGroupId

`func (o *UrlMapping) SetContentGroupId(v int64)`

SetContentGroupId sets ContentGroupId field to given value.


### GetIsOnlyAfterNotFound

`func (o *UrlMapping) GetIsOnlyAfterNotFound() bool`

GetIsOnlyAfterNotFound returns the IsOnlyAfterNotFound field if non-nil, zero value otherwise.

### GetIsOnlyAfterNotFoundOk

`func (o *UrlMapping) GetIsOnlyAfterNotFoundOk() (*bool, bool)`

GetIsOnlyAfterNotFoundOk returns a tuple with the IsOnlyAfterNotFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnlyAfterNotFound

`func (o *UrlMapping) SetIsOnlyAfterNotFound(v bool)`

SetIsOnlyAfterNotFound sets IsOnlyAfterNotFound field to given value.


### GetIsRegex

`func (o *UrlMapping) GetIsRegex() bool`

GetIsRegex returns the IsRegex field if non-nil, zero value otherwise.

### GetIsRegexOk

`func (o *UrlMapping) GetIsRegexOk() (*bool, bool)`

GetIsRegexOk returns a tuple with the IsRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRegex

`func (o *UrlMapping) SetIsRegex(v bool)`

SetIsRegex sets IsRegex field to given value.


### GetIsMatchFullUrl

`func (o *UrlMapping) GetIsMatchFullUrl() bool`

GetIsMatchFullUrl returns the IsMatchFullUrl field if non-nil, zero value otherwise.

### GetIsMatchFullUrlOk

`func (o *UrlMapping) GetIsMatchFullUrlOk() (*bool, bool)`

GetIsMatchFullUrlOk returns a tuple with the IsMatchFullUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMatchFullUrl

`func (o *UrlMapping) SetIsMatchFullUrl(v bool)`

SetIsMatchFullUrl sets IsMatchFullUrl field to given value.


### GetIsMatchQueryString

`func (o *UrlMapping) GetIsMatchQueryString() bool`

GetIsMatchQueryString returns the IsMatchQueryString field if non-nil, zero value otherwise.

### GetIsMatchQueryStringOk

`func (o *UrlMapping) GetIsMatchQueryStringOk() (*bool, bool)`

GetIsMatchQueryStringOk returns a tuple with the IsMatchQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMatchQueryString

`func (o *UrlMapping) SetIsMatchQueryString(v bool)`

SetIsMatchQueryString sets IsMatchQueryString field to given value.


### GetIsPattern

`func (o *UrlMapping) GetIsPattern() bool`

GetIsPattern returns the IsPattern field if non-nil, zero value otherwise.

### GetIsPatternOk

`func (o *UrlMapping) GetIsPatternOk() (*bool, bool)`

GetIsPatternOk returns a tuple with the IsPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPattern

`func (o *UrlMapping) SetIsPattern(v bool)`

SetIsPattern sets IsPattern field to given value.


### GetIsTrailingSlashOptional

`func (o *UrlMapping) GetIsTrailingSlashOptional() bool`

GetIsTrailingSlashOptional returns the IsTrailingSlashOptional field if non-nil, zero value otherwise.

### GetIsTrailingSlashOptionalOk

`func (o *UrlMapping) GetIsTrailingSlashOptionalOk() (*bool, bool)`

GetIsTrailingSlashOptionalOk returns a tuple with the IsTrailingSlashOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrailingSlashOptional

`func (o *UrlMapping) SetIsTrailingSlashOptional(v bool)`

SetIsTrailingSlashOptional sets IsTrailingSlashOptional field to given value.


### GetIsProtocolAgnostic

`func (o *UrlMapping) GetIsProtocolAgnostic() bool`

GetIsProtocolAgnostic returns the IsProtocolAgnostic field if non-nil, zero value otherwise.

### GetIsProtocolAgnosticOk

`func (o *UrlMapping) GetIsProtocolAgnosticOk() (*bool, bool)`

GetIsProtocolAgnosticOk returns a tuple with the IsProtocolAgnostic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtocolAgnostic

`func (o *UrlMapping) SetIsProtocolAgnostic(v bool)`

SetIsProtocolAgnostic sets IsProtocolAgnostic field to given value.


### GetName

`func (o *UrlMapping) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UrlMapping) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UrlMapping) SetName(v string)`

SetName sets Name field to given value.


### GetPrecedence

`func (o *UrlMapping) GetPrecedence() int32`

GetPrecedence returns the Precedence field if non-nil, zero value otherwise.

### GetPrecedenceOk

`func (o *UrlMapping) GetPrecedenceOk() (*int32, bool)`

GetPrecedenceOk returns a tuple with the Precedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedence

`func (o *UrlMapping) SetPrecedence(v int32)`

SetPrecedence sets Precedence field to given value.


### GetDeletedAt

`func (o *UrlMapping) GetDeletedAt() int64`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *UrlMapping) GetDeletedAtOk() (*int64, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *UrlMapping) SetDeletedAt(v int64)`

SetDeletedAt sets DeletedAt field to given value.


### GetNote

`func (o *UrlMapping) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *UrlMapping) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *UrlMapping) SetNote(v string)`

SetNote sets Note field to given value.


### GetLabel

`func (o *UrlMapping) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UrlMapping) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UrlMapping) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetInternallyCreated

`func (o *UrlMapping) GetInternallyCreated() bool`

GetInternallyCreated returns the InternallyCreated field if non-nil, zero value otherwise.

### GetInternallyCreatedOk

`func (o *UrlMapping) GetInternallyCreatedOk() (*bool, bool)`

GetInternallyCreatedOk returns a tuple with the InternallyCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternallyCreated

`func (o *UrlMapping) SetInternallyCreated(v bool)`

SetInternallyCreated sets InternallyCreated field to given value.


### GetCosObjectType

`func (o *UrlMapping) GetCosObjectType() string`

GetCosObjectType returns the CosObjectType field if non-nil, zero value otherwise.

### GetCosObjectTypeOk

`func (o *UrlMapping) GetCosObjectTypeOk() (*string, bool)`

GetCosObjectTypeOk returns a tuple with the CosObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosObjectType

`func (o *UrlMapping) SetCosObjectType(v string)`

SetCosObjectType sets CosObjectType field to given value.


### GetCdnPurgeEmbargoTime

`func (o *UrlMapping) GetCdnPurgeEmbargoTime() int64`

GetCdnPurgeEmbargoTime returns the CdnPurgeEmbargoTime field if non-nil, zero value otherwise.

### GetCdnPurgeEmbargoTimeOk

`func (o *UrlMapping) GetCdnPurgeEmbargoTimeOk() (*int64, bool)`

GetCdnPurgeEmbargoTimeOk returns a tuple with the CdnPurgeEmbargoTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnPurgeEmbargoTime

`func (o *UrlMapping) SetCdnPurgeEmbargoTime(v int64)`

SetCdnPurgeEmbargoTime sets CdnPurgeEmbargoTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


