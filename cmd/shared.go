package cmd

import (
	"strings"

	"github.com/lindell/connect-iq-sdk-manager-cli/internal/manager"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/storage"
	"github.com/pkg/errors"
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

func checkExlusivity(options map[string]bool) error {
	enabledOptions := []string{}
	for name, val := range options {
		if val {
			enabledOptions = append(enabledOptions, name)
		}
	}

	if len(enabledOptions) <= 1 {
		return nil
	}

	return errors.Errorf("can't define more than one of %s", strings.Join(enabledOptions, ", "))
}
