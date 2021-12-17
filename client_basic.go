package huwlte

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"golang.org/x/xerrors"
)

const (
	headerRequestVerificationToken    = "__RequestVerificationToken"
	headerRequestVerificationTokenOne = "__RequestVerificationTokenone"
	headerRequestVerificationTokenTwo = "__RequestVerificationTokentwo"
)

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
			time.Sleep(time.Second)
			continue
		}

		return err
	}
}

func (client *Client) Get(ctx context.Context, path string, dst interface{}) error {
	return client.withSessionRetry(ctx, func(ctx context.Context) error {
		return client.get(ctx, path, dst)
	})
}

func (client *Client) get(ctx context.Context, path string, dst interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, client.getURL(path), nil)
	if err != nil {
		return xerrors.Errorf("create request: %w", err)
	}

	if client.session.Cookie != "" {
		req.Header.Set("Cookie", client.session.Cookie)
	}

	if client.session.HasToken() {
		req.Header.Set(headerRequestVerificationToken, client.session.Tokens[0])
	}

	res, err := client.doer.Do(req)
	if err != nil {
		return xerrors.Errorf("do request: %w", err)
	}
	defer res.Body.Close()

	if err := client.proccessResponse(res.Body, dst); err != nil {
		return xerrors.Errorf("process response: %w", err)
	}

	return nil
}

func (client *Client) proccessResponse(r io.Reader, dst interface{}) error {
	content, err := ioutil.ReadAll(r)
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
		if w, ok := dst.(io.Writer); ok {
			if _, err := w.Write(content); err != nil {
				return xerrors.Errorf("write response: %w", err)
			}
			return nil
		}

		if err := envelope.decode(dst); err != nil {
			return xerrors.Errorf("decode response: %w", err)
		}
	}

	return nil
}

func (client *Client) post(
	ctx context.Context,
	path string,
	data interface{},
	refershCSRF bool,
	dst interface{},
) error {
	body := &bytes.Buffer{}

	if data != nil {
		if err := xml.NewEncoder(body).Encode(data); err != nil {
			return xerrors.Errorf("encode request: %w", err)
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, client.getURL(path), body)
	if err != nil {
		return xerrors.Errorf("create request: %w", err)
	}

	if client.session.Cookie != "" {
		req.Header.Set("Cookie", client.session.Cookie)
	}

	var csrfToken string

	if client.session.HasToken() {
		if client.session.HasMultipleTokens() {
			csrfToken = client.session.PopToken()
		} else {
			csrfToken = client.session.Tokens[0]
		}
	}

	if csrfToken != "" {
		req.Header.Set(headerRequestVerificationToken, csrfToken)
	}

	res, err := client.doer.Do(req)
	if err != nil {
		return xerrors.Errorf("do request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return xerrors.Errorf("response status code not 200: %d", res.StatusCode)
	}

	if err := client.proccessResponse(res.Body, dst); err != nil {
		return xerrors.Errorf("process response: %w", err)
	}

	if refershCSRF {
		client.session.ResetTokens()
	}

	if v := res.Header.Get("Set-Cookie"); v != "" {
		client.session.Cookie = v
	}

	if rvt1 := res.Header.Get(headerRequestVerificationTokenOne); rvt1 != "" {
		client.session.AddToken(rvt1)
		if rvt2 := res.Header.Get(headerRequestVerificationTokenTwo); rvt2 != "" {
			client.session.AddToken(rvt2)
		}
	} else if rvt := res.Header.Get(headerRequestVerificationToken); rvt != "" {
		client.session.AddToken(rvt)
	}

	return nil
}
