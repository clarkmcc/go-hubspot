# \BasicAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmAssociations202609FromObjectTypeToObjectTypeLabelsAssociationTypeIdCrmAssociations202609BetaFromObjectTypeToObjectTypeLabelsAssociationTypeId**](BasicAPI.md#DeleteCrmAssociations202609FromObjectTypeToObjectTypeLabelsAssociationTypeIdCrmAssociations202609BetaFromObjectTypeToObjectTypeLabelsAssociationTypeId) | **Delete** /crm/associations/2026-09/{fromObjectType}/{toObjectType}/labels/{associationTypeId} | Delete
[**DeleteCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeToObjectId**](BasicAPI.md#DeleteCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeToObjectId) | **Delete** /crm/objects/2026-09/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId} | Delete associations between two records
[**DeleteCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectId**](BasicAPI.md#DeleteCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectId) | **Delete** /crm/objects/2026-09/{objectType}/{objectId} | Archive
[**GetCrmAssociations202609DefinitionsConfigurationsAllCrmAssociations202609BetaDefinitionsConfigurationsAll**](BasicAPI.md#GetCrmAssociations202609DefinitionsConfigurationsAllCrmAssociations202609BetaDefinitionsConfigurationsAll) | **Get** /crm/associations/2026-09/definitions/configurations/all | Read All
[**GetCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectType**](BasicAPI.md#GetCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectType) | **Get** /crm/associations/2026-09/definitions/configurations/{fromObjectType}/{toObjectType} | Read
[**GetCrmObjects202609ObjectTypeCrmObjects202609BetaObjectType**](BasicAPI.md#GetCrmObjects202609ObjectTypeCrmObjects202609BetaObjectType) | **Get** /crm/objects/2026-09/{objectType} | List
[**GetCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectType**](BasicAPI.md#GetCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectType) | **Get** /crm/objects/2026-09/{objectType}/{objectId}/associations/{toObjectType} | Retrieve all associations by object type
[**GetCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectId**](BasicAPI.md#GetCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectId) | **Get** /crm/objects/2026-09/{objectType}/{objectId} | Read
[**GetLabels**](BasicAPI.md#GetLabels) | **Get** /crm/associations/2026-09/{fromObjectType}/{toObjectType}/labels | Read
[**PatchCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectId**](BasicAPI.md#PatchCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectId) | **Patch** /crm/objects/2026-09/{objectType}/{objectId} | Update
[**PostCrmAssociations202609FromObjectTypeToObjectTypeLabelsCrmAssociations202609BetaFromObjectTypeToObjectTypeLabels**](BasicAPI.md#PostCrmAssociations202609FromObjectTypeToObjectTypeLabelsCrmAssociations202609BetaFromObjectTypeToObjectTypeLabels) | **Post** /crm/associations/2026-09/{fromObjectType}/{toObjectType}/labels | Create
[**PostCrmAssociations202609UsageHighUsageReportUserIdCrmAssociations202609BetaUsageHighUsageReportUserId**](BasicAPI.md#PostCrmAssociations202609UsageHighUsageReportUserIdCrmAssociations202609BetaUsageHighUsageReportUserId) | **Post** /crm/associations/2026-09/usage/high-usage-report/{userId} | Report high usage
[**PostCrmObjects202609ObjectTypeCrmObjects202609BetaObjectType**](BasicAPI.md#PostCrmObjects202609ObjectTypeCrmObjects202609BetaObjectType) | **Post** /crm/objects/2026-09/{objectType} | Create
[**PutCrmAssociations202609FromObjectTypeToObjectTypeLabelsCrmAssociations202609BetaFromObjectTypeToObjectTypeLabels**](BasicAPI.md#PutCrmAssociations202609FromObjectTypeToObjectTypeLabelsCrmAssociations202609BetaFromObjectTypeToObjectTypeLabels) | **Put** /crm/associations/2026-09/{fromObjectType}/{toObjectType}/labels | Update
[**PutCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeToObjectId**](BasicAPI.md#PutCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeToObjectId) | **Put** /crm/objects/2026-09/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId} | Associate records (labelled)



## DeleteCrmAssociations202609FromObjectTypeToObjectTypeLabelsAssociationTypeIdCrmAssociations202609BetaFromObjectTypeToObjectTypeLabelsAssociationTypeId

> DeleteCrmAssociations202609FromObjectTypeToObjectTypeLabelsAssociationTypeIdCrmAssociations202609BetaFromObjectTypeToObjectTypeLabelsAssociationTypeId(ctx, associationTypeId, fromObjectType, toObjectType).Execute()

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
	r, err := apiClient.BasicAPI.DeleteCrmAssociations202609FromObjectTypeToObjectTypeLabelsAssociationTypeIdCrmAssociations202609BetaFromObjectTypeToObjectTypeLabelsAssociationTypeId(context.Background(), associationTypeId, fromObjectType, toObjectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.DeleteCrmAssociations202609FromObjectTypeToObjectTypeLabelsAssociationTypeIdCrmAssociations202609BetaFromObjectTypeToObjectTypeLabelsAssociationTypeId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCrmAssociations202609FromObjectTypeToObjectTypeLabelsAssociationTypeIdCrmAssociations202609BetaFromObjectTypeToObjectTypeLabelsAssociationTypeIdRequest struct via the builder pattern


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
**objectId** | **string** | The unique identifier of the initial object from which the association will be removed. | 
**objectType** | **string** | The type of the initial object from which the association will be removed. | 
**toObjectId** | **string** | The unique identifier of the target object to which the association will be removed. | 
**toObjectType** | **string** | The type of the target object to which the association will be removed. | 

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


## DeleteCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectId

> DeleteCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectId(ctx, objectId, objectType).Execute()

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
	r, err := apiClient.BasicAPI.DeleteCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectId(context.Background(), objectId, objectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.DeleteCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectIdRequest struct via the builder pattern


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


## GetCrmAssociations202609DefinitionsConfigurationsAllCrmAssociations202609BetaDefinitionsConfigurationsAll

> CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging GetCrmAssociations202609DefinitionsConfigurationsAllCrmAssociations202609BetaDefinitionsConfigurationsAll(ctx).Execute()

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
	resp, r, err := apiClient.BasicAPI.GetCrmAssociations202609DefinitionsConfigurationsAllCrmAssociations202609BetaDefinitionsConfigurationsAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetCrmAssociations202609DefinitionsConfigurationsAllCrmAssociations202609BetaDefinitionsConfigurationsAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmAssociations202609DefinitionsConfigurationsAllCrmAssociations202609BetaDefinitionsConfigurationsAll`: CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetCrmAssociations202609DefinitionsConfigurationsAllCrmAssociations202609BetaDefinitionsConfigurationsAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmAssociations202609DefinitionsConfigurationsAllCrmAssociations202609BetaDefinitionsConfigurationsAllRequest struct via the builder pattern


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


## GetCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectType

> CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging GetCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectType(ctx, fromObjectType, toObjectType).Execute()

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
	resp, r, err := apiClient.BasicAPI.GetCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectType(context.Background(), fromObjectType, toObjectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectType`: CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** | The type of the source object in the association. | 
**toObjectType** | **string** | The type of the target object in the association. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmAssociations202609DefinitionsConfigurationsFromObjectTypeToObjectTypeCrmAssociations202609BetaDefinitionsConfigurationsFromObjectTypeToObjectTypeRequest struct via the builder pattern


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


## GetCrmObjects202609ObjectTypeCrmObjects202609BetaObjectType

> CollectionResponseSimplePublicObjectWithAssociationsForwardPaging GetCrmObjects202609ObjectTypeCrmObjects202609BetaObjectType(ctx, objectType).After(after).Archived(archived).Associations(associations).Limit(limit).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()

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
	resp, r, err := apiClient.BasicAPI.GetCrmObjects202609ObjectTypeCrmObjects202609BetaObjectType(context.Background(), objectType).After(after).Archived(archived).Associations(associations).Limit(limit).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetCrmObjects202609ObjectTypeCrmObjects202609BetaObjectType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmObjects202609ObjectTypeCrmObjects202609BetaObjectType`: CollectionResponseSimplePublicObjectWithAssociationsForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetCrmObjects202609ObjectTypeCrmObjects202609BetaObjectType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of CRM object to list. This parameter specifies which object type&#39;s records will be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmObjects202609ObjectTypeCrmObjects202609BetaObjectTypeRequest struct via the builder pattern


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


## GetCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectType

> CollectionResponseMultiAssociatedObjectWithLabelForwardPaging GetCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectType(ctx, objectId, objectType, toObjectType).After(after).Limit(limit).Execute()

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
	resp, r, err := apiClient.BasicAPI.GetCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectType(context.Background(), objectId, objectType, toObjectType).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectType`: CollectionResponseMultiAssociatedObjectWithLabelForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectType`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeRequest struct via the builder pattern


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


## GetCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectId

> SimplePublicObjectWithAssociations GetCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectId(ctx, objectId, objectType).Archived(archived).Associations(associations).IdProperty(idProperty).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()

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
	resp, r, err := apiClient.BasicAPI.GetCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectId(context.Background(), objectId, objectType).Archived(archived).Associations(associations).IdProperty(idProperty).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectId`: SimplePublicObjectWithAssociations
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The unique identifier of the object to read. | 
**objectType** | **string** | The type of the object to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectIdRequest struct via the builder pattern


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


## PatchCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectId

> SimplePublicObject PatchCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectId(ctx, objectId, objectType).SimplePublicObjectInput(simplePublicObjectInput).IdProperty(idProperty).Execute()

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
	resp, r, err := apiClient.BasicAPI.PatchCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectId(context.Background(), objectId, objectType).SimplePublicObjectInput(simplePublicObjectInput).IdProperty(idProperty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PatchCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectId`: SimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PatchCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The unique identifier of the object to update. | 
**objectType** | **string** | The type of the object to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCrmObjects202609ObjectTypeObjectIdCrmObjects202609BetaObjectTypeObjectIdRequest struct via the builder pattern


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


## PostCrmAssociations202609FromObjectTypeToObjectTypeLabelsCrmAssociations202609BetaFromObjectTypeToObjectTypeLabels

> CollectionResponseAssociationSpecWithLabelNoPaging PostCrmAssociations202609FromObjectTypeToObjectTypeLabelsCrmAssociations202609BetaFromObjectTypeToObjectTypeLabels(ctx, fromObjectType, toObjectType).PublicAssociationDefinitionCreateRequest(publicAssociationDefinitionCreateRequest).Execute()

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
	resp, r, err := apiClient.BasicAPI.PostCrmAssociations202609FromObjectTypeToObjectTypeLabelsCrmAssociations202609BetaFromObjectTypeToObjectTypeLabels(context.Background(), fromObjectType, toObjectType).PublicAssociationDefinitionCreateRequest(publicAssociationDefinitionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostCrmAssociations202609FromObjectTypeToObjectTypeLabelsCrmAssociations202609BetaFromObjectTypeToObjectTypeLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmAssociations202609FromObjectTypeToObjectTypeLabelsCrmAssociations202609BetaFromObjectTypeToObjectTypeLabels`: CollectionResponseAssociationSpecWithLabelNoPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PostCrmAssociations202609FromObjectTypeToObjectTypeLabelsCrmAssociations202609BetaFromObjectTypeToObjectTypeLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** | The type of the source object in the association. | 
**toObjectType** | **string** | The type of the target object in the association. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmAssociations202609FromObjectTypeToObjectTypeLabelsCrmAssociations202609BetaFromObjectTypeToObjectTypeLabelsRequest struct via the builder pattern


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


## PostCrmAssociations202609UsageHighUsageReportUserIdCrmAssociations202609BetaUsageHighUsageReportUserId

> ReportCreationResponse PostCrmAssociations202609UsageHighUsageReportUserIdCrmAssociations202609BetaUsageHighUsageReportUserId(ctx, userId).Execute()

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
	resp, r, err := apiClient.BasicAPI.PostCrmAssociations202609UsageHighUsageReportUserIdCrmAssociations202609BetaUsageHighUsageReportUserId(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostCrmAssociations202609UsageHighUsageReportUserIdCrmAssociations202609BetaUsageHighUsageReportUserId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmAssociations202609UsageHighUsageReportUserIdCrmAssociations202609BetaUsageHighUsageReportUserId`: ReportCreationResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PostCrmAssociations202609UsageHighUsageReportUserIdCrmAssociations202609BetaUsageHighUsageReportUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | The unique identifier of the user requesting the report. This is an integer value. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmAssociations202609UsageHighUsageReportUserIdCrmAssociations202609BetaUsageHighUsageReportUserIdRequest struct via the builder pattern


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


## PostCrmObjects202609ObjectTypeCrmObjects202609BetaObjectType

> SimplePublicObject PostCrmObjects202609ObjectTypeCrmObjects202609BetaObjectType(ctx, objectType).SimplePublicObjectInputForCreate(simplePublicObjectInputForCreate).Execute()

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
	resp, r, err := apiClient.BasicAPI.PostCrmObjects202609ObjectTypeCrmObjects202609BetaObjectType(context.Background(), objectType).SimplePublicObjectInputForCreate(simplePublicObjectInputForCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostCrmObjects202609ObjectTypeCrmObjects202609BetaObjectType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmObjects202609ObjectTypeCrmObjects202609BetaObjectType`: SimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PostCrmObjects202609ObjectTypeCrmObjects202609BetaObjectType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of CRM object to create, such as contacts, companies, or deals. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmObjects202609ObjectTypeCrmObjects202609BetaObjectTypeRequest struct via the builder pattern


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


## PutCrmAssociations202609FromObjectTypeToObjectTypeLabelsCrmAssociations202609BetaFromObjectTypeToObjectTypeLabels

> PutCrmAssociations202609FromObjectTypeToObjectTypeLabelsCrmAssociations202609BetaFromObjectTypeToObjectTypeLabels(ctx, fromObjectType, toObjectType).PublicAssociationDefinitionUpdateRequest(publicAssociationDefinitionUpdateRequest).Execute()

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
	r, err := apiClient.BasicAPI.PutCrmAssociations202609FromObjectTypeToObjectTypeLabelsCrmAssociations202609BetaFromObjectTypeToObjectTypeLabels(context.Background(), fromObjectType, toObjectType).PublicAssociationDefinitionUpdateRequest(publicAssociationDefinitionUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PutCrmAssociations202609FromObjectTypeToObjectTypeLabelsCrmAssociations202609BetaFromObjectTypeToObjectTypeLabels``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutCrmAssociations202609FromObjectTypeToObjectTypeLabelsCrmAssociations202609BetaFromObjectTypeToObjectTypeLabelsRequest struct via the builder pattern


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


## PutCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeToObjectId

> LabelsBetweenObjectPair PutCrmObjects202609ObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCrmObjects202609BetaObjectTypeObjectIdAssociationsToObjectTypeToObjectId(ctx, objectId, objectType, toObjectId, toObjectType).AssociationSpec(associationSpec).Execute()

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
**objectId** | **string** | The unique identifier of the source object. | 
**objectType** | **string** | The type of the source object to associate. | 
**toObjectId** | **string** | The unique identifier of the target object. | 
**toObjectType** | **string** | The type of the target object to associate. | 

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

