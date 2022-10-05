# \BehavioralEventsTrackingApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Send**](BehavioralEventsTrackingApi.md#Send) | **Post** /events/v3/send | Sends Custom Behavioral Event



## Send

> Send(ctx).BehavioralEventHttpCompletionRequest(behavioralEventHttpCompletionRequest).Execute()

Sends Custom Behavioral Event



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
    behavioralEventHttpCompletionRequest := *openapiclient.NewBehavioralEventHttpCompletionRequest("EventName_example", map[string]string{"key": "Inner_example"}) // BehavioralEventHttpCompletionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BehavioralEventsTrackingApi.Send(context.Background()).BehavioralEventHttpCompletionRequest(behavioralEventHttpCompletionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BehavioralEventsTrackingApi.Send``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **behavioralEventHttpCompletionRequest** | [**BehavioralEventHttpCompletionRequest**](BehavioralEventHttpCompletionRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

