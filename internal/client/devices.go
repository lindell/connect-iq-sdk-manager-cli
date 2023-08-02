package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/lindell/connect-iq-manager/internal/connectiq"
	"github.com/pkg/errors"
)

type DeviceInfo struct {
	DeviceUUID            string `json:"deviceUuid"`
	PartNumber            string `json:"partNumber"`
	Name                  string `json:"name"`
	ProductInfoFileExists bool   `json:"productInfoFileExists"`
	CiqInfoFileExists     bool   `json:"ciqInfoFileExists"`
	Upcoming              bool   `json:"upcoming"`
	ProductInfoHash       string `json:"productInfoHash"`
	CiqInfoHash           string `json:"ciqInfoHash"`
	Group                 string `json:"group"`
	DisplayName           string `json:"displayName"`
	LastUpdateTime        string `json:"lastUpdateTime"`
	Hidden                bool   `json:"hidden"`
	Faceit2Capable        bool   `json:"faceit2Capable"`
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

	resp, err := http.DefaultClient.Do(req)
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

	return deviceInfo, nil
}

func DownloadDevice(ctx context.Context, device DeviceInfo) (io.ReadCloser, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(deviceDownloadURL, url.PathEscape(device.PartNumber)), nil)
	if err != nil {
		return nil, err
	}
	if err := addReqCredentials(ctx, req); err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
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
