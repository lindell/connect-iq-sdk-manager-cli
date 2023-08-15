package manager

import (
	"context"
	"os"
	"path/filepath"

	"github.com/Masterminds/semver/v3"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/connectiq"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (m *Manager) SetSDK(ctx context.Context, semverConstraint *semver.Constraints) error {
	sdk, err := latestMatchingSDK(ctx, semverConstraint)
	if err != nil {
		return err
	}

	log.Infof("Setting %s to be the current SDK", sdk.Version)

	sdkDir, err := sdkPath(sdk)
	if err != nil {
		return err
	}

	if _, err := os.Stat(sdkDir); os.IsNotExist(err) {
		log.Info("SDK does not exist, downloading it")
		if err := downloadSDK(ctx, sdk); err != nil {
			return err
		}
	}

	garminRoot, err := connectiq.RootGarminFolder()
	if err != nil {
		return err
	}
	currentSDKFilePath := filepath.Join(garminRoot, "current-sdk.cfg")

	if err := os.WriteFile(currentSDKFilePath, []byte(sdkDir+string(filepath.Separator)), 0644); err != nil { //nolint: gosec
		return errors.WithMessage(err, "could not write to current sdk file")
	}

	return nil
}
