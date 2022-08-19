package hubspot

import (
	"context"
	"fmt"
	"net/url"
)

var ContextKey = "hubspot.authorizer"

// WithAuthorizer appends the given authorizer to the context
func WithAuthorizer(ctx context.Context, authorizer Authorizer) context.Context {
	return context.WithValue(ctx, ContextKey, authorizer)
}

// AuthorizationRequest represents a request that could be authorized but is not yet authorized.
// It only contains the fields that an authorizer may need to modify.
type AuthorizationRequest struct {
	QueryParams url.Values
	FormParams  url.Values
	Headers     map[string]string
}

// Authorizer knows how to authorize API requests to the HubSpot API by modifying the request context
type Authorizer interface {
	Apply(request AuthorizationRequest)
}

var _ Authorizer = &TokenAuthorizer{}

type TokenAuthorizer struct {
	Token string
}

func (a *TokenAuthorizer) Apply(request AuthorizationRequest) {
	request.Headers["Authorization"] = fmt.Sprintf("Bearer %v", a.Token)
}

func NewTokenAuthorizer(token string) *TokenAuthorizer {
	return &TokenAuthorizer{Token: token}
}

type APIKeyAuthorizer struct {
	Key string
}

func (a *APIKeyAuthorizer) Apply(request AuthorizationRequest) {
	request.QueryParams.Set("hapikey", a.Key)
}

func NewAPIKeyAuthorizer(key string) *APIKeyAuthorizer {
	return &APIKeyAuthorizer{Key: key}
}

var _ Authorizer = &noopAuthorizer{}

// noopAuthorizer is a default authorizer that does not modify the request
type noopAuthorizer struct{}

func (n noopAuthorizer) Apply(request AuthorizationRequest) {}
