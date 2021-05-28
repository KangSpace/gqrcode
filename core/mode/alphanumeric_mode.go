package mode

import (
	"github.com/gqrcode/core/cons"
	"github.com/gqrcode/util"
	"regexp"
)

// Define Alphanumeric mode handle
// Alphanumeric mode encodes data from `a set of 45 characters`,
// i.e.:
// 		10 numeric digits (0-9)(byte values 30 hex to 39 hex),
//		26 alphabetic characters (A-Z)(byte values 41 hex to 5A hex),
//		9 symbols (SP,$,%,*,+,-,.,/,:)(byte values 20 hex,24 hex, 25 hex, 2A hex 2B hex, 2D to 2F hex, 3A hex respectively)(SP is space.)
// Normally, `two input characters are represented by 11 bits`.
//
// Tips:
// 1. Alphanumeric mode is not available in Version M1 Micro QR Code Symbol.
//

var alphanumericChars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:"
const (
	AlphaNumericRegExpType = "^[0-9A-Z \\$%\\*\\+\\-\\.\\/\\:]+$"
)

type AlphanumericMode struct {
	*AbstractMode
}

func NewAlphanumericMode() *AlphanumericMode{
	return &AlphanumericMode{&AbstractMode{Name: cons.AlphanumericMode}}
}

func (n *AlphanumericMode) DataEncode(qr *QRCodeStruct) (dataStream *util.DataStream){
	panic("implement me")
}

func (n *AlphanumericMode) GetMode() *AbstractMode{
	return n.AbstractMode
}

func(m *AlphanumericMode) IsSupport(data string) bool{
	numericMatch, _ := regexp.Compile(AlphaNumericRegExpType)
	return numericMatch.MatchString(data)
}