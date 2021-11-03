# \SingleSendApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMarketingV3TransactionalSingleEmailSendSendEmail**](SingleSendApi.md#PostMarketingV3TransactionalSingleEmailSendSendEmail) | **Post** /marketing/v3/transactional/single-email/send | Send a single transactional email asynchronously.



## PostMarketingV3TransactionalSingleEmailSendSendEmail

> EmailSendStatusView PostMarketingV3TransactionalSingleEmailSendSendEmail(ctx).PublicSingleSendRequestEgg(publicSingleSendRequestEgg).Execute()

Send a single transactional email asynchronously.



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
    publicSingleSendRequestEgg := *openapiclient.NewPublicSingleSendRequestEgg(*openapiclient.NewPublicSingleSendEmail([]string{"ReplyTo_example"}, []string{"Cc_example"}, []string{"Bcc_example"}), map[string]string{"key": "Inner_example"}, map[string]interface{}(123), int32(123)) // PublicSingleSendRequestEgg | A request object describing the email to send. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SingleSendApi.PostMarketingV3TransactionalSingleEmailSendSendEmail(context.Background()).PublicSingleSendRequestEgg(publicSingleSendRequestEgg).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SingleSendApi.PostMarketingV3TransactionalSingleEmailSendSendEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMarketingV3TransactionalSingleEmailSendSendEmail`: EmailSendStatusView
    fmt.Fprintf(os.Stdout, "Response from `SingleSendApi.PostMarketingV3TransactionalSingleEmailSendSendEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3TransactionalSingleEmailSendSendEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicSingleSendRequestEgg** | [**PublicSingleSendRequestEgg**](PublicSingleSendRequestEgg.md) | A request object describing the email to send. | 

### Return type

[**EmailSendStatusView**](EmailSendStatusView.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

