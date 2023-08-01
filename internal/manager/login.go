package manager

import (
	"context"

	"github.com/lindell/connect-iq-manager/internal/client"
	"github.com/lindell/connect-iq-manager/internal/storage"
	"github.com/pkg/errors"
)

type Manager struct {
	Store *storage.Store
}

func (m *Manager) Login(ctx context.Context, username, password string) error {
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
