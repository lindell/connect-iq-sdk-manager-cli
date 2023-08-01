package cmd

import (
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

// RootCmd is the root command containing all subcommands
func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "connect-iq-manager",
		Short: "CLI Template", // TODO: Change name
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initializeConfig(cmd) // Bind configs that are not flags
		},
	}

	cmd.AddCommand(VersionCmd())
	configureConfig(cmd)

	return cmd
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
