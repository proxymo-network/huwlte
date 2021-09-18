package huwlte

import (
	"context"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/xerrors"
)

type session struct {
	XMLName xml.Name `xml:"response"`
	Cookie  string   `xml:"SesInfo"`
	Token   string   `xml:"TokInfo"`
}

func (s *session) empty() bool {
	return s.Cookie == "" || s.Token == ""
}

func (client *Client) getURL(path string) string {
	return strings.TrimSuffix(client.baseURL, "/") + "/" + strings.TrimPrefix(path, "/")
}

func (client *Client) getSession(ctx context.Context) error {
	if err := client.get(ctx, "/api/webserver/SesTokInfo", &client.session); err != nil {
		return xerrors.Errorf("get session: %w", err)
	}
	return nil
}

func (client *Client) withSessionRetry(ctx context.Context, f func(ctx context.Context) error) error {
	if client.session.empty() {
		if err := client.getSession(ctx); err != nil {
			return xerrors.Errorf("get session: %w", err)
		}
	}

	for {
		err := f(ctx)

		var modemErr *Error

		if errors.As(err, &modemErr) {
			if modemErr.Code != ErrorCodeCSRF {
				return err
			}

			if err := client.getSession(ctx); err != nil {
				return xerrors.Errorf("get session: %w", err)
			}

			continue
		} else {
			return err
		}
	}
}

func (client *Client) get(ctx context.Context, path string, dst interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, client.getURL(path), nil)
	if err != nil {
		return xerrors.Errorf("create request: %w", err)
	}

	if client.session.Cookie != "" {
		req.Header.Set("Cookie", client.session.Cookie)
	}

	if client.session.Token != "" {
		req.Header.Set("__RequestVerificationToken", client.session.Token)
	}

	res, err := client.doer.Do(req)
	if err != nil {
		return xerrors.Errorf("do request: %w", err)
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return xerrors.Errorf("read response: %w", err)
	}

	envelope, err := parseResponseEnvelope(content)
	if err != nil {
		return xerrors.Errorf("parse response envelope: %w", err)
	}

	if err := envelope.toErr(); err != nil {
		return err
	}

	if dst != nil {
		if err := envelope.decode(dst); err != nil {
			return xerrors.Errorf("decode response: %w", err)
		}
	}

	return nil
}
