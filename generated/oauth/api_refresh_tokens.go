/*
OAuthService

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oauth

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// RefreshTokensApiService RefreshTokensApi service
type RefreshTokensApiService service

type ApiDeleteOauthV1RefreshTokensTokenArchiveRefreshTokenRequest struct {
	ctx        context.Context
	ApiService *RefreshTokensApiService
	token      string
}

func (r ApiDeleteOauthV1RefreshTokensTokenArchiveRefreshTokenRequest) Execute() (*Error, *http.Response, error) {
	return r.ApiService.DeleteOauthV1RefreshTokensTokenArchiveRefreshTokenExecute(r)
}

/*
DeleteOauthV1RefreshTokensTokenArchiveRefreshToken Method for DeleteOauthV1RefreshTokensTokenArchiveRefreshToken

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token
 @return ApiDeleteOauthV1RefreshTokensTokenArchiveRefreshTokenRequest
*/
func (a *RefreshTokensApiService) DeleteOauthV1RefreshTokensTokenArchiveRefreshToken(ctx context.Context, token string) ApiDeleteOauthV1RefreshTokensTokenArchiveRefreshTokenRequest {
	return ApiDeleteOauthV1RefreshTokensTokenArchiveRefreshTokenRequest{
		ApiService: a,
		ctx:        ctx,
		token:      token,
	}
}

// Execute executes the request
//  @return Error
func (a *RefreshTokensApiService) DeleteOauthV1RefreshTokensTokenArchiveRefreshTokenExecute(r ApiDeleteOauthV1RefreshTokensTokenArchiveRefreshTokenRequest) (*Error, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Error
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RefreshTokensApiService.DeleteOauthV1RefreshTokensTokenArchiveRefreshToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/v1/refresh-tokens/{token}"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterToString(r.token, "")), -1)

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

type ApiGetOauthV1RefreshTokensTokenGetRefreshTokenRequest struct {
	ctx        context.Context
	ApiService *RefreshTokensApiService
	token      string
}

func (r ApiGetOauthV1RefreshTokensTokenGetRefreshTokenRequest) Execute() (*RefreshTokenInfoResponse, *http.Response, error) {
	return r.ApiService.GetOauthV1RefreshTokensTokenGetRefreshTokenExecute(r)
}

/*
GetOauthV1RefreshTokensTokenGetRefreshToken Method for GetOauthV1RefreshTokensTokenGetRefreshToken

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token
 @return ApiGetOauthV1RefreshTokensTokenGetRefreshTokenRequest
*/
func (a *RefreshTokensApiService) GetOauthV1RefreshTokensTokenGetRefreshToken(ctx context.Context, token string) ApiGetOauthV1RefreshTokensTokenGetRefreshTokenRequest {
	return ApiGetOauthV1RefreshTokensTokenGetRefreshTokenRequest{
		ApiService: a,
		ctx:        ctx,
		token:      token,
	}
}

// Execute executes the request
//  @return RefreshTokenInfoResponse
func (a *RefreshTokensApiService) GetOauthV1RefreshTokensTokenGetRefreshTokenExecute(r ApiGetOauthV1RefreshTokensTokenGetRefreshTokenRequest) (*RefreshTokenInfoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RefreshTokenInfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RefreshTokensApiService.GetOauthV1RefreshTokensTokenGetRefreshToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/v1/refresh-tokens/{token}"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterToString(r.token, "")), -1)

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