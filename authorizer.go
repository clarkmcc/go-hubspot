package hubspot

import (
	"context"
	"fmt"
	"net/http"
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

// ApplyAuthorizerFromContext looks up an Authorizer stored on ctx (via
// WithAuthorizer) and applies it to req's headers and query string. It is a
// convenience used by the generated clients so the per-client code does not need
// to reference net/url (which the generated prepareRequest shadows with a local
// variable). It is a no-op if no Authorizer is present on the context.
func ApplyAuthorizerFromContext(ctx context.Context, req *http.Request) {
	if ctx == nil || req == nil {
		return
	}
	authz, ok := ctx.Value(ContextKey).(Authorizer)
	if !ok || authz == nil {
		return
	}
	r := AuthorizationRequest{
		QueryParams: req.URL.Query(),
		FormParams:  url.Values{},
		Headers:     map[string]string{},
	}
	authz.Apply(r)
	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}
	if len(r.QueryParams) > 0 {
		req.URL.RawQuery = r.QueryParams.Encode()
	}
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
