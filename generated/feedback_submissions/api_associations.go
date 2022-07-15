/*
Feedback Submissions

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package feedback_submissions

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// AssociationsApiService AssociationsApi service
type AssociationsApiService service

type ApiDeleteCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveRequest struct {
	ctx                  context.Context
	ApiService           *AssociationsApiService
	feedbackSubmissionId string
	toObjectType         string
	toObjectId           string
	associationType      string
}

func (r ApiDeleteCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveExecute(r)
}

/*
DeleteCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive Remove an association between two feedback submissions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param feedbackSubmissionId
 @param toObjectType
 @param toObjectId
 @param associationType
 @return ApiDeleteCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveRequest
*/
func (a *AssociationsApiService) DeleteCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive(ctx context.Context, feedbackSubmissionId string, toObjectType string, toObjectId string, associationType string) ApiDeleteCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveRequest {
	return ApiDeleteCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveRequest{
		ApiService:           a,
		ctx:                  ctx,
		feedbackSubmissionId: feedbackSubmissionId,
		toObjectType:         toObjectType,
		toObjectId:           toObjectId,
		associationType:      associationType,
	}
}

// Execute executes the request
func (a *AssociationsApiService) DeleteCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveExecute(r ApiDeleteCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssociationsApiService.DeleteCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/feedback_submissions/{feedbackSubmissionId}/associations/{toObjectType}/{toObjectId}/{associationType}"
	localVarPath = strings.Replace(localVarPath, "{"+"feedbackSubmissionId"+"}", url.PathEscape(parameterToString(r.feedbackSubmissionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterToString(r.toObjectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectId"+"}", url.PathEscape(parameterToString(r.toObjectId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"associationType"+"}", url.PathEscape(parameterToString(r.associationType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeGetPageRequest struct {
	ctx                  context.Context
	ApiService           *AssociationsApiService
	feedbackSubmissionId string
	toObjectType         string
	after                *string
	limit                *int32
}

// The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results.
func (r ApiGetCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeGetPageRequest) After(after string) ApiGetCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeGetPageRequest {
	r.after = &after
	return r
}

// The maximum number of results to display per page.
func (r ApiGetCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeGetPageRequest) Limit(limit int32) ApiGetCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeGetPageRequest {
	r.limit = &limit
	return r
}

func (r ApiGetCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeGetPageRequest) Execute() (*CollectionResponseAssociatedIdForwardPaging, *http.Response, error) {
	return r.ApiService.GetCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeGetPageExecute(r)
}

/*
GetCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeGetPage List associations of a feedback submission by type

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param feedbackSubmissionId
 @param toObjectType
 @return ApiGetCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeGetPageRequest
*/
func (a *AssociationsApiService) GetCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeGetPage(ctx context.Context, feedbackSubmissionId string, toObjectType string) ApiGetCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeGetPageRequest {
	return ApiGetCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeGetPageRequest{
		ApiService:           a,
		ctx:                  ctx,
		feedbackSubmissionId: feedbackSubmissionId,
		toObjectType:         toObjectType,
	}
}

// Execute executes the request
//  @return CollectionResponseAssociatedIdForwardPaging
func (a *AssociationsApiService) GetCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeGetPageExecute(r ApiGetCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeGetPageRequest) (*CollectionResponseAssociatedIdForwardPaging, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionResponseAssociatedIdForwardPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssociationsApiService.GetCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeGetPage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/feedback_submissions/{feedbackSubmissionId}/associations/{toObjectType}"
	localVarPath = strings.Replace(localVarPath, "{"+"feedbackSubmissionId"+"}", url.PathEscape(parameterToString(r.feedbackSubmissionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterToString(r.toObjectType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPutCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateRequest struct {
	ctx                  context.Context
	ApiService           *AssociationsApiService
	feedbackSubmissionId string
	toObjectType         string
	toObjectId           string
	associationType      string
}

func (r ApiPutCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateRequest) Execute() (*SimplePublicObjectWithAssociations, *http.Response, error) {
	return r.ApiService.PutCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateExecute(r)
}

/*
PutCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate Associate a feedback submission with another object

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param feedbackSubmissionId
 @param toObjectType
 @param toObjectId
 @param associationType
 @return ApiPutCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateRequest
*/
func (a *AssociationsApiService) PutCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate(ctx context.Context, feedbackSubmissionId string, toObjectType string, toObjectId string, associationType string) ApiPutCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateRequest {
	return ApiPutCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateRequest{
		ApiService:           a,
		ctx:                  ctx,
		feedbackSubmissionId: feedbackSubmissionId,
		toObjectType:         toObjectType,
		toObjectId:           toObjectId,
		associationType:      associationType,
	}
}

// Execute executes the request
//  @return SimplePublicObjectWithAssociations
func (a *AssociationsApiService) PutCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateExecute(r ApiPutCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateRequest) (*SimplePublicObjectWithAssociations, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SimplePublicObjectWithAssociations
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssociationsApiService.PutCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/feedback_submissions/{feedbackSubmissionId}/associations/{toObjectType}/{toObjectId}/{associationType}"
	localVarPath = strings.Replace(localVarPath, "{"+"feedbackSubmissionId"+"}", url.PathEscape(parameterToString(r.feedbackSubmissionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterToString(r.toObjectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectId"+"}", url.PathEscape(parameterToString(r.toObjectId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"associationType"+"}", url.PathEscape(parameterToString(r.associationType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}