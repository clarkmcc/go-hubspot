# \CardsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3ExtensionsCardsAppIdCardIdArchive**](CardsApi.md#DeleteCrmV3ExtensionsCardsAppIdCardIdArchive) | **Delete** /crm/v3/extensions/cards/{appId}/{cardId} | Delete a card
[**GetCrmV3ExtensionsCardsAppIdCardIdGetById**](CardsApi.md#GetCrmV3ExtensionsCardsAppIdCardIdGetById) | **Get** /crm/v3/extensions/cards/{appId}/{cardId} | Get a card.
[**GetCrmV3ExtensionsCardsAppIdGetAll**](CardsApi.md#GetCrmV3ExtensionsCardsAppIdGetAll) | **Get** /crm/v3/extensions/cards/{appId} | Get all cards
[**PatchCrmV3ExtensionsCardsAppIdCardIdUpdate**](CardsApi.md#PatchCrmV3ExtensionsCardsAppIdCardIdUpdate) | **Patch** /crm/v3/extensions/cards/{appId}/{cardId} | Update a card
[**PostCrmV3ExtensionsCardsAppIdCreate**](CardsApi.md#PostCrmV3ExtensionsCardsAppIdCreate) | **Post** /crm/v3/extensions/cards/{appId} | Create a new card



## DeleteCrmV3ExtensionsCardsAppIdCardIdArchive

> DeleteCrmV3ExtensionsCardsAppIdCardIdArchive(ctx, appId, cardId).Execute()

Delete a card



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
    appId := int32(56) // int32 | The ID of the target app.
    cardId := "cardId_example" // string | The ID of the card to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardsApi.DeleteCrmV3ExtensionsCardsAppIdCardIdArchive(context.Background(), appId, cardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.DeleteCrmV3ExtensionsCardsAppIdCardIdArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 
**cardId** | **string** | The ID of the card to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3ExtensionsCardsAppIdCardIdArchiveRequest struct via the builder pattern


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


## GetCrmV3ExtensionsCardsAppIdCardIdGetById

> CardResponse GetCrmV3ExtensionsCardsAppIdCardIdGetById(ctx, appId, cardId).Execute()

Get a card.



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
    appId := int32(56) // int32 | The ID of the target app.
    cardId := "cardId_example" // string | The ID of the target card.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardsApi.GetCrmV3ExtensionsCardsAppIdCardIdGetById(context.Background(), appId, cardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.GetCrmV3ExtensionsCardsAppIdCardIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3ExtensionsCardsAppIdCardIdGetById`: CardResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.GetCrmV3ExtensionsCardsAppIdCardIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 
**cardId** | **string** | The ID of the target card. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ExtensionsCardsAppIdCardIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CardResponse**](CardResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3ExtensionsCardsAppIdGetAll

> CardListResponse GetCrmV3ExtensionsCardsAppIdGetAll(ctx, appId).Execute()

Get all cards



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
    appId := int32(56) // int32 | The ID of the target app.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardsApi.GetCrmV3ExtensionsCardsAppIdGetAll(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.GetCrmV3ExtensionsCardsAppIdGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3ExtensionsCardsAppIdGetAll`: CardListResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.GetCrmV3ExtensionsCardsAppIdGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ExtensionsCardsAppIdGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CardListResponse**](CardListResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCrmV3ExtensionsCardsAppIdCardIdUpdate

> CardResponse PatchCrmV3ExtensionsCardsAppIdCardIdUpdate(ctx, appId, cardId).CardPatchRequest(cardPatchRequest).Execute()

Update a card



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
    appId := int32(56) // int32 | The ID of the target app.
    cardId := "cardId_example" // string | The ID of the card to update.
    cardPatchRequest := *openapiclient.NewCardPatchRequest() // CardPatchRequest | Card definition fields to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardsApi.PatchCrmV3ExtensionsCardsAppIdCardIdUpdate(context.Background(), appId, cardId).CardPatchRequest(cardPatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.PatchCrmV3ExtensionsCardsAppIdCardIdUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCrmV3ExtensionsCardsAppIdCardIdUpdate`: CardResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.PatchCrmV3ExtensionsCardsAppIdCardIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 
**cardId** | **string** | The ID of the card to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCrmV3ExtensionsCardsAppIdCardIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cardPatchRequest** | [**CardPatchRequest**](CardPatchRequest.md) | Card definition fields to be updated. | 

### Return type

[**CardResponse**](CardResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ExtensionsCardsAppIdCreate

> CardResponse PostCrmV3ExtensionsCardsAppIdCreate(ctx, appId).CardCreateRequest(cardCreateRequest).Execute()

Create a new card



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
    appId := int32(56) // int32 | The ID of the target app.
    cardCreateRequest := *openapiclient.NewCardCreateRequest("Title_example", *openapiclient.NewCardFetchBody("TargetUrl_example", []openapiclient.CardObjectTypeBody{*openapiclient.NewCardObjectTypeBody("Name_example", []string{"PropertiesToSend_example"})}), *openapiclient.NewCardDisplayBody([]openapiclient.CardDisplayProperty{*openapiclient.NewCardDisplayProperty("Name_example", "Label_example", "DataType_example", []openapiclient.DisplayOption{*openapiclient.NewDisplayOption("Name_example", "Label_example", "Type_example")})}), *openapiclient.NewCardActions([]string{"BaseUrls_example"})) // CardCreateRequest | The new card definition.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardsApi.PostCrmV3ExtensionsCardsAppIdCreate(context.Background(), appId).CardCreateRequest(cardCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.PostCrmV3ExtensionsCardsAppIdCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ExtensionsCardsAppIdCreate`: CardResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.PostCrmV3ExtensionsCardsAppIdCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExtensionsCardsAppIdCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cardCreateRequest** | [**CardCreateRequest**](CardCreateRequest.md) | The new card definition. | 

### Return type

[**CardResponse**](CardResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

