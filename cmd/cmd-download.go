package cmd

import (
	"context"

	"github.com/lindell/connect-iq-sdk-manager-cli/internal/manager"
	"github.com/spf13/cobra"
)

// DownloadCmd downloads devices
func DownloadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download",
		Short: "Download resources.",
		Long:  "Download resources.",
		Args:  cobra.NoArgs,
		RunE:  download,
	}

	configureDeviceCmd(cmd)

	return cmd
}

func download(cmd *cobra.Command, _ []string) error {
	ctx := context.Background()

	mngr, err := NewManager()
	if err != nil {
		return err
	}

	deviceFilters, err := getDeviceFilters(cmd)
	if err != nil {
		return err
	}

	return mngr.Download(ctx, manager.DownloadConfig{
		DeviceFilters: deviceFilters,
	})
}
