package client

import (
	"fmt"
	"net/http"
)

type unexpectedStatusCodeError struct {
	code int
}

func (e unexpectedStatusCodeError) Error() string {
	return fmt.Sprintf("unexpected status code: %d", e.code)
}

func (e unexpectedStatusCodeError) Temporary() bool {
	switch e.code {
	case http.StatusRequestTimeout,
		http.StatusInternalServerError,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout:
		return true
	}
	return false
}

func (e unexpectedStatusCodeError) NotFound() bool {
	return e.code == http.StatusNotFound
}
