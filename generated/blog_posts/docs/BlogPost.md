# BlogPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the Blog Post. | 
**Slug** | **string** | The path of the this blog post. This field is appended to the domain to construct the url of this post. | 
**ContentGroupId** | **string** | The ID of the parent Blog this Blog Post is associated with. | 
**Campaign** | **string** | The GUID of the marketing campaign this Blog Post is a part of. | 
**CategoryId** | **int32** | ID of the type of object this is. Should always . | 
**State** | **string** | An ENUM descibing the current state of this Blog Post. | 
**TemplatePath** | **string** |  | 
**Name** | **string** | The internal name of the blog post. | 
**MabExperimentId** | **string** |  | 
**Archived** | **bool** | If True, the post will not show up in your dashboard, although the post could still be live. | 
**AuthorName** | **string** | The name of the user that updated this blog post. | 
**AbTestId** | **string** |  | 
**CreatedById** | **string** | The ID of the user that created this blog post. | 
**UpdatedById** | **string** | The ID of the user that updated this blog post. | 
**Domain** | **string** | The domain this Blog Post will resolve to. If null, the Blog Post will default to the domain of the ParentBlog. | 
**Subcategory** | **string** |  | 
**AbStatus** | **string** |  | 
**FolderId** | **string** |  | 
**WidgetContainers** | **map[string]map[string]interface{}** | A data structure containing the data for all the modules inside the containers for this post. This will only be populated if the page has widget containers. | 
**Widgets** | **map[string]map[string]interface{}** | A data structure containing the data for all the modules for this page. | 
**Language** | **string** | The explicitly defined language of the Blog Post. If null, the Blog Post will default to the language of the ParentBlog. | 
**TranslatedFromId** | **string** | ID of the primary blog post this object was translated from. | 
**DynamicPageHubDbTableId** | **string** |  | 
**BlogAuthorId** | **string** | The ID of the Blog Author associated with this Blog Post. | 
**TagIds** | **[]int64** | List of IDs for the tags associated with this Blog Post. | 
**PostBody** | **string** | The HTML of the main post body. | 
**PostSummary** | **string** | The summary of the blog post that will appear on the main listing page. | 
**RssBody** | **string** | The contents of the RSS body for this Blog Post. | 
**RssSummary** | **string** | The contents of the RSS summary for this Blog Post. | 
**EnableGoogleAmpOutputOverride** | **bool** | Boolean to allow overriding the AMP settings for the blog. | 
**HtmlTitle** | **string** | The html title of this Blog Post. | 
**PageRedirected** | **bool** |  | 
**PageExpiryEnabled** | **bool** |  | 
**PageExpiryDate** | **int64** |  | 
**PageExpiryRedirectId** | **int64** |  | 
**PageExpiryRedirectUrl** | **string** |  | 
**UseFeaturedImage** | **bool** | Boolean to determine if this post should use a featuredImage. | 
**Password** | **string** | Set this to create a password protected page. Entering the password will be required to view the page. | 
**AttachedStylesheets** | **[]map[string]map[string]interface{}** | List of stylesheets to attach to this blog post. These stylesheets are attached to just this page. Order of precedence is bottom to top, just like in the HTML. | 
**IncludeDefaultCustomCss** | **bool** | Boolean to determine whether or not the Primary CSS Files should be applied. | 
**EnableDomainStylesheets** | **bool** | Boolean to determine whether or not the styles from the template should be applied. | 
**EnableLayoutStylesheets** | **bool** | Boolean to determine whether or not the styles from the template should be applied. | 
**MetaDescription** | **string** | A description that goes in &lt;meta&gt; tag on the page. | 
**PublishImmediately** | **bool** | Set this to true if you want to be published immediately when the schedule publish endpoint is called, and to ignore the publish_date setting. | 
**HeadHtml** | **string** | Custom HTML for embed codes, javascript, etc. that goes in the &lt;head&gt; tag of the page. | 
**FooterHtml** | **string** | Custom HTML for embed codes, javascript that should be placed before the &lt;/body&gt; tag of the page. | 
**ContentTypeCategory** | **string** | An ENUM descibing the type of this object. Should always be BLOG_POST. | 
**CurrentState** | **string** | A generated ENUM descibing the current state of this Blog Post. Should always match state. | 
**LinkRelCanonicalUrl** | **string** | Optional override to set the URL to be used in the rel&#x3D;canonical link tag on the page. | 
**FeaturedImage** | **string** | The featuredImage of this Blog Post. | 
**FeaturedImageAltText** | **string** | Alt Text of the featuredImage. | 
**PublicAccessRulesEnabled** | **bool** | Boolean to determine whether or not to respect publicAccessRules. | 
**PublicAccessRules** | **[]map[string]interface{}** | Rules for require member registration to access private content. | 
**LayoutSections** | [**map[string]LayoutSection**](LayoutSection.md) |  | 
**ThemeSettingsValues** | **map[string]map[string]interface{}** |  | 
**Url** | **string** | A generated field representing the URL of this blog post. | 
**PublishDate** | **time.Time** | The date (ISO8601 format) the blog post is to be published at. | 
**DeletedAt** | **time.Time** | The timestamp (ISO8601 format) when this Blog Post was deleted. | 
**CreatedAt** | **time.Time** | The timestamp (ISO8601 format) when this blog post was created. | 
**Published** | **bool** | Boolean describing if this Blog Post is published. | 
**UpdatedAt** | **time.Time** | The timestamp (ISO8601 format) when this Blog Post was last updated. | 

## Methods

### NewBlogPost

`func NewBlogPost(id string, slug string, contentGroupId string, campaign string, categoryId int32, state string, templatePath string, name string, mabExperimentId string, archived bool, authorName string, abTestId string, createdById string, updatedById string, domain string, subcategory string, abStatus string, folderId string, widgetContainers map[string]map[string]interface{}, widgets map[string]map[string]interface{}, language string, translatedFromId string, dynamicPageHubDbTableId string, blogAuthorId string, tagIds []int64, postBody string, postSummary string, rssBody string, rssSummary string, enableGoogleAmpOutputOverride bool, htmlTitle string, pageRedirected bool, pageExpiryEnabled bool, pageExpiryDate int64, pageExpiryRedirectId int64, pageExpiryRedirectUrl string, useFeaturedImage bool, password string, attachedStylesheets []map[string]map[string]interface{}, includeDefaultCustomCss bool, enableDomainStylesheets bool, enableLayoutStylesheets bool, metaDescription string, publishImmediately bool, headHtml string, footerHtml string, contentTypeCategory string, currentState string, linkRelCanonicalUrl string, featuredImage string, featuredImageAltText string, publicAccessRulesEnabled bool, publicAccessRules []map[string]interface{}, layoutSections map[string]LayoutSection, themeSettingsValues map[string]map[string]interface{}, url string, publishDate time.Time, deletedAt time.Time, createdAt time.Time, published bool, updatedAt time.Time, ) *BlogPost`

NewBlogPost instantiates a new BlogPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlogPostWithDefaults

`func NewBlogPostWithDefaults() *BlogPost`

NewBlogPostWithDefaults instantiates a new BlogPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlogPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlogPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlogPost) SetId(v string)`

SetId sets Id field to given value.


### GetSlug

`func (o *BlogPost) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BlogPost) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BlogPost) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetContentGroupId

`func (o *BlogPost) GetContentGroupId() string`

GetContentGroupId returns the ContentGroupId field if non-nil, zero value otherwise.

### GetContentGroupIdOk

`func (o *BlogPost) GetContentGroupIdOk() (*string, bool)`

GetContentGroupIdOk returns a tuple with the ContentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentGroupId

`func (o *BlogPost) SetContentGroupId(v string)`

SetContentGroupId sets ContentGroupId field to given value.


### GetCampaign

`func (o *BlogPost) GetCampaign() string`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *BlogPost) GetCampaignOk() (*string, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *BlogPost) SetCampaign(v string)`

SetCampaign sets Campaign field to given value.


### GetCategoryId

`func (o *BlogPost) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *BlogPost) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *BlogPost) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.


### GetState

`func (o *BlogPost) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BlogPost) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BlogPost) SetState(v string)`

SetState sets State field to given value.


### GetTemplatePath

`func (o *BlogPost) GetTemplatePath() string`

GetTemplatePath returns the TemplatePath field if non-nil, zero value otherwise.

### GetTemplatePathOk

`func (o *BlogPost) GetTemplatePathOk() (*string, bool)`

GetTemplatePathOk returns a tuple with the TemplatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePath

`func (o *BlogPost) SetTemplatePath(v string)`

SetTemplatePath sets TemplatePath field to given value.


### GetName

`func (o *BlogPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlogPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlogPost) SetName(v string)`

SetName sets Name field to given value.


### GetMabExperimentId

`func (o *BlogPost) GetMabExperimentId() string`

GetMabExperimentId returns the MabExperimentId field if non-nil, zero value otherwise.

### GetMabExperimentIdOk

`func (o *BlogPost) GetMabExperimentIdOk() (*string, bool)`

GetMabExperimentIdOk returns a tuple with the MabExperimentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMabExperimentId

`func (o *BlogPost) SetMabExperimentId(v string)`

SetMabExperimentId sets MabExperimentId field to given value.


### GetArchived

`func (o *BlogPost) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *BlogPost) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *BlogPost) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetAuthorName

`func (o *BlogPost) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *BlogPost) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *BlogPost) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.


### GetAbTestId

`func (o *BlogPost) GetAbTestId() string`

GetAbTestId returns the AbTestId field if non-nil, zero value otherwise.

### GetAbTestIdOk

`func (o *BlogPost) GetAbTestIdOk() (*string, bool)`

GetAbTestIdOk returns a tuple with the AbTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbTestId

`func (o *BlogPost) SetAbTestId(v string)`

SetAbTestId sets AbTestId field to given value.


### GetCreatedById

`func (o *BlogPost) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *BlogPost) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *BlogPost) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetUpdatedById

`func (o *BlogPost) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *BlogPost) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *BlogPost) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetDomain

`func (o *BlogPost) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *BlogPost) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *BlogPost) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetSubcategory

`func (o *BlogPost) GetSubcategory() string`

GetSubcategory returns the Subcategory field if non-nil, zero value otherwise.

### GetSubcategoryOk

`func (o *BlogPost) GetSubcategoryOk() (*string, bool)`

GetSubcategoryOk returns a tuple with the Subcategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcategory

`func (o *BlogPost) SetSubcategory(v string)`

SetSubcategory sets Subcategory field to given value.


### GetAbStatus

`func (o *BlogPost) GetAbStatus() string`

GetAbStatus returns the AbStatus field if non-nil, zero value otherwise.

### GetAbStatusOk

`func (o *BlogPost) GetAbStatusOk() (*string, bool)`

GetAbStatusOk returns a tuple with the AbStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbStatus

`func (o *BlogPost) SetAbStatus(v string)`

SetAbStatus sets AbStatus field to given value.


### GetFolderId

`func (o *BlogPost) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *BlogPost) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *BlogPost) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.


### GetWidgetContainers

`func (o *BlogPost) GetWidgetContainers() map[string]map[string]interface{}`

GetWidgetContainers returns the WidgetContainers field if non-nil, zero value otherwise.

### GetWidgetContainersOk

`func (o *BlogPost) GetWidgetContainersOk() (*map[string]map[string]interface{}, bool)`

GetWidgetContainersOk returns a tuple with the WidgetContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetContainers

`func (o *BlogPost) SetWidgetContainers(v map[string]map[string]interface{})`

SetWidgetContainers sets WidgetContainers field to given value.


### GetWidgets

`func (o *BlogPost) GetWidgets() map[string]map[string]interface{}`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *BlogPost) GetWidgetsOk() (*map[string]map[string]interface{}, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgets

`func (o *BlogPost) SetWidgets(v map[string]map[string]interface{})`

SetWidgets sets Widgets field to given value.


### GetLanguage

`func (o *BlogPost) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *BlogPost) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *BlogPost) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetTranslatedFromId

`func (o *BlogPost) GetTranslatedFromId() string`

GetTranslatedFromId returns the TranslatedFromId field if non-nil, zero value otherwise.

### GetTranslatedFromIdOk

`func (o *BlogPost) GetTranslatedFromIdOk() (*string, bool)`

GetTranslatedFromIdOk returns a tuple with the TranslatedFromId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedFromId

`func (o *BlogPost) SetTranslatedFromId(v string)`

SetTranslatedFromId sets TranslatedFromId field to given value.


### GetDynamicPageHubDbTableId

`func (o *BlogPost) GetDynamicPageHubDbTableId() string`

GetDynamicPageHubDbTableId returns the DynamicPageHubDbTableId field if non-nil, zero value otherwise.

### GetDynamicPageHubDbTableIdOk

`func (o *BlogPost) GetDynamicPageHubDbTableIdOk() (*string, bool)`

GetDynamicPageHubDbTableIdOk returns a tuple with the DynamicPageHubDbTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicPageHubDbTableId

`func (o *BlogPost) SetDynamicPageHubDbTableId(v string)`

SetDynamicPageHubDbTableId sets DynamicPageHubDbTableId field to given value.


### GetBlogAuthorId

`func (o *BlogPost) GetBlogAuthorId() string`

GetBlogAuthorId returns the BlogAuthorId field if non-nil, zero value otherwise.

### GetBlogAuthorIdOk

`func (o *BlogPost) GetBlogAuthorIdOk() (*string, bool)`

GetBlogAuthorIdOk returns a tuple with the BlogAuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogAuthorId

`func (o *BlogPost) SetBlogAuthorId(v string)`

SetBlogAuthorId sets BlogAuthorId field to given value.


### GetTagIds

`func (o *BlogPost) GetTagIds() []int64`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *BlogPost) GetTagIdsOk() (*[]int64, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *BlogPost) SetTagIds(v []int64)`

SetTagIds sets TagIds field to given value.


### GetPostBody

`func (o *BlogPost) GetPostBody() string`

GetPostBody returns the PostBody field if non-nil, zero value otherwise.

### GetPostBodyOk

`func (o *BlogPost) GetPostBodyOk() (*string, bool)`

GetPostBodyOk returns a tuple with the PostBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostBody

`func (o *BlogPost) SetPostBody(v string)`

SetPostBody sets PostBody field to given value.


### GetPostSummary

`func (o *BlogPost) GetPostSummary() string`

GetPostSummary returns the PostSummary field if non-nil, zero value otherwise.

### GetPostSummaryOk

`func (o *BlogPost) GetPostSummaryOk() (*string, bool)`

GetPostSummaryOk returns a tuple with the PostSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostSummary

`func (o *BlogPost) SetPostSummary(v string)`

SetPostSummary sets PostSummary field to given value.


### GetRssBody

`func (o *BlogPost) GetRssBody() string`

GetRssBody returns the RssBody field if non-nil, zero value otherwise.

### GetRssBodyOk

`func (o *BlogPost) GetRssBodyOk() (*string, bool)`

GetRssBodyOk returns a tuple with the RssBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssBody

`func (o *BlogPost) SetRssBody(v string)`

SetRssBody sets RssBody field to given value.


### GetRssSummary

`func (o *BlogPost) GetRssSummary() string`

GetRssSummary returns the RssSummary field if non-nil, zero value otherwise.

### GetRssSummaryOk

`func (o *BlogPost) GetRssSummaryOk() (*string, bool)`

GetRssSummaryOk returns a tuple with the RssSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssSummary

`func (o *BlogPost) SetRssSummary(v string)`

SetRssSummary sets RssSummary field to given value.


### GetEnableGoogleAmpOutputOverride

`func (o *BlogPost) GetEnableGoogleAmpOutputOverride() bool`

GetEnableGoogleAmpOutputOverride returns the EnableGoogleAmpOutputOverride field if non-nil, zero value otherwise.

### GetEnableGoogleAmpOutputOverrideOk

`func (o *BlogPost) GetEnableGoogleAmpOutputOverrideOk() (*bool, bool)`

GetEnableGoogleAmpOutputOverrideOk returns a tuple with the EnableGoogleAmpOutputOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGoogleAmpOutputOverride

`func (o *BlogPost) SetEnableGoogleAmpOutputOverride(v bool)`

SetEnableGoogleAmpOutputOverride sets EnableGoogleAmpOutputOverride field to given value.


### GetHtmlTitle

`func (o *BlogPost) GetHtmlTitle() string`

GetHtmlTitle returns the HtmlTitle field if non-nil, zero value otherwise.

### GetHtmlTitleOk

`func (o *BlogPost) GetHtmlTitleOk() (*string, bool)`

GetHtmlTitleOk returns a tuple with the HtmlTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlTitle

`func (o *BlogPost) SetHtmlTitle(v string)`

SetHtmlTitle sets HtmlTitle field to given value.


### GetPageRedirected

`func (o *BlogPost) GetPageRedirected() bool`

GetPageRedirected returns the PageRedirected field if non-nil, zero value otherwise.

### GetPageRedirectedOk

`func (o *BlogPost) GetPageRedirectedOk() (*bool, bool)`

GetPageRedirectedOk returns a tuple with the PageRedirected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageRedirected

`func (o *BlogPost) SetPageRedirected(v bool)`

SetPageRedirected sets PageRedirected field to given value.


### GetPageExpiryEnabled

`func (o *BlogPost) GetPageExpiryEnabled() bool`

GetPageExpiryEnabled returns the PageExpiryEnabled field if non-nil, zero value otherwise.

### GetPageExpiryEnabledOk

`func (o *BlogPost) GetPageExpiryEnabledOk() (*bool, bool)`

GetPageExpiryEnabledOk returns a tuple with the PageExpiryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageExpiryEnabled

`func (o *BlogPost) SetPageExpiryEnabled(v bool)`

SetPageExpiryEnabled sets PageExpiryEnabled field to given value.


### GetPageExpiryDate

`func (o *BlogPost) GetPageExpiryDate() int64`

GetPageExpiryDate returns the PageExpiryDate field if non-nil, zero value otherwise.

### GetPageExpiryDateOk

`func (o *BlogPost) GetPageExpiryDateOk() (*int64, bool)`

GetPageExpiryDateOk returns a tuple with the PageExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageExpiryDate

`func (o *BlogPost) SetPageExpiryDate(v int64)`

SetPageExpiryDate sets PageExpiryDate field to given value.


### GetPageExpiryRedirectId

`func (o *BlogPost) GetPageExpiryRedirectId() int64`

GetPageExpiryRedirectId returns the PageExpiryRedirectId field if non-nil, zero value otherwise.

### GetPageExpiryRedirectIdOk

`func (o *BlogPost) GetPageExpiryRedirectIdOk() (*int64, bool)`

GetPageExpiryRedirectIdOk returns a tuple with the PageExpiryRedirectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageExpiryRedirectId

`func (o *BlogPost) SetPageExpiryRedirectId(v int64)`

SetPageExpiryRedirectId sets PageExpiryRedirectId field to given value.


### GetPageExpiryRedirectUrl

`func (o *BlogPost) GetPageExpiryRedirectUrl() string`

GetPageExpiryRedirectUrl returns the PageExpiryRedirectUrl field if non-nil, zero value otherwise.

### GetPageExpiryRedirectUrlOk

`func (o *BlogPost) GetPageExpiryRedirectUrlOk() (*string, bool)`

GetPageExpiryRedirectUrlOk returns a tuple with the PageExpiryRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageExpiryRedirectUrl

`func (o *BlogPost) SetPageExpiryRedirectUrl(v string)`

SetPageExpiryRedirectUrl sets PageExpiryRedirectUrl field to given value.


### GetUseFeaturedImage

`func (o *BlogPost) GetUseFeaturedImage() bool`

GetUseFeaturedImage returns the UseFeaturedImage field if non-nil, zero value otherwise.

### GetUseFeaturedImageOk

`func (o *BlogPost) GetUseFeaturedImageOk() (*bool, bool)`

GetUseFeaturedImageOk returns a tuple with the UseFeaturedImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFeaturedImage

`func (o *BlogPost) SetUseFeaturedImage(v bool)`

SetUseFeaturedImage sets UseFeaturedImage field to given value.


### GetPassword

`func (o *BlogPost) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BlogPost) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BlogPost) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetAttachedStylesheets

`func (o *BlogPost) GetAttachedStylesheets() []map[string]map[string]interface{}`

GetAttachedStylesheets returns the AttachedStylesheets field if non-nil, zero value otherwise.

### GetAttachedStylesheetsOk

`func (o *BlogPost) GetAttachedStylesheetsOk() (*[]map[string]map[string]interface{}, bool)`

GetAttachedStylesheetsOk returns a tuple with the AttachedStylesheets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedStylesheets

`func (o *BlogPost) SetAttachedStylesheets(v []map[string]map[string]interface{})`

SetAttachedStylesheets sets AttachedStylesheets field to given value.


### GetIncludeDefaultCustomCss

`func (o *BlogPost) GetIncludeDefaultCustomCss() bool`

GetIncludeDefaultCustomCss returns the IncludeDefaultCustomCss field if non-nil, zero value otherwise.

### GetIncludeDefaultCustomCssOk

`func (o *BlogPost) GetIncludeDefaultCustomCssOk() (*bool, bool)`

GetIncludeDefaultCustomCssOk returns a tuple with the IncludeDefaultCustomCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDefaultCustomCss

`func (o *BlogPost) SetIncludeDefaultCustomCss(v bool)`

SetIncludeDefaultCustomCss sets IncludeDefaultCustomCss field to given value.


### GetEnableDomainStylesheets

`func (o *BlogPost) GetEnableDomainStylesheets() bool`

GetEnableDomainStylesheets returns the EnableDomainStylesheets field if non-nil, zero value otherwise.

### GetEnableDomainStylesheetsOk

`func (o *BlogPost) GetEnableDomainStylesheetsOk() (*bool, bool)`

GetEnableDomainStylesheetsOk returns a tuple with the EnableDomainStylesheets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDomainStylesheets

`func (o *BlogPost) SetEnableDomainStylesheets(v bool)`

SetEnableDomainStylesheets sets EnableDomainStylesheets field to given value.


### GetEnableLayoutStylesheets

`func (o *BlogPost) GetEnableLayoutStylesheets() bool`

GetEnableLayoutStylesheets returns the EnableLayoutStylesheets field if non-nil, zero value otherwise.

### GetEnableLayoutStylesheetsOk

`func (o *BlogPost) GetEnableLayoutStylesheetsOk() (*bool, bool)`

GetEnableLayoutStylesheetsOk returns a tuple with the EnableLayoutStylesheets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLayoutStylesheets

`func (o *BlogPost) SetEnableLayoutStylesheets(v bool)`

SetEnableLayoutStylesheets sets EnableLayoutStylesheets field to given value.


### GetMetaDescription

`func (o *BlogPost) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *BlogPost) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *BlogPost) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.


### GetPublishImmediately

`func (o *BlogPost) GetPublishImmediately() bool`

GetPublishImmediately returns the PublishImmediately field if non-nil, zero value otherwise.

### GetPublishImmediatelyOk

`func (o *BlogPost) GetPublishImmediatelyOk() (*bool, bool)`

GetPublishImmediatelyOk returns a tuple with the PublishImmediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishImmediately

`func (o *BlogPost) SetPublishImmediately(v bool)`

SetPublishImmediately sets PublishImmediately field to given value.


### GetHeadHtml

`func (o *BlogPost) GetHeadHtml() string`

GetHeadHtml returns the HeadHtml field if non-nil, zero value otherwise.

### GetHeadHtmlOk

`func (o *BlogPost) GetHeadHtmlOk() (*string, bool)`

GetHeadHtmlOk returns a tuple with the HeadHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadHtml

`func (o *BlogPost) SetHeadHtml(v string)`

SetHeadHtml sets HeadHtml field to given value.


### GetFooterHtml

`func (o *BlogPost) GetFooterHtml() string`

GetFooterHtml returns the FooterHtml field if non-nil, zero value otherwise.

### GetFooterHtmlOk

`func (o *BlogPost) GetFooterHtmlOk() (*string, bool)`

GetFooterHtmlOk returns a tuple with the FooterHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterHtml

`func (o *BlogPost) SetFooterHtml(v string)`

SetFooterHtml sets FooterHtml field to given value.


### GetContentTypeCategory

`func (o *BlogPost) GetContentTypeCategory() string`

GetContentTypeCategory returns the ContentTypeCategory field if non-nil, zero value otherwise.

### GetContentTypeCategoryOk

`func (o *BlogPost) GetContentTypeCategoryOk() (*string, bool)`

GetContentTypeCategoryOk returns a tuple with the ContentTypeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypeCategory

`func (o *BlogPost) SetContentTypeCategory(v string)`

SetContentTypeCategory sets ContentTypeCategory field to given value.


### GetCurrentState

`func (o *BlogPost) GetCurrentState() string`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *BlogPost) GetCurrentStateOk() (*string, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *BlogPost) SetCurrentState(v string)`

SetCurrentState sets CurrentState field to given value.


### GetLinkRelCanonicalUrl

`func (o *BlogPost) GetLinkRelCanonicalUrl() string`

GetLinkRelCanonicalUrl returns the LinkRelCanonicalUrl field if non-nil, zero value otherwise.

### GetLinkRelCanonicalUrlOk

`func (o *BlogPost) GetLinkRelCanonicalUrlOk() (*string, bool)`

GetLinkRelCanonicalUrlOk returns a tuple with the LinkRelCanonicalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkRelCanonicalUrl

`func (o *BlogPost) SetLinkRelCanonicalUrl(v string)`

SetLinkRelCanonicalUrl sets LinkRelCanonicalUrl field to given value.


### GetFeaturedImage

`func (o *BlogPost) GetFeaturedImage() string`

GetFeaturedImage returns the FeaturedImage field if non-nil, zero value otherwise.

### GetFeaturedImageOk

`func (o *BlogPost) GetFeaturedImageOk() (*string, bool)`

GetFeaturedImageOk returns a tuple with the FeaturedImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedImage

`func (o *BlogPost) SetFeaturedImage(v string)`

SetFeaturedImage sets FeaturedImage field to given value.


### GetFeaturedImageAltText

`func (o *BlogPost) GetFeaturedImageAltText() string`

GetFeaturedImageAltText returns the FeaturedImageAltText field if non-nil, zero value otherwise.

### GetFeaturedImageAltTextOk

`func (o *BlogPost) GetFeaturedImageAltTextOk() (*string, bool)`

GetFeaturedImageAltTextOk returns a tuple with the FeaturedImageAltText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedImageAltText

`func (o *BlogPost) SetFeaturedImageAltText(v string)`

SetFeaturedImageAltText sets FeaturedImageAltText field to given value.


### GetPublicAccessRulesEnabled

`func (o *BlogPost) GetPublicAccessRulesEnabled() bool`

GetPublicAccessRulesEnabled returns the PublicAccessRulesEnabled field if non-nil, zero value otherwise.

### GetPublicAccessRulesEnabledOk

`func (o *BlogPost) GetPublicAccessRulesEnabledOk() (*bool, bool)`

GetPublicAccessRulesEnabledOk returns a tuple with the PublicAccessRulesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAccessRulesEnabled

`func (o *BlogPost) SetPublicAccessRulesEnabled(v bool)`

SetPublicAccessRulesEnabled sets PublicAccessRulesEnabled field to given value.


### GetPublicAccessRules

`func (o *BlogPost) GetPublicAccessRules() []map[string]interface{}`

GetPublicAccessRules returns the PublicAccessRules field if non-nil, zero value otherwise.

### GetPublicAccessRulesOk

`func (o *BlogPost) GetPublicAccessRulesOk() (*[]map[string]interface{}, bool)`

GetPublicAccessRulesOk returns a tuple with the PublicAccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAccessRules

`func (o *BlogPost) SetPublicAccessRules(v []map[string]interface{})`

SetPublicAccessRules sets PublicAccessRules field to given value.


### GetLayoutSections

`func (o *BlogPost) GetLayoutSections() map[string]LayoutSection`

GetLayoutSections returns the LayoutSections field if non-nil, zero value otherwise.

### GetLayoutSectionsOk

`func (o *BlogPost) GetLayoutSectionsOk() (*map[string]LayoutSection, bool)`

GetLayoutSectionsOk returns a tuple with the LayoutSections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutSections

`func (o *BlogPost) SetLayoutSections(v map[string]LayoutSection)`

SetLayoutSections sets LayoutSections field to given value.


### GetThemeSettingsValues

`func (o *BlogPost) GetThemeSettingsValues() map[string]map[string]interface{}`

GetThemeSettingsValues returns the ThemeSettingsValues field if non-nil, zero value otherwise.

### GetThemeSettingsValuesOk

`func (o *BlogPost) GetThemeSettingsValuesOk() (*map[string]map[string]interface{}, bool)`

GetThemeSettingsValuesOk returns a tuple with the ThemeSettingsValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeSettingsValues

`func (o *BlogPost) SetThemeSettingsValues(v map[string]map[string]interface{})`

SetThemeSettingsValues sets ThemeSettingsValues field to given value.


### GetUrl

`func (o *BlogPost) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BlogPost) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BlogPost) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPublishDate

`func (o *BlogPost) GetPublishDate() time.Time`

GetPublishDate returns the PublishDate field if non-nil, zero value otherwise.

### GetPublishDateOk

`func (o *BlogPost) GetPublishDateOk() (*time.Time, bool)`

GetPublishDateOk returns a tuple with the PublishDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishDate

`func (o *BlogPost) SetPublishDate(v time.Time)`

SetPublishDate sets PublishDate field to given value.


### GetDeletedAt

`func (o *BlogPost) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BlogPost) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BlogPost) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### GetCreatedAt

`func (o *BlogPost) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BlogPost) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BlogPost) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetPublished

`func (o *BlogPost) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *BlogPost) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *BlogPost) SetPublished(v bool)`

SetPublished sets Published field to given value.


### GetUpdatedAt

`func (o *BlogPost) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BlogPost) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BlogPost) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


