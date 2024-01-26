# \EventsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventsV3EventsGetPage**](EventsApi.md#GetEventsV3EventsGetPage) | **Get** /events/v3/events/ | 



## GetEventsV3EventsGetPage

> CollectionResponseExternalUnifiedEvent GetEventsV3EventsGetPage(ctx).ObjectType(objectType).EventType(eventType).OccurredAfter(occurredAfter).OccurredBefore(occurredBefore).ObjectId(objectId).IndexTableName(indexTableName).IndexSpecificMetadata(indexSpecificMetadata).After(after).Before(before).Limit(limit).Sort(sort).ObjectPropertyPropname(objectPropertyPropname).PropertyPropname(propertyPropname).Id(id).Execute()



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
    objectType := "objectType_example" // string |  (optional)
    eventType := "eventType_example" // string |  (optional)
    occurredAfter := time.Now() // time.Time |  (optional)
    occurredBefore := time.Now() // time.Time |  (optional)
    objectId := int64(789) // int64 |  (optional)
    indexTableName := "indexTableName_example" // string |  (optional)
    indexSpecificMetadata := "indexSpecificMetadata_example" // string |  (optional)
    after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
    before := "before_example" // string |  (optional)
    limit := int32(56) // int32 | The maximum number of results to display per page. (optional)
    sort := []string{"Inner_example"} // []string |  (optional)
    objectPropertyPropname := map[string]interface{}{ ... } // map[string]interface{} |  (optional)
    propertyPropname := map[string]interface{}{ ... } // map[string]interface{} |  (optional)
    id := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetEventsV3EventsGetPage(context.Background()).ObjectType(objectType).EventType(eventType).OccurredAfter(occurredAfter).OccurredBefore(occurredBefore).ObjectId(objectId).IndexTableName(indexTableName).IndexSpecificMetadata(indexSpecificMetadata).After(after).Before(before).Limit(limit).Sort(sort).ObjectPropertyPropname(objectPropertyPropname).PropertyPropname(propertyPropname).Id(id).Execute()
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
 **objectType** | **string** |  | 
 **eventType** | **string** |  | 
 **occurredAfter** | **time.Time** |  | 
 **occurredBefore** | **time.Time** |  | 
 **objectId** | **int64** |  | 
 **indexTableName** | **string** |  | 
 **indexSpecificMetadata** | **string** |  | 
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **before** | **string** |  | 
 **limit** | **int32** | The maximum number of results to display per page. | 
 **sort** | **[]string** |  | 
 **objectPropertyPropname** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **propertyPropname** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **id** | **[]string** |  | 

### Return type

[**CollectionResponseExternalUnifiedEvent**](CollectionResponseExternalUnifiedEvent.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

