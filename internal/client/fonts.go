package client

import (
	"context"
	"io"
	"net/http"
)

var fontDownloadURL = urlMustParse("https://api.gcs.garmin.com/ciq-product-onboarding/fonts/font")

func DownloadFont(ctx context.Context, fontFilename string) (io.ReadCloser, error) {
	u := cloneURL(fontDownloadURL)
	q := u.Query()
	q.Set("fontName", fontFilename)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
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
