![](assets/go-hubspot-client-banner-01.png)
This is a HubSpot Go client generated using [openapi-generator](https://github.com/OpenAPITools/openapi-generator). It includes packages for every OpenAPI spec listed in HubSpot's API [directory](https://api.hubspot.com/api-catalog-public/v1/apis).

## Installing
```shell
go get github.com/clarkmcc/go-hubspot
```

## Re-generating Clients
This package includes a client generator that runs the `openapi-generator` binary (must be located in your PATH) and re-generates the Go clients using the most up-to-date API specs. Any changes to the specs should be committed to this repository through a PR. I'll do my best to keep this up to date with the most recent changes. Assuming that you have `openapi-generator` installed, just run the following to re-generate the clients.

```shell
go generate
```

## Using the clients
The clients are broken out into their own packages by name. Currently, `openapi-generator` produces mutilated method names like `GetCrmV3ObjectsContactsContactIdGetById`. This will need to be fixed in the `openapi-generator` project at some point.
```go
client := contacts.NewAPIClient(contacts.NewConfiguration())
```

## Authorization
Authorization is done by passing context values in with each request. In order to facilitate this, I've created a simple `Authorizer` interface and an API key implementation (more implementations to come).
```go
ctx := context.Background()
authorizer := hubspot.NewAPIKeyAuthorizer("aa-bb-cc-dd")
client.DoRequest(authorizer.Apply(ctx), ...) // this request is now authorized
```