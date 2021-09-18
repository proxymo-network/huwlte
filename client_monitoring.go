package huwlte

import (
	"context"
	"encoding/xml"
)

type ClientMonitoring struct {
	*Client
}

type MonitoringStatus struct {
	XMLName              xml.Name `xml:"response" human:"-"`
	ConnectionStatus     string   `xml:"ConnectionStatus"`
	WifiConnectionStatus string   `xml:"WifiConnectionStatus"`
	SignalStrength       string   `xml:"SignalStrength"`
	SignalIcon           string   `xml:"SignalIcon"`
	CurrentNetworkType   string   `xml:"CurrentNetworkType"`
	CurrentServiceDomain string   `xml:"CurrentServiceDomain"`
	RoamingStatus        string   `xml:"RoamingStatus"`
	BatteryStatus        string   `xml:"BatteryStatus"`
	BatteryLevel         string   `xml:"BatteryLevel"`
	BatteryPercent       string   `xml:"BatteryPercent"`
	SimLockStatus        string   `xml:"simlockStatus"`
	PrimaryDNS           string   `xml:"PrimaryDns"`
	SecondaryDNS         string   `xml:"SecondaryDns"`
	PrimaryIPv6DNS       string   `xml:"PrimaryIPv6Dns"`
	SecondaryIPv6DNS     string   `xml:"SecondaryIPv6Dns"`
	CurrentWiFiUser      string   `xml:"CurrentWifiUser"`
	TotalWiFiUser        string   `xml:"TotalWifiUser"`
	CurrentTotalWiFiUser string   `xml:"currenttotalwifiuser"`
	ServiceStatus        string   `xml:"ServiceStatus"`
	SimStatus            string   `xml:"SimStatus"`
	WifiStatus           string   `xml:"WifiStatus"`
	CurrentNetworkTypeEx string   `xml:"CurrentNetworkTypeEx"`
	Maxsignal            string   `xml:"maxsignal"`
	WifiInDoorOnly       string   `xml:"wifiindooronly"`
	WiFiFrequence        string   `xml:"wififrequence"`
	Classify             string   `xml:"classify"`
	Flymode              string   `xml:"flymode"`
	CellRoam             string   `xml:"cellroam"`
	LTECAStatus          string   `xml:"ltecastatus"`
}

func (monitoring *ClientMonitoring) Status(ctx context.Context) (*MonitoringStatus, error) {
	var result MonitoringStatus

	if err := monitoring.withSessionRetry(ctx, func(ctx context.Context) error {
		return monitoring.Get(ctx, "/api/monitoring/status", &result)
	}); err != nil {
		return nil, err
	}

	return &result, nil
}
