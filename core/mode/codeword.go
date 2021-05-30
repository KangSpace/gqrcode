package mode

import (
	"github.com/gqrcode/core/cons"
	"github.com/gqrcode/core/model"
	"github.com/gqrcode/util"
)

// Define Code Words handle

// BuildCodeWords :Build the codeword for DataStream
// Page 40,7.4.10 Bit stream to codeword conversion
// The resulting message bit stream shall then `be divided into codewords`.
// All codewords are `8 bits in length`, except for the final data symbol character in `Micro QR Code version M1 and M3 symbols,which is 4 bits in length`.
// If the bit stream length is such that it does not end at a codeword boundary,padding bits with binary value 0 shall be added after the
//	final bit(least significant bit) of data stream to extend it to the codeword boundary.
// The message bit stream shall then be extended to fill the data capacity of the symbol corresponding to the Version and Error Correction Level,
//  as defined in Table8, by adding the Pad Codewords 11101100 and -00010001 alternately.
// For Micro QR Code versions M1 and M3 symbols, the final data codeword is 4 bits long.
// The Pad Codeword used in the final data symbol character position in Micro QR Code version M1 and M3 symbols shall be represented as 0000.
// The resulting series of codewords,the data codeword sequence,is then process as described in 7.5 to add error correction codewords to the message.
// In certain versions of symbol,it may be necessary to add 3,4 or 7 Remainder bits(all zeros) to end of the message,
//  after the final error correction codeword,in order exactly to fill the symbol capacity(see Table1)

// PadCodewords : defined the model.QRCODE and model.MICRO_QRCODE Pad Codewords

type CodewordBit = int

const (
	QrcodeCodewordBit          CodewordBit = 8
	MicroQrcodeM1m3CodewordBit CodewordBit = 4
)

var PadCodewords = map[cons.Format][][]byte{
	// 8 bit length for QRCode symbols
	cons.QRCODE: {{1,1,1,0,1,1,0,0},{0,0,0,1,0,0,0,1}},
	// 4 bits length for Micro QR Code M1 and M3 symbols
	cons.MicroQrcode: {{0,0,0,0},{0,0,0,0}},
}

func (m *AbstractMode) BuildCodewords(qr *QRCodeStruct,dataStream *util.DataStream){
	version := qr.Version
	versionId := version.Id

	// Micro QR Code version M1 and M3 symbols,which is 4 bits in length
	if model.VERSION_M1 == versionId || model.VERSION_M3 == versionId {
		m.buildDataCodewords(qr,dataStream, cons.MicroQrcode, MicroQrcodeM1m3CodewordBit)
	}else{
		//	all codewords are `8 bits in length`
		m.buildDataCodewords(qr,dataStream, cons.QRCODE, QrcodeCodewordBit)
	}
}

func (m *AbstractMode) buildDataCodewords(qr *QRCodeStruct,dataStream *util.DataStream,format cons.Format,codewordBit CodewordBit)  {
	dataStreamLen := dataStream.GetCount()
	codewordsRemain := dataStreamLen % codewordBit
	if codewordsRemain >0 {
		dataStream.AddBit(nil, codewordBit - codewordsRemain)
	}
	version := qr.Version
	ec := qr.ErrorCorrection
	versionDataCapacity := version.GetVersionSymbolCharsAndInputDataCapacity(ec.Level)
	numberOfDataBits := versionDataCapacity.NumberOfDataBits
	dataStreamLen = dataStream.GetCount()
	padCodewordsBitLen := numberOfDataBits - dataStreamLen
	// fill the Pad Codewords
	if padCodewordsBitLen >0{
		//logger.Info("buildDataCodewords fill the Pad Codewords: numberOfDataBits:"+ strconv.Itoa(numberOfDataBits)+
		//	" dataStreamLen:"+strconv.Itoa(dataStreamLen)+", data ramin:"+ strconv.Itoa(padCodewordsBitLen%codewordBit))
		for i:=0;i< padCodewordsBitLen/codewordBit;i++{
			dataStream.AddBit(PadCodewords[format][i%2] ,codewordBit)
		}
	}
}