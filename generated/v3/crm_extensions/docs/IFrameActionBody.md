# IFrameActionBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [default to "IFRAME"]
**Width** | **int32** |  | 
**Height** | **int32** |  | 
**Url** | **string** |  | 
**Label** | Pointer to **string** |  | [optional] 
**PropertyNamesIncluded** | **[]string** |  | 

## Methods

### NewIFrameActionBody

`func NewIFrameActionBody(type_ string, width int32, height int32, url string, propertyNamesIncluded []string, ) *IFrameActionBody`

NewIFrameActionBody instantiates a new IFrameActionBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIFrameActionBodyWithDefaults

`func NewIFrameActionBodyWithDefaults() *IFrameActionBody`

NewIFrameActionBodyWithDefaults instantiates a new IFrameActionBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IFrameActionBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IFrameActionBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IFrameActionBody) SetType(v string)`

SetType sets Type field to given value.


### GetWidth

`func (o *IFrameActionBody) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *IFrameActionBody) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *IFrameActionBody) SetWidth(v int32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *IFrameActionBody) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *IFrameActionBody) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *IFrameActionBody) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetUrl

`func (o *IFrameActionBody) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IFrameActionBody) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IFrameActionBody) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetLabel

`func (o *IFrameActionBody) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *IFrameActionBody) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *IFrameActionBody) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *IFrameActionBody) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPropertyNamesIncluded

`func (o *IFrameActionBody) GetPropertyNamesIncluded() []string`

GetPropertyNamesIncluded returns the PropertyNamesIncluded field if non-nil, zero value otherwise.

### GetPropertyNamesIncludedOk

`func (o *IFrameActionBody) GetPropertyNamesIncludedOk() (*[]string, bool)`

GetPropertyNamesIncludedOk returns a tuple with the PropertyNamesIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyNamesIncluded

`func (o *IFrameActionBody) SetPropertyNamesIncluded(v []string)`

SetPropertyNamesIncluded sets PropertyNamesIncluded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


