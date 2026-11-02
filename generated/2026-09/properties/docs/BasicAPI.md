# \BasicAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProperties**](BasicAPI.md#DeleteProperties) | **Delete** /crm/properties/2026-09/{objectType}/{propertyName} | Archive a property
[**DeletePropertiesGroups**](BasicAPI.md#DeletePropertiesGroups) | **Delete** /crm/properties/2026-09/{objectType}/groups/{groupName} | Archive a property group
[**GetProperties**](BasicAPI.md#GetProperties) | **Get** /crm/properties/2026-09/{objectType} | Read all properties
[**GetPropertiesByObjectTypeByPropertyName**](BasicAPI.md#GetPropertiesByObjectTypeByPropertyName) | **Get** /crm/properties/2026-09/{objectType}/{propertyName} | Read a property
[**GetPropertiesByObjectTypeGroupsByGroupName**](BasicAPI.md#GetPropertiesByObjectTypeGroupsByGroupName) | **Get** /crm/properties/2026-09/{objectType}/groups/{groupName} | Read a property group
[**GetPropertiesGroups**](BasicAPI.md#GetPropertiesGroups) | **Get** /crm/properties/2026-09/{objectType}/groups | Read all property groups
[**PatchProperties**](BasicAPI.md#PatchProperties) | **Patch** /crm/properties/2026-09/{objectType}/{propertyName} | Update a property
[**PatchPropertiesGroups**](BasicAPI.md#PatchPropertiesGroups) | **Patch** /crm/properties/2026-09/{objectType}/groups/{groupName} | Update a property group
[**PostProperties**](BasicAPI.md#PostProperties) | **Post** /crm/properties/2026-09/{objectType} | Create a property
[**PostPropertiesGroups**](BasicAPI.md#PostPropertiesGroups) | **Post** /crm/properties/2026-09/{objectType}/groups | Create a property group



## DeleteProperties

> DeleteProperties(ctx, objectType, propertyName).Execute()

Archive a property



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
	objectType := "objectType_example" // string | The type of CRM object that the property belongs to, such as contacts, companies, or deals.
	propertyName := "propertyName_example" // string | The unique name of the property to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicAPI.DeleteProperties(context.Background(), objectType, propertyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.DeleteProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of CRM object that the property belongs to, such as contacts, companies, or deals. | 
**propertyName** | **string** | The unique name of the property to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePropertiesRequest struct via the builder pattern


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


## DeletePropertiesGroups

> DeletePropertiesGroups(ctx, groupName, objectType).Execute()

Archive a property group



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
	groupName := "groupName_example" // string | The unique name of the property group to be deleted.
	objectType := "objectType_example" // string | The type of object for which the property group is defined, such as contacts, companies, or deals.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicAPI.DeletePropertiesGroups(context.Background(), groupName, objectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.DeletePropertiesGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** | The unique name of the property group to be deleted. | 
**objectType** | **string** | The type of object for which the property group is defined, such as contacts, companies, or deals. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePropertiesGroupsRequest struct via the builder pattern


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


## GetProperties

> CollectionResponsePropertyNoPaging GetProperties(ctx, objectType).Archived(archived).DataSensitivity(dataSensitivity).Locale(locale).Properties(properties).Execute()

Read all properties



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
	objectType := "objectType_example" // string | The type of CRM object for which properties are being retrieved, such as 'contacts', 'companies', etc.
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)
	dataSensitivity := "dataSensitivity_example" // string | Filter properties based on data sensitivity level. Valid values are 'non_sensitive', 'sensitive', and 'highly_sensitive'. Defaults to 'non_sensitive'. (optional) (default to "non_sensitive")
	locale := "locale_example" // string | Specify the locale to retrieve properties in a particular language or regional format. (optional)
	properties := "properties_example" // string | A comma-separated list of property names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetProperties(context.Background(), objectType).Archived(archived).DataSensitivity(dataSensitivity).Locale(locale).Properties(properties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProperties`: CollectionResponsePropertyNoPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of CRM object for which properties are being retrieved, such as &#39;contacts&#39;, &#39;companies&#39;, etc. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]
 **dataSensitivity** | **string** | Filter properties based on data sensitivity level. Valid values are &#39;non_sensitive&#39;, &#39;sensitive&#39;, and &#39;highly_sensitive&#39;. Defaults to &#39;non_sensitive&#39;. | [default to &quot;non_sensitive&quot;]
 **locale** | **string** | Specify the locale to retrieve properties in a particular language or regional format. | 
 **properties** | **string** | A comma-separated list of property names to include in the response. | 

### Return type

[**CollectionResponsePropertyNoPaging**](CollectionResponsePropertyNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPropertiesByObjectTypeByPropertyName

> Property GetPropertiesByObjectTypeByPropertyName(ctx, objectType, propertyName).Archived(archived).DataSensitivity(dataSensitivity).Locale(locale).Properties(properties).Execute()

Read a property



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
	objectType := "objectType_example" // string | The type of object the property belongs to, such as 'contacts' or 'deals'.
	propertyName := "propertyName_example" // string | The name of the property to retrieve.
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)
	dataSensitivity := "dataSensitivity_example" // string | The level of data sensitivity to filter properties by. Valid values are 'non_sensitive', 'sensitive', and 'highly_sensitive'. Defaults to 'non_sensitive'. (optional) (default to "non_sensitive")
	locale := "locale_example" // string | The locale to use for property labels and descriptions. (optional)
	properties := "properties_example" // string | A comma-separated list of property names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetPropertiesByObjectTypeByPropertyName(context.Background(), objectType, propertyName).Archived(archived).DataSensitivity(dataSensitivity).Locale(locale).Properties(properties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetPropertiesByObjectTypeByPropertyName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPropertiesByObjectTypeByPropertyName`: Property
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetPropertiesByObjectTypeByPropertyName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of object the property belongs to, such as &#39;contacts&#39; or &#39;deals&#39;. | 
**propertyName** | **string** | The name of the property to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPropertiesByObjectTypeByPropertyNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]
 **dataSensitivity** | **string** | The level of data sensitivity to filter properties by. Valid values are &#39;non_sensitive&#39;, &#39;sensitive&#39;, and &#39;highly_sensitive&#39;. Defaults to &#39;non_sensitive&#39;. | [default to &quot;non_sensitive&quot;]
 **locale** | **string** | The locale to use for property labels and descriptions. | 
 **properties** | **string** | A comma-separated list of property names to include in the response. | 

### Return type

[**Property**](Property.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPropertiesByObjectTypeGroupsByGroupName

> PropertyGroup GetPropertiesByObjectTypeGroupsByGroupName(ctx, groupName, objectType).Locale(locale).Execute()

Read a property group



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
	groupName := "groupName_example" // string | The name of the property group to retrieve details for.
	objectType := "objectType_example" // string | The type of CRM object for which the property group is being retrieved, such as contacts or deals.
	locale := "locale_example" // string | The language locale to use for the response. This can affect the language of the labels and descriptions returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetPropertiesByObjectTypeGroupsByGroupName(context.Background(), groupName, objectType).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetPropertiesByObjectTypeGroupsByGroupName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPropertiesByObjectTypeGroupsByGroupName`: PropertyGroup
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetPropertiesByObjectTypeGroupsByGroupName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** | The name of the property group to retrieve details for. | 
**objectType** | **string** | The type of CRM object for which the property group is being retrieved, such as contacts or deals. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPropertiesByObjectTypeGroupsByGroupNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **locale** | **string** | The language locale to use for the response. This can affect the language of the labels and descriptions returned. | 

### Return type

[**PropertyGroup**](PropertyGroup.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPropertiesGroups

> CollectionResponsePropertyGroupNoPaging GetPropertiesGroups(ctx, objectType).Locale(locale).Execute()

Read all property groups



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
	objectType := "objectType_example" // string | The type of CRM object for which to retrieve property groups, such as 'contacts' or 'deals'.
	locale := "locale_example" // string | The language locale to use for property group labels. This can help in displaying labels in the preferred language. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetPropertiesGroups(context.Background(), objectType).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetPropertiesGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPropertiesGroups`: CollectionResponsePropertyGroupNoPaging
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetPropertiesGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of CRM object for which to retrieve property groups, such as &#39;contacts&#39; or &#39;deals&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPropertiesGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locale** | **string** | The language locale to use for property group labels. This can help in displaying labels in the preferred language. | 

### Return type

[**CollectionResponsePropertyGroupNoPaging**](CollectionResponsePropertyGroupNoPaging.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchProperties

> Property PatchProperties(ctx, objectType, propertyName).PropertyUpdate(propertyUpdate).Execute()

Update a property



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
	objectType := "objectType_example" // string | The type of object for which the property is being updated, such as contacts, companies, or deals.
	propertyName := "propertyName_example" // string | The unique name of the property to update.
	propertyUpdate := *openapiclient.NewPropertyUpdate() // PropertyUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PatchProperties(context.Background(), objectType, propertyName).PropertyUpdate(propertyUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PatchProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProperties`: Property
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PatchProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of object for which the property is being updated, such as contacts, companies, or deals. | 
**propertyName** | **string** | The unique name of the property to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **propertyUpdate** | [**PropertyUpdate**](PropertyUpdate.md) |  | 

### Return type

[**Property**](Property.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPropertiesGroups

> PropertyGroup PatchPropertiesGroups(ctx, groupName, objectType).PropertyGroupUpdate(propertyGroupUpdate).Execute()

Update a property group



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
	groupName := "groupName_example" // string | The name of the property group to update.
	objectType := "objectType_example" // string | The type of object for which the property group is being updated, such as contacts, companies, or deals.
	propertyGroupUpdate := *openapiclient.NewPropertyGroupUpdate() // PropertyGroupUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PatchPropertiesGroups(context.Background(), groupName, objectType).PropertyGroupUpdate(propertyGroupUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PatchPropertiesGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchPropertiesGroups`: PropertyGroup
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PatchPropertiesGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** | The name of the property group to update. | 
**objectType** | **string** | The type of object for which the property group is being updated, such as contacts, companies, or deals. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPropertiesGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **propertyGroupUpdate** | [**PropertyGroupUpdate**](PropertyGroupUpdate.md) |  | 

### Return type

[**PropertyGroup**](PropertyGroup.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProperties

> Property PostProperties(ctx, objectType).PropertyCreate(propertyCreate).Execute()

Create a property



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
	objectType := "objectType_example" // string | The type of object for which the property is being created, such as contacts, companies, or deals.
	propertyCreate := *openapiclient.NewPropertyCreate("FieldType_example", "GroupName_example", "Label_example", "Name_example", "Type_example") // PropertyCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PostProperties(context.Background(), objectType).PropertyCreate(propertyCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProperties`: Property
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PostProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of object for which the property is being created, such as contacts, companies, or deals. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **propertyCreate** | [**PropertyCreate**](PropertyCreate.md) |  | 

### Return type

[**Property**](Property.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPropertiesGroups

> PropertyGroup PostPropertiesGroups(ctx, objectType).PropertyGroupCreate(propertyGroupCreate).Execute()

Create a property group



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
	objectType := "objectType_example" // string | The type of object for which the property group is being created, such as contacts, companies, or deals.
	propertyGroupCreate := *openapiclient.NewPropertyGroupCreate("Label_example", "Name_example") // PropertyGroupCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PostPropertiesGroups(context.Background(), objectType).PropertyGroupCreate(propertyGroupCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostPropertiesGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPropertiesGroups`: PropertyGroup
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PostPropertiesGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of object for which the property group is being created, such as contacts, companies, or deals. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPropertiesGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **propertyGroupCreate** | [**PropertyGroupCreate**](PropertyGroupCreate.md) |  | 

### Return type

[**PropertyGroup**](PropertyGroup.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

