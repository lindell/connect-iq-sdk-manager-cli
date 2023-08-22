package cmd

import "github.com/spf13/cobra"

// AgreementCmd is the root agreement command
func AgreementCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "agreement",
		Short: "Handles the agreement.",
	}

	cmd.AddCommand(AgreementViewCmd())
	cmd.AddCommand(AgreementAcceptCmd())

	return cmd
}
