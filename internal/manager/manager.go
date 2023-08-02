package manager

import (
	"context"

	"github.com/lindell/connect-iq-manager/internal/connectiq"
	"github.com/lindell/connect-iq-manager/internal/storage"
)

type Manager struct {
	Store *storage.Store
}

func (m *Manager) setTokenToCtx(ctx context.Context) (context.Context, error) {
	token, err := m.Store.GetToken()
	if err != nil {
		return ctx, err
	}
	return connectiq.SetContextToken(ctx, token), nil
}
