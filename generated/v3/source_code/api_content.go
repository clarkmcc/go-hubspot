/*
CMS Source Code

Endpoints for interacting with files in the CMS Developer File System. These files include HTML templates, CSS, JS, modules, and other assets which are used to create CMS content.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package source_code

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/clarkmcc/go-hubspot"
	"net/url"
	"os"
	"strings"
)

// ContentApiService ContentApi service
type ContentApiService service

type ApiContentArchiveRequest struct {
	ctx         context.Context
	ApiService  *ContentApiService
	environment string
	path        string
}

func (r ApiContentArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.ContentArchiveExecute(r)
}

/*
ContentArchive Delete a file

Deletes the file at the specified path in the specified environment.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environment The environment of the file (\"draft\" or \"published\").
 @param path The file system location of the file.
 @return ApiContentArchiveRequest
*/
func (a *ContentApiService) ContentArchive(ctx context.Context, environment string, path string) ApiContentArchiveRequest {
	return ApiContentArchiveRequest{
		ApiService:  a,
		ctx:         ctx,
		environment: environment,
		path:        path,
	}
}

// Execute executes the request
func (a *ContentApiService) ContentArchiveExecute(r ApiContentArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentApiService.ContentArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/source-code/{environment}/content/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"environment"+"}", url.PathEscape(parameterToString(r.environment, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", url.PathEscape(parameterToString(r.path, "")), -1)

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

type ApiContentCreateRequest struct {
	ctx         context.Context
	ApiService  *ContentApiService
	environment string
	path        string
	file        **os.File
}

// The file to upload.
func (r ApiContentCreateRequest) File(file *os.File) ApiContentCreateRequest {
	r.file = &file
	return r
}

func (r ApiContentCreateRequest) Execute() (*AssetFileMetadata, *http.Response, error) {
	return r.ApiService.ContentCreateExecute(r)
}

/*
ContentCreate Create a file

Creates a file at the specified path in the specified environment. Accepts multipart/form-data content type. Throws an error if a file already exists at the specified path.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environment The environment of the file (\"draft\" or \"published\").
 @param path The file system location of the file.
 @return ApiContentCreateRequest

Deprecated
*/
func (a *ContentApiService) ContentCreate(ctx context.Context, environment string, path string) ApiContentCreateRequest {
	return ApiContentCreateRequest{
		ApiService:  a,
		ctx:         ctx,
		environment: environment,
		path:        path,
	}
}

// Execute executes the request
//  @return AssetFileMetadata
// Deprecated
func (a *ContentApiService) ContentCreateExecute(r ApiContentCreateRequest) (*AssetFileMetadata, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AssetFileMetadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentApiService.ContentCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/source-code/{environment}/content/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"environment"+"}", url.PathEscape(parameterToString(r.environment, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", url.PathEscape(parameterToString(r.path, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	var fileLocalVarFormFileName string
	var fileLocalVarFileName string
	var fileLocalVarFileBytes []byte

	fileLocalVarFormFileName = "file"

	var fileLocalVarFile *os.File
	if r.file != nil {
		fileLocalVarFile = *r.file
	}
	if fileLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(fileLocalVarFile)
		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
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

type ApiGetCmsV3SourceCodeEnvironmentContentPathDownloadRequest struct {
	ctx         context.Context
	ApiService  *ContentApiService
	environment string
	path        string
}

func (r ApiGetCmsV3SourceCodeEnvironmentContentPathDownloadRequest) Execute() (*Error, *http.Response, error) {
	return r.ApiService.GetCmsV3SourceCodeEnvironmentContentPathDownloadExecute(r)
}

/*
GetCmsV3SourceCodeEnvironmentContentPathDownload Download a file

Downloads the byte contents of the file at the specified path in the specified environment.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environment The environment of the file (\"draft\" or \"published\").
 @param path The file system location of the file.
 @return ApiGetCmsV3SourceCodeEnvironmentContentPathDownloadRequest
*/
func (a *ContentApiService) GetCmsV3SourceCodeEnvironmentContentPathDownload(ctx context.Context, environment string, path string) ApiGetCmsV3SourceCodeEnvironmentContentPathDownloadRequest {
	return ApiGetCmsV3SourceCodeEnvironmentContentPathDownloadRequest{
		ApiService:  a,
		ctx:         ctx,
		environment: environment,
		path:        path,
	}
}

// Execute executes the request
//  @return Error
func (a *ContentApiService) GetCmsV3SourceCodeEnvironmentContentPathDownloadExecute(r ApiGetCmsV3SourceCodeEnvironmentContentPathDownloadRequest) (*Error, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Error
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentApiService.GetCmsV3SourceCodeEnvironmentContentPathDownload")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/source-code/{environment}/content/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"environment"+"}", url.PathEscape(parameterToString(r.environment, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", url.PathEscape(parameterToString(r.path, "")), -1)

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

type ApiPutCmsV3SourceCodeEnvironmentContentPathCreateOrUpdateRequest struct {
	ctx         context.Context
	ApiService  *ContentApiService
	environment string
	path        string
	file        **os.File
}

// The file to upload.
func (r ApiPutCmsV3SourceCodeEnvironmentContentPathCreateOrUpdateRequest) File(file *os.File) ApiPutCmsV3SourceCodeEnvironmentContentPathCreateOrUpdateRequest {
	r.file = &file
	return r
}

func (r ApiPutCmsV3SourceCodeEnvironmentContentPathCreateOrUpdateRequest) Execute() (*AssetFileMetadata, *http.Response, error) {
	return r.ApiService.PutCmsV3SourceCodeEnvironmentContentPathCreateOrUpdateExecute(r)
}

/*
PutCmsV3SourceCodeEnvironmentContentPathCreateOrUpdate Create or update a file

Upserts a file at the specified path in the specified environment. Accepts multipart/form-data content type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environment The environment of the file (\"draft\" or \"published\").
 @param path The file system location of the file.
 @return ApiPutCmsV3SourceCodeEnvironmentContentPathCreateOrUpdateRequest
*/
func (a *ContentApiService) PutCmsV3SourceCodeEnvironmentContentPathCreateOrUpdate(ctx context.Context, environment string, path string) ApiPutCmsV3SourceCodeEnvironmentContentPathCreateOrUpdateRequest {
	return ApiPutCmsV3SourceCodeEnvironmentContentPathCreateOrUpdateRequest{
		ApiService:  a,
		ctx:         ctx,
		environment: environment,
		path:        path,
	}
}

// Execute executes the request
//  @return AssetFileMetadata
func (a *ContentApiService) PutCmsV3SourceCodeEnvironmentContentPathCreateOrUpdateExecute(r ApiPutCmsV3SourceCodeEnvironmentContentPathCreateOrUpdateRequest) (*AssetFileMetadata, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AssetFileMetadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentApiService.PutCmsV3SourceCodeEnvironmentContentPathCreateOrUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/source-code/{environment}/content/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"environment"+"}", url.PathEscape(parameterToString(r.environment, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", url.PathEscape(parameterToString(r.path, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	var fileLocalVarFormFileName string
	var fileLocalVarFileName string
	var fileLocalVarFileBytes []byte

	fileLocalVarFormFileName = "file"

	var fileLocalVarFile *os.File
	if r.file != nil {
		fileLocalVarFile = *r.file
	}
	if fileLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(fileLocalVarFile)
		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
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
