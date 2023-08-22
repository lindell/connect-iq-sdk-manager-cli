package cmd

import (
	"context"

	"github.com/Masterminds/semver/v3"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/manager"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const versionDesc = `The version argument can be a specific version or a semver-range.
For example: ^6.2.0 or >=4.0.0 or 4.2.1`

// SDKDownloadCmd downloads an sdk
func SDKDownloadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download version",
		Short: "Download an SDK. Without setting it as the current one.",
		Long: `Download an SDK. Without setting it as the current one.

` + versionDesc,
		Args: cobra.ExactArgs(1),
		RunE: downloadSDKs,
	}

	return cmd
}

func downloadSDKs(_ *cobra.Command, args []string) error {
	ctx := context.Background()
	mngr, ctx, err := manager.NewManager(ctx, manager.ManagerConfig{
		SkipLoginRequired: true,
	})
	if err != nil {
		return err
	}

	semverConstraint, err := semver.NewConstraint(args[0])
	if err != nil {
		return errors.WithMessage(err, "could not parse version contraint")
	}

	return mngr.DownloadSDK(ctx, semverConstraint)
}
