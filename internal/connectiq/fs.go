package connectiq

import (
	"os"
	"path"
	"regexp"
	"runtime"

	"github.com/pkg/errors"
)

// Get the folder where Garmin information is stored
func RootGarminFolder() (string, error) {
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

func SDKsFolder() (string, error) {
	root, err := RootGarminFolder()
	if err != nil {
		return "", errors.WithMessage(err, "could not find sdks folder")
	}
	return path.Join(root, "Sdks"), nil
}

var sdkVersionRegexp = regexp.MustCompile(`^connectiq-sdk-\w+-([\d.]+)-`)

func SDKVersionFromFilename(name string) string {
	match := sdkVersionRegexp.FindStringSubmatch(name)
	if match == nil {
		return ""
	}

	return match[1]
}
