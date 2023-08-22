package manager

import (
	"context"
	"fmt"

	"github.com/lindell/connect-iq-sdk-manager-cli/internal/client"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/connectiq"
)

func (m *Manager) ViewAgreement(ctx context.Context) error {
	fmt.Println("Agreement:", connectiq.AgreementURL)

	hash, err := client.AgreementHash(ctx)
	if err != nil {
		fmt.Printf("Could not fetch current hash: %s\n", err)
	} else {
		fmt.Printf("Current Hash: %s\n", hash)
	}

	return nil
}
