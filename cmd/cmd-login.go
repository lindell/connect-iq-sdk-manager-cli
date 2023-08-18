package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// LoginCmd logs the user in to the garmin service
func LoginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login the user to the Garmin.",
		Long: `Login the user to the Garmin.
The credentials can either be set via the --username and --password config,
via GARMIN_USERNAME and GARMIN_PASSWORD environment variables,
or be inputed interactively.`,
		Args: cobra.NoArgs,
		RunE: login,
	}

	cmd.Flags().StringP("username", "", "", "The Garmin username.")
	cmd.Flags().StringP("password", "", "", "The Garmin password.")

	return cmd
}

func login(cmd *cobra.Command, _ []string) error {
	flag := cmd.Flags()

	username, password := getLoginCredentials(flag)

	manager := NewManager()

	ctx := context.Background()

	if username != "" && password != "" {
		log.Debug("Using credentials to simulate oauth login")
		err := manager.LoginWithCredentials(ctx, username, password)
		if err != nil {
			return err
		}
	} else {
		log.Debug("Using Oauth flow to login")
		err := manager.LoginWithOauth(ctx)
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
