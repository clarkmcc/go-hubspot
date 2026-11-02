# \BasicAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteContacts**](BasicAPI.md#DeleteContacts) | **Delete** /crm/objects/2026-09/contacts/{contactId} | Archive a contact
[**DeleteRoot**](BasicAPI.md#DeleteRoot) | **Delete** /crm/objects/2026-09/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId} | Delete
[**GetContacts**](BasicAPI.md#GetContacts) | **Get** /crm/objects/2026-09/contacts | Retrieve contacts
[**GetContactsByContactId**](BasicAPI.md#GetContactsByContactId) | **Get** /crm/objects/2026-09/contacts/{contactId} | Retrieve a contact
[**GetContactsByContactIdByToObjectType**](BasicAPI.md#GetContactsByContactIdByToObjectType) | **Get** /crm/objects/2026-09/contacts/{contactId}/associations/{toObjectType} | List associations
[**PatchContacts**](BasicAPI.md#PatchContacts) | **Patch** /crm/objects/2026-09/contacts/{contactId} | Update a contact
[**PostContacts**](BasicAPI.md#PostContacts) | **Post** /crm/objects/2026-09/contacts | Create a contact
[**PutRoot**](BasicAPI.md#PutRoot) | **Put** /crm/objects/2026-09/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId} | Create



## DeleteContacts

> DeleteContacts(ctx, contactId).Execute()

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
	r, err := apiClient.BasicAPI.DeleteContacts(context.Background(), contactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.DeleteContacts``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteContactsRequest struct via the builder pattern


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
	objectId := "objectId_example" // string | The unique identifier of the first object.
	objectType := "objectType_example" // string | The type of the first object involved in the association.
	toObjectId := "toObjectId_example" // string | The unique identifier of the second object.
	toObjectType := "toObjectType_example" // string | The type of the second object involved in the association.

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
**objectId** | **string** | The unique identifier of the first object. | 
**objectType** | **string** | The type of the first object involved in the association. | 
**toObjectId** | **string** | The unique identifier of the second object. | 
**toObjectType** | **string** | The type of the second object involved in the association. | 

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


## GetContacts

> CollectionResponseSimplePublicObjectWithAssociationsForwardPaging GetContacts(ctx).After(after).Archived(archived).Associations(associations).Limit(limit).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()

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
	resp, r, err := apiClient.BasicAPI.GetContacts(context.Background()).After(after).Archived(archived).Associations(associations).Limit(limit).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContacts`: CollectionResponseSimplePublicObjectWithAssociationsForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContactsRequest struct via the builder pattern


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


## GetContactsByContactId

> SimplePublicObjectWithAssociations GetContactsByContactId(ctx, contactId).Archived(archived).Associations(associations).IdProperty(idProperty).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()

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
	resp, r, err := apiClient.BasicAPI.GetContactsByContactId(context.Background(), contactId).Archived(archived).Associations(associations).IdProperty(idProperty).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetContactsByContactId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactsByContactId`: SimplePublicObjectWithAssociations
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetContactsByContactId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactsByContactIdRequest struct via the builder pattern


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


## GetContactsByContactIdByToObjectType

> CollectionResponseMultiAssociatedObjectWithLabelForwardPaging GetContactsByContactIdByToObjectType(ctx, contactId, toObjectType).After(after).Limit(limit).Execute()

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
	resp, r, err := apiClient.BasicAPI.GetContactsByContactIdByToObjectType(context.Background(), contactId, toObjectType).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetContactsByContactIdByToObjectType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactsByContactIdByToObjectType`: CollectionResponseMultiAssociatedObjectWithLabelForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetContactsByContactIdByToObjectType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 
**toObjectType** | **string** | The type of the associated objects to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactsByContactIdByToObjectTypeRequest struct via the builder pattern


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


## PatchContacts

> SimplePublicObject PatchContacts(ctx, contactId).SimplePublicObjectInput(simplePublicObjectInput).IdProperty(idProperty).Execute()

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
	resp, r, err := apiClient.BasicAPI.PatchContacts(context.Background(), contactId).SimplePublicObjectInput(simplePublicObjectInput).IdProperty(idProperty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PatchContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchContacts`: SimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PatchContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchContactsRequest struct via the builder pattern


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


## PostContacts

> SimplePublicObject PostContacts(ctx).SimplePublicObjectInputForCreate(simplePublicObjectInputForCreate).Execute()

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
	resp, r, err := apiClient.BasicAPI.PostContacts(context.Background()).SimplePublicObjectInputForCreate(simplePublicObjectInputForCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostContacts`: SimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PostContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostContactsRequest struct via the builder pattern


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
**objectId** | **string** | The unique identifier of the first object involved in the association. | 
**objectType** | **string** | The type of the first object involved in the association. | 
**toObjectId** | **string** | The unique identifier of the second object involved in the association. | 
**toObjectType** | **string** | The type of the second object involved in the association. | 

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

