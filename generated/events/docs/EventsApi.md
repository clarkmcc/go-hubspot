# \EventsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventsV3EventsGetPage**](EventsApi.md#GetEventsV3EventsGetPage) | **Get** /events/v3/events | Returns a collection of events matching a query.



## GetEventsV3EventsGetPage

> CollectionResponseExternalUnifiedEvent GetEventsV3EventsGetPage(ctx).OccurredAfter(occurredAfter).OccurredBefore(occurredBefore).ObjectType(objectType).ObjectId(objectId).EventType(eventType).After(after).Before(before).Limit(limit).Sort(sort).Execute()

Returns a collection of events matching a query.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    occurredAfter := time.Now() // time.Time | The starting time as an ISO 8601 timestamp. (optional)
    occurredBefore := time.Now() // time.Time | The ending time as an ISO 8601 timestamp. (optional)
    objectType := "objectType_example" // string | The type of object being selected. Valid values are hubspot named object types (e.g. `contact`). (optional)
    objectId := int64(789) // int64 | The id of the selected object. If not present, then the `objectProperty` parameter is required. (optional)
    eventType := "eventType_example" // string | Limits the response to the specified event type.  For example `&eventType=e_visited_page` returns only `e_visited_page` events.  If not present all event types are returned. (optional)
    after := "after_example" // string | An additional parameter that may be used to get the next `limit` set of results. (optional)
    before := "before_example" // string |  (optional)
    limit := int32(56) // int32 | The maximum number of events to return, defaults to 20. (optional)
    sort := []string{"Inner_example"} // []string | Selects the sort field and order. Defaults to ascending, prefix with `-` for descending order. `occurredAt` is the only field supported for sorting. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetEventsV3EventsGetPage(context.Background()).OccurredAfter(occurredAfter).OccurredBefore(occurredBefore).ObjectType(objectType).ObjectId(objectId).EventType(eventType).After(after).Before(before).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEventsV3EventsGetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventsV3EventsGetPage`: CollectionResponseExternalUnifiedEvent
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetEventsV3EventsGetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsV3EventsGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **occurredAfter** | **time.Time** | The starting time as an ISO 8601 timestamp. | 
 **occurredBefore** | **time.Time** | The ending time as an ISO 8601 timestamp. | 
 **objectType** | **string** | The type of object being selected. Valid values are hubspot named object types (e.g. &#x60;contact&#x60;). | 
 **objectId** | **int64** | The id of the selected object. If not present, then the &#x60;objectProperty&#x60; parameter is required. | 
 **eventType** | **string** | Limits the response to the specified event type.  For example &#x60;&amp;eventType&#x3D;e_visited_page&#x60; returns only &#x60;e_visited_page&#x60; events.  If not present all event types are returned. | 
 **after** | **string** | An additional parameter that may be used to get the next &#x60;limit&#x60; set of results. | 
 **before** | **string** |  | 
 **limit** | **int32** | The maximum number of events to return, defaults to 20. | 
 **sort** | **[]string** | Selects the sort field and order. Defaults to ascending, prefix with &#x60;-&#x60; for descending order. &#x60;occurredAt&#x60; is the only field supported for sorting. | 

### Return type

[**CollectionResponseExternalUnifiedEvent**](CollectionResponseExternalUnifiedEvent.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

