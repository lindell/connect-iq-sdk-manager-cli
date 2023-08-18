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

	cmd.Flags().BoolP("include-fonts", "F", false, "Download the fonts used for simulating the downloaded devices.")

	return cmd
}

func download(cmd *cobra.Command, _ []string) error {
	ctx := context.Background()

	includeFonts, _ := cmd.Flags().GetBool("include-fonts")

	mngr := NewManager()

	deviceFilters, err := getDeviceFilters(cmd)
	if err != nil {
		return err
	}

	return mngr.Download(ctx, manager.DownloadConfig{
		DeviceFilters: deviceFilters,
		IncludeFonts:  includeFonts,
	})
}
