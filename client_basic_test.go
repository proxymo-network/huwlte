package huwlte

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func copyTestdataJSON(t *testing.T, path string, w io.Writer) {
	t.Helper()

	f, err := os.Open(path)
	require.NoError(t, err)
	t.Cleanup(func() {
		f.Close()
	})

	_, err = io.Copy(w, f)
	require.NoError(t, err)
}

func TestClient_getSession(t *testing.T) {
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "/api/webserver/SesTokInfo", r.URL.String())

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "text/html")

			copyTestdataJSON(t, "testdata/api-webserver-SesTokInfo.xml", w)
		}))
		defer server.Close()

		client := New(server.URL)

		err := client.getSession(ctx)

		assert.NoError(t, err, "should be no error")
		assert.NotZero(t, client.session.Cookie, "should have a session cookie")
		assert.NotZero(t, client.session.Tokens, "should have a session token")
	})
}

func TestClient_withSessionRetry(t *testing.T) {
	ctx := context.Background()

	var (
		getSessionCalls          int
		getBasicInformationCalls int
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.String() {
		case "/api/webserver/SesTokInfo":
			getSessionCalls++

			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "/api/webserver/SesTokInfo", r.URL.String())

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "text/html")

			copyTestdataJSON(t, "testdata/api-webserver-SesTokInfo.xml", w)
		case "/api/device/basic_information":
			getBasicInformationCalls++

			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "/api/device/basic_information", r.URL.String())

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "text/html")

			if cookie := r.Header.Get("Cookie"); cookie == "" || cookie == "invalid" {
				copyTestdataJSON(t, "testdata/error.xml", w)
			}

			copyTestdataJSON(t, "testdata/api-device-basic_information.xml", w)
		}
	}))
	defer server.Close()

	client := New(server.URL)

	err := client.withSessionRetry(ctx, func(ctx context.Context) error {
		return client.get(ctx, "/api/device/basic_information", nil)
	})

	assert.NoError(t, err, "should be no error")
	assert.Equal(t, 1, getBasicInformationCalls, "should have called getSession twice")
	assert.Equal(t, 1, getSessionCalls, "should have called getSession once")

	client.session.Cookie = "invalid"
	getBasicInformationCalls = 0
	getSessionCalls = 0

	err = client.withSessionRetry(ctx, func(ctx context.Context) error {
		return client.get(ctx, "/api/device/basic_information", nil)
	})
	assert.NoError(t, err, "should be no error")

	assert.Equal(t, 1, getSessionCalls, "should have called get session one")
	assert.Equal(t, 2, getBasicInformationCalls, "should have called get basic informatin one")
}
