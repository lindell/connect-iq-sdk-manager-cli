package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/connectiq"
)

type tokenResponse struct {
	AccessToken           string `json:"access_token"`
	TokenType             string `json:"token_type"`
	ExpiresIn             int    `json:"expires_in"`
	Scope                 string `json:"scope"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn string `json:"refresh_token_expires_in"`
	CustomerID            string `json:"customerId"`
}

const loginURL = `https://sso.garmin.com/sso/signin?service=https%3A%2F%2Fsso.garmin.com%2Fsso%2Fembed&clientId=ConnectIqSdkManager`
const exchangeURL = "https://services.garmin.com/api/oauth/token"

func setHeaders(header http.Header) {
	header.Set("origin", "https://sso.garmin.com")
}

var ticketRegexp = regexp.MustCompile("ticket=([A-Za-z0-9-]+)")

func Login(ctx context.Context, username, password string) (connectiq.Token, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return connectiq.Token{}, err
	}

	client := &http.Client{
		Jar: jar,
	}

	ticket, err := login(ctx, client, username, password)
	if err != nil {
		return connectiq.Token{}, err
	}

	token, err := exchangeTicket(ctx, client, ticket)
	if err != nil {
		return connectiq.Token{}, err
	}

	return connectiq.Token{
		AccessToken: token.AccessToken,
		ExpiresAt:   time.Now().Add(time.Duration(token.ExpiresIn) * time.Second),
	}, nil
}

func login(ctx context.Context, client *http.Client, username, password string) (ticket string, err error) {
	data, err := formValuesFromInitialRequest(ctx, client)
	if err != nil {
		return "", err
	}

	data.Set("username", username)
	data.Set("password", password)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, loginURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	setHeaders(req.Header)
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("could not login, status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	results := ticketRegexp.FindStringSubmatch(string(body))
	if results == nil {
		return "", errors.New("could not find ticket") // TODO: Return actual error
	}
	ticket = results[1]

	return ticket, nil
}

func exchangeTicket(ctx context.Context, client *http.Client, ticket string) (tokenResponse, error) {
	data := url.Values{}
	data.Set("grant_type", "service_ticket")
	data.Set("client_id", "CIQ_SDK_MANAGER")
	data.Set("service_ticket", ticket)
	data.Set("service_url", "https://sso.garmin.com/sso/embed")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, exchangeURL, strings.NewReader(data.Encode()))
	if err != nil {
		return tokenResponse{}, err
	}

	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return tokenResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return tokenResponse{}, fmt.Errorf("could not exchange ticket for token, status code: %d", resp.StatusCode)
	}

	var tokenResp tokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return tokenResponse{}, err
	}

	return tokenResp, nil
}

// formValuesFromInitialRequest gets (hidden) input fields that can be used in the login request
func formValuesFromInitialRequest(ctx context.Context, client *http.Client) (url.Values, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, loginURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not load initial login page, status code: %d", resp.StatusCode)
	}

	// Copy inputs (importantly CSRF token) into next request
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	data := url.Values{}
	doc.Find("#login-form input").Each(func(i int, s *goquery.Selection) {
		if name, exist := s.Attr("name"); exist {
			data.Add(name, s.AttrOr("value", ""))
		}
	})

	return data, nil
}
