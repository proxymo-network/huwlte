package huwlte

import (
	"context"
	"encoding/xml"
)

type ClientDialup struct {
	*Client
}

type mobileSwitchResponse struct {
	XMLName    xml.Name `xml:"response"`
	Dataswitch int      `xml:"dataswitch"`
}

// MobileSwitch returns status of mobile connection.
func (dialup *ClientDialup) MobileSwitch(ctx context.Context) (bool, error) {
	var result mobileSwitchResponse

	if err := dialup.withSessionRetry(ctx, func(ctx context.Context) error {
		return dialup.Get(ctx, "/api/dialup/mobile-dataswitch", &result)
	}); err != nil {
		return false, err
	}

	return result.Dataswitch == 1, nil
}

// SetMobileSwitch sets mobile connection status.
func (dialup *ClientDialup) SetMobileSwitch(ctx context.Context, status bool) error {
	var data struct {
		XMLName xml.Name `xml:"request"`

		Dataswitch int `xml:"dataswitch"`
	}

	if status {
		data.Dataswitch = 1
	}

	if err := dialup.withSessionRetry(ctx, func(ctx context.Context) error {
		return dialup.post(ctx, "/api/dialup/mobile-dataswitch", data, false, nil)
	}); err != nil {
		return err
	}

	return nil
}
