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
	device.withSessionRetry(ctx, func(ctx context.Context) error {
		return device.get(ctx, "/api/device/basic_information", &result)
	})
	return &result, nil
}
