package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

const sdksURL = "https://developer.garmin.com/downloads/connect-iq/sdks/sdks.json"

type SDK struct {
	Version string `json:"version"`
	Title   string `json:"title"`
	Release string `json:"release"`
	Mac     string `json:"mac"`
	Windows string `json:"windows"`
	Linux   string `json:"linux"`
}

func GetSDKs(ctx context.Context) ([]SDK, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, sdksURL, nil)
	if err != nil {
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

	var sdks []SDK
	if err := json.NewDecoder(resp.Body).Decode(&sdks); err != nil {
		return nil, errors.WithMessage(err, "could not decode sdk information")
	}

	return sdks, nil
}
