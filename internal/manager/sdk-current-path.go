package manager

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lindell/connect-iq-sdk-manager-cli/internal/connectiq"
)

func (m *Manager) CurrentSDKPath(_ context.Context, binary bool) error {
	if _, err := os.Stat(connectiq.CurrentSDKPath); os.IsNotExist(err) {
		fmt.Println("No SDK set as current")
		return nil
	} else if err != nil {
		return err
	}

	bb, err := os.ReadFile(connectiq.CurrentSDKPath)
	if err != nil {
		return err
	}
	path := string(bb)

	if binary {
		path = filepath.Join(path, "bin")
	}

	fmt.Printf("%s\n", path)
	return nil
}
