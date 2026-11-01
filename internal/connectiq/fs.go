package connectiq

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"

	"github.com/pkg/errors"
)

func init() {
	if _, err := rootGarminFolder(); err != nil {
		fmt.Fprintf(os.Stderr, "could not find Garmin folder: %s", err)
		os.Exit(1)
	}
}

var RootPath, _ = rootGarminFolder()
var SDKsPath = filepath.Join(RootPath, "Sdks")
var FontsPath = filepath.Join(RootPath, "Fonts")
var DevicesPath = filepath.Join(RootPath, "Devices")
var CurrentSDKPath = filepath.Join(RootPath, "current-sdk.cfg")
var ConfigPath = filepath.Join(RootPath, "sdkmanager-config.ini")

// Get the folder where Garmin information is stored
func rootGarminFolder() (string, error) {
	if runtime.GOOS == "windows" {
		appDataFolder := os.Getenv("APPDATA")
		if appDataFolder == "" {
			return "", errors.New("could not find appdata folder")
		}
		return filepath.Join(appDataFolder, "Garmin", "ConnectIQ"), nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.WithMessage(err, "could not get the home directory")
	}
	if runtime.GOOS == "darwin" {
		homeDir = filepath.Join(homeDir, "Library", "Application Support", "Garmin")
	} else {
		homeDir = filepath.Join(homeDir, ".Garmin")
	}
	return filepath.Join(homeDir, "ConnectIQ"), nil
}

var sdkVersionRegexp = regexp.MustCompile(`^connectiq-sdk-\w+-([\d.]+)-`)

func SDKVersionFromFilename(name string) string {
	match := sdkVersionRegexp.FindStringSubmatch(name)
	if match == nil {
		return ""
	}

	return match[1]
}
