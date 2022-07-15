# \BasicApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMarketingV3MarketingEventsEventsExternalEventIdArchive**](BasicApi.md#DeleteMarketingV3MarketingEventsEventsExternalEventIdArchive) | **Delete** /marketing/v3/marketing-events/events/{externalEventId} | Delete a marketing event
[**GetMarketingV3MarketingEventsEventsExternalEventIdGetById**](BasicApi.md#GetMarketingV3MarketingEventsEventsExternalEventIdGetById) | **Get** /marketing/v3/marketing-events/events/{externalEventId} | Get a marketing event
[**PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate**](BasicApi.md#PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate) | **Patch** /marketing/v3/marketing-events/events/{externalEventId} | Update a marketing event
[**PostMarketingV3MarketingEventsEventsCreate**](BasicApi.md#PostMarketingV3MarketingEventsEventsCreate) | **Post** /marketing/v3/marketing-events/events | Create a marketing event
[**PostMarketingV3MarketingEventsEventsExternalEventIdCancelDoCancel**](BasicApi.md#PostMarketingV3MarketingEventsEventsExternalEventIdCancelDoCancel) | **Post** /marketing/v3/marketing-events/events/{externalEventId}/cancel | Mark a marketing event as cancelled
[**PutMarketingV3MarketingEventsEventsExternalEventIdReplace**](BasicApi.md#PutMarketingV3MarketingEventsEventsExternalEventIdReplace) | **Put** /marketing/v3/marketing-events/events/{externalEventId} | Create or update a marketing event



## DeleteMarketingV3MarketingEventsEventsExternalEventIdArchive

> DeleteMarketingV3MarketingEventsEventsExternalEventIdArchive(ctx, externalEventId).ExternalAccountId(externalAccountId).Execute()

Delete a marketing event



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
    externalEventId := "externalEventId_example" // string | The id of the marketing event to delete
    externalAccountId := "externalAccountId_example" // string | The account id associated with the marketing event

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BasicApi.DeleteMarketingV3MarketingEventsEventsExternalEventIdArchive(context.Background(), externalEventId).ExternalAccountId(externalAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BasicApi.DeleteMarketingV3MarketingEventsEventsExternalEventIdArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketingV3MarketingEventsEventsExternalEventIdArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** | The account id associated with the marketing event | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3MarketingEventsEventsExternalEventIdGetById

> MarketingEventPublicReadResponse GetMarketingV3MarketingEventsEventsExternalEventIdGetById(ctx, externalEventId).ExternalAccountId(externalAccountId).Execute()

Get a marketing event



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
    externalEventId := "externalEventId_example" // string | The id of the marketing event to return
    externalAccountId := "externalAccountId_example" // string | The account id associated with the marketing event

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BasicApi.GetMarketingV3MarketingEventsEventsExternalEventIdGetById(context.Background(), externalEventId).ExternalAccountId(externalAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BasicApi.GetMarketingV3MarketingEventsEventsExternalEventIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketingV3MarketingEventsEventsExternalEventIdGetById`: MarketingEventPublicReadResponse
    fmt.Fprintf(os.Stdout, "Response from `BasicApi.GetMarketingV3MarketingEventsEventsExternalEventIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3MarketingEventsEventsExternalEventIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** | The account id associated with the marketing event | 

### Return type

[**MarketingEventPublicReadResponse**](MarketingEventPublicReadResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate

> MarketingEventPublicDefaultResponse PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate(ctx, externalEventId).ExternalAccountId(externalAccountId).MarketingEventUpdateRequestParams(marketingEventUpdateRequestParams).Execute()

Update a marketing event



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
    externalEventId := "externalEventId_example" // string | The id of the marketing event to update
    externalAccountId := "externalAccountId_example" // string | The account id associated with the marketing event
    marketingEventUpdateRequestParams := *openapiclient.NewMarketingEventUpdateRequestParams() // MarketingEventUpdateRequestParams | The details of the marketing event to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BasicApi.PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate(context.Background(), externalEventId).ExternalAccountId(externalAccountId).MarketingEventUpdateRequestParams(marketingEventUpdateRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BasicApi.PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate`: MarketingEventPublicDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `BasicApi.PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMarketingV3MarketingEventsEventsExternalEventIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** | The account id associated with the marketing event | 
 **marketingEventUpdateRequestParams** | [**MarketingEventUpdateRequestParams**](MarketingEventUpdateRequestParams.md) | The details of the marketing event to update | 

### Return type

[**MarketingEventPublicDefaultResponse**](MarketingEventPublicDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsEventsCreate

> MarketingEventDefaultResponse PostMarketingV3MarketingEventsEventsCreate(ctx).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()

Create a marketing event



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
    marketingEventCreateRequestParams := *openapiclient.NewMarketingEventCreateRequestParams("EventName_example", "EventOrganizer_example", "ExternalAccountId_example", "ExternalEventId_example") // MarketingEventCreateRequestParams | The details of the marketing event to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BasicApi.PostMarketingV3MarketingEventsEventsCreate(context.Background()).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BasicApi.PostMarketingV3MarketingEventsEventsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMarketingV3MarketingEventsEventsCreate`: MarketingEventDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `BasicApi.PostMarketingV3MarketingEventsEventsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketingEventCreateRequestParams** | [**MarketingEventCreateRequestParams**](MarketingEventCreateRequestParams.md) | The details of the marketing event to create | 

### Return type

[**MarketingEventDefaultResponse**](MarketingEventDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsEventsExternalEventIdCancelDoCancel

> MarketingEventDefaultResponse PostMarketingV3MarketingEventsEventsExternalEventIdCancelDoCancel(ctx, externalEventId).ExternalAccountId(externalAccountId).Execute()

Mark a marketing event as cancelled



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
    externalEventId := "externalEventId_example" // string | The id of the marketing event to mark as cancelled
    externalAccountId := "externalAccountId_example" // string | The account id associated with the marketing event

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BasicApi.PostMarketingV3MarketingEventsEventsExternalEventIdCancelDoCancel(context.Background(), externalEventId).ExternalAccountId(externalAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BasicApi.PostMarketingV3MarketingEventsEventsExternalEventIdCancelDoCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMarketingV3MarketingEventsEventsExternalEventIdCancelDoCancel`: MarketingEventDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `BasicApi.PostMarketingV3MarketingEventsEventsExternalEventIdCancelDoCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event to mark as cancelled | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsExternalEventIdCancelDoCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** | The account id associated with the marketing event | 

### Return type

[**MarketingEventDefaultResponse**](MarketingEventDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMarketingV3MarketingEventsEventsExternalEventIdReplace

> MarketingEventPublicDefaultResponse PutMarketingV3MarketingEventsEventsExternalEventIdReplace(ctx, externalEventId).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()

Create or update a marketing event



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
    externalEventId := "externalEventId_example" // string | The id of the marketing event to upsert
    marketingEventCreateRequestParams := *openapiclient.NewMarketingEventCreateRequestParams("EventName_example", "EventOrganizer_example", "ExternalAccountId_example", "ExternalEventId_example") // MarketingEventCreateRequestParams | The details of the marketing event to upsert

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BasicApi.PutMarketingV3MarketingEventsEventsExternalEventIdReplace(context.Background(), externalEventId).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BasicApi.PutMarketingV3MarketingEventsEventsExternalEventIdReplace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutMarketingV3MarketingEventsEventsExternalEventIdReplace`: MarketingEventPublicDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `BasicApi.PutMarketingV3MarketingEventsEventsExternalEventIdReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event to upsert | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMarketingV3MarketingEventsEventsExternalEventIdReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketingEventCreateRequestParams** | [**MarketingEventCreateRequestParams**](MarketingEventCreateRequestParams.md) | The details of the marketing event to upsert | 

### Return type

[**MarketingEventPublicDefaultResponse**](MarketingEventPublicDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

