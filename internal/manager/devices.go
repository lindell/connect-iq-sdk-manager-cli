package manager

import (
	"context"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/lindell/connect-iq-sdk-manager-cli/internal/client"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/connectiq"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/datetime"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// DeviceFilters describes which devices should be filtered
type DeviceFilters struct {
	Manifest    string   // Path to the manifest file where devices are specified
	DeviceNames []string // The name of the device to download
}

func filterDevices(deviceInfos []client.DeviceInfo, deviceFilters DeviceFilters) ([]client.DeviceInfo, error) {
	switch {
	case deviceFilters.Manifest != "":
		deviceNames, err := manifestDeviceNames(deviceFilters.Manifest)
		if err != nil {
			return nil, err
		}
		return filterDeviceNames(deviceInfos, deviceNames)

	case len(deviceFilters.DeviceNames) > 0:
		return filterDeviceNames(deviceInfos, deviceFilters.DeviceNames)
	}

	return deviceInfos, nil
}

func filterDeviceNames(devices []client.DeviceInfo, deviceNames []string) ([]client.DeviceInfo, error) {
	// Create a map for fast lookup of device names
	deviceNameMap := map[string]client.DeviceInfo{}
	for _, device := range devices {
		deviceNameMap[device.Name] = device
	}

	newDevices := make([]client.DeviceInfo, 0, len(deviceNames))
	for _, deviceName := range deviceNames {
		if device, exist := deviceNameMap[deviceName]; exist {
			newDevices = append(newDevices, device)
		} else {
			return nil, errors.Errorf("device %s was defined to be downloaded, but could not be found", deviceName)
		}
	}

	return newDevices, nil
}

func getDeviceInfo(ctx context.Context) ([]client.DeviceInfo, error) {
	deviceInfos, err := client.GetDeviceInfo(ctx)
	if err != nil {
		return nil, err
	}
	if err := storeDeviceFirstSeen(deviceInfos); err != nil {
		log.Errorf("Could not store first seen: %s", err)
	}
	return deviceInfos, nil
}

// storeDeviceFirstSeen stores the first time a device was seen in the config file
// This is not used by the CLI, but is stored by the GUI
func storeDeviceFirstSeen(devices []client.DeviceInfo) error {
	firstSeenKeys := make([]string, len(devices))
	for i, d := range devices {
		firstSeenKeys[i] = fmt.Sprintf("%s-first-seen", d.Name)
	}

	alreadySeen := connectiq.LoadConfigVals(firstSeenKeys...)

	notYetSeens := []connectiq.ConfigEntity{}
	for i := range alreadySeen {
		if alreadySeen[i] == "" {
			notYetSeens = append(notYetSeens, connectiq.ConfigEntity{
				Key:   firstSeenKeys[i],
				Value: datetime.Now().String(),
			})
		}
	}

	return connectiq.StoreConfigKeyVals(notYetSeens...)
}

// manifest contains data from the manifest.xml file
type manifest struct {
	Application struct {
		Products struct {
			Text    string `xml:",chardata"`
			Product []struct {
				ID string `xml:"id,attr"`
			} `xml:"product"`
		} `xml:"products"`
	} `xml:"application"`
}

func manifestDeviceNames(manifestPath string) ([]string, error) {
	bb, err := os.ReadFile(manifestPath)
	if err != nil {
		return nil, errors.WithMessage(err, "could not read manifest file")
	}

	var manifest manifest
	if err := xml.Unmarshal(bb, &manifest); err != nil {
		return nil, errors.WithMessage(err, "could not unmarshal manifest file")
	}
	products := manifest.Application.Products.Product

	deviceNames := make([]string, len(products))
	for i, product := range products {
		deviceNames[i] = product.ID
	}

	return deviceNames, nil
}
