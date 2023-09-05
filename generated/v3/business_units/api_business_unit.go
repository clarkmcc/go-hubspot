/*
Business Unit

Retrieve Business Unit information.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package business_units

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// BusinessUnitApiService BusinessUnitApi service
type BusinessUnitApiService service

type ApiGetBusinessUnitsV3BusinessUnitsUserUserIdRequest struct {
	ctx        context.Context
	ApiService *BusinessUnitApiService
	userId     string
	properties *[]string
	name       *[]string
}

// The names of properties to optionally include in the response body. The only valid value is &#x60;logoMetadata&#x60;.
func (r ApiGetBusinessUnitsV3BusinessUnitsUserUserIdRequest) Properties(properties []string) ApiGetBusinessUnitsV3BusinessUnitsUserUserIdRequest {
	r.properties = &properties
	return r
}

// The names of Business Units to retrieve. If empty or not provided, then all associated Business Units will be returned.
func (r ApiGetBusinessUnitsV3BusinessUnitsUserUserIdRequest) Name(name []string) ApiGetBusinessUnitsV3BusinessUnitsUserUserIdRequest {
	r.name = &name
	return r
}

func (r ApiGetBusinessUnitsV3BusinessUnitsUserUserIdRequest) Execute() (*CollectionResponsePublicBusinessUnitNoPaging, *http.Response, error) {
	return r.ApiService.GetBusinessUnitsV3BusinessUnitsUserUserIdExecute(r)
}

/*
GetBusinessUnitsV3BusinessUnitsUserUserId Get Business Units for a user

Get Business Units identified by `userId`. The `userId` refers to the user’s ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId Identifier of user to retrieve.
 @return ApiGetBusinessUnitsV3BusinessUnitsUserUserIdRequest
*/
func (a *BusinessUnitApiService) GetBusinessUnitsV3BusinessUnitsUserUserId(ctx context.Context, userId string) ApiGetBusinessUnitsV3BusinessUnitsUserUserIdRequest {
	return ApiGetBusinessUnitsV3BusinessUnitsUserUserIdRequest{
		ApiService: a,
		ctx:        ctx,
		userId:     userId,
	}
}

// Execute executes the request
//  @return CollectionResponsePublicBusinessUnitNoPaging
func (a *BusinessUnitApiService) GetBusinessUnitsV3BusinessUnitsUserUserIdExecute(r ApiGetBusinessUnitsV3BusinessUnitsUserUserIdRequest) (*CollectionResponsePublicBusinessUnitNoPaging, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionResponsePublicBusinessUnitNoPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BusinessUnitApiService.GetBusinessUnitsV3BusinessUnitsUserUserId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/business-units/v3/business-units/user/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.properties != nil {
		t := *r.properties
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("properties", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("properties", parameterToString(t, "multi"))
		}
	}
	if r.name != nil {
		t := *r.name
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("name", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("name", parameterToString(t, "multi"))
		}
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
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
