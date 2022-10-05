# \SyncApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SyncCreateContact**](SyncApi.md#SyncCreateContact) | **Post** /crm/v3/extensions/accounting/sync/{appId}/contacts | Import contacts
[**SyncCreateProduct**](SyncApi.md#SyncCreateProduct) | **Post** /crm/v3/extensions/accounting/sync/{appId}/products | Import products



## SyncCreateContact

> ActionResponse SyncCreateContact(ctx, appId).SyncContactsRequest(syncContactsRequest).Execute()

Import contacts



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
    appId := int32(56) // int32 | The ID of the accounting app. This is the identifier of the application created in your HubSpot developer portal.
    syncContactsRequest := *openapiclient.NewSyncContactsRequest("acct-app-123", []openapiclient.UpdatedContact{*openapiclient.NewUpdatedContact("UPDATE", time.Now(), "johndoe@company.com", "acct-app-123")}) // SyncContactsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyncApi.SyncCreateContact(context.Background(), appId).SyncContactsRequest(syncContactsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyncApi.SyncCreateContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncCreateContact`: ActionResponse
    fmt.Fprintf(os.Stdout, "Response from `SyncApi.SyncCreateContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the accounting app. This is the identifier of the application created in your HubSpot developer portal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncCreateContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **syncContactsRequest** | [**SyncContactsRequest**](SyncContactsRequest.md) |  | 

### Return type

[**ActionResponse**](ActionResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncCreateProduct

> ActionResponse SyncCreateProduct(ctx, appId).SyncProductsRequest(syncProductsRequest).Execute()

Import products



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
    appId := int32(56) // int32 | The ID of the accounting app. This is the identifier of the application created in your HubSpot developer portal.
    syncProductsRequest := *openapiclient.NewSyncProductsRequest("AccountId_example", []openapiclient.UpdatedProduct{*openapiclient.NewUpdatedProduct("SyncAction_example", time.Now(), float32(123), "Id_example", map[string]string{"key": "Inner_example"})}) // SyncProductsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyncApi.SyncCreateProduct(context.Background(), appId).SyncProductsRequest(syncProductsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyncApi.SyncCreateProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncCreateProduct`: ActionResponse
    fmt.Fprintf(os.Stdout, "Response from `SyncApi.SyncCreateProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the accounting app. This is the identifier of the application created in your HubSpot developer portal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncCreateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **syncProductsRequest** | [**SyncProductsRequest**](SyncProductsRequest.md) |  | 

### Return type

[**ActionResponse**](ActionResponse.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

