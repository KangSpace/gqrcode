package mode

import (
	"fmt"
	"github.com/gqrcode/core/cons"
	"github.com/gqrcode/core/logger"
	"github.com/gqrcode/util"
	"regexp"
	"strconv"
)

// Define Numeric mode handle
// Numeric mode encodes data from the decimal digit set `(0-9)(byte values 30hex to 39hex)`.
// Normally,3 data characters are represented by 10bits;
// Tips
// 1. Numeric mode max capacity is 7089

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
func (nm *NumericMode)  DataEncode(qr *QRCodeStruct) (dataStream *util.DataStream){
	//dataStream = util.NewDataStream(16)
	//// Mode indicator bits
	//nm.buildModeIndicator(qr,dataStream)
	//// Data character count indicator bits
	//_,_,numberOfDataBits := nm.buildCharacterCountIndicator(qr,dataStream)
	//// Data bits
	//nm.buildDataBits(qr,dataStream)
	//
	//terminatorBitsLen := numberOfDataBits - dataStream.GetCount()
	//// add Terminator
	//nm.buildTerminators(qr,dataStream,terminatorBitsLen)
	return nm.BuildEncodeData(qr,nm.buildDataBits)
}

// Page 33, 7.4.3 Numeric mode
// The input data string is `divided into groups of 3(three) digits`, and each group is converted to its `10-bit binary equivalent`.
// If the number of input digits is not an exact multiple of 3(three),`the final one or two digits are converted to 4 or 7 bits respectively`.
func (m *NumericMode) buildDataBits(qr *QRCodeStruct,dataStream *util.DataStream) (dataBitLen int) {
	inputData := qr.Data
	dataLen := len(inputData)

	// divided into groups of 3(three) digits
	for i:=0;i<dataLen;i+=3{
		var group string
		if i + 3<= dataLen{
			group = inputData[i:i+3]
		}else{
			// the last group
			group = inputData[i:dataLen]
		}

		// group is converted to its `10-bit binary equivalent`.
		bit10, err := strconv.Atoi(group)
		if err != nil {
			err := fmt.Errorf("buildDataBits: \"%s\" can not be encoded as %s . %s", inputData, cons.NumericMode,err.Error())
			logger.Error("buildDataBits: " + err.Error())
			panic(err)
		}
		//
		groupBitsLen := 10
		// the final one or two digits are converted to 4 or 7 bits respectively
		switch len(group) % 3 {
		case 1:
			groupBitsLen = 4
		case 2:
			groupBitsLen = 7
		}
		dataStream.AddIntBit(bit10,groupBitsLen)
		dataBitLen += groupBitsLen
	}
	return dataBitLen
}


func (nm *NumericMode) GetMode() *AbstractMode{
	return nm.AbstractMode
}

func(nm *NumericMode) IsSupport(data string) bool{
	match, _ := regexp.Compile(NumericRegExpType)
	return match.MatchString(data)
}

