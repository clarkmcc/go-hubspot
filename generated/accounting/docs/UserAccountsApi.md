# \UserAccountsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3ExtensionsAccountingUserAccountsAccountIdArchive**](UserAccountsApi.md#DeleteCrmV3ExtensionsAccountingUserAccountsAccountIdArchive) | **Delete** /crm/v3/extensions/accounting/user-accounts/{accountId} | Delete user account
[**PutCrmV3ExtensionsAccountingUserAccountsReplace**](UserAccountsApi.md#PutCrmV3ExtensionsAccountingUserAccountsReplace) | **Put** /crm/v3/extensions/accounting/user-accounts | Create a user account



## DeleteCrmV3ExtensionsAccountingUserAccountsAccountIdArchive

> DeleteCrmV3ExtensionsAccountingUserAccountsAccountIdArchive(ctx, accountId).Execute()

Delete user account



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
    accountId := "accountId_example" // string | The ID of the user account to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAccountsApi.DeleteCrmV3ExtensionsAccountingUserAccountsAccountIdArchive(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAccountsApi.DeleteCrmV3ExtensionsAccountingUserAccountsAccountIdArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the user account to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3ExtensionsAccountingUserAccountsAccountIdArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmV3ExtensionsAccountingUserAccountsReplace

> PutCrmV3ExtensionsAccountingUserAccountsReplace(ctx).CreateUserAccountRequestExternal(createUserAccountRequestExternal).Execute()

Create a user account



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
    createUserAccountRequestExternal := *openapiclient.NewCreateUserAccountRequestExternal("AccountId_example", "AccountName_example", "CurrencyCode_example") // CreateUserAccountRequestExternal | The external accounting system user account information.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAccountsApi.PutCrmV3ExtensionsAccountingUserAccountsReplace(context.Background()).CreateUserAccountRequestExternal(createUserAccountRequestExternal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAccountsApi.PutCrmV3ExtensionsAccountingUserAccountsReplace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV3ExtensionsAccountingUserAccountsReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserAccountRequestExternal** | [**CreateUserAccountRequestExternal**](CreateUserAccountRequestExternal.md) | The external accounting system user account information. | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

