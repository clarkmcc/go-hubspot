# IntegratorObjectResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Title** | **string** |  | 
**LinkUrl** | Pointer to **string** |  | [optional] 
**Tokens** | [**[]ObjectToken**](ObjectToken.md) |  | 
**Actions** | [**[]IntegratorObjectResultActionsInner**](IntegratorObjectResultActionsInner.md) |  | 

## Methods

### NewIntegratorObjectResult

`func NewIntegratorObjectResult(id string, title string, tokens []ObjectToken, actions []IntegratorObjectResultActionsInner, ) *IntegratorObjectResult`

NewIntegratorObjectResult instantiates a new IntegratorObjectResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegratorObjectResultWithDefaults

`func NewIntegratorObjectResultWithDefaults() *IntegratorObjectResult`

NewIntegratorObjectResultWithDefaults instantiates a new IntegratorObjectResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegratorObjectResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegratorObjectResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegratorObjectResult) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *IntegratorObjectResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IntegratorObjectResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IntegratorObjectResult) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetLinkUrl

`func (o *IntegratorObjectResult) GetLinkUrl() string`

GetLinkUrl returns the LinkUrl field if non-nil, zero value otherwise.

### GetLinkUrlOk

`func (o *IntegratorObjectResult) GetLinkUrlOk() (*string, bool)`

GetLinkUrlOk returns a tuple with the LinkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUrl

`func (o *IntegratorObjectResult) SetLinkUrl(v string)`

SetLinkUrl sets LinkUrl field to given value.

### HasLinkUrl

`func (o *IntegratorObjectResult) HasLinkUrl() bool`

HasLinkUrl returns a boolean if a field has been set.

### GetTokens

`func (o *IntegratorObjectResult) GetTokens() []ObjectToken`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *IntegratorObjectResult) GetTokensOk() (*[]ObjectToken, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *IntegratorObjectResult) SetTokens(v []ObjectToken)`

SetTokens sets Tokens field to given value.


### GetActions

`func (o *IntegratorObjectResult) GetActions() []IntegratorObjectResultActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *IntegratorObjectResult) GetActionsOk() (*[]IntegratorObjectResultActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *IntegratorObjectResult) SetActions(v []IntegratorObjectResultActionsInner)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


