package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd is the root command containing all subcommands
func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "connect-iq-sdk-manager",
		Short: "CLI to download connectIQ resources",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Bind configs that are not flags
			if err := initializeConfig(cmd); err != nil {
				return err
			}

			return logFlagInit(cmd)
		},
	}

	cmd.AddCommand(VersionCmd())
	cmd.AddCommand(LoginCmd())
	cmd.AddCommand(DeviceCmd())
	cmd.AddCommand(SdkCmd())
	cmd.AddCommand(AgreementCmd())

	configureConfig(cmd)
	configureLogging(cmd)

	return cmd
}
