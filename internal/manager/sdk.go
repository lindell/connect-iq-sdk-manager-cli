package manager

import (
	"sort"

	"github.com/Masterminds/semver/v3"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/client"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type sdkWithVersion struct {
	client.SDK
	version *semver.Version
}

func filterAndSortSDKs(sdks []client.SDK, constraints *semver.Constraints) []client.SDK {
	sdksWithVersions := make([]sdkWithVersion, len(sdks))
	for i, s := range sdks {
		v, err := semver.NewVersion(s.Version)
		if err != nil {
			log.Error(errors.WithMessage(err, "could not parse sdk version"))
			continue
		}

		sdksWithVersions[i] = sdkWithVersion{
			SDK:     s,
			version: v,
		}
	}

	sdksWithVersions = filterSDKs(sdksWithVersions, constraints)
	sortSDKs(sdksWithVersions)

	newSDKs := make([]client.SDK, len(sdksWithVersions))
	for i, s := range sdksWithVersions {
		newSDKs[i] = s.SDK
	}
	return newSDKs
}

func filterSDKs(sdks []sdkWithVersion, constraints *semver.Constraints) []sdkWithVersion {
	filteredSDKs := []sdkWithVersion{}
	for _, s := range sdks {
		if constraints.Check(s.version) {
			filteredSDKs = append(filteredSDKs, s)
		}
	}
	return filteredSDKs
}

func sortSDKs(sdks []sdkWithVersion) {
	sort.Slice(sdks, func(i, j int) bool {
		return sdks[i].version.GreaterThan(sdks[j].version)
	})
}
