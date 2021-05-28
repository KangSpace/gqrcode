package mode

import (
	"errors"
	"github.com/gqrcode/core/cons"
	"github.com/gqrcode/util"
	"github.com/gqrcode/util/reedsolomon"
	"strconv"
)

// Define Error Correction here

// ErrorCorrection :QRCode Error Correction(Reed-Solomon codes)
// (referred to as L,M,Q and H increasing order of capacity)
// L 7%
// M 15%
// L 25%
// H 30%
type ErrorCorrection struct {
	Level cons.ErrorCorrectionLevel
	// Name in (L, M, Q, H)
	Name string	`json:"name"`
	// Capacity in 0.07, 0.15, 0.25, 0.30
	RecoveryRate float32 `json:"recoveryRate"`
	RecoveryPercent string `json:"recoveryPercent"`
	rs *reedsolomon.ReedSolomonEncoder
}

// NewErrorCorrection : Create a new error correction
func NewErrorCorrection(level cons.ErrorCorrectionLevel) *ErrorCorrection{
	var ec *ErrorCorrection
	rs := newReedSolomonEncoder()
	switch level {
	case cons.L:
		ec = &ErrorCorrection{cons.L,"L",0.07,"7%",rs}
	case cons.M:
		ec = &ErrorCorrection{cons.M,"M",0.15,"15%",rs}
	case cons.Q:
		ec = &ErrorCorrection{cons.Q,"Q",0.25,"25%",rs}
	case cons.H:
		ec = &ErrorCorrection{cons.H,"H",0.30,"30%",rs}
	case cons.NONE:
		// NONE_EC : None Error Correction for Micro QRCode M1
		ec =  &ErrorCorrection{cons.NONE,"-",0,"0",rs}
	default:
		panic(errors.New("specified level:"+ strconv.Itoa(level)+" is not support"))
	}
	return ec
}

func newReedSolomonEncoder() *reedsolomon.ReedSolomonEncoder{
	return reedsolomon.NewReedSolomonEncoder(reedsolomon.NewGaloisField(285, 256, 0));
}

// CalcECCodewords :Get Error Correction Codewords by Reed-Solomon algorithms
func (ec *ErrorCorrection) CalcECCodewords(data []byte,ecCodewordsCount int) []byte{
	dataIntArray := util.ByteArrayToIntArray(data)
	eccIntArray := ec.rs.Encode(dataIntArray, ecCodewordsCount)
	return util.IntArrayToByteArray(eccIntArray)
}