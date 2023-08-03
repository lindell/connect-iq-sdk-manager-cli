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
	path, err := connectiq.RootGarminFolder()
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
		return errors.WithMessage(err, "could not create Garmin root folder")
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

func (s *Store) GetToken() (connectiq.Token, error) {
	bb, err := os.ReadFile(path.Join(s.rootPath, tokenFilename))
	if err != nil {
		return connectiq.Token{}, errors.WithMessage(err, "could not read token file")
	}

	var token connectiq.Token
	if err := json.Unmarshal(bb, &token); err != nil {
		return connectiq.Token{}, errors.WithMessage(err, "could not decode token")
	}

	return token, nil
}

func ensureFolderExist(path string) error {
	return os.MkdirAll(path, 0755)
}
