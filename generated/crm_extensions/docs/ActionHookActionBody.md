# ActionHookActionBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [default to "ACTION_HOOK"]
**Confirmation** | Pointer to [**ActionConfirmationBody**](ActionConfirmationBody.md) |  | [optional] 
**HttpMethod** | **string** |  | 
**Url** | **string** |  | 
**Label** | Pointer to **string** |  | [optional] 
**PropertyNamesIncluded** | **[]string** |  | 

## Methods

### NewActionHookActionBody

`func NewActionHookActionBody(type_ string, httpMethod string, url string, propertyNamesIncluded []string, ) *ActionHookActionBody`

NewActionHookActionBody instantiates a new ActionHookActionBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionHookActionBodyWithDefaults

`func NewActionHookActionBodyWithDefaults() *ActionHookActionBody`

NewActionHookActionBodyWithDefaults instantiates a new ActionHookActionBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ActionHookActionBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionHookActionBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionHookActionBody) SetType(v string)`

SetType sets Type field to given value.


### GetConfirmation

`func (o *ActionHookActionBody) GetConfirmation() ActionConfirmationBody`

GetConfirmation returns the Confirmation field if non-nil, zero value otherwise.

### GetConfirmationOk

`func (o *ActionHookActionBody) GetConfirmationOk() (*ActionConfirmationBody, bool)`

GetConfirmationOk returns a tuple with the Confirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmation

`func (o *ActionHookActionBody) SetConfirmation(v ActionConfirmationBody)`

SetConfirmation sets Confirmation field to given value.

### HasConfirmation

`func (o *ActionHookActionBody) HasConfirmation() bool`

HasConfirmation returns a boolean if a field has been set.

### GetHttpMethod

`func (o *ActionHookActionBody) GetHttpMethod() string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *ActionHookActionBody) GetHttpMethodOk() (*string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *ActionHookActionBody) SetHttpMethod(v string)`

SetHttpMethod sets HttpMethod field to given value.


### GetUrl

`func (o *ActionHookActionBody) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ActionHookActionBody) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ActionHookActionBody) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetLabel

`func (o *ActionHookActionBody) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ActionHookActionBody) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ActionHookActionBody) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ActionHookActionBody) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPropertyNamesIncluded

`func (o *ActionHookActionBody) GetPropertyNamesIncluded() []string`

GetPropertyNamesIncluded returns the PropertyNamesIncluded field if non-nil, zero value otherwise.

### GetPropertyNamesIncludedOk

`func (o *ActionHookActionBody) GetPropertyNamesIncludedOk() (*[]string, bool)`

GetPropertyNamesIncludedOk returns a tuple with the PropertyNamesIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyNamesIncluded

`func (o *ActionHookActionBody) SetPropertyNamesIncluded(v []string)`

SetPropertyNamesIncluded sets PropertyNamesIncluded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


