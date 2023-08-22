package manager

import (
	"context"

	"github.com/lindell/connect-iq-sdk-manager-cli/internal/client"
)

type DeviceListConfig struct {
	DeviceFilters DeviceFilters
}

func (m *Manager) ListDevices(ctx context.Context, config DeviceListConfig) error {
	deviceInfos, err := getDeviceInfo(ctx)
	if err != nil {
		return err
	}

	deviceInfos, err = filterDevices(deviceInfos, config.DeviceFilters)
	if err != nil {
		return err
	}

	table := createTable()
	table.SetHeader([]string{"NAME", "GROUP"})
	for _, d := range deviceInfos {
		table.Append(deviceToRow(d))
	}
	table.Render()

	return nil
}

func deviceToRow(device client.DeviceInfo) []string {
	return []string{
		device.Name,
		device.Group,
	}
}
