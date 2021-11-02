# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletecrmv3extensionsaccountinguserAccountsaccountIdArchive**](UserAccountsApi.md#Deletecrmv3extensionsaccountinguserAccountsaccountIdArchive) | **Delete** /crm/v3/extensions/accounting/user-accounts/{accountId} | Delete user account
[**Putcrmv3extensionsaccountinguserAccountsReplace**](UserAccountsApi.md#Putcrmv3extensionsaccountinguserAccountsReplace) | **Put** /crm/v3/extensions/accounting/user-accounts | Create a user account

# **Deletecrmv3extensionsaccountinguserAccountsaccountIdArchive**
> Deletecrmv3extensionsaccountinguserAccountsaccountIdArchive(ctx, accountId)
Delete user account

Deletes a user account from HubSpot, meaning that HubSpot will no longer send requests to the external accounting system for this user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| The ID of the user account to delete. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Putcrmv3extensionsaccountinguserAccountsReplace**
> Putcrmv3extensionsaccountinguserAccountsReplace(ctx, body)
Create a user account

Creates an account which contains the information about the account in the external accounting system.  This *must* be called after a user connects their HubSpot account to the external accounting system, as there is no other way for HubSpot to obtain the external account details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateUserAccountRequestExternal**](CreateUserAccountRequestExternal.md)| The external accounting system user account information. | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

