# UrlMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of this URL redirect. | 
**RoutePrefix** | **string** | The target incoming URL, path, or pattern to match for redirection. | 
**Destination** | **string** | The destination URL, where the target URL should be redirected if it matches the &#x60;routePrefix&#x60;. | 
**RedirectStyle** | **int32** | The type of redirect to create. Options include: 301 (permanent), 302 (temporary), or 305 (proxy). Find more details [here](https://knowledge.hubspot.com/cos-general/how-to-redirect-a-hubspot-page). | 
**IsOnlyAfterNotFound** | **bool** | Whether the URL redirect mapping should apply only if a live page on the URL isn&#39;t found. If False, the URL redirect mapping will take precedence over any existing page. | 
**IsMatchFullUrl** | **bool** | Whether the &#x60;routePrefix&#x60; should match on the entire URL, including the domain. | 
**IsMatchQueryString** | **bool** | Whether the &#x60;routePrefix&#x60; should match on the entire URL path, including the query string. | 
**IsPattern** | **bool** | Whether the &#x60;routePrefix&#x60; should match based on pattern. | 
**IsTrailingSlashOptional** | **bool** | Whether a trailing slash will be ignored. | 
**IsProtocolAgnostic** | **bool** | Whether the &#x60;routePrefix&#x60; should match both HTTP and HTTPS protocols. | 
**Precedence** | **int32** | Used to prioritize URL redirection. If a given URL matches more than one redirect, the one with the **lower** precedence will be used. | 
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUrlMapping

`func NewUrlMapping(id string, routePrefix string, destination string, redirectStyle int32, isOnlyAfterNotFound bool, isMatchFullUrl bool, isMatchQueryString bool, isPattern bool, isTrailingSlashOptional bool, isProtocolAgnostic bool, precedence int32, ) *UrlMapping`

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

`func (o *UrlMapping) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UrlMapping) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UrlMapping) SetId(v string)`

SetId sets Id field to given value.


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


### GetCreated

`func (o *UrlMapping) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UrlMapping) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UrlMapping) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *UrlMapping) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *UrlMapping) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *UrlMapping) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *UrlMapping) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *UrlMapping) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


