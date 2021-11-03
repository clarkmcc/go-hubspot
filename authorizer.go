package hubspot

import (
	"context"
	"github.com/clarkmcc/go-hubspot/generated/contacts"
)

// Authorizer knows how to authorize API requests to the HubSpot API by modifying the request context
type Authorizer interface {
	Apply(ctx context.Context) context.Context
}

// Does APIKeyAuthorizer implement Authorizer?
var _ Authorizer = &APIKeyAuthorizer{}

// APIKeyAuthorizer is an Authorizer that knows how to apply API keys to requests
type APIKeyAuthorizer struct {
	Key string
}

func (a *APIKeyAuthorizer) Apply(ctx context.Context) context.Context {
	// We'll use the types from the contacts generated package for now since the authorizer stuff here
	// is the same across all the generated packages.
	return context.WithValue(ctx, contacts.ContextAPIKeys, map[string]contacts.APIKey{
		"hapikey": {Key: a.Key},
	})
}

func NewAPIKeyAuthorizer(key string) *APIKeyAuthorizer {
	return &APIKeyAuthorizer{Key: key}
}

var _ Authorizer = &noopAuthorizer{}

// noopAuthorizer is a default authorizer that does not modify the request context
type noopAuthorizer struct{}

func (n noopAuthorizer) Apply(ctx context.Context) context.Context {
	return ctx
}
