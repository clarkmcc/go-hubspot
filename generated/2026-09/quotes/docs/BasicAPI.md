# \BasicAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteQuotes**](BasicAPI.md#DeleteQuotes) | **Delete** /crm/objects/2026-09/quotes/{quoteId} | Archive
[**DeleteRoot**](BasicAPI.md#DeleteRoot) | **Delete** /crm/objects/2026-09/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId} | Delete
[**GetQuotes**](BasicAPI.md#GetQuotes) | **Get** /crm/objects/2026-09/quotes | List
[**GetQuotesByQuoteId**](BasicAPI.md#GetQuotesByQuoteId) | **Get** /crm/objects/2026-09/quotes/{quoteId} | Read
[**GetQuotesByQuoteIdByToObjectType**](BasicAPI.md#GetQuotesByQuoteIdByToObjectType) | **Get** /crm/objects/2026-09/quotes/{quoteId}/associations/{toObjectType} | List associations
[**PatchQuotes**](BasicAPI.md#PatchQuotes) | **Patch** /crm/objects/2026-09/quotes/{quoteId} | Update
[**PostQuotes**](BasicAPI.md#PostQuotes) | **Post** /crm/objects/2026-09/quotes | Create
[**PutRoot**](BasicAPI.md#PutRoot) | **Put** /crm/objects/2026-09/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId} | Create Association



## DeleteQuotes

> DeleteQuotes(ctx, quoteId).Execute()

Archive



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	quoteId := "quoteId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicAPI.DeleteQuotes(context.Background(), quoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.DeleteQuotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoot

> DeleteRoot(ctx, objectId, objectType, toObjectId, toObjectType).Execute()

Delete



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	objectId := "objectId_example" // string | The unique identifier of the source object from which the association is being removed.
	objectType := "objectType_example" // string | The type of the source object from which the association is being removed.
	toObjectId := "toObjectId_example" // string | The unique identifier of the target object to which the association is being removed.
	toObjectType := "toObjectType_example" // string | The type of the target object to which the association is being removed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicAPI.DeleteRoot(context.Background(), objectId, objectType, toObjectId, toObjectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.DeleteRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The unique identifier of the source object from which the association is being removed. | 
**objectType** | **string** | The type of the source object from which the association is being removed. | 
**toObjectId** | **string** | The unique identifier of the target object to which the association is being removed. | 
**toObjectType** | **string** | The type of the target object to which the association is being removed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuotes

> CollectionResponseSimplePublicObjectWithAssociationsForwardPaging GetQuotes(ctx).After(after).Archived(archived).Associations(associations).Limit(limit).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()

List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)
	associations := []string{"Inner_example"} // []string | A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored. (optional)
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional) (default to 10)
	properties := []string{"Inner_example"} // []string | A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored. (optional)
	propertiesWithHistory := []string{"Inner_example"} // []string | A comma separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored. Usage of this parameter will reduce the maximum number of objects that can be read by a single request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetQuotes(context.Background()).After(after).Archived(archived).Associations(associations).Limit(limit).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetQuotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuotes`: CollectionResponseSimplePublicObjectWithAssociationsForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetQuotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]
 **associations** | **[]string** | A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored. | 
 **limit** | **int32** | The maximum number of results to display per page. | [default to 10]
 **properties** | **[]string** | A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored. | 
 **propertiesWithHistory** | **[]string** | A comma separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored. Usage of this parameter will reduce the maximum number of objects that can be read by a single request. | 

### Return type

[**CollectionResponseSimplePublicObjectWithAssociationsForwardPaging**](CollectionResponseSimplePublicObjectWithAssociationsForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuotesByQuoteId

> SimplePublicObjectWithAssociations GetQuotesByQuoteId(ctx, quoteId).Archived(archived).Associations(associations).IdProperty(idProperty).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()

Read



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	quoteId := "quoteId_example" // string | 
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)
	associations := []string{"Inner_example"} // []string | A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored. (optional)
	idProperty := "idProperty_example" // string | The name of a property whose values are unique for this object type (optional)
	properties := []string{"Inner_example"} // []string | A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored. (optional)
	propertiesWithHistory := []string{"Inner_example"} // []string | A comma separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetQuotesByQuoteId(context.Background(), quoteId).Archived(archived).Associations(associations).IdProperty(idProperty).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetQuotesByQuoteId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuotesByQuoteId`: SimplePublicObjectWithAssociations
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetQuotesByQuoteId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuotesByQuoteIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]
 **associations** | **[]string** | A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored. | 
 **idProperty** | **string** | The name of a property whose values are unique for this object type | 
 **properties** | **[]string** | A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored. | 
 **propertiesWithHistory** | **[]string** | A comma separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored. | 

### Return type

[**SimplePublicObjectWithAssociations**](SimplePublicObjectWithAssociations.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuotesByQuoteIdByToObjectType

> CollectionResponseMultiAssociatedObjectWithLabelForwardPaging GetQuotesByQuoteIdByToObjectType(ctx, quoteId, toObjectType).After(after).Limit(limit).Execute()

List associations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	quoteId := "quoteId_example" // string | 
	toObjectType := "toObjectType_example" // string | The type of the associated objects to retrieve.
	after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetQuotesByQuoteIdByToObjectType(context.Background(), quoteId, toObjectType).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetQuotesByQuoteIdByToObjectType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuotesByQuoteIdByToObjectType`: CollectionResponseMultiAssociatedObjectWithLabelForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetQuotesByQuoteIdByToObjectType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 
**toObjectType** | **string** | The type of the associated objects to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuotesByQuoteIdByToObjectTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to display per page. | [default to 500]

### Return type

[**CollectionResponseMultiAssociatedObjectWithLabelForwardPaging**](CollectionResponseMultiAssociatedObjectWithLabelForwardPaging.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchQuotes

> SimplePublicObject PatchQuotes(ctx, quoteId).SimplePublicObjectInput(simplePublicObjectInput).IdProperty(idProperty).Execute()

Update



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	quoteId := "quoteId_example" // string | 
	simplePublicObjectInput := *openapiclient.NewSimplePublicObjectInput(map[string]string{"key": "Inner_example"}) // SimplePublicObjectInput | 
	idProperty := "idProperty_example" // string | The name of a property whose values are unique for this object type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PatchQuotes(context.Background(), quoteId).SimplePublicObjectInput(simplePublicObjectInput).IdProperty(idProperty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PatchQuotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchQuotes`: SimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PatchQuotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **simplePublicObjectInput** | [**SimplePublicObjectInput**](SimplePublicObjectInput.md) |  | 
 **idProperty** | **string** | The name of a property whose values are unique for this object type | 

### Return type

[**SimplePublicObject**](SimplePublicObject.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostQuotes

> SimplePublicObject PostQuotes(ctx).SimplePublicObjectInputForCreate(simplePublicObjectInputForCreate).Execute()

Create



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	simplePublicObjectInputForCreate := *openapiclient.NewSimplePublicObjectInputForCreate([]openapiclient.PublicAssociationsForObject{*openapiclient.NewPublicAssociationsForObject(*openapiclient.NewPublicObjectId("Id_example"), []openapiclient.AssociationSpec{*openapiclient.NewAssociationSpec("AssociationCategory_example", int32(123))})}, map[string]string{"key": "Inner_example"}) // SimplePublicObjectInputForCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PostQuotes(context.Background()).SimplePublicObjectInputForCreate(simplePublicObjectInputForCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostQuotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQuotes`: SimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PostQuotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **simplePublicObjectInputForCreate** | [**SimplePublicObjectInputForCreate**](SimplePublicObjectInputForCreate.md) |  | 

### Return type

[**SimplePublicObject**](SimplePublicObject.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRoot

> LabelsBetweenObjectPair PutRoot(ctx, objectId, objectType, toObjectId, toObjectType).AssociationSpec(associationSpec).Execute()

Create Association



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	objectId := "objectId_example" // string | The unique identifier of the source object for the association.
	objectType := "objectType_example" // string | The type of the source object for the association.
	toObjectId := "toObjectId_example" // string | The unique identifier of the target object for the association.
	toObjectType := "toObjectType_example" // string | The type of the target object for the association.
	associationSpec := []openapiclient.AssociationSpec{*openapiclient.NewAssociationSpec("AssociationCategory_example", int32(123))} // []AssociationSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PutRoot(context.Background(), objectId, objectType, toObjectId, toObjectType).AssociationSpec(associationSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PutRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRoot`: LabelsBetweenObjectPair
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PutRoot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The unique identifier of the source object for the association. | 
**objectType** | **string** | The type of the source object for the association. | 
**toObjectId** | **string** | The unique identifier of the target object for the association. | 
**toObjectType** | **string** | The type of the target object for the association. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **associationSpec** | [**[]AssociationSpec**](AssociationSpec.md) |  | 

### Return type

[**LabelsBetweenObjectPair**](LabelsBetweenObjectPair.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

