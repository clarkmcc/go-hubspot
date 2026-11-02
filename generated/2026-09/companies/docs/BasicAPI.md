# \BasicAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanies**](BasicAPI.md#DeleteCompanies) | **Delete** /crm/objects/2026-09/companies/{companyId} | Archive a company
[**DeleteRoot**](BasicAPI.md#DeleteRoot) | **Delete** /crm/objects/2026-09/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId} | Delete
[**GetCompanies**](BasicAPI.md#GetCompanies) | **Get** /crm/objects/2026-09/companies | Retrieve companies
[**GetCompaniesByCompanyId**](BasicAPI.md#GetCompaniesByCompanyId) | **Get** /crm/objects/2026-09/companies/{companyId} | Retrieve a company
[**GetCompaniesByCompanyIdByToObjectType**](BasicAPI.md#GetCompaniesByCompanyIdByToObjectType) | **Get** /crm/objects/2026-09/companies/{companyId}/associations/{toObjectType} | List associations
[**PatchCompanies**](BasicAPI.md#PatchCompanies) | **Patch** /crm/objects/2026-09/companies/{companyId} | Update a company
[**PostCompanies**](BasicAPI.md#PostCompanies) | **Post** /crm/objects/2026-09/companies | Create a company
[**PutRoot**](BasicAPI.md#PutRoot) | **Put** /crm/objects/2026-09/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId} | Create



## DeleteCompanies

> DeleteCompanies(ctx, companyId).Execute()

Archive a company



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
	companyId := "companyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicAPI.DeleteCompanies(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.DeleteCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompaniesRequest struct via the builder pattern


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
	objectId := "objectId_example" // string | The unique identifier of the first object involved in the association.
	objectType := "objectType_example" // string | The type of the first object involved in the association.
	toObjectId := "toObjectId_example" // string | The unique identifier of the second object involved in the association.
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
**objectId** | **string** | The unique identifier of the first object involved in the association. | 
**objectType** | **string** | The type of the first object involved in the association. | 
**toObjectId** | **string** | The unique identifier of the second object involved in the association. | 
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


## GetCompanies

> CollectionResponseSimplePublicObjectWithAssociationsForwardPaging GetCompanies(ctx).After(after).Archived(archived).Associations(associations).Limit(limit).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()

Retrieve companies



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
	propertiesWithHistory := []string{"Inner_example"} // []string | A comma separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored. Usage of this parameter will reduce the maximum number of companies that can be read by a single request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetCompanies(context.Background()).After(after).Archived(archived).Associations(associations).Limit(limit).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanies`: CollectionResponseSimplePublicObjectWithAssociationsForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetCompanies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]
 **associations** | **[]string** | A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored. | 
 **limit** | **int32** | The maximum number of results to display per page. | [default to 10]
 **properties** | **[]string** | A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored. | 
 **propertiesWithHistory** | **[]string** | A comma separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored. Usage of this parameter will reduce the maximum number of companies that can be read by a single request. | 

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


## GetCompaniesByCompanyId

> SimplePublicObjectWithAssociations GetCompaniesByCompanyId(ctx, companyId).Archived(archived).Associations(associations).IdProperty(idProperty).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()

Retrieve a company



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
	companyId := "companyId_example" // string | 
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)
	associations := []string{"Inner_example"} // []string | A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored. (optional)
	idProperty := "idProperty_example" // string | The name of a property whose values are unique for this object type (optional)
	properties := []string{"Inner_example"} // []string | A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored. (optional)
	propertiesWithHistory := []string{"Inner_example"} // []string | A comma separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetCompaniesByCompanyId(context.Background(), companyId).Archived(archived).Associations(associations).IdProperty(idProperty).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetCompaniesByCompanyId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompaniesByCompanyId`: SimplePublicObjectWithAssociations
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetCompaniesByCompanyId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesByCompanyIdRequest struct via the builder pattern


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


## GetCompaniesByCompanyIdByToObjectType

> CollectionResponseMultiAssociatedObjectWithLabelForwardPaging GetCompaniesByCompanyIdByToObjectType(ctx, companyId, toObjectType).After(after).Limit(limit).Execute()

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
	companyId := "companyId_example" // string | 
	toObjectType := "toObjectType_example" // string | The type of the object to which associations are being retrieved.
	after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetCompaniesByCompanyIdByToObjectType(context.Background(), companyId, toObjectType).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetCompaniesByCompanyIdByToObjectType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompaniesByCompanyIdByToObjectType`: CollectionResponseMultiAssociatedObjectWithLabelForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetCompaniesByCompanyIdByToObjectType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**toObjectType** | **string** | The type of the object to which associations are being retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesByCompanyIdByToObjectTypeRequest struct via the builder pattern


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


## PatchCompanies

> SimplePublicObject PatchCompanies(ctx, companyId).SimplePublicObjectInput(simplePublicObjectInput).IdProperty(idProperty).Execute()

Update a company



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
	companyId := "companyId_example" // string | 
	simplePublicObjectInput := *openapiclient.NewSimplePublicObjectInput(map[string]string{"key": "Inner_example"}) // SimplePublicObjectInput | 
	idProperty := "idProperty_example" // string | The name of a property whose values are unique for this object type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PatchCompanies(context.Background(), companyId).SimplePublicObjectInput(simplePublicObjectInput).IdProperty(idProperty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PatchCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanies`: SimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PatchCompanies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompaniesRequest struct via the builder pattern


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


## PostCompanies

> SimplePublicObject PostCompanies(ctx).SimplePublicObjectInputForCreate(simplePublicObjectInputForCreate).Execute()

Create a company



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
	resp, r, err := apiClient.BasicAPI.PostCompanies(context.Background()).SimplePublicObjectInputForCreate(simplePublicObjectInputForCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanies`: SimplePublicObject
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PostCompanies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompaniesRequest struct via the builder pattern


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

