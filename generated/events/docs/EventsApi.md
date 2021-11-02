# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Geteventsv3eventsGetPage**](EventsApi.md#Geteventsv3eventsGetPage) | **Get** /events/v3/events | Returns a collection of events matching a query.

# **Geteventsv3eventsGetPage**
> CollectionResponseExternalUnifiedEvent Geteventsv3eventsGetPage(ctx, optional)
Returns a collection of events matching a query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventsApiGeteventsv3eventsGetPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventsApiGeteventsv3eventsGetPageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **occurredAfter** | **optional.Time**| The starting time as an ISO 8601 timestamp. | 
 **occurredBefore** | **optional.Time**| The ending time as an ISO 8601 timestamp. | 
 **objectType** | **optional.String**| The type of object being selected. Valid values are hubspot named object types (e.g. &#x60;contact&#x60;). | 
 **objectId** | **optional.Int64**| The id of the selected object. If not present, then the &#x60;objectProperty&#x60; parameter is required. | 
 **eventType** | **optional.String**| Limits the response to the specified event type.  For example &#x60;&amp;eventType&#x3D;e_visited_page&#x60; returns only &#x60;e_visited_page&#x60; events.  If not present all event types are returned. | 
 **after** | **optional.String**| An additional parameter that may be used to get the next &#x60;limit&#x60; set of results. | 
 **before** | **optional.String**|  | 
 **limit** | **optional.Int32**| The maximum number of events to return, defaults to 20. | 
 **sort** | [**optional.Interface of []string**](string.md)| Selects the sort field and order. Defaults to ascending, prefix with &#x60;-&#x60; for descending order. &#x60;occurredAt&#x60; is the only field supported for sorting. | 

### Return type

[**CollectionResponseExternalUnifiedEvent**](CollectionResponseExternalUnifiedEvent.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

