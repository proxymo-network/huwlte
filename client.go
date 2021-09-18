package huwlte

import (
	"net/http"
)

type ClientOpt func(*Client)

// WithDoer sets the HTTP client used to make requests.
func WithDoer(doer *http.Client) ClientOpt {
	return func(c *Client) {
		c.doer = doer
	}
}

// Clients it's XML API wrapper.
type Client struct {
	baseURL string
	doer    *http.Client
	session session

	Device *ClientDevice
}

// NewClient creates a new Client instance.
func NewClient(baseURL string, opts ...ClientOpt) *Client {
	c := &Client{
		baseURL: baseURL,
		doer:    http.DefaultClient,
	}
	for _, opt := range opts {
		opt(c)
	}

	c.Device = &ClientDevice{c}

	return c
}

// getDoer returns the HTTP client used to make requests.
// If it's not set, it returns the default HTTP client.
func (client *Client) getDoer() *http.Client {
	if client.doer != nil {
		return client.doer
	}

	return http.DefaultClient
}
