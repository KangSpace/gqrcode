package mode

import (
	"errors"
	"fmt"
	"github.com/gqrcode/core/cons"
	"github.com/gqrcode/core/model"
	"github.com/gqrcode/core/output"
)

var QRCODE_FORMART = map[string]string{
	"QRCODE":"QRCode",
	"MICRO_QRCODE":"Micro QRCode",
}


// QRCodeStruct :
// contain 3 subtypes:
// 1. QRCode with QRCodeModel1 , QRCodeModel2
// 2. Micro QRCode
type QRCodeStruct struct {
	// Data :the input data
	Data string `json:"data"`
	// QRCode Format: QRCode or Micro QRCode
	Format string `json:"format"`
	Mode Mode                       `json:"mode"`
	Version *model.Version `json:"version"`
	ErrorCorrection *ErrorCorrection `json:"errorCorrection"`
	//FinderPattern *model.FinderPattern `json:"finderPattern"`
	//Separator *model.Separator        `json:"separator"`
	AlignmentPattern *model.AlignmentPattern `json:"alignmentPattern"`
	TimingPattern *model.TimingPattern       `json:"timingPattern"`
	QuietZone *model.QuietZone               `json:"quietZone"`
}

// NewQRCodeStruct :create new QRCodeStruct
// param: data,the input data
// param: format, the QRCode Format(cons.Format), value in (cons.QRCODE,cons.QrcodeModel1,cons.QrcodeModel2,cons.MicroQrcode)
// param: version, the QRCode Version(model.Version), value take by model.NewVersion(model.VersionId),versionId in (model.VERSION1 to model.VERSION40,and model.VERSION_M1 to model.VERSION_M4)
// param: mode,the encode mode(mode.Mode), value in (mode.NumericMode)
// param: ec,the Error Correction(cons.ErrorCorrectionLevel), value in (cons.L,cons.M,cons.Q,cons.H,cons.NONE)
// return: mode.QRCodeStruct
// return: error
func NewQRCodeStruct(data string, format cons.Format, version *model.Version,mode Mode,ec *ErrorCorrection,qz *model.QuietZone) *QRCodeStruct {
	qr := new(QRCodeStruct)
	qr.Data = data
	qr.Format = format
	qr.Version = version
	qr.Mode = mode
	qr.ErrorCorrection = ec
	qr.AlignmentPattern = model.NewAlignmentPattern(version)
	qr.TimingPattern = model.NewTimingPattern()
	qr.QuietZone = qz
	return qr
}

// QRCodeModel1 :QRCode model1
type QRCodeModel1 struct {
	// Format: QRCode model1
	QRCodeStruct `json:"qrCode"`
}
// QRCodeModel2 :QRCode model2
type QRCodeModel2 struct {
	// Format: QRCode model2
	QRCodeStruct `json:"qrCode"`
}

// MicroQRCode :Micro QRCode
type MicroQRCode struct {
	// Format: Micro QRCode
	QRCodeStruct `json:"qrCode"`
}
// GetModuleSize :
func (qr *QRCodeStruct) GetModuleSize() int{
	return qr.Version.GetModuleSize() + qr.QuietZone.GetQuietZoneSize()
}

// Encode :
func (qr *QRCodeStruct) Encode(out output.Output,fileName string) (err error){
	defer func(){
		if rec := recover(); rec != nil {
			switch x := rec.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New(fmt.Sprintf("%v", x))
			}
		}
	}()
	return qr.buildQRCode(out).Save(fileName)
}


// BuildQRCode : Handle QRCode Data step by step , and write data to out.
// Core handle entrance.
// return: output.Output,return the new output by lowestPenaltyOut
func (qr *QRCodeStruct) buildQRCode(out output.Output) output.Output{
	if out.GetBaseOutput().Size == output. AUTO_SIZE{
		out.Init(qr.Version,qr.QuietZone)
	}
	ds := qr.Mode.DataEncode(qr)
	mode := qr.Mode.GetMode()
	// build data bits to codewords
	mode.BuildCodewords(qr,ds)
	// build final message
	finalMessage := mode.BuildFinalErrorCorrectionCodewords(qr,ds)
	// draw image
	return mode.BuildModuleInMatrix(qr,finalMessage,out)
}