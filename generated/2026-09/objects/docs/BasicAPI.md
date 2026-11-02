# \BasicAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRoot**](BasicAPI.md#DeleteRoot) | **Delete** /crm/objects/2026-09/{objectType}/{objectId} | Archive
[**GetByObjectTypeByObjectId**](BasicAPI.md#GetByObjectTypeByObjectId) | **Get** /crm/objects/2026-09/{objectType}/{objectId} | Read
[**GetRoot**](BasicAPI.md#GetRoot) | **Get** /crm/objects/2026-09/{objectType} | List
[**PatchRoot**](BasicAPI.md#PatchRoot) | **Patch** /crm/objects/2026-09/{objectType}/{objectId} | Update
[**PostRoot**](BasicAPI.md#PostRoot) | **Post** /crm/objects/2026-09/{objectType} | Create



## DeleteRoot

> DeleteRoot(ctx, objectId, objectType).Execute()

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
	objectId := "objectId_example" // string | The unique identifier of the object to be archived.
	objectType := "objectType_example" // string | The type of the object to be archived.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicAPI.DeleteRoot(context.Background(), objectId, objectType).Execute()
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
**objectId** | **string** | The unique identifier of the object to be archived. | 
**objectType** | **string** | The type of the object to be archived. | 

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


## GetByObjectTypeByObjectId

> SimplePublicObjectWithAssociations GetByObjectTypeByObjectId(ctx, objectId, objectType).Archived(archived).Associations(associations).IdProperty(idProperty).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()

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
	objectId := "objectId_example" // string | The unique identifier of the object to retrieve.
	objectType := "objectType_example" // string | The type of the object to retrieve.
	archived := true // bool | Whether to return only results that have been archived. Defaults to false. (optional) (default to false)
	associations := []string{"Inner_example"} // []string | A comma-separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored. (optional)
	idProperty := "idProperty_example" // string | The name of a property whose values are unique for this object type. (optional)
	properties := []string{"Inner_example"} // []string | A comma-separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored. (optional)
	propertiesWithHistory := []string{"Inner_example"} // []string | A comma-separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetByObjectTypeByObjectId(context.Background(), objectId, objectType).Archived(archived).Associations(associations).IdProperty(idProperty).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetByObjectTypeByObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetByObjectTypeByObjectId`: SimplePublicObjectWithAssociations
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetByObjectTypeByObjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The unique identifier of the object to retrieve. | 
**objectType** | **string** | The type of the object to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByObjectTypeByObjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **archived** | **bool** | Whether to return only results that have been archived. Defaults to false. | [default to false]
 **associations** | **[]string** | A comma-separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored. | 
 **idProperty** | **string** | The name of a property whose values are unique for this object type. | 
 **properties** | **[]string** | A comma-separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored. | 
 **propertiesWithHistory** | **[]string** | A comma-separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored. | 

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


## GetRoot

> CollectionResponseSimplePublicObjectWithAssociationsForwardPaging GetRoot(ctx, objectType).After(after).Archived(archived).Associations(associations).Limit(limit).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()

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
	objectType := "objectType_example" // string | The type of CRM object to retrieve, such as contacts, companies, or deals.
	after := "after_example" // string | A cursor token for pagination. Use the value from the previous response's paging.next.after field. (optional)
	archived := true // bool | Whether to include archived records in the response. Defaults to false. (optional) (default to false)
	associations := []string{"Inner_example"} // []string | A comma-separated list of object types to retrieve associated IDs for. If any specified associations do not exist, they will be ignored. (optional)
	limit := int32(56) // int32 | The maximum number of results to display per page. Defaults to 10. (optional) (default to 10)
	properties := []string{"Inner_example"} // []string | A comma-separated list of property names to include in the response. If any specified properties are not present on the requested object(s), they will be ignored. (optional)
	propertiesWithHistory := []string{"Inner_example"} // []string | A comma-separated list of properties to include in the response along with their history of previous values. If any specified properties are not present on the requested object(s), they will be ignored. Using this parameter reduces the maximum number of objects that can be read in a single request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetRoot(context.Background(), objectType).After(after).Archived(archived).Associations(associations).Limit(limit).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoot`: CollectionResponseSimplePublicObjectWithAssociationsForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetRoot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of CRM object to retrieve, such as contacts, companies, or deals. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | A cursor token for pagination. Use the value from the previous response&#39;s paging.next.after field. | 
 **archived** | **bool** | Whether to include archived records in the response. Defaults to false. | [default to false]
 **associations** | **[]string** | A comma-separated list of object types to retrieve associated IDs for. If any specified associations do not exist, they will be ignored. | 
 **limit** | **int32** | The maximum number of results to display per page. Defaults to 10. | [default to 10]
 **properties** | **[]string** | A comma-separated list of property names to include in the response. If any specified properties are not present on the requested object(s), they will be ignored. | 
 **propertiesWithHistory** | **[]string** | A comma-separated list of properties to include in the response along with their history of previous values. If any specified properties are not present on the requested object(s), they will be ignored. Using this parameter reduces the maximum number of objects that can be read in a single request. | 

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


## PatchRoot

> SimplePublicObject PatchRoot(ctx, objectId, objectType).SimplePublicObjectInput(simplePublicObjectInput).IdProperty(idProperty).Execute()

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
	objectId := "objectId_example" // string | The unique identifier of the object to update.
	objectType := "objectType_example" // string | The type of the object to update.
	simplePublicObjectInput := *openapiclient.NewSimplePublicObjectInput(map[string]string{"key": "Inner_example"}) // SimplePublicObjectInput | 
	idProperty := "idProperty_example" // string | The name of a property whose values are unique for this object type. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PatchRoot(context.Background(), objectId, objectType).SimplePublicObjectInput(simplePublicObjectInput).IdProperty(idProperty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PatchRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRoot`: SimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PatchRoot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The unique identifier of the object to update. | 
**objectType** | **string** | The type of the object to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **simplePublicObjectInput** | [**SimplePublicObjectInput**](SimplePublicObjectInput.md) |  | 
 **idProperty** | **string** | The name of a property whose values are unique for this object type. | 

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


## PostRoot

> SimplePublicObject PostRoot(ctx, objectType).SimplePublicObjectInputForCreate(simplePublicObjectInputForCreate).Execute()

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
	objectType := "objectType_example" // string | The type of CRM object to create, such as contacts, companies, or deals.
	simplePublicObjectInputForCreate := *openapiclient.NewSimplePublicObjectInputForCreate([]openapiclient.PublicAssociationsForObject{*openapiclient.NewPublicAssociationsForObject(*openapiclient.NewPublicObjectId("Id_example"), []openapiclient.AssociationSpec{*openapiclient.NewAssociationSpec("AssociationCategory_example", int32(123))})}, map[string]string{"key": "Inner_example"}) // SimplePublicObjectInputForCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PostRoot(context.Background(), objectType).SimplePublicObjectInputForCreate(simplePublicObjectInputForCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRoot`: SimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PostRoot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of CRM object to create, such as contacts, companies, or deals. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRootRequest struct via the builder pattern


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

