# VersionBlogPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | [**BlogPost**](BlogPost.md) |  | 
**User** | [**VersionUser**](VersionUser.md) |  | 
**Id** | **string** | The id of the version. | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewVersionBlogPost

`func NewVersionBlogPost(object BlogPost, user VersionUser, id string, updatedAt time.Time, ) *VersionBlogPost`

NewVersionBlogPost instantiates a new VersionBlogPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionBlogPostWithDefaults

`func NewVersionBlogPostWithDefaults() *VersionBlogPost`

NewVersionBlogPostWithDefaults instantiates a new VersionBlogPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *VersionBlogPost) GetObject() BlogPost`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *VersionBlogPost) GetObjectOk() (*BlogPost, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *VersionBlogPost) SetObject(v BlogPost)`

SetObject sets Object field to given value.


### GetUser

`func (o *VersionBlogPost) GetUser() VersionUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *VersionBlogPost) GetUserOk() (*VersionUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *VersionBlogPost) SetUser(v VersionUser)`

SetUser sets User field to given value.


### GetId

`func (o *VersionBlogPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VersionBlogPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VersionBlogPost) SetId(v string)`

SetId sets Id field to given value.


### GetUpdatedAt

`func (o *VersionBlogPost) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VersionBlogPost) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VersionBlogPost) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


