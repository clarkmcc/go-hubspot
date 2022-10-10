# ContentSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the content. | 
**Score** | **float32** | The matching score of the document. | 
**Type** | **string** | The type of document. Can be &#x60;SITE_PAGE&#x60;, &#x60;LANDING_PAGE&#x60;, &#x60;BLOG_POST&#x60;, &#x60;LISTING_PAGE&#x60;, or &#x60;KNOWLEDGE_ARTICLE&#x60;. | 
**Domain** | **string** | The domain the document is hosted on. | 
**Url** | **string** | The url of the document. | 
**FeaturedImageUrl** | Pointer to **string** | URL of the featured image. | [optional] 
**Language** | Pointer to **string** | The document&#39;s language. | [optional] 
**Title** | Pointer to **string** | The title of the returned document. | [optional] 
**Description** | Pointer to **string** | The result&#39;s description. The content will be determined by the value of &#x60;length&#x60; in the request. | [optional] 
**Category** | Pointer to **string** | For knowledge articles, the category of the article. | [optional] 
**Subcategory** | Pointer to **string** | For knowledge articles, the subcategory of the article. | [optional] 
**AuthorFullName** | Pointer to **string** | Name of the author. | [optional] 
**Tags** | Pointer to **[]string** | If a blog post, the tags associated with it. | [optional] 
**TableId** | Pointer to **int64** | If a dynamic page, the ID of the HubDB table. | [optional] 
**RowId** | Pointer to **int64** | If a dynamic page, the row ID in the HubDB table. | [optional] 
**PublishedDate** | Pointer to **int64** | The date the content was published. | [optional] 
**CombinedId** | Pointer to **string** | The ID of the document in HubSpot. | [optional] 

## Methods

### NewContentSearchResult

`func NewContentSearchResult(id int32, score float32, type_ string, domain string, url string, ) *ContentSearchResult`

NewContentSearchResult instantiates a new ContentSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentSearchResultWithDefaults

`func NewContentSearchResultWithDefaults() *ContentSearchResult`

NewContentSearchResultWithDefaults instantiates a new ContentSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContentSearchResult) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContentSearchResult) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContentSearchResult) SetId(v int32)`

SetId sets Id field to given value.


### GetScore

`func (o *ContentSearchResult) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ContentSearchResult) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ContentSearchResult) SetScore(v float32)`

SetScore sets Score field to given value.


### GetType

`func (o *ContentSearchResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentSearchResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentSearchResult) SetType(v string)`

SetType sets Type field to given value.


### GetDomain

`func (o *ContentSearchResult) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ContentSearchResult) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ContentSearchResult) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetUrl

`func (o *ContentSearchResult) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ContentSearchResult) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ContentSearchResult) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetFeaturedImageUrl

`func (o *ContentSearchResult) GetFeaturedImageUrl() string`

GetFeaturedImageUrl returns the FeaturedImageUrl field if non-nil, zero value otherwise.

### GetFeaturedImageUrlOk

`func (o *ContentSearchResult) GetFeaturedImageUrlOk() (*string, bool)`

GetFeaturedImageUrlOk returns a tuple with the FeaturedImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedImageUrl

`func (o *ContentSearchResult) SetFeaturedImageUrl(v string)`

SetFeaturedImageUrl sets FeaturedImageUrl field to given value.

### HasFeaturedImageUrl

`func (o *ContentSearchResult) HasFeaturedImageUrl() bool`

HasFeaturedImageUrl returns a boolean if a field has been set.

### GetLanguage

`func (o *ContentSearchResult) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ContentSearchResult) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ContentSearchResult) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ContentSearchResult) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetTitle

`func (o *ContentSearchResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ContentSearchResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ContentSearchResult) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ContentSearchResult) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ContentSearchResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContentSearchResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContentSearchResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContentSearchResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *ContentSearchResult) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ContentSearchResult) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ContentSearchResult) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ContentSearchResult) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetSubcategory

`func (o *ContentSearchResult) GetSubcategory() string`

GetSubcategory returns the Subcategory field if non-nil, zero value otherwise.

### GetSubcategoryOk

`func (o *ContentSearchResult) GetSubcategoryOk() (*string, bool)`

GetSubcategoryOk returns a tuple with the Subcategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcategory

`func (o *ContentSearchResult) SetSubcategory(v string)`

SetSubcategory sets Subcategory field to given value.

### HasSubcategory

`func (o *ContentSearchResult) HasSubcategory() bool`

HasSubcategory returns a boolean if a field has been set.

### GetAuthorFullName

`func (o *ContentSearchResult) GetAuthorFullName() string`

GetAuthorFullName returns the AuthorFullName field if non-nil, zero value otherwise.

### GetAuthorFullNameOk

`func (o *ContentSearchResult) GetAuthorFullNameOk() (*string, bool)`

GetAuthorFullNameOk returns a tuple with the AuthorFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorFullName

`func (o *ContentSearchResult) SetAuthorFullName(v string)`

SetAuthorFullName sets AuthorFullName field to given value.

### HasAuthorFullName

`func (o *ContentSearchResult) HasAuthorFullName() bool`

HasAuthorFullName returns a boolean if a field has been set.

### GetTags

`func (o *ContentSearchResult) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ContentSearchResult) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ContentSearchResult) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ContentSearchResult) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTableId

`func (o *ContentSearchResult) GetTableId() int64`

GetTableId returns the TableId field if non-nil, zero value otherwise.

### GetTableIdOk

`func (o *ContentSearchResult) GetTableIdOk() (*int64, bool)`

GetTableIdOk returns a tuple with the TableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableId

`func (o *ContentSearchResult) SetTableId(v int64)`

SetTableId sets TableId field to given value.

### HasTableId

`func (o *ContentSearchResult) HasTableId() bool`

HasTableId returns a boolean if a field has been set.

### GetRowId

`func (o *ContentSearchResult) GetRowId() int64`

GetRowId returns the RowId field if non-nil, zero value otherwise.

### GetRowIdOk

`func (o *ContentSearchResult) GetRowIdOk() (*int64, bool)`

GetRowIdOk returns a tuple with the RowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowId

`func (o *ContentSearchResult) SetRowId(v int64)`

SetRowId sets RowId field to given value.

### HasRowId

`func (o *ContentSearchResult) HasRowId() bool`

HasRowId returns a boolean if a field has been set.

### GetPublishedDate

`func (o *ContentSearchResult) GetPublishedDate() int64`

GetPublishedDate returns the PublishedDate field if non-nil, zero value otherwise.

### GetPublishedDateOk

`func (o *ContentSearchResult) GetPublishedDateOk() (*int64, bool)`

GetPublishedDateOk returns a tuple with the PublishedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedDate

`func (o *ContentSearchResult) SetPublishedDate(v int64)`

SetPublishedDate sets PublishedDate field to given value.

### HasPublishedDate

`func (o *ContentSearchResult) HasPublishedDate() bool`

HasPublishedDate returns a boolean if a field has been set.

### GetCombinedId

`func (o *ContentSearchResult) GetCombinedId() string`

GetCombinedId returns the CombinedId field if non-nil, zero value otherwise.

### GetCombinedIdOk

`func (o *ContentSearchResult) GetCombinedIdOk() (*string, bool)`

GetCombinedIdOk returns a tuple with the CombinedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedId

`func (o *ContentSearchResult) SetCombinedId(v string)`

SetCombinedId sets CombinedId field to given value.

### HasCombinedId

`func (o *ContentSearchResult) HasCombinedId() bool`

HasCombinedId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


