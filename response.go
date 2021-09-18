package huwlte

import (
	"encoding/xml"

	"golang.org/x/xerrors"
)

type responseEnvelope struct {
	XMLName xml.Name
	Content []byte `xml:"-"`
}

func parseResponseEnvelope(content []byte) (*responseEnvelope, error) {
	env := &responseEnvelope{
		Content: content,
	}
	if err := xml.Unmarshal(content, &env); err != nil {
		return nil, err
	}
	return env, nil
}

type responseError struct {
	XMLName xml.Name `xml:"error"`
	Code    int      `xml:"code"`
	Message string   `xml:"message"`
}

func (env *responseEnvelope) isError() bool {
	return env.XMLName.Local == "error"
}

func (env *responseEnvelope) decode(dst interface{}) error {
	return xml.Unmarshal(env.Content, dst)
}

func (env *responseEnvelope) toErr() error {
	if !env.isError() {
		return nil
	}

	var responseErr responseError
	if err := env.decode(&responseErr); err != nil {
		return xerrors.Errorf("decode response error: %w", err)
	}

	return newError(responseErr.Code, responseErr.Message)
}
