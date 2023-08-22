package client

import (
	"context"
	"crypto/md5"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/connectiq"
	"github.com/pkg/errors"
)

// AgreementHash returns the MD5 hash of the agreement the same way the GUI does it
func AgreementHash(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, connectiq.AgreementURL, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code, status code: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	disc, err := doc.Find(".disclaimer").Html()
	if err != nil {
		return "", errors.WithMessage(err, "could not get disclaimer text")
	}

	hash := md5.New() //nolint:gosec
	hash.Write([]byte(disc))
	md5hex := fmt.Sprintf("%x", hash.Sum(nil))

	return strings.ToUpper(md5hex), nil
}
