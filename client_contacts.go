/*
Copyright (C) 2020 Print Tracker, LLC - All Rights Reserved

Unauthorized copying of this file, via any medium is strictly prohibited
as this source code is proprietary and confidential. Dissemination of this
information or reproduction of this material is strictly forbidden unless
prior written permission is obtained from Print Tracker, LLC.
*/

package hubspot

import (
	"context"
	"github.com/clarkmcc/go-hubspot/generated/contacts"
)

type ClientContactsInterface interface {
	GetByID(ctx context.Context, id string, params *GetContactByIdParams) (*SimplePublicObjectWithAssociations, error)
}

type GetContactByIdParams = contacts.BasicApiGetcrmv3objectscontactscontactIdGetByIdOpts
type SimplePublicObjectWithAssociations = contacts.SimplePublicObjectWithAssociations

var _ ClientContactsInterface = &clientContacts{}

// clientContacts is a wrapper around the generated contacts.APIClient
type clientContacts struct {
	client   *Client
	internal *contacts.APIClient
}

func (c *clientContacts) GetByID(ctx context.Context, id string, params *GetContactByIdParams) (*SimplePublicObjectWithAssociations, error) {
	ctx = c.client.authorizer.Apply(ctx)
	obj, _, err := c.internal.BasicApi.Getcrmv3objectscontactscontactIdGetById(ctx, id, params)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func NewClientContacts(client *Client) *clientContacts {
	return &clientContacts{client: client, internal: contacts.NewAPIClient(contacts.NewConfiguration())}
}
