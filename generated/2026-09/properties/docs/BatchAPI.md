# \BatchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostPropertiesBatchArchive**](BatchAPI.md#PostPropertiesBatchArchive) | **Post** /crm/properties/2026-09/{objectType}/batch/archive | Archive a batch of properties
[**PostPropertiesBatchCreate**](BatchAPI.md#PostPropertiesBatchCreate) | **Post** /crm/properties/2026-09/{objectType}/batch/create | Create a batch of properties
[**PostPropertiesBatchRead**](BatchAPI.md#PostPropertiesBatchRead) | **Post** /crm/properties/2026-09/{objectType}/batch/read | Read a batch of properties



## PostPropertiesBatchArchive

> PostPropertiesBatchArchive(ctx, objectType).BatchInputPropertyName(batchInputPropertyName).Execute()

Archive a batch of properties



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
	objectType := "objectType_example" // string | The type of object whose properties are to be archived. This is a required parameter and identifies the specific object type within HubSpot.
	batchInputPropertyName := *openapiclient.NewBatchInputPropertyName([]openapiclient.PropertyName{*openapiclient.NewPropertyName("Name_example")}) // BatchInputPropertyName | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchAPI.PostPropertiesBatchArchive(context.Background(), objectType).BatchInputPropertyName(batchInputPropertyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostPropertiesBatchArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of object whose properties are to be archived. This is a required parameter and identifies the specific object type within HubSpot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPropertiesBatchArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputPropertyName** | [**BatchInputPropertyName**](BatchInputPropertyName.md) |  | 

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


## PostPropertiesBatchCreate

> BatchResponseProperty PostPropertiesBatchCreate(ctx, objectType).BatchInputPropertyCreate(batchInputPropertyCreate).Execute()

Create a batch of properties



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
	objectType := "objectType_example" // string | The type of CRM object for which the properties are being created, such as contacts, companies, or deals.
	batchInputPropertyCreate := *openapiclient.NewBatchInputPropertyCreate([]openapiclient.PropertyCreate{*openapiclient.NewPropertyCreate("FieldType_example", "GroupName_example", "Label_example", "Name_example", "Type_example")}) // BatchInputPropertyCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostPropertiesBatchCreate(context.Background(), objectType).BatchInputPropertyCreate(batchInputPropertyCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostPropertiesBatchCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPropertiesBatchCreate`: BatchResponseProperty
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostPropertiesBatchCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of CRM object for which the properties are being created, such as contacts, companies, or deals. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPropertiesBatchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputPropertyCreate** | [**BatchInputPropertyCreate**](BatchInputPropertyCreate.md) |  | 

### Return type

[**BatchResponseProperty**](BatchResponseProperty.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPropertiesBatchRead

> BatchResponseProperty PostPropertiesBatchRead(ctx, objectType).BatchReadInputPropertyName(batchReadInputPropertyName).Locale(locale).Execute()

Read a batch of properties



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
	objectType := "objectType_example" // string | The type of CRM object for which properties are being retrieved. This is a required path parameter.
	batchReadInputPropertyName := *openapiclient.NewBatchReadInputPropertyName(false, "DataSensitivity_example", []openapiclient.PropertyName{*openapiclient.NewPropertyName("Name_example")}) // BatchReadInputPropertyName | 
	locale := "locale_example" // string | A string specifying the locale to use for the response. This can be used to tailor the response to a specific language or region. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostPropertiesBatchRead(context.Background(), objectType).BatchReadInputPropertyName(batchReadInputPropertyName).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostPropertiesBatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPropertiesBatchRead`: BatchResponseProperty
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostPropertiesBatchRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of CRM object for which properties are being retrieved. This is a required path parameter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPropertiesBatchReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchReadInputPropertyName** | [**BatchReadInputPropertyName**](BatchReadInputPropertyName.md) |  | 
 **locale** | **string** | A string specifying the locale to use for the response. This can be used to tailor the response to a specific language or region. | 

### Return type

[**BatchResponseProperty**](BatchResponseProperty.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

