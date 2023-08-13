package manager

import (
	"context"

	"github.com/lindell/connect-iq-sdk-manager-cli/internal/client"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/oauth"
	"github.com/pkg/errors"
)

func (m *Manager) LoginWithCredentials(ctx context.Context, username, password string) error {
	token, err := client.Login(ctx, username, password)
	if err != nil {
		return errors.WithMessage(err, "could not login")
	}

	err = m.Store.StoreToken(token)
	if err != nil {
		return errors.WithMessage(err, "could not store token")
	}

	return nil
}

func (m *Manager) LoginWithOauth(ctx context.Context) error {
	ticket, serviceURL, err := oauth.GetToken(ctx)
	if err != nil {
		return err
	}

	token, err := client.ExchangeTicket(ctx, ticket, serviceURL)
	if err != nil {
		return err
	}

	err = m.Store.StoreToken(token)
	if err != nil {
		return errors.WithMessage(err, "could not store token")
	}

	return nil
}
