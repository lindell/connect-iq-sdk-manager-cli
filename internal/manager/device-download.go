package manager

import (
	"context"
	"os"
	"path"

	"github.com/lindell/connect-iq-sdk-manager-cli/internal/client"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/connectiq"
	log "github.com/sirupsen/logrus"
)

type DownloadConfig struct {
	DeviceFilters DeviceFilters
}

func (m *Manager) Download(ctx context.Context, config DownloadConfig) error {
	var err error
	if ctx, err = m.setTokenToCtx(ctx); err != nil {
		return err
	}

	deviceInfos, err := client.GetDeviceInfo(ctx)
	if err != nil {
		return err
	}

	deviceInfos, err = filterDevices(deviceInfos, config.DeviceFilters)
	if err != nil {
		return err
	}

	log.Infof("Downloading %d devices.", len(deviceInfos))

	for _, device := range deviceInfos {
		log := log.WithField("device", device.Name)
		if err := m.fetchDevice(ctx, log, device); err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager) fetchDevice(ctx context.Context, log log.FieldLogger, device client.DeviceInfo) error {
	rootFolder, err := connectiq.RootGarminFolder()
	if err != nil {
		return err
	}
	deviceDir := path.Join(rootFolder, "Devices", device.Name)

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

	// Save the zip to a temporary file
	return fetchAndExtract(r, deviceDir)
}
