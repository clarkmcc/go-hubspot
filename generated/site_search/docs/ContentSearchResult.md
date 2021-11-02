# ContentSearchResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the content. | [default to null]
**Score** | **float64** | The matching score of the document. | [default to null]
**Type_** | **string** | The type of document. Can be &#x60;SITE_PAGE&#x60;, &#x60;LANDING_PAGE&#x60;, &#x60;BLOG_POST&#x60;, &#x60;LISTING_PAGE&#x60;, or &#x60;KNOWLEDGE_ARTICLE&#x60;. | [default to null]
**Domain** | **string** | The domain the document is hosted on. | [default to null]
**Url** | **string** | The url of the document. | [default to null]
**FeaturedImageUrl** | **string** | URL of the featured image. | [optional] [default to null]
**Language** | **string** | The document&#x27;s language. | [optional] [default to null]
**Title** | **string** | The title of the returned document. | [optional] [default to null]
**Description** | **string** | The result&#x27;s description. The content will be determined by the value of &#x60;length&#x60; in the request. | [optional] [default to null]
**Category** | **string** | For knowledge articles, the category of the article. | [optional] [default to null]
**Subcategory** | **string** | For knowledge articles, the subcategory of the article. | [optional] [default to null]
**AuthorFullName** | **string** | Name of the author. | [optional] [default to null]
**Tags** | **[]string** | If a blog post, the tags associated with it. | [optional] [default to null]
**TableId** | **int64** | If a dynamic page, the ID of the HubDB table. | [optional] [default to null]
**RowId** | **int64** | If a dynamic page, the row ID in the HubDB table. | [optional] [default to null]
**PublishedDate** | **int64** | The date the content was published. | [optional] [default to null]
**CombinedId** | **string** | The ID of the document in HubSpot. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

