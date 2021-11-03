# BlogAuthor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the Blog Author. | 
**Email** | **string** | Email address of the Blog Author. | 
**Bio** | **string** | A short biography of the blog author. | 
**Website** | **string** | URL to the website of the Blog Author. | 
**Twitter** | **string** | URL or username of the Twitter account associated with the Blog Author. This will be normalized into the Twitter url for said user. | 
**Facebook** | **string** | URL to the Blog Author&#39;s Facebook page. | 
**Linkedin** | **string** | URL to the blog author&#39;s LinkedIn page. | 
**Avatar** | **string** | URL to the blog author&#39;s avatar, if supplying a custom one. | 
**DisplayName** | **string** | The full name of the Blog Author to be displayed. | 
**DeletedAt** | **time.Time** | The timestamp (ISO8601 format) when this Blog Author was deleted. | 
**CreatedAt** | **time.Time** | The timestamp (ISO8601 format) when this Blog Author was created. | 
**UpdatedAt** | **time.Time** | The timestamp (ISO8601 format) when this Blog Author was last updated. | 

## Methods

### NewBlogAuthor

`func NewBlogAuthor(id string, email string, bio string, website string, twitter string, facebook string, linkedin string, avatar string, displayName string, deletedAt time.Time, createdAt time.Time, updatedAt time.Time, ) *BlogAuthor`

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


### GetCreatedAt

`func (o *BlogAuthor) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BlogAuthor) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BlogAuthor) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BlogAuthor) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BlogAuthor) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BlogAuthor) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


