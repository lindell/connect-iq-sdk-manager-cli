package cmd

import (
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/manager"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/storage"
)

func NewManager() (manager.Manager, error) {
	store, err := storage.NewStore()
	if err != nil {
		return manager.Manager{}, err
	}

	return manager.Manager{
		Store: store,
	}, nil
}
