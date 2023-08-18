package manager

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lindell/connect-iq-sdk-manager-cli/internal/connectiq"
)

func (m *Manager) CurrentSDKPath(_ context.Context, binary bool) error {
	garminRoot, err := connectiq.RootGarminFolder()
	if err != nil {
		return err
	}
	currentSDKFilePath := filepath.Join(garminRoot, "current-sdk.cfg")

	if _, err := os.Stat(currentSDKFilePath); os.IsNotExist(err) {
		fmt.Println("No SDK set as current")
		return nil
	} else if err != nil {
		return err
	}

	bb, err := os.ReadFile(currentSDKFilePath)
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
