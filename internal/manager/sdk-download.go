package manager

import (
	"context"
	"os"
	"path"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/client"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/connectiq"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (m *Manager) DownloadSDK(ctx context.Context, semverConstraint *semver.Constraints) error {
	sdk, err := latestMatchingSDK(ctx, semverConstraint)
	if err != nil {
		return err
	}

	log.Infof("Downloading %s", sdk.Version)

	sdkDir, err := sdkPath(sdk)
	if err != nil {
		return err
	}

	if _, err := os.Stat(sdkDir); !os.IsNotExist(err) {
		log.Info("SDK folder already exist, skip download")
		return nil
	}

	return downloadSDK(ctx, sdk)
}

func downloadSDK(ctx context.Context, sdk client.SDK) error {
	r, err := client.DownloadSDK(ctx, sdk)
	if err != nil {
		return err
	}
	defer r.Close()

	sdkDir, err := sdkPath(sdk)
	if err != nil {
		return err
	}

	filename, err := sdk.Filename()
	if err != nil {
		return err
	}

	_, err = fetchAndExtract(r, sdkDir, filename)
	return err
}

func sdkPath(sdk client.SDK) (string, error) {
	filename, err := sdk.Filename()
	if err != nil {
		return "", err
	}

	filename = strings.TrimSuffix(filename, ".zip")
	filename = strings.TrimSuffix(filename, ".dmg")

	return path.Join(connectiq.SDKsPath, filename), nil
}

func latestMatchingSDK(ctx context.Context, semverConstraint *semver.Constraints) (client.SDK, error) {
	// Find the SDK to download
	sdks, err := client.GetSDKs(ctx)
	if err != nil {
		return client.SDK{}, errors.WithMessage(err, "could not fetch SDKs")
	}
	sdks = filterAndSortSDKs(sdks, semverConstraint)
	if len(sdks) == 0 {
		return client.SDK{}, errors.Errorf("no SDKs matched %q", semverConstraint.String())
	}
	return sdks[0], nil
}
