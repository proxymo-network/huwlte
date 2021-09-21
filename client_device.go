package huwlte

import (
	"context"
	"encoding/xml"
)

type DeviceBasicInformation struct {
	XMLName               xml.Name `xml:"response" human:"-"`
	ProductFamily         string   `xml:"productfamily"`
	Classify              string   `xml:"classify"`
	Multimode             int8     `xml:"multimode"`
	RestoreDefaultStatus  int8     `xml:"restore_default_status"`
	AutoupdateGuideStatus int8     `xml:"autoupdate_guide_status"`
	SimSavePinEnable      int8     `xml:"sim_save_pin_enable"`
	Name                  string   `xml:"devicename"`
	SoftwareVersion       string   `xml:"SoftwareVersion"`
	WebUIVersion          string   `xml:"WebUIVersion"`
}

type ClientDevice struct {
	*Client
}

// BasicInformation returns the basic information of the device.
func (device *Client) BasicInformation(ctx context.Context) (*DeviceBasicInformation, error) {
	var result DeviceBasicInformation

	if err := device.withSessionRetry(ctx, func(ctx context.Context) error {
		return device.get(ctx, "/api/device/basic_information", &result)
	}); err != nil {
		return nil, err
	}

	return &result, nil
}

// Cotnrol power of device. If v = 1 reboot, device.
func (device *Client) Control(ctx context.Context, v int) error {
	var req = struct {
		XMLName xml.Name `xml:"request"`
		Control int      `xml:"Control"`
	}{
		Control: v,
	}

	if err := device.withSessionRetry(ctx, func(ctx context.Context) error {
		return device.post(ctx, "/api/device/control", req, false, nil)
	}); err != nil {
		return err
	}

	return nil
}

type DeviceInformation struct {
	XMLName         xml.Name `xml:"response"`
	DeviceName      string   `xml:"DeviceName"`
	SerialNumber    string   `xml:"SerialNumber"`
	IMEI            string   `xml:"Imei"`
	IMSI            string   `xml:"Imsi"`
	ICCID           string   `xml:"Iccid"`
	MSISDN          string   `xml:"Msisdn"`
	HardwareVersion string   `xml:"HardwareVersion"`
	SoftwareVersion string   `xml:"SoftwareVersion"`
	WebUIVersion    string   `xml:"WebUIVersion"`
	MacAddress1     string   `xml:"MacAddress1"`
	MacAddress2     string   `xml:"MacAddress2"`
	ProductFamily   string   `xml:"ProductFamily"`
	Classify        string   `xml:"Classify"`
	SupportMode     string   `xml:"supportmode"`
	Workmode        string   `xml:"workmode"`
	WanIPAddress    string   `xml:"WanIPAddress"`
	WanIPv6Address  string   `xml:"WanIPv6Address"`
}

// BasicInformation returns the basic information of the device.
func (device *Client) Information(ctx context.Context) (*DeviceInformation, error) {
	var result DeviceInformation

	if err := device.withSessionRetry(ctx, func(ctx context.Context) error {
		return device.get(ctx, "/api/device/information", &result)
	}); err != nil {
		return nil, err
	}

	return &result, nil
}
