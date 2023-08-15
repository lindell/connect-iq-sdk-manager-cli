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

	// Find where to place the sdk in the filesystem
	rootFolder, err := connectiq.RootGarminFolder()
	if err != nil {
		return err
	}
	filename, err := sdk.Filename()
	if err != nil {
		return err
	}
	sdkDir := path.Join(rootFolder, "Sdks", strings.TrimSuffix(filename, ".zip"))

	if _, err := os.Stat(sdkDir); !os.IsNotExist(err) {
		log.Info("SDK folder already exist, skip download")
		return nil
	}

	r, err := client.DownloadSDK(ctx, sdk)
	if err != nil {
		return err
	}
	defer r.Close()

	return fetchAndExtract(r, sdkDir)
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
