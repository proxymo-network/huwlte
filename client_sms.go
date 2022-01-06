package huwlte

import (
	"context"
	"encoding/xml"
)

type ClientSMS struct {
	base *Client
}

type SMSCount struct {
	XMLName      xml.Name `xml:"response" human:"-"`
	LocalUnread  int      `xml:"LocalUnread"`
	LocalInbox   int      `xml:"LocalInbox"`
	LocalOutbox  int      `xml:"LocalOutbox"`
	LocalDraft   int      `xml:"LocalDraft"`
	LocalDeleted int      `xml:"LocalDeleted"`
	SimUnread    int      `xml:"SimUnread"`
	SimInbox     int      `xml:"SimInbox"`
	SimOutbox    int      `xml:"SimOutbox"`
	SimDraft     int      `xml:"SimDraft"`
	LocalMax     int      `xml:"LocalMax"`
	SimMax       int      `xml:"SimMax"`
	SimUsed      int      `xml:"SimUsed"`
	NewMsg       int      `xml:"NewMsg"`
}

type SMSList struct {
	XMLName  xml.Name      `xml:"response" human:"-"`
	Count    int           `xml:"Count"`
	Messages []SMSListItem `xml:"Messages>Message"`
}

type SMSListItem struct {
	XMLName  xml.Name `xml:"Message" human:"-"`
	Smstat   int      `xml:"Smstat"`
	Index    int      `xml:"Index"`
	Phone    string   `xml:"Phone"`
	Content  string   `xml:"Content"`
	Date     string   `xml:"Date"`
	Sca      string   `xml:"Sca"`
	SaveType int      `xml:"SaveType"`
	Priority int      `xml:"Priority"`
	SmsType  int      `xml:"SmsType"`
}

type SMSListOptions struct {
	XMLName         xml.Name `xml:"request"`
	PageIndex       int      `xml:"PageIndex"`
	ReadCount       int      `xml:"ReadCount"`
	BoxType         int      `xml:"BoxType"`
	SortType        int      `xml:"SortType"`
	Ascending       int      `xml:"Ascending"`
	UnreadPreferred int      `xml:"UnreadPreferred"`
}

func (sms *ClientSMS) List(ctx context.Context, opts SMSListOptions) (*SMSList, error) {
	var result SMSList

	if err := sms.base.withSessionRetry(ctx, func(ctx context.Context) error {
		return sms.base.post(ctx, "/api/sms/sms-list", opts, false, &result)
	}); err != nil {
		return nil, err
	}

	return &result, nil
}

func (sms *ClientSMS) Count(ctx context.Context) (*SMSCount, error) {
	var result SMSCount

	if err := sms.base.withSessionRetry(ctx, func(ctx context.Context) error {
		return sms.base.get(ctx, "/api/sms/sms-count", &result)
	}); err != nil {
		return nil, err
	}

	return &result, nil
}
