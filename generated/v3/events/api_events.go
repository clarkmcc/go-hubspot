/*
HubSpot Events API

API for accessing CRM object events.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package events

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/clarkmcc/go-hubspot"
	"net/url"
	"reflect"
	"time"
)

// EventsApiService EventsApi service
type EventsApiService service

type ApiGetPageRequest struct {
	ctx            context.Context
	ApiService     *EventsApiService
	occurredAfter  *time.Time
	occurredBefore *time.Time
	objectType     *string
	objectId       *int64
	eventType      *string
	after          *string
	before         *string
	limit          *int32
	sort           *[]string
}

// The starting time as an ISO 8601 timestamp.
func (r ApiGetPageRequest) OccurredAfter(occurredAfter time.Time) ApiGetPageRequest {
	r.occurredAfter = &occurredAfter
	return r
}

// The ending time as an ISO 8601 timestamp.
func (r ApiGetPageRequest) OccurredBefore(occurredBefore time.Time) ApiGetPageRequest {
	r.occurredBefore = &occurredBefore
	return r
}

// The type of object being selected. Valid values are hubspot named object types (e.g. &#x60;contact&#x60;).
func (r ApiGetPageRequest) ObjectType(objectType string) ApiGetPageRequest {
	r.objectType = &objectType
	return r
}

// The id of the selected object. If not present, then the &#x60;objectProperty&#x60; parameter is required.
func (r ApiGetPageRequest) ObjectId(objectId int64) ApiGetPageRequest {
	r.objectId = &objectId
	return r
}

// Limits the response to the specified event type.  For example &#x60;&amp;eventType&#x3D;e_visited_page&#x60; returns only &#x60;e_visited_page&#x60; events.  If not present all event types are returned.
func (r ApiGetPageRequest) EventType(eventType string) ApiGetPageRequest {
	r.eventType = &eventType
	return r
}

// An additional parameter that may be used to get the next &#x60;limit&#x60; set of results.
func (r ApiGetPageRequest) After(after string) ApiGetPageRequest {
	r.after = &after
	return r
}

func (r ApiGetPageRequest) Before(before string) ApiGetPageRequest {
	r.before = &before
	return r
}

// The maximum number of events to return, defaults to 20.
func (r ApiGetPageRequest) Limit(limit int32) ApiGetPageRequest {
	r.limit = &limit
	return r
}

// Selects the sort field and order. Defaults to ascending, prefix with &#x60;-&#x60; for descending order. &#x60;occurredAt&#x60; is the only field supported for sorting.
func (r ApiGetPageRequest) Sort(sort []string) ApiGetPageRequest {
	r.sort = &sort
	return r
}

func (r ApiGetPageRequest) Execute() (*CollectionResponseExternalUnifiedEvent, *http.Response, error) {
	return r.ApiService.GetPageExecute(r)
}

/*
GetPage Returns a collection of events matching a query.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPageRequest
*/
func (a *EventsApiService) GetPage(ctx context.Context) ApiGetPageRequest {
	return ApiGetPageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return CollectionResponseExternalUnifiedEvent
func (a *EventsApiService) GetPageExecute(r ApiGetPageRequest) (*CollectionResponseExternalUnifiedEvent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionResponseExternalUnifiedEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.GetPage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/v3/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.occurredAfter != nil {
		localVarQueryParams.Add("occurredAfter", parameterToString(*r.occurredAfter, ""))
	}
	if r.occurredBefore != nil {
		localVarQueryParams.Add("occurredBefore", parameterToString(*r.occurredBefore, ""))
	}
	if r.objectType != nil {
		localVarQueryParams.Add("objectType", parameterToString(*r.objectType, ""))
	}
	if r.objectId != nil {
		localVarQueryParams.Add("objectId", parameterToString(*r.objectId, ""))
	}
	if r.eventType != nil {
		localVarQueryParams.Add("eventType", parameterToString(*r.eventType, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.before != nil {
		localVarQueryParams.Add("before", parameterToString(*r.before, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.sort != nil {
		t := *r.sort
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("sort", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("sort", parameterToString(t, "multi"))
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