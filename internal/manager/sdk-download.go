package manager

import (
	"context"
	"io"
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
	// Find the SDK to download
	sdks, err := client.GetSDKs(ctx)
	if err != nil {
		return errors.WithMessage(err, "could not fetch SDKs")
	}
	sdks = filterAndSortSDKs(sdks, semverConstraint)
	if len(sdks) == 0 {
		return errors.Errorf("no SDKs matched %q", semverConstraint.String())
	}
	sdk := sdks[0]

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

	// Save the zip to a temporary file
	f, err := os.CreateTemp(os.TempDir(), "sdk-*.zip")
	if err != nil {
		return errors.WithMessage(err, "could not create tmp device file")
	}
	defer os.Remove(f.Name())
	defer f.Close()
	if _, err := io.Copy(f, r); err != nil {
		return err
	}

	log.Info("Extracting sdk zip")
	return unzip(f.Name(), sdkDir)
}
