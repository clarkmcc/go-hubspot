# \BusinessUnitApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBusinessUnitsV3BusinessUnitsUserUserId**](BusinessUnitApi.md#GetBusinessUnitsV3BusinessUnitsUserUserId) | **Get** /business-units/v3/business-units/user/{userId} | Get Business Units for a user



## GetBusinessUnitsV3BusinessUnitsUserUserId

> CollectionResponsePublicBusinessUnitNoPaging GetBusinessUnitsV3BusinessUnitsUserUserId(ctx, userId).Properties(properties).Name(name).Execute()

Get Business Units for a user



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
    userId := "userId_example" // string | Identifier of user to retrieve.
    properties := []string{"Inner_example"} // []string | The names of properties to optionally include in the response body. The only valid value is `logoMetadata`. (optional)
    name := []string{"Inner_example"} // []string | The names of Business Units to retrieve. If empty or not provided, then all associated Business Units will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BusinessUnitApi.GetBusinessUnitsV3BusinessUnitsUserUserId(context.Background(), userId).Properties(properties).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessUnitApi.GetBusinessUnitsV3BusinessUnitsUserUserId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBusinessUnitsV3BusinessUnitsUserUserId`: CollectionResponsePublicBusinessUnitNoPaging
    fmt.Fprintf(os.Stdout, "Response from `BusinessUnitApi.GetBusinessUnitsV3BusinessUnitsUserUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Identifier of user to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinessUnitsV3BusinessUnitsUserUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **properties** | **[]string** | The names of properties to optionally include in the response body. The only valid value is &#x60;logoMetadata&#x60;. | 
 **name** | **[]string** | The names of Business Units to retrieve. If empty or not provided, then all associated Business Units will be returned. | 

### Return type

[**CollectionResponsePublicBusinessUnitNoPaging**](CollectionResponsePublicBusinessUnitNoPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

