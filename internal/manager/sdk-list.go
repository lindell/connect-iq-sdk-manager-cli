package manager

import (
	"context"
	"os"

	"github.com/Masterminds/semver/v3"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/client"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/connectiq"
	"github.com/pkg/errors"
)

func (m *Manager) ListSDKs(ctx context.Context, semverConstraint *semver.Constraints) error {
	sdks, err := client.GetSDKs(ctx)
	if err != nil {
		return errors.WithMessage(err, "could not fetch SDKs")
	}
	sdks = filterAndSortSDKs(sdks, semverConstraint)

	installedSDKs, err := installedSDKs()
	if err != nil {
		return err
	}

	table := createTable()
	table.SetHeader([]string{"VERSION", "RELEASE", "INSTALLED"})
	for _, s := range sdks {
		table.Append(sdkToRow(s, installedSDKs))
	}
	table.Render()

	return nil
}

func sdkToRow(device client.SDK, installedSDKs mapset.Set[string]) []string {
	installed := "No"
	if installedSDKs.Contains(device.Version) {
		installed = "Yes"
	}

	return []string{
		device.Version,
		device.Release,
		installed,
	}
}

func installedSDKs() (mapset.Set[string], error) {
	sdksFolder, err := connectiq.SDKsFolder()
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(sdksFolder)
	if err != nil {
		return nil, errors.WithMessage(err, "could not read sdks folder")
	}

	versions := mapset.NewSet[string]()
	for _, e := range entries {
		version := connectiq.SDKVersionFromFilename(e.Name())
		if version != "" {
			versions.Add(version)
		}
	}

	return versions, nil
}
