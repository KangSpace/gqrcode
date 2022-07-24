package mode

import (
	"errors"
	"fmt"
	"github.com/KangSpace/gqrcode/core/cons"
	"github.com/KangSpace/gqrcode/core/logger"
	"github.com/KangSpace/gqrcode/core/model"
	"github.com/KangSpace/gqrcode/core/output"
	"io"
)

var QrcodeFormart = map[string]string{
	"QRCODE":       "QRCode",
	"MICRO_QRCODE": "Micro QRCode",
}

// QRCodeStruct :
// contain 3 subtypes:
// 1. QRCode with QRCodeModel1 , QRCodeModel2
// 2. Micro QRCode
type QRCodeStruct struct {
	// Data :the input data
	Data string `json:"data"`
	// QRCode Format: QRCode or Micro QRCode
	Format          string           `json:"format"`
	Mode            Mode             `json:"mode"`
	Version         *model.Version   `json:"version"`
	ErrorCorrection *ErrorCorrection `json:"errorCorrection"`
	//FinderPattern *model.FinderPattern `json:"finderPattern"`
	//Separator *model.Separator        `json:"separator"`
	AlignmentPattern *model.AlignmentPattern `json:"alignmentPattern"`
	TimingPattern    *model.TimingPattern    `json:"timingPattern"`
	QuietZone        *model.QuietZone        `json:"quietZone"`
	Mask             int                     `json:"mask"`
}

// NewQRCodeStruct :create new QRCodeStruct
// param: data,the input data
// param: format, the QRCode Format(cons.Format), value in (cons.QRCODE,cons.QrcodeModel1,cons.QrcodeModel2,cons.MicroQrcode)
// param: version, the QRCode Version(model.Version), value take by model.NewVersion(model.VersionId),versionId in (model.VERSION1 to model.VERSION40,and model.VersionM1 to model.VersionM4)
// param: mode,the encode mode(mode.Mode), value in (mode.NumericMode)
// param: ec,the Error Correction(cons.ErrorCorrectionLevel), value in (cons.L,cons.M,cons.Q,cons.H,cons.NONE)
// return: mode.QRCodeStruct
// return: error
func NewQRCodeStruct(data string, format cons.Format, version *model.Version, mode Mode, ec *ErrorCorrection, qz *model.QuietZone) *QRCodeStruct {
	qr := new(QRCodeStruct)
	qr.Data = data
	qr.Format = format
	qr.Version = version
	qr.Mode = mode
	qr.ErrorCorrection = ec
	qr.AlignmentPattern = model.NewAlignmentPattern(version)
	qr.TimingPattern = model.NewTimingPattern(version)
	qr.QuietZone = qz
	return qr
}

// resetQRCodeStruct :create new QRCodeStruct from QRCodeStruct
// param: data,the input data
// param: format, the QRCode Format(cons.Format), value in (cons.QRCODE,cons.QrcodeModel1,cons.QrcodeModel2,cons.MicroQrcode)
// param: version, the QRCode Version(model.Version), value take by model.NewVersion(model.VersionId),versionId in (model.VERSION1 to model.VERSION40,and model.VersionM1 to model.VersionM4)
// param: mode,the encode mode(mode.Mode), value in (mode.NumericMode)
// param: ec,the Error Correction(cons.ErrorCorrectionLevel), value in (cons.L,cons.M,cons.Q,cons.H,cons.NONE)
// return: mode.QRCodeStruct
// return: error
func resetQRCodeStruct(qr *QRCodeStruct) *QRCodeStruct {
	qr.AlignmentPattern = model.NewAlignmentPattern(qr.Version)
	qr.TimingPattern = model.NewTimingPattern(qr.Version)
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
func (qr *QRCodeStruct) GetModuleSize() int {
	return qr.Version.GetModuleSize() + qr.QuietZone.GetQuietZoneSize()
}

// GetMaskCount :
func (qr *QRCodeStruct) GetMaskCount() int {
	if qr.Version.Id > 0 {
		return 8
	}
	return 4
}

// GetBitLen :
func (qr *QRCodeStruct) GetBitLen() int {
	if qr.Version.Id == model.VersionM1 || qr.Version.Id == model.VersionM3 {
		return 4
	}
	return 8
}

// SetMask :
func (qr *QRCodeStruct) SetMask(mask int) {
	qr.Mask = mask
}

// Encode :
func (qr *QRCodeStruct) Encode(out output.Output, fileName string) (err error) {
	if out, err := qr.innerEncode(out); err == nil {
		return out.Save(fileName)
	}
	return err
}

// EncodeToWriter :
func (qr *QRCodeStruct) EncodeToWriter(out output.Output, writer io.Writer) (err error) {
	if out, err = qr.innerEncode(out); err == nil {
		return out.SaveToWriter(writer)
	}
	return err
}

// EncodeToBase64 :
func (qr *QRCodeStruct) EncodeToBase64(out output.Output) (base64Str string, err error) {
	if out, err := qr.innerEncode(out); err == nil {
		return out.SaveToBase64()
	}
	return "", err
}

func (qr *QRCodeStruct) innerEncode(out output.Output) (out_ output.Output, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			switch x := rec.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New(fmt.Sprintf("%v", x))
			}
			logger.Error(err)
		}
	}()
	return qr.buildQRCode(out), nil
}

// BuildQRCode : Handle QRCode Data step by step , and write data to out.
// Core handle entrance.
// return: output.Output,return the new output by lowestPenaltyOut
func (qr *QRCodeStruct) buildQRCode(out output.Output) output.Output {
	if out.GetBaseOutput().Size == output.AUTO_SIZE {
		out.Init(qr.Version, qr.QuietZone)
	} else {
		//qr.resetQRCodeVersionByInputData(out.GetBaseOutput().Size)
	}
	ds := qr.Mode.DataEncode(qr)
	mode := qr.Mode.GetMode()
	// build data bits to codewords
	mode.BuildCodewords(qr, ds)
	// build final message
	finalMessage := mode.BuildFinalErrorCorrectionCodewords(qr, ds)
	// draw image
	return mode.BuildModuleInMatrix(qr, finalMessage, out)
}

// resetQRCodeVersionByInputData : reset version and mode by input data on specific qrcode size
func (qr *QRCodeStruct) resetQRCodeVersionByInputData(qrcodeSize int) {
	// 若存在指定大小, 则取与大小成比例的版本和错误级别
	version, m, ecLevel := GetBestSizeVersionByInputDataLength(qr.Data, qr.QuietZone, qrcodeSize, qr.Format, qr.Version.Id)
	qr.Version = version
	qr.Mode = m
	if ecLevel != qr.ErrorCorrection.Level {
		qr.ErrorCorrection = NewErrorCorrection(ecLevel)
	}
	resetQRCodeStruct(qr)
}

// GetBestSizeVersionByInputDataLength : Get best qrcode size for version by input data.
func GetBestSizeVersionByInputDataLength(data string, quietZone *model.QuietZone, qrcodeSize int, format cons.Format, srcVId model.VersionId) (*model.Version, Mode, cons.ErrorCorrectionLevel) {
	maxVId := model.VERSION40
	//minVId := model.VERSION1
	// set vid start value
	minVId := srcVId
	condition := func(vId model.VersionId) bool {
		return vId <= maxVId
	}
	vIdUpdate := func(vId model.VersionId) int {
		vId++
		return vId
	}
	if format == cons.MicroQrcode {
		//minVId = model.VersionM1
		maxVId = model.VersionM4
		condition = func(vId model.VersionId) bool {
			return vId >= maxVId
		}
		vIdUpdate = func(vId model.VersionId) int {
			vId--
			return vId
		}
	}
	dataLen := len(data)
	// fetch all supported modes
	modes, err := GetSupportedModes(data)
	if err != nil {
		panic(err)
	}
	var candidateVersion *model.Version
	var candidateMode Mode
	var candidateEcLevel cons.ErrorCorrectionLevel
versionCicle:
	for vId := minVId; condition(vId); vId = vIdUpdate(vId) {
		if ecMap, ok := model.VersionSymbolCharsAndInputDataCapacityMap[vId]; ok {
			for ecLevel := cons.H; ecLevel > 0; ecLevel-- {
				if capacity, ok := ecMap[ecLevel]; ok {
					for _, modeObj := range modes {
						mode_ := (modeObj).GetMode().Name
						if size, ok := capacity.DataCapacity[mode_]; ok && size >= dataLen {
							tempVersion := model.GetVersion(vId)
							totalModuleSize := tempVersion.GetTotalModuleSize(quietZone)
							isSizeMatch := (qrcodeSize+totalModuleSize)%totalModuleSize == 0
							if candidateVersion == nil || isSizeMatch {
								candidateVersion = tempVersion
								candidateEcLevel = ecLevel
								candidateMode = modeObj
								// use min version for qrcode
								if isSizeMatch {
									break versionCicle
								}
							}
						}
					}
				}
			}
		}
	}
	if candidateVersion == nil {
		panic(errors.New("can not find DataCapacity for data:" + data))
	}
	return candidateVersion, candidateMode, candidateEcLevel
}

func (qr *QRCodeStruct) IsMicroQRCode() bool {
	return cons.MicroQrcode == qr.Format
}
func (qr *QRCodeStruct) IsQRCodeModel2() bool {
	return cons.QrcodeModel1 == qr.Format
}
func (qr *QRCodeStruct) IsQRCodeModel1() bool {
	return cons.QrcodeModel2 == qr.Format
}

// Decode :
func (qr *QRCodeStruct) Decode(fileName string) (err error) {
	// TODO
	return err
}
