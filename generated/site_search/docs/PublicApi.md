# \PublicApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCmsV3SiteSearchIndexedDataContentIdGetById**](PublicApi.md#GetCmsV3SiteSearchIndexedDataContentIdGetById) | **Get** /cms/v3/site-search/indexed-data/{contentId} | Get indexed properties.
[**GetCmsV3SiteSearchSearchSearch**](PublicApi.md#GetCmsV3SiteSearchSearchSearch) | **Get** /cms/v3/site-search/search | Search your site.



## GetCmsV3SiteSearchIndexedDataContentIdGetById

> IndexedData GetCmsV3SiteSearchIndexedDataContentIdGetById(ctx, contentId).Type_(type_).Execute()

Get indexed properties.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    contentId := "contentId_example" // string | ID of the target document when searching for indexed properties.
    type_ := "type__example" // string | The type of document. Can be one of `SITE_PAGE`, `BLOG_POST`, or `KNOWLEDGE_ARTICLE`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicApi.GetCmsV3SiteSearchIndexedDataContentIdGetById(context.Background(), contentId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.GetCmsV3SiteSearchIndexedDataContentIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3SiteSearchIndexedDataContentIdGetById`: IndexedData
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.GetCmsV3SiteSearchIndexedDataContentIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentId** | **string** | ID of the target document when searching for indexed properties. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3SiteSearchIndexedDataContentIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | The type of document. Can be one of &#x60;SITE_PAGE&#x60;, &#x60;BLOG_POST&#x60;, or &#x60;KNOWLEDGE_ARTICLE&#x60;. | 

### Return type

[**IndexedData**](IndexedData.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3SiteSearchSearchSearch

> PublicSearchResults GetCmsV3SiteSearchSearchSearch(ctx).Q(q).Limit(limit).Offset(offset).Language(language).MatchPrefix(matchPrefix).Autocomplete(autocomplete).PopularityBoost(popularityBoost).BoostLimit(boostLimit).MinScore(minScore).BoostRecent(boostRecent).TableId(tableId).HubdbQuery(hubdbQuery).Domain(domain).Type_(type_).PathPrefix(pathPrefix).Property(property).Length(length).GroupId(groupId).Execute()

Search your site.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    q := "q_example" // string | The term to search for. (optional)
    limit := int32(56) // int32 | Specifies the number of results to be returned in a single response. Defaults to `10`. Maximum value is `100`. (optional)
    offset := int32(56) // int32 | Used to page through the results. If there are more results than specified by the `limit` parameter, you will need to use the value of offset returned in the previous request to get the next set of results. (optional)
    language := "language_example" // string | Specifies the language of content to be searched. This value must be a valid [ISO 639-1 language code](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) (e.g. `es` for Spanish) (optional)
    matchPrefix := true // bool | Inverts the behavior of the pathPrefix filter when set to `false`. Defaults to `true`. (optional)
    autocomplete := true // bool | Specifies whether or not you are showing autocomplete results. Defaults to false. (optional)
    popularityBoost := float32(8.14) // float32 | Specifies how strongly a result is boosted based on its view count. Defaults to 1.0. (optional)
    boostLimit := float32(8.14) // float32 | Specifies the maximum amount a result will be boosted based on its view count. Defaults to 5.0. Read more about elasticsearch boosting [here](https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-boost.html). (optional)
    minScore := float32(8.14) // float32 | Specifies the minimum search score threshold for returned results. This value is intentionally set low by default in order to return many results. Increase this for higher precision, but less recall. (optional)
    boostRecent := "boostRecent_example" // string | Specifies a relative time window where scores of documents published outside this time window decay. This can only be used for blog posts. For example, boostRecent=10d will boost documents published within the last 10 days. Supported timeunits are ms (milliseconds), s (seconds), m (minutes), h (hours), d (days). (optional)
    tableId := int64(789) // int64 | Specifies a specific HubDB table to search. Only returns results from the specified table. Can be used in tandem with the `hubdbQuery` parameter to further filter results. (optional)
    hubdbQuery := "hubdbQuery_example" // string | Specify a HubDB query to further filter the search results. (optional)
    domain := []string{"Inner_example"} // []string | A domain to match search results for. Multiple domains can be provided with &. (optional)
    type_ := []string{"Type_example"} // []string | Specifies the type of content to search. Can be one or more of SITE_PAGE, LANDING_PAGE, BLOG_POST, LISTING_PAGE, and KNOWLEDGE_ARTICLE. Defaults to all content types except LANDING_PAGE and KNOWLEDGE_ARTICLE (optional)
    pathPrefix := []string{"Inner_example"} // []string | Specifies a path prefix to filter search results. Will only return results with URL paths that start with the specified parameter. Can be used multiple times. (optional)
    property := []string{"Inner_example"} // []string | Specifies which properties to include in the search. Options include `title`, `description`, and `html`. All properties will be searched by default. (optional)
    length := "length_example" // string | Specifies the length of the search results. Can be set to `LONG` or `SHORT`. `SHORT` will return the first 128 characters of the content's meta description. `LONG` will build a more detailed content snippet based on the html/content of the page. (optional)
    groupId := []int64{int64(123)} // []int64 | Specifies which blog(s) to be searched by blog ID. Can be used multiple times to search more than one blog. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicApi.GetCmsV3SiteSearchSearchSearch(context.Background()).Q(q).Limit(limit).Offset(offset).Language(language).MatchPrefix(matchPrefix).Autocomplete(autocomplete).PopularityBoost(popularityBoost).BoostLimit(boostLimit).MinScore(minScore).BoostRecent(boostRecent).TableId(tableId).HubdbQuery(hubdbQuery).Domain(domain).Type_(type_).PathPrefix(pathPrefix).Property(property).Length(length).GroupId(groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.GetCmsV3SiteSearchSearchSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3SiteSearchSearchSearch`: PublicSearchResults
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.GetCmsV3SiteSearchSearchSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3SiteSearchSearchSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | The term to search for. | 
 **limit** | **int32** | Specifies the number of results to be returned in a single response. Defaults to &#x60;10&#x60;. Maximum value is &#x60;100&#x60;. | 
 **offset** | **int32** | Used to page through the results. If there are more results than specified by the &#x60;limit&#x60; parameter, you will need to use the value of offset returned in the previous request to get the next set of results. | 
 **language** | **string** | Specifies the language of content to be searched. This value must be a valid [ISO 639-1 language code](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) (e.g. &#x60;es&#x60; for Spanish) | 
 **matchPrefix** | **bool** | Inverts the behavior of the pathPrefix filter when set to &#x60;false&#x60;. Defaults to &#x60;true&#x60;. | 
 **autocomplete** | **bool** | Specifies whether or not you are showing autocomplete results. Defaults to false. | 
 **popularityBoost** | **float32** | Specifies how strongly a result is boosted based on its view count. Defaults to 1.0. | 
 **boostLimit** | **float32** | Specifies the maximum amount a result will be boosted based on its view count. Defaults to 5.0. Read more about elasticsearch boosting [here](https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-boost.html). | 
 **minScore** | **float32** | Specifies the minimum search score threshold for returned results. This value is intentionally set low by default in order to return many results. Increase this for higher precision, but less recall. | 
 **boostRecent** | **string** | Specifies a relative time window where scores of documents published outside this time window decay. This can only be used for blog posts. For example, boostRecent&#x3D;10d will boost documents published within the last 10 days. Supported timeunits are ms (milliseconds), s (seconds), m (minutes), h (hours), d (days). | 
 **tableId** | **int64** | Specifies a specific HubDB table to search. Only returns results from the specified table. Can be used in tandem with the &#x60;hubdbQuery&#x60; parameter to further filter results. | 
 **hubdbQuery** | **string** | Specify a HubDB query to further filter the search results. | 
 **domain** | **[]string** | A domain to match search results for. Multiple domains can be provided with &amp;. | 
 **type_** | **[]string** | Specifies the type of content to search. Can be one or more of SITE_PAGE, LANDING_PAGE, BLOG_POST, LISTING_PAGE, and KNOWLEDGE_ARTICLE. Defaults to all content types except LANDING_PAGE and KNOWLEDGE_ARTICLE | 
 **pathPrefix** | **[]string** | Specifies a path prefix to filter search results. Will only return results with URL paths that start with the specified parameter. Can be used multiple times. | 
 **property** | **[]string** | Specifies which properties to include in the search. Options include &#x60;title&#x60;, &#x60;description&#x60;, and &#x60;html&#x60;. All properties will be searched by default. | 
 **length** | **string** | Specifies the length of the search results. Can be set to &#x60;LONG&#x60; or &#x60;SHORT&#x60;. &#x60;SHORT&#x60; will return the first 128 characters of the content&#39;s meta description. &#x60;LONG&#x60; will build a more detailed content snippet based on the html/content of the page. | 
 **groupId** | **[]int64** | Specifies which blog(s) to be searched by blog ID. Can be used multiple times to search more than one blog. | 

### Return type

[**PublicSearchResults**](PublicSearchResults.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

