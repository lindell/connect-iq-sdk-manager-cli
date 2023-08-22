package cmd

import (
	"context"

	"github.com/lindell/connect-iq-sdk-manager-cli/internal/manager"
	"github.com/spf13/cobra"
)

// SDKCurrentPathCmd print the path to the currently active SDK
func SDKCurrentPathCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "current-path",
		Short: "Print the path to the currently active SDK",
		Args:  cobra.NoArgs,
		RunE:  currentSDKPath,
	}

	cmd.Flags().BoolP("bin", "", false, "Print binary path")

	return cmd
}

func currentSDKPath(cmd *cobra.Command, _ []string) error {
	ctx := context.Background()
	mngr, ctx, err := manager.NewManager(ctx, manager.Config{
		SkipLoginRequired: true,
	})
	if err != nil {
		return err
	}

	bin, _ := cmd.Flags().GetBool("bin")

	return mngr.CurrentSDKPath(ctx, bin)
}
