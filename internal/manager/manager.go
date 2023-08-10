package manager

import (
	"context"
	"errors"
	"time"

	"github.com/lindell/connect-iq-sdk-manager-cli/internal/client"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/connectiq"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/storage"
	log "github.com/sirupsen/logrus"
)

type Manager struct {
	Store *storage.Store
}

func (m *Manager) setTokenToCtx(ctx context.Context) (context.Context, error) {
	token, err := m.Store.GetToken()
	if err != nil {
		if errors.Is(err, storage.ErrTokenNotFound) {
			return ctx, errors.New("not logged in, please do so with the `login` command")
		}
		return ctx, err
	}

	if token.RefreshTokenExpiresAt.Before(time.Now()) {
		return ctx, errors.New("token has expired")
	}

	// Refresh the token if it is about to expire
	if token.ExpiresAt.Before(time.Now().AddDate(0, 0, 1)) {
		log.Info("Refresh token is about to expire, try to refresh")
		newToken, err := m.refreshAndSaveToken(ctx, token)
		if err != nil {
			log.Errorf("Could not refresh token: %s", err)
		} else {
			token = newToken
		}
	}

	return connectiq.SetContextToken(ctx, token), nil
}

func (m *Manager) refreshAndSaveToken(ctx context.Context, token connectiq.Token) (connectiq.Token, error) {
	newToken, err := client.RefreshToken(ctx, token.RefreshToken)
	if err != nil {
		return connectiq.Token{}, err
	}

	if err := m.Store.StoreToken(newToken); err != nil {
		return connectiq.Token{}, err
	}

	return newToken, nil
}
