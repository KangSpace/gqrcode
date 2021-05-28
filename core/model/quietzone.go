package model

import "github.com/gqrcode/core/cons"

// Define Quiet Zone here

type QuietZone struct {
	// Times:
	// QRCode minimum: 4x
	// Micro QRCode minimum: 2x
	Size int `json:"size"`
}

// NewDefaultQuietZone :new default quiet zone by version
func NewDefaultQuietZone(version *Version) *QuietZone{
	qz := new(QuietZone)
	// QRCode
	if version.Id>0 {
		qz.Size = 4
	}else{
		// Micro QRCode
		qz.Size = 2
	}
	return qz
}

func NewQuietZone(size int) *QuietZone{
	return &QuietZone{Size: size}
}

func (qz *QuietZone) GetDefaultPixelSize() int{
	return qz.Size * cons.DefaultPixelSizePerModule
}

// NoneQuietZone : define quiet zone for zero
var NoneQuietZone = NewQuietZone(0)
// AutoQuietZone : auto quiet zone, quiet zone size by NewDefaultQuietZone()
var AutoQuietZone = NewQuietZone(-1)
