package oauth

import (
	"bytes"
	"context"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"

	log "github.com/sirupsen/logrus"
)

//go:embed done.html
var doneHTML []byte

const (
	urlDone  = "/done"
	urlLogin = "/login"
)

// GetToken asks the user to authenticate with the Oauth flow.
// It will spin up a server to listen for the callback that it will close once the ticket is received
// If successful, it returns the token, as well as the service-url used, that is needed when exchanging the ticket
func GetToken(ctx context.Context) (token string, serviceURL string, err error) {
	// Start listening on random port
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return "", "", err
	}
	defer l.Close()

	port := l.Addr().(*net.TCPAddr).Port
	host := fmt.Sprintf("http://localhost:%d", port)

	log.Debugf("Listening on %s\n", host)
	fmt.Printf("Please login at: %s\n", host+urlLogin)

	ticketCh := make(chan string, 1)

	mux := http.NewServeMux()
	mux.HandleFunc(urlDone, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ticket := r.URL.Query().Get("ticket")
		log.Debug("Received ticket")

		buf := bytes.NewBuffer(doneHTML)
		_, _ = io.Copy(rw, buf)

		ticketCh <- ticket
	}))
	mux.Handle(urlLogin, http.RedirectHandler(loginURL(host+urlDone), http.StatusSeeOther))

	server := http.Server{
		ReadTimeout:       time.Second * 10,
		ReadHeaderTimeout: time.Second * 10,
		WriteTimeout:      time.Second * 10,
		IdleTimeout:       time.Second * 10,
		Handler:           mux,
	}

	go func() {
		if err := server.Serve(l); err != http.ErrServerClosed {
			log.Error(err)
		}
	}()

	ticket, ok := <-ticketCh
	if !ok {
		return "", "", errors.New("unexpected close while listening for ticket")
	}

	return ticket, host + urlDone, server.Shutdown(ctx)
}

func loginURL(localURL string) string {
	u, err := url.Parse("https://sso.garmin.com/sso/signin?gauthHost=https%3A%2F%2Fsso.garmin.com%2Fsso&locale=en&id=gauth-widget&cssUrl=https%3A%2F%2Fdeveloper.garmin.com%2Fdownloads%2Fconnect-iq%2Fsdk-manager-login.css&clientId=ConnectIqSdkManager&rememberMeShown=false&rememberMeChecked=false&createAccountShown=true&openCreateAccount=false&displayNameShown=false&consumeServiceTicket=true&initialFocus=true&embedWidget=true&generateExtraServiceTicket=false&generateTwoExtraServiceTickets=false&generateNoServiceTicket=false&globalOptInShown=false&globalOptInChecked=false&mobile=false&connectLegalTerms=false&showTermsOfUse=false&showPrivacyPolicy=false&showConnectLegalAge=false&locationPromptShown=false&showPassword=true&useCustomHeader=false&mfaRequired=false&performMFACheck=false&rememberMyBrowserShown=false&rememberMyBrowserChecked=false")
	if err != nil {
		panic(err)
	}

	q := u.Query()
	q.Add("service", localURL)
	q.Add("source", localURL)
	q.Add("redirectAfterAccountLoginUrl", localURL)
	q.Add("redirectAfterAccountCreationUrl", localURL)
	u.RawQuery = q.Encode()

	return u.String()
}
