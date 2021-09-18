package huwlte

type NetworkMode string

const (
	NetworkModeAuto     NetworkMode = "00"
	NetworkMode2GOnly   NetworkMode = "01"
	NetworkMode3GOnly   NetworkMode = "02"
	NetworkMode4GOnly   NetworkMode = "03"
	NetworkMode4G3GAuto NetworkMode = "0302"
)

const (
	NetworkBandBC0A       = 0x01
	NetworkBandBC0B       = 0x02
	NetworkBandBC1        = 0x04
	NetworkBandBC2        = 0x08
	NetworkBandBC3        = 0x10
	NetworkBandBC4        = 0x20
	NetworkBandBC5        = 0x40
	NetworkBandGSM1800    = 0x80
	NetworkBandGSM900     = 0x100
	NetworkBandBC6        = 0x400
	NetworkBandBC7        = 0x800
	NetworkBandBC8        = 0x1000
	NetworkBandBC9        = 0x2000
	NetworkBandBC10       = 0x4000
	NetworkBandBC11       = 0x8000
	NetworkBandGSM850     = 0x80000
	NetworkBandGSM1900    = 0x200000
	NetworkBandUMTSB12100 = 0x400000
	NetworkBandUMTSB21900 = 0x800000
	NetworkBandBC12       = 0x10000000
	NetworkBandBC13       = 0x20000000
	NetworkBandUMTSB5850  = 0x40000000
	NetworkBandBC14       = 0x80000000
	NetworkBandUMTSB8900  = 0x2000000000000
	NetworkBandAll        = 0x3ffffffffffffff
)
