# TokenResponseIF

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefreshToken** | **string** |  | 
**ExpiresIn** | **int32** |  | 
**AccessToken** | **string** |  | 
**IdToken** | Pointer to **string** |  | [optional] 
**TokenType** | **string** |  | 

## Methods

### NewTokenResponseIF

`func NewTokenResponseIF(refreshToken string, expiresIn int32, accessToken string, tokenType string, ) *TokenResponseIF`

NewTokenResponseIF instantiates a new TokenResponseIF object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenResponseIFWithDefaults

`func NewTokenResponseIFWithDefaults() *TokenResponseIF`

NewTokenResponseIFWithDefaults instantiates a new TokenResponseIF object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefreshToken

`func (o *TokenResponseIF) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *TokenResponseIF) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *TokenResponseIF) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.


### GetExpiresIn

`func (o *TokenResponseIF) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *TokenResponseIF) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *TokenResponseIF) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.


### GetAccessToken

`func (o *TokenResponseIF) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TokenResponseIF) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TokenResponseIF) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetIdToken

`func (o *TokenResponseIF) GetIdToken() string`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *TokenResponseIF) GetIdTokenOk() (*string, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *TokenResponseIF) SetIdToken(v string)`

SetIdToken sets IdToken field to given value.

### HasIdToken

`func (o *TokenResponseIF) HasIdToken() bool`

HasIdToken returns a boolean if a field has been set.

### GetTokenType

`func (o *TokenResponseIF) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *TokenResponseIF) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *TokenResponseIF) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


