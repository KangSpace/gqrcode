package mode

import (
	"github.com/KangSpace/gqrcode/core/cons"
	"github.com/KangSpace/gqrcode/util"
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
// 2. Alphanumeric mode max capacity is 4296

var alphanumericChars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:"

// AlphanumericModeCodingTable :Page 34,Table 5-Encoding/decoding table for Alphanumeric mode
var AlphanumericModeCodingTable = map[string]int{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
	"A": 10, "B": 11, "C": 12, "D": 13, "E": 14, "F": 15, "G": 16, "H": 17, "I": 18, "J": 19, "K": 20, "L": 21, "M": 22, "N": 23, "O": 24, "P": 25, "Q": 26, "R": 27, "S": 28, "T": 29, "U": 30, "V": 31, "W": 32, "X": 33, "Y": 34, "Z": 35,
	" ": 36, "$": 37, "%": 38, "*": 39, "+": 40, "-": 41, ".": 42, "/": 43, ":": 44,
}

const AlphaNumericRegExpType = "^[0-9A-Z \\$%\\*\\+\\-\\.\\/\\:]+$"

type AlphanumericMode struct {
	*AbstractMode
}

func NewAlphanumericMode() *AlphanumericMode {
	return &AlphanumericMode{&AbstractMode{Name: cons.AlphanumericMode}}
}

func (am *AlphanumericMode) DataEncode(qr *QRCodeStruct) (dataStream *util.DataStream) {
	return am.BuildEncodeData(qr, am.buildDataBits)
}

func getCodingValByRaw(raw string) int {
	return AlphanumericModeCodingTable[raw]
}

// Page 35, 7.4.4 Alphanumeric mode
// Input data characters are divided into groups of 2(two) characters which are encoded as 11-bit binary codes.
// The character value of the first character is multiplied by 45 and the character value of the second digit is added to the product.
// The sum is then converted to an 11-bit binary number. If the number of input data characters is not a multiple of 2(two),
//   the character value of the final character is encoded as a 6-bit binary number.
// The binary data is then concatenated and prefixed with the mode indicator and the character count indicator.
// The mode indicator in the Alphanumeric mode has either 4 bits for QRCode symbols or the number of bits defined in Table2 for Micro QRCode symbols,
//	 abd the character count indicator has the number of  bits defined in Table3.
func (am *AlphanumericMode) buildDataBits(qr *QRCodeStruct, dataStream *util.DataStream) (dataBitLen int) {
	inputData := qr.Data
	dataLen := len(inputData)
	charsNumPerGroup := 2
	// divided into groups of 2(two) digits
	for i := 0; i < dataLen; i += charsNumPerGroup {
		var group string
		if i+charsNumPerGroup <= dataLen {
			group = inputData[i : i+charsNumPerGroup]
		} else {
			// the last group
			group = inputData[i:dataLen]
		}
		// group is converted to its `11-bit binary equivalent`.
		groupBitsLen := 11
		// the value of first character
		bit11 := getCodingValByRaw(group[:1])
		if len(group) == 1 {
			// the final one character is converted to 6 bits respectively
			groupBitsLen = 6
		} else {
			bit11 = bit11*45 + getCodingValByRaw(group[1:])
		}
		dataStream.AddIntBit(bit11, groupBitsLen)
		dataBitLen += groupBitsLen
	}
	return dataBitLen
}

func (n *AlphanumericMode) GetMode() *AbstractMode {
	return n.AbstractMode
}

func (m *AlphanumericMode) IsSupport(data string) bool {
	match, _ := regexp.Compile(AlphaNumericRegExpType)
	return match.MatchString(data)
}
