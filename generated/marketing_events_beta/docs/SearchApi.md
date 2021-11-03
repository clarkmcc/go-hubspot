# \SearchApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMarketingV3MarketingEventsEventsSearch**](SearchApi.md#GetMarketingV3MarketingEventsEventsSearch) | **Get** /marketing/v3/marketing-events/events/search | Search for marketing events



## GetMarketingV3MarketingEventsEventsSearch

> CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging GetMarketingV3MarketingEventsEventsSearch(ctx).Q(q).Execute()

Search for marketing events



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
    q := "q_example" // string | The partial event id to search for

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchApi.GetMarketingV3MarketingEventsEventsSearch(context.Background()).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetMarketingV3MarketingEventsEventsSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketingV3MarketingEventsEventsSearch`: CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetMarketingV3MarketingEventsEventsSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3MarketingEventsEventsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | The partial event id to search for | 

### Return type

[**CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging**](CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

