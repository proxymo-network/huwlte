package huwlte

import (
	"strings"

	"golang.org/x/xerrors"
)

type NetworkMode string

const (
	NetworkModeAuto     NetworkMode = "00"
	NetworkMode2GOnly   NetworkMode = "01"
	NetworkMode3GOnly   NetworkMode = "02"
	NetworkMode4GOnly   NetworkMode = "03"
	NetworkMode4G3GAuto NetworkMode = "0302"
)

var (
	networkModeStr = map[NetworkMode]string{
		NetworkModeAuto:     "auto",
		NetworkMode2GOnly:   "2g_only",
		NetworkMode3GOnly:   "3_gonly",
		NetworkMode4GOnly:   "4g_only",
		NetworkMode4G3GAuto: "4g_3g_auto",
	}
)

func ParseNetworkMode(mode string) (NetworkMode, error) {
	for k, v := range networkModeStr {
		if v == mode {
			return k, nil
		}
	}
	return "", xerrors.Errorf("unknown network mode: %s", mode)
}

func (mode NetworkMode) String() string {
	return networkModeStr[mode]
}

type NetworkBand int64

const (
	NetworkBandBC0A       NetworkBand = 0x01
	NetworkBandBC0B       NetworkBand = 0x02
	NetworkBandBC1        NetworkBand = 0x04
	NetworkBandBC2        NetworkBand = 0x08
	NetworkBandBC3        NetworkBand = 0x10
	NetworkBandBC4        NetworkBand = 0x20
	NetworkBandBC5        NetworkBand = 0x40
	NetworkBandGSM1800    NetworkBand = 0x80
	NetworkBandGSM900     NetworkBand = 0x100
	NetworkBandBC6        NetworkBand = 0x400
	NetworkBandBC7        NetworkBand = 0x800
	NetworkBandBC8        NetworkBand = 0x1000
	NetworkBandBC9        NetworkBand = 0x2000
	NetworkBandBC10       NetworkBand = 0x4000
	NetworkBandBC11       NetworkBand = 0x8000
	NetworkBandGSM850     NetworkBand = 0x80000
	NetworkBandGSM1900    NetworkBand = 0x200000
	NetworkBandUMTSB12100 NetworkBand = 0x400000
	NetworkBandUMTSB21900 NetworkBand = 0x800000
	NetworkBandBC12       NetworkBand = 0x10000000
	NetworkBandBC13       NetworkBand = 0x20000000
	NetworkBandUMTSB5850  NetworkBand = 0x40000000
	NetworkBandBC14       NetworkBand = 0x80000000
	NetworkBandUMTSB8900  NetworkBand = 0x2000000000000
	NetworkBandAll        NetworkBand = 0x3ffffffffffffff
)

var (
	networkBandStr = map[NetworkBand]string{
		NetworkBandBC0A:       "BC0A",
		NetworkBandBC0B:       "BC0B",
		NetworkBandBC1:        "BC1",
		NetworkBandBC2:        "BC2",
		NetworkBandBC3:        "BC3",
		NetworkBandBC4:        "BC4",
		NetworkBandBC5:        "BC5",
		NetworkBandGSM1800:    "GSM1800",
		NetworkBandGSM900:     "GSM900",
		NetworkBandBC6:        "BC6",
		NetworkBandBC7:        "BC7",
		NetworkBandBC8:        "BC8",
		NetworkBandBC9:        "BC9",
		NetworkBandBC10:       "BC10",
		NetworkBandBC11:       "BC11",
		NetworkBandGSM850:     "GSM850",
		NetworkBandGSM1900:    "GSM1900",
		NetworkBandUMTSB12100: "UMTSB12100",
		NetworkBandUMTSB21900: "UMTSB21900",
		NetworkBandBC12:       "BC12",
		NetworkBandBC13:       "BC13",
		NetworkBandUMTSB5850:  "UMTSB5850",
		NetworkBandBC14:       "BC14",
		NetworkBandUMTSB8900:  "UMTSB8900",
		NetworkBandAll:        "ALL",
	}

	NetworkBandList = []NetworkBand{
		NetworkBandBC0A,
		NetworkBandBC0B,
		NetworkBandBC1,
		NetworkBandBC2,
		NetworkBandBC3,
		NetworkBandBC4,
		NetworkBandBC5,
		NetworkBandGSM1800,
		NetworkBandGSM900,
		NetworkBandBC6,
		NetworkBandBC7,
		NetworkBandBC8,
		NetworkBandBC9,
		NetworkBandBC10,
		NetworkBandBC11,
		NetworkBandGSM850,
		NetworkBandGSM1900,
		NetworkBandUMTSB12100,
		NetworkBandUMTSB21900,
		NetworkBandBC12,
		NetworkBandBC13,
		NetworkBandUMTSB5850,
		NetworkBandBC14,
		NetworkBandUMTSB8900,
		NetworkBandAll,
	}
)

func ParseNetworkBand(band string) (NetworkBand, error) {
	band = strings.ToLower(band)
	for k, v := range networkBandStr {
		if v == band {
			return k, nil
		}
	}
	return NetworkBand(-1), xerrors.Errorf("unknown network band: %s", band)
}

func (band NetworkBand) String() string {
	return networkBandStr[band]
}
