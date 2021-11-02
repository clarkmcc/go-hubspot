package hubspot

import (
	"context"
)

type Client struct {
	authorizer Authorizer

	// All the wrapper clients
	contacts *clientContacts
}

// Authorize can be used to apply the same authorization scheme to clients not managed under this type. For
// example if there is an API endpoint that is not currently supported by this wrapper client, you can use
// the generated API endpoint directly, then call this function to apply the same authorization schemes used
// on this client, to requests to the generated endpoints.
func (c *Client) Authorize(ctx context.Context) context.Context {
	return c.authorizer.Apply(ctx)
}

// Contacts returns the Contacts' wrapper client
func (c *Client) Contacts() ClientContactsInterface {
	return c.contacts
}

func (c *Client) WithAPIKey(key string) *Client {
	c.authorizer = &APIKeyAuthorizer{Key: key}
	return c
}

// NewClient creates a new client with a default authorizer
func NewClient() *Client {
	c := &Client{
		authorizer: noopAuthorizer{},
	}
	c.contacts = NewClientContacts(c)
	return c
}
