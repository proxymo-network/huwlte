package huwlte

import (
	"context"
	"encoding/xml"
	"fmt"
)

type ClientMonitoring struct {
	*Client
}

type MonitoringStatusConnectionStatus int16

const (
	MonitoringStatusConnectionConnecting    MonitoringStatusConnectionStatus = 900
	MonitoringStatusConnectionConnected     MonitoringStatusConnectionStatus = 901
	MonitoringStatusConnectionDisconnected  MonitoringStatusConnectionStatus = 902
	MonitoringStatusConnectionDisconnecting MonitoringStatusConnectionStatus = 903
)

func (v MonitoringStatusConnectionStatus) String() string {
	switch v {
	case MonitoringStatusConnectionConnecting:
		return "connecting"
	case MonitoringStatusConnectionConnected:
		return "connected"
	case MonitoringStatusConnectionDisconnected:
		return "disconnected"
	case MonitoringStatusConnectionDisconnecting:
		return "disconnecting"
	default:
		return fmt.Sprintf("unknown(%d)", v)
	}
}

type MonitoringStatusNetworkType int16

const (
	MonitoringStatusNetworkTypeNoService MonitoringStatusNetworkType = iota
	MonitoringStatusNetworkTypeGSM
	MonitoringStatusNetworkTypeGPRS
	MonitoringStatusNetworkTypeEDGE
	MonitoringStatusNetworkTypeWCDMA
	MonitoringStatusNetworkTypeHSDPA
	MonitoringStatusNetworkTypeHSUPA
	MonitoringStatusNetworkTypeHSPA
	MonitoringStatusNetworkTypeTDSCDMA
	MonitoringStatusNetworkTypeHSPAPlus
	MonitoringStatusNetworkTypeEVDORev0
	MonitoringStatusNetworkTypeEVDORevA
	MonitoringStatusNetworkTypeEVDORevB
	MonitoringStatusNetworkType1xRTT
	MonitoringStatusNetworkTypeUMB
	MonitoringStatusNetworkType1xEVDV
	MonitoringStatusNetworkType3xRTT
	MonitoringStatusNetworkTypeHSPAPlus64QAM
	MonitoringStatusNetworkTypeHSPAPlusMIMO
	MonitoringStatusNetworkTypeLTE
)

func (v MonitoringStatusNetworkType) String() string {
	switch v {
	case MonitoringStatusNetworkTypeNoService:
		return "No Service"
	case MonitoringStatusNetworkTypeGSM:
		return "GSM"
	case MonitoringStatusNetworkTypeGPRS:
		return "GPRS"
	case MonitoringStatusNetworkTypeEDGE:
		return "EDGE"
	case MonitoringStatusNetworkTypeWCDMA:
		return "WCDMA"
	case MonitoringStatusNetworkTypeHSDPA:
		return "HSDPA"
	case MonitoringStatusNetworkTypeHSUPA:
		return "HSUPA"
	case MonitoringStatusNetworkTypeHSPA:
		return "HSPA"
	case MonitoringStatusNetworkTypeTDSCDMA:
		return "TDSCDMA"
	case MonitoringStatusNetworkTypeHSPAPlus:
		return "HSPA Plus"
	case MonitoringStatusNetworkTypeEVDORev0:
		return "EVDORev0"
	case MonitoringStatusNetworkTypeEVDORevA:
		return "EVDORevA"
	case MonitoringStatusNetworkTypeEVDORevB:
		return "EVDORevB"
	case MonitoringStatusNetworkType1xRTT:
		return "1xRTT"
	case MonitoringStatusNetworkTypeUMB:
		return "UMB"
	case MonitoringStatusNetworkType1xEVDV:
		return "1xEVDV"
	case MonitoringStatusNetworkType3xRTT:
		return "3xRTT"
	case MonitoringStatusNetworkTypeHSPAPlus64QAM:
		return "HSPA Plus 64QAM"
	case MonitoringStatusNetworkTypeHSPAPlusMIMO:
		return "HSPA Plus MIMO"
	case MonitoringStatusNetworkTypeLTE:
		return "LTE"
	default:
		return fmt.Sprintf("unknown(%d)", v)
	}
}

type MonitoringStatusService int8

const (
	MonitoringStatusServiceAvailable MonitoringStatusService = 2
)

type MonitoringStatus struct {
	XMLName xml.Name `xml:"response" human:"-"`

	// Represents current connection link.
	// Use MonitoringStatusConnectionStatus constants to get the string representation and equality.
	ConnectionStatus     MonitoringStatusConnectionStatus `xml:"ConnectionStatus"`
	WifiConnectionStatus string                           `xml:"WifiConnectionStatus"`
	SignalStrength       string                           `xml:"SignalStrength"`
	SignalIcon           int                              `xml:"SignalIcon"`

	// Represents current network type.
	CurrentNetworkType   MonitoringStatusNetworkType `xml:"CurrentNetworkType"`
	CurrentServiceDomain int                         `xml:"CurrentServiceDomain"`
	RoamingStatus        int                         `xml:"RoamingStatus"`
	BatteryStatus        string                      `xml:"BatteryStatus"`
	BatteryLevel         string                      `xml:"BatteryLevel"`
	BatteryPercent       string                      `xml:"BatteryPercent"`
	SimLockStatus        string                      `xml:"simlockStatus"`
	PrimaryDNS           string                      `xml:"PrimaryDns"`
	SecondaryDNS         string                      `xml:"SecondaryDns"`
	PrimaryIPv6DNS       string                      `xml:"PrimaryIPv6Dns"`
	SecondaryIPv6DNS     string                      `xml:"SecondaryIPv6Dns"`
	CurrentWiFiUser      string                      `xml:"CurrentWifiUser"`
	TotalWiFiUser        string                      `xml:"TotalWifiUser"`
	CurrentTotalWiFiUser string                      `xml:"currenttotalwifiuser"`
	ServiceStatus        string                      `xml:"ServiceStatus"`
	SimStatus            string                      `xml:"SimStatus"`
	WifiStatus           string                      `xml:"WifiStatus"`
	CurrentNetworkTypeEx string                      `xml:"CurrentNetworkTypeEx"`
	MaxSignal            int                         `xml:"maxsignal"`
	WifiInDoorOnly       string                      `xml:"wifiindooronly"`
	WiFiFrequence        string                      `xml:"wififrequence"`
	Classify             string                      `xml:"classify"`
	Flymode              string                      `xml:"flymode"`
	CellRoam             string                      `xml:"cellroam"`
	LTECAStatus          string                      `xml:"ltecastatus"`
}

func (monitoring *ClientMonitoring) Status(ctx context.Context) (*MonitoringStatus, error) {
	var result MonitoringStatus

	if err := monitoring.withSessionRetry(ctx, func(ctx context.Context) error {
		return monitoring.get(ctx, "/api/monitoring/status", &result)
	}); err != nil {
		return nil, err
	}

	return &result, nil
}
