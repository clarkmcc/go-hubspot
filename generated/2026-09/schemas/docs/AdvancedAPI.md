# \AdvancedAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostSchemasByObjectType**](AdvancedAPI.md#PostSchemasByObjectType) | **Post** /crm-object-schemas/2026-09/schemas/{objectType}/associations | Create an association for a specified object type.



## PostSchemasByObjectType

> AssociationDefinition PostSchemasByObjectType(ctx, objectType).AssociationDefinitionEgg(associationDefinitionEgg).Execute()

Create an association for a specified object type.



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
	objectType := "objectType_example" // string | The unique identifier of the object type for which the association is being created.
	associationDefinitionEgg := *openapiclient.NewAssociationDefinitionEgg("FromObjectTypeId_example", "ToObjectTypeId_example") // AssociationDefinitionEgg | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedAPI.PostSchemasByObjectType(context.Background(), objectType).AssociationDefinitionEgg(associationDefinitionEgg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedAPI.PostSchemasByObjectType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSchemasByObjectType`: AssociationDefinition
	fmt.Fprintf(os.Stdout, "Response from `AdvancedAPI.PostSchemasByObjectType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The unique identifier of the object type for which the association is being created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSchemasByObjectTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **associationDefinitionEgg** | [**AssociationDefinitionEgg**](AssociationDefinitionEgg.md) |  | 

### Return type

[**AssociationDefinition**](AssociationDefinition.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

