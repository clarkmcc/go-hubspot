# RefreshTokenInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**User** | Pointer to **string** |  | [optional] 
**HubDomain** | Pointer to **string** |  | [optional] 
**Scopes** | **[]string** |  | 
**HubId** | **int32** |  | 
**ClientId** | **string** |  | 
**UserId** | **int32** |  | 
**TokenType** | **string** |  | 

## Methods

### NewRefreshTokenInfoResponse

`func NewRefreshTokenInfoResponse(token string, scopes []string, hubId int32, clientId string, userId int32, tokenType string, ) *RefreshTokenInfoResponse`

NewRefreshTokenInfoResponse instantiates a new RefreshTokenInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshTokenInfoResponseWithDefaults

`func NewRefreshTokenInfoResponseWithDefaults() *RefreshTokenInfoResponse`

NewRefreshTokenInfoResponseWithDefaults instantiates a new RefreshTokenInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *RefreshTokenInfoResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RefreshTokenInfoResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RefreshTokenInfoResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetUser

`func (o *RefreshTokenInfoResponse) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RefreshTokenInfoResponse) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RefreshTokenInfoResponse) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *RefreshTokenInfoResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetHubDomain

`func (o *RefreshTokenInfoResponse) GetHubDomain() string`

GetHubDomain returns the HubDomain field if non-nil, zero value otherwise.

### GetHubDomainOk

`func (o *RefreshTokenInfoResponse) GetHubDomainOk() (*string, bool)`

GetHubDomainOk returns a tuple with the HubDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubDomain

`func (o *RefreshTokenInfoResponse) SetHubDomain(v string)`

SetHubDomain sets HubDomain field to given value.

### HasHubDomain

`func (o *RefreshTokenInfoResponse) HasHubDomain() bool`

HasHubDomain returns a boolean if a field has been set.

### GetScopes

`func (o *RefreshTokenInfoResponse) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *RefreshTokenInfoResponse) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *RefreshTokenInfoResponse) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetHubId

`func (o *RefreshTokenInfoResponse) GetHubId() int32`

GetHubId returns the HubId field if non-nil, zero value otherwise.

### GetHubIdOk

`func (o *RefreshTokenInfoResponse) GetHubIdOk() (*int32, bool)`

GetHubIdOk returns a tuple with the HubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubId

`func (o *RefreshTokenInfoResponse) SetHubId(v int32)`

SetHubId sets HubId field to given value.


### GetClientId

`func (o *RefreshTokenInfoResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *RefreshTokenInfoResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *RefreshTokenInfoResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetUserId

`func (o *RefreshTokenInfoResponse) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RefreshTokenInfoResponse) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RefreshTokenInfoResponse) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetTokenType

`func (o *RefreshTokenInfoResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *RefreshTokenInfoResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *RefreshTokenInfoResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


