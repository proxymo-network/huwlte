package huwlte

import (
	"context"
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientDevice_BasicInformation(t *testing.T) {
	ctx := context.Background()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.String() {
		case "/api/webserver/SesTokInfo":
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "/api/webserver/SesTokInfo", r.URL.String())

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "text/html")

			copyTestdataJSON(t, "testdata/api-webserver-SesTokInfo.xml", w)
		case "/api/device/basic_information":
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
	info, err := client.Device.BasicInformation(ctx)

	assert.NoError(t, err)

	assert.Equal(t, &DeviceBasicInformation{
		XMLName:               xml.Name{Space: "", Local: "response"},
		ProductFamily:         "LTE",
		Classify:              "wingle",
		Multimode:             0,
		RestoreDefaultStatus:  0,
		AutoupdateGuideStatus: 0,
		SimSavePinEnable:      0,
		Name:                  "E8372",
		SoftwareVersion:       "21.333.03.00.00",
		WebUIVersion:          "17.100.21.02.03",
	}, info)
}
