package manager

import (
	"context"
	"time"

	"github.com/lindell/connect-iq-sdk-manager-cli/internal/client"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/connectiq"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/datetime"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/storage"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Manager struct {
	Store *storage.Store
}

type ManagerConfig struct {
	SkipAgreementCheck bool
	SkipLoginRequired  bool
}

func NewManager(ctx context.Context, config ManagerConfig) (Manager, context.Context, error) {
	store := storage.NewStore()
	mngr := Manager{
		Store: store,
	}

	if !config.SkipAgreementCheck {
		if err := checkAgreement(ctx); err != nil {
			return Manager{}, ctx, err
		}
	}

	if !config.SkipLoginRequired {
		var err error
		ctx, err = mngr.setTokenToCtx(ctx)
		if err != nil {
			return Manager{}, ctx, err
		}
	}

	return mngr, ctx, nil
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

func checkAgreement(ctx context.Context) error {
	vals := connectiq.LoadConfigVals("agreement-hash", "agreement-hash-verified-at")
	hash := vals[0]
	acceptedAtStr := vals[1]

	if hash == "" {
		return errors.New("agreement is not accepted. Please do so with the `agreement` command")
	}

	// If acceptance was made within one our, don't make a request to make sure the agreement hasn't changed.
	if acceptedAtStr != "" {
		acceptedAt, err := datetime.Parse(acceptedAtStr)
		if err != nil {
			return err
		}
		if time.Now().Before(acceptedAt.Time().Add(time.Hour)) {
			return nil
		}
	}

	currentHash, err := client.AgreementHash(ctx)
	if err != nil {
		return err
	}

	if currentHash != hash {
		return errors.Errorf("accepted agreement is to old. Old hash is %q, new hash is %q", hash, currentHash)
	}

	return nil
}
