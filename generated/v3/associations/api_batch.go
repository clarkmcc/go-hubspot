/*
Associations

Associations define the relationships between objects in HubSpot. These endpoints allow you to create, read, and remove associations.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package associations

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/clarkmcc/go-hubspot"
	"net/url"
	"strings"
)

// BatchApiService BatchApi service
type BatchApiService service

type ApiBatchArchiveRequest struct {
	ctx                         context.Context
	ApiService                  *BatchApiService
	fromObjectType              string
	toObjectType                string
	batchInputPublicAssociation *BatchInputPublicAssociation
}

func (r ApiBatchArchiveRequest) BatchInputPublicAssociation(batchInputPublicAssociation BatchInputPublicAssociation) ApiBatchArchiveRequest {
	r.batchInputPublicAssociation = &batchInputPublicAssociation
	return r
}

func (r ApiBatchArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.BatchArchiveExecute(r)
}

/*
BatchArchive Archive a batch of associations

Remove the associations between all pairs of objects identified in the request body.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fromObjectType
 @param toObjectType
 @return ApiBatchArchiveRequest
*/
func (a *BatchApiService) BatchArchive(ctx context.Context, fromObjectType string, toObjectType string) ApiBatchArchiveRequest {
	return ApiBatchArchiveRequest{
		ApiService:     a,
		ctx:            ctx,
		fromObjectType: fromObjectType,
		toObjectType:   toObjectType,
	}
}

// Execute executes the request
func (a *BatchApiService) BatchArchiveExecute(r ApiBatchArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchApiService.BatchArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/associations/{fromObjectType}/{toObjectType}/batch/archive"
	localVarPath = strings.Replace(localVarPath, "{"+"fromObjectType"+"}", url.PathEscape(parameterToString(r.fromObjectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterToString(r.toObjectType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.batchInputPublicAssociation
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
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

type ApiBatchCreateRequest struct {
	ctx                         context.Context
	ApiService                  *BatchApiService
	fromObjectType              string
	toObjectType                string
	batchInputPublicAssociation *BatchInputPublicAssociation
}

func (r ApiBatchCreateRequest) BatchInputPublicAssociation(batchInputPublicAssociation BatchInputPublicAssociation) ApiBatchCreateRequest {
	r.batchInputPublicAssociation = &batchInputPublicAssociation
	return r
}

func (r ApiBatchCreateRequest) Execute() (*BatchResponsePublicAssociation, *http.Response, error) {
	return r.ApiService.BatchCreateExecute(r)
}

/*
BatchCreate Create a batch of associations

Associate all pairs of objects identified in the request body.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fromObjectType
 @param toObjectType
 @return ApiBatchCreateRequest
*/
func (a *BatchApiService) BatchCreate(ctx context.Context, fromObjectType string, toObjectType string) ApiBatchCreateRequest {
	return ApiBatchCreateRequest{
		ApiService:     a,
		ctx:            ctx,
		fromObjectType: fromObjectType,
		toObjectType:   toObjectType,
	}
}

// Execute executes the request
//  @return BatchResponsePublicAssociation
func (a *BatchApiService) BatchCreateExecute(r ApiBatchCreateRequest) (*BatchResponsePublicAssociation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponsePublicAssociation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchApiService.BatchCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/associations/{fromObjectType}/{toObjectType}/batch/create"
	localVarPath = strings.Replace(localVarPath, "{"+"fromObjectType"+"}", url.PathEscape(parameterToString(r.fromObjectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterToString(r.toObjectType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.batchInputPublicAssociation
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
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

type ApiBatchReadRequest struct {
	ctx                      context.Context
	ApiService               *BatchApiService
	fromObjectType           string
	toObjectType             string
	batchInputPublicObjectId *BatchInputPublicObjectId
}

func (r ApiBatchReadRequest) BatchInputPublicObjectId(batchInputPublicObjectId BatchInputPublicObjectId) ApiBatchReadRequest {
	r.batchInputPublicObjectId = &batchInputPublicObjectId
	return r
}

func (r ApiBatchReadRequest) Execute() (*BatchResponsePublicAssociationMulti, *http.Response, error) {
	return r.ApiService.BatchReadExecute(r)
}

/*
BatchRead Read a batch of associations

Get the IDs of all `{toObjectType}` objects associated with those specified in the request body.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fromObjectType
 @param toObjectType
 @return ApiBatchReadRequest
*/
func (a *BatchApiService) BatchRead(ctx context.Context, fromObjectType string, toObjectType string) ApiBatchReadRequest {
	return ApiBatchReadRequest{
		ApiService:     a,
		ctx:            ctx,
		fromObjectType: fromObjectType,
		toObjectType:   toObjectType,
	}
}

// Execute executes the request
//  @return BatchResponsePublicAssociationMulti
func (a *BatchApiService) BatchReadExecute(r ApiBatchReadRequest) (*BatchResponsePublicAssociationMulti, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponsePublicAssociationMulti
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchApiService.BatchRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/associations/{fromObjectType}/{toObjectType}/batch/read"
	localVarPath = strings.Replace(localVarPath, "{"+"fromObjectType"+"}", url.PathEscape(parameterToString(r.fromObjectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterToString(r.toObjectType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.batchInputPublicObjectId
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
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
