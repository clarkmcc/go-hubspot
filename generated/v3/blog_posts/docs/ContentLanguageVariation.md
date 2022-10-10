# ContentLanguageVariation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**State** | **string** |  | 
**AuthorName** | **string** |  | 
**Password** | **string** |  | 
**PublicAccessRulesEnabled** | **bool** |  | 
**PublicAccessRules** | **[]map[string]interface{}** |  | 
**Campaign** | **string** |  | 
**TagIds** | Pointer to **[]int64** |  | [optional] 
**ArchivedInDashboard** | **bool** |  | 
**Created** | **time.Time** |  | 
**Updated** | **time.Time** |  | 
**PublishDate** | **time.Time** |  | 

## Methods

### NewContentLanguageVariation

`func NewContentLanguageVariation(id int64, name string, slug string, state string, authorName string, password string, publicAccessRulesEnabled bool, publicAccessRules []map[string]interface{}, campaign string, archivedInDashboard bool, created time.Time, updated time.Time, publishDate time.Time, ) *ContentLanguageVariation`

NewContentLanguageVariation instantiates a new ContentLanguageVariation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentLanguageVariationWithDefaults

`func NewContentLanguageVariationWithDefaults() *ContentLanguageVariation`

NewContentLanguageVariationWithDefaults instantiates a new ContentLanguageVariation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContentLanguageVariation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContentLanguageVariation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContentLanguageVariation) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ContentLanguageVariation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentLanguageVariation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentLanguageVariation) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *ContentLanguageVariation) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ContentLanguageVariation) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ContentLanguageVariation) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetState

`func (o *ContentLanguageVariation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ContentLanguageVariation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ContentLanguageVariation) SetState(v string)`

SetState sets State field to given value.


### GetAuthorName

`func (o *ContentLanguageVariation) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *ContentLanguageVariation) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *ContentLanguageVariation) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.


### GetPassword

`func (o *ContentLanguageVariation) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ContentLanguageVariation) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ContentLanguageVariation) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPublicAccessRulesEnabled

`func (o *ContentLanguageVariation) GetPublicAccessRulesEnabled() bool`

GetPublicAccessRulesEnabled returns the PublicAccessRulesEnabled field if non-nil, zero value otherwise.

### GetPublicAccessRulesEnabledOk

`func (o *ContentLanguageVariation) GetPublicAccessRulesEnabledOk() (*bool, bool)`

GetPublicAccessRulesEnabledOk returns a tuple with the PublicAccessRulesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAccessRulesEnabled

`func (o *ContentLanguageVariation) SetPublicAccessRulesEnabled(v bool)`

SetPublicAccessRulesEnabled sets PublicAccessRulesEnabled field to given value.


### GetPublicAccessRules

`func (o *ContentLanguageVariation) GetPublicAccessRules() []map[string]interface{}`

GetPublicAccessRules returns the PublicAccessRules field if non-nil, zero value otherwise.

### GetPublicAccessRulesOk

`func (o *ContentLanguageVariation) GetPublicAccessRulesOk() (*[]map[string]interface{}, bool)`

GetPublicAccessRulesOk returns a tuple with the PublicAccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAccessRules

`func (o *ContentLanguageVariation) SetPublicAccessRules(v []map[string]interface{})`

SetPublicAccessRules sets PublicAccessRules field to given value.


### GetCampaign

`func (o *ContentLanguageVariation) GetCampaign() string`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *ContentLanguageVariation) GetCampaignOk() (*string, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *ContentLanguageVariation) SetCampaign(v string)`

SetCampaign sets Campaign field to given value.


### GetTagIds

`func (o *ContentLanguageVariation) GetTagIds() []int64`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *ContentLanguageVariation) GetTagIdsOk() (*[]int64, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *ContentLanguageVariation) SetTagIds(v []int64)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *ContentLanguageVariation) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetArchivedInDashboard

`func (o *ContentLanguageVariation) GetArchivedInDashboard() bool`

GetArchivedInDashboard returns the ArchivedInDashboard field if non-nil, zero value otherwise.

### GetArchivedInDashboardOk

`func (o *ContentLanguageVariation) GetArchivedInDashboardOk() (*bool, bool)`

GetArchivedInDashboardOk returns a tuple with the ArchivedInDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedInDashboard

`func (o *ContentLanguageVariation) SetArchivedInDashboard(v bool)`

SetArchivedInDashboard sets ArchivedInDashboard field to given value.


### GetCreated

`func (o *ContentLanguageVariation) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ContentLanguageVariation) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ContentLanguageVariation) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *ContentLanguageVariation) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ContentLanguageVariation) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ContentLanguageVariation) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### GetPublishDate

`func (o *ContentLanguageVariation) GetPublishDate() time.Time`

GetPublishDate returns the PublishDate field if non-nil, zero value otherwise.

### GetPublishDateOk

`func (o *ContentLanguageVariation) GetPublishDateOk() (*time.Time, bool)`

GetPublishDateOk returns a tuple with the PublishDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishDate

`func (o *ContentLanguageVariation) SetPublishDate(v time.Time)`

SetPublishDate sets PublishDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


