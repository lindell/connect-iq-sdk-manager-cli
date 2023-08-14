package cmd

import (
	"context"

	"github.com/lindell/connect-iq-sdk-manager-cli/internal/manager"
	"github.com/spf13/cobra"
)

// DeviceDownloadCmd downloads devices
func DeviceDownloadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download",
		Short: "Download resources.",
		Long:  "Download resources.",
		Args:  cobra.NoArgs,
		RunE:  download,
	}

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
