package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/lindell/connect-iq-sdk-manager-cli/internal/connectiq"
	"github.com/pkg/errors"
)

type DeviceInfo struct {
	DeviceUUID            string             `json:"deviceUuid"`
	PartNumber            string             `json:"partNumber"`
	Name                  string             `json:"name"`
	ProductInfoFileExists bool               `json:"productInfoFileExists"`
	CiqInfoFileExists     bool               `json:"ciqInfoFileExists"`
	Upcoming              bool               `json:"upcoming"`
	ProductInfoHash       string             `json:"productInfoHash"`
	CiqInfoHash           string             `json:"ciqInfoHash"`
	Group                 string             `json:"group"`
	DisplayName           string             `json:"displayName"`
	LastUpdateTime        connectiq.DateTime `json:"lastUpdateTime"`
	Hidden                bool               `json:"hidden"`
	Faceit2Capable        bool               `json:"faceit2Capable"`
}

var devicesURL = "https://api.gcs.garmin.com/ciq-product-onboarding/devices?sdkManagerVersion=1.0.5"
var deviceDownloadURL = "https://api.gcs.garmin.com/ciq-product-onboarding/devices/%s/ciqInfo" // %s is partNumber

func GetDeviceInfo(ctx context.Context) ([]DeviceInfo, error) {
	req, err := http.NewRequest(http.MethodGet, devicesURL, nil)
	if err != nil {
		return nil, err
	}
	if err := addReqCredentials(ctx, req); err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code, status code: %d", resp.StatusCode)
	}

	var deviceInfo []DeviceInfo
	if err := json.NewDecoder(resp.Body).Decode(&deviceInfo); err != nil {
		return nil, errors.WithMessage(err, "could not decode device information")
	}

	return removeDuplicateDevices(deviceInfo)
}

// removeDuplicateDevices removes any dublicate devices
// For some reason, the API returns multiple devices with the same name, this function removes those
func removeDuplicateDevices(devices []DeviceInfo) ([]DeviceInfo, error) {
	deviceMap := map[string]DeviceInfo{}
	for _, device := range devices {
		if !device.CiqInfoFileExists {
			continue
		}

		// If the device already exist, choose the one that was updated last
		if existingDevice, exist := deviceMap[device.Name]; !exist {
			deviceMap[device.Name] = device
		} else {
			existingTime := existingDevice.LastUpdateTime.Time()
			newTime := device.LastUpdateTime.Time()

			if newTime.After(existingTime) {
				deviceMap[device.Name] = device
			}
		}
	}

	newDevices := make([]DeviceInfo, 0, len(deviceMap))
	for _, device := range deviceMap {
		newDevices = append(newDevices, device)
	}
	return newDevices, nil
}

func DownloadDevice(ctx context.Context, device DeviceInfo) (io.ReadCloser, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(deviceDownloadURL, url.PathEscape(device.PartNumber)), nil)
	if err != nil {
		return nil, err
	}
	if err := addReqCredentials(ctx, req); err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if err := expectStatusCode(resp, http.StatusOK); err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func addReqCredentials(ctx context.Context, req *http.Request) error {
	token, ok := connectiq.GetContextToken(ctx)
	if !ok {
		return errors.New("could not find credentials in token")
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	return nil
}
