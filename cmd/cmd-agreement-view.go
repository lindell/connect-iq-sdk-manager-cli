package cmd

import (
	"context"

	"github.com/lindell/connect-iq-sdk-manager-cli/internal/manager"
	"github.com/spf13/cobra"
)

// AgreementViewCmd gives a link to the agreement
func AgreementViewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "view",
		Short: "View the SDK agreement.",
		Long:  "View the the Garmin CONNECT IQ SDK License Agreement and CONNECT IQ Application Developer Agreement.",
		Args:  cobra.NoArgs,
		RunE:  viewAgreement,
	}

	return cmd
}

func viewAgreement(_ *cobra.Command, _ []string) error {
	ctx := context.Background()
	mngr, ctx, err := manager.NewManager(ctx, manager.ManagerConfig{
		SkipAgreementCheck: true,
		SkipLoginRequired:  true,
	})
	if err != nil {
		return err
	}

	return mngr.ViewAgreement(ctx)
}
