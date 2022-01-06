package huwlte

import (
	"context"
	"net/http"

	"golang.org/x/xerrors"
)

type ClientOpt func(*Client)

// WithDoer sets the HTTP client used to make requests.
func WithDoer(doer *http.Client) ClientOpt {
	return func(c *Client) {
		c.doer = doer
	}
}

// WithStorage sets the storage used to store the modem session.
func WithStorage(name string, storage SessionStorage) ClientOpt {
	return func(client *Client) {
		client.sessionStorage = storage
		client.sessionName = name
	}
}

// WithAdminAuth sets the credentials used to authenticate, when need.
func WithAdminAuth(username, password string) ClientOpt {
	return func(c *Client) {
		c.adminLogin = username
		c.adminPass = password
	}
}

// Clients it's XML API wrapper.
type Client struct {
	baseURL string
	doer    *http.Client
	session Session

	adminLogin string
	adminPass  string

	sessionName    string
	sessionStorage SessionStorage

	Device     *ClientDevice
	User       *ClientUser
	Monitoring *ClientMonitoring
	Dialup     *ClientDialup
	Net        *ClientNet
	SMS        *ClientSMS
}

// New creates a new Client instance.
func New(baseURL string, opts ...ClientOpt) *Client {
	c := &Client{
		baseURL: baseURL,
		doer:    http.DefaultClient,
	}
	for _, opt := range opts {
		opt(c)
	}

	c.Device = &ClientDevice{c}
	c.User = &ClientUser{c}
	c.Monitoring = &ClientMonitoring{c}
	c.Dialup = &ClientDialup{c}
	c.Net = &ClientNet{c}
	c.SMS = &ClientSMS{c}

	return c
}

func (client *Client) SaveSession(ctx context.Context) error {
	if client.sessionStorage == nil {
		return nil
	}

	return client.sessionStorage.Save(ctx, client.sessionName, &client.session)
}

func (client *Client) ResetSession(ctx context.Context) error {
	client.session = Session{}

	if client.sessionStorage != nil {
		return client.sessionStorage.Remove(ctx, client.sessionName)
	}

	return nil
}

func (client *Client) LoadSession(ctx context.Context) error {
	if client.sessionStorage == nil {
		return nil
	}

	session, err := client.sessionStorage.Load(ctx, client.sessionName)
	if err != nil {
		return xerrors.Errorf("load session: %w", err)
	}

	client.session = *session

	return nil
}

// getDoer returns the HTTP client used to make requests.
// If it's not set, it returns the default HTTP client.
func (client *Client) getDoer() *http.Client {
	if client.doer != nil {
		return client.doer
	}

	return http.DefaultClient
}
