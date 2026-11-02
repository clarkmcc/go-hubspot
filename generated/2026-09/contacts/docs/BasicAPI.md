# \BasicAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectId**](BasicAPI.md#DeleteCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectId) | **Delete** /crm/objects/2026-09/contacts/{contactId} | Archive a contact
[**DeleteCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeToObjectId**](BasicAPI.md#DeleteCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeToObjectId) | **Delete** /crm/objects/2026-09/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId} | Delete
[**GetCrmObjects202609ContactsContactIdAssociationsToObjectTypeCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectType**](BasicAPI.md#GetCrmObjects202609ContactsContactIdAssociationsToObjectTypeCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectType) | **Get** /crm/objects/2026-09/contacts/{contactId}/associations/{toObjectType} | List associations
[**GetCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectId**](BasicAPI.md#GetCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectId) | **Get** /crm/objects/2026-09/contacts/{contactId} | Retrieve a contact
[**GetCrmObjects202609ContactsCrmObjects202609BetaObjectType**](BasicAPI.md#GetCrmObjects202609ContactsCrmObjects202609BetaObjectType) | **Get** /crm/objects/2026-09/contacts | Retrieve contacts
[**PatchCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectId**](BasicAPI.md#PatchCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectId) | **Patch** /crm/objects/2026-09/contacts/{contactId} | Update a contact
[**PostCrmObjects202609ContactsCrmObjects202609BetaObjectType**](BasicAPI.md#PostCrmObjects202609ContactsCrmObjects202609BetaObjectType) | **Post** /crm/objects/2026-09/contacts | Create a contact
[**PutCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeToObjectId**](BasicAPI.md#PutCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeToObjectId) | **Put** /crm/objects/2026-09/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId} | Create



## DeleteCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectId

> DeleteCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectId(ctx, contactId).Execute()

Archive a contact



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
	contactId := "contactId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicAPI.DeleteCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectId(context.Background(), contactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.DeleteCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectIdRequest struct via the builder pattern


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


## DeleteCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeToObjectId

> DeleteCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeToObjectId(ctx, objectId, objectType, toObjectId, toObjectType).Execute()

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
	objectId := "objectId_example" // string | The unique identifier of the first object.
	objectType := "objectType_example" // string | The type of the first object involved in the association.
	toObjectId := "toObjectId_example" // string | The unique identifier of the second object.
	toObjectType := "toObjectType_example" // string | The type of the second object involved in the association.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicAPI.DeleteCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeToObjectId(context.Background(), objectId, objectType, toObjectId, toObjectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.DeleteCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeToObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The unique identifier of the first object. | 
**objectType** | **string** | The type of the first object involved in the association. | 
**toObjectId** | **string** | The unique identifier of the second object. | 
**toObjectType** | **string** | The type of the second object involved in the association. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeToObjectIdRequest struct via the builder pattern


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


## GetCrmObjects202609ContactsContactIdAssociationsToObjectTypeCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectType

> CollectionResponseMultiAssociatedObjectWithLabelForwardPaging GetCrmObjects202609ContactsContactIdAssociationsToObjectTypeCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectType(ctx, contactId, toObjectType).After(after).Limit(limit).Execute()

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
	contactId := "contactId_example" // string | 
	toObjectType := "toObjectType_example" // string | The type of the associated objects to retrieve.
	after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetCrmObjects202609ContactsContactIdAssociationsToObjectTypeCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectType(context.Background(), contactId, toObjectType).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetCrmObjects202609ContactsContactIdAssociationsToObjectTypeCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmObjects202609ContactsContactIdAssociationsToObjectTypeCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectType`: CollectionResponseMultiAssociatedObjectWithLabelForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetCrmObjects202609ContactsContactIdAssociationsToObjectTypeCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 
**toObjectType** | **string** | The type of the associated objects to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmObjects202609ContactsContactIdAssociationsToObjectTypeCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeRequest struct via the builder pattern


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


## GetCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectId

> SimplePublicObjectWithAssociations GetCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectId(ctx, contactId).Archived(archived).Associations(associations).IdProperty(idProperty).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()

Retrieve a contact



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
	contactId := "contactId_example" // string | 
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)
	associations := []string{"Inner_example"} // []string | A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored. (optional)
	idProperty := "idProperty_example" // string | The name of a property whose values are unique for this object type (optional)
	properties := []string{"Inner_example"} // []string | A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored. (optional)
	propertiesWithHistory := []string{"Inner_example"} // []string | A comma separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectId(context.Background(), contactId).Archived(archived).Associations(associations).IdProperty(idProperty).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectId`: SimplePublicObjectWithAssociations
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectIdRequest struct via the builder pattern


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


## GetCrmObjects202609ContactsCrmObjects202609BetaObjectType

> CollectionResponseSimplePublicObjectWithAssociationsForwardPaging GetCrmObjects202609ContactsCrmObjects202609BetaObjectType(ctx).After(after).Archived(archived).Associations(associations).Limit(limit).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()

Retrieve contacts



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
	propertiesWithHistory := []string{"Inner_example"} // []string | A comma separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored. Usage of this parameter will reduce the maximum number of contacts that can be read by a single request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetCrmObjects202609ContactsCrmObjects202609BetaObjectType(context.Background()).After(after).Archived(archived).Associations(associations).Limit(limit).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetCrmObjects202609ContactsCrmObjects202609BetaObjectType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmObjects202609ContactsCrmObjects202609BetaObjectType`: CollectionResponseSimplePublicObjectWithAssociationsForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetCrmObjects202609ContactsCrmObjects202609BetaObjectType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmObjects202609ContactsCrmObjects202609BetaObjectTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]
 **associations** | **[]string** | A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored. | 
 **limit** | **int32** | The maximum number of results to display per page. | [default to 10]
 **properties** | **[]string** | A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored. | 
 **propertiesWithHistory** | **[]string** | A comma separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored. Usage of this parameter will reduce the maximum number of contacts that can be read by a single request. | 

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


## PatchCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectId

> SimplePublicObject PatchCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectId(ctx, contactId).SimplePublicObjectInput(simplePublicObjectInput).IdProperty(idProperty).Execute()

Update a contact



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
	contactId := "contactId_example" // string | 
	simplePublicObjectInput := *openapiclient.NewSimplePublicObjectInput(map[string]string{"key": "Inner_example"}) // SimplePublicObjectInput | 
	idProperty := "idProperty_example" // string | The name of a property whose values are unique for this object type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PatchCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectId(context.Background(), contactId).SimplePublicObjectInput(simplePublicObjectInput).IdProperty(idProperty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PatchCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectId`: SimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PatchCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCrmObjects202609ContactsContactIdCrmObjects202609BetaObjectTypeObjectIdRequest struct via the builder pattern


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


## PostCrmObjects202609ContactsCrmObjects202609BetaObjectType

> SimplePublicObject PostCrmObjects202609ContactsCrmObjects202609BetaObjectType(ctx).SimplePublicObjectInputForCreate(simplePublicObjectInputForCreate).Execute()

Create a contact



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
	resp, r, err := apiClient.BasicAPI.PostCrmObjects202609ContactsCrmObjects202609BetaObjectType(context.Background()).SimplePublicObjectInputForCreate(simplePublicObjectInputForCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostCrmObjects202609ContactsCrmObjects202609BetaObjectType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmObjects202609ContactsCrmObjects202609BetaObjectType`: SimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PostCrmObjects202609ContactsCrmObjects202609BetaObjectType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmObjects202609ContactsCrmObjects202609BetaObjectTypeRequest struct via the builder pattern


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


## PutCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeToObjectId

> LabelsBetweenObjectPair PutCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeToObjectId(ctx, objectId, objectType, toObjectId, toObjectType).AssociationSpec(associationSpec).Execute()

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
	objectId := "objectId_example" // string | The unique identifier of the first object involved in the association.
	objectType := "objectType_example" // string | The type of the first object involved in the association.
	toObjectId := "toObjectId_example" // string | The unique identifier of the second object involved in the association.
	toObjectType := "toObjectType_example" // string | The type of the second object involved in the association.
	associationSpec := []openapiclient.AssociationSpec{*openapiclient.NewAssociationSpec("AssociationCategory_example", int32(123))} // []AssociationSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PutCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeToObjectId(context.Background(), objectId, objectType, toObjectId, toObjectType).AssociationSpec(associationSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PutCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeToObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeToObjectId`: LabelsBetweenObjectPair
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PutCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeToObjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The unique identifier of the first object involved in the association. | 
**objectType** | **string** | The type of the first object involved in the association. | 
**toObjectId** | **string** | The unique identifier of the second object involved in the association. | 
**toObjectType** | **string** | The type of the second object involved in the association. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeToObjectIdRequest struct via the builder pattern


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

