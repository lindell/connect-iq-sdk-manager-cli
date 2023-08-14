package cmd

import "github.com/spf13/cobra"

// DeviceCmd is the root device command
func DeviceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "device",
		Short: "Handle devices",
	}

	cmd.AddCommand(DeviceListCmd())
	cmd.AddCommand(DeviceDownloadCmd())

	configureDeviceCmd(cmd)

	return cmd
}
