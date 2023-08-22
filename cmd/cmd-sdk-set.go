package cmd

import (
	"context"

	"github.com/Masterminds/semver/v3"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/manager"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// SDKSetCmd sets the current version of the SDK
func SDKSetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set version",
		Short: "Set which SDK version to be used.",
		Long: `Set which SDK version to be used. If it does not exist, it will be downloaded.

` + versionDesc,
		Args: cobra.ExactArgs(1),
		RunE: setSDK,
	}

	return cmd
}

func setSDK(_ *cobra.Command, args []string) error {
	ctx := context.Background()
	mngr, ctx, err := manager.NewManager(ctx, manager.Config{
		SkipLoginRequired: true,
	})
	if err != nil {
		return err
	}

	semverConstraint, err := semver.NewConstraint(args[0])
	if err != nil {
		return errors.WithMessage(err, "could not parse version constraint")
	}

	return mngr.SetSDK(ctx, semverConstraint)
}
