# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getcmsv3siteSearchindexedDatacontentIdGetById**](PublicApi.md#Getcmsv3siteSearchindexedDatacontentIdGetById) | **Get** /cms/v3/site-search/indexed-data/{contentId} | Get indexed properties.
[**Getcmsv3siteSearchsearchSearch**](PublicApi.md#Getcmsv3siteSearchsearchSearch) | **Get** /cms/v3/site-search/search | Search your site.

# **Getcmsv3siteSearchindexedDatacontentIdGetById**
> IndexedData Getcmsv3siteSearchindexedDatacontentIdGetById(ctx, contentId, optional)
Get indexed properties.

For a given account and document ID (page ID, blog post ID, HubDB row ID, etc.), return all indexed data for that document. This is useful when debugging why a particular document is not returned from a custom search.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentId** | **string**| ID of the target document when searching for indexed properties. | 
 **optional** | ***PublicApiGetcmsv3siteSearchindexedDatacontentIdGetByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicApiGetcmsv3siteSearchindexedDatacontentIdGetByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| The type of document. Can be one of &#x60;SITE_PAGE&#x60;, &#x60;BLOG_POST&#x60;, or &#x60;KNOWLEDGE_ARTICLE&#x60;. | 

### Return type

[**IndexedData**](IndexedData.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcmsv3siteSearchsearchSearch**
> PublicSearchResults Getcmsv3siteSearchsearchSearch(ctx, optional)
Search your site.

Returns any website content matching the given search criteria for a given HubSpot account. Searches can be filtered by content type, domain, or URL path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PublicApiGetcmsv3siteSearchsearchSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicApiGetcmsv3siteSearchsearchSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| The term to search for. | 
 **limit** | **optional.Int32**| Specifies the number of results to be returned in a single response. Defaults to &#x60;10&#x60;. Maximum value is &#x60;100&#x60;. | 
 **offset** | **optional.Int32**| Used to page through the results. If there are more results than specified by the &#x60;limit&#x60; parameter, you will need to use the value of offset returned in the previous request to get the next set of results. | 
 **language** | **optional.String**| Specifies the language of content to be searched. This value must be a valid [ISO 639-1 language code](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) (e.g. &#x60;es&#x60; for Spanish) | 
 **matchPrefix** | **optional.Bool**| Inverts the behavior of the pathPrefix filter when set to &#x60;false&#x60;. Defaults to &#x60;true&#x60;. | 
 **autocomplete** | **optional.Bool**| Specifies whether or not you are showing autocomplete results. Defaults to false. | 
 **popularityBoost** | **optional.Float64**| Specifies how strongly a result is boosted based on its view count. Defaults to 1.0. | 
 **boostLimit** | **optional.Float64**| Specifies the maximum amount a result will be boosted based on its view count. Defaults to 5.0. Read more about elasticsearch boosting [here](https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-boost.html). | 
 **minScore** | **optional.Float64**| Specifies the minimum search score threshold for returned results. This value is intentionally set low by default in order to return many results. Increase this for higher precision, but less recall. | 
 **boostRecent** | **optional.String**| Specifies a relative time window where scores of documents published outside this time window decay. This can only be used for blog posts. For example, boostRecent&#x3D;10d will boost documents published within the last 10 days. Supported timeunits are ms (milliseconds), s (seconds), m (minutes), h (hours), d (days). | 
 **tableId** | **optional.Int64**| Specifies a specific HubDB table to search. Only returns results from the specified table. Can be used in tandem with the &#x60;hubdbQuery&#x60; parameter to further filter results. | 
 **hubdbQuery** | **optional.String**| Specify a HubDB query to further filter the search results. | 
 **domain** | [**optional.Interface of []string**](string.md)| A domain to match search results for. Multiple domains can be provided with &amp;. | 
 **type_** | [**optional.Interface of []string**](string.md)| Specifies the type of content to search. Can be one or more of SITE_PAGE, LANDING_PAGE, BLOG_POST, LISTING_PAGE, and KNOWLEDGE_ARTICLE. Defaults to all content types except LANDING_PAGE and KNOWLEDGE_ARTICLE | 
 **pathPrefix** | [**optional.Interface of []string**](string.md)| Specifies a path prefix to filter search results. Will only return results with URL paths that start with the specified parameter. Can be used multiple times. | 
 **property** | [**optional.Interface of []string**](string.md)| Specifies which properties to include in the search. Options include &#x60;title&#x60;, &#x60;description&#x60;, and &#x60;html&#x60;. All properties will be searched by default. | 
 **length** | **optional.String**| Specifies the length of the search results. Can be set to &#x60;LONG&#x60; or &#x60;SHORT&#x60;. &#x60;SHORT&#x60; will return the first 128 characters of the content&#x27;s meta description. &#x60;LONG&#x60; will build a more detailed content snippet based on the html/content of the page. | 
 **groupId** | [**optional.Interface of []int64**](int64.md)| Specifies which blog(s) to be searched by blog ID. Can be used multiple times to search more than one blog. | 

### Return type

[**PublicSearchResults**](PublicSearchResults.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

