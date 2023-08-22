package manager

import (
	"context"

	"github.com/lindell/connect-iq-sdk-manager-cli/internal/client"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/connectiq"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/datetime"
	"github.com/pkg/errors"
)

func (m *Manager) AcceptAgreement(ctx context.Context, hash string) error {
	currentHash, err := client.AgreementHash(ctx)
	if err != nil {
		return errors.WithMessage(err, "could not get agreement hash")
	}

	if hash != "" {
		if hash != currentHash {
			return errors.Errorf("the set agreement hash %q does not match the current one %q", hash, currentHash)
		}
	}

	return connectiq.StoreConfigKeyVals(
		connectiq.ConfigEntity{Key: "agreement-hash", Value: currentHash},
		// This is not a GUI field, but a field used by the CLI to not have to recheck the agreement every request
		connectiq.ConfigEntity{Key: "agreement-hash-verified-at", Value: datetime.Now().String()},
	)
}
