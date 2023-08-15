package cmd

import (
	"context"

	"github.com/Masterminds/semver/v3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// SDKDownloadCmd downloads an sdk
func SDKDownloadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download version",
		Short: "Download SDK",
		Long: `Download SDK.

The version argument can be a specific version or a semver-range.
For example: ^6.2.0 or >=4.0.0 or 4.2.1`,
		Args: cobra.ExactArgs(1),
		RunE: downloadSDKs,
	}

	return cmd
}

func downloadSDKs(_ *cobra.Command, args []string) error {
	ctx := context.Background()

	mngr, err := NewManager()
	if err != nil {
		return err
	}

	semverConstraint, err := semver.NewConstraint(args[0])
	if err != nil {
		return errors.WithMessage(err, "could not parse version contraint")
	}

	return mngr.DownloadSDK(ctx, semverConstraint)
}
