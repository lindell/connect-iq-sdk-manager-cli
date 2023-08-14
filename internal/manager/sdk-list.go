package manager

import (
	"context"

	"github.com/Masterminds/semver/v3"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/client"
	"github.com/pkg/errors"
)

func (m *Manager) ListSdks(ctx context.Context, semverConstraint *semver.Constraints) error {
	sdks, err := client.GetSDKs(ctx)
	if err != nil {
		return errors.WithMessage(err, "could not fetch SDKs")
	}
	sdks = filterAndSortSDKs(sdks, semverConstraint)

	table := createTable()
	table.SetHeader([]string{"VERSION", "RELEASE"})
	for _, s := range sdks {
		table.Append(sdkToRow(s))
	}
	table.Render()

	return nil
}

func sdkToRow(device client.SDK) []string {
	return []string{
		device.Version,
		device.Release,
	}
}
