package huwlte

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"errors"

	"golang.org/x/xerrors"
)

type ClientUser struct {
	*Client
}

type UserStateType int

const (
	UserStateLoggedIn  UserStateType = 0
	UserStateLoggedOut UserStateType = -1
	UserStateRepeat    UserStateType = -2
)

func (typ UserStateType) String() string {
	switch typ {
	case UserStateLoggedIn:
		return "logged_in"
	case UserStateLoggedOut:
		return "logged_out"
	case UserStateRepeat:
		return "repeat"
	default:
		return "unknown"
	}
}

type UserPasswordType int

const (
	UserPasswordTypeBase64 UserPasswordType = 0
	UserPasswordTypeSHA256 UserPasswordType = 4
)

func (typ UserPasswordType) String() string {
	switch typ {
	case UserPasswordTypeBase64:
		return "base64"
	case UserPasswordTypeSHA256:
		return "sha256"
	default:
		return "unknown"
	}
}

type UserStateLogin struct {
	XMLName            xml.Name         `xml:"response" human:"-"`
	State              UserStateType    `xml:"State"`
	Username           string           `xml:"Username"`
	PasswordType       UserPasswordType `xml:"password_type"`
	ExternPasswordType int              `xml:"extern_password_type"`
	Firstlogin         int              `xml:"firstlogin"`
}

func (auth *ClientUser) StateLogin(ctx context.Context) (*UserStateLogin, error) {
	var result UserStateLogin

	if err := auth.withSessionRetry(ctx, func(ctx context.Context) error {
		return auth.get(ctx, "/api/user/state-login", &result)
	}); err != nil {
		return nil, err
	}

	return &result, nil
}

func (auth *ClientUser) Login(ctx context.Context, username string, password string, relogin bool) error {
	if username == "" {
		username = "admin"
	}

	state, err := auth.StateLogin(ctx)

	var modemErr *Error
	if errors.As(err, &modemErr) && modemErr.Code == ErrorCodeNotSupported {
		return nil
	}

	if state.State == UserStateLoggedIn && !relogin {
		return nil
	}

	return auth.login(ctx, username, password, state.PasswordType)
}

func stringSHA256(v string) string {
	h := sha256.New()
	h.Write([]byte(v))
	return hex.EncodeToString(h.Sum(nil))
}

func stringBase64(v string) string {
	return base64.StdEncoding.EncodeToString([]byte(v))
}

func generateConcentratedPasswordSHA256(username, password, verificationToken string) string {
	buf := &bytes.Buffer{}

	buf.WriteString(username)
	buf.WriteString(stringBase64(stringSHA256(password)))
	buf.WriteString(verificationToken)

	return stringBase64(stringSHA256(buf.String()))
}

type userLoginRequest struct {
	XMLName      xml.Name `xml:"request"`
	Username     string   `xml:"Username"`
	Password     string   `xml:"Password"`
	PasswordType int      `xml:"password_type"`
}

func (auth *ClientUser) login(ctx context.Context, username, password string, pt UserPasswordType) error {
	if password != "" {
		switch pt {
		case UserPasswordTypeSHA256:
			password = generateConcentratedPasswordSHA256(username, password, auth.session.Tokens[0])
		default:
			password = base64.StdEncoding.EncodeToString([]byte(password))
		}
	}

	if err := auth.withSessionRetry(ctx, func(ctx context.Context) error {
		return auth.post(ctx, "/api/user/login", &userLoginRequest{
			Username:     username,
			Password:     password,
			PasswordType: int(pt),
		}, true, nil)
	}); err != nil {
		return xerrors.Errorf("post /api/user/login: %w", err)
	}

	return nil
}
