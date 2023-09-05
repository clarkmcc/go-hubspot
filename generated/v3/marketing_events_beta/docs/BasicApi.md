# \BasicApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](BasicApi.md#Create) | **Post** /marketing/v3/marketing-events/events | Create a marketing event
[**ExternalArchive**](BasicApi.md#ExternalArchive) | **Delete** /marketing/v3/marketing-events/events/{externalEventId} | Delete a marketing event
[**ExternalCancel**](BasicApi.md#ExternalCancel) | **Post** /marketing/v3/marketing-events/events/{externalEventId}/cancel | Mark a marketing event as cancelled
[**ExternalGetByID**](BasicApi.md#ExternalGetByID) | **Get** /marketing/v3/marketing-events/events/{externalEventId} | Get a marketing event
[**ExternalReplace**](BasicApi.md#ExternalReplace) | **Put** /marketing/v3/marketing-events/events/{externalEventId} | Create or update a marketing event
[**ExternalUpdate**](BasicApi.md#ExternalUpdate) | **Patch** /marketing/v3/marketing-events/events/{externalEventId} | Update a marketing event



## Create

> MarketingEventDefaultResponse Create(ctx).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()

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
    marketingEventCreateRequestParams := *openapiclient.NewMarketingEventCreateRequestParams("ExternalAccountId_example", "EventOrganizer_example", "ExternalEventId_example", "EventName_example") // MarketingEventCreateRequestParams | The details of the marketing event to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BasicApi.Create(context.Background()).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BasicApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: MarketingEventDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `BasicApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketingEventCreateRequestParams** | [**MarketingEventCreateRequestParams**](MarketingEventCreateRequestParams.md) | The details of the marketing event to create | 

### Return type

[**MarketingEventDefaultResponse**](MarketingEventDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalArchive

> ExternalArchive(ctx, externalEventId).ExternalAccountId(externalAccountId).Execute()

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
    resp, r, err := apiClient.BasicApi.ExternalArchive(context.Background(), externalEventId).ExternalAccountId(externalAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BasicApi.ExternalArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExternalArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** | The account id associated with the marketing event | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalCancel

> MarketingEventDefaultResponse ExternalCancel(ctx, externalEventId).ExternalAccountId(externalAccountId).Execute()

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
    resp, r, err := apiClient.BasicApi.ExternalCancel(context.Background(), externalEventId).ExternalAccountId(externalAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BasicApi.ExternalCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExternalCancel`: MarketingEventDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `BasicApi.ExternalCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event to mark as cancelled | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** | The account id associated with the marketing event | 

### Return type

[**MarketingEventDefaultResponse**](MarketingEventDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalGetByID

> MarketingEventPublicReadResponse ExternalGetByID(ctx, externalEventId).ExternalAccountId(externalAccountId).Execute()

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
    resp, r, err := apiClient.BasicApi.ExternalGetByID(context.Background(), externalEventId).ExternalAccountId(externalAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BasicApi.ExternalGetByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExternalGetByID`: MarketingEventPublicReadResponse
    fmt.Fprintf(os.Stdout, "Response from `BasicApi.ExternalGetByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalGetByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** | The account id associated with the marketing event | 

### Return type

[**MarketingEventPublicReadResponse**](MarketingEventPublicReadResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalReplace

> MarketingEventPublicDefaultResponse ExternalReplace(ctx, externalEventId).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()

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
    marketingEventCreateRequestParams := *openapiclient.NewMarketingEventCreateRequestParams("ExternalAccountId_example", "EventOrganizer_example", "ExternalEventId_example", "EventName_example") // MarketingEventCreateRequestParams | The details of the marketing event to upsert

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BasicApi.ExternalReplace(context.Background(), externalEventId).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BasicApi.ExternalReplace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExternalReplace`: MarketingEventPublicDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `BasicApi.ExternalReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event to upsert | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketingEventCreateRequestParams** | [**MarketingEventCreateRequestParams**](MarketingEventCreateRequestParams.md) | The details of the marketing event to upsert | 

### Return type

[**MarketingEventPublicDefaultResponse**](MarketingEventPublicDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalUpdate

> MarketingEventPublicDefaultResponse ExternalUpdate(ctx, externalEventId).ExternalAccountId(externalAccountId).MarketingEventUpdateRequestParams(marketingEventUpdateRequestParams).Execute()

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
    resp, r, err := apiClient.BasicApi.ExternalUpdate(context.Background(), externalEventId).ExternalAccountId(externalAccountId).MarketingEventUpdateRequestParams(marketingEventUpdateRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BasicApi.ExternalUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExternalUpdate`: MarketingEventPublicDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `BasicApi.ExternalUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** | The account id associated with the marketing event | 
 **marketingEventUpdateRequestParams** | [**MarketingEventUpdateRequestParams**](MarketingEventUpdateRequestParams.md) | The details of the marketing event to update | 

### Return type

[**MarketingEventPublicDefaultResponse**](MarketingEventPublicDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

