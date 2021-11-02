# {{classname}}

All URIs are relative to *https://api.hubapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getcrmv3objectsfeedbackSubmissionsfeedbackSubmissionIdassociationstoObjectTypeGetAll**](AssociationsApi.md#Getcrmv3objectsfeedbackSubmissionsfeedbackSubmissionIdassociationstoObjectTypeGetAll) | **Get** /crm/v3/objects/feedback_submissions/{feedbackSubmissionId}/associations/{toObjectType} | List associations of a feedback submission by type

# **Getcrmv3objectsfeedbackSubmissionsfeedbackSubmissionIdassociationstoObjectTypeGetAll**
> CollectionResponseAssociatedIdForwardPaging Getcrmv3objectsfeedbackSubmissionsfeedbackSubmissionIdassociationstoObjectTypeGetAll(ctx, feedbackSubmissionId, toObjectType, optional)
List associations of a feedback submission by type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **feedbackSubmissionId** | **string**|  | 
  **toObjectType** | **string**|  | 
 **optional** | ***AssociationsApiGetcrmv3objectsfeedbackSubmissionsfeedbackSubmissionIdassociationstoObjectTypeGetAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssociationsApiGetcrmv3objectsfeedbackSubmissionsfeedbackSubmissionIdassociationstoObjectTypeGetAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **optional.String**| The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **optional.Int32**| The maximum number of results to display per page. | [default to 500]

### Return type

[**CollectionResponseAssociatedIdForwardPaging**](CollectionResponseAssociatedIdForwardPaging.md)

### Authorization

[hapikey](../README.md#hapikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

