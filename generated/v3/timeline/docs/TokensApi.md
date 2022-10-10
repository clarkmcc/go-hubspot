# \TokensApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TemplatesTokensArchive**](TokensApi.md#TemplatesTokensArchive) | **Delete** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId}/tokens/{tokenName} | Removes a token from the event template
[**TemplatesTokensCreate**](TokensApi.md#TemplatesTokensCreate) | **Post** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId}/tokens | Adds a token to an existing event template
[**TemplatesTokensUpdate**](TokensApi.md#TemplatesTokensUpdate) | **Put** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId}/tokens/{tokenName} | Updates an existing token on an event template



## TemplatesTokensArchive

> TemplatesTokensArchive(ctx, eventTemplateId, tokenName, appId).Execute()

Removes a token from the event template



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
    eventTemplateId := "eventTemplateId_example" // string | The event template ID.
    tokenName := "tokenName_example" // string | The token name.
    appId := int32(56) // int32 | The ID of the target app.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensApi.TemplatesTokensArchive(context.Background(), eventTemplateId, tokenName, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.TemplatesTokensArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**tokenName** | **string** | The token name. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesTokensArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesTokensCreate

> TimelineEventTemplateToken TemplatesTokensCreate(ctx, eventTemplateId, appId).TimelineEventTemplateToken(timelineEventTemplateToken).Execute()

Adds a token to an existing event template



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
    eventTemplateId := "eventTemplateId_example" // string | The event template ID.
    appId := int32(56) // int32 | The ID of the target app.
    timelineEventTemplateToken := *openapiclient.NewTimelineEventTemplateToken("Label_example", []openapiclient.TimelineEventTemplateTokenOption{*openapiclient.NewTimelineEventTemplateTokenOption("Value_example", "Label_example")}, "Name_example", "Type_example") // TimelineEventTemplateToken | The new token definition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensApi.TemplatesTokensCreate(context.Background(), eventTemplateId, appId).TimelineEventTemplateToken(timelineEventTemplateToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.TemplatesTokensCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesTokensCreate`: TimelineEventTemplateToken
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.TemplatesTokensCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesTokensCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timelineEventTemplateToken** | [**TimelineEventTemplateToken**](TimelineEventTemplateToken.md) | The new token definition. | 

### Return type

[**TimelineEventTemplateToken**](TimelineEventTemplateToken.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesTokensUpdate

> TimelineEventTemplateToken TemplatesTokensUpdate(ctx, eventTemplateId, tokenName, appId).TimelineEventTemplateTokenUpdateRequest(timelineEventTemplateTokenUpdateRequest).Execute()

Updates an existing token on an event template



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
    eventTemplateId := "eventTemplateId_example" // string | The event template ID.
    tokenName := "tokenName_example" // string | The token name.
    appId := int32(56) // int32 | The ID of the target app.
    timelineEventTemplateTokenUpdateRequest := *openapiclient.NewTimelineEventTemplateTokenUpdateRequest("Label_example", []openapiclient.TimelineEventTemplateTokenOption{*openapiclient.NewTimelineEventTemplateTokenOption("Value_example", "Label_example")}) // TimelineEventTemplateTokenUpdateRequest | The updated token definition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensApi.TemplatesTokensUpdate(context.Background(), eventTemplateId, tokenName, appId).TimelineEventTemplateTokenUpdateRequest(timelineEventTemplateTokenUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.TemplatesTokensUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesTokensUpdate`: TimelineEventTemplateToken
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.TemplatesTokensUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**tokenName** | **string** | The token name. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesTokensUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **timelineEventTemplateTokenUpdateRequest** | [**TimelineEventTemplateTokenUpdateRequest**](TimelineEventTemplateTokenUpdateRequest.md) | The updated token definition. | 

### Return type

[**TimelineEventTemplateToken**](TimelineEventTemplateToken.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

