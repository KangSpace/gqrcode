package mode

import (
	"github.com/gqrcode/util"
)

// Define Error Correction Block here


// ECBlock :Error correction block unit
type ECBlock struct {
	// per block data codewords
	DataCodewords []byte
	// per block error correction codewords
	ErrorCorrectionCodewords []byte
}

// ECBlockList :ECBlock array
type ECBlockList = []*ECBlock

// ConvertDataBitsToECBlocks :convert the data stream bits to ECBlockList to build final message.
// Page 52,7.5.2 Generating the error correction codewords
// struct: {ECBlockList{ DataCodewords:{1,1,1,1},ErrorCorrectionCodewords:{1,1,1,1}}},....}
func (m *AbstractMode) ConvertDataBitsToECBlocks(qr *QRCodeStruct,dataStream *util.DataStream) ECBlockList {
	byteOutChan := dataStream.IteratorByte()
	ec := qr.ErrorCorrection
	versionSymbolCharsAndInputDataCapacity := qr.Version.GetVersionSymbolCharsAndInputDataCapacity(ec.Level)
	ecBlockCapacity:= versionSymbolCharsAndInputDataCapacity.ErrorCorrectionBlockCapacity
	ecBlocks := make(ECBlockList, ecBlockCapacity.GetTotalECBlocksCount())
	g1ECBlockCount := ecBlockCapacity.NoECBlocksG1
	// group 1 error block codeword handle
	for g1i := 0; g1i < g1ECBlockCount; g1i++ {
		ecb := new(ECBlock)
		ecb.DataCodewords = make([]byte, ecBlockCapacity.NoDataCodewordsPerBlockG1)
		for cw := 0; cw < ecBlockCapacity.NoDataCodewordsPerBlockG1; cw++ {
			ecb.DataCodewords[cw] = <-byteOutChan
		}
		ecb.ErrorCorrectionCodewords = ec.CalcECCodewords(ecb.DataCodewords,ecBlockCapacity.NoECCodewordsPerBlock)
		ecBlocks[g1i] = ecb
	}
	// group 2 error block codeword handle
	for g2i := 0; g2i < ecBlockCapacity.NoECBlocksG2; g2i++ {
		ecb := new(ECBlock)
		ecb.DataCodewords = make([]byte, ecBlockCapacity.NoDataCodewordsPerBlockG2)
		for cw := 0; cw < ecBlockCapacity.NoDataCodewordsPerBlockG2; cw++ {
			ecb.DataCodewords[cw] = <-byteOutChan
		}
		ecb.ErrorCorrectionCodewords = ec.CalcECCodewords(ecb.DataCodewords,ecBlockCapacity.NoECCodewordsPerBlock)
		ecBlocks[g1ECBlockCount +g2i] = ecb
	}
	return ecBlocks
}

// InterleaveECBlocks : interleave the ECBlocks.
// Page 53,7.6 Constructing the final message codeword sequence.
// return the final bits array(1 bit in byte).
func (m *AbstractMode) InterleaveECBlocks(qr *QRCodeStruct,ecBlocks ECBlockList) []util.Bit{
	versionFinalCodewordCapacity := qr.Version.GetVersionFinalCodewordCapacity()
	ec := qr.ErrorCorrection

	vscaidc := qr.Version.GetVersionSymbolCharsAndInputDataCapacity(ec.Level)
	finalDataCodewordCount := versionFinalCodewordCapacity.DataCapacityCodewords
	reminderBitsLen := versionFinalCodewordCapacity.RemainderBits
	finalCodewords := make([]byte, 0, finalDataCodewordCount)

	maxPerBlockCodewordCount := 0
	if vscaidc.ErrorCorrectionBlockCapacity.NoDataCodewordsPerBlockG1 >
		vscaidc.ErrorCorrectionBlockCapacity.NoDataCodewordsPerBlockG2{
		maxPerBlockCodewordCount = vscaidc.ErrorCorrectionBlockCapacity.NoDataCodewordsPerBlockG1
	} else{
		maxPerBlockCodewordCount = vscaidc.ErrorCorrectionBlockCapacity.NoDataCodewordsPerBlockG2
	}

	rowsCount := len(ecBlocks)
	// build data codewords
	// Page 53,7.6 Constructing the final message codeword sequence
	for col:=0 ; col< maxPerBlockCodewordCount; col++{
		for row:=0 ; row <rowsCount ; row ++{
			if len(ecBlocks[row].DataCodewords) > col{
				finalCodewords = append(finalCodewords,ecBlocks[row].DataCodewords[col])
			}
		}
	}
	eccPerBlockCount := vscaidc.ErrorCorrectionBlockCapacity.NoECCodewordsPerBlock
	// build error correction codewords
	for col:=0 ; col< eccPerBlockCount; col++{
		for row:=0 ; row <rowsCount ; row ++{
			finalCodewords = append(finalCodewords,ecBlocks[row].ErrorCorrectionCodewords[col])
		}
	}

	finalBitsLen := len(finalCodewords) * 8 + reminderBitsLen
	// convert codewords to bit array and add Remainder bits
	return util.ByteArrayTo8BitArrayWithCount(finalCodewords,finalBitsLen)
}

