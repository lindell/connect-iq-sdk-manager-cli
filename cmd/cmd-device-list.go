package cmd

import (
	"context"

	"github.com/lindell/connect-iq-sdk-manager-cli/internal/manager"
	"github.com/spf13/cobra"
)

// DeviceListCmd list devices
func DeviceListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List devices",
		Args:  cobra.NoArgs,
		RunE:  listDevices,
	}

	return cmd
}

func listDevices(cmd *cobra.Command, _ []string) error {
	ctx := context.Background()

	mngr, err := NewManager()
	if err != nil {
		return err
	}

	deviceFilters, err := getDeviceFilters(cmd)
	if err != nil {
		return err
	}

	return mngr.ListDevices(ctx, manager.DeviceListConfig{
		DeviceFilters: deviceFilters,
	})
}
