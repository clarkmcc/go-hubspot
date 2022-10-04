# AccessTokenInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**User** | Pointer to **string** |  | [optional] 
**HubDomain** | Pointer to **string** |  | [optional] 
**Scopes** | **[]string** |  | 
**ScopeToScopeGroupPks** | **[]int32** |  | 
**TrialScopes** | **[]string** |  | 
**TrialScopeToScopeGroupPks** | **[]int32** |  | 
**HubId** | **int32** |  | 
**AppId** | **int32** |  | 
**ExpiresIn** | **int32** |  | 
**UserId** | **int32** |  | 
**TokenType** | **string** |  | 

## Methods

### NewAccessTokenInfoResponse

`func NewAccessTokenInfoResponse(token string, scopes []string, scopeToScopeGroupPks []int32, trialScopes []string, trialScopeToScopeGroupPks []int32, hubId int32, appId int32, expiresIn int32, userId int32, tokenType string, ) *AccessTokenInfoResponse`

NewAccessTokenInfoResponse instantiates a new AccessTokenInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenInfoResponseWithDefaults

`func NewAccessTokenInfoResponseWithDefaults() *AccessTokenInfoResponse`

NewAccessTokenInfoResponseWithDefaults instantiates a new AccessTokenInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *AccessTokenInfoResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AccessTokenInfoResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AccessTokenInfoResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetUser

`func (o *AccessTokenInfoResponse) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AccessTokenInfoResponse) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AccessTokenInfoResponse) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *AccessTokenInfoResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetHubDomain

`func (o *AccessTokenInfoResponse) GetHubDomain() string`

GetHubDomain returns the HubDomain field if non-nil, zero value otherwise.

### GetHubDomainOk

`func (o *AccessTokenInfoResponse) GetHubDomainOk() (*string, bool)`

GetHubDomainOk returns a tuple with the HubDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubDomain

`func (o *AccessTokenInfoResponse) SetHubDomain(v string)`

SetHubDomain sets HubDomain field to given value.

### HasHubDomain

`func (o *AccessTokenInfoResponse) HasHubDomain() bool`

HasHubDomain returns a boolean if a field has been set.

### GetScopes

`func (o *AccessTokenInfoResponse) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AccessTokenInfoResponse) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AccessTokenInfoResponse) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetScopeToScopeGroupPks

`func (o *AccessTokenInfoResponse) GetScopeToScopeGroupPks() []int32`

GetScopeToScopeGroupPks returns the ScopeToScopeGroupPks field if non-nil, zero value otherwise.

### GetScopeToScopeGroupPksOk

`func (o *AccessTokenInfoResponse) GetScopeToScopeGroupPksOk() (*[]int32, bool)`

GetScopeToScopeGroupPksOk returns a tuple with the ScopeToScopeGroupPks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeToScopeGroupPks

`func (o *AccessTokenInfoResponse) SetScopeToScopeGroupPks(v []int32)`

SetScopeToScopeGroupPks sets ScopeToScopeGroupPks field to given value.


### GetTrialScopes

`func (o *AccessTokenInfoResponse) GetTrialScopes() []string`

GetTrialScopes returns the TrialScopes field if non-nil, zero value otherwise.

### GetTrialScopesOk

`func (o *AccessTokenInfoResponse) GetTrialScopesOk() (*[]string, bool)`

GetTrialScopesOk returns a tuple with the TrialScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialScopes

`func (o *AccessTokenInfoResponse) SetTrialScopes(v []string)`

SetTrialScopes sets TrialScopes field to given value.


### GetTrialScopeToScopeGroupPks

`func (o *AccessTokenInfoResponse) GetTrialScopeToScopeGroupPks() []int32`

GetTrialScopeToScopeGroupPks returns the TrialScopeToScopeGroupPks field if non-nil, zero value otherwise.

### GetTrialScopeToScopeGroupPksOk

`func (o *AccessTokenInfoResponse) GetTrialScopeToScopeGroupPksOk() (*[]int32, bool)`

GetTrialScopeToScopeGroupPksOk returns a tuple with the TrialScopeToScopeGroupPks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialScopeToScopeGroupPks

`func (o *AccessTokenInfoResponse) SetTrialScopeToScopeGroupPks(v []int32)`

SetTrialScopeToScopeGroupPks sets TrialScopeToScopeGroupPks field to given value.


### GetHubId

`func (o *AccessTokenInfoResponse) GetHubId() int32`

GetHubId returns the HubId field if non-nil, zero value otherwise.

### GetHubIdOk

`func (o *AccessTokenInfoResponse) GetHubIdOk() (*int32, bool)`

GetHubIdOk returns a tuple with the HubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubId

`func (o *AccessTokenInfoResponse) SetHubId(v int32)`

SetHubId sets HubId field to given value.


### GetAppId

`func (o *AccessTokenInfoResponse) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AccessTokenInfoResponse) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AccessTokenInfoResponse) SetAppId(v int32)`

SetAppId sets AppId field to given value.


### GetExpiresIn

`func (o *AccessTokenInfoResponse) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *AccessTokenInfoResponse) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *AccessTokenInfoResponse) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.


### GetUserId

`func (o *AccessTokenInfoResponse) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AccessTokenInfoResponse) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AccessTokenInfoResponse) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetTokenType

`func (o *AccessTokenInfoResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *AccessTokenInfoResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *AccessTokenInfoResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


