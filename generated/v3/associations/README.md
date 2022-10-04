# Go API client for associations

Associations define the relationships between objects in HubSpot. These endpoints allow you to create, read, and remove associations.

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
import associations "github.com/GIT_USER_ID/GIT_REPO_ID"
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
ctx := context.WithValue(context.Background(), associations.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), associations.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), associations.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), associations.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.hubapi.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BatchApi* | [**PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchArchiveArchive**](docs/BatchApi.md#postcrmv3associationsfromobjecttypetoobjecttypebatcharchivearchive) | **Post** /crm/v3/associations/{fromObjectType}/{toObjectType}/batch/archive | Archive a batch of associations
*BatchApi* | [**PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchCreateCreate**](docs/BatchApi.md#postcrmv3associationsfromobjecttypetoobjecttypebatchcreatecreate) | **Post** /crm/v3/associations/{fromObjectType}/{toObjectType}/batch/create | Create a batch of associations
*BatchApi* | [**PostCrmV3AssociationsFromObjectTypeToObjectTypeBatchReadRead**](docs/BatchApi.md#postcrmv3associationsfromobjecttypetoobjecttypebatchreadread) | **Post** /crm/v3/associations/{fromObjectType}/{toObjectType}/batch/read | Read a batch of associations
*TypesApi* | [**GetCrmV3AssociationsFromObjectTypeToObjectTypeTypesGetAll**](docs/TypesApi.md#getcrmv3associationsfromobjecttypetoobjecttypetypesgetall) | **Get** /crm/v3/associations/{fromObjectType}/{toObjectType}/types | List association types


## Documentation For Models

 - [AssociatedId](docs/AssociatedId.md)
 - [BatchInputPublicAssociation](docs/BatchInputPublicAssociation.md)
 - [BatchInputPublicObjectId](docs/BatchInputPublicObjectId.md)
 - [BatchResponsePublicAssociation](docs/BatchResponsePublicAssociation.md)
 - [BatchResponsePublicAssociationMulti](docs/BatchResponsePublicAssociationMulti.md)
 - [CollectionResponsePublicAssociationDefiniton](docs/CollectionResponsePublicAssociationDefiniton.md)
 - [Error](docs/Error.md)
 - [ErrorCategory](docs/ErrorCategory.md)
 - [ErrorDetail](docs/ErrorDetail.md)
 - [NextPage](docs/NextPage.md)
 - [Paging](docs/Paging.md)
 - [PublicAssociation](docs/PublicAssociation.md)
 - [PublicAssociationDefiniton](docs/PublicAssociationDefiniton.md)
 - [PublicAssociationMulti](docs/PublicAssociationMulti.md)
 - [PublicObjectId](docs/PublicObjectId.md)
 - [StandardError](docs/StandardError.md)


## Documentation For Authorization



### hapikey

- **Type**: API key
- **API key parameter name**: hapikey
- **Location**: URL query string

Note, each API key must be added to a map of `map[string]APIKey` where the key is: hapikey and passed in as the auth context for each request.


### oauth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://app.hubspot.com/oauth/authorize
- **Scopes**: 
 - **crm.objects.quotes.write**: Quotes
 - **crm.objects.contacts.read**:  
 - **crm.objects.contacts.write**:  
 - **crm.objects.companies.write**:  
 - **crm.objects.companies.read**:  
 - **crm.objects.line_items.write**: Line Items
 - **crm.objects.quotes.read**: Quotes
 - **crm.objects.deals.write**:  
 - **crm.objects.line_items.read**: Line Items
 - **crm.objects.deals.read**:  

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
 - **crm.objects.custom.read**: View custom object records
 - **contacts**: Read from and write to my Contacts
 - **crm.objects.custom.write**: Change custom object records
 - **e-commerce**: e-commerce
 - **media_bridge.read**: Read media and media events
 - **tickets**: Read and write tickets

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


