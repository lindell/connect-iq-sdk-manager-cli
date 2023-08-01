package cmd

import (
	"fmt"
	"runtime"
	"time"

	"github.com/spf13/cobra"
)

// VersionCmd prints the version
func VersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Get the version of connect-iq-manager.",
		Long:  "Get the version of connect-iq-manager.",
		Args:  cobra.NoArgs,
		Run:   version,
	}

	return cmd
}

// Version is the current version (set by main.go)
var Version string

// BuildDate is the time the build was made (set by main.go)
var BuildDate time.Time

// Commit is the commit the build was made on (set by main.go)
var Commit string

func version(_ *cobra.Command, _ []string) {
	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("Release-Date: %s\n", BuildDate.Format("2006-01-02"))
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Arch: %s\n", runtime.GOARCH)
	fmt.Printf("Commit: %s\n", Commit)
}
