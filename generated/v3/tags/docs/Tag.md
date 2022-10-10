# Tag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the Blog Tag. | 
**Name** | **string** | The name of the tag. | 
**Language** | **string** | The explicitly defined ISO 639 language code of the tag. | 
**TranslatedFromId** | **int64** | ID of the primary tag this object was translated from. | 
**Created** | **time.Time** |  | 
**Updated** | **time.Time** |  | 
**DeletedAt** | **time.Time** | The timestamp (ISO8601 format) when this Blog Tag was deleted. | 

## Methods

### NewTag

`func NewTag(id string, name string, language string, translatedFromId int64, created time.Time, updated time.Time, deletedAt time.Time, ) *Tag`

NewTag instantiates a new Tag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagWithDefaults

`func NewTagWithDefaults() *Tag`

NewTagWithDefaults instantiates a new Tag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Tag) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tag) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tag) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Tag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tag) SetName(v string)`

SetName sets Name field to given value.


### GetLanguage

`func (o *Tag) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Tag) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Tag) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetTranslatedFromId

`func (o *Tag) GetTranslatedFromId() int64`

GetTranslatedFromId returns the TranslatedFromId field if non-nil, zero value otherwise.

### GetTranslatedFromIdOk

`func (o *Tag) GetTranslatedFromIdOk() (*int64, bool)`

GetTranslatedFromIdOk returns a tuple with the TranslatedFromId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedFromId

`func (o *Tag) SetTranslatedFromId(v int64)`

SetTranslatedFromId sets TranslatedFromId field to given value.


### GetCreated

`func (o *Tag) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Tag) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Tag) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *Tag) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Tag) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Tag) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### GetDeletedAt

`func (o *Tag) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Tag) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Tag) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


