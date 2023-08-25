package client

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

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

	err := unexpectedStatusCodeError{code: resp.StatusCode}

	var errorResp errorResponse
	if err := json.NewDecoder(resp.Body).Decode(&errorResp); err != nil || errorResp.Error == "" {
		return err
	}

	return errors.WithMessagef(err, "%s: %s", errorResp.Error, errorResp.Message)
}

func urlMustParse(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return u
}

func cloneURL(u *url.URL) *url.URL {
	u2 := *u
	return &u2
}
