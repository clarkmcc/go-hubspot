# \CardsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CardsArchive**](CardsApi.md#CardsArchive) | **Delete** /crm/v3/extensions/cards/{appId}/{cardId} | Delete a card
[**CardsCreate**](CardsApi.md#CardsCreate) | **Post** /crm/v3/extensions/cards/{appId} | Create a new card
[**CardsGetAll**](CardsApi.md#CardsGetAll) | **Get** /crm/v3/extensions/cards/{appId} | Get all cards
[**CardsGetByID**](CardsApi.md#CardsGetByID) | **Get** /crm/v3/extensions/cards/{appId}/{cardId} | Get a card.
[**CardsUpdate**](CardsApi.md#CardsUpdate) | **Patch** /crm/v3/extensions/cards/{appId}/{cardId} | Update a card



## CardsArchive

> CardsArchive(ctx, appId, cardId).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.CardsArchive(context.Background(), appId, cardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.CardsArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCardsArchiveRequest struct via the builder pattern


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


## CardsCreate

> CardResponse CardsCreate(ctx, appId).CardCreateRequest(cardCreateRequest).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.CardsCreate(context.Background(), appId).CardCreateRequest(cardCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.CardsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CardsCreate`: CardResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.CardsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCardsCreateRequest struct via the builder pattern


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


## CardsGetAll

> CardListResponse CardsGetAll(ctx, appId).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.CardsGetAll(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.CardsGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CardsGetAll`: CardListResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.CardsGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCardsGetAllRequest struct via the builder pattern


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


## CardsGetByID

> CardResponse CardsGetByID(ctx, appId, cardId).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.CardsGetByID(context.Background(), appId, cardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.CardsGetByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CardsGetByID`: CardResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.CardsGetByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 
**cardId** | **string** | The ID of the target card. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCardsGetByIDRequest struct via the builder pattern


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


## CardsUpdate

> CardResponse CardsUpdate(ctx, appId, cardId).CardPatchRequest(cardPatchRequest).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.CardsUpdate(context.Background(), appId, cardId).CardPatchRequest(cardPatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.CardsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CardsUpdate`: CardResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.CardsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 
**cardId** | **string** | The ID of the card to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCardsUpdateRequest struct via the builder pattern


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

