package mode

import (
	"github.com/gqrcode/core/cons"
	"github.com/gqrcode/util"
	"regexp"
)

// Define Byte mode handle:
// In this mode, data is encoded at `8 bits` per character.
// The default character set for byte mode is ISO-8859-1, and when possible,you should convert your input text to this character set.
//
// Tips:
// 1. Byte mode is not available in Version M1 or M2 Micro QR Code Symbol.
// 2. Byte mode max capacity is 2953

const ByteRegExpType = "^.*$"

type ByteMode struct {
	*AbstractMode
}

func NewByteModeMode() *ByteMode{
	return &ByteMode{&AbstractMode{Name: cons.ByteMode}}
}

func (bm *ByteMode) DataEncode(qr *QRCodeStruct) (dataStream *util.DataStream){
	return bm.BuildEncodeData(qr,bm.buildDataBits)
}

func (bm *ByteMode) buildDataBits(qr *QRCodeStruct,dataStream *util.DataStream) (dataBitLen int) {
	inputData := qr.Data
	groupBitsLen := 8
	for _,char:= range []byte(inputData){
		bit8:= int(char)
		dataStream.AddIntBit(bit8,groupBitsLen)
		dataBitLen += groupBitsLen
	}
	return dataBitLen
}


func (n *ByteMode) GetMode() *AbstractMode{
	return n.AbstractMode
}

func(m *ByteMode) IsSupport(data string) bool{
	match, _ := regexp.Compile(ByteRegExpType)
	return match.MatchString(data)
}