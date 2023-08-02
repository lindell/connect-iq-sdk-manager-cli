package cmd

import (
	"github.com/lindell/connect-iq-manager/internal/manager"
	"github.com/lindell/connect-iq-manager/internal/storage"
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
