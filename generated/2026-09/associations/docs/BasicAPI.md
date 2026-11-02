# \BasicAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteByObjectTypeByObjectIdByToObjectTypeByToObjectId**](BasicAPI.md#DeleteByObjectTypeByObjectIdByToObjectTypeByToObjectId) | **Delete** /crm/objects/2026-09/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId} | Delete associations between two records
[**DeleteLabels**](BasicAPI.md#DeleteLabels) | **Delete** /crm/associations/2026-09/{fromObjectType}/{toObjectType}/labels/{associationTypeId} | Delete
[**DeleteRoot**](BasicAPI.md#DeleteRoot) | **Delete** /crm/objects/2026-09/{objectType}/{objectId} | Archive
[**GetByObjectTypeByObjectId**](BasicAPI.md#GetByObjectTypeByObjectId) | **Get** /crm/objects/2026-09/{objectType}/{objectId} | Read
[**GetByObjectTypeByObjectIdByToObjectType**](BasicAPI.md#GetByObjectTypeByObjectIdByToObjectType) | **Get** /crm/objects/2026-09/{objectType}/{objectId}/associations/{toObjectType} | Retrieve all associations by object type
[**GetDefinitionsConfigurations**](BasicAPI.md#GetDefinitionsConfigurations) | **Get** /crm/associations/2026-09/definitions/configurations/{fromObjectType}/{toObjectType} | Read
[**GetDefinitionsConfigurationsAll**](BasicAPI.md#GetDefinitionsConfigurationsAll) | **Get** /crm/associations/2026-09/definitions/configurations/all | Read All
[**GetLabels**](BasicAPI.md#GetLabels) | **Get** /crm/associations/2026-09/{fromObjectType}/{toObjectType}/labels | Read
[**GetRoot**](BasicAPI.md#GetRoot) | **Get** /crm/objects/2026-09/{objectType} | List
[**PatchRoot**](BasicAPI.md#PatchRoot) | **Patch** /crm/objects/2026-09/{objectType}/{objectId} | Update
[**PostLabels**](BasicAPI.md#PostLabels) | **Post** /crm/associations/2026-09/{fromObjectType}/{toObjectType}/labels | Create
[**PostRoot**](BasicAPI.md#PostRoot) | **Post** /crm/objects/2026-09/{objectType} | Create
[**PostUsageHighUsageReport**](BasicAPI.md#PostUsageHighUsageReport) | **Post** /crm/associations/2026-09/usage/high-usage-report/{userId} | Report high usage
[**PutLabels**](BasicAPI.md#PutLabels) | **Put** /crm/associations/2026-09/{fromObjectType}/{toObjectType}/labels | Update
[**PutRoot**](BasicAPI.md#PutRoot) | **Put** /crm/objects/2026-09/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId} | Associate records (labelled)



## DeleteByObjectTypeByObjectIdByToObjectTypeByToObjectId

> DeleteByObjectTypeByObjectIdByToObjectTypeByToObjectId(ctx, objectId, objectType, toObjectId, toObjectType).Execute()

Delete associations between two records



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
	objectId := "objectId_example" // string | The unique identifier of the initial object from which the association will be removed.
	objectType := "objectType_example" // string | The type of the initial object from which the association will be removed.
	toObjectId := "toObjectId_example" // string | The unique identifier of the target object to which the association will be removed.
	toObjectType := "toObjectType_example" // string | The type of the target object to which the association will be removed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicAPI.DeleteByObjectTypeByObjectIdByToObjectTypeByToObjectId(context.Background(), objectId, objectType, toObjectId, toObjectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.DeleteByObjectTypeByObjectIdByToObjectTypeByToObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The unique identifier of the initial object from which the association will be removed. | 
**objectType** | **string** | The type of the initial object from which the association will be removed. | 
**toObjectId** | **string** | The unique identifier of the target object to which the association will be removed. | 
**toObjectType** | **string** | The type of the target object to which the association will be removed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteByObjectTypeByObjectIdByToObjectTypeByToObjectIdRequest struct via the builder pattern


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


## DeleteLabels

> DeleteLabels(ctx, associationTypeId, fromObjectType, toObjectType).Execute()

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
	associationTypeId := int32(56) // int32 | The unique identifier of the association type to be deleted.
	fromObjectType := "fromObjectType_example" // string | The type of the source object for the association to be deleted.
	toObjectType := "toObjectType_example" // string | The type of the target object for the association to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicAPI.DeleteLabels(context.Background(), associationTypeId, fromObjectType, toObjectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.DeleteLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**associationTypeId** | **int32** | The unique identifier of the association type to be deleted. | 
**fromObjectType** | **string** | The type of the source object for the association to be deleted. | 
**toObjectType** | **string** | The type of the target object for the association to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLabelsRequest struct via the builder pattern


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
	objectType := "objectType_example" // string | The type of the object to be archived, such as contacts, companies, or deals.

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
**objectType** | **string** | The type of the object to be archived, such as contacts, companies, or deals. | 

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
	objectId := "objectId_example" // string | The unique identifier of the object to read.
	objectType := "objectType_example" // string | The type of the object to read.
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)
	associations := []string{"Inner_example"} // []string | A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored. (optional)
	idProperty := "idProperty_example" // string | The name of a property whose values are unique for this object type (optional)
	properties := []string{"Inner_example"} // []string | A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored. (optional)
	propertiesWithHistory := []string{"Inner_example"} // []string | A comma separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored. (optional)

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
**objectId** | **string** | The unique identifier of the object to read. | 
**objectType** | **string** | The type of the object to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByObjectTypeByObjectIdRequest struct via the builder pattern


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


## GetByObjectTypeByObjectIdByToObjectType

> CollectionResponseMultiAssociatedObjectWithLabelForwardPaging GetByObjectTypeByObjectIdByToObjectType(ctx, objectId, objectType, toObjectType).After(after).Limit(limit).Execute()

Retrieve all associations by object type



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
	objectId := "objectId_example" // string | The unique identifier of the object for which associations are being listed.
	objectType := "objectType_example" // string | The type of the object for which associations are being listed.
	toObjectType := "toObjectType_example" // string | The type of the associated object to retrieve.
	after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetByObjectTypeByObjectIdByToObjectType(context.Background(), objectId, objectType, toObjectType).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetByObjectTypeByObjectIdByToObjectType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetByObjectTypeByObjectIdByToObjectType`: CollectionResponseMultiAssociatedObjectWithLabelForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetByObjectTypeByObjectIdByToObjectType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The unique identifier of the object for which associations are being listed. | 
**objectType** | **string** | The type of the object for which associations are being listed. | 
**toObjectType** | **string** | The type of the associated object to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByObjectTypeByObjectIdByToObjectTypeRequest struct via the builder pattern


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


## GetDefinitionsConfigurations

> CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging GetDefinitionsConfigurations(ctx, fromObjectType, toObjectType).Execute()

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
	fromObjectType := "fromObjectType_example" // string | The type of the source object in the association.
	toObjectType := "toObjectType_example" // string | The type of the target object in the association.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetDefinitionsConfigurations(context.Background(), fromObjectType, toObjectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetDefinitionsConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefinitionsConfigurations`: CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetDefinitionsConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** | The type of the source object in the association. | 
**toObjectType** | **string** | The type of the target object in the association. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefinitionsConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging**](CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefinitionsConfigurationsAll

> CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging GetDefinitionsConfigurationsAll(ctx).Execute()

Read All



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetDefinitionsConfigurationsAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetDefinitionsConfigurationsAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefinitionsConfigurationsAll`: CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetDefinitionsConfigurationsAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefinitionsConfigurationsAllRequest struct via the builder pattern


### Return type

[**CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging**](CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLabels

> CollectionResponseAssociationSpecWithLabelNoPaging GetLabels(ctx, fromObjectType, toObjectType).Execute()

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
	fromObjectType := "fromObjectType_example" // string | The type of the source object in the association.
	toObjectType := "toObjectType_example" // string | The type of the target object in the association.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetLabels(context.Background(), fromObjectType, toObjectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLabels`: CollectionResponseAssociationSpecWithLabelNoPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** | The type of the source object in the association. | 
**toObjectType** | **string** | The type of the target object in the association. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CollectionResponseAssociationSpecWithLabelNoPaging**](CollectionResponseAssociationSpecWithLabelNoPaging.md)

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
	objectType := "objectType_example" // string | The type of CRM object to list. This parameter specifies which object type's records will be retrieved.
	after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)
	associations := []string{"Inner_example"} // []string | A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored. (optional)
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional) (default to 10)
	properties := []string{"Inner_example"} // []string | A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored. (optional)
	propertiesWithHistory := []string{"Inner_example"} // []string | A comma separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored. Usage of this parameter will reduce the maximum number of objects that can be read by a single request. (optional)

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
**objectType** | **string** | The type of CRM object to list. This parameter specifies which object type&#39;s records will be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRootRequest struct via the builder pattern


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
	idProperty := "idProperty_example" // string | The name of a property whose values are unique for this object type (optional)

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


## PostLabels

> CollectionResponseAssociationSpecWithLabelNoPaging PostLabels(ctx, fromObjectType, toObjectType).PublicAssociationDefinitionCreateRequest(publicAssociationDefinitionCreateRequest).Execute()

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
	fromObjectType := "fromObjectType_example" // string | The type of the source object in the association.
	toObjectType := "toObjectType_example" // string | The type of the target object in the association.
	publicAssociationDefinitionCreateRequest := *openapiclient.NewPublicAssociationDefinitionCreateRequest("Label_example", "Name_example") // PublicAssociationDefinitionCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PostLabels(context.Background(), fromObjectType, toObjectType).PublicAssociationDefinitionCreateRequest(publicAssociationDefinitionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLabels`: CollectionResponseAssociationSpecWithLabelNoPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PostLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** | The type of the source object in the association. | 
**toObjectType** | **string** | The type of the target object in the association. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **publicAssociationDefinitionCreateRequest** | [**PublicAssociationDefinitionCreateRequest**](PublicAssociationDefinitionCreateRequest.md) |  | 

### Return type

[**CollectionResponseAssociationSpecWithLabelNoPaging**](CollectionResponseAssociationSpecWithLabelNoPaging.md)

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


## PostUsageHighUsageReport

> ReportCreationResponse PostUsageHighUsageReport(ctx, userId).Execute()

Report high usage



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
	userId := int32(56) // int32 | The unique identifier of the user requesting the report. This is an integer value.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PostUsageHighUsageReport(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostUsageHighUsageReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUsageHighUsageReport`: ReportCreationResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PostUsageHighUsageReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | The unique identifier of the user requesting the report. This is an integer value. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsageHighUsageReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReportCreationResponse**](ReportCreationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLabels

> PutLabels(ctx, fromObjectType, toObjectType).PublicAssociationDefinitionUpdateRequest(publicAssociationDefinitionUpdateRequest).Execute()

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
	fromObjectType := "fromObjectType_example" // string | The type of the source object in the association.
	toObjectType := "toObjectType_example" // string | The type of the target object in the association.
	publicAssociationDefinitionUpdateRequest := *openapiclient.NewPublicAssociationDefinitionUpdateRequest(int32(123), "Label_example") // PublicAssociationDefinitionUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicAPI.PutLabels(context.Background(), fromObjectType, toObjectType).PublicAssociationDefinitionUpdateRequest(publicAssociationDefinitionUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PutLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** | The type of the source object in the association. | 
**toObjectType** | **string** | The type of the target object in the association. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **publicAssociationDefinitionUpdateRequest** | [**PublicAssociationDefinitionUpdateRequest**](PublicAssociationDefinitionUpdateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRoot

> LabelsBetweenObjectPair PutRoot(ctx, objectId, objectType, toObjectId, toObjectType).AssociationSpec(associationSpec).Execute()

Associate records (labelled)



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
	objectId := "objectId_example" // string | The unique identifier of the source object.
	objectType := "objectType_example" // string | The type of the source object to associate.
	toObjectId := "toObjectId_example" // string | The unique identifier of the target object.
	toObjectType := "toObjectType_example" // string | The type of the target object to associate.
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
**objectId** | **string** | The unique identifier of the source object. | 
**objectType** | **string** | The type of the source object to associate. | 
**toObjectId** | **string** | The unique identifier of the target object. | 
**toObjectType** | **string** | The type of the target object to associate. | 

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

