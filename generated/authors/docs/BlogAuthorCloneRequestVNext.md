# BlogAuthorCloneRequestVNext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the object to be cloned. | 
**Language** | Pointer to **string** | Language of newly cloned object. | [optional] 
**PrimaryLanguage** | Pointer to **string** | Primary language in multi-language group. | [optional] 
**BlogAuthor** | [**BlogAuthor**](BlogAuthor.md) |  | 

## Methods

### NewBlogAuthorCloneRequestVNext

`func NewBlogAuthorCloneRequestVNext(id string, blogAuthor BlogAuthor, ) *BlogAuthorCloneRequestVNext`

NewBlogAuthorCloneRequestVNext instantiates a new BlogAuthorCloneRequestVNext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlogAuthorCloneRequestVNextWithDefaults

`func NewBlogAuthorCloneRequestVNextWithDefaults() *BlogAuthorCloneRequestVNext`

NewBlogAuthorCloneRequestVNextWithDefaults instantiates a new BlogAuthorCloneRequestVNext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlogAuthorCloneRequestVNext) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlogAuthorCloneRequestVNext) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlogAuthorCloneRequestVNext) SetId(v string)`

SetId sets Id field to given value.


### GetLanguage

`func (o *BlogAuthorCloneRequestVNext) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *BlogAuthorCloneRequestVNext) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *BlogAuthorCloneRequestVNext) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *BlogAuthorCloneRequestVNext) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPrimaryLanguage

`func (o *BlogAuthorCloneRequestVNext) GetPrimaryLanguage() string`

GetPrimaryLanguage returns the PrimaryLanguage field if non-nil, zero value otherwise.

### GetPrimaryLanguageOk

`func (o *BlogAuthorCloneRequestVNext) GetPrimaryLanguageOk() (*string, bool)`

GetPrimaryLanguageOk returns a tuple with the PrimaryLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLanguage

`func (o *BlogAuthorCloneRequestVNext) SetPrimaryLanguage(v string)`

SetPrimaryLanguage sets PrimaryLanguage field to given value.

### HasPrimaryLanguage

`func (o *BlogAuthorCloneRequestVNext) HasPrimaryLanguage() bool`

HasPrimaryLanguage returns a boolean if a field has been set.

### GetBlogAuthor

`func (o *BlogAuthorCloneRequestVNext) GetBlogAuthor() BlogAuthor`

GetBlogAuthor returns the BlogAuthor field if non-nil, zero value otherwise.

### GetBlogAuthorOk

`func (o *BlogAuthorCloneRequestVNext) GetBlogAuthorOk() (*BlogAuthor, bool)`

GetBlogAuthorOk returns a tuple with the BlogAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogAuthor

`func (o *BlogAuthorCloneRequestVNext) SetBlogAuthor(v BlogAuthor)`

SetBlogAuthor sets BlogAuthor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


