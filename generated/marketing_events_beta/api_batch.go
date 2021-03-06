/*
Marketing Events Extension

These APIs allow you to interact with HubSpot's Marketing Events Extension. It allows you to: * Create, Read or update Marketing Event information in HubSpot * Specify whether a HubSpot contact has registered, attended or cancelled a registration to a Marketing Event. * Specify a URL that can be called to get the details of a Marketing Event.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_events_beta

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// BatchApiService BatchApi service
type BatchApiService service

type ApiPostMarketingV3MarketingEventsEventsDeleteArchiveRequest struct {
	ctx                                              context.Context
	ApiService                                       *BatchApiService
	batchInputMarketingEventExternalUniqueIdentifier *BatchInputMarketingEventExternalUniqueIdentifier
}

// The details of the marketing events to delete
func (r ApiPostMarketingV3MarketingEventsEventsDeleteArchiveRequest) BatchInputMarketingEventExternalUniqueIdentifier(batchInputMarketingEventExternalUniqueIdentifier BatchInputMarketingEventExternalUniqueIdentifier) ApiPostMarketingV3MarketingEventsEventsDeleteArchiveRequest {
	r.batchInputMarketingEventExternalUniqueIdentifier = &batchInputMarketingEventExternalUniqueIdentifier
	return r
}

func (r ApiPostMarketingV3MarketingEventsEventsDeleteArchiveRequest) Execute() (*Error, *http.Response, error) {
	return r.ApiService.PostMarketingV3MarketingEventsEventsDeleteArchiveExecute(r)
}

/*
PostMarketingV3MarketingEventsEventsDeleteArchive Delete multiple marketing events

Bulk delete a number of marketing events in HubSpot

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostMarketingV3MarketingEventsEventsDeleteArchiveRequest
*/
func (a *BatchApiService) PostMarketingV3MarketingEventsEventsDeleteArchive(ctx context.Context) ApiPostMarketingV3MarketingEventsEventsDeleteArchiveRequest {
	return ApiPostMarketingV3MarketingEventsEventsDeleteArchiveRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return Error
func (a *BatchApiService) PostMarketingV3MarketingEventsEventsDeleteArchiveExecute(r ApiPostMarketingV3MarketingEventsEventsDeleteArchiveRequest) (*Error, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Error
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchApiService.PostMarketingV3MarketingEventsEventsDeleteArchive")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/marketing/v3/marketing-events/events/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputMarketingEventExternalUniqueIdentifier == nil {
		return localVarReturnValue, nil, reportError("batchInputMarketingEventExternalUniqueIdentifier is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.batchInputMarketingEventExternalUniqueIdentifier
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

type ApiPostMarketingV3MarketingEventsEventsUpsertDoUpsertRequest struct {
	ctx                                         context.Context
	ApiService                                  *BatchApiService
	batchInputMarketingEventCreateRequestParams *BatchInputMarketingEventCreateRequestParams
}

// The details of the marketing events to upsert
func (r ApiPostMarketingV3MarketingEventsEventsUpsertDoUpsertRequest) BatchInputMarketingEventCreateRequestParams(batchInputMarketingEventCreateRequestParams BatchInputMarketingEventCreateRequestParams) ApiPostMarketingV3MarketingEventsEventsUpsertDoUpsertRequest {
	r.batchInputMarketingEventCreateRequestParams = &batchInputMarketingEventCreateRequestParams
	return r
}

func (r ApiPostMarketingV3MarketingEventsEventsUpsertDoUpsertRequest) Execute() (*BatchResponseMarketingEventPublicDefaultResponse, *http.Response, error) {
	return r.ApiService.PostMarketingV3MarketingEventsEventsUpsertDoUpsertExecute(r)
}

/*
PostMarketingV3MarketingEventsEventsUpsertDoUpsert Create or update multiple marketing events

Upset multiple Marketing Event. If there is an existing Marketing event with the specified id, it will be updated; otherwise a new event will be created.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostMarketingV3MarketingEventsEventsUpsertDoUpsertRequest
*/
func (a *BatchApiService) PostMarketingV3MarketingEventsEventsUpsertDoUpsert(ctx context.Context) ApiPostMarketingV3MarketingEventsEventsUpsertDoUpsertRequest {
	return ApiPostMarketingV3MarketingEventsEventsUpsertDoUpsertRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return BatchResponseMarketingEventPublicDefaultResponse
func (a *BatchApiService) PostMarketingV3MarketingEventsEventsUpsertDoUpsertExecute(r ApiPostMarketingV3MarketingEventsEventsUpsertDoUpsertRequest) (*BatchResponseMarketingEventPublicDefaultResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponseMarketingEventPublicDefaultResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchApiService.PostMarketingV3MarketingEventsEventsUpsertDoUpsert")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/marketing/v3/marketing-events/events/upsert"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputMarketingEventCreateRequestParams == nil {
		return localVarReturnValue, nil, reportError("batchInputMarketingEventCreateRequestParams is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.batchInputMarketingEventCreateRequestParams
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
