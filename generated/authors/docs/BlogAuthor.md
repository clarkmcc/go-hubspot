# BlogAuthor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the Blog Author. | 
**FullName** | **string** |  | 
**Email** | **string** | Email address of the Blog Author. | 
**Slug** | **string** |  | 
**Language** | **string** | The explicitly defined ISO 639 language code of the blog author. | 
**TranslatedFromId** | **int64** | ID of the primary blog author this object was translated from. | 
**Name** | **string** |  | 
**DisplayName** | **string** | The full name of the Blog Author to be displayed. | 
**Bio** | **string** | A short biography of the blog author. | 
**Website** | **string** | URL to the website of the Blog Author. | 
**Twitter** | **string** | URL or username of the Twitter account associated with the Blog Author. This will be normalized into the Twitter url for said user. | 
**Facebook** | **string** | URL to the Blog Author&#39;s Facebook page. | 
**Linkedin** | **string** | URL to the blog author&#39;s LinkedIn page. | 
**Avatar** | **string** | URL to the blog author&#39;s avatar, if supplying a custom one. | 
**Created** | **time.Time** |  | 
**Updated** | **time.Time** |  | 
**DeletedAt** | **time.Time** | The timestamp (ISO8601 format) when this Blog Author was deleted. | 

## Methods

### NewBlogAuthor

`func NewBlogAuthor(id string, fullName string, email string, slug string, language string, translatedFromId int64, name string, displayName string, bio string, website string, twitter string, facebook string, linkedin string, avatar string, created time.Time, updated time.Time, deletedAt time.Time, ) *BlogAuthor`

NewBlogAuthor instantiates a new BlogAuthor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlogAuthorWithDefaults

`func NewBlogAuthorWithDefaults() *BlogAuthor`

NewBlogAuthorWithDefaults instantiates a new BlogAuthor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlogAuthor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlogAuthor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlogAuthor) SetId(v string)`

SetId sets Id field to given value.


### GetFullName

`func (o *BlogAuthor) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *BlogAuthor) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *BlogAuthor) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetEmail

`func (o *BlogAuthor) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BlogAuthor) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BlogAuthor) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetSlug

`func (o *BlogAuthor) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BlogAuthor) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BlogAuthor) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetLanguage

`func (o *BlogAuthor) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *BlogAuthor) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *BlogAuthor) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetTranslatedFromId

`func (o *BlogAuthor) GetTranslatedFromId() int64`

GetTranslatedFromId returns the TranslatedFromId field if non-nil, zero value otherwise.

### GetTranslatedFromIdOk

`func (o *BlogAuthor) GetTranslatedFromIdOk() (*int64, bool)`

GetTranslatedFromIdOk returns a tuple with the TranslatedFromId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedFromId

`func (o *BlogAuthor) SetTranslatedFromId(v int64)`

SetTranslatedFromId sets TranslatedFromId field to given value.


### GetName

`func (o *BlogAuthor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlogAuthor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlogAuthor) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *BlogAuthor) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BlogAuthor) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BlogAuthor) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetBio

`func (o *BlogAuthor) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *BlogAuthor) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *BlogAuthor) SetBio(v string)`

SetBio sets Bio field to given value.


### GetWebsite

`func (o *BlogAuthor) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *BlogAuthor) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *BlogAuthor) SetWebsite(v string)`

SetWebsite sets Website field to given value.


### GetTwitter

`func (o *BlogAuthor) GetTwitter() string`

GetTwitter returns the Twitter field if non-nil, zero value otherwise.

### GetTwitterOk

`func (o *BlogAuthor) GetTwitterOk() (*string, bool)`

GetTwitterOk returns a tuple with the Twitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitter

`func (o *BlogAuthor) SetTwitter(v string)`

SetTwitter sets Twitter field to given value.


### GetFacebook

`func (o *BlogAuthor) GetFacebook() string`

GetFacebook returns the Facebook field if non-nil, zero value otherwise.

### GetFacebookOk

`func (o *BlogAuthor) GetFacebookOk() (*string, bool)`

GetFacebookOk returns a tuple with the Facebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebook

`func (o *BlogAuthor) SetFacebook(v string)`

SetFacebook sets Facebook field to given value.


### GetLinkedin

`func (o *BlogAuthor) GetLinkedin() string`

GetLinkedin returns the Linkedin field if non-nil, zero value otherwise.

### GetLinkedinOk

`func (o *BlogAuthor) GetLinkedinOk() (*string, bool)`

GetLinkedinOk returns a tuple with the Linkedin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedin

`func (o *BlogAuthor) SetLinkedin(v string)`

SetLinkedin sets Linkedin field to given value.


### GetAvatar

`func (o *BlogAuthor) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *BlogAuthor) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *BlogAuthor) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### GetCreated

`func (o *BlogAuthor) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BlogAuthor) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BlogAuthor) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *BlogAuthor) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *BlogAuthor) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *BlogAuthor) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### GetDeletedAt

`func (o *BlogAuthor) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BlogAuthor) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BlogAuthor) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


