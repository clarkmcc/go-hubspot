# \SyncApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmV3ExtensionsAccountingSyncAppIdContactsCreateContact**](SyncApi.md#PostCrmV3ExtensionsAccountingSyncAppIdContactsCreateContact) | **Post** /crm/v3/extensions/accounting/sync/{appId}/contacts | Import contacts
[**PostCrmV3ExtensionsAccountingSyncAppIdProductsCreateProduct**](SyncApi.md#PostCrmV3ExtensionsAccountingSyncAppIdProductsCreateProduct) | **Post** /crm/v3/extensions/accounting/sync/{appId}/products | Import products



## PostCrmV3ExtensionsAccountingSyncAppIdContactsCreateContact

> ActionResponse PostCrmV3ExtensionsAccountingSyncAppIdContactsCreateContact(ctx, appId).SyncContactsRequest(syncContactsRequest).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyncApi.PostCrmV3ExtensionsAccountingSyncAppIdContactsCreateContact(context.Background(), appId).SyncContactsRequest(syncContactsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyncApi.PostCrmV3ExtensionsAccountingSyncAppIdContactsCreateContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ExtensionsAccountingSyncAppIdContactsCreateContact`: ActionResponse
    fmt.Fprintf(os.Stdout, "Response from `SyncApi.PostCrmV3ExtensionsAccountingSyncAppIdContactsCreateContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the accounting app. This is the identifier of the application created in your HubSpot developer portal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExtensionsAccountingSyncAppIdContactsCreateContactRequest struct via the builder pattern


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


## PostCrmV3ExtensionsAccountingSyncAppIdProductsCreateProduct

> ActionResponse PostCrmV3ExtensionsAccountingSyncAppIdProductsCreateProduct(ctx, appId).SyncProductsRequest(syncProductsRequest).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyncApi.PostCrmV3ExtensionsAccountingSyncAppIdProductsCreateProduct(context.Background(), appId).SyncProductsRequest(syncProductsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyncApi.PostCrmV3ExtensionsAccountingSyncAppIdProductsCreateProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ExtensionsAccountingSyncAppIdProductsCreateProduct`: ActionResponse
    fmt.Fprintf(os.Stdout, "Response from `SyncApi.PostCrmV3ExtensionsAccountingSyncAppIdProductsCreateProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the accounting app. This is the identifier of the application created in your HubSpot developer portal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExtensionsAccountingSyncAppIdProductsCreateProductRequest struct via the builder pattern


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

