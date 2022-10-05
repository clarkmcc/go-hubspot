/*
Custom Workflow Actions

Create custom workflow actions

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package actions

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// FunctionsApiService FunctionsApi service
type FunctionsApiService service

type ApiFunctionsArchiveRequest struct {
	ctx          context.Context
	ApiService   *FunctionsApiService
	definitionId string
	functionType string
	functionId   string
	appId        int32
}

func (r ApiFunctionsArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.FunctionsArchiveExecute(r)
}

/*
FunctionsArchive Delete a custom action function

Delete a function for a custom workflow action. This will remove the function itself as well as removing the association between the function and the custom action. This can't be undone.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param definitionId The ID of the custom workflow action
 @param functionType The type of function. This determines when the function will be called.
 @param functionId The ID qualifier for the function. This is used to specify which input field a function is associated with for `PRE_FETCH_OPTIONS` and `POST_FETCH_OPTIONS` function types.
 @param appId
 @return ApiFunctionsArchiveRequest
*/
func (a *FunctionsApiService) FunctionsArchive(ctx context.Context, definitionId string, functionType string, functionId string, appId int32) ApiFunctionsArchiveRequest {
	return ApiFunctionsArchiveRequest{
		ApiService:   a,
		ctx:          ctx,
		definitionId: definitionId,
		functionType: functionType,
		functionId:   functionId,
		appId:        appId,
	}
}

// Execute executes the request
func (a *FunctionsApiService) FunctionsArchiveExecute(r ApiFunctionsArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FunctionsApiService.FunctionsArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"definitionId"+"}", url.PathEscape(parameterToString(r.definitionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"functionType"+"}", url.PathEscape(parameterToString(r.functionType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"functionId"+"}", url.PathEscape(parameterToString(r.functionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
			}
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

type ApiFunctionsArchiveByTypeRequest struct {
	ctx          context.Context
	ApiService   *FunctionsApiService
	definitionId string
	functionType string
	appId        int32
}

func (r ApiFunctionsArchiveByTypeRequest) Execute() (*http.Response, error) {
	return r.ApiService.FunctionsArchiveByTypeExecute(r)
}

/*
FunctionsArchiveByType Delete a custom action function

Delete a function for a custom workflow action. This will remove the function itself as well as removing the association between the function and the custom action. This can't be undone.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param definitionId The ID of the custom workflow action.
 @param functionType The type of function. This determines when the function will be called.
 @param appId
 @return ApiFunctionsArchiveByTypeRequest
*/
func (a *FunctionsApiService) FunctionsArchiveByType(ctx context.Context, definitionId string, functionType string, appId int32) ApiFunctionsArchiveByTypeRequest {
	return ApiFunctionsArchiveByTypeRequest{
		ApiService:   a,
		ctx:          ctx,
		definitionId: definitionId,
		functionType: functionType,
		appId:        appId,
	}
}

// Execute executes the request
func (a *FunctionsApiService) FunctionsArchiveByTypeExecute(r ApiFunctionsArchiveByTypeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FunctionsApiService.FunctionsArchiveByType")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}/{definitionId}/functions/{functionType}"
	localVarPath = strings.Replace(localVarPath, "{"+"definitionId"+"}", url.PathEscape(parameterToString(r.definitionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"functionType"+"}", url.PathEscape(parameterToString(r.functionType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
			}
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

type ApiFunctionsCreateOrReplaceRequest struct {
	ctx          context.Context
	ApiService   *FunctionsApiService
	definitionId string
	functionType string
	functionId   string
	appId        int32
	body         *string
}

// The function source code. Must be valid JavaScript code.
func (r ApiFunctionsCreateOrReplaceRequest) Body(body string) ApiFunctionsCreateOrReplaceRequest {
	r.body = &body
	return r
}

func (r ApiFunctionsCreateOrReplaceRequest) Execute() (*ActionFunctionIdentifier, *http.Response, error) {
	return r.ApiService.FunctionsCreateOrReplaceExecute(r)
}

/*
FunctionsCreateOrReplace Create or replace a custom action function

Creates or replaces a function for a custom workflow action.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param definitionId The ID of the custom workflow action.
 @param functionType The type of function. This determines when the function will be called.
 @param functionId The ID qualifier for the function. This is used to specify which input field a function is associated with for `PRE_FETCH_OPTIONS` and `POST_FETCH_OPTIONS` function types.
 @param appId
 @return ApiFunctionsCreateOrReplaceRequest
*/
func (a *FunctionsApiService) FunctionsCreateOrReplace(ctx context.Context, definitionId string, functionType string, functionId string, appId int32) ApiFunctionsCreateOrReplaceRequest {
	return ApiFunctionsCreateOrReplaceRequest{
		ApiService:   a,
		ctx:          ctx,
		definitionId: definitionId,
		functionType: functionType,
		functionId:   functionId,
		appId:        appId,
	}
}

// Execute executes the request
//  @return ActionFunctionIdentifier
func (a *FunctionsApiService) FunctionsCreateOrReplaceExecute(r ApiFunctionsCreateOrReplaceRequest) (*ActionFunctionIdentifier, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ActionFunctionIdentifier
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FunctionsApiService.FunctionsCreateOrReplace")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"definitionId"+"}", url.PathEscape(parameterToString(r.definitionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"functionType"+"}", url.PathEscape(parameterToString(r.functionType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"functionId"+"}", url.PathEscape(parameterToString(r.functionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/plain"}

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
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
			}
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

type ApiFunctionsCreateOrReplaceByTypeRequest struct {
	ctx          context.Context
	ApiService   *FunctionsApiService
	definitionId string
	functionType string
	appId        int32
	body         *string
}

// The function source code. Must be valid JavaScript code.
func (r ApiFunctionsCreateOrReplaceByTypeRequest) Body(body string) ApiFunctionsCreateOrReplaceByTypeRequest {
	r.body = &body
	return r
}

func (r ApiFunctionsCreateOrReplaceByTypeRequest) Execute() (*ActionFunctionIdentifier, *http.Response, error) {
	return r.ApiService.FunctionsCreateOrReplaceByTypeExecute(r)
}

/*
FunctionsCreateOrReplaceByType Create or replace a custom action function

Creates or replaces a function for a custom workflow action.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param definitionId The ID of the custom workflow action.
 @param functionType The type of function. This determines when the function will be called.
 @param appId
 @return ApiFunctionsCreateOrReplaceByTypeRequest
*/
func (a *FunctionsApiService) FunctionsCreateOrReplaceByType(ctx context.Context, definitionId string, functionType string, appId int32) ApiFunctionsCreateOrReplaceByTypeRequest {
	return ApiFunctionsCreateOrReplaceByTypeRequest{
		ApiService:   a,
		ctx:          ctx,
		definitionId: definitionId,
		functionType: functionType,
		appId:        appId,
	}
}

// Execute executes the request
//  @return ActionFunctionIdentifier
func (a *FunctionsApiService) FunctionsCreateOrReplaceByTypeExecute(r ApiFunctionsCreateOrReplaceByTypeRequest) (*ActionFunctionIdentifier, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ActionFunctionIdentifier
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FunctionsApiService.FunctionsCreateOrReplaceByType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}/{definitionId}/functions/{functionType}"
	localVarPath = strings.Replace(localVarPath, "{"+"definitionId"+"}", url.PathEscape(parameterToString(r.definitionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"functionType"+"}", url.PathEscape(parameterToString(r.functionType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/plain"}

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
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
			}
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

type ApiFunctionsGetByIDRequest struct {
	ctx          context.Context
	ApiService   *FunctionsApiService
	definitionId string
	functionType string
	functionId   string
	appId        int32
}

func (r ApiFunctionsGetByIDRequest) Execute() (*ActionFunction, *http.Response, error) {
	return r.ApiService.FunctionsGetByIDExecute(r)
}

/*
FunctionsGetByID Get a custom action function

Returns the given function for a custom workflow action.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param definitionId The ID of the custom workflow action.
 @param functionType The type of function. This determines when the function will be called.
 @param functionId The ID qualifier for the function. This is used to specify which input field a function is associated with for `PRE_FETCH_OPTIONS` and `POST_FETCH_OPTIONS` function types.
 @param appId
 @return ApiFunctionsGetByIDRequest
*/
func (a *FunctionsApiService) FunctionsGetByID(ctx context.Context, definitionId string, functionType string, functionId string, appId int32) ApiFunctionsGetByIDRequest {
	return ApiFunctionsGetByIDRequest{
		ApiService:   a,
		ctx:          ctx,
		definitionId: definitionId,
		functionType: functionType,
		functionId:   functionId,
		appId:        appId,
	}
}

// Execute executes the request
//  @return ActionFunction
func (a *FunctionsApiService) FunctionsGetByIDExecute(r ApiFunctionsGetByIDRequest) (*ActionFunction, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ActionFunction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FunctionsApiService.FunctionsGetByID")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"definitionId"+"}", url.PathEscape(parameterToString(r.definitionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"functionType"+"}", url.PathEscape(parameterToString(r.functionType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"functionId"+"}", url.PathEscape(parameterToString(r.functionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
			}
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

type ApiFunctionsGetByTypeRequest struct {
	ctx          context.Context
	ApiService   *FunctionsApiService
	definitionId string
	functionType string
	appId        int32
}

func (r ApiFunctionsGetByTypeRequest) Execute() (*ActionFunction, *http.Response, error) {
	return r.ApiService.FunctionsGetByTypeExecute(r)
}

/*
FunctionsGetByType Get a custom action function

Returns the given function for a custom workflow action.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param definitionId The ID of the custom workflow action.
 @param functionType The type of function. This determines when the function will be called.
 @param appId
 @return ApiFunctionsGetByTypeRequest
*/
func (a *FunctionsApiService) FunctionsGetByType(ctx context.Context, definitionId string, functionType string, appId int32) ApiFunctionsGetByTypeRequest {
	return ApiFunctionsGetByTypeRequest{
		ApiService:   a,
		ctx:          ctx,
		definitionId: definitionId,
		functionType: functionType,
		appId:        appId,
	}
}

// Execute executes the request
//  @return ActionFunction
func (a *FunctionsApiService) FunctionsGetByTypeExecute(r ApiFunctionsGetByTypeRequest) (*ActionFunction, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ActionFunction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FunctionsApiService.FunctionsGetByType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}/{definitionId}/functions/{functionType}"
	localVarPath = strings.Replace(localVarPath, "{"+"definitionId"+"}", url.PathEscape(parameterToString(r.definitionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"functionType"+"}", url.PathEscape(parameterToString(r.functionType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
			}
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

type ApiFunctionsGetPageRequest struct {
	ctx          context.Context
	ApiService   *FunctionsApiService
	definitionId string
	appId        int32
}

func (r ApiFunctionsGetPageRequest) Execute() (*CollectionResponseActionFunctionIdentifierNoPaging, *http.Response, error) {
	return r.ApiService.FunctionsGetPageExecute(r)
}

/*
FunctionsGetPage Get all custom action functions

Returns a list of all functions that are associated with the given custom workflow action.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param definitionId The ID of the custom workflow action.
 @param appId
 @return ApiFunctionsGetPageRequest
*/
func (a *FunctionsApiService) FunctionsGetPage(ctx context.Context, definitionId string, appId int32) ApiFunctionsGetPageRequest {
	return ApiFunctionsGetPageRequest{
		ApiService:   a,
		ctx:          ctx,
		definitionId: definitionId,
		appId:        appId,
	}
}

// Execute executes the request
//  @return CollectionResponseActionFunctionIdentifierNoPaging
func (a *FunctionsApiService) FunctionsGetPageExecute(r ApiFunctionsGetPageRequest) (*CollectionResponseActionFunctionIdentifierNoPaging, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionResponseActionFunctionIdentifierNoPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FunctionsApiService.FunctionsGetPage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}/{definitionId}/functions"
	localVarPath = strings.Replace(localVarPath, "{"+"definitionId"+"}", url.PathEscape(parameterToString(r.definitionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
			}
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
