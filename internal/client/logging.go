package client

import (
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// loggingRoundTripper logs a request-response
type loggingRoundTripper struct {
	Next http.RoundTripper
}

// RoundTrip logs a request-response
func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	req, _ := httputil.DumpRequestOut(r, true)

	var roundTripper http.RoundTripper
	if l.Next != nil {
		roundTripper = l.Next
	} else {
		roundTripper = http.DefaultTransport
	}

	start := time.Now()
	resp, err := roundTripper.RoundTrip(r)
	took := time.Since(start)

	var res []byte
	if resp != nil {
		getBody := true

		cdHeader := resp.Header.Get("Content-Disposition")
		if strings.HasPrefix(cdHeader, "attachment") {
			getBody = false
		}

		res, _ = httputil.DumpResponse(resp, getBody)
	}

	logger := log.WithFields(log.Fields{
		"host":     r.Host,
		"took":     took,
		"request":  string(req),
		"response": string(res),
	})
	logger.Trace("http request")

	return resp, err
}
