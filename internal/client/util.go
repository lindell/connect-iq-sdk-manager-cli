package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

func mustParseURL(rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return u
}

type errorResponse struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

// expectStatusCode assumes a status code, and will otherwise try to decode the error message
func expectStatusCode(resp *http.Response, status ...int) error {
	for _, s := range status {
		if resp.StatusCode == s {
			return nil
		}
	}

	var errorResp errorResponse
	if err := json.NewDecoder(resp.Body).Decode(&errorResp); err != nil || errorResp.Error == "" {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return errors.Errorf("%s: %s", errorResp.Error, errorResp.Message)
}
