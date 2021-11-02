# HubSpot Go Client
This is a HubSpot Go client generated using [swagger-codegen](https://github.com/swagger-api/swagger-codegen#getting-started). It includes packages for every OpenAPI spec listed in HubSpot's API [directory](https://api.hubspot.com/api-catalog-public/v1/apis).

## Re-generating Clients
This package includes a client generator that runs the `swagger-codegen` binary (must be located in your PATH) and re-generates the Go clients using the most up-to-date API specs. Any changes to the specs should be committed to this repository through a PR. I'll do my best to keep this up to date with the most recent changes. Assuming that you have `swagger-codegen` installed, just run the following to re-generate the clients.

```shell
go generate
```

## Using the clients
The clients are broken out into their own packages by name. Currently, `swagger-codegen` produces mutilated method names like `Deletecrmv3objectscontactscontactIdArchive`. This will need to be fixed in the `swagger-codegen` project at some point.
```go
client := contacts.NewAPIClient(contacts.NewConfiguration())
```

## Client Wrappers
I've been toying with the idea of creating more idiomatic wrapper functions for each client. I've started down this path with the `client.go` and `client_contacts.go` files in the root of the repository. I'm currently exploring options for generating more idiomatic wrappers for each generated method rather than hand-writing each one.

## Authorization
Authorization is done by passing context values in with each request. In order to facilitate this, I've created a simple `Authorizer` interface and an API key implementation (more implementations to come).
```go
ctx := context.Background()
authorizer := hubspot.NewAPIKeyAuthorizer("aa-bb-cc-dd")
client.DoRequest(authorizer.Apply(ctx), ...) // this request is now authorized
```