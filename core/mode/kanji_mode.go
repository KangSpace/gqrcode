package mode

import (
	"fmt"
	"github.com/KangSpace/gqrcode/core/cons"
	"github.com/KangSpace/gqrcode/util"
	"strconv"
)

// Define Kanji mode handle
// The Kanji mode efficiently encodes Kanji characters in accordance with the Shift JIS system based on JIS X 0209.
// The Shift JIS values are shifted from the JIS X 0208 values.
// JIS X 0208 gives details of the shift coded representation.
// Each two-byte character value is compacted to a 13-bit binary codeword.
//
// Tips:
// 1. Kanji mode is not available in Version M1 Or M2 Micro QR Code Symbol.
// 2. Kanji mode max capacity is 1817
// 3. Kanji mode need input ShiftJIS coding format string

// KanjiRegExpType : double-byte HEX: \x8140-\x9FFC or \xE040-\xEBBF
//const KanjiRegExpType  = "^([\x81\x40-\x9F\xFC]|[\xE0\x40-\xEB\xBF])*$"

const hex8140 = 0x8140
const hexC140 = 0xC140
const hexE040 = 0xE040

type KanjiMode struct {
	*AbstractMode
}

func NewKanjiModeMode() *KanjiMode {
	return &KanjiMode{&AbstractMode{Name: cons.KanjiMode}}
}

func (km *KanjiMode) DataEncode(qr *QRCodeStruct) (dataStream *util.DataStream) {
	return km.BuildEncodeData(qr, km.buildDataBits)
}

func (km *KanjiMode) buildDataBits(qr *QRCodeStruct, dataStream *util.DataStream) (dataBitLen int) {
	inputData := qr.Data
	loop := len(inputData) / 2
	twoByteChain := util.IteratorTwoByteEach(inputData)
	for i := 0; i < loop; i++ {
		twoByteIntArr := <-twoByteChain
		twoByteInt := twoByteIntArr[0] | twoByteIntArr[1]
		groupBitsLen := 13
		var bit13 uint16
		var subResult uint16

		fmt.Printf("a:%v %v %v\n", twoByteInt, strconv.FormatInt(int64(twoByteInt), 16), strconv.FormatInt(int64(twoByteIntArr[0]), 16))
		if is8140To9FFC(twoByteInt) {
			subResult = twoByteInt - hex8140
		} else {
			subResult = twoByteInt - hexC140
		}
		mostSignByte := subResult >> 8
		leastSignByte := uint16(uint8(subResult))
		fmt.Printf("subResult:%v mostSignByte:%v leastSignByte:%v \n", strconv.FormatInt(int64(subResult), 16), strconv.FormatInt(int64(mostSignByte), 16), strconv.FormatInt(int64(leastSignByte), 16))
		bit13 = mostSignByte*0xC0 + leastSignByte
		dataStream.AddIntBit16(bit13, groupBitsLen)
		dataBitLen += groupBitsLen
	}
	return dataBitLen
}

func (km *KanjiMode) GetMode() *AbstractMode {
	return km.AbstractMode
}

func (km *KanjiMode) IsSupport(data string) bool {
	dataLen := len(data)
	if dataLen%2 > 0 {
		return false
	}
	loop := dataLen / 2
	for i := 0; i < loop; i++ {
		twoByteInt := <-util.IteratorTwoByte(data)
		if !(is8140To9FFC(twoByteInt)) &&
			!(isE040ToEBBF(twoByteInt)) {

			return false
		}
	}
	return true
}

func is8140To9FFC(twoByteInt uint16) bool {
	return hex8140 <= twoByteInt && twoByteInt <= 0x9FFC
}
func isE040ToEBBF(twoByteInt uint16) bool {
	return hexE040 <= twoByteInt && twoByteInt <= 0xEBBF
}
