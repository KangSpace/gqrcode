package cons

import (
	"bufio"
	"github.com/gqrcode/util"
	"os"
)

// ErrorCorrectionLevel : value in (L, M, Q, H, NONE).
// Micro QRCODE have not ErrorCorrection, so set NONE.
type ErrorCorrectionLevel = int

// ErrorCorrectionLevel constants
const (
	L    ErrorCorrectionLevel = 1
	M    ErrorCorrectionLevel = 2
	Q    ErrorCorrectionLevel = 3
	H    ErrorCorrectionLevel = 4
	NONE ErrorCorrectionLevel = -1   //-1
)

type ModeType = string

const (
	NumericMode          ModeType = "Numeric"
	AlphanumericMode     ModeType = "Alphanumeric"
	ByteMode             ModeType = "Byte"
	EciMode              ModeType = "ECI"
	Fnc1Mode            ModeType = "FNC1"
	Fnc1ModeP1           ModeType = "FNC1_P1"
	Fnc1ModeP2           ModeType = "FNC1_P2"
	KanjiMode            ModeType = "Kanji"
	StructuredAppendMode ModeType = "Structured Append"
)

type Format = string
const (
	QRCODE       Format = "QRCode"
	QrcodeModel1 Format = "QRCode Model1"
	QrcodeModel2 Format = "QRCode Model2"
	MicroQrcode  Format = "Micro QRCode"

)

const (
	// DefaultPixelSizePerModule : default pixel size is 4 pixels for per module(is BaseOutput.Size = AUTO_SIZE)
	DefaultPixelSizePerModule = 4
)

// FormatInformationBitsMap :List of all Format Information Strings
// struct: {ErrorCorrectionLevel:{mask:{info bits array}}}
var FormatInformationBitsMap = map[ErrorCorrectionLevel]map[int][]util.Bit{
	L: {0: {1,1,1,0,1,1,1,1,1,0,0,0,1,0,0},1:{1,1,1,0,0,1,0,1,1,1,1,0,0,1,1},2:{1,1,1,1,1,0,1,1,0,1,0,1,0,1,0},3:{1,1,1,1,0,0,0,1,0,0,1,1,1,0,1},4:{1,1,0,0,1,1,0,0,0,1,0,1,1,1,1},5:{1,1,0,0,0,1,1,0,0,0,1,1,0,0,0},6:{1,1,0,1,1,0,0,0,1,0,0,0,0,0,1},7:{1,1,0,1,0,0,1,0,1,1,1,0,1,1,0}},
	M: {0: {1,0,1,0,1,0,0,0,0,0,1,0,0,1,0},1:{1,0,1,0,0,0,1,0,0,1,0,0,1,0,1},2:{1,0,1,1,1,1,0,0,1,1,1,1,1,0,0},3:{1,0,1,1,0,1,1,0,1,0,0,1,0,1,1},4:{1,0,0,0,1,0,1,1,1,1,1,1,0,0,1},5:{1,0,0,0,0,0,0,1,1,0,0,1,1,1,0},6:{1,0,0,1,1,1,1,1,0,0,1,0,1,1,1},7:{1,0,0,1,0,1,0,1,0,1,0,0,0,0,0}},
	Q: {0: {0,1,1,0,1,0,1,0,1,0,1,1,1,1,1},1:{0,1,1,0,0,0,0,0,1,1,0,1,0,0,0},2:{0,1,1,1,1,1,1,0,0,1,1,0,0,0,1},3:{0,1,1,1,0,1,0,0,0,0,0,0,1,1,0},4:{0,1,0,0,1,0,0,1,0,1,1,0,1,0,0},5:{0,1,0,0,0,0,1,1,0,0,0,0,0,1,1},6:{0,1,0,1,1,1,0,1,1,0,1,1,0,1,0},7:{0,1,0,1,0,1,1,1,1,1,0,1,1,0,1}},
	H: {0: {0,0,1,0,1,1,0,1,0,0,0,1,0,0,1},1:{0,0,1,0,0,1,1,1,0,1,1,1,1,1,0},2:{0,0,1,1,1,0,0,1,1,1,0,0,1,1,1},3:{0,0,1,1,0,0,1,1,1,0,1,0,0,0,0},4:{0,0,0,0,1,1,1,0,1,1,0,0,0,1,0},5:{0,0,0,0,0,1,0,0,1,0,1,0,1,0,1},6:{0,0,0,1,1,0,1,0,0,0,0,1,1,0,0},7:{0,0,0,1,0,0,0,0,0,1,1,1,0,1,1}},
}

// VersionInformationBitsLen :Version Information bit length.
const VersionInformationBitsLen = 18

// VersionInformationBitsMap :List of all Version Information Strings
// struct: {versionId:{info bits array}}
var VersionInformationBitsMap = map[int][]util.Bit{
	7: {0,0,0,1,1,1,1,1,0,0,1,0,0,1,0,1,0,0},
	8: {0,0,1,0,0,0,0,1,0,1,1,0,1,1,1,1,0,0},
	9: {0,0,1,0,0,1,1,0,1,0,1,0,0,1,1,0,0,1},
	10:{0,0,1,0,1,0,0,1,0,0,1,1,0,1,0,0,1,1},
	11:{0,0,1,0,1,1,1,0,1,1,1,1,1,1,0,1,1,0},
	12:{0,0,1,1,0,0,0,1,1,1,0,1,1,0,0,0,1,0},
	13:{0,0,1,1,0,1,1,0,0,0,0,1,0,0,0,1,1,1},
	14:{0,0,1,1,1,0,0,1,1,0,0,0,0,0,1,1,0,1},
	15:{0,0,1,1,1,1,1,0,0,1,0,0,1,0,1,0,0,0},
	16:{0,1,0,0,0,0,1,0,1,1,0,1,1,1,1,0,0,0},
	17:{0,1,0,0,0,1,0,1,0,0,0,1,0,1,1,1,0,1},
	18:{0,1,0,0,1,0,1,0,1,0,0,0,0,1,0,1,1,1},
	19:{0,1,0,0,1,1,0,1,0,1,0,0,1,1,0,0,1,0},
	20:{0,1,0,1,0,0,1,0,0,1,1,0,1,0,0,1,1,0},
	21:{0,1,0,1,0,1,0,1,1,0,1,0,0,0,0,0,1,1},
	22:{0,1,0,1,1,0,1,0,0,0,1,1,0,0,1,0,0,1},
	23:{0,1,0,1,1,1,0,1,1,1,1,1,1,0,1,1,0,0},
	24:{0,1,1,0,0,0,1,1,1,0,1,1,0,0,0,1,0,0},
	25:{0,1,1,0,0,1,0,0,0,1,1,1,1,0,0,0,0,1},
	26:{0,1,1,0,1,0,1,1,1,1,1,0,1,0,1,0,1,1},
	27:{0,1,1,0,1,1,0,0,0,0,1,0,0,0,1,1,1,0},
	28:{0,1,1,1,0,0,1,1,0,0,0,0,0,1,1,0,1,0},
	29:{0,1,1,1,0,1,0,0,1,1,0,0,1,1,1,1,1,1},
	30:{0,1,1,1,1,0,1,1,0,1,0,1,1,1,0,1,0,1},
	31:{0,1,1,1,1,1,0,0,1,0,0,1,0,1,0,0,0,0},
	32:{1,0,0,0,0,0,1,0,0,1,1,1,0,1,0,1,0,1},
	33:{1,0,0,0,0,1,0,1,1,0,1,1,1,1,0,0,0,0},
	34:{1,0,0,0,1,0,1,0,0,0,1,0,1,1,1,0,1,0},
	35:{1,0,0,0,1,1,0,1,1,1,1,0,0,1,1,1,1,1},
	36:{1,0,0,1,0,0,1,0,1,1,0,0,0,0,1,0,1,1},
	37:{1,0,0,1,0,1,0,1,0,0,0,0,1,0,1,1,1,0},
	38:{1,0,0,1,1,0,1,0,1,0,0,1,1,0,0,1,0,0},
	39:{1,0,0,1,1,1,0,1,0,1,0,1,0,0,0,0,0,1},
	40:{1,0,1,0,0,0,1,1,0,0,0,1,1,0,1,0,0,1},
}

var GQRcodeVersion = getGQRCodeVersion

func getGQRCodeVersion() string{
	pwd ,_:=os.Getwd()
	versionFile := pwd +"/../../gqrcode.version"
	if file,err := os.Open(versionFile); err == nil{
		version,_,_:= bufio.NewReader(file).ReadLine()
		return string(version)
	}
	return ""
}