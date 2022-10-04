# \BatchApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMarketingV3MarketingEventsEventsDeleteArchive**](BatchApi.md#PostMarketingV3MarketingEventsEventsDeleteArchive) | **Post** /marketing/v3/marketing-events/events/delete | Delete multiple marketing events
[**PostMarketingV3MarketingEventsEventsUpsertDoUpsert**](BatchApi.md#PostMarketingV3MarketingEventsEventsUpsertDoUpsert) | **Post** /marketing/v3/marketing-events/events/upsert | Create or update multiple marketing events



## PostMarketingV3MarketingEventsEventsDeleteArchive

> Error PostMarketingV3MarketingEventsEventsDeleteArchive(ctx).BatchInputMarketingEventExternalUniqueIdentifier(batchInputMarketingEventExternalUniqueIdentifier).Execute()

Delete multiple marketing events



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
    batchInputMarketingEventExternalUniqueIdentifier := *openapiclient.NewBatchInputMarketingEventExternalUniqueIdentifier([]openapiclient.MarketingEventExternalUniqueIdentifier{*openapiclient.NewMarketingEventExternalUniqueIdentifier(int32(123), "ExternalAccountId_example", "ExternalEventId_example")}) // BatchInputMarketingEventExternalUniqueIdentifier | The details of the marketing events to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostMarketingV3MarketingEventsEventsDeleteArchive(context.Background()).BatchInputMarketingEventExternalUniqueIdentifier(batchInputMarketingEventExternalUniqueIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostMarketingV3MarketingEventsEventsDeleteArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMarketingV3MarketingEventsEventsDeleteArchive`: Error
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.PostMarketingV3MarketingEventsEventsDeleteArchive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsDeleteArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputMarketingEventExternalUniqueIdentifier** | [**BatchInputMarketingEventExternalUniqueIdentifier**](BatchInputMarketingEventExternalUniqueIdentifier.md) | The details of the marketing events to delete | 

### Return type

[**Error**](Error.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsEventsUpsertDoUpsert

> BatchResponseMarketingEventPublicDefaultResponse PostMarketingV3MarketingEventsEventsUpsertDoUpsert(ctx).BatchInputMarketingEventCreateRequestParams(batchInputMarketingEventCreateRequestParams).Execute()

Create or update multiple marketing events



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
    batchInputMarketingEventCreateRequestParams := *openapiclient.NewBatchInputMarketingEventCreateRequestParams([]openapiclient.MarketingEventCreateRequestParams{*openapiclient.NewMarketingEventCreateRequestParams("EventName_example", "EventOrganizer_example", "ExternalAccountId_example", "ExternalEventId_example")}) // BatchInputMarketingEventCreateRequestParams | The details of the marketing events to upsert

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostMarketingV3MarketingEventsEventsUpsertDoUpsert(context.Background()).BatchInputMarketingEventCreateRequestParams(batchInputMarketingEventCreateRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostMarketingV3MarketingEventsEventsUpsertDoUpsert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMarketingV3MarketingEventsEventsUpsertDoUpsert`: BatchResponseMarketingEventPublicDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.PostMarketingV3MarketingEventsEventsUpsertDoUpsert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsUpsertDoUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputMarketingEventCreateRequestParams** | [**BatchInputMarketingEventCreateRequestParams**](BatchInputMarketingEventCreateRequestParams.md) | The details of the marketing events to upsert | 

### Return type

[**BatchResponseMarketingEventPublicDefaultResponse**](BatchResponseMarketingEventPublicDefaultResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

