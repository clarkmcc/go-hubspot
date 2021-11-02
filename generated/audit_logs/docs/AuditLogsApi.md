# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getcmsv3auditLogsGetPage**](AuditLogsApi.md#Getcmsv3auditLogsGetPage) | **Get** /cms/v3/audit-logs/ | Query audit logs

# **Getcmsv3auditLogsGetPage**
> CollectionResponsePublicAuditLog Getcmsv3auditLogsGetPage(ctx, optional)
Query audit logs

Returns audit logs based on filters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuditLogsApiGetcmsv3auditLogsGetPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuditLogsApiGetcmsv3auditLogsGetPageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectId** | [**optional.Interface of []string**](string.md)| Comma separated list of object ids to filter by. | 
 **userId** | [**optional.Interface of []string**](string.md)| Comma separated list of user ids to filter by. | 
 **after** | **optional.String**| Timestamp after which audit logs will be returned | 
 **before** | **optional.String**| Timestamp before which audit logs will be returned | 
 **sort** | [**optional.Interface of []string**](string.md)| The sort direction for the audit logs. (Can only sort by timestamp). | 
 **eventType** | [**optional.Interface of []string**](string.md)| Comma separated list of event types to filter by (CREATED, UPDATED, PUBLISHED, DELETED, UNPUBLISHED). | 
 **limit** | **optional.Int32**| The number of logs to return. | 
 **objectType** | [**optional.Interface of []string**](string.md)| Comma separated list of object types to filter by (BLOG, LANDING_PAGE, DOMAIN, HUBDB_TABLE etc.) | 

### Return type

[**CollectionResponsePublicAuditLog**](CollectionResponsePublicAuditLog.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

