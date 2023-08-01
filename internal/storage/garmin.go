package storage

import (
	"os"
	"path"
	"runtime"

	"github.com/pkg/errors"
)

func rootGarminFolder() (string, error) {
	if runtime.GOOS == "windows" {
		appDataFolder := os.Getenv("APPDATA")
		if appDataFolder == "" {
			return "", errors.New("could not find appdata folder")
		}
		return path.Join(appDataFolder, "Garmin", "ConnectIQ"), nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.WithMessage(err, "could not get the home directory")
	}
	return path.Join(homeDir, ".Garmin", "ConnectIQ"), nil
}
