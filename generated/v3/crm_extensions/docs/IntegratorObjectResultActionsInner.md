# IntegratorObjectResultActionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [default to "IFRAME"]
**Confirmation** | Pointer to [**ActionConfirmationBody**](ActionConfirmationBody.md) |  | [optional] 
**HttpMethod** | **string** |  | 
**Url** | **string** |  | 
**Label** | Pointer to **string** |  | [optional] 
**PropertyNamesIncluded** | **[]string** |  | 
**Width** | **int32** |  | 
**Height** | **int32** |  | 

## Methods

### NewIntegratorObjectResultActionsInner

`func NewIntegratorObjectResultActionsInner(type_ string, httpMethod string, url string, propertyNamesIncluded []string, width int32, height int32, ) *IntegratorObjectResultActionsInner`

NewIntegratorObjectResultActionsInner instantiates a new IntegratorObjectResultActionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegratorObjectResultActionsInnerWithDefaults

`func NewIntegratorObjectResultActionsInnerWithDefaults() *IntegratorObjectResultActionsInner`

NewIntegratorObjectResultActionsInnerWithDefaults instantiates a new IntegratorObjectResultActionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IntegratorObjectResultActionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegratorObjectResultActionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegratorObjectResultActionsInner) SetType(v string)`

SetType sets Type field to given value.


### GetConfirmation

`func (o *IntegratorObjectResultActionsInner) GetConfirmation() ActionConfirmationBody`

GetConfirmation returns the Confirmation field if non-nil, zero value otherwise.

### GetConfirmationOk

`func (o *IntegratorObjectResultActionsInner) GetConfirmationOk() (*ActionConfirmationBody, bool)`

GetConfirmationOk returns a tuple with the Confirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmation

`func (o *IntegratorObjectResultActionsInner) SetConfirmation(v ActionConfirmationBody)`

SetConfirmation sets Confirmation field to given value.

### HasConfirmation

`func (o *IntegratorObjectResultActionsInner) HasConfirmation() bool`

HasConfirmation returns a boolean if a field has been set.

### GetHttpMethod

`func (o *IntegratorObjectResultActionsInner) GetHttpMethod() string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *IntegratorObjectResultActionsInner) GetHttpMethodOk() (*string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *IntegratorObjectResultActionsInner) SetHttpMethod(v string)`

SetHttpMethod sets HttpMethod field to given value.


### GetUrl

`func (o *IntegratorObjectResultActionsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IntegratorObjectResultActionsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IntegratorObjectResultActionsInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetLabel

`func (o *IntegratorObjectResultActionsInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *IntegratorObjectResultActionsInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *IntegratorObjectResultActionsInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *IntegratorObjectResultActionsInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPropertyNamesIncluded

`func (o *IntegratorObjectResultActionsInner) GetPropertyNamesIncluded() []string`

GetPropertyNamesIncluded returns the PropertyNamesIncluded field if non-nil, zero value otherwise.

### GetPropertyNamesIncludedOk

`func (o *IntegratorObjectResultActionsInner) GetPropertyNamesIncludedOk() (*[]string, bool)`

GetPropertyNamesIncludedOk returns a tuple with the PropertyNamesIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyNamesIncluded

`func (o *IntegratorObjectResultActionsInner) SetPropertyNamesIncluded(v []string)`

SetPropertyNamesIncluded sets PropertyNamesIncluded field to given value.


### GetWidth

`func (o *IntegratorObjectResultActionsInner) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *IntegratorObjectResultActionsInner) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *IntegratorObjectResultActionsInner) SetWidth(v int32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *IntegratorObjectResultActionsInner) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *IntegratorObjectResultActionsInner) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *IntegratorObjectResultActionsInner) SetHeight(v int32)`

SetHeight sets Height field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


