package cmd

import (
	"context"

	"github.com/Masterminds/semver/v3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// SDKListCmd list devices
func SDKListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list [version]",
		Short: "List SDKs.",
		Long: `List SDKs.

To only list certain versions. The version argument can be used with a semver-range.
For example: ^6.2.0 or >=4.0.0`,
		Args: cobra.MaximumNArgs(1),
		RunE: listSdks,
	}

	return cmd
}

func listSdks(_ *cobra.Command, args []string) (err error) {
	ctx := context.Background()

	mngr := NewManager()

	semverConstraint, _ := semver.NewConstraint("*")
	if len(args) > 0 {
		semverConstraint, err = semver.NewConstraint(args[0])
		if err != nil {
			return errors.WithMessage(err, "could not parse version constraint")
		}
	}

	return mngr.ListSDKs(ctx, semverConstraint)
}
