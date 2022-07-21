package main

import (
	"errors"
	"fmt"
	"github.com/KangSpace/gqrcode/core/cons"
	"github.com/KangSpace/gqrcode/core/logger"
	"github.com/KangSpace/gqrcode/core/mode"
	"github.com/KangSpace/gqrcode/core/model"
)

// NewQRCode0 :
// Two step encode data to QRCode:
// 1. qrcode := NewQRCode()
// 3. qrcode.encode
//
// param content,the input data
// param format, the QRCode Format(common.Format), value in (common.QRCODE,common.QRCODE_MODEL1,common.QrcodeModel2,common.MICRO_QRCODE)
// param ec,the Error Correction(common.ErrorCorrectionLevel), value in (common.L,common.M,common.Q,common.H,common.NONE)
// param m,the encode mode(mode.Mode), value in (mode.NumericMode)
// param quietZoneSize,the quiet zone size for qr code, if zero, then no quiet zone, value in ()
// return: mode.QRCodeStruct
// return: error
func NewQRCode0(content string, format cons.Format, ec *mode.ErrorCorrection, m mode.Mode, quietZone *model.QuietZone) (qr *mode.QRCodeStruct, err error) {
	// panic handle
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
			qr = nil
		}
	}()
	if format == "" {
		// default is QRCode Model2
		format = cons.QrcodeModel2
	}
	da := &DataAnalyzer{content}
	version, ec, m, err := da.analyze(format, ec, m)
	if err != nil {
		return nil, err
	}

	if quietZone == nil {
		quietZone = model.NoneQuietZone
	} else if quietZone == model.AutoQuietZone {
		quietZone = model.NewDefaultQuietZone(version)
	}
	///model.NewQuietZone(version)
	return mode.NewQRCodeStruct(content, format, version, m, ec, quietZone), nil
}

// NewQRCode :Create a QRCode(Model 2).
func NewQRCode(content string) (*mode.QRCodeStruct, error) {
	return NewQRCode0(content, "", nil, nil, nil)
}

// NewQRCodeAutoQuiet :Create a QRCode(Model 2) with model.AutoQuietZone.
func NewQRCodeAutoQuiet(content string) (*mode.QRCodeStruct, error) {
	return NewQRCode0(content, "", nil, nil, model.AutoQuietZone)
}

// NewMicroQRCode :Create a Micro QRCode.
func NewMicroQRCode(content string) (*mode.QRCodeStruct, error) {
	return NewQRCode0(content, cons.MicroQrcode, nil, nil, nil)
}

// NewMicroQRCodeAutoQuiet :Create a Micro QRCode with model.AutoQuietZone.
func NewMicroQRCodeAutoQuiet(content string) (*mode.QRCodeStruct, error) {
	return NewQRCode0(content, cons.MicroQrcode, nil, nil, model.AutoQuietZone)
}

// DataAnalyzer : struct for Data Analysis handle
type DataAnalyzer struct {
	Data string `json:"data"`
}

// Analyze the input data and determine the smallest version.
// This is the first step for general a QRCode.
// Mode select strategy by Page 107(PDF) optimisation of bit stream length.
// param: data, not null
// param: ec, nullable, error correction, if null, choose appropriate ErrorCorrection from H to L with full fill data.
// param: format, nullable, QRCode Format @see model.Format , default is cons.QrcodeModel2
// param: mode, nullable, encode mode, if null, choose appropriate encode by data type.
func (da *DataAnalyzer) analyze(format cons.Format, ec *mode.ErrorCorrection, m mode.Mode) (*model.Version, *mode.ErrorCorrection, mode.Mode, error) {
	// Choose mode
	var err error
	if m == nil {
		if m, err = getMode(da.Data); err != nil {
			return nil, nil, nil, err
		}
	}
	dataLen := len(da.Data)
	var version *model.Version
	var ecLevel cons.ErrorCorrectionLevel
	if ec != nil {
		ecLevel = ec.Level
	}
	// Choose version and error correction level.
	version, ecLevel = model.GetVersionByInputDataLength(format, dataLen, m.GetMode().Name, ecLevel)
	if ec == nil {
		ec = mode.NewErrorCorrection(ecLevel)
	}
	return version, ec, m, nil
}

// 获取支持的内容模式
// param: data
func getMode(data string) (mode.Mode, error) {
	for _, mode := range mode.SupportModes {
		if mode.IsSupport(data) {
			return mode, nil
		}
	}
	return nil, errors.New("please check the input data,can not find a valid Mode for data:" + data)
}
