# UrlMappingCreateRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Precedence** | Pointer to **int32** |  | [optional] 
**IsOnlyAfterNotFound** | Pointer to **bool** |  | [optional] 
**IsMatchFullUrl** | Pointer to **bool** |  | [optional] 
**IsMatchQueryString** | Pointer to **bool** |  | [optional] 
**IsPattern** | Pointer to **bool** |  | [optional] 
**IsTrailingSlashOptional** | Pointer to **bool** |  | [optional] 
**IsProtocolAgnostic** | Pointer to **bool** |  | [optional] 
**RoutePrefix** | **string** |  | 
**Destination** | **string** |  | 
**RedirectStyle** | **int32** |  | 

## Methods

### NewUrlMappingCreateRequestBody

`func NewUrlMappingCreateRequestBody(routePrefix string, destination string, redirectStyle int32, ) *UrlMappingCreateRequestBody`

NewUrlMappingCreateRequestBody instantiates a new UrlMappingCreateRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUrlMappingCreateRequestBodyWithDefaults

`func NewUrlMappingCreateRequestBodyWithDefaults() *UrlMappingCreateRequestBody`

NewUrlMappingCreateRequestBodyWithDefaults instantiates a new UrlMappingCreateRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrecedence

`func (o *UrlMappingCreateRequestBody) GetPrecedence() int32`

GetPrecedence returns the Precedence field if non-nil, zero value otherwise.

### GetPrecedenceOk

`func (o *UrlMappingCreateRequestBody) GetPrecedenceOk() (*int32, bool)`

GetPrecedenceOk returns a tuple with the Precedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedence

`func (o *UrlMappingCreateRequestBody) SetPrecedence(v int32)`

SetPrecedence sets Precedence field to given value.

### HasPrecedence

`func (o *UrlMappingCreateRequestBody) HasPrecedence() bool`

HasPrecedence returns a boolean if a field has been set.

### GetIsOnlyAfterNotFound

`func (o *UrlMappingCreateRequestBody) GetIsOnlyAfterNotFound() bool`

GetIsOnlyAfterNotFound returns the IsOnlyAfterNotFound field if non-nil, zero value otherwise.

### GetIsOnlyAfterNotFoundOk

`func (o *UrlMappingCreateRequestBody) GetIsOnlyAfterNotFoundOk() (*bool, bool)`

GetIsOnlyAfterNotFoundOk returns a tuple with the IsOnlyAfterNotFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnlyAfterNotFound

`func (o *UrlMappingCreateRequestBody) SetIsOnlyAfterNotFound(v bool)`

SetIsOnlyAfterNotFound sets IsOnlyAfterNotFound field to given value.

### HasIsOnlyAfterNotFound

`func (o *UrlMappingCreateRequestBody) HasIsOnlyAfterNotFound() bool`

HasIsOnlyAfterNotFound returns a boolean if a field has been set.

### GetIsMatchFullUrl

`func (o *UrlMappingCreateRequestBody) GetIsMatchFullUrl() bool`

GetIsMatchFullUrl returns the IsMatchFullUrl field if non-nil, zero value otherwise.

### GetIsMatchFullUrlOk

`func (o *UrlMappingCreateRequestBody) GetIsMatchFullUrlOk() (*bool, bool)`

GetIsMatchFullUrlOk returns a tuple with the IsMatchFullUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMatchFullUrl

`func (o *UrlMappingCreateRequestBody) SetIsMatchFullUrl(v bool)`

SetIsMatchFullUrl sets IsMatchFullUrl field to given value.

### HasIsMatchFullUrl

`func (o *UrlMappingCreateRequestBody) HasIsMatchFullUrl() bool`

HasIsMatchFullUrl returns a boolean if a field has been set.

### GetIsMatchQueryString

`func (o *UrlMappingCreateRequestBody) GetIsMatchQueryString() bool`

GetIsMatchQueryString returns the IsMatchQueryString field if non-nil, zero value otherwise.

### GetIsMatchQueryStringOk

`func (o *UrlMappingCreateRequestBody) GetIsMatchQueryStringOk() (*bool, bool)`

GetIsMatchQueryStringOk returns a tuple with the IsMatchQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMatchQueryString

`func (o *UrlMappingCreateRequestBody) SetIsMatchQueryString(v bool)`

SetIsMatchQueryString sets IsMatchQueryString field to given value.

### HasIsMatchQueryString

`func (o *UrlMappingCreateRequestBody) HasIsMatchQueryString() bool`

HasIsMatchQueryString returns a boolean if a field has been set.

### GetIsPattern

`func (o *UrlMappingCreateRequestBody) GetIsPattern() bool`

GetIsPattern returns the IsPattern field if non-nil, zero value otherwise.

### GetIsPatternOk

`func (o *UrlMappingCreateRequestBody) GetIsPatternOk() (*bool, bool)`

GetIsPatternOk returns a tuple with the IsPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPattern

`func (o *UrlMappingCreateRequestBody) SetIsPattern(v bool)`

SetIsPattern sets IsPattern field to given value.

### HasIsPattern

`func (o *UrlMappingCreateRequestBody) HasIsPattern() bool`

HasIsPattern returns a boolean if a field has been set.

### GetIsTrailingSlashOptional

`func (o *UrlMappingCreateRequestBody) GetIsTrailingSlashOptional() bool`

GetIsTrailingSlashOptional returns the IsTrailingSlashOptional field if non-nil, zero value otherwise.

### GetIsTrailingSlashOptionalOk

`func (o *UrlMappingCreateRequestBody) GetIsTrailingSlashOptionalOk() (*bool, bool)`

GetIsTrailingSlashOptionalOk returns a tuple with the IsTrailingSlashOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrailingSlashOptional

`func (o *UrlMappingCreateRequestBody) SetIsTrailingSlashOptional(v bool)`

SetIsTrailingSlashOptional sets IsTrailingSlashOptional field to given value.

### HasIsTrailingSlashOptional

`func (o *UrlMappingCreateRequestBody) HasIsTrailingSlashOptional() bool`

HasIsTrailingSlashOptional returns a boolean if a field has been set.

### GetIsProtocolAgnostic

`func (o *UrlMappingCreateRequestBody) GetIsProtocolAgnostic() bool`

GetIsProtocolAgnostic returns the IsProtocolAgnostic field if non-nil, zero value otherwise.

### GetIsProtocolAgnosticOk

`func (o *UrlMappingCreateRequestBody) GetIsProtocolAgnosticOk() (*bool, bool)`

GetIsProtocolAgnosticOk returns a tuple with the IsProtocolAgnostic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtocolAgnostic

`func (o *UrlMappingCreateRequestBody) SetIsProtocolAgnostic(v bool)`

SetIsProtocolAgnostic sets IsProtocolAgnostic field to given value.

### HasIsProtocolAgnostic

`func (o *UrlMappingCreateRequestBody) HasIsProtocolAgnostic() bool`

HasIsProtocolAgnostic returns a boolean if a field has been set.

### GetRoutePrefix

`func (o *UrlMappingCreateRequestBody) GetRoutePrefix() string`

GetRoutePrefix returns the RoutePrefix field if non-nil, zero value otherwise.

### GetRoutePrefixOk

`func (o *UrlMappingCreateRequestBody) GetRoutePrefixOk() (*string, bool)`

GetRoutePrefixOk returns a tuple with the RoutePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutePrefix

`func (o *UrlMappingCreateRequestBody) SetRoutePrefix(v string)`

SetRoutePrefix sets RoutePrefix field to given value.


### GetDestination

`func (o *UrlMappingCreateRequestBody) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *UrlMappingCreateRequestBody) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *UrlMappingCreateRequestBody) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetRedirectStyle

`func (o *UrlMappingCreateRequestBody) GetRedirectStyle() int32`

GetRedirectStyle returns the RedirectStyle field if non-nil, zero value otherwise.

### GetRedirectStyleOk

`func (o *UrlMappingCreateRequestBody) GetRedirectStyleOk() (*int32, bool)`

GetRedirectStyleOk returns a tuple with the RedirectStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectStyle

`func (o *UrlMappingCreateRequestBody) SetRedirectStyle(v int32)`

SetRedirectStyle sets RedirectStyle field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


