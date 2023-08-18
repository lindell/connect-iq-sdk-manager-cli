package manager

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/client"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/connectiq"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (m *Manager) downloadFonts(ctx context.Context, deviceInfos []client.DeviceInfo) error {
	fontFilenames, err := fontsFromDevices(deviceInfos)
	if err != nil {
		return err
	}
	log.Infof("Installing %d fonts", len(fontFilenames))

	existingFonts, err := installedFonts()
	if err != nil {
		return errors.WithMessage(err, "could not find which fonts are already installed")
	}

	for _, fontFilename := range fontFilenames {
		log := log.WithField("font", fontFilename)

		if existingFonts.Contains(fontFilename) {
			log.Info("Font is already installed")
			continue
		} else if strings.Contains(fontFilename, " ") {
			// This seems to be the same behavior as the GUI since these file are never downloaded
			log.Info("Font is contains spaces and can't be downloaded")
			continue
		}

		if err := m.fetchFont(ctx, log, fontFilename); err != nil {
			return errors.WithMessagef(err, "unable to download font: %q", fontFilename)
		}
	}

	return nil
}

func (m *Manager) fetchFont(ctx context.Context, log log.FieldLogger, font string) error {
	log.Info("Downloading font zip")
	r, err := client.DownloadFont(ctx, font)
	if err != nil {
		return err
	}
	defer r.Close()

	md5hash, err := fetchAndExtract(r, connectiq.FontsPath)
	if err != nil {
		return err
	}

	// The GUI sdk manager saves a file containing the md5 hash of the downloaded (zip) file.
	// It is unclear what it is for, since it does not seem to check this with a HEAD request or similar.
	// But to keep compatibility we will also save this file.
	hashFile := filepath.Join(connectiq.FontsPath, font+".md5")
	if err := os.WriteFile(hashFile, []byte(md5hash), 0644); err != nil { //nolint:gosec
		return errors.WithMessage(err, "could not write md5 hash to file")
	}

	return nil
}

// installedFonts returns a list of all installed fonts
func installedFonts() (mapset.Set[string], error) {
	entries, err := os.ReadDir(connectiq.FontsPath)
	if os.IsNotExist(err) {
		return mapset.NewSet[string](), os.Mkdir(connectiq.FontsPath, 0755)
	} else if err != nil {
		return nil, errors.WithMessage(err, "could not read fonts")
	}

	set := mapset.NewSet[string]()
	for _, e := range entries {
		nameWithoutExt := strings.TrimSuffix(e.Name(), filepath.Ext(e.Name()))
		set.Add(nameWithoutExt)
	}
	return set, nil
}

// fontsFromDevices returns all fonts that are included in a list of devices
func fontsFromDevices(devices []client.DeviceInfo) ([]string, error) {
	fontSet := mapset.NewSet[string]()
	for _, device := range devices {
		newFonts, err := fontsFromDevice(device)
		if err != nil {
			return nil, err
		}
		fontSet.Append(newFonts...)
	}
	return fontSet.ToSlice(), nil
}

func fontsFromDevice(device client.DeviceInfo) ([]string, error) {
	simPath := filepath.Join(connectiq.DevicesPath, device.Name, "simulator.json")

	bb, err := os.ReadFile(simPath)
	if err != nil {
		return nil, errors.WithMessage(err, "could not read simulator file")
	}

	var simData connectiq.SimulatorFile
	if err := json.Unmarshal(bb, &simData); err != nil {
		return nil, errors.WithMessage(err, "could not decode simulator file")
	}

	fontSet := mapset.NewSet[string]()
	for _, font := range simData.Fonts {
		for _, font := range font.Fonts {
			fontSet.Add(font.Filename)
		}
	}
	return fontSet.ToSlice(), nil
}
