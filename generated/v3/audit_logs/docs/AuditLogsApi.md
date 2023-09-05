# \AuditLogsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPage**](AuditLogsApi.md#GetPage) | **Get** /cms/v3/audit-logs/ | Query audit logs



## GetPage

> CollectionResponsePublicAuditLog GetPage(ctx).UserId(userId).EventType(eventType).ObjectType(objectType).ObjectId(objectId).After(after).Before(before).Limit(limit).Sort(sort).Execute()

Query audit logs



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
    userId := []string{"Inner_example"} // []string | Comma separated list of user ids to filter by. (optional)
    eventType := []string{"Inner_example"} // []string | Comma separated list of event types to filter by (CREATED, UPDATED, PUBLISHED, DELETED, UNPUBLISHED). (optional)
    objectType := []string{"Inner_example"} // []string | Comma separated list of object types to filter by (BLOG, LANDING_PAGE, DOMAIN, HUBDB_TABLE etc.) (optional)
    objectId := []string{"Inner_example"} // []string | Comma separated list of object ids to filter by. (optional)
    after := "after_example" // string | Timestamp after which audit logs will be returned (optional)
    before := "before_example" // string | Timestamp before which audit logs will be returned (optional)
    limit := int32(56) // int32 | The number of logs to return. (optional)
    sort := []string{"Inner_example"} // []string | The sort direction for the audit logs. (Can only sort by timestamp). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuditLogsApi.GetPage(context.Background()).UserId(userId).EventType(eventType).ObjectType(objectType).ObjectId(objectId).After(after).Before(before).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsApi.GetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPage`: CollectionResponsePublicAuditLog
    fmt.Fprintf(os.Stdout, "Response from `AuditLogsApi.GetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **[]string** | Comma separated list of user ids to filter by. | 
 **eventType** | **[]string** | Comma separated list of event types to filter by (CREATED, UPDATED, PUBLISHED, DELETED, UNPUBLISHED). | 
 **objectType** | **[]string** | Comma separated list of object types to filter by (BLOG, LANDING_PAGE, DOMAIN, HUBDB_TABLE etc.) | 
 **objectId** | **[]string** | Comma separated list of object ids to filter by. | 
 **after** | **string** | Timestamp after which audit logs will be returned | 
 **before** | **string** | Timestamp before which audit logs will be returned | 
 **limit** | **int32** | The number of logs to return. | 
 **sort** | **[]string** | The sort direction for the audit logs. (Can only sort by timestamp). | 

### Return type

[**CollectionResponsePublicAuditLog**](CollectionResponsePublicAuditLog.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

