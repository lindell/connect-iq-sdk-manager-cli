package cmd

import (
	"context"
	"errors"
	"regexp"
	"strings"

	"github.com/lindell/connect-iq-sdk-manager-cli/internal/manager"
	"github.com/spf13/cobra"
)

// AgreementAcceptCmd gives a link to the agreement
func AgreementAcceptCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accept",
		Short: "Accept the SDK agreement.",
		Long: `Accept the the Garmin CONNECT IQ SDK License Agreement and CONNECT IQ Application Developer Agreement.
		
You can either accept the latest, just read agreement. Or accept a previously read agreement, if used in for example CI/CD.`,
		Args: cobra.NoArgs,
		RunE: acceptAgreement,
	}

	cmd.Flags().StringP("agreement-hash", "H", "", "The hash of a previously read agreement.")

	return cmd
}

var md5regexp = regexp.MustCompile("^[a-fA-F0-9]{32}$")

func acceptAgreement(cmd *cobra.Command, _ []string) error {
	ctx := context.Background()
	mngr, ctx, err := manager.NewManager(ctx, manager.ManagerConfig{
		SkipAgreementCheck: true,
		SkipLoginRequired:  true,
	})
	if err != nil {
		return err
	}

	hash, _ := cmd.Flags().GetString("agreement-hash")
	if hash != "" {
		if !md5regexp.MatchString(hash) {
			return errors.New("hash is not an md5 hash")
		}
	}

	return mngr.AcceptAgreement(ctx, strings.ToUpper(hash))
}
