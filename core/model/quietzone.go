package model

import "github.com/KangSpace/gqrcode/core/cons"

// Define Quiet Zone here

type QuietZone struct {
	// Times:
	// QRCode minimum: 4x
	// Micro QRCode minimum: 2x
	Multiple int `json:"multiple"`
}

var (
	// QuietZoneZero : no quiet zone
	QuietZoneZero = &QuietZone{0}
	// QuietZoneOne : one multiple quiet zone
	QuietZoneOne = &QuietZone{1}
	// QuietZoneTwo : two multiple quiet zone
	QuietZoneTwo = &QuietZone{2}
	// QuietZoneFour : four multiple quiet zone
	QuietZoneFour = &QuietZone{4}
	QuietZones    = []*QuietZone{QuietZoneZero, QuietZoneOne, QuietZoneTwo, QuietZoneFour}
)

// NewDefaultQuietZone :new default quiet zone by version
func NewDefaultQuietZone(version *Version) *QuietZone {
	return NewPopularQuietZone()
}

// NewStandardQuietZone :new standard quiet zone by version, QRCode is 4 multiple, Micro QRCode is 2.
func NewStandardQuietZone(version *Version) *QuietZone {
	qz := new(QuietZone)
	// QRCode
	if version.Id > 0 {
		qz.Multiple = 4
	} else {
		// Micro QRCode
		qz.Multiple = 2
	}
	return qz
}

// NewPopularQuietZone :new popular quiet zone, it's 2 multiple size.
func NewPopularQuietZone() *QuietZone {
	return GetQuietZone(2)
}

func NewQuietZone(size int) *QuietZone {
	return &QuietZone{Multiple: size}
}

// GetQuietZoneSize :Quiet zone size is QuietZone.Multiple * 2.
func (qz *QuietZone) GetQuietZoneSize() int {
	return qz.Multiple * 2
}

func (qz *QuietZone) GetDefaultPixelSize() int {
	return qz.GetQuietZoneSize() * cons.DefaultPixelSizePerModule
}

// NoneQuietZone : define quiet zone for zero
var NoneQuietZone = NewQuietZone(0)

// AutoQuietZone : auto quiet zone, quiet zone size by NewDefaultQuietZone()
var AutoQuietZone = NewQuietZone(-1)

// GetQuietZone : Get quiet zone by quietZoneMultiple.
func GetQuietZone(quietZoneMultiple int) *QuietZone {
	for _, zone := range QuietZones {
		if zone.Multiple == quietZoneMultiple {
			return zone
		}
	}
	return NewQuietZone(quietZoneMultiple)
}
