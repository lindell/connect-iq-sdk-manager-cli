package cmd

import (
	"context"

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

	return cmd
}

func currentSDKPath(_ *cobra.Command, _ []string) error {
	ctx := context.Background()

	mngr, err := NewManager()
	if err != nil {
		return err
	}

	return mngr.CurrentSDKPath(ctx)
}
