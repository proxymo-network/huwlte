package huwlte

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"errors"
	"os"
	"path"
	"strings"

	"golang.org/x/xerrors"
)

type Session struct {
	XMLName xml.Name `xml:"response" json:"-"`
	Cookie  string   `xml:"SesInfo"`
	Tokens  []string `xml:"TokInfo"`
}

func (s *Session) HasToken() bool {
	return len(s.Tokens) > 0
}

func (s *Session) AddToken(v string) {
	s.Tokens = append(s.Tokens, v)
}

func (s *Session) Reset() {
	s.Cookie = ""
	s.Tokens = []string{}
}

func (s *Session) ResetTokens() {
	s.Tokens = []string{}
}

func (s *Session) HasMultipleTokens() bool {
	return len(s.Tokens) > 1
}

func (s *Session) Token() string {
	return s.Tokens[0]
}

func (s *Session) PopToken() string {
	var token string

	token, s.Tokens = s.Tokens[0], s.Tokens[1:]

	return token
}

func (s *Session) empty() bool {
	return s.Cookie == "" || len(s.Tokens) == 0
}

type SessionStorage interface {
	Save(ctx context.Context, name string, s *Session) error
	Load(ctx context.Context, name string) (*Session, error)
	Remove(ctx context.Context, name string) error
}

type LocalSessionStorage struct {
	Dir string
}

func (ls *LocalSessionStorage) Load(ctx context.Context, name string) (*Session, error) {
	filename := path.Join(ls.Dir, name+".json")
	f, err := os.Open(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return &Session{}, nil
		}
		return nil, xerrors.Errorf("open file: %w", err)
	}
	defer f.Close()

	var s Session
	if err := json.NewDecoder(f).Decode(&s); err != nil {
		return nil, xerrors.Errorf("decode session: %w", err)
	}

	return &s, nil
}

func (ls *LocalSessionStorage) Save(ctx context.Context, name string, s *Session) error {
	filename := path.Join(ls.Dir, name+".json")
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		return xerrors.Errorf("open file: %w", err)
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(s); err != nil {
		return xerrors.Errorf("encode session: %w", err)
	}

	return nil
}

func (ls *LocalSessionStorage) Remove(ctx context.Context, name string) error {
	filename := path.Join(ls.Dir, name+".json")

	if err := os.Remove(filename); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return xerrors.Errorf("remove file: %w", err)
	}

	return nil
}

func SessionName(parts []string) string {
	v := strings.Join(parts, "|")
	h := sha256.New()
	h.Write([]byte(v))
	return hex.EncodeToString(h.Sum(nil))
}
