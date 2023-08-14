package cmd

import (
	"os"

	"github.com/lindell/connect-iq-sdk-manager-cli/internal/manager"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const defaultManifest = "manifest.xml"

func configureDeviceCmd(cmd *cobra.Command) {
	flags := cmd.PersistentFlags()

	flags.StringSliceP("device", "d", nil, "The device(s) that should be used.")
	flags.StringP("manifest", "m", "", "Path to the manifest file.")
	flags.BoolP("download-all", "", false, "Path to the manifest file.")
}

func getDeviceFilters(cmd *cobra.Command) (manager.DeviceFilters, error) {
	devices, _ := cmd.Flags().GetStringSlice("device")
	manifest, _ := cmd.Flags().GetString("manifest")
	downloadAll, _ := cmd.Flags().GetBool("download-all")

	err := checkExlusivity(map[string]bool{
		"devices":      len(devices) > 0,
		"manifest":     manifest != "",
		"download-all": downloadAll,
	})
	if err != nil {
		return manager.DeviceFilters{}, err
	}

	switch {
	case downloadAll:
		return manager.DeviceFilters{}, nil
	case len(devices) > 0:
		return manager.DeviceFilters{DeviceNames: devices}, nil
	case manifest != "":
		if !fileExists(manifest) {
			return manager.DeviceFilters{}, errors.Errorf("could not find manifest file %s", manifest)
		}
		return manager.DeviceFilters{
			Manifest: manifest,
		}, nil
	}

	// If no filters are defined, but an manifest.xml file exist in the current dir, use it.
	if fileExists(defaultManifest) {
		log.Infof("Found %s file, using it to download devices", defaultManifest)
		return manager.DeviceFilters{
			Manifest: defaultManifest,
		}, nil
	}

	log.Infof("Did not find any %s file, downloading all devices", defaultManifest)
	return manager.DeviceFilters{}, nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
