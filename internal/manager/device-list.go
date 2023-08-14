package manager

import (
	"context"
	"os"

	"github.com/lindell/connect-iq-sdk-manager-cli/internal/client"
	"github.com/olekukonko/tablewriter"
)

type DeviceListConfig struct {
	DeviceFilters DeviceFilters
}

func (m *Manager) ListDevices(ctx context.Context, config DeviceListConfig) error {
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

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NAME", "GROUP"})
	table.SetAutoWrapText(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)

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
