# Go API client for deals

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v3
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import deals "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), deals.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), deals.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), deals.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), deals.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.hubapi.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AssociationsApi* | [**DeleteCrmV3ObjectsDealsDealIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive**](docs/AssociationsApi.md#deletecrmv3objectsdealsdealidassociationstoobjecttypetoobjectidassociationtypearchive) | **Delete** /crm/v3/objects/deals/{dealId}/associations/{toObjectType}/{toObjectId}/{associationType} | Remove an association between two deals
*AssociationsApi* | [**GetCrmV3ObjectsDealsDealIdAssociationsToObjectTypeGetAll**](docs/AssociationsApi.md#getcrmv3objectsdealsdealidassociationstoobjecttypegetall) | **Get** /crm/v3/objects/deals/{dealId}/associations/{toObjectType} | List associations of a deal by type
*AssociationsApi* | [**PutCrmV3ObjectsDealsDealIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate**](docs/AssociationsApi.md#putcrmv3objectsdealsdealidassociationstoobjecttypetoobjectidassociationtypecreate) | **Put** /crm/v3/objects/deals/{dealId}/associations/{toObjectType}/{toObjectId}/{associationType} | Associate a deal with another object
*BasicApi* | [**DeleteCrmV3ObjectsDealsDealIdArchive**](docs/BasicApi.md#deletecrmv3objectsdealsdealidarchive) | **Delete** /crm/v3/objects/deals/{dealId} | Archive
*BasicApi* | [**GetCrmV3ObjectsDealsDealIdGetById**](docs/BasicApi.md#getcrmv3objectsdealsdealidgetbyid) | **Get** /crm/v3/objects/deals/{dealId} | Read
*BasicApi* | [**GetCrmV3ObjectsDealsGetPage**](docs/BasicApi.md#getcrmv3objectsdealsgetpage) | **Get** /crm/v3/objects/deals | List
*BasicApi* | [**PatchCrmV3ObjectsDealsDealIdUpdate**](docs/BasicApi.md#patchcrmv3objectsdealsdealidupdate) | **Patch** /crm/v3/objects/deals/{dealId} | Update
*BasicApi* | [**PostCrmV3ObjectsDealsCreate**](docs/BasicApi.md#postcrmv3objectsdealscreate) | **Post** /crm/v3/objects/deals | Create
*BatchApi* | [**PostCrmV3ObjectsDealsBatchArchiveArchive**](docs/BatchApi.md#postcrmv3objectsdealsbatcharchivearchive) | **Post** /crm/v3/objects/deals/batch/archive | Archive a batch of deals by ID
*BatchApi* | [**PostCrmV3ObjectsDealsBatchCreateCreate**](docs/BatchApi.md#postcrmv3objectsdealsbatchcreatecreate) | **Post** /crm/v3/objects/deals/batch/create | Create a batch of deals
*BatchApi* | [**PostCrmV3ObjectsDealsBatchReadRead**](docs/BatchApi.md#postcrmv3objectsdealsbatchreadread) | **Post** /crm/v3/objects/deals/batch/read | Read a batch of deals by internal ID, or unique property values
*BatchApi* | [**PostCrmV3ObjectsDealsBatchUpdateUpdate**](docs/BatchApi.md#postcrmv3objectsdealsbatchupdateupdate) | **Post** /crm/v3/objects/deals/batch/update | Update a batch of deals
*PublicObjectApi* | [**PostCrmV3ObjectsDealsMergeMerge**](docs/PublicObjectApi.md#postcrmv3objectsdealsmergemerge) | **Post** /crm/v3/objects/deals/merge | Merge two deals with same type
*SearchApi* | [**PostCrmV3ObjectsDealsSearchDoSearch**](docs/SearchApi.md#postcrmv3objectsdealssearchdosearch) | **Post** /crm/v3/objects/deals/search | 


## Documentation For Models

 - [AssociatedId](docs/AssociatedId.md)
 - [BatchInputSimplePublicObjectBatchInput](docs/BatchInputSimplePublicObjectBatchInput.md)
 - [BatchInputSimplePublicObjectId](docs/BatchInputSimplePublicObjectId.md)
 - [BatchInputSimplePublicObjectInput](docs/BatchInputSimplePublicObjectInput.md)
 - [BatchReadInputSimplePublicObjectId](docs/BatchReadInputSimplePublicObjectId.md)
 - [BatchResponseSimplePublicObject](docs/BatchResponseSimplePublicObject.md)
 - [BatchResponseSimplePublicObjectWithErrors](docs/BatchResponseSimplePublicObjectWithErrors.md)
 - [CollectionResponseAssociatedId](docs/CollectionResponseAssociatedId.md)
 - [CollectionResponseAssociatedIdForwardPaging](docs/CollectionResponseAssociatedIdForwardPaging.md)
 - [CollectionResponseSimplePublicObjectWithAssociationsForwardPaging](docs/CollectionResponseSimplePublicObjectWithAssociationsForwardPaging.md)
 - [CollectionResponseWithTotalSimplePublicObjectForwardPaging](docs/CollectionResponseWithTotalSimplePublicObjectForwardPaging.md)
 - [Error](docs/Error.md)
 - [ErrorCategory](docs/ErrorCategory.md)
 - [ErrorDetail](docs/ErrorDetail.md)
 - [Filter](docs/Filter.md)
 - [FilterGroup](docs/FilterGroup.md)
 - [ForwardPaging](docs/ForwardPaging.md)
 - [NextPage](docs/NextPage.md)
 - [Paging](docs/Paging.md)
 - [PreviousPage](docs/PreviousPage.md)
 - [PublicMergeInput](docs/PublicMergeInput.md)
 - [PublicObjectSearchRequest](docs/PublicObjectSearchRequest.md)
 - [SimplePublicObject](docs/SimplePublicObject.md)
 - [SimplePublicObjectBatchInput](docs/SimplePublicObjectBatchInput.md)
 - [SimplePublicObjectId](docs/SimplePublicObjectId.md)
 - [SimplePublicObjectInput](docs/SimplePublicObjectInput.md)
 - [SimplePublicObjectWithAssociations](docs/SimplePublicObjectWithAssociations.md)
 - [StandardError](docs/StandardError.md)
 - [ValueWithTimestamp](docs/ValueWithTimestamp.md)


## Documentation For Authorization



### oauth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://app.hubspot.com/oauth/authorize
- **Scopes**: 
 - **crm.objects.deals.read**:  
 - **crm.objects.deals.write**:  

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


### oauth2_legacy


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://app.hubspot.com/oauth/authorize
- **Scopes**: 
 - **contacts**: Read from and write to my Contacts

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



