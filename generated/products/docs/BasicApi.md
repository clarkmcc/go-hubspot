# \BasicApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3ObjectsProductsProductIdArchive**](BasicApi.md#DeleteCrmV3ObjectsProductsProductIdArchive) | **Delete** /crm/v3/objects/products/{productId} | Archive
[**GetCrmV3ObjectsProductsGetPage**](BasicApi.md#GetCrmV3ObjectsProductsGetPage) | **Get** /crm/v3/objects/products | List
[**GetCrmV3ObjectsProductsProductIdGetById**](BasicApi.md#GetCrmV3ObjectsProductsProductIdGetById) | **Get** /crm/v3/objects/products/{productId} | Read
[**PatchCrmV3ObjectsProductsProductIdUpdate**](BasicApi.md#PatchCrmV3ObjectsProductsProductIdUpdate) | **Patch** /crm/v3/objects/products/{productId} | Update
[**PostCrmV3ObjectsProductsCreate**](BasicApi.md#PostCrmV3ObjectsProductsCreate) | **Post** /crm/v3/objects/products | Create



## DeleteCrmV3ObjectsProductsProductIdArchive

> DeleteCrmV3ObjectsProductsProductIdArchive(ctx, productId).Execute()

Archive



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    productId := "productId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BasicApi.DeleteCrmV3ObjectsProductsProductIdArchive(context.Background(), productId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BasicApi.DeleteCrmV3ObjectsProductsProductIdArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3ObjectsProductsProductIdArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3ObjectsProductsGetPage

> CollectionResponseSimplePublicObjectWithAssociationsForwardPaging GetCrmV3ObjectsProductsGetPage(ctx).Limit(limit).After(after).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Associations(associations).Archived(archived).Execute()

List



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := int32(56) // int32 | The maximum number of results to display per page. (optional) (default to 10)
    after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
    properties := []string{"Inner_example"} // []string | A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored. (optional)
    propertiesWithHistory := []string{"Inner_example"} // []string | A comma separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored. Usage of this parameter will reduce the maximum number of objects that can be read by a single request. (optional)
    associations := []string{"Inner_example"} // []string | A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored. (optional)
    archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BasicApi.GetCrmV3ObjectsProductsGetPage(context.Background()).Limit(limit).After(after).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Associations(associations).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BasicApi.GetCrmV3ObjectsProductsGetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3ObjectsProductsGetPage`: CollectionResponseSimplePublicObjectWithAssociationsForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `BasicApi.GetCrmV3ObjectsProductsGetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ObjectsProductsGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of results to display per page. | [default to 10]
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **properties** | **[]string** | A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored. | 
 **propertiesWithHistory** | **[]string** | A comma separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored. Usage of this parameter will reduce the maximum number of objects that can be read by a single request. | 
 **associations** | **[]string** | A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored. | 
 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

[**CollectionResponseSimplePublicObjectWithAssociationsForwardPaging**](CollectionResponseSimplePublicObjectWithAssociationsForwardPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3ObjectsProductsProductIdGetById

> SimplePublicObjectWithAssociations GetCrmV3ObjectsProductsProductIdGetById(ctx, productId).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Associations(associations).Archived(archived).IdProperty(idProperty).Execute()

Read



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    productId := "productId_example" // string | 
    properties := []string{"Inner_example"} // []string | A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored. (optional)
    propertiesWithHistory := []string{"Inner_example"} // []string | A comma separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored. (optional)
    associations := []string{"Inner_example"} // []string | A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored. (optional)
    archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)
    idProperty := "idProperty_example" // string | The name of a property whose values are unique for this object type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BasicApi.GetCrmV3ObjectsProductsProductIdGetById(context.Background(), productId).Properties(properties).PropertiesWithHistory(propertiesWithHistory).Associations(associations).Archived(archived).IdProperty(idProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BasicApi.GetCrmV3ObjectsProductsProductIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3ObjectsProductsProductIdGetById`: SimplePublicObjectWithAssociations
    fmt.Fprintf(os.Stdout, "Response from `BasicApi.GetCrmV3ObjectsProductsProductIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ObjectsProductsProductIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **properties** | **[]string** | A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored. | 
 **propertiesWithHistory** | **[]string** | A comma separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored. | 
 **associations** | **[]string** | A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored. | 
 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]
 **idProperty** | **string** | The name of a property whose values are unique for this object type | 

### Return type

[**SimplePublicObjectWithAssociations**](SimplePublicObjectWithAssociations.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCrmV3ObjectsProductsProductIdUpdate

> SimplePublicObject PatchCrmV3ObjectsProductsProductIdUpdate(ctx, productId).SimplePublicObjectInput(simplePublicObjectInput).IdProperty(idProperty).Execute()

Update



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    productId := "productId_example" // string | 
    simplePublicObjectInput := *openapiclient.NewSimplePublicObjectInput(map[string]string{"key": "Inner_example"}) // SimplePublicObjectInput | 
    idProperty := "idProperty_example" // string | The name of a property whose values are unique for this object type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BasicApi.PatchCrmV3ObjectsProductsProductIdUpdate(context.Background(), productId).SimplePublicObjectInput(simplePublicObjectInput).IdProperty(idProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BasicApi.PatchCrmV3ObjectsProductsProductIdUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCrmV3ObjectsProductsProductIdUpdate`: SimplePublicObject
    fmt.Fprintf(os.Stdout, "Response from `BasicApi.PatchCrmV3ObjectsProductsProductIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCrmV3ObjectsProductsProductIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **simplePublicObjectInput** | [**SimplePublicObjectInput**](SimplePublicObjectInput.md) |  | 
 **idProperty** | **string** | The name of a property whose values are unique for this object type | 

### Return type

[**SimplePublicObject**](SimplePublicObject.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ObjectsProductsCreate

> SimplePublicObject PostCrmV3ObjectsProductsCreate(ctx).SimplePublicObjectInput(simplePublicObjectInput).Execute()

Create



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    simplePublicObjectInput := *openapiclient.NewSimplePublicObjectInput(map[string]string{"key": "Inner_example"}) // SimplePublicObjectInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BasicApi.PostCrmV3ObjectsProductsCreate(context.Background()).SimplePublicObjectInput(simplePublicObjectInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BasicApi.PostCrmV3ObjectsProductsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ObjectsProductsCreate`: SimplePublicObject
    fmt.Fprintf(os.Stdout, "Response from `BasicApi.PostCrmV3ObjectsProductsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsProductsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **simplePublicObjectInput** | [**SimplePublicObjectInput**](SimplePublicObjectInput.md) |  | 

### Return type

[**SimplePublicObject**](SimplePublicObject.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

