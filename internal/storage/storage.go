package storage

import (
	"encoding/json"
	"os"
	"path"

	"github.com/lindell/connect-iq-manager/internal/connectiq"
	"github.com/pkg/errors"
)

type Store struct {
	rootPath string
}

func NewStore() (*Store, error) {
	path, err := rootGarminFolder()
	if err != nil {
		return nil, err
	}

	return &Store{
		rootPath: path,
	}, nil
}

var tokenFilename = "token.json"

func (s *Store) StoreToken(token connectiq.Token) error {
	if err := ensureFolderExist(s.rootPath); err != nil {
		return err
	}

	encodedData, err := json.Marshal(token)
	if err != nil {
		return err
	}

	err = os.WriteFile(path.Join(s.rootPath, tokenFilename), encodedData, 0600)
	if err != nil {
		return errors.WithMessage(err, "could not write token to file")
	}

	return nil
}

func ensureFolderExist(path string) error {
	return os.MkdirAll(path, 0600)
}
