package mode

import (
	"github.com/gqrcode/core/cons"
	"github.com/gqrcode/util"
	"regexp"
)

// Define Numeric mode handle
// Numeric mode encodes data from the decimal digit set `(0-9)(byte values 30hex to 39hex)`.
// Normally,3 data characters are represented by 10bits;

var numericChars = "0123456789"

const (
	NumericRegExpType      = "^[0-9]+$"
)

type NumericMode struct {
	*AbstractMode
}

func NewNumericMode() *NumericMode{
	return &NumericMode{&AbstractMode{Name: cons.NumericMode}}
}

// DataEncode : Data encoding
// For any number of data characters the length of the bit stream in Numeric mode is given by the following formula:
// B = M + C + 10(D DIV 3) + R
// where:
// 	B	number of bits in bit stream
//  M	number of bits in mod indicator(4 for QRCode symbols,or as shown in Table 2 for Micro QR code symbols)
//	C 	number of bits in character count indicator(from Table3)
//  D	number of input data characters
//	R	0 if(D MOD 3) = 0
//	R	4 if(D MOD 3) = 1
//	R	7 if(D MOD 3) = 2
func (nm *NumericMode) DataEncode(qr *QRCodeStruct) (dataStream *util.DataStream){
	dataStream = util.NewDataStream(16)
	// Mode indicator bits
	nm.buildModeIndicator(qr,dataStream)
	// Data character count indicator bits
	_,_,numberOfDataBits := nm.buildCharacterCountIndicator(qr,dataStream)
	// Data bits
	nm.buildDataBits(qr,dataStream)

	terminatorBitsLen := numberOfDataBits - dataStream.GetCount()
	// add Terminator
	nm.buildTerminators(qr,dataStream,terminatorBitsLen)
	return dataStream
}

func (nm *NumericMode) GetMode() *AbstractMode{
	return nm.AbstractMode
}

func(nm *NumericMode) IsSupport(data string) bool{
	numericMatch, _ := regexp.Compile(NumericRegExpType)
	return numericMatch.MatchString(data)
}

