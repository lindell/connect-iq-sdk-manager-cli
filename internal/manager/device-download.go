package manager

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/lindell/connect-iq-sdk-manager-cli/internal/client"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/connectiq"
	log "github.com/sirupsen/logrus"
)

type DownloadConfig struct {
	DeviceFilters DeviceFilters
	IncludeFonts  bool
}

func (m *Manager) Download(ctx context.Context, config DownloadConfig) error {
	deviceInfos, err := getDeviceInfo(ctx)
	if err != nil {
		return err
	}

	deviceInfos, err = filterDevices(deviceInfos, config.DeviceFilters)
	if err != nil {
		return err
	}

	log.Infof("Downloading %d devices", len(deviceInfos))

	for _, device := range deviceInfos {
		log := log.WithField("device", device.Name)
		if err := m.fetchDevice(ctx, log, device); err != nil {
			return err
		}
	}

	if config.IncludeFonts {
		if err := m.downloadFonts(ctx, deviceInfos); err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager) fetchDevice(ctx context.Context, log log.FieldLogger, device client.DeviceInfo) error {
	deviceDir := path.Join(connectiq.DevicesPath, device.Name)

	if _, err := os.Stat(deviceDir); !os.IsNotExist(err) {
		log.Info("Device folder already exist")
		return nil
	}

	log.Info("Downloading device zip")
	r, err := client.DownloadDevice(ctx, device)
	if err != nil {
		return err
	}
	defer r.Close()

	hash, err := fetchAndExtract(r, deviceDir)
	if err := connectiq.StoreConfigKeyVal(fmt.Sprintf("%s-hash", device.Name), hash); err != nil {
		log.Errorf("Could not store hash: %s", err)
	}
	return err
}
