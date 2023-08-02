package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"golang.org/x/term"
)

// LoginCmd logs the user in to the garmin service
func LoginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login the user to the Garmin.",
		Long:  "Login the user to the Garmin.",
		Args:  cobra.NoArgs,
		RunE:  login,
	}

	cmd.Flags().StringP("username", "", "", "The Garmin username.")
	cmd.Flags().StringP("password", "", "", "The Garmin password.")

	return cmd
}

func login(cmd *cobra.Command, _ []string) error {
	flag := cmd.Flags()

	username, password, err := getLoginCredentials(flag)
	if err != nil {
		return errors.WithMessage(err, "could not get credentials")
	}

	if username == "" || password == "" {
		return errors.New("username and password has to be set")
	}

	manager, err := NewManager()
	if err != nil {
		return err
	}

	ctx := context.Background()

	err = manager.Login(ctx, username, password)
	if err != nil {
		return err
	}

	fmt.Println("Successfully logged in")

	return nil
}

func getLoginCredentials(flag *pflag.FlagSet) (username, password string, err error) {
	username, _ = flag.GetString("username")
	password, _ = flag.GetString("password")

	if username == "" {
		fmt.Print("Enter Username: ")
		reader := bufio.NewReader(os.Stdin)
		username, err = reader.ReadString('\n')
		if err != nil {
			return "", "", err
		}
	}

	if password == "" {
		fmt.Print("Enter Password: ")
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return "", "", err
		}
		fmt.Print("\n")
		password = string(bytePassword)
	}

	return strings.TrimSpace(username), strings.TrimSpace(password), nil
}
