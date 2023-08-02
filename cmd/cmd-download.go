package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

// DownloadCmd downloads devices
func DownloadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download",
		Short: "Download resources.",
		Long:  "Download resources.",
		Args:  cobra.NoArgs,
		RunE:  download,
	}

	return cmd
}

func download(_ *cobra.Command, _ []string) error {
	ctx := context.Background()

	manager, err := NewManager()
	if err != nil {
		return err
	}

	return manager.Download(ctx)
}
