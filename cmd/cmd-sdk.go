package cmd

import "github.com/spf13/cobra"

// SdkCmd is the root sdk command
func SdkCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sdk",
		Short: "Handle sdks",
	}

	cmd.AddCommand(SDKListCmd())
	cmd.AddCommand(SDKDownloadCmd())

	return cmd
}
