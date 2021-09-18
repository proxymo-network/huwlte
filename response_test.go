package huwlte

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func loadTestdataJSON(t *testing.T, path string) []byte {
	t.Helper()

	f, err := os.Open(path)
	assert.NoError(t, err)
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	assert.NoError(t, err)

	return data
}

func TestResponseEnvelope_Error(t *testing.T) {
	xml := loadTestdataJSON(t, "testdata/error.xml")

	ev, err := parseResponseEnvelope(xml)
	require.NoError(t, err)
	assert.True(t, ev.isError())

	err = ev.toErr()
	require.NotNil(t, err)

	var serr *Error
	require.ErrorAs(t, err, &serr)
	assert.Equal(t, ErrorCodeCSRF, serr.Code)
	assert.Equal(t, "csrf", serr.Message)
}
