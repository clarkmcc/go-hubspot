# \GDPRApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmV3ObjectsLineItemsGdprDeletePurge**](GDPRApi.md#PostCrmV3ObjectsLineItemsGdprDeletePurge) | **Post** /crm/v3/objects/line_items/gdpr-delete | GDPR DELETE



## PostCrmV3ObjectsLineItemsGdprDeletePurge

> PostCrmV3ObjectsLineItemsGdprDeletePurge(ctx).PublicGdprDeleteInput(publicGdprDeleteInput).Execute()

GDPR DELETE



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
    publicGdprDeleteInput := *openapiclient.NewPublicGdprDeleteInput("ObjectId_example") // PublicGdprDeleteInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GDPRApi.PostCrmV3ObjectsLineItemsGdprDeletePurge(context.Background()).PublicGdprDeleteInput(publicGdprDeleteInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GDPRApi.PostCrmV3ObjectsLineItemsGdprDeletePurge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsLineItemsGdprDeletePurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicGdprDeleteInput** | [**PublicGdprDeleteInput**](PublicGdprDeleteInput.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

