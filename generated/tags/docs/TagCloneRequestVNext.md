# TagCloneRequestVNext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the object to be cloned. | 
**Language** | Pointer to **string** | Target language of new variant. | [optional] 
**PrimaryLanguage** | Pointer to **string** | Language of primary blog tag to clone. | [optional] 
**Name** | **string** | Name of newly cloned blog tag. | 

## Methods

### NewTagCloneRequestVNext

`func NewTagCloneRequestVNext(id string, name string, ) *TagCloneRequestVNext`

NewTagCloneRequestVNext instantiates a new TagCloneRequestVNext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagCloneRequestVNextWithDefaults

`func NewTagCloneRequestVNextWithDefaults() *TagCloneRequestVNext`

NewTagCloneRequestVNextWithDefaults instantiates a new TagCloneRequestVNext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TagCloneRequestVNext) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagCloneRequestVNext) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagCloneRequestVNext) SetId(v string)`

SetId sets Id field to given value.


### GetLanguage

`func (o *TagCloneRequestVNext) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *TagCloneRequestVNext) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *TagCloneRequestVNext) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *TagCloneRequestVNext) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPrimaryLanguage

`func (o *TagCloneRequestVNext) GetPrimaryLanguage() string`

GetPrimaryLanguage returns the PrimaryLanguage field if non-nil, zero value otherwise.

### GetPrimaryLanguageOk

`func (o *TagCloneRequestVNext) GetPrimaryLanguageOk() (*string, bool)`

GetPrimaryLanguageOk returns a tuple with the PrimaryLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLanguage

`func (o *TagCloneRequestVNext) SetPrimaryLanguage(v string)`

SetPrimaryLanguage sets PrimaryLanguage field to given value.

### HasPrimaryLanguage

`func (o *TagCloneRequestVNext) HasPrimaryLanguage() bool`

HasPrimaryLanguage returns a boolean if a field has been set.

### GetName

`func (o *TagCloneRequestVNext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagCloneRequestVNext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagCloneRequestVNext) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


