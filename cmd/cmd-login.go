package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/lindell/connect-iq-sdk-manager-cli/internal/manager"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// LoginCmd logs the user in to the garmin service
func LoginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login to be able to use some parts of the manager.",
		Long: `Login to be able to use some parts of the manager.

If used as is, you will be asked to login via the Garmin SSO OAuth flow.
Credentials can also be set via the --username and --password config,
or GARMIN_USERNAME and GARMIN_PASSWORD environment variables.`,
		Args: cobra.NoArgs,
		RunE: login,
	}

	cmd.Flags().StringP("username", "", "", "The Garmin username.")
	cmd.Flags().StringP("password", "", "", "The Garmin password.")

	return cmd
}

func login(cmd *cobra.Command, _ []string) error {
	ctx := context.Background()
	mngr, ctx, err := manager.NewManager(ctx, manager.ManagerConfig{
		SkipLoginRequired: true,
	})
	if err != nil {
		return err
	}

	username, password := getLoginCredentials(cmd.Flags())

	if username != "" && password != "" {
		log.Debug("Using credentials to simulate oauth login")
		err := mngr.LoginWithCredentials(ctx, username, password)
		if err != nil {
			return err
		}
	} else {
		log.Debug("Using Oauth flow to login")
		err := mngr.LoginWithOauth(ctx)
		if err != nil {
			return err
		}
	}

	fmt.Println("Successfully logged in")

	return nil
}

func getLoginCredentials(flag *pflag.FlagSet) (username, password string) {
	username, _ = flag.GetString("username")
	password, _ = flag.GetString("password")

	if username == "" {
		username = os.Getenv("GARMIN_USERNAME")
	}
	if password == "" {
		password = os.Getenv("GARMIN_PASSWORD")
	}

	return strings.TrimSpace(username), strings.TrimSpace(password)
}
